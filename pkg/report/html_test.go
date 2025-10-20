// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package report_test

import (
	"strings"
	"testing"

	"github.com/ramendr/ramenctl/pkg/helpers"
	"github.com/ramendr/ramenctl/pkg/report"
)

func TestRenderBuild(t *testing.T) {
	build := report.Build{Version: "v0.12.0", Commit: "548f48a2e4d85f042df2c5cbaf58acff618873ff"}
	out := &strings.Builder{}
	build.Render(out)
	expected := `<div class="object" id="build">
<h2>Build</h2>
<dl>
  <div class="property">
    <dt>version</dt>
    <dd>v0.12.0</dd>
  </div>
  <div class="property">
    <dt>commit</dt>
    <dd>548f48a2e4d85f042df2c5cbaf58acff618873ff</dd>
  </div>
</dl>
</div>
`
	actual := out.String()
	if actual != expected {
		t.Fatalf("html mismatch\n%s", helpers.UnifiedDiff(t, expected, actual))
	}
}
