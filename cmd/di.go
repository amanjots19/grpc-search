package main

type diContainer struct {
	server func() *server
}

func newDIContainer() *diContainer {
	dic := &diContainer{}

	dic.server = newServerDIProvider()
	return &diContainer{}
}
