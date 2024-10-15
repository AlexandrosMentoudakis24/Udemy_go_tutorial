package prices

import (
	"fmt"

	"example.com/price_calculator/conversion"
	"example.com/price_calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager iomanager.IOManager `json:"-"`
	InputPrices []float64 `json:"input_prices"`
	TaxRate float64 `json:"tax_rate"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: iom,
		InputPrices: []float64{ 10, 20, 30 },
		TaxRate: taxRate,
	}
}

func (t *TaxIncludedPriceJob) CalculateTaxedPrices(doneChan chan bool, errorChan chan error) {
	err := t.loadData()

	if err != nil {
		errorChan <- err

		return
	}

	result := make(map[string]string, len(t.InputPrices))

	for _, price := range t.InputPrices {
		taxIncludedPrice := price * (1 + t.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	t.TaxIncludedPrices = result

	err = t.IOManager.WriteJSONToFile(t)

	if err != nil {
		errorChan <- err

		return
	}

	doneChan <- true
}

func (t *TaxIncludedPriceJob) loadData() error {
	lines, err := t.IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		return err
	}


	t.InputPrices = prices

	return nil
}







