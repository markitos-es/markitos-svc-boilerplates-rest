package domain

import "errors"

var ErrBoilerplateNotFound error = errors.New("boilerplate not found")
var ErrBoilerplateBadRequest error = errors.New("bad request")
var ErrInvalidBoilerplateId error = errors.New("invalid boilerplate id")
var ErrInvalidBoilerplateName error = errors.New("invalid boilerplate name")
var ErrInvalidPageNumber error = errors.New("invalid page number")
var ErrInvalidPageSize error = errors.New("invalid page size")
