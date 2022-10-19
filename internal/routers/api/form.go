package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isther/clinic/internal/dao"
	"github.com/isther/clinic/internal/model"
	"github.com/sirupsen/logrus"
)

type UserFormApi struct{}

func NewUserFormApi() *UserFormApi {
	return &UserFormApi{}
}

func (api *UserFormApi) Create(c *gin.Context) {
	var newForm model.Form
	if err := c.ShouldBind(&newForm); err != nil {
		logrus.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if tx := dao.DB.Create(&model.FormSql{
		Form: newForm,
	}); tx.Error != nil {
		logrus.Error(tx.Error)
		c.JSON(http.StatusOK, gin.H{
			"msg": tx.Error,
		})
		return
	}

	logrus.Info(fmt.Sprintf("Create Form: %#v", newForm))
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func (api *UserFormApi) Query(c *gin.Context) {
	type query struct {
		ID string `json:"id"`
	}

	var newQuery = new(query)
	if err := c.ShouldBind(newQuery); err != nil {
		logrus.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	var formSqls []model.FormSql
	if tx := dao.DB.Where("student_id = ?", newQuery.ID).Find(&formSqls); tx.Error != nil {
		logrus.Error(tx.Error)
		c.JSON(http.StatusOK, gin.H{
			"msg": tx.Error,
		})
		return
	}

	logrus.Info(fmt.Sprintf("Query Form: %#v", formSqls))
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
		"res": formSqls,
	})
}
