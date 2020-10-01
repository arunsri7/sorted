package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Request struct {
	Numbers []int
}

type Response struct {
	SortedArray [][]int
}

func bubbleSort(a []int) [][]int {
	//2-D array to store the sorted array after each iteration
	var values [][]int
	for j := 0; j < len(a); j++ {
		var b []int
		count := 0
		for i := range a {
			if i == len(a)-1 {
				continue
			}
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				count++
			}
		}
		//To avoid looping over the sorted array
		if count == 0 {
			return values
		}
		//Appending the value of 'a' to 'values'
		b = make([]int, len(a))
		copy(b, a)
		values = append(values, b)
	}
	return values
}

func handler(w http.ResponseWriter, r *http.Request) {
	var req Request
	json.NewDecoder(r.Body).Decode(&req)
	tmp := bubbleSort(req.Numbers)
	res := Response{
		SortedArray: tmp,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/bubbleSort", handler)
	log.Fatal(http.ListenAndServe(":8000", router))
}
