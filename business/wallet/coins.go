package wallet

type Currency interface {
	GetAmount(string) string
}

type usdt struct {
}

func (e usdt) GetAmount(string) string{
	panic("implement me")
}

type btc struct {
}

func (b btc) GetAmount(string)string {
	panic("implement me")
}

type ars struct {
}

func (a ars) GetAmount(string) string {
	panic("implement me")
}

type errCoin struct {
}

func (e errCoin) GetAmount(string) string {
	panic("implement me")
}

func CurrencyFactory(currency string) Currency {
	switch currency {
	case "usdt":
		return &usdt{}
	case "btc":
		return &btc{}
	case "ars":
		return &ars{}
	default:
		return &errCoin{}
	}
}
