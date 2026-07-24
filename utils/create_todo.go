package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) CreateTodo(c *gin.Context) {
	var todo Todo
	err := c.BindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, ok := c.Get("user_id")
	if !ok {
		fmt.Println("c.Get user_id failed, ok =", ok)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userIdFloat, ok := userId.(float64)
	if !ok {
		fmt.Printf("type assertion failed, userId = %v, type = %T\n", userId, userId)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	todo.UserID = int(userIdFloat)

	if todo.Deadline == "" {
		err = a.DB.QueryRow("INSERT INTO todos (title, description, user_id) VALUES ($1, $2, $3) RETURNING id, status", todo.Title, todo.Description, todo.UserID).Scan(&todo.ID, &todo.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusAccepted, todo)
		return
	} else {
		err = a.DB.QueryRow("INSERT INTO todos (title, description, deadline, user_id) VALUES ($1, $2, $3, $4) RETURNING id, status", todo.Title, todo.Description, todo.Deadline, todo.UserID).Scan(&todo.ID, &todo.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, todo)
}
