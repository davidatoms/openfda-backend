package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

// Struct for parsing FDA response
type FDAResponse struct {
    Meta    map[string]interface{}   `json:"meta"`
    Results []map[string]interface{} `json:"results"`
}

func fetchFDAData(applicant string) (*FDAResponse, error) {
    url := fmt.Sprintf("https://api.fda.gov/device/510k.json?search=applicant:'%s'&limit=10", applicant)
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch data: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to read response body: %v", err)
    }

    var fdaResponse FDAResponse
    err = json.Unmarshal(body, &fdaResponse)
    if err != nil {
        return nil, fmt.Errorf("failed to parse JSON: %v", err)
    }

    return &fdaResponse, nil
}

func main() {
    applicant := "Blackrock Neurotech"
    data, err := fetchFDAData(applicant)
    if err != nil {
        log.Fatalf("Error fetching FDA data: %v", err)
    }

    for _, result := range data.Results {
        fmt.Printf("Device Name: %v\n", result["device_name"])
        fmt.Printf("Applicant: %v\n", result["applicant"])
        fmt.Printf("Decision Date: %v\n", result["decision_date"])
        fmt.Println("-----")
    }
}
