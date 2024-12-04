package cast

import "strconv"

func ToInt(arg interface{}) int {
	var val int
	var err error
	val, err = strconv.Atoi(arg.(string))
	if err != nil {
		print("Error converting string to int: " + err.Error())
	}
	return val
}

func IntToByteSlice(in int) []byte {
	return []byte(strconv.Itoa(in))
}

func ToRunes(v rune) rune {
	return v
}
