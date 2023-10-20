package v2

import (
	"github.com/diy0663/go_project_packages/response"
	"github.com/diy0663/gohub/app/models/project"
	"github.com/gin-gonic/gin"
)

// todo 注意, 需要自行保存文件之后自动import 包进来
// todo 控制器这里有个版本控制 类似 v1, v2
// todo 还需要自行去添加对应的路由,以及是否加中间件等

type ProjectsController struct {

	// 需要确保 v2 文件夹下也有一个 BaseAPIController
	// BaseAPIController
}

func (ctrl *ProjectsController) Index(c *gin.Context) {
	projects := project.All()
	response.Data(c, projects)
}

func (ctrl *ProjectsController) Show(c *gin.Context) {
	projectModel := project.Get(c.Param("id"))
	if projectModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, projectModel)
}

func (ctrl *ProjectsController) Store(c *gin.Context) {
	/*
	   request := requests.ProjectRequest{}

	   	if ok := requests.RequestValidate(c, &request, requests.ProjectSave); !ok {
	   	    return
	   	}

	   	projectModel := project.Project{
	   	    FieldName:      request.FieldName,
	   	}

	   projectModel.Create()

	   	if projectModel.ID > 0 {
	   	    response.Created(c, projectModel)
	   	} else {

	   	    response.Abort500(c, "创建失败，请稍后尝试~")
	   	}
	*/
}

func (ctrl *ProjectsController) Update(c *gin.Context) {

	/*
	   projectModel := project.Get(c.Param("id"))

	   	if projectModel.ID == 0 {
	   	    response.Abort404(c)
	   	    return
	   	}

	   	if ok := policies.CanModifyProject(c, projectModel); !ok {
	   	    response.Abort403(c)
	   	    return
	   	}

	   request := requests.ProjectRequest{}
	   bindOk, errs := requests.Validate(c, &request, requests.ProjectSave)

	   	if !bindOk {
	   	    return
	   	}

	   	if len(errs) > 0 {
	   	    response.ValidationError(c, errs)
	   	    return
	   	}

	   projectModel.FieldName = request.FieldName
	   rowsAffected := projectModel.Save()

	   	if rowsAffected > 0 {
	   	    response.Data(c, projectModel)
	   	} else {

	   	    response.Abort500(c, "更新失败，请稍后尝试~")
	   	}
	*/
}

func (ctrl *ProjectsController) Delete(c *gin.Context) {

	/*
	   projectModel := project.Get(c.Param("id"))
	   if projectModel.ID == 0 {
	       response.Abort404(c)
	       return
	   }

	   if ok := policies.CanModifyProject(c, projectModel); !ok {
	       response.Abort403(c)
	       return
	   }

	   rowsAffected := projectModel.Delete()
	   if rowsAffected > 0 {
	       response.Success(c)
	       return
	   }

	*/
	response.Abort500(c, "删除失败，请稍后尝试~")
}
