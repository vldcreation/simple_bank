package util

const (
	USD = "USD"
	EUR = "EUR"
	JPY = "JPY"
	GBP = "GBP"
	AUD = "AUD"
	CAD = "CAD"
	CHF = "CHF"
	CNY = "CNY"
	SEK = "SEK"
	NZD = "NZD"
	IDR = "IDR"
)

// AllCurrencies returns all supported currencies.
func AllCurrencies() []string {
	return []string{
		USD,
		EUR,
		JPY,
		GBP,
		AUD,
		CAD,
		CHF,
		CNY,
		SEK,
		NZD,
		IDR,
	}
}

// IsSupportedCurrency checks if the given currency is supported.
func IsSupportedCurrency(currency string) bool {
	currencies := AllCurrencies()
	for _, c := range currencies {
		if c == currency {
			return true
		}
	}
	return false
}
