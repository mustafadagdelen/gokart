package coupon

import (
	"testing"
	"time"
)

func TestRateCouponIsApplicableForValidParameters(t *testing.T) {
	minPurchaseAmount := float64(0)
	rate := int(30)
	couponCode := "TEST-COUPON"
	finishDate := time.Now().AddDate(0, 0, 10)
	f := RateCoupon{minPurchaseAmount, rate, couponCode, finishDate}

	totalCartAmount := float64(200)

	if f.IsApplicable(totalCartAmount) == false {
		t.Errorf("Coupon did not work properly")
	}
}

func TestRateCouponIsApplicableForInValidMinPurchaseAmout(t *testing.T) {
	minPurchaseAmount := float64(150)
	rate := int(20)
	couponCode := "TEST-COUPON"
	finishDate := time.Now().AddDate(0, 0, 40)
	f := RateCoupon{minPurchaseAmount, rate, couponCode, finishDate}

	totalCartAmount := float64(100)

	if f.IsApplicable(totalCartAmount) == true {
		t.Errorf("Total cart amount should greater then min purchase amount")
	}
}

func TestRateCouponIsApplicableForInValidDateParameters(t *testing.T) {
	minPurchaseAmount := float64(150)
	rate := int(20)
	couponCode := "TEST-COUPON"
	finishDate := time.Now().AddDate(0, 0, -10)
	f := RateCoupon{minPurchaseAmount, rate, couponCode, finishDate}

	totalCartAmount := float64(100)

	if f.IsApplicable(totalCartAmount) == true {
		t.Errorf("Total cart amount should greater then min purchase amount")
	}
}

func TestRateCouponApplyDiscountForValidParameters(t *testing.T) {
	f := RateCoupon{MinPurchaseAmount: 100, Rate: 20, CouponCode: "TEST-COUPON", FinishDate: time.Now().AddDate(0, 0, 10)}

	priceAfterDiscountApply := f.ApplyDiscount(200)

	expected := float64(160)

	if priceAfterDiscountApply != expected {
		t.Errorf("Rate coupon implementation is wrong for valid parameters. Expected %v. But found  %v", expected, priceAfterDiscountApply)
	}
}

func TestRateCouponApplyDiscountForInValidParameters(t *testing.T) {
	f := RateCoupon{MinPurchaseAmount: 100, Rate: 20, CouponCode: "TEST-COUPON", FinishDate: time.Now().AddDate(0, 0, 10)}

	totalCartAmount := float64(100)
	priceAfterDiscountApply := f.ApplyDiscount(totalCartAmount)

	if priceAfterDiscountApply != totalCartAmount {
		t.Errorf("Rate coupon implementation is wrong. Expected : %b . But found %b", totalCartAmount, priceAfterDiscountApply)
	}
}