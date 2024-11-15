package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/cmackenzie1/go-uuid"
)

func run(version int, count int, uppercase bool) error {
	for i := 0; i < count; i++ {
		var id uuid.UUID
		var err error
		if version == 4 {
			id, err = uuid.NewV4()
		} else {
			id, err = uuid.NewV7()
		}
		if err != nil {
			return err
		}
		if uppercase {
			fmt.Println(strings.ToUpper(id.String()))
		} else {
			fmt.Println(id)
		}
	}
	return nil
}

func main() {
	v := flag.Int("v", 4, "UUID version to generate.Supported versions are 4 and 7.")
	n := flag.Int("c", 1, "Number of UUIDs to generate.")
	u := flag.Bool("u", false, "Print UUIDs in uppercase.")
	flag.Parse()

	if *v != 4 && *v != 7 {
		fmt.Printf("Unsupported UUID version: %d\n", *v)
		os.Exit(1)
	}

	if *n < 1 {
		fmt.Printf("Number of UUIDs to generate must be greater than 0.\n")
		os.Exit(1)
	}

	if err := run(*v, *n, *u); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
