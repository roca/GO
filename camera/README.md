kubectl patch deployment camera -p "{\"spec\":{\"template\":{\"metadata\":{\"labels\":{\"date\":\"`date +'%s'`\"}}}}}"

Rebuild master:

flash --hostname master https://github.com/hypriot/image-builder-rpi/releases/download/v1.9.0/hypriotos-rpi-v1.9.0.img.zip
edit /boot/user-data as needed

Go edit /boot/cmdline.txt with your favorite editor, or use
sudo nano /boot/cmdline
and add this at the very end. Don't press enter.
cgroup_enable=cpuset cgroup_enable=memory

As root:

        apt-get update && apt-get install -y apt-transport-https curl
        curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
        cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
        deb http://apt.kubernetes.io/ kubernetes-xenial main
        EOF
        apt-get update
        apt-get install -y kubelet kubeadm kubectl

        on master: swapoff -a && kubeadm init --pod-network-cidr 10.244.0.0/16 --apiserver-advertise-address=192.168.1.100

As non-root

   mkdir -p $HOME/.kube
   sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
   sudo chown $(id -u):$(id -g) $HOME/.kube/config
   export KUBECONFIG=$HOME/.kube/config
   echo 'export KUBECONFIG=$HOME/.kube/config' >> ~/.profile

    kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
    kubectl label node master nginx-controller=traefik
    kubectl apply -f https://raw.githubusercontent.com/hypriot/rpi-traefik/master/traefik-k8s-example.yaml


    sudo apt-get install weavedconnectd
    sudo weavedinstaller


   

