package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/bid", func(w http.ResponseWriter, r *http.Request) {
		amount := r.FormValue("amount")
		response := "<p>Ставка " + amount + " ₽ принята!</p>"
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(response))
	})

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
