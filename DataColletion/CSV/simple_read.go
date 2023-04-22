package CSV

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// Define the expected return record type
type CSVRecord struct {
	SepalLength float64
	SepalWidth  float64
	PetalLength float64
	PetalWidth  float64
	Species     string
	ParseError  error
}

func Simple() {
	//Open iris file
	f, err := os.Open("./Data/DataCollection/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//Create Reader
	reader := csv.NewReader(f)
	//Read 5 fields in each row
	reader.FieldsPerRecord = 5

	var csvData []CSVRecord
	for {
		//read data by line
		record, err := reader.Read()
		//finish reading
		if err == io.EOF {
			break
		}

		var csvRecord CSVRecord
		for index, value := range record {
			//check species
			if index == 4 {
				if value == "" {
					log.Printf("Unexpected type in column &d\n", index)
					csvRecord.ParseError = fmt.Errorf("Empty string value")
					break
				}
				csvRecord.Species = value
				continue
			}
			//check value
			var floatV float64
			if floatV, err = strconv.ParseFloat(value, 64); err != nil {
				log.Printf("Unexpected type in column %d\n", index)
				csvRecord.ParseError = fmt.Errorf("Couldn't parse float")
				break
			}

			switch index {
			case 0:
				csvRecord.SepalLength = floatV
			case 1:
				csvRecord.SepalWidth = floatV
			case 2:
				csvRecord.PetalLength = floatV
			case 3:
				csvRecord.PetalWidth = floatV
			}
		}

		if csvRecord.ParseError == nil {
			csvData = append(csvData, csvRecord)
		}

	}
	for i := 0; i < 5; i++ {
		fmt.Println(csvData[i])
	}
}
