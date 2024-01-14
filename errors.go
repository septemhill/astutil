package goastutil

import "errors"

var (
	ErrNotAppendPrepend = errors.New("cannot prepend/append stmt on block stmt")
	ErrNotBlockStmt     = errors.New("cannot add stmt to non-block stmt")
	ErrInvalidStmt      = errors.New("cannot parse invalid stmt")
	ErrInvalidDecl      = errors.New("cannot parse invalid decl")
)
