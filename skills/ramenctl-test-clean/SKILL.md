---
name: ramenctl-test-clean
description: >-
  Clean up after ramenctl test run using ramenctl test clean. Use when the user
  wants to clean up test resources, remove test artifacts, or reset after a
  test run.
---

# ramenctl test clean

Delete resources created by `ramenctl test run`.

## Prerequisites

- A previous `ramenctl test run` was executed
- If the test failed due to an infrastructure issue (e.g., rbd-mirror
  down), **restore the infrastructure first** before cleaning. Cleaning
  before restoring may fail or leave leftovers.

## Workflow

### Step 1: Run test clean

Use the same output directory and config file as the test run:

```console
$ ramenctl test clean -o <output-dir>
```

Use `--config <file>` if the config file is not the default `config.yaml`.

The command unprotects and undeploys test applications, then cleans the
test environment (channels, namespaces).

### Step 2: Check the result

**On success:**

```console
✅ passed (N passed, 0 failed, 0 skipped)
```

The output directory gains additional files after clean:

| File | Purpose |
|------|---------|
| `test-clean.yaml` | Clean operation report |
| `test-clean.log` | Detailed log |

## Notes

- Always clean up after test runs to avoid leftover resources affecting
  subsequent tests.
- If clean fails, check the log for details. You may need to manually
  delete stuck resources.
