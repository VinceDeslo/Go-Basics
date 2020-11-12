package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Data format
type Document struct {
	Id    string `json:"Id"`
	Title string `json:"Title"`
	Desc  string `json:"Desc"`
	Info  string `json:"Info"`
}

// Global for storing of mock data
var Documents []Document

// Main program entry point
func main() {

	// Mock data
	Documents = []Document{
		Document{Id: "1", Title: "Doc1", Desc: "Doc1_desc", Info: "Some info"},
		Document{Id: "2", Title: "Doc2", Desc: "Doc2_desc", Info: "Some info"},
	}
	fmt.Println("Rest API Server starting.")
	handleRequests()
}

// Root page endpoint
func rootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home endpoint")
}

// Request handler
func handleRequests() {

	// Instance of the mux router
	router := mux.NewRouter().StrictSlash(true)

	// Routes
	router.HandleFunc("/", rootPage)
	router.HandleFunc("/documents", returnAllDocuments)
	router.HandleFunc("/document/{id}", deleteDocument).Methods("DELETE")
	router.HandleFunc("/document/{id}", updateDocument).Methods("PUT")
	router.HandleFunc("/document/{id}", returnDocument)
	router.HandleFunc("/document", createDocument).Methods("POST")

	// Run server
	log.Fatal(http.ListenAndServe(":10000", router))
}

// READ: all data endpoint
func returnAllDocuments(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All documents endpoint")
	json.NewEncoder(w).Encode(Documents)
}

// READ: single data entry endpoint
func returnDocument(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Single document endpoint")

	// Obtain key from request route variables
	vars := mux.Vars(r)
	id := vars["id"]

	// Fetch correct document
	for _, doc := range Documents {
		if doc.Id == id {
			json.NewEncoder(w).Encode(doc)
		}
	}
}

// CREATE: data creation endpoint
func createDocument(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Document creation endpoint")

	// Fetch request body
	body, _ := ioutil.ReadAll(r.Body)

	// Encode body into JSON article and store it
	var doc Document
	json.Unmarshal(body, &doc)
	Documents = append(Documents, doc)
	json.NewEncoder(w).Encode(doc)
}

// UPDATE: data updating endpoint
func updateDocument(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update document endpoint")

	// Obtain key from request route variables
	vars := mux.Vars(r)
	id := vars["id"]

	// Fetch request body
	body, _ := ioutil.ReadAll(r.Body)

	// Fetch correct document
	for index, doc := range Documents {
		if doc.Id == id {

			// Encode body into JSON article and store it
			json.Unmarshal(body, &doc)
			Documents[index] = doc
			json.NewEncoder(w).Encode(doc)
		}
	}
}

// DELETE: data deletion endpoint
func deleteDocument(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete document endpoint")

	// Obtain key from request route variables
	vars := mux.Vars(r)
	id := vars["id"]

	// Fetch correct document
	for index, doc := range Documents {
		if doc.Id == id {
			Documents = append(Documents[:index], Documents[index+1:]...)
		}
	}
}
