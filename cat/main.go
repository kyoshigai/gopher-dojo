package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	n := flag.Bool("n", false, "Number the output lines starting at 1")
	flag.Parse()
	filePaths := flag.Args()
	lineNums := 1

	for _, filePath := range filePaths {
		func() {
			f, err := os.Open(filePath)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			sc := bufio.NewScanner(f)
			for sc.Scan() {
				text := sc.Text()
				if *n {
					fmt.Printf("%d: %s\n", lineNums, text)
				} else {
					fmt.Println(text)
				}
				lineNums++
			}
		}()
	}
}
