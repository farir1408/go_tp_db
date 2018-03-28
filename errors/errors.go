package errors

import "github.com/pkg/errors"


var ForumNotFound = errors.New("ForumNotFound")
var ForumIsExist = errors.New("ForumIsExist")
var ThreadIsExist = errors.New("ThreadIsExist")
var UserNotFound = errors.New("UserNotFound")
var UserIsExist = errors.New("UserIsExist")
var UserUpdateConflict = errors.New("UserUpdateConflict")

