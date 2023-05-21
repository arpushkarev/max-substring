package cli

type cli struct {
	url string
}

func NewCli() (*cli, error) {

	return &cli{
		url: "",
	}, nil
}
