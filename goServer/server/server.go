package server

import (
	"encoding/json"
	"fmt"
	store "goServer/entity"
	"goServer/internal/service"
	"io"
	"net/http"
)

var (
	V           = &service.DataStore{
		Books: make(map[string]*store.Book),
	}
)

func StartServer(srv *http.Server) error{
	http.HandleFunc("/index", Index)
	http.HandleFunc("/book/get", GetBookHandler)
	http.HandleFunc("/book/delete", DelBookHandler)
	http.HandleFunc("/book", CreateBook)
	fmt.Println("http server start....")
	error := srv.ListenAndServe()
	return error
}


func Index(w http.ResponseWriter, r *http.Request){
	io.WriteString(w,"hello world")
}

func  CreateBook(w http.ResponseWriter, r *http.Request){
	dec := json.NewDecoder(r.Body)
	var book store.Book
	if err := dec.Decode(&book); err != nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}


	if err := V.Create(&book); err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func  GetBookHandler(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()

	id := query.Get("id")
	if id == "" {
		http.Error(w, "no id found in request", http.StatusBadRequest)
		return
	}

	book, err := V.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response(w, book)
}


func DelBookHandler(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()

	id := query.Get("id")
	if id == "" {
		http.Error(w, "no id found in request", http.StatusBadRequest)
		return
	}

	err := V.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func response(w http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}


