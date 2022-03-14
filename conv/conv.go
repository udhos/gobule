package conv

import (
	"strconv"
	"strings"
)

/*
// InterfaceList converts []string to []interface{}.
func InterfaceList(strList []string) []interface{} {
	list := make([]interface{}, 0, len(strList))
	for _, s := range strList {
		list = append(list, s)
	}
	return list
}
*/

// VersionToNumber converts version numbers to number.
// 1,2,3 => 100020003
func VersionToNumber(v1, v2, v3 int64) int64 {
	return v1*100000000 + v2*10000 + v3
}

// VersionToString converts version string to string.
// "1.2.3" => "100020003"
func VersionToString(version string) string {
	var v1, v2, v3 int64
	s := strings.Split(version, ".")

	if len(s) > 0 {
		v1 = atoi(s[0])
	}
	if len(s) > 1 {
		v2 = atoi(s[1])
	}
	if len(s) > 2 {
		v3 = atoi(s[2])
	}

	v := VersionToNumber(v1, v2, v3)
	return strconv.FormatInt(v, 10)
}

func atoi(s string) int64 {
	v, _ := strconv.ParseInt(s, 10, 64)
	return v
}
