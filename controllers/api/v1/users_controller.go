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

	u, err := models.UserFindByEmail(vo.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		resp.InternalErr()
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// register new user
		u, err = models.UserCreate(vo.Email, vo.Password)
		if err != nil {
			resp.ParametersErr(err)
			return
		}
	} else {
		// check use password
		u.Password = vo.Password
		if !u.ValidPassword() {
			resp.UnauthorizedErr()
			return
		}
	}

	// login
	if s, err := u.GenerateSession(); err != nil {
		resp.ParametersErr(err)
	} else {
		resp.Created(s)
	}
}
