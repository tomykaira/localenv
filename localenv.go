package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/oleiade/trousseau"
)

const (
	SEPARATOR = "/%/"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Command is not specified.")
		os.Exit(1)
	}

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	switch os.Args[1] {
	case "get":
		if len(os.Args) < 3 {
			log.Fatal("get key")
		}
		key := os.Args[2]

		tr, err := trousseau.OpenTrousseau(trousseau.InferStorePath())
		if err != nil {
			log.Fatal(err)
		}

		store, err := tr.Decrypt()
		if err != nil {
			log.Fatal(err)
		}

		keys := store.Data.Keys()
		dir := pwd
		for dir != "." && dir != "/" {
			for _, k := range keys {
				if k == dir+SEPARATOR+key {
					data, err := store.Data.Get(k)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Print(data)
					os.Exit(0)
				}
			}
			dir = path.Dir(dir)
		}
		os.Exit(1)

	case "set":
		if len(os.Args) < 4 {
			log.Fatal("set key value")
		}
		trousseau.SetAction(pwd+SEPARATOR+os.Args[2], os.Args[3], "")

	case "list":
		tr, err := trousseau.OpenTrousseau(trousseau.InferStorePath())
		if err != nil {
			log.Fatal(err)
		}

		store, err := tr.Decrypt()
		if err != nil {
			log.Fatal(err)
		}

		keys := store.Data.Keys()
		dir := pwd
		results := make(map[string]interface{})
		for dir != "." && dir != "/" {
			for _, k := range keys {
				if strings.HasPrefix(k, dir+SEPARATOR) {
					key := strings.TrimPrefix(k, dir+SEPARATOR)
					if _, ok := results[key]; !ok && !strings.HasPrefix(key, "_") {
						results[key], err = store.Data.Get(k)
						if err != nil {
							log.Fatal(err)
						}
					}
				}
			}
			dir = path.Dir(dir)
		}

		for k, v := range results {
			fmt.Printf("'%s'='%s'\n", k, v)
		}

	default:
		log.Fatal("Unknown command " + os.Args[1])
	}
}
