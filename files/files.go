package files

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

type DirEntry struct {
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
}

func ListFiles(w http.ResponseWriter, r *http.Request) {

	// state
	files, err := os.ReadDir(".")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonableFiles := make([]DirEntry, 0, len(files))
	for _, entry := range files {
		jsonableFiles = append(jsonableFiles, DirEntry{entry.Name(), entry.IsDir()})
	}

	// representation
	filesJson, err := json.MarshalIndent(jsonableFiles, " ", "")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// transfer
	w.Write(filesJson)
}

type File struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func ShowFile(w http.ResponseWriter, r *http.Request) {

	cwd, err := os.Getwd()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)

	filename := cwd + "/" + vars["file"]

	search := r.URL.Query().Get("text")

	if search != "" {

		// state
		file, err := os.Open(filename)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		scanner := bufio.NewScanner(file)

		number := 1

		// representation
		lines := strings.Builder{}

		for scanner.Scan() {

			text := scanner.Text()

			if strings.Contains(text, search) {
				lines.WriteString(fmt.Sprintf("%d\t %s\n", number, text))
			}

			number++
		}

		w.Write([]byte(lines.String()))
		return

	}

	// state
	file, err := os.ReadFile(filename)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json, err := json.MarshalIndent(File{Name: vars["file"], Content: string(file)}, " ", "")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// transfer
	w.Write(json)
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func CreateFile(w http.ResponseWriter, r *http.Request) {

	// representation
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	createIntent := File{}

	err = json.Unmarshal(body, &createIntent)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if createIntent.Name == "" || createIntent.Content == "" {
		http.Error(w, "empty payload", http.StatusUnprocessableEntity)
		return
	}

	cwd, err := os.Getwd()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileName := cwd + "/" + createIntent.Name

	if fileExists(fileName) {
		http.Error(w, "file already exists", http.StatusUnprocessableEntity)
		return
	}

	// state
	err = os.WriteFile(fileName, []byte(createIntent.Content), fs.ModePerm)

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

	if updateIntent.Content == "" {
		http.Error(w, "empty payload", http.StatusUnprocessableEntity)
		return
	}

	cwd, err := os.Getwd()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)

	fileName := cwd + "/" + vars["file"]

	if !fileExists(fileName) {
		http.Error(w, "file doesn't exists", http.StatusInternalServerError)
		return
	}

	err = os.WriteFile(fileName, []byte(updateIntent.Content), fs.ModeAppend)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {

	cwd, err := os.Getwd()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)

	fileName := cwd + "/" + vars["file"]

	if !fileExists(fileName) {
		http.Error(w, "file doesn't exists", http.StatusInternalServerError)
		return
	}

	err = os.Remove(fileName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
