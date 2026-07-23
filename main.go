package main

import (
	"Task-Tracker-V2/utils"
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", os.Getenv("PSGX_DB_CONN"))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	defer db.Close()

	app := &utils.App{DB: db}

	router := gin.Default()
	router.POST("/register", app.RegisterUser)
	router.Run(":8080")
}

// func (a *utils.App) HelloHandler(c *gin.Context) {
// 	c.String(http.StatusOK, "jello")
// }

// func (a *utils.App) ByeHandler(c *gin.Context) {
// 	c.String(http.StatusOK, "bye")
// }

// func (a *App) CountHandler(c *gin.Context) {
// 	var count int
// 	err := a.DB.QueryRow("SELECT COUNT(*) FROM todos").Scan(&count)
// 	if err != nil {
// 		c.String(http.StatusInternalServerError, "failed to count todos")
// 		return
// 	}
// 	c.String(http.StatusOK, fmt.Sprintf("count: %d", count))
// }

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		panic(err)
// 	}

// 	db, err := sql.Open("postgres", os.Getenv("PSGX_DB_CONN"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Connected to database")

// 	app := &App{DB: db}

// 	router := gin.Default()
// 	router.GET("/hello", app.HelloHandler)
// 	router.GET("/bye", app.ByeHandler)
// 	router.GET("/count", app.CountHandler)
// 	router.Run()

// }
