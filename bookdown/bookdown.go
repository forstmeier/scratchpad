package main

import (
	"bytes"
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/russross/blackfriday"
	"gopkg.in/go-playground/colors.v1"
)

type sliceFlag []string

func (sf *sliceFlag) String() string {
	return strings.Join(*sf, " ")
}

func (sf *sliceFlag) Set(value string) error {
	*sf = append(*sf, value)
	return nil
}

type content struct {
	Title    string
	Body     template.HTML
	Desc     string
	Keywords string
	Author   string
	Next     string
	Prev     string
}

func main() {
	title := flag.String("title", "BookDown", "book title displayed in browser tab")
	dest := flag.String("dest", "", "html file output directory location")
	color := flag.String("color", "#DB5525", "color of the top ribbon and links")
	darktheme := flag.Bool("darktheme", false, "generate html file output in dark mode")
	desc := flag.String("desc", "", "meta tag for the content description")
	keywords := flag.String("keywords", "", "meta tag for page keywords")
	author := flag.String("author", "", "meta tag for the content author name")

	var exclude sliceFlag
	flag.Var(&exclude, "exclude", "files to keep out of html generation")

	flag.Parse()

	files := flag.Args()

	box := packr.NewBox(".")
	htmlFile, err := box.FindString("resources/bookdown.html")
	if err != nil {
		log.Fatal(err)
	}
	cssFile, err := box.FindString("resources/bookdown.css")
	if err != nil {
		log.Fatal(err)
	}
	faviconFile, err := box.FindString("resources/favicon.ico")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := colors.ParseHEX(*color); err != nil {
		log.Fatal(err)
	}

	if *dest != "" {
		if src, err := os.Stat(*dest); os.IsNotExist(err) {
			log.Fatal("output directory does not exist")
		} else if !src.IsDir() {
			log.Fatal("non-directory provided in dest flag")
		}
	}

	if *color != "" {
		cssFile = strings.Replace(cssFile, "#DB5525", *color, -1)
	}

	if *darktheme {
		cssFile = strings.Replace(cssFile, "#E8E8E8", "#424242", -1)
		cssFile = strings.Replace(cssFile, "#000000", "#FFFFFF", -1)
	}

	if err := ioutil.WriteFile(filepath.Join(*dest, "bookdown.css"), []byte(cssFile), 0644); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(filepath.Join(*dest, "favicon.ico"), []byte(faviconFile), 0644); err != nil {
		log.Fatal(err)
	}

	if len(files) == 0 {
		if err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			notREADME := strings.ToLower(filepath.Base(path)) != "readme.md"
			notLICENSE := strings.ToLower(filepath.Base(path)) != "license.md"
			if filepath.Ext(path) == ".md" && notREADME && notLICENSE {
				files = append(files, path)
			}
			return nil
		}); err != nil {
			log.Fatal(err)
		}
	}

	if len(exclude) != 0 {
		for _, exld := range exclude {
			for i, file := range files {
				if file == exld {
					files = append(files[:i], files[i+1:]...)
				}
			}
		}
	}

	for i, file := range files {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		body := template.HTML(blackfriday.MarkdownCommon(b))

		file = strings.Replace(file, ".md", ".html", -1)

		next, prev := "", ""
		if i < len(files)-1 && file != "index.html" {
			next = strings.Replace(files[i+1], ".md", ".html", -1)
		}
		if i > 0 && file != "index.html" {
			prev = strings.Replace(files[i-1], ".md", ".html", -1)
		}

		cont := content{
			Title:    *title,
			Body:     body,
			Desc:     *desc,
			Keywords: *keywords,
			Author:   *author,
			Next:     next,
			Prev:     prev,
		}

		var buffer bytes.Buffer

		tmpl := template.Must(template.New("bookdown").Parse(htmlFile))

		if err := tmpl.Execute(&buffer, cont); err != nil {
			log.Fatal(err)
		}

		if err := ioutil.WriteFile(filepath.Join(*dest, filepath.Base(file)), buffer.Bytes(), 0644); err != nil {
			log.Fatal(err)
		}
	}
}
