package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func get_index_page(filename string) (string, error) {

	var s_html string
	html, err := os.ReadFile(filename)
	if err != nil {
		s_html = ""
	}
	s_html = string(html)
	return s_html, err
}

func index(w http.ResponseWriter, req *http.Request) {
	page, err := get_index_page("src/index.html")
	if err != nil {
		page = "<!DOCTYPE html> <h1> Page not found </h1>"
	}
	fmt.Fprintf(w, page)
}

func script(w http.ResponseWriter, req *http.Request) {
	page, err := get_index_page("src/script.js")
	if err != nil {
		page = "<!DOCTYPE html> <h1> Page not found </h1>"
	}
	fmt.Fprintf(w, page)
}

func style(w http.ResponseWriter, req *http.Request) {
	page, err := get_index_page("src/styles/style.css")
	if err != nil {
		page = "<!DOCTYPE html> <h1> Page not found </h1>"
	}
	fmt.Fprintf(w, page)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/script.js", script)
	http.HandleFunc("/styles/style.css", style)
	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed \n")
	} else if err != nil {
		fmt.Printf("error starting server: %s", err)
		os.Exit(1)
	}
	return
	// x1, x2, code := solver.Solve_quadratic(1, 2, 1)
	// if code == -1 {
	// 	fmt.Println("Ну, в этот раз Вам не повезло")
	// 	return
	// } else if code == 1 {
	// 	fmt.Printf("Решений нет\n")
	// } else {
	// 	fmt.Printf("x1 = %f\nx2 = %f\n", x1, x2)
	// }
}
