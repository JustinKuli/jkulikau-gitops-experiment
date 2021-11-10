package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		pprint(req)
		fmt.Fprintf(w, "recieved\n")
	})

	fmt.Println("Starting server...")
	http.ListenAndServe(":8080", nil)
}

func pprint(req *http.Request) {
	fmt.Printf("Method: <%s>\n", req.Method)

	fmt.Println("========== Header start ==========")
	for key, values := range req.Header {
		for i, val := range values {
			fmt.Printf("Key <%s>, Value <%v>: <%s>\n", key, i, val)
		}
	}
	fmt.Println("========== Header end ============")

	fmt.Println("========== Body start ==========")
	b, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error while reading body:", err)
	} else {
		fmt.Printf("%s\n", b)
	}
	fmt.Println("========== Body end ============")
}
