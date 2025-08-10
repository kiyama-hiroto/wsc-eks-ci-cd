#!/bin/bash

# Acesso a variável de ambiente IMAGE_URI que é passada pelo CodeDeploy
IMAGE_URI=$(cat imageDetail.json | jq -r .ImageURI)

# Substitui o placeholder no app.yaml com a nova URI da imagem
sed -i "s|APP_IMAGE|$IMAGE_URI|g" conf/app.yaml

# Aplica as configurações no cluster EKS
kubectl apply -f conf/app.yaml
kubectl apply -f conf/ingress.yaml
kubectl apply -f conf/service.yaml