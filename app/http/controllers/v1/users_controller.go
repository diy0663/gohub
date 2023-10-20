package v1

import (
	"github.com/diy0663/go_project_packages/response"
	"github.com/diy0663/gohub/pkg/auth"
	"github.com/gin-gonic/gin"
)

// todo 注意, 需要自行保存文件之后自动import 包进来
// todo 控制器这里有个版本控制 类似 v1, v2
// todo 还需要自行去添加对应的路由,以及是否加中间件等

type UsersController struct {
	BaseAPIController
}

func (user_controller *UsersController) CurrentUser(c *gin.Context) {
	userData := auth.CurrentUser(c)
	response.Data(c, userData)

}

// func (ctrl *UsersController) Index(c *gin.Context) {
// 	users := user.All()
// 	response.Data(c, users)
// }

// func (ctrl *UsersController) Show(c *gin.Context) {
// 	userModel := user.Get(c.Param("id"))
// 	if userModel.ID == 0 {
// 		response.Abort404(c)
// 		return
// 	}
// 	response.Data(c, userModel)
// }

// func (ctrl *UsersController) Store(c *gin.Context) {
// 	/*
// 	   request := requests.UserRequest{}

// 	   	if ok := requests.RequestValidate(c, &request, requests.UserSave); !ok {
// 	   	    return
// 	   	}

// 	   	userModel := user.User{
// 	   	    FieldName:      request.FieldName,
// 	   	}

// 	   userModel.Create()

// 	   	if userModel.ID > 0 {
// 	   	    response.Created(c, userModel)
// 	   	} else {

// 	   	    response.Abort500(c, "创建失败，请稍后尝试~")
// 	   	}
// 	*/
// }

// func (ctrl *UsersController) Update(c *gin.Context) {

// 	/*
// 	   userModel := user.Get(c.Param("id"))

// 	   	if userModel.ID == 0 {
// 	   	    response.Abort404(c)
// 	   	    return
// 	   	}

// 	   	if ok := policies.CanModifyUser(c, userModel); !ok {
// 	   	    response.Abort403(c)
// 	   	    return
// 	   	}

// 	   request := requests.UserRequest{}
// 	   bindOk, errs := requests.Validate(c, &request, requests.UserSave)

// 	   	if !bindOk {
// 	   	    return
// 	   	}

// 	   	if len(errs) > 0 {
// 	   	    response.ValidationError(c, errs)
// 	   	    return
// 	   	}

// 	   userModel.FieldName = request.FieldName
// 	   rowsAffected := userModel.Save()

// 	   	if rowsAffected > 0 {
// 	   	    response.Data(c, userModel)
// 	   	} else {

// 	   	    response.Abort500(c, "更新失败，请稍后尝试~")
// 	   	}
// 	*/
// }

// func (ctrl *UsersController) Delete(c *gin.Context) {

// 	/*
// 	   userModel := user.Get(c.Param("id"))
// 	   if userModel.ID == 0 {
// 	       response.Abort404(c)
// 	       return
// 	   }

// 	   if ok := policies.CanModifyUser(c, userModel); !ok {
// 	       response.Abort403(c)
// 	       return
// 	   }

// 	   rowsAffected := userModel.Delete()
// 	   if rowsAffected > 0 {
// 	       response.Success(c)
// 	       return
// 	   }

// 	*/
// 	response.Abort500(c, "删除失败，请稍后尝试~")
// }
