# goastutil

The `goastutil` package provides utility functions for working with Go abstract
syntax tree (AST).

## Installation

To install the `goastutil` package, run the following command:

```shell
go get github.com/septemhill/goastutil
```

# Usage

To use the goastutil package, import it in your Go code:

```go
import "github.com/septemhill/goastutil"
```

You can then use the utility functions provided by the goastutil package in your
code.

# Comparison with `Jennifier` Package

[`Jennifier`](https://github.com/dave/jennifer) is a good and awesome code generator tool for Go. In most of cases
you should use it rather than `goastutil`. Unless you want to do some Source-to-Source (from Go to Go) code transformation and generation from AST, then we serve for you.


# TODO

- [ ] Introduce AST modify functionality
  - [ ] Imports
    - [X] Add
    - [ ] Remove
  - [ ] Declaration
    - [X] Add
    - [ ] Remove
  - [ ] Statement
    - [X] Prepend
    - [X] Append
    - [ ] Remove
  - [ ] Expression
    - [ ] Modify value
    - [ ] Variable renaming
  - [ ] Comments
- [ ] Add generic support

# Contributing

Contributions to the `goastutil` package are welcome! If you find any issues or
have any improvements to suggest, please open an issue or submit a pull request.
