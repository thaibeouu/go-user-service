package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

//Person with id and name
type Person struct {
	ID   int
	Name string
}

var db = initDb()

func getTime(c *gin.Context) {
	currentTime := time.Now().Format("02/01/2006 15:04:05")
	content := gin.H{"time": currentTime}
	c.JSON(200, content)
}

func insert(c *gin.Context) {
	name := c.PostForm("name")
	dbIns, err := db.Prepare("insert into people (name) values(?);")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = dbIns.Exec(name)

	if err != nil {
		fmt.Print(err.Error())
	}
	// dbIns.Close()
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf(" %s created", name),
	})
}
func index(c *gin.Context) {
	var (
		person Person
		people []Person
	)
	rows, err := db.Query("select * from people")
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&person.ID, &person.Name)
		people = append(people, person)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	// rows.Close()
	c.JSON(http.StatusOK, people)
}

func initDb() *sql.DB {
	db, err := sql.Open("mysql", "root:vinhhien@tcp(127.0.0.1:3306)/service")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	app := gin.Default()
	app.GET("/time", getTime)
	app.POST("/insert", insert)
	app.GET("/", index)
	fmt.Println(app.Run(":8000"))
}
