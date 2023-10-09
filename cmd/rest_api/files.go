package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/fotonmoton/golang_networking/pkg/files"
	"github.com/gorilla/mux"
)

func ListFiles(w http.ResponseWriter, r *http.Request) {

	files, err := files.ListFiles(".")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// representation
	filesJson, err := json.MarshalIndent(files, " ", "")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// transfer
	w.Write(filesJson)
}

func ShowFile(w http.ResponseWriter, r *http.Request) {

	name := mux.Vars(r)["file"]

	search := r.URL.Query().Get("text")

	if search != "" {

		found, err := files.Search(name, search)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json, err := json.MarshalIndent(found, " ", "")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(json)
		return
	}

	file, err := files.ShowFile(name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json, err := json.MarshalIndent(file, " ", "")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// transfer
	w.Write(json)
}

func CreateFile(w http.ResponseWriter, r *http.Request) {

	// representation
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	createIntent := files.File{}

	err = json.Unmarshal(body, &createIntent)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = files.CreateFile(createIntent)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// transfer
}

type UpdateFilePayload struct {
	Content string `json:"content"`
}

func UpdateFile(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updateIntent := UpdateFilePayload{}

	err = json.Unmarshal(body, &updateIntent)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = files.UpdateFile(mux.Vars(r)["file"], []byte(updateIntent.Content))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {

	err := files.DeleteFile(mux.Vars(r)["file"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
