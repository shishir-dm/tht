## ArgoCD Configuration Repo

This repository holds all the configuration that instructs ArgoCD what, where and how to deploy a helm chart.

For more details see [Deployments](../presentation/deployments.md)

This repo has the following structure
```
├── README.md
└── aws
    ├── dev
    │   └── apps
    │       ├── service-a.yaml
    │       └── service-b.yaml
    └── prod
        └── apps
            ├── service-a.yaml
            └── service-b.yaml
```

This structure is meant for argocd to deploy different helm charts on different cluster environments. For example everything under `aws/dev/apps` is meant to deployed on the AWS Dev EKS cluster.

To configure this ArgoCd also uses an Application and it would look like
```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: argocd-dev-app-manager
  namespace: argocd
  finalizers:
    - resources-finalizer.argocd.argoproj.io
spec:
  project: argocd
  source:
    repoURL: https://github.com/example-org/argo-config.git
    targetRevision: master
    path: aws/dev/apps
  destination:
    server: https://kubernetes.default.svc
    namespace: argocd
  syncPolicy:
    automated:
      prune: true
      selfHeal: true
      allowEmpty: true
    syncOptions:
    - Validate=true
    - CreateNamespace=true
    - PrunePropagationPolicy=foreground
    - PruneLast=true
    retry:
      limit: 1
```
This Application tells ArgoCD to deploy all Applications under `aws/dev/apps` to the Dev EKS Cluster.