package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var n int
	var c int
	flag.IntVar(&n, "n", 10, "display first n lines")
	flag.IntVar(&c, "c", 0, "display first c bytes")
	flag.Parse()

	if flag.NArg() > 0 {
		for i := 0; i < flag.NArg(); i++ {
			filename := flag.Arg(i)
			if flag.NArg() > 1 {
				fmt.Printf("==> %s <==\n", filename)
			}
			handleFile(filename, n, c)
			if flag.NArg() > 1 && i < flag.NArg()-1 {
				fmt.Println()
			}
		}
	} else {
		handleFile("", n, c)
	}

}

func handleFile(filename string, nFlag int, cFlag int) {
	var f *os.File
	var err error
	if filename == "" {
		f = os.Stdin
	} else {
		f, err = os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			if err = f.Close(); err != nil {
				log.Fatal(err)
			}
		}()
	}
	if cFlag == 0 {
		readLines(f, nFlag)
	} else {
		readBytes(f, cFlag)
		fmt.Println()
	}
}

func readLines(f *os.File, n int) {
	s := bufio.NewScanner(f)

	linesRead := 0
	for linesRead < n {
		if !s.Scan() {
			break
		}
		err := s.Err()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(s.Text())
		linesRead++
	}

}

func readBytes(f *os.File, n int) {
	b := make([]byte, 1)
	bytesRead := 0
	for bytesRead < n {
		n, err := f.Read(b)
		if n == 0 {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(b[0]))
		bytesRead++
	}
}
