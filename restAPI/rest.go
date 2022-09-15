package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Student struct {
	Id int
	Name string
	Age int
	Score int
}

var students map[int]Student
var lastId int

func MkaeWebHandler() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/students", GetStudentListHandler).Methods("GET")
	students = make(map[int]Student)
	students[1] = Student{1, "aaa", 16, 87}
	students[2] = Student{2, "bbb", 18, 98}
	lastId = 2

	return mux
}

