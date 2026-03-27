package scanner

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
)

type ScanResult struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
	Service  string `json:"service"`
	Banner   string `json:"banner,omitempty"`
}

// SaveCSV writes results to a CSV file
func SaveCSV(filename string, results []ScanResult) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Host", "Port", "Protocol", "Service", "Banner"})
	for _, r := range results {
		writer.Write([]string{r.Host, strconv.Itoa(r.Port), r.Protocol, r.Service, r.Banner})
	}
	return nil
}

// SaveJSON writes results to a JSON file
func SaveJSON(filename string, results []ScanResult) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, _ := json.MarshalIndent(results, "", "  ")
	file.Write(data)
	return nil
}