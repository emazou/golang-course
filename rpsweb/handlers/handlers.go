package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"rpsweb/rps"
	"strconv"
)

const (
	templateDir  = "templates/"
	baseTemplate = templateDir + "base.html"
)

type Player struct {
	Name string
}

var player Player

func Index(w http.ResponseWriter, r *http.Request) {
	restartGame()
	renderTemplate(w, baseTemplate, "index.html", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	restartGame()
	renderTemplate(w, baseTemplate, "new-game.html", nil)
}

func Game(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
			return
		}
		player.Name = r.Form.Get("name")
	}
	if player.Name == "" {
		http.Redirect(w, r, "/new", http.StatusSeeOther)
		return
	}
	renderTemplate(w, baseTemplate, "game.html", player)
}

func PlayGame(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("choice"))
	result := rps.PlayRound(playerChoice)
	out, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func About(w http.ResponseWriter, r *http.Request) {
	restartGame()
	renderTemplate(w, baseTemplate, "about.html", nil)
}

func renderTemplate(w http.ResponseWriter, base string, page string, data any) {
	tpl := template.Must(template.ParseFiles(base, templateDir+page))
	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func restartGame() {
	player.Name = ""
	rps.ComputerScore = 0
	rps.PlayerScore = 0
}
