package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/icbd/gohighlights/controllers/api"
	"github.com/icbd/gohighlights/models"
)

/**
POST /api/v1/marks
{
	url: "url without query params and anchor",
	tag: "color or custom tag string",
	hash_key: "uuid generate by frontend",
 	selection: ""
}
*/
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
PATCH /api/v1/marks/:hash_key
{
	tag: "new tag value"
}
*/
func MarksUpdate(c *gin.Context) {
	resp := api.New(c)
	u := api.CurrentUser(c)

	vo := models.MarkUpdateVO{}
	if err := c.BindJSON(&vo); err != nil {
		resp.ParametersErr(err)
		return
	}

	if mark, err := u.UpdateMark(c.Param("hash_key"), &vo); err != nil {
		resp.ParametersErr(err)
	} else {
		resp.OK(mark)
	}
}

/**
DELETE /api/v1/marks/:hash_key
*/
func MarksDestroy(c *gin.Context) {
	resp := api.New(c)
	u := api.CurrentUser(c)

	if err := u.DestroyMark(c.Param("hash_key")); err != nil {
		resp.ParametersErr(err)
	} else {
		resp.NoContent()
	}
}

/**
GET /api/v1/marks/query?url=encodeURIComponent(btoa(url))
*/
func MarksQuery(c *gin.Context) {
	resp := api.New(c)
	u := api.CurrentUser(c)
	marks := u.MarkQuery(c.Query("url"))
	resp.OK(marks)
}

/**
GET /api/v1/marks?page=1&size=10
*/
func MarksIndex(c *gin.Context) {
	resp := api.New(c)
	u := api.CurrentUser(c)

	pagination := models.Pagination
	if err := c.Bind(&pagination); err != nil {
		resp.ParametersErr(err)
		return
	}

	resp.OK(gin.H{
		"total": u.MarksTotal(),
		"items": u.MarksAll(pagination),
	})
}
