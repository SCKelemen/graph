package flake

type ICodec interface {
	Encode(input int) string
	Decode(input string) int
}

type Base54Codec struct {
	EncodeMap map[int]rune
	DecodeMap map[rune]int
	CanonicalizationMap map[rune]rune
}

func (c Base54Codec) Encode(input int) string {
	runemap := map[int]rune{
		0: '0',
		1: '1',
		2: '2',
		3: '3',
		4: '4',
		5: '5',
		6: '6',
		7: '7',
		8: '8',
		9: '9',
		10: 'A',		
		11: 'C',
		12: 'D',
		13: 'E',
		14: 'F',
		15: 'G',
		16: 'H',		
		17: 'J',
		18: 'K',
		19: 'L',
		20: 'M',
		21: 'N',		
		22: 'P',
		23: 'Q',
		24: 'R',
		25: 'S',
		26: 'T',
		27: 'U',
		28: 'V',
		29: 'W',
		30: 'X',
		31: 'Y',
		32: 'Z',
		33: 'a',
		34: 'b',
		35: 'c',
		36: 'd',
		37: 'e',
		38: 'f',
		39: 'g',
		40: 'h',		
		41: 'j',
		42: 'k',		
		43: 'm',
		44: 'n',
		45: 'o',
		46: 'p',
		47: 'q',
		48: 'r',
		49: 's',
		50: 't',
		51: 'u',
		52: 'v',
		53: 'w',
		54: 'x',
		55: 'y',
		56: 'z',
	}

}

func (c )

func (c codec) CanonicalizeRune(r rune) rune {
	canonicalizationMap := map[rune]rune{
		'0': '0',
		'1': '1',
		'2': '2',
		'3': '3',
		'4': '4',
		'5': '5',
		'6': '6',
		'7': '7',
		'8': '8',
		'9': '9',
		'A': 'A',
		'B': '8',
		'C': 'C',
		'D': 'D',
		'E': 'E',
		'F': 'F',
		'G': 'G',
		'H': 'H',
		'I': '1',
		'J': 'J',
		'K': 'K',
		'L': 'L',
		'M': 'M',
		'N': 'N',
		'O': '0',
		'P': 'P',
		'Q': 'Q',
		'R': 'R',
		'S': 'S',
		'T': 'T',
		'U': 'U',
		'V': 'V',
		'W': 'W',
		'X': 'X',
		'Y': 'Y',
		'Z': 'Z',
		'a': 'a',
		'b': 'b',
		'c': 'c',
		'd': 'd',
		'e': 'e',
		'f': 'f',
		'g': 'g',
		'h': 'h',
		'i': '1',
		'j': 'j',
		'k': 'k',
		'l': '1',
		'm': 'm',
		'n': 'n',
		'o': 'o',
		'p': 'p',
		'q': 'q',
		'r': 'r',
		's': 's',
		't': 't',
		'u': 'u',
		'v': 'v',
		'w': 'w',
		'x': 'x',
		'y': 'y',
		'z': 'z',
	}

	if char, ok := canonicalizationMap[r]; ok {
		return char
	}
	return rune(0)
}


func NewBase64Codec() ICodec {

}