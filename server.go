package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":80", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	nome := os.Getenv("NOME")
	idade := os.Getenv("IDADE")

	w.Write([]byte("<h1>Fullcycle app go!!</h1>"))
	w.Write([]byte("<h1>Loadbalancer com configmap</h1>"))
	w.Write([]byte("<h1>Testado v2</h1>"))
	fmt.Fprintf(w, "Salve, eu sou o %s. e eu tenho %s anos de idade.", nome, idade)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("minhaFamilia/familia.txt")
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo", string(data))
	}
	fmt.Fprintf(w, "Minha familia: %s. ", string(data))
}

func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	fmt.Fprintf(w, "User: %s. Password %s.", user, password)
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(startedAt)
	if duration.Seconds() < 10 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("duration: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}
}
