package main

import (
	"fmt"
	"io/ioutil"
	"flag"
	"regexp"
	"os"
)

func main () {
	var prefix = flag.String("prefix", "", "Prefix string you want to remove.")

	flag.Parse()

	files, err := ioutil.ReadDir("./")

	if err != nil {
		fmt.Printf("Something wrong during reading files from current directory %d", err)
		os.Exit(1)
	}

	r := regexp.MustCompile(fmt.Sprintf("^%s(.+)$", *prefix))
	renamedCount := 0

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()

		if r.MatchString(fileName) {
			newFileName := r.ReplaceAllString(fileName, "$1")
			err := os.Rename(fileName, newFileName)

			if err != nil {
				fmt.Printf("Error occurred when renaming %s to %s: %d\n", fileName, newFileName, err)
				continue
			}

			renamedCount++
			fmt.Printf("%s -> %s\n", fileName, newFileName)
			// renamedMessages := append(renamedMessages, fmt.Sprintf("%d -> %d", fileName, newFileName))
		}
	}

	fmt.Printf("%d files are renamed", renamedCount)
}
