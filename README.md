# Lit - A Fun Git Clone from Scratch in Golang

Lit is a lightweight, Git-like version control system implemented from scratch in Go. It aims to help developers understand Git's internals by reimplementing its core functionality, including repository initialization, object storage, commits, and more.

## Features

- Simple and educational Git-like functionality
- Initialize a Lit repository (`lit init`)
- Add files to the staging area (`lit add`)
- Commit changes (`lit commit`)
- View commit logs (`lit log`)
- And much more

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/lit.git
cd lit

# Initialize the Go module
go mod tidy

# Build the CLI
go build -o lit
```

## Getting Started

```bash
./lit help
```

## Contributions

Contributions are welcome! Feel free to submit issues, feature requests, or pull requests.

## License

MIT License
