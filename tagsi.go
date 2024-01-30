package main

import (
	"flag"
	"fmt"
	"regexp"
)

var (
	// Capture everything between the backticks denoting the struct tag
	StructTagPattern = regexp.MustCompile("\x60([a-zA-Z0-9: -,\"]+)\x60")
)

type Args struct {
	Buf *string
}

func ProcessArgs() *Args {
	buf := flag.String("buf", "", "Struct Tag Line")
	flag.Parse()

	return &Args{
		Buf: buf,
	}
}

//func GetStructTags(buf string) []string {
//
//}

func main() {
	args := ProcessArgs()

	matches := StructTagPattern.FindStringSubmatch(*args.Buf)
	if matches != nil {
		fmt.Println(matches[1])
	}

}
