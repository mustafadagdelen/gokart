package catalog

import (
	"gokart/domain/campaign"
	"testing"
	"time"
)

func TestCategoryNew(t *testing.T) {
	var categoryTitle string
	categoryTitle = "Book"

	var category = NewCategory(categoryTitle)

	if categoryTitle != category.title {
		t.Errorf("Error on Category.New function.")
	}
}

func TestSetParentCategory(t *testing.T) {
	noteBookCategory := NewCategory("Notebook")
	stationeryCategory := NewCategory("Stationary")

	noteBookCategory.SetParentCategory(stationeryCategory)

	if noteBookCategory.parentCategory == nil || noteBookCategory.parentCategory.title != stationeryCategory.title {
		t.Errorf("Error on Category.SetParentCategory function.")
	}
}

func TestAddAmountCampaign(t *testing.T) {
	var (
		amountCampaign = campaign.AmountCampaign{MinProductQuantity: 3, Amount: 20, CampaignCode: "TEST-AMOUNT-CAMPAIGN"}
		bookCategory   = NewCategory("Book")
	)

	result := bookCategory.AddCampaign(amountCampaign)

	if result == false {
		t.Error("Campaign could not added to category")
	}

	campaignCount := len(bookCategory.campaigns)

	expected := 1
	if campaignCount < expected {
		t.Error("Campaign count less then expected")
	}
}

func TestAddRateCampaign(t *testing.T) {
	var (
		rateCampaign = campaign.RateCampaign{MinProductQuantity: 3, Rate: 20, CampaignCode: "TEST-RATE-CAMPAIGN"}
		bookCategory = NewCategory("Book")
	)

	result := bookCategory.AddCampaign(rateCampaign)

	if result == false {
		t.Error("Campaign could not added to category")
	}

	campaignCount := len(bookCategory.campaigns)

	expected := 1
	if campaignCount < expected {
		t.Error("Campaign count less then expected")
	}
}

func TestAddMixedCampaign(t *testing.T) {
	var (
		amountCampaign = campaign.AmountCampaign{MinProductQuantity: 3, Amount: 20, CampaignCode: "TEST-AMOUNT-CAMPAIGN"}
		rateCampaign   = campaign.RateCampaign{MinProductQuantity: 3, Rate: 20, CampaignCode: "TEST-RATE-CAMPAIGN"}
		bookCategory   = NewCategory("Book")
	)

	rateCampaignAddResult := bookCategory.AddCampaign(rateCampaign)

	if rateCampaignAddResult == false {
		t.Error("Rate Campaign could not added to category")
	}

	amountCampaignAddResult := bookCategory.AddCampaign(amountCampaign)

	if amountCampaignAddResult == false {
		t.Error("Amount Campaign could not added to category")
	}

	campaignCount := len(bookCategory.campaigns)

	expected := 2
	if campaignCount < expected {
		t.Error("Campaign count less then expected")
	}
}

func TestAddSameCampaign(t *testing.T) {
	var (
		amountCampaign = campaign.AmountCampaign{MinProductQuantity: 3, Amount: 20, CampaignCode: "TEST-CAMPAIGN"}

		rateCampaign = campaign.RateCampaign{MinProductQuantity: 3, Rate: 20, CampaignCode: "TEST-CAMPAIGN"}

		bookCategory = NewCategory("Book")
	)

	rateCampaignAddResult := bookCategory.AddCampaign(rateCampaign)

	if rateCampaignAddResult == false {
		t.Error("Rate Campaign could not added to category")
	}

	amountCampaignAddResult := bookCategory.AddCampaign(amountCampaign)

	if amountCampaignAddResult == true {
		t.Error("Same campaign code add should return false")
	}

	campaignCount := len(bookCategory.campaigns)

	expected := 1
	if campaignCount < expected {
		t.Error("Campaign count less then expected")
	}
}

func TestGetParentCategories(t *testing.T) {
	var (
		wearCategory = NewCategory("Wear")
		menCategory  = NewCategory("Men")
		jeanCategory = NewCategory("Jean")
	)

	menCategory.SetParentCategory(wearCategory)
	jeanCategory.SetParentCategory(menCategory)

	categories := jeanCategory.GetCategoryTree()

	actual := len(categories)
	expected := 3

	if actual != expected {
		t.Errorf("Category Tree is wrong")
	}
}

func TestGetAvailableCampaigns(t *testing.T) {
	var (
		category       = NewCategory("Book")
		amountCampaign = campaign.AmountCampaign{MinProductQuantity: 3, Amount: 20, CampaignCode: "TEST-AMOUNT-CAMPAIGN", FinishDate: time.Now().AddDate(0, 0, 10)}

		rateCampaign = campaign.RateCampaign{MinProductQuantity: 5, Rate: 20, CampaignCode: "TEST-RATE-CAMPAIGN", FinishDate: time.Now().AddDate(0, 0, 10)}
	)

	result := category.AddCampaign(amountCampaign)

	t.Logf("Amount campaign add result : %v", result)

	result = category.AddCampaign(rateCampaign)
	t.Logf("Rate campaign add result : %v", result)

	currentProductCount := 4
	availableCampaigns := category.GetAvailableCampaigns(currentProductCount)

	expected := 1
	actual := len(availableCampaigns)

	if actual != expected {
		t.Errorf("GetAvailableCampaigns returned wrong result. Expected %v found %v", expected, actual)
	}
}
