package common

import "strings"

func StringSplice(strLIST ...string) string {
	var build strings.Builder
	for _,val := range strLIST{
		build.WriteString(val)
	}
	return build.String()
}
