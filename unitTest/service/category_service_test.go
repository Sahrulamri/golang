package service

import (
	// "belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T){
	categoryRepository.Mock.On("FindByID", "123").Return(nil)
	category, err := categoryService.Get("123")

	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestcaategoryService_GetFound(t *testing.T){
	category := entity.Category{ID: 1, Name: "Laptop" }
	categoryRepository.Mock.On("FindByID", "123").Return(&category)

	result, err := categoryService.Get("123")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.ID, result.ID)
	assert.Equal(t, category.Name, result.Name)
}