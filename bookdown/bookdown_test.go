package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

var binary = "bookdown"

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func cleanupFiles() {
	files, err := ioutil.ReadDir(".")
	errorCheck(err)

	for _, file := range files {
		if strings.Contains(file.Name(), "test-") || filepath.Ext(file.Name()) == ".css" || filepath.Ext(file.Name()) == ".ico" {
			if err := os.Remove(file.Name()); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func cleanupDir() {
	if _, err := os.Stat("tatooine/"); !os.IsNotExist(err) {
		if err := os.Remove("tatooine/"); err != nil {
			log.Fatal(err)
		}
	}
}

func writeTestFiles(input map[string][]byte) {
	for name, content := range input {
		if err := ioutil.WriteFile(name, content, 0644); err != nil {
			log.Fatal(err.Error())
		}
	}
}

func countHTMLFiles(files []os.FileInfo) int {
	count := 0
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".html" {
			count++
		}
	}
	return count
}

func checkHTMLContent(exist bool) string {
	content := ""
	if exist {
		b, err := ioutil.ReadFile("test-book.html")
		errorCheck(err)

		content = string(b)
	}
	return content
}

func checkCSSContent(exist bool) string {
	content := ""
	if exist {
		b, err := ioutil.ReadFile("bookdown.css")
		errorCheck(err)

		content = string(b)
	}
	return content
}

func TestMain(m *testing.M) {
	build := exec.Command("go", "build", "bookdown.go")
	if err := build.Run(); err != nil {
		log.Fatalf("failure generating binary: %s\n", err.Error())
	}

	if err := os.Mkdir("tatooine/", 0644); err != nil {
		log.Fatal(err)
	}

	code := m.Run()

	if err := os.Remove(binary); err != nil {
		log.Fatalf("failure removing ginary: %s\n", err.Error())
	}

	os.Exit(code)
}

func TestExecutable(t *testing.T) {
	tests := []struct {
		desc    string
		input   map[string][]byte
		output  []string
		args    []string
		status  string
		exist   bool
		count   int
		content []string
	}{
		{
			desc:    "no files scenario",
			input:   make(map[string][]byte),
			output:  []string{},
			args:    []string{},
			status:  "",
			exist:   false,
			count:   0,
			content: []string{""}},
		{
			desc: "one book file scenario",
			input: map[string][]byte{
				"test-book.md": []byte(`
				# Tatooine
				> A desert world
				Mostly scum and villainy
				`),
			},
			output:  []string{"test-book.html"},
			args:    []string{},
			status:  "",
			exist:   true,
			count:   1,
			content: []string{"BookDown", "Tatooine"},
		},
		{
			desc: "title arg and book file scenario",
			input: map[string][]byte{
				"test-book.md": []byte(`
				# Naboo
				> A lush world of peace
				- Humans (land-based)
				- Gungans (sea/swamp-based)
				`),
			},
			output:  []string{"test-book.html"},
			args:    []string{"--title", "Naboo"},
			status:  "",
			exist:   true,
			count:   1,
			content: []string{"Naboo"},
		},
		{
			desc: "title arg, book file, and chapter file scenario",
			input: map[string][]byte{
				"test-book.md": []byte(`
				# Coruscant
				> A world-spanning metropolis
				`),
				"test-chapter-1.md": []byte(`
				# Battle of Coruscant
				| GAR | CIS |
				| --- | --- |
				| Kenobi | Tyranus |
				| Skywalker | Grievous |
				`),
			},
			output:  []string{"test-book.html", "test-chapter-1.html"},
			args:    []string{"--title", "Coruscant"},
			status:  "",
			exist:   true,
			count:   2,
			content: []string{"Coruscant"},
		},
		{
			desc: "dest arg to non-existent directory scenario",
			input: map[string][]byte{
				"test-book.md": []byte(`
				# Geonosis
				> A desert world
				The first engagement of the Clone Wars
				`),
			},
			output:  []string{},
			args:    []string{"--dest", "mos-eisley"},
			status:  "exit status 1",
			exist:   false,
			count:   0,
			content: []string{},
		},
		{
			desc: "dest arg to existing directory scenario",
			input: map[string][]byte{
				"test-book.md": []byte(`
				# Utapau
				> A windy planet of sinkholes
				- [x] pleasant greetings
				- [x] bold ones
				- [ ] war (unless you brought it with you)
				`),
			},
			output:  []string{},
			args:    []string{"--dest", "tatooine"},
			status:  "",
			exist:   false,
			count:   0,
			content: []string{},
		},
		{
			desc: "color arg scenario",
			input: map[string][]byte{
				"test-book.md": []byte(`
				# Mustafar
				> A planet of lava
				Battle of the heroes
				`),
			},
			output:  []string{"test-book.html"},
			args:    []string{"--color", "#DC143C"},
			status:  "",
			exist:   true,
			count:   1,
			content: []string{},
		},
		{
			desc: "exclude arg scenario",
			input: map[string][]byte{
				"test-book.md": []byte(`
				# Mygeeto
				> A world of ice crystals
				`),
				"test-chapter-1.md": []byte(`
				- Caar Damask + Darth Tenebrous = Darth Plagueis
				`),
				"test-chapter-2.md": []byte(`
				- Ki-Adi-Mundi + Order 66 = dead
				`),
			},
			output:  []string{"test-chapter-1.html", "test-chapter-2.html"},
			args:    []string{"--exclude", "test-book.md"},
			status:  "",
			exist:   false,
			count:   2,
			content: []string{},
		},
		{
			desc: "page traversal hyperlinks",
			input: map[string][]byte{
				"test-book.md": []byte(`
				# Felucia
				> Living world of fungi
				`),
				"test-chapter-1.md": []byte(`
				- Battle of Felucia
				`),
				"test-chapter-2.md": []byte(`
				- Hunt for Shaak Ti
				`),
			},
			output:  []string{"test-book.html", "test-chapter-1.html", "test-chapter-2.html"},
			args:    []string{},
			status:  "",
			exist:   true,
			count:   3,
			content: []string{"Next"},
		},
		{
			desc: "meta tag args",
			input: map[string][]byte{
				"test-book.md": []byte(`
				# Saleucami
				> An arid oasis planet
				`),
			},
			output:  []string{"test-book.html"},
			args:    []string{"--author", "Stass Allie", "--keywords", "fallen"},
			status:  "",
			exist:   true,
			count:   1,
			content: []string{"Stass Allie", "fallen"},
		},
		{
			desc: "dark theme output",
			input: map[string][]byte{
				"test-book.md": []byte(`
				# Kashyyyk
				> The Wookie world
				- Worshyr trees
				- More worshyr trees
				- All the worshyr trees
				`),
			},
			output:  []string{"test-book.html"},
			args:    []string{"--darktheme"},
			status:  "",
			exist:   true,
			count:   1,
			content: []string{},
		},
	}

	for _, test := range tests {
		writeTestFiles(test.input)

		run := exec.Command("./"+binary, test.args...)
		if err := run.Run(); err != nil {
			cleanupFiles()
			cleanupDir()
		}

		t.Run("check for output files", func(t *testing.T) {
			files, err := ioutil.ReadDir(".")
			errorCheck(err)

			for _, expected := range test.output {
				result, file := func(rec []os.FileInfo, exp string) (bool, string) {
					for _, r := range rec {
						if r.Name() == exp {
							return true, exp
						}
					}
					return false, exp
				}(files, expected)
				if !result {
					t.Errorf("%s - %s", file, test.desc) // TEMP
				}
			}
		})

		t.Run("check output file count", func(t *testing.T) {
			files, err := ioutil.ReadDir(".")
			errorCheck(err)

			received := countHTMLFiles(files)

			if received != test.count {
				t.Errorf("desc: %s, expected: %d, received: %d", test.desc, test.count, received)
			}
		})

		t.Run("check test-book.html file content", func(t *testing.T) {
			received := checkHTMLContent(test.exist)

			for _, expected := range test.content {
				if !strings.Contains(received, expected) {
					t.Errorf("desc: %s, expected: %s, received: %s", test.desc, expected, received)
				}
			}
		})

		t.Run("check bookdown.css file content", func(t *testing.T) {
			received := checkCSSContent(test.exist)

			for _, arg := range test.args {
				if strings.Contains(arg, "#") {
					if strings.Contains(received, "#DB5525") {
						t.Errorf("desc: %s, expected: %s, received: %s", test.desc, arg, "#DB5525")
					}
				}

				if arg == "darktheme" {
					if strings.Contains(received, "#FFFFFF") || strings.Contains(received, "#424242") {
						t.Errorf("desc: %s, expected: %s, received: %s", test.desc, arg, "#DB5525")
					}
				}
			}
		})

		cleanupFiles()
	}

	cleanupFiles()
	cleanupDir()
}
