package v1

import (
	"github.com/diy0663/go_project_packages/response"
	"github.com/diy0663/gohub/app/models/category"
	"github.com/diy0663/gohub/app/requests"
	"github.com/gin-gonic/gin"
)

// todo 注意, 需要自行保存文件之后自动import 包进来
// todo 控制器这里有个版本控制 类似 v1, v2
// todo 还需要自行去添加对应的路由,以及是否加中间件等

type CategoriesController struct {
	BaseAPIController
}

func (ctrl *CategoriesController) Index(c *gin.Context) {

	// 分页的参数验证
	request := requests.PaginationRequest{}
	if ok := requests.RequestValidate(c, &request, requests.Pagination); !ok {

		return
	}

	data, pager := category.Paginate(c, 20)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})

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

	if ok := requests.RequestValidate(c, &request, requests.CategorySave); !ok {
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

// 更新数据
func (ctrl *CategoriesController) Update(c *gin.Context) {

	request := requests.CategoryRequest{}
	bindOk := requests.RequestValidate(c, &request, requests.CategorySave)

	if !bindOk {
		return
	}

	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}
	// if ok := policies.CanModifyCategory(c, categoryModel); !ok {
	// 	response.Abort403(c)
	// 	return
	// }

	categoryModel.Name = request.Name
	categoryModel.Description = request.Description
	rowsAffected := categoryModel.Save()

	if rowsAffected > 0 {
		response.Data(c, categoryModel)
	} else {

		response.Abort500(c, "更新失败，请稍后尝试~")
	}

}

func (ctrl *CategoriesController) Delete(c *gin.Context) {

	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}

	// if ok := policies.CanModifyCategory(c, categoryModel); !ok {
	// 	response.Abort403(c)
	// 	return
	// }

	rowsAffected := categoryModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
