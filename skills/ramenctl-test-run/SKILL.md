---
name: ramenctl-test-run
description: >-
  Run disaster recovery tests using ramenctl test run. Use when the user wants
  to test DR flow, verify failover/relocate works, or run end-to-end DR tests.
---

# ramenctl test run

Run a disaster recovery flow test with one or more tiny applications defined
in the configuration file.

## Prerequisites

- `ramenctl` installed
- A configuration file with both common options and test options configured
  (see ramenctl-init skill). The test command needs `repo`, `drPolicy`,
  `pvcSpecs`, `deployers`, and `tests` sections.
- Clusters fully configured for DR (DRPolicy, DRClusters, S3, storage
  replication).

## Workflow

### Step 1: Verify the configuration

Help the user check their config file has the test sections:

- **drPolicy** - Must match an actual DRPolicy in the hub cluster
- **pvcSpecs** - Storage class names must match the managed clusters
- **tests** - Each test specifies workload + deployer + pvcSpec

Available deployer types: `appset`, `subscr`, `disapp`, `disapp-recipe`,
`disapp-recipe-check`, `disapp-recipe-exec`, `disapp-recipe-check-exec`.

### Step 2: Run the test

Use `<env>/test` for the output directory, e.g., `myenv/test`.

```console
$ ramenctl test run -o <output-dir>
```

Use `--config <file>` if the config file is not the default `config.yaml`.

The test runs through the full DR flow for each test case:
deploy, protect, failover, relocate, unprotect, undeploy.

**This typically takes 15-20 minutes per test case.** Multiple test cases
run in parallel.

### Step 3: Check the result

**On success:**

```console
✅ passed (N passed, 0 failed, 0 skipped)
```

**On failure** the command automatically gathers diagnostic data from all
clusters and S3 into `test-run.data/`:

```console
❌ failed (N passed, M failed, 0 skipped)
```

### Step 4: Inspect the report

The output directory contains:

| File | Purpose |
|------|---------|
| `test-run.yaml` | Machine-readable test report |
| `test-run.log` | Detailed log |
| `test-run.data/` | Gathered data (only present on failure) |

Useful commands:

Overall status:

```console
$ yq '.status' < <output-dir>/test-run.yaml
```

Individual step status and durations:

```console
$ yq '.steps[-1].items' < <output-dir>/test-run.yaml
```

Test events from the log:

```console
$ grep -E '(INFO|ERROR).+<app-name>' <output-dir>/test-run.log
```

The `test-run.yaml` report contains the config used, each test step with
its status and duration, and a summary with pass/fail/skip counts.

### Step 5: Handle failures

If the test failed:

1. Check `test-run.yaml` to see which step failed
2. Inspect `test-run.data/` for gathered resources and logs — it has the
   same structure as validate-application gathered data (one directory per
   cluster plus `s3/`)
3. Check DRPC status, VRG conditions, and operator logs
4. Fix the underlying issue
5. **Always clean up** before re-running (see ramenctl-test-clean skill)

### Step 6: Clean up

After the test (pass or fail), always clean up:

```console
$ ramenctl test clean -o <output-dir>
```

## Notes

- Pressing Ctrl+C cancels gracefully. The current step may take up to
  10 minutes to complete. Partial results are saved but data is not
  gathered for incomplete tests.
- To report DR issues, archive the entire output directory.
