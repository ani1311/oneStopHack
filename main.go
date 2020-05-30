package main

import (
	"./models"
	"./website/hackerEarth"
	"fmt"
	"net/http"
	"text/template"
)

type Data struct {
	Challenges []models.Challenge
}

func main() {
	http.HandleFunc("/AllChallenges", allChallenges)
	http.ListenAndServe(":8000", nil)
}

func allChallenges(w http.ResponseWriter, r *http.Request) {
	challenges := hackerEarth.GetChallanges()
	data := Data{Challenges: challenges}
	t, err := template.ParseFiles("templates/allChallenges.html")
	if err != nil {
		fmt.Println(err)
	}
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}
