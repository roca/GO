kubectl patch deployment camera -p "{\"spec\":{\"template\":{\"metadata\":{\"labels\":{\"date\":\"`date +'%s'`\"}}}}}"

Rebuild master:

    1  curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
    2  echo "deb http://apt.kubernetes.io/ kubernetes-xenial main" > /etc/apt/sources.list.d/kubernetes.list
    3  apt-get update && apt-get install -y kubelet=1.7.1-00 kubeadm=1.7.1-00 kubectl=1.7.1-00 kubernetes-cni=0.5.1-00 --force-yes
    or  apt-get update && apt-get install -y kubelet=1.7.0-00 kubeadm=1.7.0-00 kubectl=1.7.0-00 kubernetes-cni=0.5.1-00 --force-yes
    4  kubeadm reset
    5  kubeadm init --pod-network-cidr 10.244.0.0/16 --apiserver-advertise-address=192.168.1.100A

    as non-root

   mkdir -p $HOME/.kube
   sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
   sudo chown $(id -u):$(id -g) $HOME/.kube/config
   export KUBECONFIG=$HOME/.kube/config
   echo 'export KUBECONFIG=$HOME/.kube/config' >> ~/.profile

    6 curl -sSL https://rawgit.com/coreos/flannel/v0.9.1/Documentation/kube-flannel.yml | sed "s/amd64/arm/g" | kubectl create -f -
    7 kubectl label node master nginx-controller=traefik
    8 kubectl apply -f https://raw.githubusercontent.com/hypriot/rpi-traefik/master/traefik-k8s-example.yaml


    83  sudo apt-get install weavedconnectd
   84  sudo weavedinstaller