package redacted

import (
	"encoding/csv"
	"log"
	"os"
	"reflect"
)

func ReadCSV[T any](file string, header bool) []T {
	csvfile, err := os.Open(file)
	defer csvfile.Close()
	PanicOnError(err, "Couldn't open the csv file")
	r := csv.NewReader(csvfile)

	var output []T
	records, err := r.ReadAll()

	PanicOnError(err, "unable to read records from file")
	for i, record := range records {
		var TRecord T

		// skip header
		if (i > 0) && (header == true) {
			fields := reflect.TypeOf(TRecord).NumField()
			for j := 0; j < fields; j++ {
				field := reflect.TypeOf(TRecord).Field(j)
				name := field.Name

				// Assign the value from the CSV file to the field
				v := reflect.ValueOf(&TRecord).Elem().FieldByName(name)
				v.Set(reflect.ValueOf(record[j]))
			}
			output = append(output, TRecord)
		}

	}
	return output
}

func PanicOnError(err error, message string) {
	if err != nil {
		log.Panicf("%s: %s", message, err)
	}
}
