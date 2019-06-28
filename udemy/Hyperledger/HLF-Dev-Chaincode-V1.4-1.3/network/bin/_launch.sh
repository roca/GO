#PEER_MODE=net
#Command=dev-init.sh -e 
#Generated: Fri Jun 28 08:55:24 UTC 2019 
docker-compose  -f ./compose/docker-compose.base.yaml      -f ./compose/docker-compose.explorer.yaml    up -d --remove-orphans
