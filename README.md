# Pod Info

## Description

Simple Go web app which returns some env variables from pod

## Build Docker Image

```bash
docker build -t andriktr/pod-info .
docker push andriktr/pod-info
```

## Deploy to kubernetes

```bash
kubectl apply -f pod-info.yaml --namespace=<<your namespace>>
```

## Test the app

