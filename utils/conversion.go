package utils

import "strconv"

func StrToInt(val string) int {
	valC, _ := strconv.Atoi(val)
	return valC
}

func StrBool(val string) (bool, error) {
	res, err := strconv.ParseBool(val)
	return res, err
}

// func ParseBool(str string) (bool, error) {
// 	switch str {
// 	case "1", "t", "T", "true", "TRUE", "True":
// 		return true, nil
// 	case "0", "f", "F", "false", "FALSE", "False":
// 		return false, nil
// 	}
// 	return false, syntaxError("ParseBool", str)
// }
