package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/icbd/gohighlights/controllers/api"
	"github.com/icbd/gohighlights/models"
)

// GET /api/v1/users/:id
func UsersShow(c *gin.Context) {
	resp := api.New(c)
	u := api.CurrentUser(c)
	if u == nil {
		resp.NotFoundErr()
	} else {
		resp.OK(u)
	}
}

// POST /api/v1/users
func UsersCreate(c *gin.Context) {
	resp := api.New(c)
	vo := models.SessionVO{}
	if err := c.BindJSON(&vo); err != nil {
		resp.ParametersErr(err)
		return
	}

	u := models.User{Email: vo.Email, Password: vo.Password}
	if err := u.CalcPasswordHash(); err != nil {
		resp.ParametersErr(err)
		return
	}

	if err := models.DB().Create(&u).Error; err != nil {
		resp.ParametersErr(err)
		return
	}

	resp.Created(u)
}
