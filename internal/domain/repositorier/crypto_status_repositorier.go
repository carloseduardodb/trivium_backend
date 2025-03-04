package repositorier

type CryptoData struct {
	Symbol string
	Price  string
	Volume string
}

type CryptoStatusRepository interface {
	StreamCryptoData(cryptos []string) <-chan CryptoData
	Get24hVolumes(cryptos []string) (map[string]string, error)
}
