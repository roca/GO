#PEER_MODE=dev
#Command=dev-init.sh -d 
#Generated: Sun Jul 14 09:24:52 UTC 2019 
docker-compose  -f ./compose/docker-compose.base.yaml    -f ./compose/docker-compose.dev.yaml      up -d --remove-orphans
