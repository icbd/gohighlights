package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/icbd/gohighlights/controllers/middleware"
	"github.com/icbd/gohighlights/errs"
	"github.com/icbd/gohighlights/models"
	"gorm.io/gorm"
	"net/http"
)

type API struct {
	C *gin.Context
}

func New(c *gin.Context) *API {
	return &API{C: c}
}

func (a *API) OK(payload interface{}) {
	a.C.JSON(http.StatusOK, payload)
}

func (a *API) Created(payload interface{}) {
	a.C.JSON(http.StatusCreated, payload)
}

func (a *API) NoContent() {
	a.C.Writer.WriteHeader(http.StatusNoContent)
	a.C.Writer.WriteHeaderNow()
}

func (a *API) InternalErr() {
	a.C.Writer.WriteHeader(http.StatusInternalServerError)
	a.C.Writer.WriteHeaderNow()
}

func (a *API) UnauthorizedErr() {
	a.C.Writer.WriteHeader(http.StatusUnauthorized)
	a.C.Writer.WriteHeaderNow()
}

func (a *API) ForbiddenErr() {
	a.C.Writer.WriteHeader(http.StatusForbidden)
	a.C.Writer.WriteHeaderNow()
}

func (a *API) ParametersErr(err error) {
	if vErr, ok := err.(validator.ValidationErrors); ok {
		msg := make([]string, len(vErr))
		for i, e := range vErr {
			msg[i] = e.Field() + "::" + e.Tag()
		}
		a.C.JSON(http.StatusBadRequest, errs.E{C: errs.Validation, M: msg})
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		a.C.Writer.WriteHeader(http.StatusNotFound)
		a.C.Writer.WriteHeaderNow()
	} else {
		a.C.JSON(http.StatusBadRequest, errs.E{C: errs.Parameters, M: []string{err.Error()}})
	}
}

func (a *API) NotFoundErr() {
	a.C.Writer.WriteHeader(http.StatusNotFound)
	a.C.Writer.WriteHeaderNow()
}

func (a *API) Err(httpCode int, businessCode int, err error) {
	a.C.JSON(httpCode, errs.E{C: businessCode, M: []string{err.Error()}})
}

// CurrentUser Always use this method after CurrentUserMiddleware,
// so you will get a valid user pointer.
func CurrentUser(c *gin.Context) *models.User {
	if v, ok := c.Get(middleware.CurrentUser); ok {
		if u, ok := v.(*models.User); ok {
			return u
		}
	}
	return nil
}
