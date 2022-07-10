package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/sriharish/stuff/cmdgen"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, cmdgen.GenerateASCII(), r.URL.Path)
	})

	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, Modules!")
		},
	}

	fmt.Println("Calling cmd.Execute()!")
	cmd.Execute()

	http.ListenAndServeTLS(":8081", "", "", nil)
	http.ListenAndServe(":8080", nil)
}
