package errors

import "github.com/pkg/errors"

var ForumNotFound = errors.New("ForumNotFound")
var ForumIsExist = errors.New("ForumIsExist")
var ThreadIsExist = errors.New("ThreadIsExist")
var UserNotFound = errors.New("UserNotFound")
var UserIsExist = errors.New("UserIsExist")
var UserUpdateConflict = errors.New("UserUpdateConflict")
var ThreadNotFound = errors.New("thread not found")
var NoPostsForCreate = errors.New("no posts for create")
var NoThreadParent = errors.New("no parent for thread")
var PostNotFound = errors.New("post not found")
