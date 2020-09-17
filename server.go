package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var (
	templates    *template.Template
	CurrentState State
)

type State struct {
	History history
	Current []Choice
}
type history string

type Choice struct {
	Text   string
	result func(*history)
}

func init() {
	// init templates
	templates = template.Must(template.ParseFiles("index.gohtml"))
	rand.Seed(time.Now().UnixNano())
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/reset", resetHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	fmt.Println("server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// choice receiver
	if err := CurrentState.getChoice(r); err != nil {
		if len(CurrentState.History) != 0 {
			CurrentState = State{}
		}
	}
	// choice sender
	CurrentState.setChoice()
	if err := templates.ExecuteTemplate(w, "index.gohtml", CurrentState); err != nil {
		panic(err)
	}
}
func resetHandler(w http.ResponseWriter, r *http.Request) {

}
func (s *State) getChoice(r *http.Request) error {
	n, err := strconv.Atoi(r.FormValue("choice"))
	if err != nil {
		return err
	}
	s.Current[n].result(&s.History)
	return nil
}

func (s *State) setChoice() {
	n := rand.Intn(3)
	currentChoices := make([]Choice, 0, n)
	for i := 0; i <= n; i++ {
		currentChoices = append(currentChoices, makeChoice())
	}
	s.Current = currentChoices
}
func (s *State) reset() {
	s.History = ""
}
func makeChoice() Choice {
	possibilities := []string{"a", "b", "ch", "ts", "i", "o", "u", "t", "p", "qu", "e", "e", "a", "o", "i", "u", "'", "m", "n"}
	n := rand.Intn(len(possibilities))
	letter := possibilities[n]
	c := Choice{
		Text: fmt.Sprint(letter),
		result: func(h *history) {
			*h += history(possibilities[n])
		},
	}
	return c
}
