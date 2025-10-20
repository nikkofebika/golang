package simple

type SayHelloInterface interface {
	Hello(string) string
}

type HelloService struct {
	SayHelloInterface
}

type SayHello struct{}

func (sayHello *SayHello) Hello(name string) string {
	return "Hello " + name
}

func NewSayHello() *SayHello {
	return &SayHello{}
}

func NewSayHelloService(sayHelloInterface SayHelloInterface) *HelloService {
	return &HelloService{SayHelloInterface: sayHelloInterface}
}
