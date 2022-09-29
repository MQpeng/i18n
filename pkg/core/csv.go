package core

import (
	"encoding/csv"
	"log"
	"os"
)

func Export(column []string, data [][]string, file_name string)error {
    file, err := os.Create(file_name)
    if err != nil {
        log.Println("open file is failed, err: ", err)
    }

    defer file.Close()

    file.WriteString("\xEF\xBB\xBF")
	
	w := csv.NewWriter(file)
	w.Write(column)

	for _, record := range data {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
		return err;
	}

	return nil
}