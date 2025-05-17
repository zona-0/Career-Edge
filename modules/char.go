package modules

var i int
func toLowerChar(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}

func ToLower(s string) string {
	var b []byte

	b = []byte(s)
	for i = 0; i < len(b); i += 1 {
		b[i] = toLowerChar(b[i])
	}
	return string(b)
}

func EqualsIgnoreCase(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	for i = 0; i < len(a); i += 1 {
		if toLowerChar(a[i]) != toLowerChar(b[i]) {
			return false
		}
	}
	return true
}
