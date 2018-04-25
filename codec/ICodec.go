package codec

type ICodec interface {
	Encode(input int) string
	Decode(input string) int
}

type CodecAlphabet string
type CodecPad rune
