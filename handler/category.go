package handler

import (
	"context"
	"github.com/haoqihan/category/common"
	"github.com/haoqihan/category/domain/model"
	"github.com/haoqihan/category/domain/service"
	log "github.com/micro/go-micro/v2/logger"

	category "github.com/haoqihan/category/proto/category"
)

type Category struct {
	CategoryDataService service.ICategoryDataService
}

// 提供创建分类的服务
func (c *Category) CreateCategory(ctx context.Context, req *category.CategoryRequest, res *category.CreateCategoryResponse) error {
	Category := &model.Category{}
	// 赋值
	err := common.SwapTo(req, Category)
	if err != nil {
		return err
	}
	categoryId, err := c.CategoryDataService.AddCategory(Category)
	if err != nil {
		return err
	}
	res.Message = "分类添加成功"
	res.CategoryId = categoryId
	return nil

}

// 提供分类更新服务
func (c *Category) UpdateCategory(ctx context.Context, req *category.CategoryRequest, res *category.UpdateCategoryResponse) error {
	Category := &model.Category{}
	err := common.SwapTo(req, Category)
	if err != nil {
		return err
	}
	err = c.CategoryDataService.UpdateCategory(Category)
	if err != nil {
		return err
	}
	res.Message = "分类更新成功"
	return nil
}

// 提供分类删除服务
func (c *Category) DeleteCategory(ctx context.Context, req *category.DeleteCategoryRequest, res *category.DeleteCategoryResponse) error {
	err := c.CategoryDataService.DeleteCategory(req.CategoryId)
	if err != nil {
		return err
	}
	res.Message = "删除成功"
	return nil
}

// 根据分类名称查找分类
func (c *Category) FindCategoryByName(ctx context.Context, req *category.FindByNameRequest, res *category.CategoryResponse) error {
	Category, err := c.CategoryDataService.FindCategoryByName(req.CategoryName)
	if err != nil {
		return err
	}
	return common.SwapTo(Category, res)
}

// 根据分类ID查找分类
func (c *Category) FindCategoryByID(ctx context.Context, req *category.FindByIDRequest, res *category.CategoryResponse) error {
	Category, err := c.CategoryDataService.FindCategoryByID(req.CategoryId)
	if err != nil {
		return err
	}
	return common.SwapTo(Category, res)
}

func (c *Category) FindCategoryByLevel(ctx context.Context, req *category.FindByLevelRequest, res *category.FindAllCategoryResponse) error {
	categorySlice, err := c.CategoryDataService.FindCategoryByLevel(req.Level)
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, res)
	return nil

}
func (c *Category) FindCategoryParent(ctx context.Context, req *category.FindByParentRequest, res *category.FindAllCategoryResponse) error {
	categorySlice, err := c.CategoryDataService.FindCategoryByParent(req.ParentId)
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, res)
	return nil
}
func (c *Category) FindAllCategory(ctx context.Context, req *category.FindAllRequest, res *category.FindAllCategoryResponse) error {
	categorySlice, err := c.CategoryDataService.FindAllCategory()
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, res)
	return nil
}

func categoryToResponse(categorySlice []model.Category, response *category.FindAllCategoryResponse) {
	for _, cg := range categorySlice {
		cr := &category.CategoryResponse{}
		err := common.SwapTo(cg, cr)
		if err != nil {
			log.Error(err)
			break
		}
		response.Category = append(response.Category, cr)
	}

}
