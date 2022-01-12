package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Quiz struct {
	QuizID   string `json:"quizid"`
	QuizName string `json:"quizname"`
	QuizVar1 string `json:"quizvar1"`
	QuizVar2 string `json:"quizvar2"`
	QuizVar3 string `json:"quizvar3"`
	QuizVar4 string `json:"quizvar4"`
}

type JsonResponse struct {
	Type    string `json:"type"`
	Data    []Quiz `json:"data"`
	Message string `json:"message"`
}

func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func AllQuiz(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Getting movies...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query(`SELECT id, "quizName", "quizVar1", "quizVar2", "quizVar3", "quizVar4"  FROM public."Quiz"`)

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var quizes []Quiz

	// Foreach movie
	for rows.Next() {
		// var id int
		var quizID string
		var quizName string
		var quizVar1 string
		var quizVar2 string
		var quizVar3 string
		var quizVar4 string

		err = rows.Scan(&quizID, &quizName, &quizVar1, &quizVar2, &quizVar3, &quizVar4)

		// check errors
		checkErr(err)

		quizes = append(quizes, Quiz{QuizID: quizID, QuizName: quizName, QuizVar1: quizVar1, QuizVar2: quizVar2, QuizVar3: quizVar3, QuizVar4: quizVar4})
	}

	var response = JsonResponse{Type: "success", Data: quizes}

	json.NewEncoder(w).Encode(response)
}

func NewUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "New user endpoind")
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Delete user endpoind")
}

func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Update user endpoind")
}
