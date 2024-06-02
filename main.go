package main

import (
	"fmt"
	"log"
	"os"
	"text/template"

	flag "github.com/spf13/pflag"
)

const dailyTemplateFile = "./templates/daily.tmpl.md"

var (
	tmpl          *template.Template
	n             *int    = flag.IntP("number", "n", 1, "Number of documents to create")
	initialOffset *int    = flag.IntP("offset", "o", 0, "Number of offset days (from today) to start documents from")
	dir           *string = flag.StringP("dir", "C", "", "Change to dir before running the command")
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	if *dir != "" {
		err := os.Chdir(*dir)
		if err != nil {
			return err
		}
	}

	for i := 0; i < *n; i++ {
		d := getDateForOffset(i)
		err := writeFile(d)
		if err != nil {
			return err
		}
	}

	return nil

}

func writeFile(d DateObject) error {
	filename := fmt.Sprintf("%s.md", d.Date)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, d)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	var err error
	tmpl, err = template.ParseFiles(dailyTemplateFile)
	if err != nil {
		log.Fatalln(err)
	}

	flag.Parse()
	if *n < 1 {
		*n = 1
	}
	if *initialOffset < 0 {
		*initialOffset = 0
	}
}
