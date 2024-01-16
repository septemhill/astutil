package goastutil

import "errors"

var (
	ErrNotAppendPrepend = errors.New("cannot prepend/append stmt on block stmt")
	ErrNotBlockStmt     = errors.New("cannot add stmt to non-block stmt")
	ErrInvalidStmt      = errors.New("cannot parse invalid stmt")
	ErrInvalidDecl      = errors.New("cannot parse invalid decl")
	ErrCaseClause       = errors.New("cannot prepend/append case stmt out of switch stmt")
	ErrCommClause       = errors.New("cannot prepend/append comm stmt out of select stmt")
	ErrOutOfIndex       = errors.New("index out of range")
)
