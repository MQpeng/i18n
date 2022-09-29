package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/i18n/pkg/command"
	"github.com/i18n/pkg/core"
)

func main() {
	command.Parse()

	if *command.Write{
		log.Println("===========Write Start==============")
		source, err := core.ByGlob(*command.File)
		log.Println("source files: ", source)
		if err != nil{
			log.Fatalln(err);
			return
		}
		var wg sync.WaitGroup
		for _, v := range source {
			wg.Add(1)
			go func(file_name string, wg *sync.WaitGroup){
				fi, err := os.Open(file_name)
				if err != nil {
					log.Printf("Error: %s\n", err)
					return
				}
				// fi.Close()

				// build Json Map
				strMap := core.Json2Map(*command.Language)
				// log.Println("strMap", strMap)
				strMapNew := core.Json2Map(*command.NewLanguage)
				strMapNew = core.ReverseMap(&strMapNew)
				// log.Println("strMapNew", strMapNew["AN状态"])
				// Read By Line
				br := bufio.NewReader(fi)
				newContent := []string{};
				for {
					a, _, c := br.ReadLine()
					if c == io.EOF {
						break
					}
					str := core.RegexForLine(a, *command.Regex)
					log.Println("RegexForLine str", str)
					if 0 != len(str) {
						str = core.RegexForLine(str[0],"'(.*)'")
						for _, v := range str {

							sa := strings.Split(string(v), "'")
							log.Println("sa", sa)
							for _, sa_key := range sa {
								
								strKey := strings.ReplaceAll(sa_key, "'", "")
								str, ok := strMap[strKey];
								log.Println("key", strKey)
								if ok{
									str2, ok2 := strMapNew[str];
									if ok2{
										log.Println("ok2", str2)
										a = bytes.Replace(a, []byte(strKey),[]byte(str2), -1)
									}
								}
							}

						}
					}
					// log.Println(string(a))
					newContent = append(newContent, string(a));
				}
				

				// Write File
				fi.Close()

				fw,_ := os.OpenFile(file_name,os.O_RDONLY|os.O_CREATE|os.O_TRUNC, 0666);
				defer fw.Close()
				write := bufio.NewWriter(fw);
				for _, v := range newContent {
					write.WriteString(v+"\n");
				}
				err = write.Flush()
				if err != nil{
					log.Fatalln("Write to flush：" , err)
				}
				wg.Done()
			}(v, &wg)
		}
		wg.Wait()
		log.Println("===========Write End==============")
	}
}
