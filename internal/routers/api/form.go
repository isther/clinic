package api

import "github.com/gin-gonic/gin"

type UserFormApi struct{}

func NewUserFormApi() *UserFormApi {
	return &UserFormApi{}
}

func (userCon *UserFormApi) Create(c *gin.Context) {}
func (userCon *UserFormApi) Query(c *gin.Context)  {}
