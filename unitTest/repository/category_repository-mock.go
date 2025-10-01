package repository

import (
	"belajar-golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindByID(id string) *entity.Category {
	args := repository.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entity.Category)
}