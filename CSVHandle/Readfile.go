package CSVFile

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadFile(name string) [][]string {

	f, err := os.Open(name)

	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", name, err.Error())
	}

	defer f.Close()

	r := csv.NewReader(f)

	r.Comma = ';'

	rows, err := r.ReadAll()

	if err != nil {
		log.Fatalln("Cannot read CSV data:", err.Error())
	}
	// fmt.Println(rows)
	return rows

}
