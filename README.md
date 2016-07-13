# grind [![GoDoc](http://godoc.org/github.com/dimchansky/grind?status.png)](http://godoc.org/github.com/dimchansky/grind)

Grind polishes Go programs.

This version of grind is a fork from the original [rsc.io/grind](https://github.com/rsc/grind).
It contains updates that have not yet been merged into the original repository (not sure it is still maintained).

The motivation behind this fork was to merge contributions from others that allow `grind` to build on Go versions 1.4+.
The original repo only built on Go versions up to 1.3.
Updates and fixes beyond Go build versions are also welcome.

## Install
`go get -u github.com/dimchansky/grind`


## Usage
```bash
$ grind -h
usage: grind [-diff] [-v] packagepath... (or file...)
```
