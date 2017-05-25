package hex2rgb

import (
	h "encoding/hex"
	f "fmt"
	s "strings"
)

func Convert(hex string) (int, int, int) {

	if s.HasPrefix(hex, "#") {
		hex = s.Replace(hex, "#", "", 1)
	}

	if len(hex) == 3 {
		hex = f.Sprintf("%c%c%c%c%c%c", hex[0], hex[0], hex[1], hex[1], hex[2], hex[2])
	}

	d, _ := h.DecodeString(hex)

	return int(d[0]), int(d[1]), int(d[2])
}
