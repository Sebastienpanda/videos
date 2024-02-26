package controllers

import (
	"books/initializers"
	"books/models"
	"books/utils"
	"books/validators"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

func VideosCreate(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title   string
		Content string
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validators.ValidateVideos(body.Title, body.Content)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	titleUppercaseLetters := utils.ToTitleCase(body.Title)

	video := models.Video{Title: titleUppercaseLetters, Content: body.Content}

	result := initializers.DB.Create(&video)

	if result.Error != nil {
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"video":   video,
		"success": true,
	})
}

func VideosIndex(w http.ResponseWriter, r *http.Request) {
	var videos []models.Video
	initializers.DB.Find(&videos)

	json.NewEncoder(w).Encode(videos)
}

func VideosShow(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]
	var video models.Video
	dbResult := initializers.DB.First(&video, id)

	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":   "Record not found",
			"success": false,
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"video":   video,
		"success": true,
	})
}

func VideosUpdate(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	var body struct {
		Title   string
		Content string
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var video models.Video
	dbResult := initializers.DB.First(&video, id)
	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":   "Record not found",
			"success": false,
		})
		return
	}

	initializers.DB.Model(&video).Updates(models.Video{Title: body.Title, Content: body.Content})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"video":   video,
		"success": true,
	})
}
func VideosDelete(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	dbResult := initializers.DB.Delete(&models.Video{}, id)

	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":   "Record not found",
			"success": false,
		})
		return
	}
}
