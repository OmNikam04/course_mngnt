package repository

import (
	"student-courses-with-transactions/internal/datastruct"
)

type UserQuery interface {
	GetUser(id int64) (*datastruct.Course, error)
	CreateUser(course datastruct.Course) (*int64, error)
	UpdateUser(course datastruct.Course) (*datastruct.Course, error)
	DeleteUser(course datastruct.Course) (*datastruct.Course, error)
}

type userQuery struct{}

// CreateUser implements UserQuery.
func (*userQuery) CreateUser(course datastruct.Course) (*int64, error) {
	panic("unimplemented")
}

// DeleteUser implements UserQuery.
func (*userQuery) DeleteUser(course datastruct.Course) (*datastruct.Course, error) {
	panic("unimplemented")
}

// GetUser implements UserQuery.
func (*userQuery) GetUser(id int64) (*datastruct.Course, error) {
	panic("unimplemented")
}

// UpdateUser implements UserQuery.
func (*userQuery) UpdateUser(course datastruct.Course) (*datastruct.Course, error) {
	panic("unimplemented")
}
