package codec

type option func(*Codec)

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
