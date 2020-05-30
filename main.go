package main

import (
	"./models"
	"./website/CodeChef"
	"./website/hackerEarth"
	"fmt"
	"net/http"
	"text/template"
)

type Data struct {
	Challenges []models.Challenge
}

func main() {
	// CodeChef.GetChallanges()
	http.HandleFunc("/AllChallenges", allChallenges)
	http.ListenAndServe(":8000", nil)
}

func allChallenges(w http.ResponseWriter, r *http.Request) {
	challenges := hackerEarth.GetChallanges()
	challenges = append(challenges, CodeChef.GetChallanges()...)
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
