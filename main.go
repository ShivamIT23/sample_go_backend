package main

import (
	"fmt"
	"net/http"
	"log"
)

func main(){
	
	http.HandleFunc(("/"),func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w,"this is only for test");
	})

	s := &http.Server{
		Addr : ":8002",
		Handler : nil,
	}

	serving()

	log.Fatal(s.ListenAndServe())
}

func serving(){
	fmt.Println("Server is running on port 8002")
}
