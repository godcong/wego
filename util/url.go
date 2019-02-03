package util

/*URL 拼接地址 */
func URL(prefix string, uri string) string {
	end := len(prefix)
	if end > 1 && prefix[end-1] == '/' {
		prefix = prefix[:end-1]
	}
	uend := len(uri)
	if uend > 1 && uri[0] == '/' {
		uri = uri[1:]
	}
	return prefix + "/" + uri
}
