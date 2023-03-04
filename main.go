package main

import (
	"fmt"
)

type StudentDatabaseOP interface {
	AddStudent(student_db_arr StudentDatabaseStruArr, student_stru StudentDatabaseStru) bool
	DecStudent(student_db_arr StudentDatabaseStruArr, student_stru StudentDatabaseStru) bool
}

type StudentDatabaseStru struct {
	name    string
	age     int
	address string
	phone   string
}

type StudentDatabaseStruMgr struct {
}

type StudentDatabaseStruArr []StudentDatabaseStru

func MakeStudentDatabaseStruEntry() StudentDatabaseStruArr {
	return make(StudentDatabaseStruArr, 256)
}

func (StudentDatabaseStruMgr) AddStudent(student_db_arr StudentDatabaseStruArr, student_stru StudentDatabaseStru) bool {
	fmt.Println("11111")
	return true
}

func (StudentDatabaseStruMgr) DecStudent(student_db_arr StudentDatabaseStruArr, student_stru StudentDatabaseStru) bool {
	fmt.Println("22222")
	return true
}

func main() {
	student := StudentDatabaseStru{"xxxx", 23, "xxxx", "yyyyy"}
	student_db := MakeStudentDatabaseStruEntry()
	students_mgr := new(StudentDatabaseStruMgr)
	students_mgr.AddStudent(student_db, student)
	students_mgr.DecStudent(student_db, student)
	fmt.Println("Hello,World!", len(student_db))
}
