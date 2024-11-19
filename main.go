package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Recipe struct {
	Title string
	Text  string
}

var recipes = map[string]Recipe{
	"espresso": {
		Title: "Espresso",
		Text:  "1. Preheat your espresso machine.\n2. Measure 18–20 grams of finely-ground coffee.\n3. Brew and enjoy!",
	},
	"americano": {
		Title: "Americano",
		Text:  "1. Brew a single or double shot of espresso.\n2. Heat water.\n3. Mix and enjoy!",
	},
	"latte": {
		Title: "Latte",
		Text:  "1. Brew espresso.\n2. Steam milk.\n3. Mix and garnish!",
	},
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}

	query = sanitizeInput(query) // Normalize the input (case-insensitive)

	if recipe, found := recipes[query]; found {
		tmpl.Execute(w, map[string]interface{}{
			"Recipe": recipe,
		})
	} else {
		tmpl.Execute(w, map[string]interface{}{
			"Error": "Recipe not found. Please try searching for 'Espresso', 'Americano', or 'Latte'.",
		})
	}
}

func sanitizeInput(input string) string {
	// Convert to lowercase to make search case-insensitive
	return strings.ToLower(strings.TrimSpace(input))
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/search", searchHandler)

	log.Println("Server running on http://localhost:4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
