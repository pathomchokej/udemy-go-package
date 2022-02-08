package hello

import (
	"fmt"

	qq "rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func Hello() string {
	var a string = qq.Hello()
	fmt.Println("rsc.io/quote.Hello() => " + a)
	return "Hello from package."
}

func Proverb() string {
	return quoteV3.Concurrency()
}
