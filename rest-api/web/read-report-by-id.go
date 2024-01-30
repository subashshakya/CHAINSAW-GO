package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

func (setup OrgSetup) ReadReportHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id := params.Get("id")

	report, err := GetAllReports(contract, id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to find the report: %v", err), http.StatusInternalServerError)
	}
	reportJSON := formatJSON(report)

	w.Write([]byte(reportJSON))
}


func GetAllReports(contract *client.Contract, id string) ([]byte, error) {
	fmt.Println("FETCHING REPORT FROM THE NETWORK")
	reports, err := contract.EvaluateTransaction("ReadAsset", id)

	fmt.Println(reports)
	if err != nil {
		fmt.Printf("THERE WAS AN ERROR WHILE FETCHING REPORT.\n Reason: %s", err)
		return nil, err
	}
	return reports, nil
}

func formatJSON(data []byte) string {
	var result bytes.Buffer
	if err := json.Indent(&result, data, "", " "); err != nil {
		panic(fmt.Errorf("FAILED TO PARSE JSON: %w", err))
	}
	return result.String()
}