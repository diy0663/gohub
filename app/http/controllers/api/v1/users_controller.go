package v1

import (
	"github.com/diy0663/gohub/app/models/user"
	"github.com/diy0663/gohub/app/requests"
	"github.com/diy0663/gohub/pkg/auth"
	"github.com/diy0663/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseApiController
}

func (ctrl UsersController) CurrentUser(c *gin.Context) {

	// 经过中间件那里会判断,假如确实登录了会把整个userModel 的值设置在上下文参数 current_user 里面, 没有的话会是空结构
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}

func (ctrl *UsersController) Index(c *gin.Context) {

	// 校验关于分页的传参是否正确
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := user.Paginate(c, 10)
	response.Data(c, gin.H{
		"data":  data,
		"pager": pager,
	})

}

// func (ctrl *UsersController) Show(c *gin.Context) {
// 	userModel := user.Get(c.Param("id"))
// 	if userModel.ID == 0 {
// 		response.Abort404(c)
// 		return
// 	}
// 	response.Data(c, userModel)
// }

// func (ctrl *UsersController) Store(c *gin.Context) {

// 	request := requests.UserRequest{}
// 	if ok := requests.Validate(c, &request, requests.UserSave); !ok {
// 		return
// 	}

// 	userModel := user.User{
// 		FieldName: request.FieldName,
// 	}
// 	userModel.Create()
// 	if userModel.ID > 0 {
// 		response.Created(c, userModel)
// 	} else {
// 		response.Abort500(c, "创建失败，请稍后尝试~")
// 	}
// }

func (ctrl *UsersController) UpdateProfile(c *gin.Context) {

	userModel := auth.CurrentUser(c)
	if userModel.ID == 0 {
		response.Abort404(c)
		return
	}

	// if ok := policies.CanModifyUser(c, userModel); !ok {
	// 	response.Abort403(c)
	// 	return
	// }

	request := requests.UserUpdateProfileRequest{}
	bindOk := requests.Validate(c, &request, requests.UserUpdateProfile)
	if !bindOk {
		return
	}

	userModel.Name = request.Name
	userModel.City = request.City
	userModel.Introduction = request.Introduction

	rowsAffected := userModel.Save()
	if rowsAffected > 0 {
		response.Data(c, userModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

// func (ctrl *UsersController) Delete(c *gin.Context) {

// 	userModel := user.Get(c.Param("id"))
// 	if userModel.ID == 0 {
// 		response.Abort404(c)
// 		return
// 	}

// 	if ok := policies.CanModifyUser(c, userModel); !ok {
// 		response.Abort403(c)
// 		return
// 	}

// 	rowsAffected := userModel.Delete()
// 	if rowsAffected > 0 {
// 		response.Success(c)
// 		return
// 	}

// 	response.Abort500(c, "删除失败，请稍后尝试~")
// }
