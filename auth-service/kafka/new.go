package kafka

type UseCase interface {
	SendMsg(topic string, par int, msg string)
	ReceiveMsg(topic string, par int, msg string)
}

type implUseCase struct {
}

var _ UseCase = implUseCase{}

func New() UseCase {
	return implUseCase{}
}
