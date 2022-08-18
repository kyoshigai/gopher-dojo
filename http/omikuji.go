package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var tmpl = template.Must(template.New("message").Parse("<html><body>{{.Name}}さんの運勢は「<b>{{.Omikuji}}</b>」です</body></html>"))

type Person struct {
	Name    string
	Omikuji string
}

func DrawOmikuji() string {
	kinds := []string{"大吉", "中吉", "小吉", "末吉", "凶", "大凶"}
	rand.Seed(time.Now().UnixNano())
	return kinds[rand.Intn(len(kinds))]
}

func Handler(w http.ResponseWriter, r *http.Request) {
	p := &Person{Name: r.FormValue("p"), Omikuji: DrawOmikuji()}
	tmpl.Execute(w, p)
}

func main() {
	http.HandleFunc("/", Handler)
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatalln(err)
	}
}
