package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func executeLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Cpus: %v\n", vars["cpus"])
	fmt.Fprintf(w, "Time in seconds: %v\n", vars["time"])
	fmt.Fprintf(w, "Max number of cpus: %v\n", runtime.NumCPU())
	cpus, err := strconv.Atoi(vars["cpus"])
	fmt.Println(err)
	runtime.GOMAXPROCS(cpus)
	timeSec, err := time.ParseDuration(vars["time"] + "s")
	fmt.Println(err)

	for i := 0; i < cpus; i++ {
		go func() {
			fmt.Println("Thread " + strconv.Itoa(i) + " started")
			var A [1000000]bool
			for i := 0; i < 1000000; i++ {
				A[i] = true
			}
			time.Sleep(timeSec)
			fmt.Println("Thread " + strconv.Itoa(i) + " ended")
		}()
		<-time.After(time.Second * 5)
	}
}

func maxLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Time in seconds: %v\n", vars["time"])
	fmt.Fprintf(w, "Max number of cpus: %v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	timeSec, err := time.ParseDuration(vars["time"] + "s")
	fmt.Println(err)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			fmt.Println("Thread " + strconv.Itoa(i) + " started")
			var A [1000000]bool
			for i := 0; i < 1000000; i++ {
				A[i] = true
			}
			time.Sleep(timeSec)
			fmt.Println("Thread " + strconv.Itoa(i) + " ended")
		}()
		<-time.After(time.Second * 5)
	}
}

func main() {
	cpus := flag.Int("cpus", 1, "The number of desired cpus to be used")
	time := flag.Int("time", 10, "Time in seconds for the command to be executed")
	flag.Parse()

	fmt.Println(*cpus)
	fmt.Println(*time)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/execute/{cpus:[0-9]+}/{time:[0-9]+}", executeLink)
	router.HandleFunc("/max/{time:[0-9]+}", maxLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}
