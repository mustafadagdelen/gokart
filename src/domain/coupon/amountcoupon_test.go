package coupon

import (
	"testing"
	"time"
)

func TestAmountCouponIsApplicableForValidParameters(t *testing.T) {
	minPurchaseAmount := float64(0)
	amount := float64(30)
	couponCode := "TEST-COUPON"
	finishDate := time.Now().AddDate(0, 0, 40)

	f := AmountCoupon{minPurchaseAmount, amount, couponCode, finishDate}

	totalCartAmount := float64(200)

	if f.IsApplicable(totalCartAmount) == false {
		t.Errorf("Coupon did not work properly")
	}
}

func TestAmountCouponIsApplicableForInValidParameters(t *testing.T) {

	minPurchaseAmount := float64(150)
	amount := float64(20)
	couponCode := "TEST-COUPON"
	finishDate := time.Now().AddDate(0, 0, 40)

	f := AmountCoupon{minPurchaseAmount, amount, couponCode, finishDate}

	totalCartAmount := float64(100)

	if f.IsApplicable(totalCartAmount) == true {
		t.Errorf("Total cart amount should greater then min purchase amount")
	}
}

func TestAmountCouponApplyDiscountForValidParameters(t *testing.T) {
	f := AmountCoupon{MinPurchaseAmount: 100, Amount: 20, CouponCode: "TEST-COUPON", FinishDate: time.Now().AddDate(0, 0, 40)}

	priceAfterDiscountApply := f.ApplyDiscount(200)

	expected := float64(180)

	if priceAfterDiscountApply != expected {
		t.Errorf("Amount coupon implementation is wrong")
	}
}

func TestAmountCouponApplyDiscountForInValidParameters(t *testing.T) {
	f := AmountCoupon{MinPurchaseAmount: 100, Amount: 20, CouponCode: "TEST-COUPON", FinishDate: time.Now().AddDate(0, 0, 40)}

	totalCartAmount := float64(100)
	priceAfterDiscountApply := f.ApplyDiscount(totalCartAmount)

	if priceAfterDiscountApply != totalCartAmount {
		t.Errorf("Amount coupon implementation is wrong")
	}
}
