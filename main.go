package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
)

type Page struct {
	Title   string
	Content string
}

func main() {
	err := os.MkdirAll("docs", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	page := Page{
		Title:   "My GitHub Pages Site",
		Content: "Welcome to my static website generated with Go!",
	}

	tmpl := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            max-width: 800px;
            margin: 0 auto;
            padding: 20px;
            line-height: 1.6;
        }
        header {
            text-align: center;
            padding: 20px 0;
            border-bottom: 1px solid #eee;
        }
        main {
            padding: 20px 0;
        }
    </style>
</head>
<body>
    <header>
        <h1>{{.Title}}</h1>
    </header>
    <main>
        <p>{{.Content}}</p>
    </main>
</body>
</html>
`

	t, err := template.New("webpage").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	outputPath := filepath.Join("docs", "index.html")
	f, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = t.Execute(f, page)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Static site generated successfully in the 'docs' directory")
}
