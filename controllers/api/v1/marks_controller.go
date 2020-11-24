package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/icbd/gohighlights/controllers/api"
	"github.com/icbd/gohighlights/indices"
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

	if mark, err := models.MarkCreate(u, vo); err != nil {
		resp.ParametersErr(err)
	} else {
		if mIndex, err := indices.NewMark(mark); err == nil {
			mIndex.Fresh()
		}
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

	vo := models.MarkUpdateVO{HashKey: c.Param("hash_key")}
	if err := c.BindJSON(&vo); err != nil {
		resp.ParametersErr(err)
		return
	}

	if mark, err := models.MarkUpdate(u, &vo); err != nil {
		resp.ParametersErr(err)
	} else {
		if mIndex, err := indices.NewMark(mark); err == nil {
			mIndex.Fresh()
		}
		resp.OK(mark)
	}
}

/**
DELETE /api/v1/marks/:hash_key
*/
func MarksDestroy(c *gin.Context) {
	resp := api.New(c)
	u := api.CurrentUser(c)

	if m, err := models.MarkDestroy(u, c.Param("hash_key")); err != nil {
		resp.ParametersErr(err)
	} else {
		indices.DeleteBy(m.ID)
		resp.NoContent()
	}
}

/**
GET /api/v1/marks/query?url=encodeURIComponent(btoa(url))
*/
func MarksQuery(c *gin.Context) {
	resp := api.New(c)
	u := api.CurrentUser(c)
	marks := models.MarkQuery(u, c.Query("url"))
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
		"total": models.MarkTotal(u),
		"items": models.MarkList(u, pagination),
	})
}
