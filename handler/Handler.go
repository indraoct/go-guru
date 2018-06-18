package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/indraoct/go-guru/abstract"
	"github.com/indraoct/go-guru/model"
	"github.com/jinzhu/gorm"
	"regexp"
	"log"
	"strconv"
)

func GetTests(db *gorm.DB) gin.HandlerFunc{
	return func(ctx *gin.Context){
		var getTests  abstract.Get_Tests

		filter := make(map[string]int)

		rtn_tests := model.Tests(db,filter)
		getTests.Data = rtn_tests.Data
		getTests.Status = 1
		getTests.Message = "Success"
		getTests.Count = rtn_tests.Count

		ctx.JSON(200, getTests)
	}

}

func GetQuestions(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var getQuestions  abstract.Get_Questions

		filter := make(map[string]int)

		rtn_questions := model.Questions(db,filter)
		getQuestions.Data = rtn_questions.Data
		getQuestions.Status = 1
		getQuestions.Message = "Success"
		getQuestions.Count = rtn_questions.Count

		ctx.JSON(200, getQuestions)

	}
}


func CreateTest(db *gorm.DB) gin.HandlerFunc{
	return func(ctx *gin.Context) {

		regString, err := regexp.Compile("[^A-Za-z0-9 ]+")
		var createTests abstract.Create_Test
		var positifWeight int
		var negatifWeight int

		if(err != nil){
			log.Fatal(err)
		}

		testDescription := regString.ReplaceAllString(ctx.PostForm("test_description")," ")
		positifWeight,_ = strconv.Atoi(ctx.PostForm("positif_weight"))
		negatifWeight ,_ = strconv.Atoi(ctx.PostForm("negatif_weight"))


		createTests.Status = 0
		if(testDescription == "" || positifWeight == 0){
			createTests.Message = "test_description can not be empty, positif weight can not be zero or empty"

		}else{
			if(negatifWeight > positifWeight){
				createTests.Message = "Negatif weight can not greater than Positif weight"

			}else{
				dataMap := make(map[string]interface{})
				dataMap["t_description"] = testDescription
				dataMap["positif_weight"] = positifWeight
				dataMap["negatif_weight"] = negatifWeight

				rtn_create_test := model.CreateTest(db,dataMap)
				createTests.Status = rtn_create_test.Status
				createTests.Message = rtn_create_test.Message
			}

		}

		ctx.JSON(200, createTests)


	}

}

func CreateQuestion(db *gorm.DB) gin.HandlerFunc{
	return func(ctx *gin.Context) {

		regString, err := regexp.Compile("[^A-Za-z0-9-+/:;*,.'\"@#$&!? ]+")
		var createTests abstract.Create_Test
		var id_test,q_number int
		var option_a,option_b,option_c,option_d string

		if(err != nil){
			log.Fatal(err)
		}

		qDescription := regString.ReplaceAllString(ctx.PostForm("q_description")," ")
		id_test	,_ = strconv.Atoi(ctx.PostForm("id_test"))
		q_number,_ = strconv.Atoi(ctx.PostForm("q_number"))
		option_a = ctx.PostForm("option_a")
		option_b = ctx.PostForm("option_b")
		option_c = ctx.PostForm("option_c")
		option_d = ctx.PostForm("option_d")

		createTests.Status = 0
		if(qDescription == "") {
			createTests.Message = "qustion desc, option a, option b, option c, option d are mandatory"
		}else if(q_number == 0 || id_test == 0){
			createTests.Message = "id test and questions number can not be empty or zero"
		}else if(isIDTestExist(db,id_test) == false){
			createTests.Message = "wrong id test!!"
		}else{
				dataMap := make(map[string]interface{})
				dataMap["q_description"] = qDescription
				dataMap["id_test"] =id_test
				dataMap["q_number"] = q_number
				dataMap["option_a"] = option_a
				dataMap["option_b"] = option_b
				dataMap["option_c"] = option_c
				dataMap["option_d"] = option_d

				rtn_create_test := model.CreateQuestion(db,dataMap)
				createTests.Status = rtn_create_test.Status
				createTests.Message = rtn_create_test.Message

		}

		ctx.JSON(200, createTests)

	}

}

func CreateAnswer(db *gorm.DB)gin.HandlerFunc{
	return func(ctx *gin.Context) {

		var createAnswer abstract.Create_Answer
		var id_question,option_a,option_b,option_c,option_d int
		var listOption []int

		id_question,_ = strconv.Atoi(ctx.PostForm("id_question"))
		option_a,_ = strconv.Atoi(ctx.PostForm("option_a"))
		option_b,_ = strconv.Atoi(ctx.PostForm("option_b"))
		option_c,_ = strconv.Atoi(ctx.PostForm("option_c"))
		option_d,_ = strconv.Atoi(ctx.PostForm("option_d"))

		listOption = append(listOption,option_a)
		listOption = append(listOption,option_b)
		listOption = append(listOption,option_c)
		listOption = append(listOption,option_d)

		createAnswer.Status = 0
		if(id_question == 0){
			createAnswer.Message = "id question can not be empty or not exist"
		}else if(isIDQuestionExist(db,id_question) == false){
			createAnswer.Message = "wrong id test!!"
		}else if(isAnswerAlreadyCreate(db,id_question) == true){
			createAnswer.Message = "answer is already assign to this question!!"
		}else if(isMinimalAnswerValue(listOption,1) == false){
			createAnswer.Message = "answer value amount for any question is 1"
		}else{
			dataMap := make(map[string]interface{})
			dataMap["id_question"] = id_question
			dataMap["option_a"] = option_a
			dataMap["option_b"] = option_b
			dataMap["option_c"] = option_c
			dataMap["option_d"] = option_d

			rtn_create_answer := model.CreateAnswer(db,dataMap)
			createAnswer.Status = rtn_create_answer.Status
			createAnswer.Message = rtn_create_answer.Message

		}

		ctx.JSON(200, createAnswer)

	}

}


func ModifyTest(db *gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context) {
		c.Set("databaseConn", db)
		c.Next()
	}

}


func ModifyQuestion(db *gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context) {
		c.Set("databaseConn", db)
		c.Next()
	}

}

func StartTest(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var startStudentTests abstract.Start_Student_Test
		var id_student,id_test int

		id_student,_ = strconv.Atoi(ctx.PostForm("id_student"))
		id_test,_ = strconv.Atoi(ctx.PostForm("id_test"))

		startStudentTests.Status = 0
		if(id_student == 0 || id_test == 0){
			startStudentTests.Message = "id_student and id_test can't be empty or zero"
		}else if(isExistStudent(db,id_student) == false){
			startStudentTests.Message = "student is not exist"
		}else if(isIDTestExist(db,id_test) == false){
			startStudentTests.Message = "test is not exist"
		}else if(isStudentStartSameTest(db,id_test,id_student) == true){
			startStudentTests.Message = "can't start the same test more than 1"
		}else{
			dataMap := make(map[string]interface{})
			dataMap["id_test"] = id_test
			dataMap["id_student"] = id_student

			rtn_start_tests := model.StartStudentTest(db,dataMap)
			startStudentTests.Status = rtn_start_tests.Status
			startStudentTests.Message = rtn_start_tests.Message

		}
		ctx.JSON(200, startStudentTests)
	}
}

func InputStudentAnswer(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var inputStudentAnswers abstract.Input_Student_Answers
		var id_student,id_test,id_question,option_a,option_b,option_c,option_d int
		var listOption []int

		id_student,_ = strconv.Atoi(ctx.PostForm("id_student"))
		id_test,_ = strconv.Atoi(ctx.PostForm("id_test"))
		id_question,_ = strconv.Atoi(ctx.PostForm("id_question"))
		option_a,_ = strconv.Atoi(ctx.PostForm("option_a"))
		option_b,_ = strconv.Atoi(ctx.PostForm("option_b"))
		option_c,_ = strconv.Atoi(ctx.PostForm("option_c"))
		option_d,_ = strconv.Atoi(ctx.PostForm("option_d"))

		listOption = append(listOption,option_a)
		listOption = append(listOption,option_b)
		listOption = append(listOption,option_c)
		listOption = append(listOption,option_d)

		dataMap := make(map[string]interface{})
		dataMap["id_test"] = id_test
		dataMap["id_student"] = id_student
		dataMap["id_question"] = id_question
		dataMap["option_a"] = option_a
		dataMap["option_b"] = option_b
		dataMap["option_c"] = option_c
		dataMap["option_d"] = option_d

		inputStudentAnswers.Status = 0
		if(id_student == 0 || id_test == 0 || id_question == 0){
			inputStudentAnswers.Message = "id_student and id_test can't be empty or zero"
		}else if(isExistStudent(db,id_student) == false){
			inputStudentAnswers.Message = "student is not exist"
		}else if(isStudentStartSameTest(db,id_test,id_student) == false){
			inputStudentAnswers.Message = "you have to start test first!"
		}else if(isIDTestExist(db,id_test) == false){
			inputStudentAnswers.Message = "test is not exist"
		}else if(isIDQuestionExist(db,id_question) == false){
			inputStudentAnswers.Message = "question is not exist"
		}else if(isMinimalAnswerValue(listOption,1) == false){
			inputStudentAnswers.Message = "answer value amount for any question is 1"
		}else if(isStudentInputAnswerInSameQuestion(db,id_test,id_student,id_question)){
			inputStudentAnswers.Message = "can't answers the same question"
		}else{

			rtn_input_answers := model.InputStudentAnswer(db, dataMap)
			inputStudentAnswers.Status = rtn_input_answers.Status
			inputStudentAnswers.Message = rtn_input_answers.Message
		}


		ctx.JSON(200,inputStudentAnswers)

	}
}

func CompletionTest(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var completionTests abstract.Completion_Tests
		var id_student,id_test int

		id_student,_ = strconv.Atoi(ctx.PostForm("id_student"))
		id_test,_ = strconv.Atoi(ctx.PostForm("id_test"))

		completionTests.Status = 0
		if(id_student == 0 || id_test == 0){
			completionTests.Message = "id_student and id_test can't be empty or zero"
		}else if(isExistStudent(db,id_student) == false){
			completionTests.Message = "student is not exist"
		}else if(isIDTestExist(db,id_test) == false){
			completionTests.Message = "test is not exist"
		}else if(isStudentStartSameTest(db,id_test,id_student) == false){
			completionTests.Message = "can't complete the different test"
		}else{
			getInsight := getStudentScoreByTestID(db,id_test,id_student)
			studentScore := getInsight["student_score"]
			totalQuestion := getInsight["total_question"]
			correctAnswer := getInsight["correct_answer"]
			wrongAnswer := getInsight["wrong_answer"]
			dataMap := make(map[string]interface{})
			dataMap["id_test"] = id_test
			dataMap["id_student"] = id_student
			dataMap["student_score"] = studentScore
			dataMap["total_question"] = totalQuestion
			dataMap["correct_answer"] = correctAnswer
			dataMap["wrong_answer"] = wrongAnswer

			rtn_completion_tests := model.CompletionTest(db,dataMap)
			completionTests.Status = rtn_completion_tests.Status
			completionTests.Message = rtn_completion_tests.Message

		}

		ctx.JSON(200,completionTests)
	}
}

func StudentgetInsight(db *gorm.DB) gin.HandlerFunc {
	return func (ctx *gin.Context) {
		var getStudentInsight abstract.Get_Studdent_Insight
		var student_score abstract.Student_Score
		var id_student,id_test int

		filter := make(map[string]int)
		filter["id"],_ = strconv.Atoi(ctx.PostForm("id_student"))
		filter["id_student"],_ = strconv.Atoi(ctx.PostForm("id_student"))
		filter["id_test"],_ = strconv.Atoi(ctx.PostForm("id_test"))


		rtn_student := model.Students(db,filter)
		id_student,_ = strconv.Atoi(ctx.PostForm("id_student"))
		id_test,_ = strconv.Atoi(ctx.PostForm("id_test"))

		getInsight := getStudentScoreByTestID(db,id_test,id_student)
		student_score.Total_score = getInsight["student_score"]
		student_score.Total_Question = getInsight["total_question"]
		student_score.Correct_answer = getInsight["correct_answer"]
		student_score.Wrong_answer = getInsight["wrong_answer"]
		student_score.Percentage_correct_answer = (float64(getInsight["correct_answer"])/float64(getInsight["total_question"]))*100


		getStudentInsight.Status = 1
		getStudentInsight.Message = "Success"
		getStudentInsight.Students = rtn_student.Data
		getStudentInsight.Student_score = student_score


		ctx.JSON(200,getStudentInsight)

	}
}

func isIDTestExist(db *gorm.DB,id_test int) bool{
	filter := make(map[string]int)
	filter["id"] = id_test
	rtn_test := model.Tests(db,filter)

	if(rtn_test.Count > 0){
		return true
	}else{
		return false
	}
}

func isIDQuestionExist(db *gorm.DB,id_question int) bool{
	filter := make(map[string]int)
	filter["id"] = id_question
	rtn_question := model.Questions(db,filter)

	if(rtn_question.Count > 0){
		return true
	}else{
		return false
	}
}

func isAnswerAlreadyCreate(db *gorm.DB,id_question int) bool{
	filter := make(map[string]int)
	filter["id_question"] = id_question
	rtn_answer := model.Answers(db,filter)

	if(rtn_answer.Count > 0){
		return true
	}else{
		return false
	}
}

func isMinimalAnswerValue(listOption []int,min int) bool{

	count:=0
	for i:=0;i<len(listOption);i++{
		if(listOption[i] == 1){
			count++
		}
	}

	if(count == min){
		return true
	}else{
		return false
	}
}

func isExistStudent(db *gorm.DB,id_student int) bool{
	filter := make(map[string]int)
	filter["id"] = id_student

	rtn_student := model.Students(db,filter)

	if(rtn_student.Count > 0){
		return true
	}else{
		return false
	}

}


func isStudentStartSameTest(db *gorm.DB,id_test int, id_student int) bool{

	filter := make(map[string]int)
	filter["id_test"] = id_test
	filter["id_student"] = id_student


	rtn_student_test := model.StudentTests(db,filter)

	if(rtn_student_test.Count > 0){
		return true
	}else{
		return false
	}

}

func isStudentInputAnswerInSameQuestion(db *gorm.DB,id_test int, id_student int,id_question int) bool{

	filter := make(map[string]int)
	filter["id_test"] = id_test
	filter["id_student"] = id_student
	filter["id_question"] = id_question


	rtn_student_answers := model.StudentAnswers(db,filter)

	if(rtn_student_answers.Count > 0){
		return true
	}else{
		return false
	}

}

func getStudentAnswers(db *gorm.DB, id_test int,id_student int) []abstract.Student_Answers{

	filter := make(map[string]int)
	filter["id_test"] = id_test
	filter["id_student"] = id_student

	rtn_student_answers := model.StudentAnswers(db,filter)

	println(len(rtn_student_answers.Data))

	return rtn_student_answers.Data
}

func getAnswers(db *gorm.DB,id_question int) []abstract.Answers{
	filter := make(map[string]int)
	filter["id_question"] = id_question

	rtn_answers := model.Answers(db,filter)

	return rtn_answers.Data
}

func getTestByID(db *gorm.DB,id_test int) []abstract.Tests{
	filter := make(map[string]int)
	filter["id"] = id_test
	rtn_tests := model.Tests(db,filter)

	return rtn_tests.Data
}

func getStudentScoreByTestID(db *gorm.DB,id_test int,id_student int) map[string]int{
	var dataStudentAnswer []abstract.Student_Answers
	var dataAnswer []abstract.Answers
	var dataTest []abstract.Tests
	getInsight := make(map[string]int)
	count_positive := 0
	count_negative := 0

	dataTest = getTestByID(db,id_test)
	positive := dataTest[0].Positif_weight
	negative := dataTest[0].Negatif_weight

	dataStudentAnswer = getStudentAnswers(db,id_test,id_student)
	for i:=0;i<len(dataStudentAnswer) ;i++  {
		dataAnswer = getAnswers(db,dataStudentAnswer[i].Id_question)

		if(dataStudentAnswer[i].Option_a == dataAnswer[0].Option_a && dataStudentAnswer[i].Option_b == dataAnswer[0].Option_b && dataStudentAnswer[i].Option_c == dataAnswer[0].Option_c && dataStudentAnswer[i].Option_d == dataAnswer[0].Option_d){
			count_positive++
		}else{
			count_negative++
		}
	}

	getInsight["student_score"] = (positive*count_positive) + (negative*count_negative)
	getInsight["total_question"] = count_positive + count_negative
	getInsight["correct_answer"] = count_positive
	getInsight["wrong_answer"] = count_negative

	return getInsight

}