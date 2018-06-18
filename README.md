# go-ruangguru

Path :
{GOPATH}/src/bitbucket.org/go-guru

How to run :    
```
    go run main.go

```


DB design :

  1. **tests**  
    - id auto increment INT    
    - t_description TEXT    
    - positif_weight INT    
    - negatif_weight INT    
    - created_date DATETIME     
    - updated_date DATETIME 
  
  2. **questions**  
   	-  id INT auto increment    
   	-  id_test  INT 
	-  q_number INT 
    -  q_description TEXT   
    -  option_a  TEXT   
    -  option_b  TEXT   
    -  option_c  TEXT   
    -  option_d  TEXT   
    -  created_date DATETIME    
    -  updated_date DATETIME    

  3.  **answers**   
 	- id INT auto increment 
    - id_question   INT   
    -  option_a  INT    
    -  option_b  INT    
    -  option_c  INT    
    -  option_d  INT    

  4.  **students**  
	- id_student INT auto increment
	- name TEXT          
 	- email TEXT          
	- password TEXT         

  5. **student_answers**	 
    - id INT auto_increment     
	- id_student INT    
	- id_test   INT     
	- id_question INT     
	-  option_a  INT    
    -  option_b  INT    
    -  option_c  INT    
    -  option_d  INT    
	-  created_date DATETIME    
	-  updated_date DATETIME    
	
  6. **student_tests**   
    - id  INT auto increment        
    - id_student INT    
    - id_test   INT 
    - start_time DATETIME   
    - end_time  DATETIME    
    - student_score INT     
    - status INT (1:active,0:not active,2:completed)    
    
    
    
   # FLOW Untuk Student Test - Answer - Question 
   
   1. Student Start Test
      http://localhost:8888/post/starttest 
      method : POST 
      param : 
      - id_test          
      - id_student                          
           
   2. Student Input Answer      
      http://localhost:8888/post/inputanswer        
      method : POST             
      param :           
      - id_test         
      - id_student      
      - id_question     
      - option_a               
      - option_b              
      - option_c                 
      - option_d            
      
   3. Student Completion Test
      http://localhost:8888/post/inputanswer        
      method : POST             
      param :           
      - id_test                    
      - id_student                        
              
    

  



      


