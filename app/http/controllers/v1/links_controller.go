package v1

import (
	"github.com/diy0663/go_project_packages/response"
	"github.com/diy0663/gohub/app/models/link"
	"github.com/gin-gonic/gin"
)

// todo 注意, 需要自行保存文件之后自动import 包进来
// todo 控制器这里有个版本控制 类似 v1, v2
// todo 还需要自行去添加对应的路由,以及是否加中间件等

type LinksController struct {
	BaseAPIController
}

func (ctrl *LinksController) Index(c *gin.Context) {

	data := link.All()
	response.Data(c, data)

}

func (ctrl *LinksController) Show(c *gin.Context) {
	linkModel := link.Get(c.Param("id"))
	if linkModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, linkModel)
}

func (ctrl *LinksController) Store(c *gin.Context) {
	/*
	   request := requests.LinkRequest{}

	   	if ok := requests.RequestValidate(c, &request, requests.LinkSave); !ok {
	   	    return
	   	}

	   	linkModel := link.Link{
	   	    FieldName:      request.FieldName,
	   	}

	   linkModel.Create()

	   	if linkModel.ID > 0 {
	   	    response.Created(c, linkModel)
	   	} else {

	   	    response.Abort500(c, "创建失败，请稍后尝试~")
	   	}
	*/
}

func (ctrl *LinksController) Update(c *gin.Context) {

	/*
		    // 先校验参数是否通过,之后再去查数据库
		    request := requests.LinkRequest{}
		    bindOk := requests.RequestValidate(c, &request, requests.LinkSave)

			if !bindOk {
				return
			}

		    linkModel := link.Get(c.Param("id"))
		    if linkModel.ID == 0 {
		        response.Abort404(c)
		        return
		    }

		    if ok := policies.CanModifyLink(c, linkModel); !ok {
		        response.Abort403(c)
		        return
		    }



		    linkModel.FieldName = request.FieldName
		    rowsAffected := linkModel.Save()
		    if rowsAffected > 0 {
		        response.Data(c, linkModel)
		    } else {
		        response.Abort500(c, "更新失败，请稍后尝试~")
		    }
	*/
}

func (ctrl *LinksController) Delete(c *gin.Context) {

	/*
	   linkModel := link.Get(c.Param("id"))
	   if linkModel.ID == 0 {
	       response.Abort404(c)
	       return
	   }

	   if ok := policies.CanModifyLink(c, linkModel); !ok {
	       response.Abort403(c)
	       return
	   }

	   rowsAffected := linkModel.Delete()
	   if rowsAffected > 0 {
	       response.Success(c)
	       return
	   }

	*/
	response.Abort500(c, "删除失败，请稍后尝试~")
}
