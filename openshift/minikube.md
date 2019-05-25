# SSH Port forward minikube

### Machines:

1. minikube
1. client

### Setup

1. install minikube
1. copy the ~/.kube/config to `client`
1. edit the `~/.kube/config` to setup the following properly
  - clusters
  - context
  - users
  - current-context
1. Ensure you copy the `.minikube/client.*` and `ca.crt` from `minikube`
1. setup ssh port-forward
```
Host minikube
  Hostname <host-running-minikube>
  LocalForward <local-port> <minikube ip>:<minikube-port>
  # ... rest of the ssh config
```


