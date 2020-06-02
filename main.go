package main

import (
	"./models"
	"./website/CodeChef"
	"./website/hackerEarth"
	"./websiteData"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
	"time"
)

var challangeFile = "challenges.json"

type Data struct {
	Challenges []models.Challenge
	Websites   []websiteData.Website
}

func main() {
	http.HandleFunc("/challenges", allChallenges)
	http.HandleFunc(string("/challenges/"+websiteData.HackerEarth), hackerEarthChallenges)
	http.HandleFunc(string("/challenges/"+websiteData.CodeChef), codeChefChallenges)
	http.ListenAndServe(":8000", nil)
}

func checkAndUpdateChallenges() {
	info, _ := os.Stat(challangeFile)
	if info.ModTime().Add(time.Hour * 12).After(time.Now()) {
		return
	}
	challenges := hackerEarth.GetChallanges()
	challenges = append(challenges, CodeChef.GetChallanges()...)
	data := Data{Challenges: challenges, Websites: websiteData.Websites}
	dataJson, _ := json.Marshal(data)
	ioutil.WriteFile(challangeFile, dataJson, 7777)
}

func getAllChallenges() Data {
	checkAndUpdateChallenges()
	var data Data
	file, _ := os.Open(challangeFile)
	defer file.Close()

	fileBytes, _ := ioutil.ReadAll(file)
	json.Unmarshal(fileBytes, &data)

	return data
}

func getChallengesOf(ws websiteData.Website) Data {
	checkAndUpdateChallenges()
	var tempData Data
	file, _ := os.Open(challangeFile)
	defer file.Close()

	fileBytes, _ := ioutil.ReadAll(file)
	json.Unmarshal(fileBytes, &tempData)

	var data Data
	for _, chal := range tempData.Challenges {
		if chal.Website == ws {
			data.Challenges = append(data.Challenges, chal)
		}
	}
	data.Websites = tempData.Websites

	return data
}

func allChallenges(w http.ResponseWriter, r *http.Request) {
	data := getAllChallenges()
	t, err := template.ParseFiles("templates/allChallenges.html")
	if err != nil {
		fmt.Println(err)
	}
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

func codeChefChallenges(w http.ResponseWriter, r *http.Request) {
	data := getChallengesOf(websiteData.CodeChef)
	t, err := template.ParseFiles("templates/allChallenges.html")
	if err != nil {
		fmt.Println(err)
	}
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

func hackerEarthChallenges(w http.ResponseWriter, r *http.Request) {
	data := getChallengesOf(websiteData.HackerEarth)
	t, err := template.ParseFiles("templates/allChallenges.html")
	if err != nil {
		fmt.Println(err)
	}
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}
