package campaign

import (
	"time"
)

type RateCampaign struct {
	MinProductQuantity int
	Rate               int
	CampaignCode       string
	FinishDate         time.Time
}

func (rateCampaign RateCampaign) IsApplicable(productQuantity int) bool {
	return productQuantity > rateCampaign.MinProductQuantity && rateCampaign.FinishDate.After(time.Now())
}

func (rateCampaign RateCampaign) ApplyDiscount(amount float64, productQuantity int) float64 {
	if rateCampaign.IsApplicable(productQuantity) {
		return amount - float64(amount/100*float64(rateCampaign.Rate))
	}
	return amount
}

func (rateCampaign RateCampaign) GetDiscountAmount(totalAmount float64, productQuantity int) float64 {
	if rateCampaign.IsApplicable(productQuantity) {
		return float64(totalAmount * float64(rateCampaign.Rate) / float64(100))
	}
	return float64(0)
}

func (rateCampaign RateCampaign) GetCampaignCode() string {
	return rateCampaign.CampaignCode
}
