package wego

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func TestConfigLoader(t *testing.T) {

	m := md5.New()
	m.Write([]byte("These pretzels are making me thirsty."))

	fmt.Printf("%x", m.Sum(nil))

}
