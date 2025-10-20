// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package report

import (
	"fmt"
	"html/template"
	"strings"
)

func (b *Build) Render(out *strings.Builder) {
	startObject(out, "Build")
	renderProperty(out, "version", b.Version)
	renderProperty(out, "commit", b.Commit)
	endObject(out)
}

// startObject start a new object fragment
func startObject(out *strings.Builder, name string) {
	name = template.HTMLEscapeString(name)
	fmt.Fprintf(out, "<div class=\"object\" id=\"%s\">\n", strings.ToLower(name))
	fmt.Fprintf(out, "<h2>%s</h2>\n", name)
	out.WriteString("<dl>\n")
}

// endObject end current object fragment
func endObject(out *strings.Builder) {
	out.WriteString("</dl>\n")
	out.WriteString("</div>\n")
}

// renderProperty write an object property
func renderProperty(out *strings.Builder, name, value string) {
	out.WriteString("  <div class=\"property\">\n")
	fmt.Fprintf(out, "    <dt>%s</dt>\n", template.HTMLEscapeString(name))
	fmt.Fprintf(out, "    <dd>%s</dd>\n", template.HTMLEscapeString(value))
	out.WriteString("  </div>\n")
}
