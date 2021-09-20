package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdditionWithIntegers(t *testing.T) {
	startValue := "5"
	additionValue := "2"
	solution := "7"
	coin, ok := newCoin(startValue, 0)
	assert.True(t, ok)
	coin.Add(additionValue)
	assert.Equal(t, solution, coin.GetAmount())
}

func TestAdditionWithLongFloats(t *testing.T) {
	startValue := "5.00100001"
	additionValue := "2.11011010"
	solution := "7.11111011"
	coin, ok := newCoin(startValue, 8)
	assert.True(t, ok)
	coin.Add(additionValue)
	assert.Equal(t, solution, coin.GetAmount())
}

func TestSubtraction(t *testing.T) {
	startValue := "7.11111011"
	subValue := "2.11011010"

	solution := "5.00100001"
	coin, ok := newCoin(startValue, 8)
	assert.True(t, ok)
	coin.Sub(subValue)
	assert.Equal(t, solution, coin.GetAmount())
}
func TestNegative(t *testing.T) {
	startValue := "7.11111011"
	subValue := "90.11011010"

	coin, ok := newCoin(startValue, 8)
	assert.True(t, ok)
	coin.Sub(subValue)
	assert.True(t, coin.IsNegative())
}
