package controllers

import (
	"diary_api/config"
	"diary_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateEntryInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func GetEntries(c *gin.Context) {
	var entries []models.Entry

	result := config.DB.Find(&entries)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	if len(entries) == 0 {
		c.JSON(404, gin.H{"message": "No entries found"})
		return
	}
	c.JSON(http.StatusOK, entries)
}

func CreateEntry(c *gin.Context) {
	var input CreateEntryInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entry := models.Entry{Title: input.Title, Content: input.Content}
	result := config.DB.Create(&entry)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	// c.JSON(http.StatusOK, gin.H{"entry": entry})
	c.JSON(http.StatusCreated, entry)
}

func GetEntry(c *gin.Context) {
	entryId := c.Param("id")
	var entry models.Entry

	result := config.DB.First(&entry, entryId)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, entry)
}

func DeleteEntry(c *gin.Context) {
	entryId := c.Param("id")
	var entry models.Entry

	result := config.DB.First(&entry, entryId)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	result = config.DB.Delete(&entry)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Entry deleted successfully"})
}
