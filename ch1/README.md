# Description
Create a local Kubernetes Kind cluster 
with a nginx reverse-proxy to expose the applications
and deploy a simple hello world application.

## Cluster creation

```makefile
make requirements
```
Install kind and kubectl requirements using `brew`
<br/>
<br/>
```makefile
make create
```
Create a local Kubernetes cluster with kind and configure the kubectl CLI to communicate with it.
<br/>
<br/>
```makefile
make connect
```
Update the kubeconfig context to the cluster. If you have created the cluster with `make create`,
the context is already updated
<br/>

### Test the cluster connection
Now that you have a running Kubernetes cluster you can try to deploy a simple nginx pod:
```bash
# Create and run a Pod with the 'nginx' latest image from DockerHub in the default namespace
kubectl run web --image nginx
```
To inspect the state of the pod or watch the logs you can run these commands:
```bash
# Print information in human readable text describing the configuration and status of the Pod
kubectl describe pod web
# Print the declarative configuration of the Pod with YAML output
kubectl get pod web -o yaml
# Watch the logs of the pod. The Pod must be in Running state
kubectl logs web -f
```

To delete the test pod run:
```bash
# Delete only the pod 
kubectl delete pod web 
```

## Build and Deploy
Deploy and expose outside the cluster a simple hello world application that greets people specified in the URL path.
This playground assumes that you have configured a local kubernetes cluster with ingress nginx controller configured to
accept Ingresses objects.

```makefile
make build
```
Build the application to a docker image
<br/>
<br/>

```makefile
make deploy
```
Deploy the application to the kubernetes cluster

## Cleanup
```makefile
make delete
```
Delete the application from the cluster
<br/>
<br/>

```makefile
make delete-cluster
```
Delete the cluster
