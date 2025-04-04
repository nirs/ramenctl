// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package test

func Clean(configFile string, outputDir string) error {
	cmd, err := newCommand("test-clean", configFile, outputDir)
	if err != nil {
		return err
	}

	if !cmd.Validate() {
		return cmd.Failed()
	}

	if !cmd.CleanTests() {
		cmd.GatherData()
		return cmd.Failed()
	}

	if !cmd.Cleanup() {
		return cmd.Failed()
	}

	cmd.Passed()
	return nil
}
