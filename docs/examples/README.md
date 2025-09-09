# Example runs

This page contain examples runs for `validate clusters` and `validate
application` commands.

## Ramen not deployed

```console
% ramenctl validate clusters -o out/ramen-not-deployed
⭐ Using config "config.yaml"
⭐ Using report "out/ramen-not-deployed"

🔎 Validate config ...
   ✅ Config validated

🔎 Validate clusters ...
   ✅ Gathered data from cluster "hub"
   ✅ Gathered data from cluster "dr1"
   ✅ Gathered data from cluster "dr2"
   ❌ Issues found during validation

❌ validation failed (0 ok, 0 stale, 8 problem)
```

Report: [ramen-not-deployed.tar.gz](ramen-not-deployed.tar.gz)

## Ramen not configured

```console
% ramenctl validate clusters -o out/ramen-not-configured
⭐ Using config "config.yaml"
⭐ Using report "out/ramen-not-configured"

🔎 Validate config ...
   ✅ Config validated

🔎 Validate clusters ...
   ✅ Gathered data from cluster "hub"
   ✅ Gathered data from cluster "dr1"
   ✅ Gathered data from cluster "dr2"
   ❌ Issues found during validation

❌ validation failed (18 ok, 0 stale, 5 problem)
```

Report: [ramen-not-configured.tar.gz](ramen-not-configured.tar.gz)

## Ramen deployed and configured

```console
% ramenctl validate clusters -o out/ramen-deployed-and-configured
⭐ Using config "config.yaml"
⭐ Using report "out/ramen-deployed-and-configured"

🔎 Validate config ...
   ✅ Config validated

🔎 Validate clusters ...
   ✅ Gathered data from cluster "hub"
   ✅ Gathered data from cluster "dr2"
   ✅ Gathered data from cluster "dr1"
   ✅ Clusters validated

✅ Validation completed (36 ok, 0 stale, 0 problem)
```

Report: [ramen-deployed-and-configured.tar.gz](ramen-deployed-and-configured.tar.gz)

## Application protected

```console
% ramenctl validate application --namespace argocd --name appset-deploy-rbd -o out/application-protected
⭐ Using config "config.yaml"
⭐ Using report "out/application-protected"

🔎 Validate config ...
   ✅ Config validated

🔎 Validate application ...
   ✅ Inspected application
   ✅ Gathered data from cluster "dr2"
   ✅ Gathered data from cluster "dr1"
   ✅ Gathered data from cluster "hub"
   ✅ Application validated

✅ Validation completed (21 ok, 0 stale, 0 problem)
```

Report: [application-protected.tar.gz](application-protected.tar.gz)

## Application failing over

```console
% ramenctl validate application --namespace argocd --name appset-deploy-rbd -o out/application-failover
⭐ Using config "config.yaml"
⭐ Using report "out/application-failover"

🔎 Validate config ...
   ✅ Config validated

🔎 Validate application ...
   ✅ Inspected application
   ✅ Gathered data from cluster "dr1"
   ✅ Gathered data from cluster "dr2"
   ✅ Gathered data from cluster "hub"
   ❌ Issues found during validation

❌ validation failed (26 ok, 0 stale, 3 problem)
```

Report: [application-failing-over.tar.gz](application-failing-over.tar.gz)

## Application failed over

```console
% ramenctl validate application --namespace argocd --name appset-deploy-rbd -o out/application-failedover
⭐ Using config "config.yaml"
⭐ Using report "out/application-failedover"

🔎 Validate config ...
   ✅ Config validated

🔎 Validate application ...
   ✅ Inspected application
   ✅ Gathered data from cluster "dr1"
   ✅ Gathered data from cluster "dr2"
   ✅ Gathered data from cluster "hub"
   ✅ Application validated

✅ Validation completed (21 ok, 0 stale, 0 problem)
```

Report: [application-failedover.tar.gz](application-failedover.tar.gz)

## Application relocating

```console
% ramenctl validate application --namespace argocd --name appset-deploy-rbd -o out/application-relocating
⭐ Using config "config.yaml"
⭐ Using report "out/application-relocating"

🔎 Validate config ...
   ✅ Config validated

🔎 Validate application ...
   ✅ Inspected application
   ✅ Gathered data from cluster "dr1"
   ✅ Gathered data from cluster "dr2"
   ✅ Gathered data from cluster "hub"
   ❌ Issues found during validation

❌ validation failed (9 ok, 1 stale, 8 problem)
```

Report: [application-relocating.tar.gz](application-relocating.tar.gz)

## Application relocated

```console
% ramenctl validate application --namespace argocd --name appset-deploy-rbd -o out/application-relocated 
⭐ Using config "config.yaml"
⭐ Using report "out/application-relocated"

🔎 Validate config ...
   ✅ Config validated

🔎 Validate application ...
   ✅ Inspected application
   ✅ Gathered data from cluster "dr2"
   ✅ Gathered data from cluster "dr1"
   ✅ Gathered data from cluster "hub"
   ✅ Application validated

✅ Validation completed (21 ok, 0 stale, 0 problem)
```

Report: [application-relocated.tar.gz](application-relocated.tar.gz)
