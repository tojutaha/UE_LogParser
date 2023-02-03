package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

    argLen := len(os.Args[1:])
    if argLen == 0 {
        fmt.Print("Usage: ", filepath.Base(os.Args[0]), " Filename")
        return
    }

    path := os.Args[1]
    file, err := os.Open(path)
    if err != nil {
        fmt.Println(err)
        return
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    for scanner.Scan() {
        if strings.Contains(scanner.Text(), "Warning") {
            fmt.Println("\033[33m", scanner.Text())
        }
        if strings.Contains(scanner.Text(), "Error") {
            fmt.Println("\033[31m", scanner.Text())
        }
    }

    fmt.Println("\033[0m")
    file.Close()
}
