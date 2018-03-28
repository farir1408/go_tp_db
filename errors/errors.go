package errors

import "github.com/pkg/errors"

var UserNotFound = errors.New("UserNotFound")
var UserIsExist = errors.New("UserIsExist")
var UserUpdateConflict = errors.New("UserUpdateConflict")
