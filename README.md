# e2s
Environment Variable to Kubernetes Secret

A command-line tool to convert environment variables or .env files into Kubernetes Secret YAML format.

## Features

- Convert environment variables to Kubernetes Secret YAML
- Load variables from .env files
- Customizable secret name and namespace
- Automatic base64 encoding of values
- Cross-platform support (Linux, macOS, Windows)

## Installation

### Download Binary
Download the latest binary from the [releases page](https://github.com/takutakahashi/e2s/releases).

### Build from Source
```bash
git clone https://github.com/takutakahashi/e2s.git
cd e2s
make build
```

### Using Go Install
```bash
go install github.com/takutakahashi/e2s@latest
```

## Usage

### Basic Usage
```bash
# Convert current environment variables
e2s --name my-secret --namespace production

# Load from .env file
e2s --env-file .env --name app-config --namespace default
```

### Options
- `--env-file`: Path to .env file
- `--name`: Name of the Kubernetes Secret (default: "app-secret")
- `--namespace`: Namespace for the Kubernetes Secret (default: "default")

### Example Output
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: app-config
  namespace: production
type: Opaque
data:
  DATABASE_URL: cG9zdGdyZXM6Ly91c2VyOnBhc3N3b3JkQGxvY2FsaG9zdDo1NDMyL215ZGI=
  API_KEY: eW91ci1hcGkta2V5LWhlcmU=
```

## Development

### Prerequisites
- Go 1.21 or later
- Make (optional, for using Makefile)

### Building
```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Run tests
make test

# Run linting
make lint
```

### Docker
```bash
# Build Docker image
docker build -t e2s .

# Run with Docker
docker run --rm -e DATABASE_URL=postgres://... e2s --name my-secret
```
