package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/AllChallenges", allChallenges)
	http.ListenAndServe(":8000", nil)
}

func allChallenges(w http.ResponseWriter, r *http.Request) {
	challenges := getChllanges()
	for _, challange := range challenges {
		chal, _ := json.Marshal(challange)
		fmt.Println(string(chal))
		fmt.Fprintln(w, string(chal))
	}
}
