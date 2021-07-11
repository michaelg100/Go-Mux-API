package main

import (
    "fmt"
    "log"
    "os"
    "net/http"
    "github.com/jdkato/prose/v2"
    "github.com/gorilla/mux"
)


func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]

    fmt.Fprintf(w, "{'Key': '" + key +"'}" )


    
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)

    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/article/{id}", returnSingleArticle)

    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {

    input_arg := os.Args[1]

    // Create a new document with the default configuration:
    doc, err := prose.NewDocument(input_arg)
    if err != nil {
        log.Fatal(err)
    }

    // Iterate over the doc's tokens:
    for _, tok := range doc.Tokens() {
        fmt.Println(tok.Text, tok.Tag, tok.Label)
    }

    // Iterate over the doc's named-entities:
    for _, ent := range doc.Entities() {
        fmt.Println(ent.Text, ent.Label)
    }

    // Iterate over the doc's sentences:
    for _, sent := range doc.Sentences() {
        fmt.Println(sent.Text)
    }
    
    handleRequests()
    
}
