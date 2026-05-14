<!-- SPDX-FileCopyrightText: The RamenDR authors -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# ramenctl AI Skills

These skills teach an AI assistant how to drive **ramenctl** for disaster
recovery on your clusters. Start with the scene below — it shows what
this feels like in practice. **Init** (create a config) is the only step
that must come before everything else; after that, use whichever command
you need. Everything after the scene is detail: which skills exist, how to
name report directories, and how to install them in your tool.

## Example session

The following is a short screenplay-style scene: a human and an agent
use the ramenctl skills. (Monospace block so it reads like a script and
renders with a distinct background on GitHub and most viewers.)

```text
INT. CURSOR CHAT — DAY

USER
        Help me set up ramenctl for my clusters. My kubeconfigs are in
        ocp/: hub.yaml is the hub, c1.yaml and c2.yaml are the managed
        clusters.

AGENT
        (Runs ramenctl init, fills kubeconfigs, reads ClusterClaims on c1
        and c2, matches names on the hub, sets clusterSet.)
        Your config is ready.

USER
        Can you check the clusters?

AGENT
        (Runs ramenctl validate clusters -o ocp/clusters.)
        Validation completed (91 ok, 0 warning, 0 problem).
        (Opens validate-clusters.html in the browser.)

USER
        Validate an application too.

AGENT
        (Lists DRPCs on the hub.)
        I found two applications:
          1. rbd4 (openshift-dr-ops) — FailedOver
          2. appset-deploy-rbd (openshift-gitops) — Deployed
        Which one should I validate?

USER
        rbd4.

AGENT
        (Runs ramenctl validate application --name rbd4 --namespace
        openshift-dr-ops -o ocp/rbd4.)
        All clear. Opening validate-application.html.

USER
        You are a very skillful agent!

AGENT
        Glad it helped.
```

## Available skills

| Skill | Description |
|-------|-------------|
| [ramenctl-init](ramenctl-init/SKILL.md) | Create a configuration file for your clusters |
| [ramenctl-validate-clusters](ramenctl-validate-clusters/SKILL.md) | Validate disaster recovery cluster configuration |
| [ramenctl-validate-application](ramenctl-validate-application/SKILL.md) | Validate a DR-protected application |
| [ramenctl-gather-application](ramenctl-gather-application/SKILL.md) | Gather diagnostic data for a protected application |
| [ramenctl-test-run](ramenctl-test-run/SKILL.md) | Run disaster recovery flow tests |
| [ramenctl-test-clean](ramenctl-test-clean/SKILL.md) | Clean up after test runs |

## Output directory tips

Most commands require an `-o <output-dir>` flag. A few tips:

- Name after the environment, optionally with a date (e.g., `myenv` or
  `myenv-2026-05-15`).
- Multiple commands can share one directory — each creates its own files
  (`validate-clusters.yaml`, `validate-application.yaml`, etc.) with no
  conflicts. Running the same command again creates numbered files
  (`validate-clusters-2.yaml`, etc.).
- This makes it easy to attach all info to a bug report as one archive.
- If the archive is too large for upload limits (e.g., GitHub's 25 MB),
  use subdirectories like `myenv/clusters` and `myenv/app-name`.

## Common tasks (no fixed order)

The only step everyone shares: run **init** (create and edit a config file)
before any other ramenctl command. Validate, gather, and test all need that
configuration.

After that, there is no typical sequence — pick what matches your goal.

- **Cluster health**: Use validate clusters when you want to know whether
  DR is wired correctly across the hub and managed clusters.
- **One application**: Use validate application when you care about a
  specific protected app (after failover, relocate, or when something looks
  wrong). When investigating issues, the agent should suggest also running
  validate clusters (e.g. `<env>/clusters`) so developers get both the app
  slice and the cluster-wide DR picture for triage or bug reports.
- **End-to-end DR**: Use test run (and test clean) when you want to exercise
  the full flow with a small workload from the config.
- **Snapshot for debugging**: Developers often use gather application to
  pull a full picture of an app (namespaces, cluster resources, S3) for a bug
  report or offline inspection — with or without running validate first.

## Using with AI tools

The skills are plain markdown files with workflow instructions. They work
with any AI coding assistant that can read markdown.

### Cursor

Copy the `skills/` directory to `~/.cursor/skills/` for personal use, or
to `.cursor/skills/` in your project:

```console
$ cp -r skills/ ~/.cursor/skills/
```

Cursor discovers skills automatically from the YAML frontmatter in each
`SKILL.md` file.

### Claude Code

Add to your `CLAUDE.md` or reference directly. For example, to include
all skills:

```console
$ cat skills/*/SKILL.md >> CLAUDE.md
```

Or reference individual skills when starting a conversation:

```text
Read skills/ramenctl-validate-clusters/SKILL.md and help me validate my clusters.
```

### Other tools (Codex, Aider, OpenCode, etc.)

These are standard markdown files. You can:

- Concatenate them into whatever instruction file your tool uses
- Reference them directly in your prompts
- Copy the content into your tool's configuration

The YAML frontmatter at the top of each file (`---\nname: ...\n---`) is
standard metadata that most tools ignore gracefully.
