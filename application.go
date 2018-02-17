package wego

type Application interface {
	//Link(string) string
	GetRequest() Request
	GetPayment() Payment
}

//type application struct {
//	Config
//	Reqeust
//}

func NewApplication(config Config) Application {
	panic("s")
}
