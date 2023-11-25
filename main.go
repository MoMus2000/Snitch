package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type Snitch struct {
	Extension       string
	ListOfFileNames []string
	Keyword         string
	Priority        bool
}

// What Flags do i need?
// Snith is going to be a cli to tool, to get TODO comments and filename
// I need the file extension
func main() {
	var Extension string
	var Keyword string
	var Priority bool
	flag.StringVar(&Extension, "extension", "*.*", "Select the extension of the file to search.")
	flag.StringVar(&Keyword, "keyword", "TODO", "Select the keyword by which you want to search.")
	flag.BoolVar(&Priority, "priority", true, "Sort the keyword by priority.")

	flag.Parse()

	sn := Snitch{Extension: Extension, ListOfFileNames: make([]string, 0)}

	WalkThroughDirectory(&sn)

	fmt.Println(sn.ListOfFileNames)

	ParseFile(&sn)
}

func WalkThroughDirectory(sn *Snitch) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	files, err := os.ReadDir(dir)
	for _, file := range files {
		if file.IsDir() {
			os.Chdir(file.Name())
			WalkThroughDirectory(sn)
			os.Chdir("..")
		} else {
			wd, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			sn.ListOfFileNames = append(sn.ListOfFileNames, wd+"/"+file.Name())
		}
	}
}

// TODO: Add functionality to read todo.

func ParseFile(sn *Snitch) {
	files := sn.ListOfFileNames

	for _, file := range files {
		openedFile, err := os.OpenFile(file, os.O_RDONLY, 0755)
		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(openedFile)

		lineNumber := 1

		for scanner.Scan() {
			buffer := string(scanner.Bytes())
			if len(buffer) > 9 && buffer[0:8] == "// TODO:" {
				fmt.Println(file, lineNumber)
			}
			lineNumber += 1
		}
	}
}
