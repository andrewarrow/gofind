package main

import "fmt"
import "os"
import "path/filepath"
import "strings"
import "strconv"

var find = ""
var hits = 0
var which = 0

func visit(path string, f os.FileInfo, err error) error {
	lpath := strings.ToLower(path)
	if strings.Contains(lpath, "test") {
		return nil
	}
	if strings.Contains(path, ".kt") {
		tokens := strings.Split(path, "/")
		last := tokens[len(tokens)-1]
		more := strings.Split(last, ".")
		if strings.Contains(more[0], find) {
			hits++
			if which > 0 && which == hits {
				fmt.Println("vi", path[26:])
			} else if which == 0 {
				fmt.Println(hits, path[26:])
			}
		}
	}
	return nil
}

func main() {
	args := os.Args
	find = args[1]
	if len(args) > 2 {
		which, _ = strconv.Atoi(args[2])
	}
	pwd, _ := os.Getwd()
	filepath.Walk(pwd, visit)
}
