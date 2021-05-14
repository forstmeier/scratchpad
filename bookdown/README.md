# BookDown

> Get your book out there :book:

[![Go Report Card](https://goreportcard.com/badge/github.com/forstmeier/bookdown)](https://goreportcard.com/report/github.com/forstmeier/bookdown)

## Description

BookDown is a Go command line application that turns Markdown files into a HTML files formatted and ready for publishing somehwere like [GitHub Pages](https://pages.github.com/) to be read as a book. I originally built this to create a clean/simple format to publish a side project of my own work. [Hugo](https://gohugo.io/) has some great themes but none of them were quite what I wanted and I felt like building something from the ground up.  

## Usage

Right now the options for usage are fairly limited but can be seen below.

```
bookdown
  [ --title ]
  [ --dest ]
  [ --color ]
  [ --exclude ]
  [ --author ]
  [ --desc ]
  [ --keywords ]
  [ --darktheme ]
  [ <files> ]
```

**Note**: at least one `index.md` file should be provided either as a command line argument or in the directory in which `bookdown` is being run.  
**Note**: this documentation may not necessarily be as up-to-date as the code; to see the exact flag descriptions, run the binary with the `--help` flag.  
