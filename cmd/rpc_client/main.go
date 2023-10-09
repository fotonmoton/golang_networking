package main

import (
	"fmt"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/fotonmoton/golang_networking/pkg/files"
	"github.com/fotonmoton/golang_networking/pkg/log"
)

func GetFiles(c *rpc.Client) []files.DirEntry {
	var result []files.DirEntry

	// Make it async
	err := c.Call("Files.List", nil, &result)
	if err != nil {
		log.Error(err.Error())
	}

	return result
}

func Show(c *rpc.Client, name string) files.File {
	var result files.File

	err := c.Call("Files.Show", name, &result)
	if err != nil {
		log.Error(err.Error())
	}

	return result
}

func Create(c *rpc.Client, name string, content string) files.File {
	var result files.File

	err := c.Call("Files.Create", struct{ Name, Content string }{Name: name, Content: content}, &result)
	if err != nil {
		log.Error(err.Error())
	}

	return result
}

func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:8081")

	if err != nil {
		panic(err)
	}

	files := GetFiles(client)

	log.Info(fmt.Sprintf("%+v\n", files))

	file := Show(client, "cmd/rest_api/main.go")

	log.Info(fmt.Sprintf("%+v\n", file))

	copy := Create(client, "copy.go", file.Content)

	log.Info(fmt.Sprintf("%+v\n", copy))
}
