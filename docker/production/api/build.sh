#!/usr/bin/env bash

origin=772010606571.dkr.ecr.ap-northeast-1.amazonaws.com
target=learnmemo
ecr=$origin/production-$target

aws ecr get-login-password | docker login --username AWS --password-stdin https://$origin
docker build -t $target -f Dockerfile ../../../api
docker tag $target:latest $ecr:latest
docker push $ecr:latest
