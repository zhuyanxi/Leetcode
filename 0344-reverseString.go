package main

func reverseString(s string) string {
	var temp []byte
	for i := len(s) - 1; i >= 0; i-- {
		temp = append(temp, s[i])
	}
	return string(temp[:])
}
