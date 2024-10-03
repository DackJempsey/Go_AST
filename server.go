package main

import (
	"fmt"
	"log"
	"net/http"
    "go_ast/compler"
    "path/filepath"
    "io/ioutil"
    "html/template"
)

var templates *template.Template

func handler(w http.ResponseWriter, r *http.Request) {

    comp_return := ""
	
    if r.Method == http.MethodPost {
		// Parse the form data
		r.ParseForm()
		text := r.FormValue("text")

		// Do something with the text (e.g., print it)
		// fmt.Println("Received text:", text)
        comp_return = compler.Comp(text)
	}

	// Serve the HTML form
    srcFile := filepath.Join(".", "index.html")
	data, _ := ioutil.ReadFile(srcFile)
	fmt.Fprintf(w, string(data))

    w.Header().Set("Content-Type", "text/html")
    
    err := templates.ExecuteTemplate(w, "results.html", comp_return)
    if err != nil {
        http.Error(w, "Error rendering template.", http.StatusInternalServerError)
        log.Println("Template execution error:", err)
    }
}

func formHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    err := templates.ExecuteTemplate(w, "index.html", nil)
    
    if err != nil {
        http.Error(w, "Error rendering template.", http.StatusInternalServerError)
        log.Println("Template execution error:", err)
    }
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
        return
    }

    // Parse the form data
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Error parsing form data.", http.StatusBadRequest)
        return
    }

    // Retrieve the value from the textarea
    textboxValue := r.FormValue("textbox")

    // Process the data (you can modify this part as needed)
    processedData := processData(textboxValue)

    w.Header().Set("Content-Type", "text/html")
    err := templates.ExecuteTemplate(w, "results.html", processedData)
    if err != nil {
        http.Error(w, "Error rendering template.", http.StatusInternalServerError)
        log.Println("Template execution error:", err)
    }
}

// Dummy function to process data
func processData(input string) string {
    // Perform your data processing here
    return "You entered:\n\n" + input
}

func main() {

    templates = template.Must(template.ParseGlob(filepath.Join("./html_templates", "*.html")))


	// Register the handler function for the root URL path
	http.HandleFunc("/", formHandler)
    http.HandleFunc("/submit", handler)

	// Start the HTTP server on port 8080
	fmt.Println("Server is running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
