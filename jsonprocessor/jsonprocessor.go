package jsonprocessor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"moritzgoedicke.com/bewaesserung/excel"
	"moritzgoedicke.com/bewaesserung/models"
)

func Worker(fileCh <-chan string, excelFile *excel.ExcelFile, processedRows map[string]bool, mu *sync.Mutex, row *int, wg *sync.WaitGroup) {
	defer wg.Done()

	for filePath := range fileCh {
		fmt.Printf("Processing file: %s\n", filePath)

		deviceDataArray, err := readJSONFile(filePath)
		if err != nil {
			log.Printf("Error processing file %s: %v", filePath, err)
			continue
		}

		for _, deviceData := range deviceDataArray {
			if !isValidDeviceData(deviceData) {
				continue
			}

			devID := deviceData.Identifiers[0].DeviceIds.DeviceID
			devEUI := deviceData.Identifiers[0].DeviceIds.DeviceEUI
			receivedAt := deviceData.Data.ReceivedAt
			payload := deviceData.Data.UplinkMessage.DecodedPayload

			key := fmt.Sprintf("%s-%s-%s", devID, devEUI, receivedAt)

			mu.Lock()
			if processedRows[key] {
				mu.Unlock()
				continue
			}

			excelFile.WriteRow(*row, deviceData, payload)
			processedRows[key] = true
			*row++
			mu.Unlock()
		}
	}
}

func readJSONFile(filePath string) ([]models.DeviceData, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	var deviceDataArray []models.DeviceData
	if err := json.Unmarshal(data, &deviceDataArray); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %v", err)
	}

	return deviceDataArray, nil
}

func isValidDeviceData(deviceData models.DeviceData) bool {
	if len(deviceData.Identifiers) == 0 ||
		deviceData.Data.ReceivedAt == "" ||
		deviceData.Data.UplinkMessage.DecodedPayload == (models.DecodedPayload{}) {
		return false
	}
	return true
}
