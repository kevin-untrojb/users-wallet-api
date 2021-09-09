package wallet

type Coins interface {
	GetAmount()
}

type usdt struct {
}

func (e usdt) GetAmount() {
	panic("implement me")
}

type btc struct {
}

func (b btc) GetAmount() {
	panic("implement me")
}

type ars struct {
}

func (a ars) GetAmount() {
	panic("implement me")
}

type errCoin struct {
}

func (e errCoin) GetAmount() {
	panic("implement me")
}

func cointFactory(coin string) Coins {
	switch coin {
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
