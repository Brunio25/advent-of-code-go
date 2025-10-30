package cast

import (
	"fmt"
	"strconv"
)

// ToInt will convert an arg to an int
// Supported types:
//   - string
func ToInt(arg interface{}) int {
	var val int
	switch arg.(type) {
	case string:
		var err error
		val, err = strconv.Atoi(arg.(string))
		if err != nil {
			panic("error converting string to int " + err.Error())
		}
	default:
		panic(fmt.Sprintf("unhandled type for int casting %T", arg))
	}
	return val
}

// ToString will convert an arg to a string
// Supported types:
//   - int
//   - uint16
//   - rune
//   - byte
func ToString(arg interface{}) string {
	var str string
	switch arg.(type) {
	case int:
		str = strconv.Itoa(arg.(int))
	case uint16:
		str = strconv.FormatUint(uint64(arg.(uint16)), 10)
	case rune:
		str = string(arg.(rune))
	case byte:
		str = string(rune(arg.(byte)))
	default:
		panic(fmt.Sprintf("unsupported type for string cating %T", arg))
	}
	return str
}
