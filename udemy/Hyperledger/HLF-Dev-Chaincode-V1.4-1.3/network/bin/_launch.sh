#PEER_MODE=dev
#Command=dev-init.sh -d 
#Generated: Mon Jul 15 09:12:17 UTC 2019 
docker-compose  -f ./compose/docker-compose.base.yaml    -f ./compose/docker-compose.dev.yaml      up -d --remove-orphans
