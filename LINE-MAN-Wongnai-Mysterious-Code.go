package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func reverse(s string) string {
	rv := []rune(s)
	for i, j := 0, len(rv)-1; i < j; i, j = i+1, j-1 {
		rv[i], rv[j] = rv[j], rv[i]
	}
	return string(rv)
}
func main() {
	var whatIsIt string
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)
	fmt.Println(string(sd))
	StringReverse := strings.Split(reverse(string(sd)), ":")
	whatIsIt = strings.Join(StringReverse, " ")
	fmt.Println(whatIsIt)
}
