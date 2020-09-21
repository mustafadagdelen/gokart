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

	a := AmountCoupon{minPurchaseAmount, amount, couponCode, finishDate}

	totalCartAmount := float64(200)

	if a.IsApplicable(totalCartAmount) == false {
		t.Errorf("Coupon did not work properly")
	}
}

func TestAmountCouponIsApplicableForInValidParameters(t *testing.T) {

	minPurchaseAmount := float64(150)
	amount := float64(20)
	couponCode := "TEST-COUPON"
	finishDate := time.Now().AddDate(0, 0, 40)

	a := AmountCoupon{minPurchaseAmount, amount, couponCode, finishDate}

	totalCartAmount := float64(100)

	if a.IsApplicable(totalCartAmount) == true {
		t.Errorf("Total cart amount should greater then min purchase amount")
	}
}

func TestAmountCouponApplyDiscountForValidParameters(t *testing.T) {
	a := AmountCoupon{MinPurchaseAmount: 100, Amount: 20, CouponCode: "TEST-COUPON", FinishDate: time.Now().AddDate(0, 0, 40)}

	priceAfterDiscountApply := a.ApplyDiscount(200)

	expected := float64(180)

	if priceAfterDiscountApply != expected {
		t.Errorf("Amount coupon implementation is wrong")
	}
}

func TestAmountCouponApplyDiscountForInValidParameters(t *testing.T) {
	a := AmountCoupon{MinPurchaseAmount: 100, Amount: 20, CouponCode: "TEST-COUPON", FinishDate: time.Now().AddDate(0, 0, 40)}

	totalCartAmount := float64(80)
	priceAfterDiscountApply := a.ApplyDiscount(totalCartAmount)

	if priceAfterDiscountApply != totalCartAmount {
		t.Errorf("Amount coupon implementation is wrong")
	}
}

func TestAmountCouponGetCouponCode(t *testing.T) {
	minPurchaseAmount := float64(0)
	amount := float64(30)
	couponCode := "TEST-COUPON"
	finishDate := time.Now().AddDate(0, 0, 40)

	a := AmountCoupon{minPurchaseAmount, amount, couponCode, finishDate}

	r := a.GetCouponCode()

	if r != couponCode {
		t.Error("Amount coupon GetCouponCode returned wrong result")
	}
}
