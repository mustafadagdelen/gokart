package catalog

import (
	"gokart/domain/campaign"
)

type Category struct {
	Title          string
	parentCategory *Category
	campaigns      map[string]campaign.Campaign
}

func NewCategory(title string) *Category {
	category := new(Category)
	category.Title = title
	category.campaigns = make(map[string]campaign.Campaign)
	return category
}

func (category *Category) SetParentCategory(parentCategory *Category) {
	category.parentCategory = parentCategory
}

func (category *Category) AddCampaign(c campaign.Campaign) bool {
	campaignCode := c.GetCampaignCode()

	if _, exist := category.campaigns[campaignCode]; exist {
		return false
	}

	category.campaigns[campaignCode] = c

	return true
}

func (category *Category) GetCategoryTree() []*Category {
	categories := []*Category{}

	var currentCategory *Category
	currentCategory = category

	for currentCategory != nil {
		categories = append(categories, currentCategory)
		currentCategory = currentCategory.parentCategory
	}

	return categories
}

func (category *Category) GetAvailableCampaigns(productCount int) []campaign.Campaign {
	var campaigns []campaign.Campaign

	if len(category.campaigns) <= 0 {
		return campaigns
	}
	for _, val := range category.campaigns {
		if val.IsApplicable(productCount) {
			campaigns = append(campaigns, val)
		}
	}

	return campaigns
}
