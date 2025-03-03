// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"errors"
	"fmt"
	"os"
)

var sampleConfig = `# %s configuration file

# Clusters configurations - modify to match your clusters.
clusters:
  hub:
    kubeconfigpath: hub.kubeconfig			# Add hub cluster kubeconfig
  c1:
    kubeconfigpath: primary.kubeconfig		# Add primary managed cluster kubeconfig
  c2:
    kubeconfigpath: secondary.kubeconfig	# Add secondary managed cluster kubeconfig

# DRPolicy - modify to match your actual DRPolicy
dr-policy: dr-policy

# ClusterSet - modify to match your Open Cluster Management configuration.
clusterset: default

# PVC specifications - modify to match your storage.
pvcspecs:
- name: rbd
  storageclassname: rook-ceph-block			# Add rbd storage class name
  accessmodes: ReadWriteOnce
- name: cephfs
  storageclassname: rook-cephfs-fs1			# Add cephfs storage class name
  accessmodes: ReadWriteMany

# Tests to run - modify for your use case.
# Available workloads: "deploy"
# Available deployers: "appset", "subscr", "disapp"
tests:
- workload: deploy
  deployer: appset
  pvcspec: rbd
`

func CreateSampleConfig(filename, creator string) error {
	content := fmt.Sprintf(sampleConfig, creator)
	if err := createFile(filename, []byte(content)); err != nil {
		if errors.Is(err, os.ErrExist) {
			return fmt.Errorf("configuration file %q already exists", filename)
		}
		return fmt.Errorf("failed to create %q: %w", filename, err)
	}
	return nil
}

func createFile(name string, content []byte) error {
	f, err := os.OpenFile(name, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0o600)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write(content); err != nil {
		return err
	}
	return f.Close()
}
