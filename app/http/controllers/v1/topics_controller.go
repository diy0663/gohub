package v1

import (
	"github.com/diy0663/go_project_packages/response"
	"github.com/diy0663/gohub/app/models/topic"
	"github.com/diy0663/gohub/app/requests"
	"github.com/diy0663/gohub/pkg/auth"
	"github.com/gin-gonic/gin"
)

// todo 还需要自行去添加对应的路由,以及是否加中间件等

type TopicsController struct {
	BaseAPIController
}

func (ctrl *TopicsController) Index(c *gin.Context) {

	// 分页的参数验证
	request := requests.PaginationRequest{}
	if ok := requests.RequestValidate(c, &request, requests.Pagination); !ok {

		return
	}

	data, pager := topic.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})

}

func (ctrl *TopicsController) Show(c *gin.Context) {
	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, topicModel)
}

func (ctrl *TopicsController) Store(c *gin.Context) {

	request := requests.TopicRequest{}

	if ok := requests.RequestValidate(c, &request, requests.TopicSave); !ok {
		return
	}

	topicModel := topic.Topic{
		Title:      request.Title,
		Body:       request.Body,
		CategoryID: request.CategoryID,

		UserID: auth.CurrentUID(c),
	}

	topicModel.Create()

	if topicModel.ID > 0 {
		response.Created(c, topicModel)
	} else {

		response.Abort500(c, "创建失败，请稍后尝试~")
	}

}

func (ctrl *TopicsController) Update(c *gin.Context) {

	/*
		    // 先校验参数是否通过,之后再去查数据库
		    request := requests.TopicRequest{}
		    bindOk := requests.RequestValidate(c, &request, requests.TopicSave)

			if !bindOk {
				return
			}

		    topicModel := topic.Get(c.Param("id"))
		    if topicModel.ID == 0 {
		        response.Abort404(c)
		        return
		    }

		    if ok := policies.CanModifyTopic(c, topicModel); !ok {
		        response.Abort403(c)
		        return
		    }



		    topicModel.FieldName = request.FieldName
		    rowsAffected := topicModel.Save()
		    if rowsAffected > 0 {
		        response.Data(c, topicModel)
		    } else {
		        response.Abort500(c, "更新失败，请稍后尝试~")
		    }
	*/
}

func (ctrl *TopicsController) Delete(c *gin.Context) {

	/*
	   topicModel := topic.Get(c.Param("id"))
	   if topicModel.ID == 0 {
	       response.Abort404(c)
	       return
	   }

	   if ok := policies.CanModifyTopic(c, topicModel); !ok {
	       response.Abort403(c)
	       return
	   }

	   rowsAffected := topicModel.Delete()
	   if rowsAffected > 0 {
	       response.Success(c)
	       return
	   }

	*/
	response.Abort500(c, "删除失败，请稍后尝试~")
}
