package handlers

import (
	"Portfolio/src/blockchain"
	"Portfolio/src/data"
	"html/template"
	"net/http"
	"node"
)

type Data struct {
	Bc    blockchain.Blockchain
	Miner node.Node
}

var Bc = blockchain.InitBlockchain()
var Miner node.Node
var PeerList data.PeerList

func HandleHome(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("templates/index.html"))
	if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandleProjects(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("templates/projects.html"))
	if err := templates.ExecuteTemplate(w, "projects.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandleResume(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("templates/resume.html"))
	if err := templates.ExecuteTemplate(w, "resume.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandleSports(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("templates/sports.html"))
	if err := templates.ExecuteTemplate(w, "sports.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandleMusic(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("templates/music.html"))
	if err := templates.ExecuteTemplate(w, "music.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandleContact(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("templates/contact.html"))
	if err := templates.ExecuteTemplate(w, "contact.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandleBlockchain(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("templates/blockchain.html"))
	if err := templates.ExecuteTemplate(w, "blockchain.html", Data{Bc: *Bc, Miner: Miner}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandleStart(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("templates/blockchain.html"))
	go func() {
		Miner.StartTryingNonces(*Bc)
	}()
	if err := templates.ExecuteTemplate(w, "blockchain.html", Data{Bc: *Bc, Miner: Miner}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
