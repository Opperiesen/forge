// Package cli provides the Forge command-line entry point.
package cli

import (
	"fmt"
	"io"
)

const usage = `Forge is a GitOps engine for a single Linux container host.

Usage:
  forge version
  forge help

The reconciliation commands are not implemented yet.`

// Run executes the CLI and returns a process exit code.
func Run(args []string, stdout, stderr io.Writer, version string) int {
	if len(args) == 0 {
		fmt.Fprintln(stdout, usage)
		return 0
	}

	switch args[0] {
	case "help", "-h", "--help":
		fmt.Fprintln(stdout, usage)
		return 0
	case "version", "-v", "--version":
		fmt.Fprintf(stdout, "forge %s\n", version)
		return 0
	default:
		fmt.Fprintf(stderr, "forge: unknown command %q\n", args[0])
		fmt.Fprintln(stderr, "Run 'forge help' for usage.")
		return 2
	}
}
