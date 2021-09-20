package wallet

import (
	"fmt"
	"math/big"
)
type BigFloat interface {
	Add(x, y *BigFloat) *BigFloat
	Cmp(y *BigFloat) int
	Format(s fmt.State, format rune)
	Text(format byte, prec int) string
}


type Coin struct {
	value *big.Float
	exponent int
}
func (c *Coin) GetAmount() string{
	return c.value.Text('f', c.exponent)
}
func (c *Coin) IsNegative() bool {
	return c.value.Signbit()
}
func (c *Coin) Add( b *Coin)  {
	c.value.Add(c.value,b.value)
}

func (c *Coin) Sub( b *Coin)  {
	c.value.Sub(c.value,b.value)
}

func newCoin(valueStr string, exponent int) (*Coin, bool ) {
	val, ok := new(big.Float).SetString(valueStr)
	if !ok{
		return nil,false
	}
	return &Coin{
		val,
		exponent,
	}, true
}
