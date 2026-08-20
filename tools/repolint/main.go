// Command repolint checks repository invariants without external dependencies.
package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var (
	markdownLink = regexp.MustCompile(`\[[^\]]*\]\(([^)]+)\)`)
	actionUse    = regexp.MustCompile(`^\s*(?:-\s*)?uses:\s*([^\s#]+)`)
	commitSHA    = regexp.MustCompile(`^[0-9a-f]{40}$`)
)

var requiredPaths = []string{
	".gitattributes",
	".github/CODEOWNERS",
	".github/ISSUE_TEMPLATE/bug_report.yml",
	".github/ISSUE_TEMPLATE/feature_request.yml",
	".github/dependabot.yml",
	".github/workflows/ci.yml",
	"AGENTS.md",
	"CODE_OF_CONDUCT.md",
	"CONTRIBUTING.md",
	"LICENSE",
	"README.md",
	"SECURITY.md",
	"docs/ARCHITECTURE.md",
	"docs/MVP.md",
	"docs/VISION.md",
	"go.mod",
}

func main() {
	errors := checkRepository(".")
	if len(errors) == 0 {
		fmt.Println("repository hygiene checks passed")
		return
	}

	sort.Strings(errors)
	for _, err := range errors {
		fmt.Fprintln(os.Stderr, err)
	}
	os.Exit(1)
}

func checkRepository(root string) []string {
	var errors []string

	for _, path := range requiredPaths {
		if _, err := os.Stat(filepath.Join(root, path)); err != nil {
			errors = append(errors, fmt.Sprintf("missing required path: %s", path))
		}
	}

	walkErr := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			if entry.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}
		if !isTextFile(path) {
			return nil
		}

		relative, relErr := filepath.Rel(root, path)
		if relErr != nil {
			return relErr
		}
		content, readErr := os.ReadFile(path)
		if readErr != nil {
			return readErr
		}

		errors = append(errors, checkText(relative, content)...)
		if filepath.Ext(path) == ".md" {
			errors = append(errors, checkMarkdownLinks(root, relative, string(content))...)
		}
		if isWorkflow(relative) {
			errors = append(errors, checkActionPins(relative, string(content))...)
		}
		return nil
	})
	if walkErr != nil {
		errors = append(errors, fmt.Sprintf("walk repository: %v", walkErr))
	}

	return errors
}

func isTextFile(path string) bool {
	base := filepath.Base(path)
	switch base {
	case ".editorconfig", ".gitattributes", ".gitignore", "LICENSE":
		return true
	}
	switch filepath.Ext(path) {
	case ".go", ".md", ".mod", ".sum", ".yaml", ".yml":
		return true
	default:
		return false
	}
}

func checkText(path string, content []byte) []string {
	var errors []string
	if strings.Contains(string(content), "\r\n") {
		errors = append(errors, fmt.Sprintf("%s: contains CRLF line endings", path))
	}
	if len(content) > 0 && content[len(content)-1] != '\n' {
		errors = append(errors, fmt.Sprintf("%s: missing final newline", path))
	}
	if len(content) > 1 && content[len(content)-1] == '\n' && content[len(content)-2] == '\n' {
		errors = append(errors, fmt.Sprintf("%s: blank line at end of file", path))
	}

	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	line := 0
	for scanner.Scan() {
		line++
		if strings.TrimRight(scanner.Text(), " \t") != scanner.Text() {
			errors = append(errors, fmt.Sprintf("%s:%d: trailing whitespace", path, line))
		}
	}
	return errors
}

func checkMarkdownLinks(root, path, content string) []string {
	var errors []string
	for _, match := range markdownLink.FindAllStringSubmatch(content, -1) {
		target := strings.Trim(match[1], "<>")
		fields := strings.Fields(target)
		if len(fields) == 0 {
			continue
		}
		target = fields[0]
		if strings.HasPrefix(target, "#") {
			continue
		}
		parsed, err := url.Parse(target)
		if err != nil || parsed.Scheme != "" || parsed.Host != "" {
			continue
		}
		decoded, err := url.PathUnescape(parsed.Path)
		if err != nil {
			errors = append(errors, fmt.Sprintf("%s: invalid link %q", path, target))
			continue
		}
		resolved := filepath.Join(root, filepath.Dir(path), filepath.FromSlash(decoded))
		if _, err := os.Stat(resolved); err != nil {
			errors = append(errors, fmt.Sprintf("%s: unresolved local link %q", path, target))
		}
	}
	return errors
}

func isWorkflow(path string) bool {
	clean := filepath.ToSlash(path)
	return strings.HasPrefix(clean, ".github/workflows/") &&
		(filepath.Ext(clean) == ".yml" || filepath.Ext(clean) == ".yaml")
}

func checkActionPins(path, content string) []string {
	var errors []string
	scanner := bufio.NewScanner(strings.NewReader(content))
	line := 0
	for scanner.Scan() {
		line++
		match := actionUse.FindStringSubmatch(scanner.Text())
		if len(match) == 0 || strings.HasPrefix(match[1], "./") {
			continue
		}
		parts := strings.Split(match[1], "@")
		if len(parts) != 2 || !commitSHA.MatchString(parts[1]) {
			errors = append(errors, fmt.Sprintf("%s:%d: action must use a full commit SHA", path, line))
		}
	}
	return errors
}
