package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/icbd/gohighlights/controllers/api"
	"github.com/icbd/gohighlights/models"
)

/**
POST /marks/:hash_key/comment
{
	content: "comment content",
}
*/
func CommentsCreate(c *gin.Context) {
	resp := api.New(c)
	u := api.CurrentUser(c)

	vo := models.CommentVO{}
	if err := c.BindJSON(&vo); err != nil {
		resp.ParametersErr(err)
		return
	}

	mark, err := models.MarkFindByHashKey(u.ID, c.Param("hash_key"))
	if err != nil {
		resp.ParametersErr(err)
		return
	}

	if comment, err := models.CommentCreate(u.ID, mark.ID, vo.Content); err != nil {
		resp.ParametersErr(err)
	} else {
		resp.Created(comment)
	}
}

/**
PATCH /marks/:hash_key/comment
{
	content: "new comment content",
}
*/
func CommentsUpdate(c *gin.Context) {
	resp := api.New(c)
	u := api.CurrentUser(c)

	vo := models.CommentVO{}
	if err := c.BindJSON(&vo); err != nil {
		resp.ParametersErr(err)
		return
	}

	mark, err := models.MarkFindByHashKey(u.ID, c.Param("hash_key"))
	if err != nil {
		resp.ParametersErr(err)
		return
	}

	if comment, err := models.CommentUpdate(u.ID, mark.ID, vo.Content); err != nil {
		resp.ParametersErr(err)
	} else {
		resp.OK(comment)
	}
}

/**
DELETE /marks/:hash_key/comment
*/
func CommentsDestroy(c *gin.Context) {
	resp := api.New(c)
	u := api.CurrentUser(c)

	mark, err := models.MarkFindByHashKey(u.ID, c.Param("hash_key"))
	if err != nil {
		resp.ParametersErr(err)
		return
	}

	if _, err := models.CommentDestroy(u.ID, mark.ID); err != nil {
		resp.ParametersErr(err)
	} else {
		resp.NoContent()
	}
}
