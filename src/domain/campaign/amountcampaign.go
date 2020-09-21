package campaign

import (
	"time"
)

type AmountCampaign struct {
	Amount             float64
	MinProductQuantity int
	CampaignCode       string
	FinishDate         time.Time
}

func (amountCampaign AmountCampaign) IsApplicable(productQuantity int) bool {
	return productQuantity > amountCampaign.MinProductQuantity && amountCampaign.FinishDate.After(time.Now())
}

func (amountCampaign AmountCampaign) ApplyDiscount(amount float64, productQuantity int) float64 {
	if amountCampaign.IsApplicable(productQuantity) {
		return amount - amountCampaign.Amount
	}

	return amount
}

func (amountCampaign AmountCampaign) GetDiscountAmount(totalAmount float64, productQuantity int) float64 {
	if amountCampaign.IsApplicable(productQuantity) {
		return totalAmount
	}

	return 0
}

func (amountCampaign AmountCampaign) GetCampaignCode() string {
	return amountCampaign.CampaignCode
}
