package usecases

import (
	"student-courses-with-transactions/internal/dto"
	"student-courses-with-transactions/internal/repository"
)

type CourseService interface {
	GetCourse(courseID int64) (*dto.Course, error)
	CreateCourse(course dto.Course) (*int64, error)
	UpdateCourse(course dto.Course) (*dto.Course, error)
	DeleteCourse(courseID int64, userID int64) error
}

type courseService struct {
	dao repository.DAO
}

func NewCourseService(dao repository.DAO) CourseService {
	return &courseService{dao: dao}
}


// CreateCourse implements CourseService.
func (*courseService) CreateCourse(course dto.Course) (*int64, error) {
	panic("unimplemented")
}

// DeleteCourse implements CourseService.
func (*courseService) DeleteCourse(courseID int64, userID int64) error {
	panic("unimplemented")
}

// GetCourse implements CourseService.
func (*courseService) GetCourse(courseID int64) (*dto.Course, error) {
	panic("unimplemented")
}

// UpdateCourse implements CourseService.
func (c *courseService) UpdateCourse(course dto.Course) (*dto.Course, error) {
	panic("unimplemented")
}
