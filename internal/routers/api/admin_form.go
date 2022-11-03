package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isther/clinic/internal/dao"
	"github.com/isther/clinic/internal/model"
	"github.com/sirupsen/logrus"
)

type AdminFormApi struct{}

func NewAdminFormApi() *AdminFormApi {
	return &AdminFormApi{}
}

func (api *AdminFormApi) Status(c *gin.Context) {
	var newQuery = &struct {
		FormID string `json:"form_id"`
		Status string `json:"status"`
	}{}

	if err := c.ShouldBind(newQuery); err != nil {
		logrus.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	var formSql model.FormSql
	if tx := dao.DB.Model(&model.FormSql{}).Where("form_id = ?", newQuery.FormID).Update("status", newQuery.Status).Find(&formSql); tx.Error != nil {
		logrus.Error(tx.Error)
		c.JSON(http.StatusOK, gin.H{
			"msg": tx.Error,
		})
		return
	}

	logrus.Info(fmt.Sprintf("Done Form: %#v", formSql))
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
		"res": formSql,
	})
}

func (api *AdminFormApi) GetUnreviewed(c *gin.Context) {
	var formSqls []model.FormSql
	if tx := dao.DB.Where("status = ?", "0").Find(&formSqls); tx.Error != nil {
		logrus.Error(tx.Error)
		c.JSON(http.StatusOK, gin.H{
			"msg": tx.Error,
		})
		return
	}

	logrus.Info(fmt.Sprintf("Get Todo: %#v", formSqls))
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
		"res": formSqls,
	})
}

func (api *AdminFormApi) GetReviewed(c *gin.Context) {
	var formSqls []model.FormSql
	if tx := dao.DB.Where("status = ?", "1").Find(&formSqls); tx.Error != nil {
		logrus.Error(tx.Error)
		c.JSON(http.StatusOK, gin.H{
			"msg": tx.Error,
		})
		return
	}

	logrus.Info(fmt.Sprintf("Get Todo: %#v", formSqls))
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
		"res": formSqls,
	})
}

func (api *AdminFormApi) GetCannot(c *gin.Context) {
	var formSqls []model.FormSql
	if tx := dao.DB.Where("status = ?", "2").Find(&formSqls); tx.Error != nil {
		logrus.Error(tx.Error)
		c.JSON(http.StatusOK, gin.H{
			"msg": tx.Error,
		})
		return
	}

	logrus.Info(fmt.Sprintf("Get Todo: %#v", formSqls))
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
		"res": formSqls,
	})
}

func (api *AdminFormApi) GetDone(c *gin.Context) {
	var formSqls []model.FormSql
	if tx := dao.DB.Where("status = ?", "3").Find(&formSqls); tx.Error != nil {
		logrus.Error(tx.Error)
		c.JSON(http.StatusOK, gin.H{
			"msg": tx.Error,
		})
		return
	}

	logrus.Info(fmt.Sprintf("Get Todo: %#v", formSqls))
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
		"res": formSqls,
	})
}
