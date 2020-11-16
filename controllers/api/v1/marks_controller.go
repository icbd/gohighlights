package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/icbd/gohighlights/controllers/api"
	"github.com/icbd/gohighlights/models"
)

// POST /api/v1/marks
//{
//  url: "",
//	tag: "",
//	hash_key: "",
//  selection: ""
//}
func MarksCreate(c *gin.Context) {
	resp := api.New(c)
	u := api.CurrentUser(c)

	vo := models.MarkCreateVO{}
	if err := c.BindJSON(&vo); err != nil {
		resp.ParametersErr(err)
		return
	}

	if mark, err := u.CreateMark(vo); err != nil {
		resp.ParametersErr(err)
	} else {
		resp.Created(mark)
	}
}

/**
GET /api/v1/marks/query
{
	url: "",
}
*/
func MarksQuery(c *gin.Context) {
	resp := api.New(c)
	u := api.CurrentUser(c)
	resp.OK(u.MarkQuery(c.Query("url")))
}

/**
GET /api/v1/marks
{
	page: 1,
	size: 10
}
*/
func MarksIndex(c *gin.Context) {
	resp := api.New(c)
	u := api.CurrentUser(c)

	pagination := models.Pagination
	if err := c.BindJSON(&pagination); err != nil {
		resp.ParametersErr(err)
		return
	}

	resp.OK(gin.H{
		"total": u.MarksTotal(),
		"items": u.MarksAll(pagination),
	})
}
