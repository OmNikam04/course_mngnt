package repository

import (
	"student-courses-with-transactions/internal/datastruct"
)

type CourseQuery interface {
	GetCourse(id int64) (*datastruct.Course, error)
	CreateCourse(course datastruct.Course) (*int64, error)
	UpdateCourse(course datastruct.Course) (*datastruct.Course, error)
	DeleteCourse(course datastruct.Course) (*datastruct.Course, error)
}

type courseQuery struct{}

// CreateCourse implements CourseQuery.
func (*courseQuery) CreateCourse(course datastruct.Course) (*int64, error) {
	panic("unimplemented")
}

// DeleteCourse implements CourseQuery.
func (*courseQuery) DeleteCourse(course datastruct.Course) (*datastruct.Course, error) {
	panic("unimplemented")
}

// GetCourse implements CourseQuery.
func (*courseQuery) GetCourse(id int64) (*datastruct.Course, error) {
	panic("unimplemented")
}

// UpdateCourse implements CourseQuery.
func (*courseQuery) UpdateCourse(course datastruct.Course) (*datastruct.Course, error) {
	panic("unimplemented")
}
