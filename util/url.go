package util

import "strings"

/*URL 拼接地址 */
func URL(prefix string, uris ...string) string {
	end := len(prefix)
	if end > 1 && prefix[end-1] == '/' {
		prefix = prefix[:end-1]
	}

	var url = []string{prefix}
	for _, v := range uris {
		url = append(url, TrimSlash(v))
	}
	return strings.Join(url, "/")
}

// TrimSlash ...
func TrimSlash(s string) string {
	if size := len(s); size > 1 {
		if s[size-1] == '/' {
			s = s[:size-1]
		}
		if s[0] == '/' {
			s = s[1:]
		}
	}
	return s
}
