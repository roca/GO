#PEER_MODE=dev
#Command=dev-init.sh -d 
#Generated: Wed Jul 10 08:05:24 UTC 2019 
docker-compose  -f ./compose/docker-compose.base.yaml    -f ./compose/docker-compose.dev.yaml      up -d --remove-orphans
