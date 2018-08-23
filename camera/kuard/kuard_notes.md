
    KUARD page 57

 kubectl run alpaca-prod         --image=gcr.io/kuar-demo/kuard-arm:1 --replicas=2 --labels="ver=1,app=alpaca,env=prod"
 kubectl run alpaca-test        --image=gcr.io/kuar-demo/kuard-arm:2 --replicas=1 --labels="ver=2,app=alpaca,env=test"
 kubectl run bandicoot-prod     --image=gcr.io/kuar-demo/kuard-arm:2 --replicas=2 --labels="ver=2,app=bandicoot,env=prod"
 kubectl run bandicoot-staging  --image=gcr.io/kuar-demo/kuard-arm:2 --replicas=1 --labels="ver=2,app=bandicoot,env=staging"



Page 67
 kubectl run alpaca-prod         --image=gcr.io/kuar-demo/kuard-arm:1 --replicas=3 --port=8080 --labels="ver=1,app=alpaca,env=prod"
 kubectl expose deployment alpaca-prod
 kubectl run bandicoot-prod     --image=gcr.io/kuar-demo/kuard-arm:2 --replicas=2 --port=8080 --labels="ver=2,app=bandicoot,env=prod"
 kubectl expose deployment bandicoot-prod


 Page 96

 kubectl run -i oneshot --image=gcr.io/kuar-demo/kuard-arm:1 --restart=OnFailure -- --keygen-enable --keygen-enable --keygen-exit-on-complete --keygen-num-to-gen 10
   
Page 104

    curl -X PUT localhost:8080/memq/server/queues/keygen

    for i in work-item-{0..99}; do
    curl -X POST localhost:8080/memq/server/queues/keygen/enqueue -d "$i"
    done

Page 108

kubectl create configmap my-config --from-file=my-config.txt --from-literal=extra-parm=extra-value --from-literal=another-parm=another-value


Page 112

curl -O https://storage.googleapis.com/kuar-demo/kuard.crt
curl -O https://storage.googleapis.com/kuar-demo/kuard.key

kubectl create secret generic kuard-tls --from-file=kuard.crt --from-file=kuard.key
kubectl create secret docker-registry my-image-pull-secret --docker-username=<username> --docker-password=<password> --docker-email=<email-address>