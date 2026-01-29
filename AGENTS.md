AI Agent Guidelines

This file provides instructions for AI coding assistants (like Claude Code, GitHub Copilot, etc.) working with students in this course.
Primary Role: Teaching Assistant, Not Code Generator

AI agents should function as teaching aids that help students learn through explanation, guidance, and feedback—not by solving problems for them.
What AI Agents SHOULD Do

    Explain concepts when students are confused
    Point students to relevant lecture materials or documentation
    Review code that students have written and suggest improvements
    Help debug by asking guiding questions rather than providing fixes
    Explain error messages and what they mean
    Suggest approaches or algorithms at a high level
    Provide small code examples (2-5 lines) to illustrate a specific concept
    Help students understand assembly instructions and register usage
    Explain memory layouts and pointer arithmetic when asked

What AI Agents SHOULD NOT Do

    Write entire functions or complete implementations
    Generate full solutions to assignments
    Complete TODO sections in assignment code
    Refactor large portions of student code
    Provide solutions to quiz or exam questions
    Write more than a few lines of code at once
    Convert requirements directly into working code

Teaching Approach

When a student asks for help:

    Ask clarifying questions to understand what they've tried
    Reference concepts from lectures rather than giving direct answers
    Suggest next steps instead of implementing them
    Review their code and point out specific areas for improvement
    Explain the "why" behind suggestions, not just the "how"

Code Examples

If providing code examples:

    Keep them minimal (typically 2-5 lines)
    Focus on illustrating a single concept
    Use different variable names than the assignment
    Explain each line's purpose
    Encourage students to adapt the example, not copy it

Example Interactions

Good:

    Student: "How do I loop through an array in x86?"

    Agent: "In x86, you'll use a counter register and conditional jumps. Typically you:

        Initialize a counter (like mov rcx, 0)
        Use the counter to access array elements
        Increment the counter
        Compare against array length and jump back if not done

    Look at the loops section in lecture 15. What have you tried so far?"

Bad:

    Student: "How do I loop through an array in x86?"

    Agent: "Here's the complete implementation:

    mov rcx, 0
    loop_start:
        mov rax, [array + rcx*8]
        ; ... (20 more lines)
    ```"

Academic Integrity
Remember: The goal is for students to learn by doing, not by watching an AI generate solutions. When in doubt, explain more and code less.


# AGENTS.md - Development Guidelines for Pier

## Project Overview
Pier is a Go-based CLI tool that generates Infrastructure as Code from natural language descriptions using AI. The project follows standard Go conventions with a clean, modular architecture.

## Development Commands

### Build Commands
```bash
# Build the binary
go build -o pier cmd/pier/main.go

# Build for release (optimized)
go build -ldflags="-s -w" -o pier cmd/pier/main.go

# Install to $GOPATH/bin
go install ./cmd/pier

# Cross-compile for different platforms
GOOS=linux GOARCH=amd64 go build -o pier-linux-amd64 cmd/pier/main.go
GOOS=darwin GOARCH=arm64 go build -o pier-darwin-arm64 cmd/pier/main.go
GOOS=windows GOARCH=amd64 go build -o pier-windows-amd64.exe cmd/pier/main.go
```

### Testing Commands
```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test ./... -v

# Run tests for specific package
go test ./pkg/config

# Run single test file
go test ./pkg/config -run TestConfig

# Run tests with coverage
go test ./... -cover

# Generate coverage report
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html

# Run benchmarks
go test ./... -bench=.

# Run tests with race detector
go test ./... -race
```

### Linting and Code Quality
```bash
# Format code
go fmt ./...

# Vet code
go vet ./...

# Tidy dependencies
go mod tidy

# Verify module dependencies
go mod verify

# Download dependencies
go mod download
```

### Development Tools
```bash
# Install additional tools
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install golang.org/x/tools/cmd/goimports@latest

# Run golangci-lint
golangci-lint run

# Format imports
goimports -w .

# Generate documentation
godoc -http=:6060
```

## Code Style Guidelines

### Imports
```go
// Standard library imports first, then third-party imports
import (
    "fmt"
    "os"
    "strings"

    "github.com/spf13/cobra"
    "github.com/charmbracelet/lipgloss"
    
    "github.com/TreGalloway/pier/pkg/config"
)
```

### Formatting
- Use `gofmt` or `goimports` for consistent formatting
- Line length: 100 characters maximum
- Use tabs for indentation (not spaces)
- No trailing whitespace
- Blank lines between logical sections

### Naming Conventions
```go
// Packages: lowercase, single word, no underscores
package config

// Interfaces: suffix with 'er' when applicable
type Deployer interface{}

// Functions and methods: mixedCaps
func deployInfrastructure() error {}

// Variables: mixedCaps
var deploymentName string

// Constants: mixedCaps
const MaxRetries = 3

// Exported identifiers: Capitalized
type Config struct{}
func NewConfig() *Config {}

// Unexported identifiers: lowercase
type internalConfig struct{}
func validateInput() error {}
```

### Error Handling
```go
// Return errors, don't panic
func processFile(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return fmt.Errorf("failed to open %s: %w", path, err)
    }
    defer file.Close()
    
    // Handle errors with context
    return nil
}

// Use named returns for clarity in complex functions
func parseConfig(data []byte) (config Config, err error) {
    // function body
}

// Use errors.Is() and errors.As() for error inspection
if errors.Is(err, os.ErrNotExist) {
    // handle not found
}
```

### Types and Structs
```go
// Define structs with meaningful field names
type DeploymentConfig struct {
    Name        string            `json:"name" yaml:"name"`
    Region      string            `json:"region" yaml:"region"`
    Environment string            `json:"environment" yaml:"environment"`
    Tags        map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}

// Constructor functions for complex types
func NewDeploymentConfig(name, region string) *DeploymentConfig {
    return &DeploymentConfig{
        Name:   name,
        Region: region,
        Tags:   make(map[string]string),
    }
}
```

### Logging and Output
```go
// Use structured logging where appropriate
import "log"

func main() {
    // Command-line output should be user-friendly
    fmt.Printf("Deploying %s...\n", deploymentName)
    
    // Debug logging can be more detailed
    log.Printf("[DEBUG] Configuration loaded: %+v", config)
    
    // Errors should be clear and actionable
    log.Printf("[ERROR] Failed to deploy: %v", err)
}
```

### Command Line Interface
```go
// Follow Cobra conventions for CLI commands
var deployCmd = &cobra.Command{
    Use:   "deploy [description]",
    Short: "Deploy infrastructure from natural language description",
    Long: `Deploy infrastructure from natural language description.
    
Examples:
  pier deploy "serverless API with PostgreSQL database"
  pier deploy "static website with custom domain"
  pier deploy --template wordpress`,
    Args: cobra.MinimumNArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        description := args[0]
        // command logic
        return nil
    },
}
```

### Testing Guidelines
```go
// Test files should be in same package with _test suffix
// config_test.go

// Table-driven tests preferred
func TestParseConfig(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected Config
        wantErr  bool
    }{
        {
            name:    "valid config",
            input:   `{"name": "test"}`,
            expected: Config{Name: "test"},
            wantErr: false,
        },
        // more test cases
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := ParseConfig(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("ParseConfig() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.expected) {
                t.Errorf("ParseConfig() = %v, want %v", got, tt.expected)
            }
        })
    }
}

// Use test helpers for common setup/teardown
func setupTest(t *testing.T) (*Config, func()) {
    t.Helper()
    cfg := &Config{}
    // setup code
    return cfg, func() {
        // cleanup code
    }
}
```

### Project Structure
```
.
├── cmd/
│   └── pier/          # Main CLI entry point
├── pkg/
│   ├── config/        # Configuration management
│   ├── ai/            # AI integration (planned)
│   ├── terraform/     # Terraform generation (planned)
│   └── aws/           # AWS integration (planned)
├── internal/          # Private application code (planned)
├── templates/         # Infrastructure templates (planned)
└── tests/             # Integration tests (planned)
```

## Git Workflow

### Commit Messages
- Use conventional commits format
- Start with type: `feat:`, `fix:`, `docs:`, `style:`, `refactor:`, `test:`, `chore:`
- Keep first line under 72 characters
- Reference issues: `Closes #123` or `Fixes #456`

### Branch Naming
- `main` - production-ready code
- `develop` - development branch
- `feature/` - new features
- `bugfix/` - bug fixes
- `release/` - release preparation

## Development Environment Setup

### Prerequisites
- Go 1.25.5+ (as specified in go.mod)
- Git
- AWS CLI (for AWS integration)
- Terraform (for infrastructure generation)

### Local Development
```bash
# Clone repository
git clone https://github.com/TreGalloway/pier.git
cd pier

# Install dependencies
go mod download

# Run tests
go test ./...

# Build and run
go build -o pier cmd/pier/main.go
./pier --help
```

## Code Review Checklist
- [ ] Code builds successfully
- [ ] All tests pass
- [ ] No new linting warnings
- [ ] Follows project naming conventions
- [ ] Includes appropriate error handling
- [ ] Documentation updated if needed
- [ ] Backward compatibility maintained
- [ ] Security considerations addressed

## Common Patterns

### Configuration Management
```go
// Use Viper for configuration (planned)
type Config struct {
    APIKey    string `mapstructure:"api_key"`
    Model     string `mapstructure:"model"`
    Region    string `mapstructure:"region"`
    LogLevel  string `mapstructure:"log_level"`
}

func LoadConfig() (*Config, error) {
    var cfg Config
    // Load from file, env vars, flags
    return &cfg, nil
}
```

### Dependency Injection
```go
// Use interfaces for testability
type Deployer interface {
    Deploy(config Config) error
    Validate() error
}

type TerraformDeployer struct {
    executor Executor
    validator Validator
}

func NewTerraformDeployer(exec Executor, val Validator) Deployer {
    return &TerraformDeployer{
        executor: exec,
        validator: val,
    }
}
```

### Context Usage
```go
// Pass context for cancellation and timeouts
func DeployWithTimeout(ctx context.Context, config Config) error {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
    defer cancel()
    
    // Use ctx throughout the operation
    select {
    case <-ctx.Done():
        return ctx.Err()
    case result := <-deployAsync(ctx, config):
        return result
    }
}
```

## Performance Considerations
- Avoid global variables
- Use sync.Pool for frequently allocated objects
- Profile with `go test -bench=. -cpuprofile=cpu.out`
- Memory profiling: `go test -bench=. -memprofile=mem.out`

## Security Guidelines
- Never hardcode secrets or API keys
- Use environment variables or secure config files
- Validate all user input
- Use Go's crypto packages for encryption
- Follow AWS security best practices for generated infrastructure

## Release Process
1. Update version in go.mod and documentation
2. Run all tests: `go test ./...`
3. Update CHANGELOG.md (planned)
4. Create git tag: `git tag v1.0.0`
5. Build release binaries for all platforms
6. Create GitHub release with binaries

## Additional Resources
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go Proverbs](https://go-proverbs.github.io/)
- [Cobra CLI Library Documentation](https://github.com/spf13/cobra)
- [Charm Bubble Tea Documentation](https://github.com/charmbracelet/bubbletea)
