package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

type Profile struct {
	Name         string
	Title        string
	Skills       []Skill
	DBSkills     []Skill
	SocialLinks  []SocialLink
	ContactEmail string
}

type Skill struct {
	Name  string
	Level string
	Icon  string
}

type SocialLink struct {
	Platform string
	URL      string
	Icon     string
}

func main() {
	profile := Profile{
		Name:         "ImCluzzy",
		Title:        "Junior Разработчик (Go/Python/C#)",
		ContactEmail: "ImCluzzy@gmail.com",
		Skills: []Skill{
			{"Go", "Junior", "fab fa-golang"},
			{"Python", "Junior", "fab fa-python"},
			{"C#", "Junior", "fas fa-code"},
			{"HTML/CSS", "Intermediate", "fab fa-html5"},
			{"Git", "Intermediate", "fab fa-git-alt"},
		},
		DBSkills: []Skill{
			{"MySQL", "Intermediate", "fas fa-database"},
			{"PostgreSQL", "Intermediate", "fas fa-database"},
			{"SQLite", "Intermediate", "fas fa-database"},
			{"MariaDB", "Intermediate", "fas fa-database"},
			{"H2 Database", "Intermediate", "fas fa-database"},
		},
		SocialLinks: []SocialLink{
			{"Telegram", "https://t.me/imcluzzy", "fab fa-telegram"},
			{"GitHub", "https://github.com/imcluzzy", "fab fa-github"},
			{"VK", "https://vk.com/imcluzzy", "fab fa-vk"},
			{"Instagram", "https://instagram.com/dyatlikarsenii/", "fab fa-instagram"},
		},
	}

	tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "index.html")))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.Execute(w, profile); err != nil {
			log.Println("Ошибка рендеринга шаблона:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
