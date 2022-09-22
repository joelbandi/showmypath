package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("can't run this util on windows :(")
	} else {
		out, err := exec.Command("echo", os.Getenv("PATH")).Output()
		if err != nil {
			log.Fatal(err)
			return
		}
		paths := strings.Split(string(out), ":")
		pathsLen := len(paths)

		for i, path := range paths {
			fmt.Printf("%d. %s%s\n", i+1, indent(i+1, pathsLen), path)
		}
	}
}

var indentCache map[int]string = map[int]string{}

func indent(current int, total int) string {
	strCurrentLen := len(strconv.Itoa(current))

	if val, ok := indentCache[strCurrentLen]; ok {
		return val
	}

	strTotalLen := len(strconv.Itoa(total))
	deltaIndent := strTotalLen - strCurrentLen

	var builder strings.Builder
	builder.WriteString("")
	for i := 0; i < deltaIndent; i++ {
		builder.WriteString(" ")
	}
	result := builder.String()

	indentCache[strCurrentLen] = result

	return result
}
