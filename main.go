package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/google/go-github/v63/github"
)

func main() {
	mux := http.NewServeMux()

	type UserMap struct {
		GithubUsers []github.User
	}

	home := func(w http.ResponseWriter, r *http.Request) {

		client := github.NewClient(nil)

		profile, _, err := client.Users.Get(context.Background(), "manuellara")
		if err != nil {
			fmt.Println(err.Error())
		}

		data := UserMap{
			GithubUsers: []github.User{
				*profile,
			},
		}

		templ := template.Must(template.ParseFiles("templates/base.html", "templates/content.html"))

		templ.Execute(w, data)
	}

	mux.HandleFunc("GET /", home)

	log.Fatal(http.ListenAndServe(":8000", mux))
}
