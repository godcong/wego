package wego

type Application interface {
	Link(string) string
	GetRequest() Request
}

//type application struct {
//	Config
//	Reqeust
//}

//func NewApplication(config Config, reqeust Reqeust) Application {
//	panic("s")
//}
