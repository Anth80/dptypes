package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	gen := flag.String("gen", "", "demand paging type to generate")
	pkg := flag.String("pkg", "", "package")
	out := flag.String("out", "", "output file name without extension")
	name := flag.String("name", "", "name of generated type")
	dataType := flag.String("name", "", "data type")
	flag.Parse()

	if *gen == "" {
		fmt.Println("Mandatory option gen missing")
		os.Exit(1)
	}
	if *pkg == "" {
		fmt.Println("Mandatory option pkg missing")
		os.Exit(1)
	}
	if *out == "" {
		fmt.Println("Mandatory option out missing")
		os.Exit(1)
	}
	if *name == "" {
		fmt.Println("Mandatory option name missing")
		os.Exit(1)
	}
	if *dataType == "" {
		fmt.Println("Mandatory option type missing")
		os.Exit(1)
	}

	var contents []byte
	switch *gen {
	case "map":
		contents = genMap(*name, *pkg, *dataType)
	default:
		panic(fmt.Sprintf("Unknown gen type %v", *gen))
	}
	ioutil.WriteFile(*out, contents, 0644)
}