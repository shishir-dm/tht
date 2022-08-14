## High Level SaaS architecture

In this document we will describe in a high-level some example architectures for a SaaS solution built on Kubernetes with two individual services, 'service-a' and 'service-b' and a distributed database.

This repository is meant to serve as a bucket for multiple example repositories such as
- [repo_microservices](repo_microservices): Example repository that holds the source code of our services (Go) and the deployment infrastructure code of our services (Helm, Docker, docker-compose)
- [repo_argo-config](repo_argo-config): Example repository that holds the ArgoCD configuration that describes:
  - What helm chart to deploy (version, name etc)
  - Where to deploy this helm chart (which kubernetes clsuter)
  - How to deploy this helm chart (what values/config to use)
- [repo_service-a](repo_service-a): Example repository that holds the source and deployment infrastructure code of service-a (meant to be just a skeleton for developing service-a seperately from service-b)
- [repo_service-b](repo_service-b): Example repository that holds the source and deployment infrastructure code of service-b (meant to be just a skeleton for developing service-a seperately from service-a)

To start the presentation of these examples, please go in the following order
1. [High Level Architecture](presentation/architectures.md)
2. [Deployments](presentation/deployments.md)
3. [Developer Setup](presentation/developer.md)
4. [Compliance](presentation/compliance.md)
