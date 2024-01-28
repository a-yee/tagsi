package main

import (
	"flag"
	"fmt"
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

	fmt.Printf(*args.Buf)
}
