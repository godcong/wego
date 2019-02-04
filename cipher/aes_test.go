package cipher

import (
	"encoding/xml"
	"github.com/godcong/wego/util"
	"strings"
	"testing"
)

// TestNewDataCrypt ...
func TestNewDataCrypt(t *testing.T) {
	//{"openid":"oE_gl0Yr54fUjBhU5nBlP4hS2efo","session_key":"v4Hn+ZHjpCD9wjU53cP0hw=="}
	//encrypted pLpcOaA1Z1nWaka2JkchrbAekCCWiU6+iSxHCFer8SM4nHEhOQMwinxx8lAOmg12tHu9hW1Ah3ghiT0ULjqU+/X2maXSiYUWBMyc36QX4BZB29JMnAzm2zycaGwmUX90WKv/ZmGh6UW4YU8/kj6WALNYlM7KpEJib6I3zSqP1irgCkKYoM1Bn7IWqJI+FNlCguPPXMoPoXDdycrfOR2CGDKN6gKCFAf4OHzv0lCggH12jCy0USRoAxZRcEGo2nhBmQBwi0jePRQEoBQ3H0Cn5sOQD5SGjZWsS/2pg0k6ABUXUZI/QRt1Gi5DSMG48W/Az75b3cui5lkxN4Tq0kwXs2UV0h3qR+66qlqJIQQjYKBbdKnZhw9CQnWg4k/3Ft28uiTa4LERRYHvMwzGBLsb6wQjGXXhkN9U8CR0XvpBbbQ9Jk2OYU0JDrkG8Jwx0KY/j2jYWPWu1I0ppDJXsRAyvrpUy4AqEOIst4gjpoMWvQ0=
	//iv u+SFMW4Rifsg3MwQ+KnNRA==
	c := NewDataCrypt("wx1ad61aeef1903b93")
	m, e := c.Decrypt("pLpcOaA1Z1nWaka2JkchrbAekCCWiU6+iSxHCFer8SM4nHEhOQMwinxx8lAOmg12tHu9hW1Ah3ghiT0ULjqU+/X2maXSiYUWBMyc36QX4BZB29JMnAzm2zycaGwmUX90WKv/ZmGh6UW4YU8/kj6WALNYlM7KpEJib6I3zSqP1irgCkKYoM1Bn7IWqJI+FNlCguPPXMoPoXDdycrfOR2CGDKN6gKCFAf4OHzv0lCggH12jCy0USRoAxZRcEGo2nhBmQBwi0jePRQEoBQ3H0Cn5sOQD5SGjZWsS/2pg0k6ABUXUZI/QRt1Gi5DSMG48W/Az75b3cui5lkxN4Tq0kwXs2UV0h3qR+66qlqJIQQjYKBbdKnZhw9CQnWg4k/3Ft28uiTa4LERRYHvMwzGBLsb6wQjGXXhkN9U8CR0XvpBbbQ9Jk2OYU0JDrkG8Jwx0KY/j2jYWPWu1I0ppDJXsRAyvrpUy4AqEOIst4gjpoMWvQ0=", "u+SFMW4Rifsg3MwQ+KnNRA==", "v4Hn+ZHjpCD9wjU53cP0hw==")
	t.Log(string(m), e)

}

// TestRefundedNotify_ServeHTTP ...
func TestRefundedNotify_ServeHTTP(t *testing.T) {
	//xmlData := []byte(`<xml><return_code>SUCCESS</return_code><appid><![CDATA[wx15810ca16324c225]]></appid><mch_id><![CDATA[1516796851]]></mch_id><nonce_str><![CDATA[ad306db6cd97f12492169f2b7a805ff8]]></nonce_str><req_info><![CDATA[XXeD817ddbzJMC1Zsb10XjOnveEaVHqv99DZ/fOL9cOvaXuvp2ZjfJeFqkDBPbHlHbyGyRW7FKT2Hy7Zj8eGC/Zz4hgyX84rUJatxhpw+W9bJRyqm6xuo20dEMcLqa0CZ44J8jcSjp3bvmi9yDmnpETSorUojhnoOL3qAVOi8d8J4X5r9cOlm4MWbvDgseMBBD4c7sGrSUl9P7p1sDomr5C/p052JZjYkgWSYquQp1UzuYO0Ol6utJ9yPupX1d1OGcBnY3upVCSFsfaeXendvSD5Rzs+chyx/t2JcgohvvtqH4225lAiA+6rksMPFolSdHy4qS5kuJNEEZSKqW7igpf6hVTXkWKRW6SWuppf1MbyFYB3JwJueQdTXzUOuYBth4RMSmoKBFhdf0t60GBFpPqo1iBEsCHKDhzOAL5CKksUx5CiD8XimyyylBn5ULTnImK8qJAZCELdTsQnFcyWJfV1QtNkDXsYr/nA2AHHjvVMahPH3zKeXRPFCgEBk3JprP9zIvLVFDwWuXcfJniLpsB+sc5NR4SZptx3A9+3nHWiBnu3riIrAsgwc754kiSAarZeblrbFtlpoaQpra5ixKRTpsGCdwq8NBvEcseVxubRm+AdeEc+gIoilqP2H+Eu2Zvwoqj/UQKEXR0Bg7498j0V+0SjGiomqdXwtauZia9S9HG6SqvcWPBrfdTtEuQyv/eW9zq/QDxGQlPMfGhZq6hpvXFkZG1TRHaejaymM6AGWtBLkV1ypUTvlB4zqyzPtD6jRrRck6VYcK81U+Y7P3sWekLlTAX7C04PPRoITmpG93p+DUszbHubFYMEc+TEH1mlBOix6fpzzSbPlG0BKE/foH+9oYiW5B2i103SqjlfvoHUW+cizVMBgDTT+5OvDmrozgyFyt8CrU7Gq15KIR9xQTQZK68NaATrR11UxmnpOcGJt0+LIXH/EtZIxg5nwKys4AmEkbgRMn6xbGGz0mDw1gDya8QJCxwU4KLpz0tcst2YBes3OqlYtXcm7Ow2hpjG5dybeAc1fHYnOOmPl0M08nHAad4NwF3uPo+eIto=]]></req_info></xml>`)
	maps := util.Map{}
	//err := xml.Unmarshal(xmlData, &maps)
	//t.Log(maps)
	//t.Log(err)
	reqInfo := `XXeD817ddbzJMC1Zsb10XjOnveEaVHqv99DZ/fOL9cOvaXuvp2ZjfJeFqkDBPbHlHbyGyRW7FKT2Hy7Zj8eGC/Zz4hgyX84rUJatxhpw+W9bJRyqm6xuo20dEMcLqa0CZ44J8jcSjp3bvmi9yDmnpETSorUojhnoOL3qAVOi8d8J4X5r9cOlm4MWbvDgseMBBD4c7sGrSUl9P7p1sDomr5C/p052JZjYkgWSYquQp1UzuYO0Ol6utJ9yPupX1d1OGcBnY3upVCSFsfaeXendvSD5Rzs+chyx/t2JcgohvvtqH4225lAiA+6rksMPFolSdHy4qS5kuJNEEZSKqW7igpf6hVTXkWKRW6SWuppf1MbyFYB3JwJueQdTXzUOuYBth4RMSmoKBFhdf0t60GBFpPqo1iBEsCHKDhzOAL5CKksUx5CiD8XimyyylBn5ULTnImK8qJAZCELdTsQnFcyWJfV1QtNkDXsYr/nA2AHHjvVMahPH3zKeXRPFCgEBk3JprP9zIvLVFDwWuXcfJniLpsB+sc5NR4SZptx3A9+3nHWiBnu3riIrAsgwc754kiSAarZeblrbFtlpoaQpra5ixKRTpsGCdwq8NBvEcseVxubRm+AdeEc+gIoilqP2H+Eu2Zvwoqj/UQKEXR0Bg7498j0V+0SjGiomqdXwtauZia9S9HG6SqvcWPBrfdTtEuQyv/eW9zq/QDxGQlPMfGhZq6hpvXFkZG1TRHaejaymM6AGWtBLkV1ypUTvlB4zqyzPtD6jRrRck6VYcK81U+Y7P3sWekLlTAX7C04PPRoITmpG93p+DUszbHubFYMEc+TEH1mlBOix6fpzzSbPlG0BKE/foH+9oYiW5B2i103SqjlfvoHUW+cizVMBgDTT+5OvDmrozgyFyt8CrU7Gq15KIR9xQTQZK68NaATrR11UxmnpOcGJt0+LIXH/EtZIxg5nwKys4AmEkbgRMn6xbGGz0mDw1gDya8QJCxwU4KLpz0tcst2YBes3OqlYtXcm7Ow2hpjG5dybeAc1fHYnOOmPl0M08nHAad4NwF3uPo+eIto=`
	payKey := "aTKnSUcTkbEnhwQNdutWkQxAjnhAz2jK"

	key := strings.ToLower(string(util.SignMD5(payKey, "")))
	ecb := CryptAES256ECB()
	ecb.SetParameter("key", []byte(key))
	d, err := ecb.Decrypt([]byte(reqInfo))
	t.Error(err)
	t.Log(string(d))
	err = xml.Unmarshal(d, &maps)
	t.Log(maps)
	t.Log(err)
}
