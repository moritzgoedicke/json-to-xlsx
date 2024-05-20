# JSON to Excel Converter

This Go program reads JSON files from a directory and extracts specific values to generate an Excel file.

## Prerequisites

Before running this project, ensure you have the following installed:

-   [Go](https://golang.org/doc/install) (version 1.16 or higher)

## Installation

1. Clone the repository or download the source code.

    ```sh
    git clone https://github.com/moritzgoedicke/json-to-xlsx.git
    cd json-to-xlsx
    ```

2. Install the necessary Go packages.

    ```sh
    go mod tidy
    ```

## Directory Structure

```
json-to-xlsx/
│
├── json_files/          # Directory containing the files to be processed
│
├── main.go
│
└── README.md
```

## Usage

1. Place your JSON files in the `json_files` directory.

2. Run the Go program.

    ```sh
    go run main.go
    ```

3. After running the program, an Excel file named `output.xlsx` will be generated in the project root directory, containing the extracted values.

## Example

Assuming you have a JSON file named `example.json` in the `json_files` directory with the following content:

```json
[
    {
        "identifiers": [
            {
                "device_ids": {
                    "device_id": "beet04-s04-1"
                }
            }
        ],
        "data": {
            "received_at": "2023-05-19T12:00:00Z",
            "uplink_message": {
                "decoded_payload": {
                    "BAT": 3.7,
                    "H1": 45.0,
                    "H2": 50.0,
                    "T1": 22.5
                }
            }
        }
    }
]
```

The generated `output.xlsx` will have the following content:

| dev_ui       | received_at          | BAT | H1   | H2   | T1   |
| ------------ | -------------------- | --- | ---- | ---- | ---- |
| beet04-s04-1 | 2023-05-19T12:00:00Z | 3.7 | 45.0 | 50.0 | 22.5 |

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

-   [excelize](https://github.com/xuri/excelize) library for creating Excel files.
-   [Go programming language](https://golang.org/)
