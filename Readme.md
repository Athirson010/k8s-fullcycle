# Comandos

## Criação do cluster
-- kind create cluster --name=gitopsfc
-- kubectl cluster-info --context kind-gitopsfc

- kubectl apply -f .\k8s\gitopsfc\
- kubectl get deploy
- kubectl get service


## Kustomize
-- kustomize edit 

## Instalando ARGO CD

-- kubectl create namespace argocd
-- kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
-- kubectl get all -n argocd

-- kubectl get secret argocd-initial-admin-secret -n argocd -o jsonpath="{.data.password}" (é necessario decodificar a senha em base64)