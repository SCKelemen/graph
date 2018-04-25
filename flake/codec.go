package flake

import (
	"bytes"
	"math"
)

type ICodec interface {
	Encode(input int) string
	Decode(input string) int
}

type Codec struct {
	Alphabet CodecAlphabet
	Padding  CodecPad
}

func NewCodec(base int) ICodec {
	var codec ICodec
	switch base {
	case 57:
		codec = Codec{Alphabet: Base57Alphabet}
		break
	case 64:
		codec = Codec{Alphabet: Base64Alphabet}
	}
	return codec
}

func (c Codec) Encode(number int) string {
	if number == 0 {
		return string(c.Alphabet[0])
	}

	chars := make([]byte, 0)

	length := len(c.Alphabet)

	for number > 0 {
		result := number / length
		remainder := number % length
		chars = append(chars, c.Alphabet[remainder])
		number = result
	}

	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars)
}

func (c Codec) Decode(token string) int {
	number := 0
	idx := 0.0
	chars := []byte(c.Alphabet)

	charsLength := float64(len(chars))
	tokenLength := float64(len(token))

	for _, c := range []byte(token) {
		power := tokenLength - (idx + 1)
		index := bytes.IndexByte(chars, c)
		number += index * int(math.Pow(charsLength, power))
		idx++
	}

	return number
}

type CodecAlphabet string
type CodecPad rune

const (
	Base57Alphabet            CodecAlphabet = "0123456789ACDEFGHJKLMNPQRSTUVWXYZabcdefghjkmnopqrstuvwxyz"
	Base16Alphabet            CodecAlphabet = "0123456789ABCDEF"
	Base32Alphabet            CodecAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	Base32ExtendedHexAlphabet CodecAlphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUV"
	// Base64Alphabet implements RFC4648  https://tools.ietf.org/html/rfc4648
	Base64Alphabet    CodecAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	Base64UrlAlphabet CodecAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

	Base57Pad            CodecPad = '='
	Base32Pad            CodecPad = '='
	Base32ExtendedHexPad CodecPad = '='
	Base64Pad            CodecPad = '='
	Base64UrlPad         CodecPad = '='
)
