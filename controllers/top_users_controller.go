package controllers

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/quizzup/db"
	"github.com/quizzup/models"
)

// GetTopUsers s
func GetTopUsers(c *gin.Context) {
	// var users []models.User

	var results []models.Result

	if err := db.Get().Find(&results).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err})
	}
	resultsInOrder := make(map[uint]int)
	for _, result := range results {
		if _, ok := resultsInOrder[result.UserID]; ok {
			resultsInOrder[result.UserID] += result.Score
		} else {
			resultsInOrder[result.UserID] = result.Score
		}
	}
	n := map[int][]uint{}
	var a []int
	for k, v := range resultsInOrder {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}

	topUsers := []models.TopUser{}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for _, k := range a {
		for _, s := range n[k] {
			var user models.User
			if error := db.Get().Where("id = ?", s).Find(&user).Error; error != nil {
				c.JSON(http.StatusOK, gin.H{"message": error})
			}
			topUser := models.TopUser{}
			topUser.User = user
			topUser.Score = k
			topUsers = append(topUsers, topUser)
		}
	}

	fmt.Println(" ==== results in order ", resultsInOrder)
	fmt.Println(" ==== results in order ", topUsers)

	if len(resultsInOrder) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No Top Users yet"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"top_users": topUsers})
}

// GetTopUsersOfQuiz s
func GetTopUsersOfQuiz(c *gin.Context) {
	id := c.Param("id")
	var quiz models.Quiz
	if err := db.Get().Preload("Results").Where("id = ?", id).First(&quiz).Order("score asc").Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	finalResult := []models.TopUser{}
	for _, result := range quiz.Results {
		tUser := models.TopUser{}
		tUser.Score = result.Score
		var user models.User
		if err := db.Get().Where("id = ?", result.UserID).Find(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}
		tUser.User = user
		finalResult = append(finalResult, tUser)
	}

	c.JSON(http.StatusOK, gin.H{"top_users": finalResult})
}
