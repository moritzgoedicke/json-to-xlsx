package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "path/filepath"

    "github.com/xuri/excelize/v2"
)

type DeviceData struct {
    Identifiers []struct {
        DeviceIds struct {
            DeviceID string `json:"device_id"`
			DeviceEUI string `json:"dev_eui"`
        } `json:"device_ids"`
    } `json:"identifiers"`
    Data struct {
        ReceivedAt    string `json:"received_at"`
        UplinkMessage struct {
            DecodedPayload struct {
                BAT float64 `json:"BAT"`
                H1  float64 `json:"H1"`
                H2  float64 `json:"H2"`
                T1  float64 `json:"T1"`
            } `json:"decoded_payload"`
        } `json:"uplink_message"`
    } `json:"data"`
}

func main() {
    // Directory containing JSON files
    dir := "./json_files"
    
    // Create a new Excel file
    f := excelize.NewFile()
    
    // Create a new sheet
    index, err := f.NewSheet("Sheet1")
    if err != nil {
        log.Fatalf("Failed to create new sheet: %v", err)
    }
    
    // Set headers
    headers := []string{"device_id","dev_eui", "received_at", "BAT", "H1", "H2", "T1"}
    for i, header := range headers {
        cell := fmt.Sprintf("%s1", string('A'+i))
        f.SetCellValue("Sheet1", cell, header)
    }

    // Iterate over JSON files in the directory
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatalf("Failed to read directory: %v", err)
    }

    row := 2
    for _, file := range files {
        if filepath.Ext(file.Name()) == ".json" {
            // Read the JSON file
            path := filepath.Join(dir, file.Name())
            jsonData, err := ioutil.ReadFile(path)
            if err != nil {
                log.Printf("Failed to read file %s: %v", path, err)
                continue
            }

            // Parse JSON data as an array
            var deviceDataArray []DeviceData
            if err := json.Unmarshal(jsonData, &deviceDataArray); err != nil {
                log.Printf("Failed to unmarshal JSON data: %v", err)
                continue
            }

            // Process each item in the array
            for _, deviceData := range deviceDataArray {
                // Ensure that the necessary fields are present
                if len(deviceData.Identifiers) == 0 || 
                   deviceData.Data.ReceivedAt == "" ||
                   deviceData.Data.UplinkMessage.DecodedPayload == (struct {
                       BAT float64 `json:"BAT"`
                       H1  float64 `json:"H1"`
                       H2  float64 `json:"H2"`
                       T1  float64 `json:"T1"`
                   }{}) {
                    continue
                }
                
                devID := deviceData.Identifiers[0].DeviceIds.DeviceID
				devEUI := deviceData.Identifiers[0].DeviceIds.DeviceEUI
                receivedAt := deviceData.Data.ReceivedAt
                decodedPayload := deviceData.Data.UplinkMessage.DecodedPayload

                // Write values to Excel
                f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), devID)
                f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), devEUI)
                f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), receivedAt)
                f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), decodedPayload.BAT)
                f.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), decodedPayload.H1)
                f.SetCellValue("Sheet1", fmt.Sprintf("F%d", row), decodedPayload.H2)
                f.SetCellValue("Sheet1", fmt.Sprintf("G%d", row), decodedPayload.T1)

                row++
            }
        }
    }

    // Set active sheet
    f.SetActiveSheet(index)

    // Save the Excel file
    if err := f.SaveAs("output.xlsx"); err != nil {
        log.Fatalf("Failed to save Excel file: %v", err)
    }

    fmt.Println("Data successfully written to output.xlsx")
}
