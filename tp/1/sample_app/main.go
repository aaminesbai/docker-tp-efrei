package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
	"github.com/fatih/color"
)

var cats = []string{
	"meow",
	`|\---/|
| o_o |
 \_^_/`,
 ` /\_/\
( o.o )
 > ^ <`,
}

type CatResponse struct {
	Art string `json:"art"`
}

func getRandomCatHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	cat := cats[rand.Intn(len(cats))]

	resp := CatResponse{Art: cat}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/cat", getRandomCatHandler)
	log.Println("I'm listening on TCP port 8080")
	color.Green("Server active !")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

