package golangexamples

import (
	"github.com/ehteshamz/greetings"
)

func ConcatSlice(sliceToConcat []byte) string {
	str := string(sliceToConcat[0])
	for _, c := range sliceToConcat[1:] {
		str += "-" + string(c)
	}
	return str
}

func Encrypt(sliceToEncrypt []byte, ceaserCount int) {
	for i, c := range sliceToEncrypt {
		sliceToEncrypt[i] = (((c + byte(ceaserCount)) - 97) % 26) + 97
	}
}

func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}
