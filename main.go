package main

import (
	"log"
	"net/http"
)

func main() {
	// Servir les fichiers statiques depuis web/static à /static/
	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Page d'accueil statique (index.html à la racine du projet)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Si l'utilisateur demande exactement "/", on rend index.html
		if r.URL.Path != "/" {
			// Laisser le serveur statique gérer les autres chemins si possible
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "index.html")
	})

	addr := ":8080"
	log.Printf("Démarrage du serveur sur %s...", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}