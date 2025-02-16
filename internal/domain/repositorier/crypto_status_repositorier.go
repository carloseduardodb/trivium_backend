package repositorier

type CryptoData struct {
	Symbol string
	Price  string
	Volume string
}

type CryptoStatusRepository interface {
	StreamCryptoData(cryptos []string) <-chan CryptoData
}
