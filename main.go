package main

import (
	_ "github.com/mattn/go-sqlite3"
	"log"
	"github.com/gin-gonic/gin"
	"time"
	"github.com/indraoct/go-guru/handler"
	"github.com/jinzhu/gorm"
)

func initDB () *gorm.DB{
	//initiate DB Connection
	db, err := gorm.Open("sqlite3", "./ruangguru.db")
	db.LogMode(true)

	if(err != nil){
		log.Fatal(err)
	}

	//connection DB Pooling
	db.DB().SetMaxIdleConns(0)
	db.DB().SetMaxOpenConns(5)
	db.DB().SetConnMaxLifetime(time.Minute*5)

	return db

}

/**
* Allow all subdomain
*/
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main(){

	db := initDB()
	event := gin.Default()
	event.Use(Cors())

	data := event.Group("data")
	{

		data.GET("/gettests",handler.GetTests(db))
		data.GET("/getquestions",handler.GetQuestions(db))

	}

	post := event.Group("post")
	{

		post.POST("/createtest",handler.CreateTest(db))
		post.POST("/createquestion",handler.CreateQuestion(db))
		post.POST("/createanswer",handler.CreateAnswer(db))

		post.POST("/starttest",handler.StartTest(db))
		post.POST("/inputanswer",handler.InputStudentAnswer(db))
		post.POST("/completiontest",handler.CompletionTest(db))
		post.POST("/gettestinsight",handler.StudentgetInsight(db))

	}

	delete := event.Group("delete")
	{
		delete.DELETE("/test")
		delete.DELETE("/question")
	}

	update := event.Group("update")
	{
		update.PUT("/modifytest",handler.ModifyTest(db))
		update.PUT("/modifyquestion",handler.ModifyQuestion(db))

	}

	event.Run(":8888")

}