package main

import(
	"fmt"
	"log"
	"net/http"
	"strconv"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/thejis/database/models"
)

func GetNotes (w http.ResponseWriter, r *http.Request){
	n := new(models.Note)
	notes, err := n.GetAll()
	if err != nil{
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	j, err := json.Marshal(notes)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func CreateNotes(w http.ResponseWriter, r *http.Request){
	var note models.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = note.Create()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateNotes (w http.ResponseWriter, r *http.Request){
	var note models.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = note.Update()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteNotes(w http.ResponseWriter, r *http.Request){
	idStr := r.URL.Query().Get("id")
	if idStr == ""{
		http.Error(w, "Query id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Query id must be a number", http.StatusBadRequest)
		return
	}

	var note models.Note

	err = note.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func main(){
	router := mux.NewRouter()

	router.HandleFunc("/notes", GetNotes).Methods("GET")
	router.HandleFunc("/notes", CreateNotes).Methods("POST")
	router.HandleFunc("/notes", UpdateNotes).Methods("UPDATE")
	router.HandleFunc("/notes", DeleteNotes).Methods("DELETE")

	fmt.Println("starting server on port 8000")
	log.Println(http.ListenAndServe(":8000", router))
}
