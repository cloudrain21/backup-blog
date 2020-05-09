package commander

type Commander interface {
	Connect() error
	RunCommand(command string) error
	Close()
}
