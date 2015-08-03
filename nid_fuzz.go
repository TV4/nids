// +build gofuzz

package nid

func Fuzz(data []byte) int {
	s := string(data)

	if out := Case(s); out != "" {
		if len(s) == len(out) {
			return 1
		}
	}

	return -1
}
