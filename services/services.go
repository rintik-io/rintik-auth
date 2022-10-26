package services

type ServicesInterface interface {
	Start() error
}

func Run(services ServicesInterface) error {
	err := services.Start()

	return err
}
