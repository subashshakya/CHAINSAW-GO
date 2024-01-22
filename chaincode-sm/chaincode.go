package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type ECGReport struct {
	Id string `json:"id"`
	PatientName string `json:"patientName"`
	Sonographer string `json:"sonographer"`
	AorticRoot string `json:"aorticRoot"`
	LeftVentrical string `json:"leftVentrical"`
	RightVentrical string `json:"rightVentrical"`
	Diastole string `json:"diastole"`
	Systole string `json:"systole"`
	LVPW string `json:"lvpw"`
	LVEF string `json:"lvef"`
	LeftAtrium string `json:"leftAtrium"`
	IVS string `json:"ivs"`
	Owner string `json:"owner"`
	Age string `json:"age"`
	Sex string `json:"sex"`
	Cp string `json:"cp"`
	TrestBPS string `json:"tRestBPS"`
	Cholestrol string `json:"chol"`
	FBS string `json:"fbs"`
	RestECG string `json:"restecg"`
	Thalach string `json:"thalach"`
	Exang string `json:"exang"`
	OldPeak string `json:"oldPeak"`
	Slope string `json:"slope"`
	CA string `json:"ca"`
	Thal string `json:"thal"`
}


// putting one dummy data into the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContext) error {
	assets := []ECGReport{
		{
			Id: "12345",
			PatientName:    "John Doe",
			Sonographer:    "Dr. Smith",
			AorticRoot:     "Normal",
			LeftVentrical:  "Normal",
			RightVentrical: "Normal",
			Diastole:       "Normal",
			Systole:        "Normal",
			LVPW:           "Normal",
			LVEF:           "Normal",
			LeftAtrium:     "Normal",
			IVS:            "Normal",
			Owner:          "Patient's Owner",
			Age:            "30",
			Sex:            "Male",
			Cp:             "Chest Pain",
			TrestBPS:       "120",
			Cholestrol:     "180",
			FBS:            "100",
			RestECG:        "Normal",
			Thalach:        "150",
			Exang:          "No",
			OldPeak:        "2.5",
			Slope:          "1",
			CA:             "0",
			Thal:           "Normal",
		},
	}

	for _, asset := range assets {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}
		err = ctx.GetStub().PutState(asset.Id, assetJSON)
		if err != nil {
			return fmt.Errorf("Failed to inset data to world state due to: --> %v", err)
		}
	}

	return nil
}

func (s *SmartContract) CreateReport(
		ctx contractapi.TransactionContextInterface,
		Id string,
		PatientName string,
		Sonographer string,
		AorticRoot string,
		LeftVentrical string,
		RightVentrical string,
		Diastole string,
		Systole string,
		LVPW string,
		LVEF string,
		LeftAtrium string,
		IVS string,
		Owner string,
		Age string,
		Sex string,
		Cp string,
		TrestBPS string,
		Cholestrol string,
		FBS string,
		RestECG string,
		Thalach string,
		Exang string,
		OldPeak string,
		Slope string,
		CA string,
		Thal string,
	) error {
		exists, err := s.AssetExists(ctx, Id)
		if err != nil {
			return err
		}

		if !exists {
			return fmt.Errorf("The report %s already exists", Id)
		}

		report := ECGReport{
			Id: Id,
			PatientName: PatientName,
			Sonographer: Sonographer,
			AorticRoot: AorticRoot,
			LeftVentrical: LeftVentrical,
			RightVentrical: RightVentrical,
			Diastole: Diastole,
			Systole: Systole,
			LVPW: LVPW,
			LVEF: LVEF,
			LeftAtrium: LeftAtrium,
			IVS: IVS,
			Owner: Owner,
			Age: Age,
			Sex: Sex,
			Cp: Cp,
			TrestBPS: TrestBPS,
			Cholestrol: Cholestrol,
			FBS: FBS,
			RestECG: RestECG,
			Thalach: Thalach,
			Exang: Exang,
			OldPeak: OldPeak,
			Slope: Slope,
			CA: CA,
			Thal: Thal,
		}
		reportJSON, err := json.Marshal(report)
		if err != nil {
			return err
		}

	return ctx.GetStub().PutState(Id, reportJSON)
}

func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, id string) (*ECGReport, error) {
	reportJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("Failed to read from the world state. Due to : %v", err)
	}
	if reportJSON == nil {
		return nil, fmt.Errorf("The report of id %s does not exist", id)
	}

	var report ECGReport
	err = json.Unmarshal(reportJSON, &report)
	if err != nil {
		return nil, err
	}

	return &report, nil
}

func (s *SmartContract) UpdateReport(
	ctx contractapi.TransactionContextInterface,
	Id string,
		PatientName string,
		Sonographer string,
		AorticRoot string,
		LeftVentrical string,
		RightVentrical string,
		Diastole string,
		Systole string,
		LVPW string,
		LVEF string,
		LeftAtrium string,
		IVS string,
		Owner string,
		Age string,
		Sex string,
		Cp string,
		TrestBPS string,
		Cholestrol string,
		FBS string,
		RestECG string,
		Thalach string,
		Exang string,
		OldPeak string,
		Slope string,
		CA string,
		Thal string,
) error {
	exists, err := ctx.GetStub().GetState(Id)
	if err != nil {
		return fmt.Errorf("The report could not be updated. Due to: %s", err)
	}
	if exists == nil {
		return fmt.Errorf("The report of id: %s does not exist!!", exists);
	}

	newReport := ECGReport{
		Id: Id,
		PatientName: PatientName,
		Sonographer: Sonographer,
		AorticRoot: AorticRoot,
		LeftVentrical: LeftVentrical,
		RightVentrical: RightVentrical,
		Diastole: Diastole,
		Systole: Systole,
		LVPW: LVPW,
		LVEF: LVEF,
		LeftAtrium: LeftAtrium,
		IVS: IVS,
		Owner: Owner,
		Age: Age,
		Sex: Sex,
		Cp: Cp,
		TrestBPS: TrestBPS,
		Cholestrol: Cholestrol,
		FBS: FBS,
		RestECG: RestECG,
		Thalach: Thalach,
		Exang: Exang,
		OldPeak: OldPeak,
		Slope: Slope,
		CA: CA,
		Thal: Thal,
	}
	reportJSON, err := json.Marshal(newReport)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(Id, reportJSON)
}

func (s *SmartContract) DeleteReport(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("The report of id: %s does not exist", id)
	}
	return ctx.GetStub().DelState(id)
}

func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool ,error) {
	reportJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("Failed to read from world state: %v", err)
	}
	return reportJSON != nil, nil
}

func (s *SmartContract) TransferAsset(ctx contractapi.TransactionContextInterface, id string, newOwner string) error {
	report, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}
	report.Owner = newOwner
	reportJSON, err := json.Marshal(report)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(id, reportJSON)
}

func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*ECGReport, error) {
	resultIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultIterator.Close()
	var reports []*ECGReport
	for resultIterator.HasNext() {
		queryResponse, err := resultIterator.Next()
		if err != nil {
			return nil, err
		}

		var report ECGReport
		err = json.Unmarshal(queryResponse.Value, &report)
		if err != nil {
			return nil, err
		}
		reports = append(reports, &report)
	}
	return reports, nil
}

func main() {
  reportChaincode, err := contractapi.NewChaincode(&SmartContract{})
  if err != nil {
    log.Panicf("Error creating asset-transfer-basic chaincode: %v", err)
  }

  if err := reportChaincode.Start(); err != nil {
    log.Panicf("Error starting asset-transfer-basic chaincode: %v", err)
  }
}