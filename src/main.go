package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Relationship struct {
	Name         string
	Points       int
	Relationship string
	Inspiration  bool
	BoonOrBane   string
}

type Job struct {
	Employer string
	Job      string
	Coworker string
}

type ExtraCurricular struct {
	Id     int
	Name   string
	D4     bool
	Skills string
	Member string
}

type ReportCard struct {
	Id     int
	Year   int
	Skills []Skill
}

type Skill struct {
	Id      int
	Name    string
	D4s     int
	Rerolls int
}

type Character struct {
	Id               int
	Name             string
	ReportCards      []ReportCard
	Relationships    []Relationship
	ExtraCurriculars []ExtraCurricular
	Job              string
	Coworker         string
	Employer         string
}

func getCharacterById(id int) (Character, error) {
	db, err := sql.Open("postgres", "user=strixhaven password=changeme dbname=strixhaven sslmode=disable")
	if err != nil {
		return Character{}, err
	}
	defer db.Close()

	var character Character
	err = db.QueryRow("SELECT * FROM characters WHERE Id = $1", id).Scan(
		&character.Id,
		&character.Name,
		&character.Job,
		&character.Coworker,
		&character.Employer,
	)

	if err != nil {
		return Character{}, err
	}

	return character, nil
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/character/{id}", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		idParam := vars["id"]
		id, err := strconv.Atoi(idParam)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		character, err := getCharacterById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(character)
	}).Methods("GET")

	log.Print("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
