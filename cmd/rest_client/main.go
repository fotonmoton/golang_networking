package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/fotonmoton/golang_networking/pkg/files"
)

func List() []files.DirEntry {
	res, err := http.Get("http://localhost:8080/files")
	if err != nil {
		log.Fatalf("error making http request: %s\n", err)
	}

	fmt.Printf("client: got response!\n")

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("error reading body: %s\n", err)
	}

	fmt.Printf("client: status code: %d\n", res.StatusCode)
	fmt.Printf("client: body: %s\n", body)

	var files []files.DirEntry

	err = json.Unmarshal(body, &files)

	if err != nil {
		log.Fatalf("error decoding: %s\n", err)
	}

	return files
}

func main() {

	files := List()

	fmt.Printf("client: decoded body: %+v\n", files)
}
