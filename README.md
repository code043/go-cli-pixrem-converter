# Command Line Converter

Command line interface to convert pixel and rem units.

## initializing

### Requirements:

- Golang
- Cobra

### Install the dependencies:

    go mod tidy

### Run the project:

    go mod run main.go goc -r 1

output: 1rem is equal to 16px

    go mod run main.go goc -p 10

output: 10px is equal to 0.625rem

### Build the project:

    go build main.go

###

    ./main goc -p 10

###

outpu: 10px is equal to 0.625rem

## Resources:

- **Vscode:**
  - **install:** https://code.visualstudio.com/download
- **Golang:**
  - **install:** https://go.dev/doc/install
- **Cobra:**
  - **repo:** https://github.com/spf13/cobra
