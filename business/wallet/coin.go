package wallet

import (
	"math/big"
)

type Coin struct {
	value    *big.Float
	exponent int
}

func (c *Coin) GetAmount() string {
	return c.value.Text('f', c.exponent)
}
func (c *Coin) IsNegative() bool {
	return c.value.Signbit()
}

func (c *Coin) Add(addValue string) bool {
	b, ok := new(big.Float).SetString(addValue)
	if !ok {
		return false
	}
	c.value.Add(c.value, b)
	return true
}

func (c *Coin) Sub(subsValue string) bool {
	b, ok := new(big.Float).SetString(subsValue)
	if !ok {
		return false
	}
	c.value.Sub(c.value, b)
	return true
}

func newCoin(valueStr string, exponent int) (*Coin, bool) {
	val, ok := new(big.Float).SetString(valueStr)
	if !ok {
		return nil, false
	}
	return &Coin{
		val,
		exponent,
	}, true
}
