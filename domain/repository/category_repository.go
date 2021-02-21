package repository

import (
	"github.com/haoqihan/category/domain/model"
	"github.com/jinzhu/gorm"
)

type ICategoryRepository interface {
	InitTable() error
	FindCategoryByID(int64) (*model.Category, error)
	CreateCategory(*model.Category) (int64, error)
	DeleteCategoryByID(int64) error
	UpDateCategory(*model.Category) error
	FindAll() ([]model.Category, error)
	FindCategoryByName(string) (*model.Category, error)
	FindCategoryByLevel(uint32) ([]model.Category, error)
	FindCategoryByParent(int64) ([]model.Category, error)
}

// 创建CategoryRepository
func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &CategoryRepository{mysqlDb: db}
}

type CategoryRepository struct {
	mysqlDb *gorm.DB
}

// 初始化表
func (u *CategoryRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Category{}).Error
}

// 根据用户ID查找用户信息
func (u *CategoryRepository) FindCategoryByID(CategoryID int64) (Category *model.Category, err error) {
	Category = &model.Category{}
	return Category, u.mysqlDb.First(Category, CategoryID).Error
}

// 创建用户
func (u *CategoryRepository) CreateCategory(Category *model.Category) (CategoryID int64, err error) {
	return Category.ID, u.mysqlDb.Create(Category).Error
}

// 根据用户ID删除用户
func (u *CategoryRepository) DeleteCategoryByID(CategoryID int64) error {
	return u.mysqlDb.Where("id = ?", CategoryID).Delete(&model.Category{}).Error
}

// 更新用户信息
func (u *CategoryRepository) UpDateCategory(Category *model.Category) error {
	return u.mysqlDb.Model(Category).Update(&Category).Error
}

// 获取结果集
func (u *CategoryRepository) FindAll() (CategoryAll []model.Category, err error) {
	return CategoryAll, u.mysqlDb.Find(&CategoryAll).Error
}

// 根据分类名称进行查找
func (u *CategoryRepository) FindCategoryByName(categoryName string) (Category *model.Category, err error) {
	Category = &model.Category{}
	return Category, u.mysqlDb.Where("category_name = ?", categoryName).Find(Category).Error
}

// 根据分类等级进行查找
func (u *CategoryRepository) FindCategoryByLevel(categoryLevel uint32) (CategorySlice []model.Category, err error) {
	CategorySlice = []model.Category{}
	return CategorySlice, u.mysqlDb.Where("category_level = ?", categoryLevel).Find(CategorySlice).Error
}

// 根据父id进行查询
func (u *CategoryRepository) FindCategoryByParent(categoryParent int64) (CategorySlice []model.Category, err error) {
	CategorySlice = []model.Category{}
	return CategorySlice, u.mysqlDb.Where("category_parent = ?", categoryParent).Find(CategorySlice).Error
}
