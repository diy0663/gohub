package v1

import (
	"github.com/diy0663/gohub/app/models/category"
	"github.com/diy0663/gohub/app/requests"
	"github.com/diy0663/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	BaseApiController
}

func (ctrl *CategoriesController) Index(c *gin.Context) {
	categories := category.All()
	response.Data(c, categories)
}

func (ctrl *CategoriesController) Show(c *gin.Context) {
	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, categoryModel)
}

func (ctrl *CategoriesController) Store(c *gin.Context) {

	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}

	categoryModel := category.Category{

		Name:        request.Name,
		Description: request.Description,
	}
	categoryModel.Create()
	if categoryModel.ID > 0 {
		response.Created(c, categoryModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

// func (ctrl *CategoriesController) Update(c *gin.Context) {

// 	categoryModel := category.Get(c.Param("id"))
// 	if categoryModel.ID == 0 {
// 		response.Abort404(c)
// 		return
// 	}

// 	if ok := policies.CanModifyCategory(c, categoryModel); !ok {
// 		response.Abort403(c)
// 		return
// 	}

// 	request := requests.CategoryRequest{}
// 	bindOk, errs := requests.Validate(c, &request, requests.CategorySave)
// 	if !bindOk {
// 		return
// 	}
// 	if len(errs) > 0 {
// 		response.ValidationError(c, errs)
// 		return
// 	}

// 	categoryModel.FieldName = request.FieldName
// 	rowsAffected := categoryModel.Save()
// 	if rowsAffected > 0 {
// 		response.Data(c, categoryModel)
// 	} else {
// 		response.Abort500(c, "更新失败，请稍后尝试~")
// 	}
// }

// func (ctrl *CategoriesController) Delete(c *gin.Context) {

// 	categoryModel := category.Get(c.Param("id"))
// 	if categoryModel.ID == 0 {
// 		response.Abort404(c)
// 		return
// 	}

// 	if ok := policies.CanModifyCategory(c, categoryModel); !ok {
// 		response.Abort403(c)
// 		return
// 	}

// 	rowsAffected := categoryModel.Delete()
// 	if rowsAffected > 0 {
// 		response.Success(c)
// 		return
// 	}

// 	response.Abort500(c, "删除失败，请稍后尝试~")
// }