package part1

import (
	"testing"
)

func TestMoney(t *testing.T) {
	r := Calc(5, 2)
	if r != 10 {
		t.Errorf("Calc(5, 2) = %v, want 10", r)
	}
}

func Calc(money int, amount int) int {
	return money * amount
}
