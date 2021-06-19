#!/bin/bash

# script usage
# run_project.sh --app_name provision-envs --clean_docker true
app_name=${app_name:-provision-envs}
clean_docker=${clean_docker:-false}

while [ $# -gt 0 ]; do
   if [[ $1 == *"--"* ]]; then
        param="${1/--/}"
        declare $param="$2"
        # echo $1 $2 // Optional to see the parameter:value result
   fi
  shift
done

if [ $clean_docker = "true" ]; then
    echo "Stopping and deleting containers from $app_name iamge ..."
    docker ps -a | awk '{ print $1,$2 }' | grep $app_name | awk '{print $1 }' | xargs -I {} docker stop {} | xargs -I {} docker rm {}
    echo "Deleting unused docker images ..."
    docker image prune -a -f
fi

echo "Triggering image $app_name creation"
docker build --build-arg APP_NAME=$app_name -t provision-envs . -f deploy/Dockerfile
docker run -it --rm --name $app_name $app_name
# to create a container that stays and never finishes
# docker run -d provision-envs tail -f /dev/null