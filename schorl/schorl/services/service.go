package services

type Service interface {
	Init() error
	Start()
	Stop()
	Destroy()
}
