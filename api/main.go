package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sorted/sort"

	"github.com/gorilla/mux"
)

//Request array
type Request struct {
	Numbers []int
}

//Response
type Response struct {
	SortedArray [][]int
}

func getSelectionSort(w http.ResponseWriter, r *http.Request) {
	var req Request
	json.NewDecoder(r.Body).Decode(&req)
	tmp := sort.SelectionSort(req.Numbers)
	res := Response{
		SortedArray: tmp,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func getBubbleSort(w http.ResponseWriter, r *http.Request) {
	var req Request
	json.NewDecoder(r.Body).Decode(&req)
	tmp := sort.BubbleSort(req.Numbers)
	res := Response{
		SortedArray: tmp,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/bubbleSort", getBubbleSort)
	router.HandleFunc("/selectionSort", getSelectionSort)
	log.Fatal(http.ListenAndServe(":8000", router))
}
