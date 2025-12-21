package utils

import "go-bakcend-todo-list/enums"

type AuthContext struct {
	UserID uint
	Role   enums.UserRole
}
