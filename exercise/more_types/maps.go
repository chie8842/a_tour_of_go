package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	slice := strings.Split(s, " ")
	for _, str := range slice {
    	m[str] += 1 
  	}
	return m
}

func main() {
	wc.Test(WordCount)
}
