package api

import (
	"fmt"
	"net/http"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/isther/clinic/conf"
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

	newForm.FormID = uuid.New().String()
	newForm.Status = "0"

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
		"msg":     "ok",
		"form_id": newForm.FormID,
	})
}

func (api *UserFormApi) Query(c *gin.Context) {
	newQuery := &struct {
		StudentID string `json:"student_id"`
	}{}

	if err := c.ShouldBind(newQuery); err != nil {
		logrus.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	var formSqls []model.FormSql
	if tx := dao.DB.Where("student_id = ?", newQuery.StudentID).Find(&formSqls); tx.Error != nil {
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

func (api *UserFormApi) Upload(c *gin.Context) {
	form, _ := c.MultipartForm()

	var (
		err     error
		imgs    model.Strs
		form_id = form.Value["form_id"]
		files   = form.File["upload[]"]
	)

	for _, file := range files {
		imgName := uuid.New().String() + path.Ext(file.Filename)
		err = c.SaveUploadedFile(file, filepath.Join(conf.Server.ImgDir, imgName))
		if err == nil {
			imgs = append(imgs, imgName)
		}
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	var formSql model.FormSql
	if tx := dao.DB.Where("form_id = ?", form_id).First(&formSql); tx.Error != nil {
		logrus.Error(tx.Error)
		c.JSON(http.StatusOK, gin.H{
			"msg": tx.Error,
		})
		return
	}

	imgs = append(formSql.Imgs, imgs...)
	if tx := dao.DB.Model(&model.FormSql{}).Where("form_id = ?", form_id).Update("imgs", imgs).Find(&formSql); tx.Error != nil {
		logrus.Error(tx.Error)
		c.JSON(http.StatusOK, gin.H{
			"msg": tx.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  fmt.Sprintf("%d files uploaded!\n", len(imgs)),
		"form": formSql,
	})
}
