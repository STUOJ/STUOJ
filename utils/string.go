package utils

import (
	"encoding/json"
	"html"
	"regexp"
	"strings"

	randv2 "math/rand/v2"
)

func Senitize(s string) string {
	s = strings.TrimSpace(s)
	s = html.EscapeString(s)
	return s
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func GetRandKey(n int) string {
	key := make([]rune, n)
	for i := range key {
		key[i] = letters[randv2.IntN(len(letters))]
	}
	return string(key)
}

func ToSnakeCase(s string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
