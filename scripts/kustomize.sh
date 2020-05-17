#!/usr/bin/env bash

env

set -x
NS=test

echo "Deploying micro-starter-kit..."

# Set the image tag to the sha hash that we just built in the previous stage.
# K8S will do a rolling deployment
(   cd deploy/overlays/production;  \
    kustomize edit set docker.pkg.github.com/seidu626/audiobook/greeter-service:"$SHORT_SHA" && \
    kustomize edit set namespace $NS  && \
    kustomize build | \
    kubectl apply -f -  )
