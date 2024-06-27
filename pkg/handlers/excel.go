package handlers

import (
	"log"

	"github.com/xuri/excelize/v2"
)

const truefilePath = "C:\\Users\\vicev\\Documents\\projects\\enubeTeste\\pkg\\data\\Reconfile fornecedores.xlsx"
const trueSheetName = "Planilha1"

func readTrueColumn(column string, pageSize, pageNumber int) ([]string, error) {
	f, err := excelize.OpenFile(truefilePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println("Erro ao fechar o arquivo:", err)
		}
	}()

	rows, err := f.GetRows(trueSheetName)
	if err != nil {
		return nil, err
	}

	var values []string
	var colIndex int
	switch column {
	case "PartnerId":
		colIndex = 1
	case "PartnerName":
		colIndex = 2
	case "CustomerId":
		colIndex = 3
	case "CustomerName":
		colIndex = 4
	case "CustomerDomainName":
		colIndex = 5
	case "CustomerCountry":
		colIndex = 6
	case "MpnId":
		colIndex = 7
	case "Tier2MpnId":
		colIndex = 8
	case "InvoiceNumber":
		colIndex = 9
	case "ProductId":
		colIndex = 10
	case "SkuId":
		colIndex = 11
	case "AvailabilityId":
		colIndex = 12
	case "SkuName":
		colIndex = 13
	case "ProductName":
		colIndex = 14
	case "PublisherName":
		colIndex = 15
	case "PublisherId":
		colIndex = 16
	case "SubscriptionDescription":
		colIndex = 17
	case "SubscriptionId":
		colIndex = 18
	case "ChargeStartDate":
		colIndex = 19
	case "ChargeEndDate":
		colIndex = 20
	case "UsageDate":
		colIndex = 21
	case "MeterType":
		colIndex = 22
	case "MeterCategory":
		colIndex = 23
	case "MeterId":
		colIndex = 24
	case "MeterSubCategory":
		colIndex = 25
	case "MeterName":
		colIndex = 26
	case "MeterRegion":
		colIndex = 27
	case "Unit":
		colIndex = 28
	case "ResourceLocation":
		colIndex = 29
	case "ConsumedService":
		colIndex = 30
	case "ResourceGroup":
		colIndex = 31
	case "ResourceURI":
		colIndex = 32
	case "ChargeType":
		colIndex = 33
	case "UnitPrice":
		colIndex = 34
	case "Quantity":
		colIndex = 35
	case "UnitType":
		colIndex = 36
	case "BillingPreTaxTotal":
		colIndex = 37
	case "BillingCurrency":
		colIndex = 38
	case "PricingPreTaxTotal":
		colIndex = 39
	case "PricingCurrency":
		colIndex = 40
	case "ServiceInfo1":
		colIndex = 41
	case "ServiceInfo2":
		colIndex = 42
	case "Tags":
		colIndex = 43
	case "AdditionalInfo":
		colIndex = 44
	case "EffectiveUnitPrice":
		colIndex = 45
	case "PCToBCExchangeRate":
		colIndex = 46
	case "PCToBCExchangeRateDate":
		colIndex = 47
	case "EntitlementId":
		colIndex = 48
	case "EntitlementDescription":
		colIndex = 49
	case "PartnerEarnedCreditPercentage":
		colIndex = 50
	case "CreditPercentage":
		colIndex = 51
	case "CreditType":
		colIndex = 52
	case "BenefitOrderId":
		colIndex = 53
	case "BenefitId":
		colIndex = 54
	case "BenefitType":
		colIndex = 55
	default:
		return nil, nil
	}

	start := (pageNumber - 1) * pageSize
	end := start + pageSize

	if start >= len(rows) {
		return []string{}, nil
	}

	if end > len(rows) {
		end = len(rows)
	}

	for i := start; i < end; i++ {
		row := rows[i]
		if len(row) >= colIndex {
			values = append(values, row[colIndex-1])
		}
	}
	return values, nil
}
