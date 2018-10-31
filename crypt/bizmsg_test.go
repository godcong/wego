package crypt

import "testing"

var encodingAesKey = "TNwHN28RXXoyVxkMCUEqKuCL08eBpCKgWZTkWNVnGLu"
var token = "godcong"
var timeStamp = "1524409354"
var nonce = "542437598"
var appID = "wxbafed7010e0f4531"
var text = `<xml>
    <ToUserName><![CDATA[gh_56870ffd193b]]></ToUserName>
    <FromUserName><![CDATA[oLyBi0hSYhggnD-kOIms0IzZFqrc]]></FromUserName>
    <CreateTime>1524409354</CreateTime>
    <MsgType><![CDATA[text]]></MsgType>
    <Content><![CDATA[你好啊你好啊你好啊]]></Content>
    <MsgID>6547288321577417974</MsgID>
</xml>`

var text1 = `<xml><ToUserName><![CDATA[gh_56870ffd193b]]></ToUserName><RSAEncrypt><![CDATA[iiCKU5aC+BE0DDhjW8qWvqfQGkIgEVNSYI3SSlaLy9xq7VUKMUFW7jXH1VBX4ZpkRJLpiSoXqSyF2S7hclV37IpphXNzQpKwwP6UvoSuZNQyhF7bQraLm3QmxBV1JNt/tH5qoV1nPIwmj/tgdIDNfiTkMi8We1984Sb+T6lB6zPMsaIRTCXHdV+5/yx98veVv3MTY3nkmFCR738wxbQ1wZxqQyuHs8AYBWAByVbm5MCdrwO8KF2xxvnX1Zneng+UjbNVh9KCWllYoNIQPgGpy2y9HGlwcYNwtPRomfb/dWYr1J43aaVMIrh8KU/cJH3V0fF/zdX0yTpNAWyMhYP2fUHARpr9qBFWacbFTcAuBMaNTeFlFUvgRb/sM3G9wRkEFm1okMcDz7o4vqE03ZAwT9BPyjr3sYBpTdgq4CHj4cKgw2+W32m+PvAa/BFmLMCSWutJExu/ze4SfkJO/3xCzw==]]></RSAEncrypt></xml>`
var text2 = `<xml><Nonce>1632909179</Nonce><RSAEncrypt><![CDATA[lAqgapbsGq3hpZC29u5OJLMOwSGZCDfCWsKFV1M7Ig2ljZMMxAB9MFqpsJItJM1BjYI4ER0lmjuFYK9X4KNR4uA8J3Gng/50vZwTsHAD2TSOkkIhAXpFczAQlRFN/r790jjg6VS0ZrfUChYapVl5CvGdqDNFRskNIVX+ikXjvRM0V3ZPKE5CZp9f/JRk/iVskKOKNK9p8DApDppngz5+y2gtWWtO2NCap2v9GI1Gs5GqtoRSzC5TbOeEM/YO4lsB651PIZrGM4Dq417C8yDY8/RHMLxwt+ogoeeYq2a7+/HCmLeY8YhswhxUBuV80VNlMFVJxTfY+GBfxHoz7gRH/MxBJ/NvT8LiLbfenuA/BPiggWA/vIzNFY0XO07Q6ZZKkGZCCMa104s+V/mfca+OIuYAse9I+B4um/2nF1Y1Bso=]]></RSAEncrypt><MsgSignature><![CDATA[08d28bc8bb189eea2d9b704d9781be2057fd4f30]]></MsgSignature><TimeStamp>1524416866</TimeStamp></xml>`
var text3 = `<?xml version="1.0" encoding="UTF-8" standalone="no"?><xml><RSAEncrypt><![CDATA[8YHvi544ufqOnTylGkwEkCtB/jf8THDLV7v9Q5FctW/Z4Y0Ied5B1Ch0mKhoMJpXylnqlfOFAovhUA8WDBhQSUparcfbx/WPMLUXXJRjgbtsde4fPII0vFyAeaiwlNeoiL17zhYRISdlMd55elzVxAYG6VQ+89MOcZ0p5YwjKwZfTXPLl2ZO5ADW6tVqjFld3DfGGNOP3yRtMaWqrCQo4ASk5bpOpCuYTd5p3dXygkKv5LwQyb+MB/xdt+Z4MeVWN0Wke+HE29iJWikvKUV9d0pNU81R+8PrTrsGs/4gtI/Nl5w5JKoxwZKSYhpVzoJvgvxu+z9UkoN/81BYY/AoPkI51fcRjcAXrViDN0TR+/EeDFd0KKnuoP6X8AtTm0JD3w68dSEjmT9U8CNFxydJsF3bYh37D7LeKuhXZDMA7vqTV2PF7LfiFer8UkcGnVNP]]></RSAEncrypt><MsgSignature><![CDATA[8c0f8d64124367eccb5f292dad91955eb0cd12d8]]></MsgSignature><TimeStamp>1524421916</TimeStamp><Nonce>457570794</Nonce></xml>`

// TestBizMsg_Encrypt ...
func TestBizMsg_Encrypt(t *testing.T) {
	biz := NewBizMsg(token, encodingAesKey, appID)
	var result string
	var result0 []byte
	var err error
	result, err = biz.Encrypt(text, timeStamp, nonce)
	t.Log(result, err)
	result0, err = biz.Decrypt(text3, "8c0f8d64124367eccb5f292dad91955eb0cd12d8", "1524421916", "457570794")
	t.Log(string(result0), err)
}
