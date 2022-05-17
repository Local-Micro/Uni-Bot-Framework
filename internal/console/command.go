package console

type Command interface {
	Start(args []string) error
}
