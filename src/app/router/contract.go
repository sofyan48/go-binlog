package router

//

type Router interface {
	LogRouter()
}

type LogProcessor interface {
	Exec() error
}
