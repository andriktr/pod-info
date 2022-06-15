# Pod Info

## Description

Simple Go web app which returns some env variables from pod

## Build Docker Image

```bash
docker build --platform linux/amd64 -t andriktr/pod-info .     
docker push andriktr/pod-info
```

## Deploy to kubernetes

```bash
kubectl apply -f pod-info.yaml --namespace=<<your namespace>>
```

## Test the app

```bash
kubectl port-forward --namespace=<<your namespace>> infra services/pod-info 8080:80
curl http://localhost:8080
```