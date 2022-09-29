package core

import (
	"log"
	"regexp"
)

func RegexForLine(line []byte, pattern string) [][]byte {
	p, e := regexp.Compile(pattern);
	if e != nil{
		log.Fatalln("compile error for pattern:", pattern);
	}
	// log.Println(p.FindAllString(string(line), -1))
	return p.FindAll(line, -1)
}