package test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/i18n/pkg/core"
)

func TestRegexForLine(t *testing.T) {
	fi, err := os.Open("./code.ts")
    if err != nil {
        fmt.Printf("Error: %s\n", err)
        return
    }
    defer fi.Close()

    br := bufio.NewReader(fi)
    for {
        a, _, c := br.ReadLine()
        if c == io.EOF {
            break
        }
		str := core.RegexForLine(a, `\$L\('.*'\)`)
		if 0 != len(str) {
			str = core.RegexForLine(str[0],"'(.*)'")
			result := []string{};
			for _, v := range str {
				strings.Split(string(v),"'")
				result = append(result, string(v))
			}
			fmt.Println(result)
		}
        
    }
	t.Fatalf("")
}