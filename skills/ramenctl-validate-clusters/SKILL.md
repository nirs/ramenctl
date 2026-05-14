---
name: ramenctl-validate-clusters
description: >-
  Validate disaster recovery cluster configuration using ramenctl validate
  clusters. Use when the user wants to check clusters are ok, verify DR
  setup, check cluster health, or troubleshoot cluster-level DR problems.
---

# ramenctl validate clusters

Validate disaster recovery clusters by gathering cluster-scoped and ramen
resources from all clusters, and checking that S3 endpoints are accessible.

## Prerequisites

- `ramenctl` installed
- A configuration file with cluster kubeconfigs (see ramenctl-init skill)

## Workflow

### Step 1: Pick an output directory

Ask the user where to store the report. Suggest `<env>/clusters` or
`<env>-<date>/clusters`, e.g., `myenv/clusters`.

### Step 2: Run validate clusters

```console
$ ramenctl validate clusters -o <output-dir> --interactive=false
```

`--interactive=false` prevents ramenctl from opening a browser—you open the
report yourself in Step 4.

Use `--config <file>` if the config file is not the default `config.yaml`.

The command takes a few seconds on local clusters and about a minute on
remote clusters.

### Step 3: Check the result

**On success** the last line shows:

```console
✅ Validation completed (N ok, 0 warning, 0 problem)
```

**On problems** the summary includes warning or problem counts and the
command exits with a non-zero exit code:

```console
❌ Validation completed (N ok, M warning, P problem)
```

When validation **finishes and writes a report**, a non-zero exit always means
there are problems **in** that report. If the command **fails before** a
report is written (early error), there may be no HTML—use the log and console
output instead.

### Step 4: Open the HTML report

If `<output-dir>/validate-clusters.html` exists, open it as soon as `ramenctl
validate clusters` has returned. Do this for **both** a fully successful run and
a completed run that exited non-zero: in the latter case the report always
describes problems. If the command failed **before** the HTML file was created,
skip opening—there is nothing to show in the browser yet.

Do **not** skip this step when the HTML file exists unless the user clearly
does not want the graphical report.

**How to open:** tell them the file path, or run `open` (macOS), `xdg-open`
(Linux, when available), or open the path in Windows Explorer / `start` on
the machine that has the file.

### Step 5: Inspect the report

The output directory contains:

| File | Purpose |
|------|---------|
| `validate-clusters.yaml` | Machine and human readable report |
| `validate-clusters.html` | HTML report |
| `validate-clusters.log` | Detailed log for troubleshooting |
| `validate-clusters.data/` | Raw resources gathered from all clusters |

Read the YAML report to understand the status:

```console
$ yq '.clustersStatus' < <output-dir>/validate-clusters.yaml
```

#### What the report validates

The `clustersStatus` section in the YAML report validates the following.
Every item shows `state: ok ✅` when healthy, or a problem/warning state
with a descriptive message.

**Each managed cluster:**
- Ramen configmap: exists, controller type is `dr-cluster`, S3 store
  profiles with bucket, endpoint, region, and secret fingerprints
- Ramen deployment: exists, conditions (Available, Progressing), replicas

**Hub cluster:**
- DRClusters: exist, conditions (Fenced, Clean, Validated), phase
- DRPolicies: exist, conditions (Validated), DR clusters, scheduling
  interval
- Ramen configmap: controller type is `dr-hub`, S3 store profiles
- Ramen deployment: exists, conditions, replicas

**S3 profiles:**
- Each profile accessible from the machine running ramenctl

Secret values are validated using sanitized fingerprints, allowing
comparison across clusters without exposing secrets.

#### Gathered data structure

```console
$ tree -L3 <output-dir>/validate-clusters.data
validate-clusters.data/
├── <cluster>/
│   ├── cluster/
│   │   ├── apiextensions.k8s.io/
│   │   ├── namespaces/
│   │   ├── nodes/
│   │   ├── ramendr.openshift.io/
│   │   ├── storage.k8s.io/
│   │   └── ...
│   └── namespaces/
│       └── ramen-system/
└── ...
```

### Step 6: Troubleshoot problems

This section is an **initial draft** — DR troubleshooting is hard and
needs more playbooks over time (for example a future analyze skill).

If problems were found, help the user understand them:

1. Read the YAML report — look for `state:` values that are not `ok ✅`
2. Check the log file for error details
3. Inspect raw resources in `validate-clusters.data/`
4. Summarize the findings and suggest what the user should look at

## Notes

- To report DR issues, archive the entire output directory and upload it
  to the issue tracker.
