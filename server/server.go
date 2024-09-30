package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse the form data
		r.ParseForm()
		text := r.FormValue("text")

		// Do something with the text (e.g., print it)
		fmt.Println("Received text:", text)
	}

	// Serve the HTML form
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
    <title>GO Sandbox</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 20px;
        }
        .editor-container {
            position: relative;
            width: 100%%;
            max-width: 800px;
            margin: 0 auto;
        }
        .editable {
            width: 100%%;
            height: 500px;
            padding: 10px;
            font-size: 16px;
            border: 1px solid #ccc;
            box-sizing: border-box;
            overflow: auto;
            resize: vertical;
        }
        .submit-button {
            position: absolute;
            bottom: 20px;
            right: 20px;
            padding: 10px 20px;
            font-size: 16px;
        }
    </style>
</head>
<body>
    <h1>Write Go Code</h1>
    <form method="POST" action="/" onsubmit="prepareSubmission()">
        <div class="editor-container">
            <div class="editable" contenteditable="true" id="editor"></div>
            <button type="submit" class="submit-button">Submit</button>
        </div>
        <input type="hidden" name="text" id="hiddenInput">
    </form>
    <script>
        function prepareSubmission() {
            var editorContent = document.getElementById('editor').innerText;
            document.getElementById('hiddenInput').value = editorContent;
        }
    </script>
</body>
</html>
        `)
}

func main() {
	// Register the handler function for the root URL path
	http.HandleFunc("/", handler)

	// Start the HTTP server on port 8080
	fmt.Println("Server is running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
