package domain

import "time"

type Item struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Sum      float32 `json:"sum"`
	Quantity float32 `json:"quantity"`
}

type Items struct {
	Items []Item `json:"items"`
}

type FullReceipt []struct {
	ID        string    `json:"_id"`
	CreatedAt time.Time `json:"createdAt"`
	Ticket    struct {
		Document struct {
			Receipt struct {
				CashTotalSum            int    `json:"cashTotalSum"`
				Code                    int    `json:"code"`
				CreditSum               int    `json:"creditSum"`
				DateTime                string `json:"dateTime"`
				EcashTotalSum           int    `json:"ecashTotalSum"`
				FiscalDocumentFormatVer int    `json:"fiscalDocumentFormatVer"`
				FiscalDocumentNumber    int    `json:"fiscalDocumentNumber"`
				FiscalDriveNumber       string `json:"fiscalDriveNumber"`
				FiscalSign              int64  `json:"fiscalSign"`
				FnsURL                  string `json:"fnsUrl"`
				Items
				KktRegID            string `json:"kktRegId"`
				Nds10               int    `json:"nds10"`
				Nds18               int    `json:"nds18"`
				OperationType       int    `json:"operationType"`
				Operator            string `json:"operator"`
				PrepaidSum          int    `json:"prepaidSum"`
				ProvisionSum        int    `json:"provisionSum"`
				RequestNumber       int    `json:"requestNumber"`
				RetailPlace         string `json:"retailPlace"`
				RetailPlaceAddress  string `json:"retailPlaceAddress"`
				ShiftNumber         int    `json:"shiftNumber"`
				TaxationType        int    `json:"taxationType"`
				AppliedTaxationType int    `json:"appliedTaxationType"`
				TotalSum            int    `json:"totalSum"`
				User                string `json:"user"`
				UserInn             string `json:"userInn"`
			} `json:"receipt"`
		} `json:"document"`
	} `json:"ticket"`
}
