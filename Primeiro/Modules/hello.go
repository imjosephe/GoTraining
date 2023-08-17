package hello

import (
	quoteV3 "rsc.io/quote/v3"
)

func Hello() string {
	return quoteV3.HelloV3()
	// return quote.Hello()
}

func Proverb() string {
	return quoteV3.Concurrency()
}
