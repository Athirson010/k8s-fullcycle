# Comandos

#Criação do cluster
-- kind create cluster --name=gitopsfc
-- kubectl cluster-info --context kind-gitopsfc

kubectl apply -f .\k8s\gitopsfc\
kubectl get deploy
kubectl get service


#Kustomize
-- kustomize edit 
