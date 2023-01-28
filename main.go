package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/TwiN/go-color"
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
    var lines []string

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    file.Close()

    for _, line := range lines {
        if strings.Contains(line, "Warning") {
            fmt.Println(color.Colorize(color.Yellow, line))
        }
        if strings.Contains(line, "Error") {
            fmt.Println(color.Colorize(color.Red, line))
        }
    }
}
