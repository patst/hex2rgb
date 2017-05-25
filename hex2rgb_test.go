package hex2rgb

import "testing"

func TestConverter(t *testing.T) {
    
	r, g, b := Convert("#A972B6")
	if r != 169 || g != 114 || b != 182 {
		t.Error("Expected 169 114 182, got ", r, g, b)
	}
	
	r, g, b = Convert("#A97")
	if r != 170 || g != 153 || b != 119 {
		t.Error("Expected 170 153 119, got ", r, g, b)
	}
}
