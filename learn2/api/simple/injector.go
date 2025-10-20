//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializeService(isError bool) (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil
}

func InitializeDatabaseRepository() *DatabaseRepository {
	wire.Build(NewDatabaseRepository, NewDatabaseMysql, NewDatabasePostgres)
	return nil
}

// func InitializeFooBarService() *FooBarService {
// 	wire.Build(NewFooBarService, NewFooService, NewBarService, NewFooRepository, NewBarRepository)
// 	return nil
// }

// grouping set
var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializeFooBarService() *FooBarService {
	wire.Build(fooSet, NewFooBarService, barSet)
	return nil
}

// binding interface
var helloSet = wire.NewSet(NewSayHello, wire.Bind(new(SayHelloInterface), new(*SayHello)))

func InitializeHelloService() *HelloService {
	wire.Build(NewSayHelloService, helloSet)
	return nil
}

// cleanup
func InitializeConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}
