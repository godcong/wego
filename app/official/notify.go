package official

import (
	"github.com/godcong/wego/util"
	"net/http"
)

// Notify ...
type Notify interface {
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

// NotifyCallback ...
type NotifyCallback func(p util.Map) (util.Map, error)

// NotifyFunc ...
type NotifyFunc func(w http.ResponseWriter, req *http.Request)

/*Notify 监听 */
type messageNotify struct {
	*Account
	NotifyCallback
}

func (n *messageNotify) ServeHTTP(w http.ResponseWriter, req *http.Request) {

}
