package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func input(r io.Reader) <-chan string {
	// TODO: チャネルを作る
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			// TODO: チャネルに読み込んだ文字列を送る
			ch <- s.Text()
		}
		// TODO: チャネルを閉じる
		close(ch)
	}()
	// TODO: チャネルを返す
	return ch
}

func main() {
	ch := input(os.Stdin)
	for {
		fmt.Print(">")

		if text, ok := <-ch; !ok {
			break
		} else {
			fmt.Println(text)
		}
	}
}
