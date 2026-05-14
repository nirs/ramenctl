---
name: ramenctl-gather-application
description: >-
  Gather diagnostic data for a DR-protected application using ramenctl gather
  application. Use when the user wants to collect data for troubleshooting,
  create a bug report, or inspect application resources across clusters.
---

# ramenctl gather application

Gather data for a specific DR-protected application across the hub, managed
clusters, and S3 storage. Unlike `validate application`, this command only
collects data without validating it.

## Prerequisites

- `ramenctl` installed
- A configuration file with cluster kubeconfigs (see ramenctl-init skill)
- At least one DR-protected application in the clusters

## Workflow

### Step 1: Look up protected applications

List DRPCs on the hub using the hub kubeconfig from the ramenctl config file:

```console
$ kubectl get drpc -A --kubeconfig <hub-kubeconfig>
NAMESPACE   NAME                   AGE   PREFERREDCLUSTER   FAILOVERCLUSTER   DESIREDSTATE   CURRENTSTATE
argocd      appset-deploy-rbd      69m   dr1                dr2               Relocate       Relocated
```

Present the list and let the user choose which application to gather
data for.

### Step 2: Run gather application

Use `<env>/<app-name>` for the output directory, e.g., `myenv/myapp`.

```console
$ ramenctl gather application --name <drpc-name> --namespace <namespace> -o <output-dir>
```

Use `--config <file>` if the config file is not the default `config.yaml`.

The command takes a few seconds on local clusters and about a minute on
remote clusters.

### Step 3: Check the result

On success the command ends with:

```console
✅ Gather completed
```

On failure, check `<output-dir>/gather-application.log` for details.

### Step 4: Inspect gathered data

The output directory contains:

| File | Purpose |
|------|---------|
| `gather-application.yaml` | Report describing the gather operation |
| `gather-application.log` | Detailed log |
| `gather-application.data/` | Gathered resources |

#### Gathered data structure

One directory per cluster plus `s3/`. Each cluster directory has
`cluster/` for cluster-scoped resources and `namespaces/` for namespaced
resources. Ramen operator logs are gathered under
`namespaces/ramen-system/pods/`.

```console
$ tree -L3 <output-dir>/gather-application.data
gather-application.data/
├── dr1/
│   ├── cluster/
│   │   ├── namespaces/
│   │   ├── persistentvolumes/
│   │   └── storage.k8s.io/
│   └── namespaces/
│       ├── e2e-appset-deploy-rbd/
│       └── ramen-system/
├── dr2/
│   ├── cluster/
│   └── namespaces/
│       ├── e2e-appset-deploy-rbd/
│       └── ramen-system/
├── hub/
│   ├── cluster/
│   └── namespaces/
│       ├── argocd/
│       └── ramen-system/
└── s3/
    ├── minio-on-dr1/
    │   └── test-appset-deploy-rbd/
    └── minio-on-dr2/
        └── test-appset-deploy-rbd/
```

#### Inspecting gathered resources

Check VRG protected PVC conditions:

```console
$ yq '.status.protectedPVCs[0].conditions' < <output-dir>/gather-application.data/<cluster>/namespaces/<ns>/ramendr.openshift.io/volumereplicationgroups/<name>.yaml
```

Search ramen operator logs for errors related to the app:

```console
$ grep -E 'ERROR.+<app-name>' <output-dir>/gather-application.data/<cluster>/namespaces/ramen-system/pods/*/manager/current.log
```

## Notes

- Secrets in gathered data are automatically sanitized.
- To report DR issues, archive the entire output directory and upload it
  to the issue tracker.
- Use `ramenctl validate application` if you need both data collection
  and problem detection.
