package functionts

import (
	"crypto/sha256"
	"fmt"
	"reflect"
)

func TypeDetection(object any) string {
	typeObj := reflect.TypeOf(object)
	fmt.Printf("object - (%v) is %v\n", object, typeObj)
	return typeObj.String()
}

func AddToStringAny(str string, object any) string {
	ObjectStr := fmt.Sprintf("%v", object)
	return (str + ObjectStr)
}

func ConvertStringToRuneSlice(str string) []rune {
	return []rune(str)
}

func HashSHA256WithSalt(slice []rune, salt string) (string, error) {
	if salt == "" {
		return "", fmt.Errorf("salt cannot be empty")
	}
	mid := len(slice) / 2
	sliceWithSalt := append(slice[:mid], []rune(salt)...)
	sliceWithSalt = append(sliceWithSalt, slice[:mid]...)

	hash := sha256.New()
	hash.Write([]byte(string(sliceWithSalt)))
	return string(hash.Sum(nil)), nil
}
