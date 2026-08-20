package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCheckActionPins(t *testing.T) {
	t.Parallel()

	const pinned = "steps:\n  - uses: actions/checkout@3d3c42e5aac5ba805825da76410c181273ba90b1 # v7.0.1\n"
	if errors := checkActionPins("ci.yml", pinned); len(errors) != 0 {
		t.Fatalf("checkActionPins() returned %v for pinned action", errors)
	}

	const mutable = "steps:\n  - uses: actions/checkout@v7\n"
	if errors := checkActionPins("ci.yml", mutable); len(errors) != 1 {
		t.Fatalf("checkActionPins() returned %v, want one error", errors)
	}
}

func TestCheckMarkdownLinks(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	docs := filepath.Join(root, "docs")
	if err := os.Mkdir(docs, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(docs, "target.md"), []byte("# Target\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	content := "[local](target.md) [anchor](#section) [remote](https://example.com)"
	if errors := checkMarkdownLinks(root, "docs/source.md", content); len(errors) != 0 {
		t.Fatalf("checkMarkdownLinks() returned %v for valid links", errors)
	}

	if errors := checkMarkdownLinks(root, "docs/source.md", "[missing](absent.md)"); len(errors) != 1 {
		t.Fatalf("checkMarkdownLinks() returned %v, want one error", errors)
	}
}

func TestCheckText(t *testing.T) {
	t.Parallel()

	if errors := checkText("clean.md", []byte("clean\n")); len(errors) != 0 {
		t.Fatalf("checkText() returned %v for clean text", errors)
	}

	errors := checkText("dirty.md", []byte("trailing \r\nno-final-newline"))
	if len(errors) != 3 {
		t.Fatalf("checkText() returned %v, want CRLF, trailing whitespace, and final newline errors", errors)
	}

	if errors := checkText("blank.md", []byte("content\n\n")); len(errors) != 1 {
		t.Fatalf("checkText() returned %v, want one blank EOF error", errors)
	}
}
