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
`grind [-diff] [-v] packagepath...`

Grind rewrites the source files in the named packages. When grind rewrites a file, it prints a line to standard error giving the name of the file and the rewrites applied.

As a special case, if the arguments are a list of Go source files, they are considered to make up a single package, which is then rewritten.

If the -diff flag is set, no files are rewritten. Instead grind prints the differences a rewrite would introduce.

Grind does not make backup copies of the files that it edits. Instead, use a version control system's “diff” functionality to inspect the changes that grind makes before committing them.

Grind is a work in progress. More rewrites are planned. The initial use case for grind is cleaning up Go code that looks like C code.

## Dead Code Elimination

Grind removes unreachable (dead) code. Code is considered unreachable if it is not the target of a goto statement and is preceded by a terminating statement (see [golang.org/ref/spec#Terminating_statements](http://golang.org/ref/spec#Terminating_statements)), or a break or continue statement.

## Goto Inlining

If the target of a goto is a block of code that is only reachable by following that goto statement, grind replaces the goto with a copy of the target code and deletes the original target code.

If the target of a goto is an explicit or implicit return statement, replaces the goto with a copy of the return statement.

## Unused Label Removal

Grind removes unused labels.

## Var Declaration Insertion

Grind rewrites := declarations of complex zero types, such as:

```go
x := (*T)(nil)
y := T{} // T a struct or array type
```

to use plain var declarations, as in:
```go
var x *T
var y T
```

## Var Declaration Placement

Grind moves var declarations as close as possible to their uses. When possible, it combines the declaration with the initialization, and it splits disjoint uses of a single variable into multiple variables.

For example, consider:
```go
var i int

...

for i = 0; i < 10; i++ {
	use(i)
}

...

for i = 0; i < 10; i++ {
	otherUse(i)
}
```

Grind deletes the declaration and rewrites both loop initializers to use a combined declaration and assignment (`i := 0`).

*Limitation*: Grind does not move variable declarations into loop bodies. It may in the future, provided it can determine that the variable is dead on entry to the body and that the variable does not escape (causing heap allocation inside the loop).

*Limitation*: Grind considers variables that appear in closures off limits. A more sophisticated analysis might consider them in the future.