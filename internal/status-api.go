package internal

import (
	"fmt"
    "net/http"
	"github.com/fatih/color"
	_ "net/http/pprof"

)


func StatusAPI() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Requests per second: %f\n\n", RPS)
		fmt.Fprintf(w, "Status Codes | Requests\n")
		for code, count := range StatusCodes {
			fmt.Fprintf(w, "%s: %d\n", code, count)
		}

    })
    
	// Start messages
	Cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
	fmt.Printf("[%s] Status server started on port 8899\n", Cyan("INF"))

	go http.ListenAndServe(":8899", nil)
	
}
