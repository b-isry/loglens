package main 

import(
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "LogLense Backend is running...")
}

func main(){
	http.HandleFunc("/", homeHandler)

	port := ":8080"
	log.Printf("Starting servr on port %s", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}

}