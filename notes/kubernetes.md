## Kubernetes Arch
Pods which contain multiple containers
container has a kubenretes-proxy you can interact with
and they have a kublet from minion to master

## Services?
key, value pair to name pods but service is abstraction for group of pods
dns name and ip lookup
service proxies thing to service behind it

replica sets ensures required number are running
  - manages and makes sure pods are maintained
  - provides elasticity

deployment

## Kubernetes locally
minikube
`curl -sS https://get.k8s.io | bash`

## Get Minikube locally
`curl -Lo minikube https://storage.googleapis.com/minikube/releases/v1.0.1/minikube-darwin-amd64 && chmod +x minikube && sudo cp minikube /usr/local/bin/ && rm minikube`
found here: https://github.com/kubernetes/minikube

## Runing locally
once installed, make sure to run minikube to test locally. `minkube start`

# Health Checks
readinessProbe and livenessProbe handles the health checking and status updates for a port.

