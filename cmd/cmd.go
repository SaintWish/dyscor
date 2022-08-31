package cmd

var (
	cmds = make(map[string]Command)
	errCmdExists = errors.New("Command already exists")
)

type Command struct {

}