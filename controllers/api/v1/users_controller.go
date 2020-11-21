package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/icbd/gohighlights/controllers/api"
	"github.com/icbd/gohighlights/models"
	"gorm.io/gorm"
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

/**
Register or Login

POST /api/v1/users
{
	email: "",
	password: ""
}
*/
func UsersCreate(c *gin.Context) {
	resp := api.New(c)
	vo := models.SessionVO{}
	if err := c.ShouldBind(&vo); err != nil {
		resp.ParametersErr(err)
		return
	}

	u, err := models.FindUserByEmail(vo.Email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			resp.ParametersErr(err)
			return
		}
		// register new user
		u.Email = vo.Email
		u.Password = vo.Password
		if err := u.Create(); err != nil {
			resp.ParametersErr(err)
			return
		}
	} else {
		// check use password
		var ok bool
		if u, ok = vo.CurrentUser(); !ok {
			resp.UnauthorizedErr()
			return
		}
	}

	// login
	if s, err := u.GenerateSession(); err != nil {
		resp.ParametersErr(err)
	} else {
		s.User = *u
		resp.Created(s)
	}
}
