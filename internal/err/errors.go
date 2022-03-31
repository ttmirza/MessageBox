package err

import "errors"

var DuplicateUserError = errors.New("user already exists")

var DuplicateGroupError = errors.New("group already exists")

var UserNotFoundError = errors.New("user not found")

var MessageNotFoundError = errors.New("message not found")
