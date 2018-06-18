package model

import (
	"github.com/indraoct/go-guru/abstract"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

func Tests(db *gorm.DB,filter map[string]int) abstract.Get_Tests{
	var tests abstract.Tests
	var arr_tests []abstract.Tests
	var getTests abstract.Get_Tests
	filter_arr := make(map[string]int)

	if(len(filter) > 0){
		rows,err := db.Raw("SELECT id,t_description,positif_weight,negatif_weight,created_date,updated_date from tests WHERE id=?",filter["id"]).Rows()
		defer rows.Close()
		if(err != nil){
			getTests.Count = 0
			getTests.Status = 0
			getTests.Message = err.Error()
		}
		for rows.Next() {
			rows.Scan(&tests.Id, &tests.T_description, &tests.Positif_weight,&tests.Negatif_weight,&tests.Created_date,&tests.Updated_date)

			filter_arr["id_test"]= tests.Id
			tests.Questions = Questions(db,filter_arr).Data
			arr_tests = append(arr_tests,tests)
		}

		getTests.Data = arr_tests
		getTests.Count = len(arr_tests)


	}else {
		rows,err := db.Raw("SELECT id,t_description,positif_weight,negatif_weight,created_date,updated_date from tests").Rows()
		defer rows.Close()

		if(err != nil){
			log.Fatal(err)
		}
		for rows.Next() {
			rows.Scan(&tests.Id, &tests.T_description, &tests.Positif_weight,&tests.Negatif_weight,&tests.Created_date,&tests.Updated_date)
			filter_arr["id_test"] = tests.Id
			tests.Questions = Questions(db,filter_arr).Data
			arr_tests = append(arr_tests,tests)
		}



		getTests.Data = arr_tests
		getTests.Count = len(arr_tests)
	}


	return getTests

}

func Questions(db *gorm.DB,filter map[string]int) abstract.Get_Questions{
	var questions abstract.Questions
	var arr_questions []abstract.Questions
	var getQuestions abstract.Get_Questions

	if id_test,exist :=filter["id_test"];exist{
		rows,err := db.Raw("SELECT id,id_test,q_number,q_description,option_a,option_b,option_c,option_d,created_date,updated_date from questions WHERE id_test=?",id_test).Rows()
		defer rows.Close()
		if(err != nil){
			getQuestions.Count = 0
			getQuestions.Status = 0
			getQuestions.Message = err.Error()
		}else {
			for rows.Next() {
				rows.Scan(&questions.Id, &questions.Id_test, &questions.Q_number, &questions.Q_description, &questions.Option_a, &questions.Option_b, &questions.Option_c, &questions.Option_d, &questions.Created_date, &questions.Updated_date)
				arr_questions = append(arr_questions, questions)

			}
		}

		getQuestions.Data = arr_questions
		getQuestions.Count = len(arr_questions)

	}else if id,exist :=filter["id"];exist{
		rows,err := db.Raw("SELECT id,id_test,q_number,q_description,option_a,option_b,option_c,option_d,created_date,updated_date from questions WHERE id=?",id).Rows()
		defer rows.Close()
		if(err != nil){
			getQuestions.Count = 0
			getQuestions.Status = 0
			getQuestions.Message = err.Error()
		}else {
			for rows.Next() {
				rows.Scan(&questions.Id, &questions.Id_test, &questions.Q_number, &questions.Q_description, &questions.Option_a, &questions.Option_b, &questions.Option_c, &questions.Option_d, &questions.Created_date, &questions.Updated_date)
				arr_questions = append(arr_questions, questions)

			}
		}

		getQuestions.Data = arr_questions
		getQuestions.Count = len(arr_questions)

	}else {
		rows,err := db.Raw("SELECT id,id_test,q_number,q_description,option_a,option_b,option_c,option_d,created_date,updated_date from questions").Rows()
		defer rows.Close()
		if(err != nil){
			getQuestions.Count = 0
			getQuestions.Status = 0
			getQuestions.Message = err.Error()
		}else {
			for rows.Next() {
				rows.Scan(&questions.Id, &questions.Id_test, &questions.Q_number, &questions.Q_description, &questions.Option_a, &questions.Option_b, &questions.Option_c, &questions.Option_d, &questions.Created_date, &questions.Updated_date)
				arr_questions = append(arr_questions, questions)

			}
		}


		getQuestions.Data = arr_questions
		getQuestions.Count = len(arr_questions)
	}

	return getQuestions
}

func Answers(db *gorm.DB,filter map[string]int) abstract.Get_Answers{
	var answers abstract.Answers
	var arr_answers []abstract.Answers
	var getAnswers abstract.Get_Answers

	if id_question,exist :=filter["id_question"];exist{
		rows,err := db.Raw("SELECT id,id_question,option_a,option_b,option_c,option_d from answers WHERE id=?",id_question).Rows()
		defer rows.Close()
		if(err != nil){
			getAnswers.Count = 0
			getAnswers.Status = 0
			getAnswers.Message = err.Error()
		}
		for rows.Next() {
			rows.Scan(&answers.Id, &answers.Id_question, &answers.Option_a,&answers.Option_b,&answers.Option_c,&answers.Option_d)
			arr_answers = append(arr_answers,answers)
		}

		getAnswers.Data = arr_answers
		getAnswers.Count = len(arr_answers)


	}else{
		rows,err := db.Raw("SELECT id,id_question,option_a,option_b,option_c,option_d from answers").Rows()
		defer rows.Close()

		if(err != nil){
			log.Fatal(err)
		}
		for rows.Next() {
			rows.Scan(&answers.Id, &answers.Id_question, &answers.Option_a,&answers.Option_b,&answers.Option_c,&answers.Option_d)
			arr_answers = append(arr_answers,answers)
		}

		getAnswers.Data = arr_answers
		getAnswers.Count = len(arr_answers)
	}

	return getAnswers
}

func Students(db *gorm.DB,filter map[string]int) abstract.Get_Students{
	var students abstract.Students
	var arr_students []abstract.Students
	var getStudents abstract.Get_Students

	if(len(filter) > 0){
		rows,err := db.Raw("SELECT id,name,email,password from students WHERE id=?",filter["id"]).Rows()
		defer rows.Close()
		if(err != nil){
			getStudents.Count = 0
			getStudents.Status = 0
			getStudents.Message = err.Error()
		}
		for rows.Next() {
			rows.Scan(&students.Id, &students.Name, &students.Email,&students.Password)
			arr_students = append(arr_students,students)
		}

		getStudents.Data = arr_students
		getStudents.Count = len(arr_students)


	}else {
		rows,err := db.Raw("SELECT id,name,email,password from students").Rows()
		defer rows.Close()

		if(err != nil){
			log.Fatal(err)
		}
		for rows.Next() {
			rows.Scan(&students.Id, &students.Name, &students.Email,&students.Password)
			arr_students = append(arr_students,students)
		}

		getStudents.Data = arr_students
		getStudents.Count = len(arr_students)
	}


	return getStudents

}

func CreateTest(db *gorm.DB,dataMap map[string]interface{}) abstract.Create_Test{
	var createTest abstract.Create_Test
	var tests abstract.Tests

	createTest.Status = 1
	createTest.Message = "Success"


	t_description := dataMap["t_description"].(string)
	positif_weight := dataMap["positif_weight"].(int)
	negatif_weight := dataMap["negatif_weight"].(int)
	currentdatetime := time.Now().Format("2006-01-02 15:04:05")
	tests = abstract.Tests{T_description:t_description,Positif_weight:positif_weight,Negatif_weight:negatif_weight,Created_date:currentdatetime}

	tx := db.Begin()

	if err := tx.Create(&tests).Error; err != nil {
		tx.Rollback()
		createTest.Status = 0
		createTest.Message = err.Error()
	}
	tx.Commit()

	return createTest

}

func CreateQuestion(db *gorm.DB,dataMap map[string]interface{}) abstract.Create_Question{
	var createQuestion abstract.Create_Question
	var questions abstract.Questions

	createQuestion.Status = 1
	createQuestion.Message = "Success"

	q_number := dataMap["q_number"].(int)
	q_description := dataMap["q_description"].(string)
	id_test  := dataMap["id_test"].(int)
	option_a := dataMap["option_a"].(string)
	option_b := dataMap["option_b"].(string)
	option_c := dataMap["option_c"].(string)
	option_d := dataMap["option_d"].(string)
	currentdatetime := time.Now().Format("2006-01-02 15:04:05")
	questions = abstract.Questions{Id_test:id_test,Q_number:q_number,Q_description:q_description,Option_a:option_a,Option_b:option_b,Option_c:option_c,Option_d:option_d,Created_date:currentdatetime}

	tx := db.Begin()

	if err := tx.Create(&questions).Error; err != nil {
		tx.Rollback()
		createQuestion.Status = 0
		createQuestion.Message = err.Error()
	}
	tx.Commit()

	return createQuestion

}

func CreateAnswer(db *gorm.DB,dataMap map[string]interface{}) abstract.Create_Answer{
	var createAnswer abstract.Create_Answer
	var answers abstract.Answers

	createAnswer.Status = 1
	createAnswer.Message = "Success"

	id_question  := dataMap["id_question"].(int)
	option_a := dataMap["option_a"].(int)
	option_b := dataMap["option_b"].(int)
	option_c := dataMap["option_c"].(int)
	option_d := dataMap["option_d"].(int)

	answers = abstract.Answers{Id_question:id_question,Option_a:option_a,Option_b:option_b,Option_c:option_c,Option_d:option_d}

	tx := db.Begin()

	if err := tx.Create(&answers).Error; err != nil {
		tx.Rollback()
		createAnswer.Status = 0
		createAnswer.Message = err.Error()
	}
	tx.Commit()

	return createAnswer

}

func StudentTests(db *gorm.DB,filter map[string]int) abstract.Get_Student_Tests{
	var student_tests abstract.Student_Tests
	var arr_student_test []abstract.Student_Tests
	var getStudentTest abstract.Get_Student_Tests

	exist_id_test := false
	exist_id_student := false

	if id_test,existid_test :=filter["id_test"];existid_test{
		exist_id_test = existid_test
		filter["id_test"] = id_test
	}

	if id_student,existid_student :=filter["id_student"];existid_student{
		exist_id_student = existid_student
		filter["id_student"] = id_student
	}

	if id,exist :=filter["id"];exist{
		rows,err := db.Raw("SELECT id,id_student,id_test,start_time,end_time,student_score,status from student_tests WHERE id=?",id).Rows()
		defer rows.Close()
		if(err != nil){
			getStudentTest.Count = 0
			getStudentTest.Status = 0
			getStudentTest.Message = err.Error()
		}
		for rows.Next() {
			rows.Scan(&student_tests.Id, &student_tests.Id_student, &student_tests.Id_test,&student_tests.Start_time,&student_tests.End_time,&student_tests.Student_score,&student_tests.Status)
			arr_student_test = append(arr_student_test,student_tests)
		}

		getStudentTest.Data = arr_student_test
		getStudentTest.Count = len(arr_student_test)


	}else if(exist_id_student == true && exist_id_test == true){

		db.Where("id_test = ? AND id_student = ?",filter["id_test"],filter["id_student"]).Find(&student_tests).Scan(&arr_student_test)

		getStudentTest.Data = arr_student_test
		getStudentTest.Count = len(arr_student_test)


	}else {
		rows,err := db.Raw("SELECT id,id_student,id_test,start_time,end_time,student_score,status from student_tests").Rows()
		defer rows.Close()

		if(err != nil){
			log.Fatal(err)
		}
		for rows.Next() {
			rows.Scan(&student_tests.Id, &student_tests.Id_student, &student_tests.Id_test,&student_tests.Start_time,&student_tests.End_time,&student_tests.Student_score,&student_tests.Status)
			arr_student_test = append(arr_student_test,student_tests)
		}

		getStudentTest.Data = arr_student_test
		getStudentTest.Count = len(arr_student_test)
	}


	return getStudentTest


}

func StudentAnswers(db *gorm.DB,filter map[string]int) abstract.Get_Student_Answers{
	var student_answers abstract.Student_Answers
	var arr_student_answers []abstract.Student_Answers
	var getStudentAnswers abstract.Get_Student_Answers

	exist_id_test := false
	exist_id_student := false
	exist_id_question := false

	if id_test,existid_test :=filter["id_test"];existid_test{
		exist_id_test = existid_test
		filter["id_test"] = id_test
	}

	if id_student,existid_student :=filter["id_student"];existid_student{
		exist_id_student = existid_student
		filter["id_student"] = id_student
	}

	if id_question,existid_question :=filter["id_question"];existid_question{
		exist_id_question = existid_question
		filter["id_question"] = id_question
	}


	if id,exist :=filter["id"];exist{
		db.Raw("SELECT id,id_student,id_test,id_question,option_a,option_b,option_c,option_d,created_date,updated_date from student_answers WHERE id=?",id).Scan(arr_student_answers)
		getStudentAnswers.Data = arr_student_answers
		getStudentAnswers.Count = len(arr_student_answers)


	}else if(exist_id_student == true && exist_id_test == true && exist_id_question == true){
		db.Where("id_test = ? AND id_student = ? AND id_question =?",filter["id_test"],filter["id_student"],filter["id_question"]).Find(&student_answers).Scan(&arr_student_answers)
		getStudentAnswers.Data = arr_student_answers
		getStudentAnswers.Count = len(arr_student_answers)


	}else if(exist_id_student == true && exist_id_test == true && exist_id_question == false){
		db.Where("id_test = ? AND id_student = ?",filter["id_test"],filter["id_student"]).Find(&student_answers).Scan(&arr_student_answers)
		getStudentAnswers.Data = arr_student_answers
		getStudentAnswers.Count = len(arr_student_answers)

	}else {
		rows,err := db.Raw("SELECT id,id_student,id_test,id_question,option_a,option_b,option_c,option_d,created_date,updated_date from student_answers").Rows()
		defer rows.Close()

		if(err != nil){
			log.Fatal(err)
		}
		for rows.Next() {
			rows.Scan(&student_answers.Id_student, &student_answers.Id_test,&student_answers.Id_question,&student_answers.Option_a,&student_answers.Option_b,&student_answers.Option_c,&student_answers.Option_d,&student_answers.Created_date,&student_answers.Updated_date)
			arr_student_answers = append(arr_student_answers,student_answers)
		}

		getStudentAnswers.Data = arr_student_answers
		getStudentAnswers.Count = len(arr_student_answers)
	}


	return getStudentAnswers


}

func StartStudentTest(db *gorm.DB,dataMap map[string]interface{}) abstract.Start_Student_Test{
	var startStudentTest abstract.Start_Student_Test
	var studentTest abstract.Student_Tests

	startStudentTest.Status = 1
	startStudentTest.Message = "Success"

	id_student := dataMap["id_student"].(int)
	id_test := dataMap["id_test"].(int)
	currentdatetime := time.Now().Format("2006-01-02 15:04:05")

	studentTest = abstract.Student_Tests{Id_student:id_student,Id_test:id_test,Start_time:currentdatetime,Student_score:0,Status:1}

	tx := db.Begin()

	if err := tx.Create(&studentTest).Error; err != nil {
		tx.Rollback()
		startStudentTest.Status = 0
		startStudentTest.Message = err.Error()
	}
	tx.Commit()

	return startStudentTest
}

func InputStudentAnswer(db *gorm.DB,dataMap map[string]interface{}) abstract.Input_Student_Answers{
	var inputStudentAnswers abstract.Input_Student_Answers
	var studentAnswers abstract.Student_Answers

	inputStudentAnswers.Status = 1
	inputStudentAnswers.Message = "Success"

	id_student := dataMap["id_student"].(int)
	id_test := dataMap["id_test"].(int)
	id_question := dataMap["id_question"].(int)
	option_a := dataMap["option_a"].(int)
	option_b := dataMap["option_b"].(int)
	option_c := dataMap["option_c"].(int)
	option_d := dataMap["option_d"].(int)
	currentdatetime := time.Now().Format("2006-01-02 15:04:05")

	studentAnswers = abstract.Student_Answers{Id_student:id_student,Id_test:id_test,Id_question:id_question,Option_a:option_a,Option_b:option_b,Option_c:option_c,Option_d:option_d,Created_date:currentdatetime}

	tx := db.Begin()

	if err := tx.Create(&studentAnswers).Error; err != nil {
		tx.Rollback()
		inputStudentAnswers.Status = 0
		inputStudentAnswers.Message = err.Error()
	}
	tx.Commit()

	return inputStudentAnswers
}

func CompletionTest(db *gorm.DB,dataMap map[string]interface{}) abstract.Completion_Tests{
	var completionTests abstract.Completion_Tests
	var studentTests abstract.Student_Tests

	completionTests.Status = 1
	completionTests.Message = "Success"

	id_student := dataMap["id_student"].(int)
	id_test := dataMap["id_test"].(int)
	student_score := dataMap["student_score"].(int)
	totalQuestion := dataMap["total_question"].(int)
	correctAnswer := dataMap["correct_answer"].(int)
	wrongAnswer := dataMap["wrong_answer"].(int)
	currentdatetime := time.Now().Format("2006-01-02 15:04:05")


	studentTests = abstract.Student_Tests{End_time:currentdatetime,Student_score:student_score,Status:2,Total_question:totalQuestion,Correct_answer:correctAnswer,Wrong_answer:wrongAnswer}

	db.Table("student_tests").Where("id_test = ? AND id_student=?",id_test,id_student).Updates(studentTests)

	return completionTests


}