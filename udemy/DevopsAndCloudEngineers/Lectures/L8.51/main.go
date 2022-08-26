package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var (
		client *kubernetes.Clientset
		err    error
	)
	ctx := context.Background()
	if client, err = getClient(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	deploymentLabels, expectedReplicas, err := deploy(ctx, client)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	err = waitForPods(ctx, client, deploymentLabels, expectedReplicas)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("deployment completed %+v\n", deploymentLabels)
}

func getClient() (*kubernetes.Clientset, error) {
	var kubeconfig *string
	kubeconfig = flag.String("kubeconfig", filepath.Join(homedir.HomeDir(), ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return nil, err
	}

	// create the clientset
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientSet, nil
}

func deploy(ctx context.Context, client *kubernetes.Clientset) (map[string]string, int32, error) {
	var deployment *v1.Deployment

	appFile, err := ioutil.ReadFile("app.yaml")
	if err != nil {
		return nil, 0,fmt.Errorf("error reading yaml file: %v", err)
	}

	obj, groupVersionKind, err := scheme.Codecs.UniversalDeserializer().Decode(appFile, nil, nil)
	switch obj.(type) {
	case *v1.Deployment:
		deployment = obj.(*v1.Deployment)
	default:
		return nil, 0, fmt.Errorf("unexpected deployment type: %v", groupVersionKind)
	}

	_, err = client.AppsV1().Deployments("default").Get(ctx, "helloworld-deployment", metav1.GetOptions{})
	if err != nil && errors.IsNotFound(err) {
		deploymentResponse, err := client.AppsV1().Deployments("default").Create(ctx, deployment, metav1.CreateOptions{})
		if err != nil {
			return nil, 0, fmt.Errorf("error creating deployment: %v", err)
		}
		return deploymentResponse.Spec.Template.Labels, 0, nil
	} else if err != nil && !errors.IsNotFound(err) {
		return nil, 0, fmt.Errorf("error getting deployment: %v", err)
	}

	deploymentResponse, err := client.AppsV1().Deployments("default").Update(ctx, deployment, metav1.UpdateOptions{})
	if err != nil {
		return nil, 0, fmt.Errorf("error creating deployment: %v", err)
	}
	return deploymentResponse.Spec.Template.Labels, *deploymentResponse.Spec.Replicas, nil
}

func waitForPods(ctx context.Context, client *kubernetes.Clientset, deploymentLabels map[string]string, expectedReplicas int32) error {
	for {
		validatedLabels, err := labels.ValidatedSelectorFromSet(deploymentLabels)
		if err != nil {
			return fmt.Errorf("ValidatedSelectorFromSet error: %v", err)
		}

		podList, err := client.CoreV1().Pods("default").List(ctx, metav1.ListOptions{
			LabelSelector: validatedLabels.String(),
		})
		if err != nil {
			return fmt.Errorf("error listing pods: %v", err)
		}

		podsRunning := 0
		for _, pod := range podList.Items {
			if pod.Status.Phase == "Running" {
				podsRunning++
			}
		}
		fmt.Printf("Waiting for pods to become ready (running %d / %d)\n", podsRunning, len(podList.Items))
		if podsRunning > 0 && podsRunning == len(podList.Items) && podsRunning == int(expectedReplicas) {
			break
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
