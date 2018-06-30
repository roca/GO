kubectl patch deployment camera -p "{\"spec\":{\"template\":{\"metadata\":{\"labels\":{\"date\":\"`date +'%s'`\"}}}}}"

Rebuild master:

    1  curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
    2  echo "deb http://apt.kubernetes.io/ kubernetes-xenial main" > /etc/apt/sources.list.d/kubernetes.list
    3  apt-get update && apt-get install -y kubelet=1.7.1-00 kubeadm=1.7.1-00 kubectl=1.7.1-00 kubernetes-cni=0.5.1-00 --force-yes
    4  kubeadm reset
    5  kubeadm init --pod-network-cidr 10.244.0.0/16 --apiserver-advertise-address=192.168.1.149


    83  sudo apt-get install weavedconnectd
   84  sudo weavedinstaller