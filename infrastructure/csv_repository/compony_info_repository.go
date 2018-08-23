package csv_repository

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/YasushiKobayashi/search-list/model"
	"github.com/pkg/errors"
)

func (r *CsvRepository) WritePageInfo(csvWriter *model.CsvWriter, prefix string) error {
	writePath := fmt.Sprintf("%s_%s", prefix, r.Path)
	file, err := os.Create(writePath)
	if err != nil {
		return errors.Wrap(err, "os.Create error")
	}
	defer file.Close()

	// writer := csv.NewWriter(transform.NewWriter(osWriter, japanese.ShiftJIS.NewEncoder()))
	writer := csv.NewWriter(file)
	return r.writeCsv(writer, csvWriter)
}

func (r *CsvRepository) writeCsv(writer *csv.Writer, csvWriter *model.CsvWriter) error {
	var writeRocords [][]string

	for i, v := range csvWriter.Rows {
		if i == 0 {
			writeRocords = append(writeRocords, csvWriter.Header)
		}

		writeRocords = append(writeRocords, v)
	}

	err := writer.WriteAll(writeRocords)
	if err != nil {
		return errors.Wrap(err, "WriteAll error")
	}

	return nil
}
