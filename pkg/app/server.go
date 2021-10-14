package app


type Server interface {
	Run() error
	Stop()
}
