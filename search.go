package main

import (
	"database/sql"
	"fmt"
	"os"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)
type question struct {
	ID int
	question_name string
	category_id int
	topic_id int
	subject_id int
	subcategory_id int
	
}
func main() {
	err := godotenv.Load(".env")

	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
	dbHost := os.Getenv("DB_HOST")
	fmt.Println("DB String: ", dbHost)



	db, err := sql.Open("mysql", dbHost)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to MySQL")

	// var questionsuggestion = getAllQuestion(db);
	var keyword = "vector"
	var categoryID = 1
	var topicID = 1
	var subcategoryID = 1
	var subjectID = 1
	var questionsuggestion = getQuestionByKeyword(db,keyword,categoryID,topicID,subcategoryID,subjectID);
	fmt.Println("questionsuggestion",questionsuggestion)



	defer db.Close()
}

func getAllQuestion(db *sql.DB) (suggestion []question) {
		// select
		results, err := db.Query("SELECT ID,question_name,category_id,topic_id,subject_id,subcategory_id FROM esdb.question limit 100")
		// defer results.Close()
		if err != nil {
			panic(err.Error())
		}

		
		for results.Next() {
			var Question question
			err = results.Scan(&Question.ID, &Question.question_name, &Question.category_id,&Question.topic_id,&Question.subject_id,&Question.subcategory_id)
			if err != nil {
				panic(err.Error())
			}
			fmt.Println("--------------Question---------------",Question.ID,Question.question_name,Question.category_id,Question.topic_id,Question.subject_id,Question.subcategory_id)
			suggestion = append(suggestion, Question)
		}
		return suggestion
	
}

func getQuestionByKeyword(db *sql.DB,Keyword string, categoryID int,topicID int,subcategoryID int,subjectID int) (suggestion []question) {
	// select
	// results, err := db.Query("SELECT ID,question_name,category_id,topic_id,subject_id,subcategory_id FROM esdb.question  where question_name like ?","%"+Keyword+"%" )
	results, err := db.Query("SELECT ID,question_name,category_id,topic_id,subject_id,subcategory_id FROM esdb.question  where MATCH(question_name) AGAINST('vector') and (category_id=? OR topic_id=? OR subject_id=? OR subcategory_id=?)",categoryID,topicID,subjectID,subcategoryID)
	// defer results.Close()
	if err != nil {
		panic(err.Error())
	}

	
	for results.Next() {
		var Question question
		err = results.Scan(&Question.ID, &Question.question_name, &Question.category_id,&Question.topic_id,&Question.subject_id,&Question.subcategory_id)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("--------------Question---------------",Question.ID,Question.question_name,Question.category_id,Question.topic_id,Question.subject_id,Question.subcategory_id)
		suggestion = append(suggestion, Question)
	}
	return suggestion

}


