package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"moritzgoedicke.com/bewaesserung/excel"
	"moritzgoedicke.com/bewaesserung/jsonprocessor"
)

func main() {
	dir := "./json_files"
	outputFile := "output.xlsx"

	startTime := time.Now()

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	excelFile := excel.NewExcelFile()
	defer excelFile.Save(outputFile)

	headers := []string{"device_id", "dev_eui", "received_at", "BAT", "H1", "H2", "T1"}
	excelFile.SetHeaders(headers)

	processedRows := make(map[string]bool)
	var mu sync.Mutex
	row := 2

	fileCh := make(chan string, len(files))
	wg := &sync.WaitGroup{}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			fileCh <- filepath.Join(dir, file.Name())
		}
	}
	close(fileCh)

	numWorkers := 10
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go jsonprocessor.Worker(fileCh, excelFile, processedRows, &mu, &row, wg)
	}

	wg.Wait()

	endTime := time.Now()
	duration := endTime.Sub(startTime)

	fmt.Println("Data successfully written to", outputFile)
	fmt.Printf("Process took %s\n", duration)
}
