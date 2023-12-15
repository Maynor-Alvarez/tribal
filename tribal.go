package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

type Chuck struct {
	IconUrl string `json:"iconUrl"`
	Id      string `json:"id"`
	Url     string `json:"url"`
	Value   string `json:"value"`
}

func main() {

	http.HandleFunc("/api/chuck", getChuck)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func getChuck(w http.ResponseWriter, r *http.Request) {
	num := 25
	chucks := make([]Chuck, 0)
	var unique = make(map[string]bool)
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < num; i++ {
		wg.Add(1)
		go fetchChuck(w, &wg, &chucks, &unique, &mutex)
	}

	wg.Wait()

	jsonResp, err := json.Marshal(chucks)
	if err != nil {
		http.Error(w, "Error al obtener datos JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	_, err = w.Write(jsonResp)
	if err != nil {
		http.Error(w, "Error al responder JSON", http.StatusInternalServerError)
		return
	}
}

func fetchChuck(w http.ResponseWriter, wg *sync.WaitGroup, chucks *[]Chuck, unique *map[string]bool, mutex *sync.Mutex) {
	defer wg.Done()

	response, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		http.Error(w, "Error al obtener datos", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	var chuck Chuck
	if err := json.NewDecoder(response.Body).Decode(&chuck); err != nil {
		http.Error(w, "Error al generar JSON", http.StatusInternalServerError)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	if !(*unique)[chuck.Id] {
		(*unique)[chuck.Id] = true
		*chucks = append(*chucks, chuck)
	}
}
