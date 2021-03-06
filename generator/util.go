package generator

import (
	"sort"
	"strings"
	"unicode"
)

func toPublicName(name string) string {
	if name == "" {
		return ""
	}
	return strings.ToUpper(name[0:1]) + name[1:]
}

func concatSortedMap(m map[string]string, sep string) string {
	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	s := ""
	for _, k := range keys {
		s += m[k] + sep
	}
	return s
}

func interfaceSliceToStringSlice(iSlice []interface{}) ([]string, bool) {
	var ok bool
	stringSlice := make([]string, len(iSlice))
	for i, v := range iSlice {
		stringSlice[i], ok = v.(string)
		if !ok {
			return nil, false
		}
	}
	return stringSlice, true
}

// Make filenames snake-case, taken from https://gist.github.com/elwinar/14e1e897fdbe4d3432e1
func toSnake(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}
