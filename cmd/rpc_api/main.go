package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/fotonmoton/golang_networking/pkg/files"
)

type Files int

func (f *Files) List(_, result *[]files.DirEntry) error {
	files, err := files.ListFiles(".")

	if err != nil {
		return err
	}

	*result = files

	return nil
}

func (f *Files) Show(name string, result *files.File) error {
	file, err := files.ShowFile(name)

	if err != nil {
		return err
	}

	*result = *file

	return nil
}

func (f *Files) Create(args files.File, result *files.File) error {
	err := files.CreateFile(args)

	if err != nil {
		return err
	}

	*result = args

	return nil
}

func main() {
	files := new(Files)
	server := rpc.NewServer()
	server.Register(files)

	listener, err := net.Listen("tcp", "localhost:8081")

	if err != nil {
		panic(err)
	}

	log.Println("Server started")
	for {
		conn, err := listener.Accept()
		log.Println(`Connection accepted`)

		if err != nil {
			log.Println(err)
			continue
		}

		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
