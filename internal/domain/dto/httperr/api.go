package httperr

import "errors"

var ErrNotFound = errors.New("resource not found")
var ErrBadRequest = errors.New("bad request")
