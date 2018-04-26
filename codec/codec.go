package codec

import (
	"bytes"
	"math"
)

type Codec struct {
	Alphabet  CodecAlphabet
	Padding   CodecPad
	shouldPad bool
	safe      bool
}

func NewCodec(base int) ICodec {
	var codec ICodec
	switch base {
	case 16:
		codec = Codec{Alphabet: Base16Alphabet}
		break
	case 32:
		codec = Codec{Alphabet: Base32Alphabet}
		break
	case 57:
		codec = Codec{Alphabet: Base57Alphabet}
		break
	case 64:
		codec = Codec{Alphabet: Base64Alphabet}
		break
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

type option func(*Codec)

func (c *Codec) Option(opts ...option) {
	for _, opt := range opts {
		opt(c)
	}
}

func PadChar(char rune) option {
	return func(c *Codec) {
		c.Padding = CodecPad(char)
	}
}

func Pad(pad bool) option {
	return func(c *Codec) {
		c.shouldPad = pad
	}
}

func URLSafe(safe bool) option {
	return func(c *Codec) {
		c.safe = safe
	}
}
