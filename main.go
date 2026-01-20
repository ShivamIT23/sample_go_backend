package main

import (
	"fmt"
	"os"
	"net/http"
	"log"
)

func main(){
	
	http.HandleFunc(("/"),func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w,"this is only for test");
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8002" // Fallback for local testing
	}
	s := &http.Server{
		Addr : ":" + port,
		Handler : nil,
	}

	serving(port)

	log.Fatal(s.ListenAndServe())
}

func serving(port string){
	fmt.Println(`Server is running on port `+ port)
}
