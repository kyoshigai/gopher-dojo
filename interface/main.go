package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Print(w io.Writer, r io.Reader) error {
	br := bufio.NewReader(r)

	for {
		s, err := br.ReadString('\n')
		if s != "" {
			io.WriteString(w, s)
		}

		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
	}
}

func PrintFromString(s string) error {
	r := strings.NewReader(s)
	if err := Print(os.Stdout, r); err != nil {
		return err
	}
	return nil
}

func PrintFromFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	Print(os.Stdout, f)

	defer func() {
		if cErr := f.Close(); err != nil {
			err = cErr
		}
	}()
	return nil
}

func main() {
	err := PrintFromString("hoge\nfoo\n")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("------------------")

	err = PrintFromFile("./print_from_file_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Check whether f implements io.Reader or not.
	// f, _ := os.Open("./print_from_file_test.txt")
	// var _ io.Reader = f
}
