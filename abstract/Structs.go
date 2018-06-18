package abstract

type Tests struct {
	Id int `json:"id" form:"id"`
	T_description string `json:"t_description" form:"t_description"`
	Positif_weight int `json:"positif_weight" form:"positif_weight"`
	Negatif_weight int `json:"negatif_weight" form:"negatif_weight"`
	Created_date string `json:"created_date" form:"created_date"`
	Updated_date string `json:"updated_date" form:"updated_date"`
	Questions []Questions `json:"questions"`

}

type Get_Tests struct {
	Status int `json:"status" form:"status"`
	Message string `json:"message" form:"message"`
	Data []Tests
	Count int `json:"count" form:"count"`

}

type Create_Test struct {
	Status int `json:"status" form:"status"`
	Message string `json:"message" form:"message"`
}

type Questions struct{
	Id int `json:"id" form:"id"`
	Id_test int `json:"id_test" form:"id_test"`
	Q_number int `json:"q_number" form:"q_number"`
	Q_description string `json:"q_description" form:"q_description"`
	Option_a string `json:"option_a" form:"option_a"`
	Option_b string `json:"option_b" form:"option_b"`
	Option_c string `json:"option_c" form:"option_c"`
	Option_d string `json:"option_d" form:"option_d"`
	Created_date string `json:"created_date" form:"created_date"`
	Updated_date string `json:"updated_date" form:"updated_date"`
}

type Get_Questions struct {
	Status int `json:"status" form:"status"`
	Message string `json:"message" form:"message"`
	Data []Questions
	Count int `json:"count" form:"count"`
}

type Create_Question struct{
	Status int `json:"status" form:"status"`
	Message string `json:"message" form:"message"`
}

type Answers struct{
	Id int `json:"id"`
	Id_question   int `json:"id_question"`
	Option_a  int `json:"option_a"`
	Option_b  int `json:"option_b"`
	Option_c  int `json:"option_c"`	
	Option_d  int `json:"option_d"`

}

type Get_Answers struct {
	Status int `json:"status" form:"status"`
	Message string `json:"message" form:"message"`
	Data []Answers
	Count int `json:"count" form:"count"`
}

type Create_Answer struct{
	Status int `json:"status" form:"status"`
	Message string `json:"message" form:"message"`
}

type Students struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Get_Students struct {
	Status int `json:"status" form:"status"`
	Message string `json:"message" form:"message"`
	Data []Students
	Count int `json:"count" form:"count"`
}

type Student_Tests struct {
	Id  int `json:"id"`
	Id_student int `json:"id_student"`
	Id_test   int `json:"id_test"`
	Start_time string `json:"start_time"`
	End_time  string `json:"end_time"`
	Total_question int `json:"total_question"`
	Correct_answer int `json:"correct_answer"`
	Wrong_answer int `json:"wrong_answer"`
	Student_score int `json:"student_score"`
	Status int  `json:"status"`
}

type Get_Student_Tests struct {
	Status int `json:"status" form:"status"`
	Message string `json:"message" form:"message"`
	Data []Student_Tests
	Count int `json:"count" form:"count"`
}


type Start_Student_Test struct {
	Status int `json:"status" form:"status"`
	Message string `json:"message" form:"message"`
}

type Student_Answers struct {
	Id_student int `json:"id_student"`
	Id_test   int `json:"id_test"`
	Id_question int `json:"id_question"`
	Option_a  int `json:"option_a"`
	Option_b  int `json:"option_b"`
	Option_c  int `json:"option_c"`
	Option_d  int `json:"option_d"`
	Created_date string `json:"created_date"`
	Updated_date string `json:"updated_date"`

}

type Get_Student_Answers struct{
	Status int `json:"status" form:"status"`
	Message string `json:"message" form:"message"`
	Data []Student_Answers
	Count int `json:"count" form:"count"`

}

type Input_Student_Answers struct {
	Status int `json:"status" form:"status"`
	Message string `json:"message" form:"message"`
}

type Completion_Tests struct{
	Status int `json:"status" form:"status"`
	Message string `json:"message" form:"message"`
}

type Student_Score struct{
	Percentage_correct_answer float64 `json:"percentage_correct_answer"`
	Correct_answer int `json:"correct_answer"`
	Wrong_answer int `json:"wrong_answer"`
	Total_score int `json:"total_score"`
	Total_Question int `json:"total_question"`

}

type Get_Studdent_Insight struct {
	Status int `json:"status" form:"status"`
	Message string `json:"message" form:"message"`
	Students []Students
	Student_score Student_Score `json:"student_score"`
}
