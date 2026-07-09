[![CI](https://github.com/Ilya-Kornilov/go-project-242/actions/workflows/ci.yml/badge.svg)](https://github.com/Ilya-Kornilov/go-project-242/actions/workflows/ci.yml)

# hexlet-path-size
A command-line utility written in Go that calculates the size of files and directories.

## Usage

### Basic usage
```bash
./bin/hexlet-path-size testdata.txt
24B	testdata.txt
```
### Flags
```bash
-H, --human: Human-readable sizes (e.g., 1.9MB).
-a, --all: Include hidden files and directories (those starting with a dot).
-r, --recursive: Calculate the total size including all nested subdirectories.
```
### Examples
Sum all files recursively, including hidden ones, in human-readable format:
```bash
./bin/hexlet-path-size -H -a -r project/
31.0MB	project/
```
### Help output
```bash
NAME:
   hexlet-path-size - print size of a file or directory; 
   supports -r (recursive), -H (human-readable), -a (include hidden)

USAGE:
   hexlet-path-size [global options] <path>

GLOBAL OPTIONS:
   --recursive, -r  recursive size of directories (default: false)
   --human,     -H  human-readable sizes (auto-select unit) (default: false)
   --all,       -a  include hidden files and directories (default: false)
   --help,      -h  show help
