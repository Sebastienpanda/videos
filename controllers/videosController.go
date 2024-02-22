package controllers

import (
	"books/initializers"
	"books/models"
	"books/utils"
	"books/validators"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func VideosCreate(c *gin.Context) {

	var body struct {
		Title   string
		Content string
	}

	c.Bind(&body)

	err := validators.ValidateVideos(body.Title, body.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}
	titleUppercaseLetters := utils.ToTitleCase(body.Title)

	video := models.Video{Title: titleUppercaseLetters, Content: body.Content}

	result := initializers.DB.Create(&video)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"video":   video,
		"success": true,
	})
}

func VideosIndex(c *gin.Context) {
	var Videos []models.Video
	initializers.DB.Find(&Videos)

	c.JSON(http.StatusOK, gin.H{
		"videos":  Videos,
		"success": true,
	})
}

func VideosShow(c *gin.Context) {
	id := c.Param("id")
	var video models.Video

	dbResult := initializers.DB.First(&video, id)

	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Record not found",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"video":   video,
		"success": true,
	})
}

func VideosUpdate(c *gin.Context) {
	var body struct {
		Title   string
		Content string
	}

	c.Bind(&body)

	var video models.Video
	dbResult := initializers.DB.First(&video, c.Param("id"))
	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Record not found",
			"success": false,
		})
		return
	}

	initializers.DB.Model(&video).Updates(models.Video{Title: body.Title, Content: body.Content})

	c.JSON(http.StatusOK, gin.H{
		"video":   video,
		"success": true,
	})
}

func VideosDelete(c *gin.Context) {
	id := c.Param("id")

	dbResult := initializers.DB.Delete(&models.Video{}, id)

	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Record not found",
			"success": false,
		})
		return
	}

	c.Status(http.StatusOK)
}
