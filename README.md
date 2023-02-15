# Go updater (goit)

This is a updater for Go, made with Go just for fun and maybe useful.

## How works

First, this is a CLI with a serie of commands like install and show current version installed.

For now only show current version and install a new removing previous installation.

In the future I want to create a scrap to download page (go.dev) to retrieve all available versions and their checksums.

## Using

Clone the project using gh:

```bash
gh repo clone Patrick564/go-installer-cli
```

Using SSH

```bash
git clone git@github.com:Patrick564/go-installer-cli.git
```

Then there are 2 options:

### Trying the CLI

```bash
go run .
```

### Build the project

Run:

```bash
go build -o ./build/goit
```

And exec:

```bash
./build/goit [command]
```
