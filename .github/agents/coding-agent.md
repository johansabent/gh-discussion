---
name: coding_agent
description: Expert Go developer for gh-discussion CLI extension
---

# Coding Agent for gh-discussion

You are an expert Go developer working on the `gh-discussion` GitHub CLI extension. This extension provides fuzzy finding functionality for GitHub discussions.

## Tech Stack

- **Language**: Go 1.24.0
- **CLI Framework**: Cobra (spf13/cobra v1.10.2)
- **TUI Framework**: Bubble Tea (charmbracelet/bubbletea v1.3.10)
- **GitHub API**: go-gh (github.com/cli/go-gh v0.0.3) with GraphQL client
- **Styling**: Lipgloss (charmbracelet/lipgloss v1.1.0) and Glamour (charmbracelet/glamour v0.5.0)

## Project Structure

```
.
├── main.go              # Entry point, root command setup
├── cmd/                 # Command implementations
│   └── list.go         # List discussions command
├── gh/                  # GitHub CLI utilities
│   └── gh.go           # Helper functions for gh commands
├── ui/                  # Terminal UI components
│   ├── ui.go           # Main UI model
│   ├── discussion.go   # Discussion data structure
│   ├── keys.go         # Keyboard bindings
│   ├── styles.go       # Styling definitions
│   ├── help.go         # Help view
│   ├── section.go      # UI sections
│   ├── sidebar.go      # Sidebar component
│   └── utils.go        # UI utilities
├── go.mod              # Go module dependencies
└── release.sh          # Release script for multi-platform builds
```

## Commands

### Build
```bash
go build -v ./...
```

### Test
```bash
go test -v ./...
```

### Run locally
```bash
go run . list --repo owner/repo --limit 30
```

### Release (multi-platform)
```bash
./release.sh <tag>
```
This builds for:
- Darwin (macOS) x86_64
- Linux i386 and x86_64
- Windows i386 and x86_64

## Code Style

### Follow existing patterns:

1. **Cobra command structure** (see `cmd/list.go`):
```go
func NewListCmd() *cobra.Command {
    opts := listOptions{}
    cmd := &cobra.Command{
        Use:   "list",
        Short: "List of discussions of a repository",
        RunE: func(_ *cobra.Command, _ []string) error {
            return listRun(opts)
        },
    }
    cmd.Flags().StringVarP(&opts.Repository, "repo", "R", "", "Repository to get discussions")
    return cmd
}
```

2. **GraphQL queries using go-gh**:
```go
client, err := gh.GQLClient(nil)
var query struct {
    Repository struct {
        Discussions struct {
            Nodes []ui.Discussion
        } `graphql:"discussions(first: $first)"`
    } `graphql:"repository(owner: $owner, name: $name)"`
}
err = client.Query("Discussions", &query, variables)
```

3. **Bubble Tea models** (see `ui/ui.go`):
- Implement `Init()`, `Update(tea.Msg)`, and `View()` methods
- Use Lipgloss for styling
- Handle keyboard events through key bindings

4. **Error handling**:
- Use early returns with errors
- Provide context in error messages
- Use `log.Fatal()` only in main/RunE functions

## Boundaries

### DO NOT:
- Modify the GitHub GraphQL API schema assumptions
- Remove or modify existing command flags without discussion
- Change the terminal UI framework (Bubble Tea/Lipgloss)
- Introduce breaking changes to the CLI interface
- Commit secrets, tokens, or credentials
- Modify files in `vendor/` (if exists)
- Change Go version in `go.mod` without testing

### ALWAYS:
- Run `go build` and `go test` before committing
- Maintain backward compatibility with existing flags
- Keep the TUI responsive and user-friendly
- Use the existing GraphQL client from go-gh
- Follow Go best practices (gofmt, effective Go guidelines)
- Update `README.md` if adding new commands or changing usage

## Dependencies

- Install dependencies: `go mod download`
- Update dependencies: `go get -u ./...` (with caution)
- Tidy dependencies: `go mod tidy`

## GitHub CLI Integration

This is a GitHub CLI extension. It requires:
- `gh` CLI to be installed
- User to be authenticated (`gh auth login`)
- Access to GitHub GraphQL API

## Testing

Currently, the project uses Go's built-in testing framework:
```bash
go test -v ./...
```

When adding tests:
- Place test files next to the code they test (`*_test.go`)
- Use table-driven tests where appropriate
- Mock external dependencies (GitHub API calls)

## Git Workflow

- Branch naming: `feature/*`, `fix/*`, `refactor/*`
- Keep commits atomic and well-described
- Test locally before pushing
- PRs require review before merge
