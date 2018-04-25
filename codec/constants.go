package codec

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
