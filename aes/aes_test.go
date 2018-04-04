package aes_test

import (
	"testing"

	"github.com/godcong/wego/aes"
)

func TestNewAES128CBCDataCrypt(t *testing.T) {
	c := aes.NewAES128CBCDataCrypt("wx1ad61aeef1903b93", "012SzBR809PzDH1rAcO80MAzR80SzBRo")
	c.Decrypt("KFDDLBew/+pVFpq1NEa28s0PxLUng3MjBjv2lNv66K6ySixY8wLnuWPrznIexjEDwXluCoTV8GQWaEw8SOTXMrfrOMqexmSr0keR+nM4Pf5k1twaGIaNMuTl8WeG5UGVpHKRn+2l+eDn+zUxxztZMnH3c4yUBtzifBMf+a79alBhJSulKoNbMDFQQtAaWkYG7e6mJ/CGcDpf5tvIO42mg38plVdNrYMMjI5uObZhShGw7ShkJWsj7uXq7XUJW96uaCR+/88jZq0Vxkr3rVjnTEAyYH8ZTVQY0vt2ynVX1Ddd9fC59XrjxL9Rvgx8w3tLHPxVk1cFkMCOszjRTbDdEJU/ViOqp7/cPT4nTFOw0+UdzVqG/NPZ/IH/wjCxY8X796VV7FsMzygLi92TM+GM2a0ij/9Wla8HeMpqVqulYpYaUbg112LGYlZQSkUByrfiF6zLB984kq3p4R7SrIyxkv4ux39QxMg5F5w69iFADfQ=", "SNJV25jtoJUm0J7XEYjnHA==")
}
