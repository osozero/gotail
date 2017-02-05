package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var f = flag.String("f", "", "file path to read")

var usage = `Usage: tail -f "file_path_to_read"
`

func usageAndExit() {
	flag.Usage()
	os.Exit(2)
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, fmt.Sprintf(usage))
	}

	flag.Parse()
	if *f == "" {
		log.Println("file path parameter must be given")
		usageAndExit()
	}

	file, err := os.Open(*f)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		fmt.Println(line)
	}
}
