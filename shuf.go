package main

import (
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {

	var dir string
	flag.StringVar(&dir, "dir", "", "directory")
	flag.Parse()

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var odd []string
	var even []string

	for i := 0; i < len(files); i++ {
		if i % 2 == 0 {
			odd = append(odd, files[i].Name())
		}
		if i % 2 != 0 {
			even = append(even, files[i].Name())
		}
	}

	if len(odd) != len(even) {
		log.Fatal(err)
	}

	shufOdd := make([]string, len(odd))
	shufEven := make([]string, len(even))

	copy(shufOdd, odd)
	copy(shufEven, even)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(shufOdd), func(i, j int) {
		shufOdd[i], shufOdd[j] = shufOdd[j], shufOdd[i]
		shufEven[i], shufEven[j] = shufEven[j], shufEven[i]
	})

	err = os.Chdir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for i := range odd {
		err := os.Rename(odd[i], shufOdd[i] + ".tmp")
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := range even {
		err := os.Rename(even[i], shufEven[i] + ".tmp")
		if err != nil {
			log.Fatal(err)
		}
	}

	files, err = ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for i := range files {
		err := os.Rename(files[i].Name(), strings.TrimSuffix(files[i].Name(), filepath.Ext(files[i].Name())))
		if err != nil {
			log.Fatal(err)
		}
	}

}
