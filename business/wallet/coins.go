package wallet

type Currency interface {
	GetAmount(string) string
}

type usdt struct {
	exp int64
}
func (e usdt) GetAmount(amount string) string {
	return amount
}

type btc struct {
	amount string
	exp int64
}

func (b btc) GetAmount(amount string) string {
	return amount
}

type ars struct {
	amount string
	exp int64
}

func (a ars) GetAmount(amount string) string {
	return amount
}

type errCoin struct {
}
func (e errCoin) GetAmount(string) string {
	return "COIN IS NOT VALID"
}

func CurrencyFactory(currency string, exponent int64) Currency {
	switch currency {
	case "usdt":
		return &usdt{exp: exponent}
	case "btc":
		return &btc{exp: exponent}
	case "ars":
		return &ars{exp: exponent}
	default:
		return &errCoin{}
	}
}
