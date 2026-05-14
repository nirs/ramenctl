---
name: ramenctl-init
description: >-
  Create a ramenctl configuration file for disaster recovery clusters using
  ramenctl init. Use when the user wants to initialize, configure, or set up
  ramenctl for their clusters.
---

# ramenctl init

Create a configuration file required by all other ramenctl commands.

## Workflow

### Step 1: Ask the user about their environment

Before creating the config, ask:
- Do they have a ramen testing environment with an envfile?
- Where are their kubeconfig files? The user must specify which
  kubeconfig is the hub and which are the managed clusters (c1, c2).
- Do they want a custom config file name (default is `config.yaml`)?

### Step 2: Create the configuration file

**For real clusters:**

```console
$ ramenctl init

✅ Created config file "config.yaml" - please modify for your clusters
```

**For a ramen testing environment with an envfile:**

```console
$ ramenctl init --envfile <path-to-envfile>
⭐ Using envfile "<path-to-envfile>"

✅ Created config file "config.yaml" - please modify for your clusters
```

**To use a custom config name:**

```console
$ ramenctl init --config myenv.yaml

✅ Created config file "myenv.yaml" - please modify for your clusters
```

The command creates `config.yaml` (or the specified name) in the current
directory. It fails if the file already exists.

### Step 3: Configure clusters

Open the generated file and help the user set the kubeconfig paths:

```yaml
clusters:
  hub:
    kubeconfig: my-hub.yaml
  passive-hub:
    kubeconfig: ""
  c1:
    kubeconfig: my-c1.yaml
  c2:
    kubeconfig: my-c2.yaml
```

Set `passive-hub` kubeconfig for optional passive hub cluster, or leave
empty if not using passive hub.

### Step 4: Discover and configure clusterSet

The clusterSet can be discovered automatically using the kubeconfigs from
step 3. Follow this procedure:

**4a. Get the OCM cluster names for c1 and c2.**

Kubernetes does not have a built-in cluster name concept. In OCM, each
managed cluster reports its name via a `ClusterClaim` resource. Get the
OCM name from each managed cluster:

```console
$ kubectl get clusterclaim name --kubeconfig <c1-kubeconfig> -o jsonpath='{.spec.value}'
dr1
```

```console
$ kubectl get clusterclaim name --kubeconfig <c2-kubeconfig> -o jsonpath='{.spec.value}'
dr2
```

**If the ClusterClaim command fails**, stop and report the error to the
user. ramenctl uses the same mechanism to resolve cluster names, so it
will not work without it. Common errors:

- `error: the server doesn't have a resource type "clusterclaim"` — OCM
  klusterlet is not installed on the cluster.
- `Error from server (NotFound)` — OCM is installed but the `name`
  claim hasn't been created yet.
- Connection or authentication errors — the kubeconfig is wrong, expired,
  or lacks permissions.

**4b. List managed clusters and their clusterSets on the hub.**

```console
$ kubectl get managedclusters --kubeconfig <hub-kubeconfig> \
  -L cluster.open-cluster-management.io/clusterset
NAME            HUB ACCEPTED   MANAGED CLUSTER URLS                 JOINED   AVAILABLE   AGE   CLUSTERSET
my-c1           true           https://api.my-c1.example.com:6443   True     True        16d   dr-clusters
my-c2           true           https://api.my-c2.example.com:6443   True     True        16d   dr-clusters
local-cluster   true           https://api.my-hub.example.com:6443  True     True        16d   default
```

**4c. Find the clusterSet that contains both clusters.**

Match the OCM names from step 4a against the `NAME` column, and read the
`CLUSTERSET` column.

- If both clusters are in the same clusterSet, use that value.
- If the clusters are in different clusterSets, something is
  misconfigured. Ask the user to verify.
- Ignore the `global` clusterSet (it automatically includes all clusters
  and is not suitable for DR).

With the default OCM `ExclusiveClusterSetLabel` selector type, a cluster
belongs to exactly one clusterSet (set via the
`cluster.open-cluster-management.io/clusterset` label). If there are
multiple non-global clusterSets, present the options and let the user
choose.

**4d. Set the clusterSet in the config file.**

```yaml
clusterSet: dr-clusters
```

### Step 5: Configure test options (optional)

Ask the user if they plan to run DR tests with `ramenctl test run`. If
they only need `validate` or `gather` commands, the configuration is
complete — skip to step 6.

If the user wants to run tests, the following additional sections need
to be configured:
- **repo** - Git repository URL and branch for test workloads
- **drPolicy** - DRPolicy name in the hub cluster
- **pvcSpecs** - Storage class names and access modes
- **deployers** - Deployer configurations (appset, subscr, disapp)
- **tests** - Test cases (workload + deployer + pvcSpec combinations)

### Step 6: Done

There is no standalone config validation command yet. The configuration
will be validated automatically when running any ramenctl command
(`validate clusters`, `validate application`, `test run`, etc.).

## Reference

- Default config file: `config.yaml` (used by all commands unless overridden
  with `--config`)
- Multiple config files can coexist for different environments or test sets
- When using `--envfile`, the file is pre-populated but may still need
  adjustments
