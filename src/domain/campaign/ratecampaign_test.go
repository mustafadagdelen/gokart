package campaign

import (
	"testing"
	"time"
)

func TestRateCampaignIsApplicableForValidParameters(t *testing.T) {
	minProductQuantity := 2
	rate := int(20)
	campaignCode := "TEST-CAMPAIGN"
	finishDate := time.Now().AddDate(0, 0, 30)
	rateCampaign := RateCampaign{MinProductQuantity: minProductQuantity, Rate: rate, CampaignCode: campaignCode, FinishDate: finishDate}

	totalCartItem := 4

	if rateCampaign.IsApplicable(totalCartItem) == false {
		t.Errorf("Campaign did not work properly")
	}
}

func TestRateCampaignIsApplicableForInValidParameters(t *testing.T) {
	minProductQuantity := 3
	rate := int(20)
	campaignCode := "TEST-CAMPAIGN"
	finishDate := time.Now().AddDate(0, 0, 30)
	rateCampaign := RateCampaign{MinProductQuantity: minProductQuantity, Rate: rate, CampaignCode: campaignCode, FinishDate: finishDate}

	totalCartItem := 2

	if rateCampaign.IsApplicable(totalCartItem) == true {
		t.Errorf("Total cart amount should greater then min purchase amount")
	}
}

//TODO : This test fails due to wrong data type of price
func TestRateCampaignApplyDiscountForValidParameters(t *testing.T) {
	rateCampaign := RateCampaign{MinProductQuantity: 3, Rate: 20, CampaignCode: "TEST-CAMPAIGN", FinishDate: time.Now().AddDate(0, 0, 30)}

	totalCartAmount := float64(200)
	totalCartItem := 7
	priceAfterDiscountApply := rateCampaign.ApplyDiscount(totalCartAmount, totalCartItem)

	expected := float64(160)

	if priceAfterDiscountApply != expected {
		t.Logf("This test fails due to wrong data type of price...!")
		t.Errorf("Rate campaign implementation is wrong for valid parameters. Expected %v. But found  %v", expected, priceAfterDiscountApply)
	}
}

func TestRateCampaignApplyDiscountForInValidParameters(t *testing.T) {
	rateCampaign := RateCampaign{MinProductQuantity: 3, Rate: 20, CampaignCode: "TEST-CAMPAIGN", FinishDate: time.Now().AddDate(0, 0, 30)}

	totalCartAmount := float64(200)
	totalCartItem := 2
	priceAfterDiscountApply := rateCampaign.ApplyDiscount(totalCartAmount, totalCartItem)

	if priceAfterDiscountApply != totalCartAmount {
		t.Errorf("Rate campaign implementation is wrong. Expected : %b . But found %b", totalCartAmount, priceAfterDiscountApply)
	}
}

func TestGetRateDiscountAmountForValidParameters(t *testing.T) {
	rateCampaign := RateCampaign{MinProductQuantity: 3, Rate: 20, CampaignCode: "TEST-CAMPAIGN", FinishDate: time.Now().AddDate(0, 0, 30)}

	actual := rateCampaign.GetDiscountAmount(100, 5)
	expected := float64(80)

	if actual == expected {
		t.Errorf("Rate campaign implementation is wrong. Expected : %b . But found %b", expected, actual)
	}
}

func TestGetRateDiscountAmountForInValidParameters(t *testing.T) {
	rateCampaign := RateCampaign{MinProductQuantity: 3, Rate: 20, CampaignCode: "TEST-CAMPAIGN", FinishDate: time.Now().AddDate(0, 0, 30)}

	actual := rateCampaign.GetDiscountAmount(100, 1)
	expected := float64(0)

	if actual != expected {
		t.Errorf("Rate campaign implementation is wrong. Expected : %b . But found %b", expected, actual)
	}
}
