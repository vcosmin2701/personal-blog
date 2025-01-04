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

	err = copyDir("static", "docs")
	if err != nil {
		log.Fatal(err)
	}

	page := Page{
		Title:   "My GitHub Pages Site",
		Content: "Welcome to my static website generated with Go!",
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	outputPath := filepath.Join("docs", "index.html")
	f, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = tmpl.Execute(f, page)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Static site generated successfully in the 'docs' directory")
}

func copyDir(src string, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		return copyFile(path, dstPath)
	})
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = in.WriteTo(out)
	return err
}
