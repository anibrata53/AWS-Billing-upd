package CSVFile

import (
	costhandlers "awscostapi/costHandlers"
	datehandlers "awscostapi/dateHandlers"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func Createfile() {
	fmt.Println("Welcome to AWS cost Explorere API")

	f, err := os.Create("AWSBilling.csv")
	defer f.Close()

	if err != nil {
		log.Fatal("File not created", err)
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	if err = w.Write(datehandlers.SetHeading()); err != nil {
		log.Fatal("header setter fail", err)
	}

	for _, record := range costhandlers.GetAwsCost() {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

}
