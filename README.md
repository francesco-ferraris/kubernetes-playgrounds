# Description
This repository contains three playgrounds that show how you can create a local Kubernetes cluster with
[kind](https://kind.sigs.k8s.io/) and deploy applications and helm charts.
Installation instruction are provided for macOS devices.

## Prerequisites
- Docker
The first playground contains `Makefile` directives to install kubectl and kind CLI

## Playgrounds
Use the `Makefile` directives in each playground to setup the environment.
These are the topics covered in the playgrounds:
- Create a local Kind cluster and expose a stateless application managed by a Deployment
- Deploy and customize the same web application using a Helm Chart.
