package csv_repository

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/YasushiKobayashi/search-list/model"
	"github.com/pkg/errors"
)

func (r *CsvRepository) WriteSearchList(csvWriter *model.CsvWriter) error {
	writePath := fmt.Sprintf("search_list-%s", r.Path)
	osWriter, err := os.Create(writePath)
	if err != nil {
		return errors.Wrap(err, "os.Create error")
	}
	defer osWriter.Close()

	// writer := csv.NewWriter(transform.NewWriter(osWriter, japanese.ShiftJIS.NewEncoder()))
	writer := csv.NewWriter(osWriter)
	return r.writeCsvKeyWord(writer, csvWriter)
}

func (r *CsvRepository) writeCsvKeyWord(writer *csv.Writer, csvWriter *model.CsvWriter) error {
	var header []string = []string{"listing1", "listing2", "listing3", "listing4", "search1", "search2", "search3", "search4"}
	var writeRocords [][]string
	var contentVal []string

	for i, v := range csvWriter.Rows {
		if i == 0 {
			var headerRecord []string
			contentVal = append(csvWriter.Header, header...)
			headerRecord = append(headerRecord, contentVal...)
			writeRocords = append(writeRocords, headerRecord)
		}

		var newRecord []string
		scrapeVal := csvWriter.SearchLists[i].GetCsvVal()
		contentVal = append(v, scrapeVal...)
		newRecord = append(newRecord, contentVal...)
		writeRocords = append(writeRocords, newRecord)
	}

	err := writer.WriteAll(writeRocords)
	if err != nil {
		return errors.Wrap(err, "WriteAll error")
	}

	return nil
}
