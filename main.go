package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Foot struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	dbHost := os.Getenv("DB_HOST")
	fmt.Println(dbHost)

	db, err := gorm.Open(mysql.Open(os.Getenv("DB_DSN")), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(db)

	item := Foot{
		Id:   1,
		Name: "Foot1",
	}
	jsonData, _ := json.Marshal(item)
	fmt.Println(string(jsonData))

	jsonString := `{"id":2,"name":"Foot 2","created_at":null,"updated_at":null}`

	var item2 Foot

	if err := json.Unmarshal([]byte(jsonString), &item2); err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(item2)

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": item2,
	// 	})
	// })
	// r.Run()
}
