package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunVersion(t *testing.T) {
	t.Parallel()

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := Run([]string{"version"}, &stdout, &stderr, "test")
	if code != 0 {
		t.Fatalf("Run() code = %d, want 0", code)
	}
	if got, want := stdout.String(), "forge test\n"; got != want {
		t.Fatalf("Run() stdout = %q, want %q", got, want)
	}
	if stderr.Len() != 0 {
		t.Fatalf("Run() stderr = %q, want empty", stderr.String())
	}
}

func TestRunUnknownCommand(t *testing.T) {
	t.Parallel()

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := Run([]string{"deploy"}, &stdout, &stderr, "test")
	if code != 2 {
		t.Fatalf("Run() code = %d, want 2", code)
	}
	if stdout.Len() != 0 {
		t.Fatalf("Run() stdout = %q, want empty", stdout.String())
	}
	if !strings.Contains(stderr.String(), `unknown command "deploy"`) {
		t.Fatalf("Run() stderr = %q, want unknown command message", stderr.String())
	}
}
