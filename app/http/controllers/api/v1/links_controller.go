package v1

import (
	"github.com/diy0663/gohub/app/models/link"
	"github.com/diy0663/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

type LinksController struct {
	BaseApiController
}

func (ctrl *LinksController) Index(c *gin.Context) {

	links := link.All()
	response.Data(c, links)
}

// func (ctrl *LinksController) Show(c *gin.Context) {
// 	linkModel := link.Get(c.Param("id"))
// 	if linkModel.ID == 0 {
// 		response.Abort404(c)
// 		return
// 	}
// 	response.Data(c, linkModel)
// }

// func (ctrl *LinksController) Store(c *gin.Context) {

// 	request := requests.LinkRequest{}
// 	if ok := requests.Validate(c, &request, requests.LinkSave); !ok {
// 		return
// 	}

// 	linkModel := link.Link{
// 		FieldName: request.FieldName,
// 	}
// 	linkModel.Create()
// 	if linkModel.ID > 0 {
// 		response.Created(c, linkModel)
// 	} else {
// 		response.Abort500(c, "创建失败，请稍后尝试~")
// 	}
// }

// func (ctrl *LinksController) Update(c *gin.Context) {

// 	linkModel := link.Get(c.Param("id"))
// 	if linkModel.ID == 0 {
// 		response.Abort404(c)
// 		return
// 	}

// 	// 授权策略
// 	//if ok := policies.CanModifyLink(c, linkModel); !ok {
// 	//    response.Abort403(c)
// 	//     return
// 	// }

// 	request := requests.LinkRequest{}
// 	bindOk := requests.Validate(c, &request, requests.LinkSave)
// 	if !bindOk {
// 		return
// 	}

// 	linkModel.FieldName = request.FieldName
// 	rowsAffected := linkModel.Save()
// 	if rowsAffected > 0 {
// 		response.Data(c, linkModel)
// 	} else {
// 		response.Abort500(c, "更新失败，请稍后尝试~")
// 	}
// }

// func (ctrl *LinksController) Delete(c *gin.Context) {

// 	linkModel := link.Get(c.Param("id"))
// 	if linkModel.ID == 0 {
// 		response.Abort404(c)
// 		return
// 	}

// 	// if ok := policies.CanModifyLink(c, linkModel); !ok {
// 	//     response.Abort403(c)
// 	//     return
// 	//  }

// 	rowsAffected := linkModel.Delete()
// 	if rowsAffected > 0 {
// 		response.Success(c)
// 		return
// 	}

// 	response.Abort500(c, "删除失败，请稍后尝试~")
// }
