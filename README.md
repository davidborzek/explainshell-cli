# explainshell-cli

> CLI wrapper for explainshell.com

## Usage

```bash
explain [command]
```

### Example

Explain the command `ls -lah`:

```bash
explain ls -lah
```

Result:

```
list directory contents
__________________________________________________

-l     use a long listing format
__________________________________________________

-a, --all
       do not ignore entries starting with .
__________________________________________________

-h, --human-readable
       with -l, print sizes in human readable format (e.g., 1K 234M 2G)
__________________________________________________
```

## Installation

```bash
go install github.com/davidborzek/explainshell-cli/cmd/explain@latest
```
