package campaign

import (
	"testing"
	"time"
)

func TestAmountCampaignIsApplicableForValidParameters(t *testing.T) {
	minProductQuantity := 3
	amount := float64(30)
	campaignCode := "TEST-CAMPAIGN"
	finishDate := time.Now().AddDate(0, 0, 40)
	amountCampaign := AmountCampaign{amount, minProductQuantity, campaignCode, finishDate}

	totalCartProduct := 5

	if amountCampaign.IsApplicable(totalCartProduct) == false {
		t.Errorf("Campaign did not work properly")
	}
}

func TestAmountCampaignIsApplicableForInValidParameters(t *testing.T) {

	minProductQuantity := 3
	amount := float64(20)
	campaignCode := "TEST-CAMPAIGN"
	finishDate := time.Now().AddDate(0, 0, 40)
	amountCampaign := AmountCampaign{amount, minProductQuantity, campaignCode, finishDate}

	totalCartProduct := 1

	if amountCampaign.IsApplicable(totalCartProduct) == true {
		t.Errorf("Total cart amount should greater then min purchase amount")
	}
}

func TestAmountCampaignApplyDiscountForValidParameters(t *testing.T) {
	amountCampaign := AmountCampaign{MinProductQuantity: 3, Amount: 20, CampaignCode: "TEST-CAMPAIGN", FinishDate: time.Now().AddDate(0, 0, 40)}

	priceAfterDiscountApply := amountCampaign.ApplyDiscount(200, 100)

	expected := float64(180)

	if priceAfterDiscountApply != expected {
		t.Errorf("Amount campaign implementation is wrong")
	}
}

func TestAmountCampaignApplyDiscountForInValidParameters(t *testing.T) {
	amountCampaign := AmountCampaign{MinProductQuantity: 3, Amount: 20, CampaignCode: "TEST-CAMPAIGN", FinishDate: time.Now().AddDate(0, 0, 40)}

	totalCartAmount := float64(100)
	totalCartItemCount := 2
	priceAfterDiscountApply := amountCampaign.ApplyDiscount(totalCartAmount, totalCartItemCount)

	if priceAfterDiscountApply != totalCartAmount {
		t.Errorf("Amount campaign implementation is wrong")
	}
}

func TestAmountCampaignGetDiscountAmountForValidParameters(t *testing.T) {
	amountCampaign := AmountCampaign{MinProductQuantity: 3, Amount: 20, CampaignCode: "TEST-CAMPAIGN", FinishDate: time.Now().AddDate(0, 0, 40)}

	totalCartAmount := float64(100)
	totalCartItemCount := 5
	discountAmount := amountCampaign.GetDiscountAmount(totalCartAmount, totalCartItemCount)
	expected := float64(20)

	if discountAmount != expected {
		t.Errorf("Amount campaign get discount amount implementation is wrong")
	}
}

func TestAmountCampaignGetDiscountAmountForInValidParameters(t *testing.T) {
	amountCampaign := AmountCampaign{MinProductQuantity: 3, Amount: 20, CampaignCode: "TEST-CAMPAIGN", FinishDate: time.Now().AddDate(0, 0, 40)}

	totalCartAmount := float64(100)
	totalCartItemCount := 1
	discountAmount := amountCampaign.GetDiscountAmount(totalCartAmount, totalCartItemCount)
	expected := float64(0)

	if discountAmount != expected {
		t.Errorf("Amount campaign get discount amount implementation is wrong")
	}
}
