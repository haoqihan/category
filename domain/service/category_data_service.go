package service

import (
	"github.com/haoqihan/category/domain/model"
	"github.com/haoqihan/category/domain/repository"
)

type ICategoryDataService interface {
	AddCategory(Category *model.Category) (int64, error)
	DeleteCategory(int64) error
	UpdateCategory(Category *model.Category) (err error)
	FindCategoryByID(int64) (*model.Category, error)
	FindAllCategory() ([]model.Category, error)
	FindCategoryByName(string) (*model.Category, error)
	FindCategoryByLevel(uint32) ([]model.Category, error)
	FindCategoryByParent(int64) ([]model.Category, error)
}

// 创建实例
func NewCategoryDataService(CategoryRepository repository.ICategoryRepository) ICategoryDataService {
	return &CategoryDataService{CategoryRepository: CategoryRepository}
}

type CategoryDataService struct {
	CategoryRepository repository.ICategoryRepository
}

// 添加分类
func (u CategoryDataService) AddCategory(Category *model.Category) (int64, error) {
	return u.CategoryRepository.CreateCategory(Category)
}

// 删除分类
func (u CategoryDataService) DeleteCategory(CategoryID int64) error {
	return u.CategoryRepository.DeleteCategoryByID(CategoryID)
}

// 更新分类
func (u CategoryDataService) UpdateCategory(Category *model.Category) (err error) {
	return u.CategoryRepository.UpDateCategory(Category)
}

// 根据分类id进行查询
func (u CategoryDataService) FindCategoryByID(categoryID int64) (*model.Category, error) {
	return u.CategoryRepository.FindCategoryByID(categoryID)
}

// 查找全部信息
func (u CategoryDataService) FindAllCategory() ([]model.Category, error) {
	return u.CategoryRepository.FindAll()
}

// 根据分类名称进行查询
func (u CategoryDataService) FindCategoryByName(CategoryName string) (*model.Category, error) {
	return u.CategoryRepository.FindCategoryByName(CategoryName)
}

func (u CategoryDataService) FindCategoryByLevel(level uint32) ([]model.Category, error) {
	return u.CategoryRepository.FindCategoryByLevel(level)
}

func (u CategoryDataService) FindCategoryByParent(parent int64) ([]model.Category, error) {
	return u.CategoryRepository.FindCategoryByParent(parent)
}
