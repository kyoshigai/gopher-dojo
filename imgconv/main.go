package main

import (
	"flag"
	"fmt"
	"gopher-dojo/imgconv"
	"log"
	"os"
	"path/filepath"
)

/*
TODO:
- Add godoc
*/

func main() {
	dirPath := flag.String("d", "./", "Specify the target dir path")
	flag.Parse()

	err := filepath.Walk(*dirPath,
		func(path string, info os.FileInfo, err error) error {
			ext, err := imgconv.Ext(path)
			if err == imgconv.ErrInvalidExt {
				fmt.Printf("skip: %v\n", path)
			} else if err != nil {
				return err
			}

			file := imgconv.File{Path: path, DestExt: ext}
			file.Convert()
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
}
