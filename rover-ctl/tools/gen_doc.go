// Copyright 2025 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"log"

	"github.com/lukas016/controlplane/rover-ctl/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	if err := doc.GenMarkdownTree(cmd.NewRootCommand(), "docs"); err != nil {
		log.Fatalf("Failed to generate markdown docs: %v", err)
	}
}
