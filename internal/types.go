package internal

var Args struct {
	Input	string   `arg:"positional"`
	Type	string   `arg:"-t,--type" help:"User or Org to import" validate:"required,oneof=User user Org org"`
	Verbose	bool     `arg:"-v,--verbose" help:"verbosity level"`
	DryRun	string   `arg:"--dry-run" help:"Do not import, just print what would be imported"`
}

type githubCredentials struct {
	Username string
	Password string
}

type giteaCredentials struct {
	Username string	
	Password string
	GiteaUrl string
}

var GithubLogin = githubCredentials{
	Username: "aneurinprice",
	Password: "#########################################",
}

var GiteaLogin = giteaCredentials{
	Username:	"aneurinprice",
	Password:	"#########################################",
	GiteaUrl:	"#########################################",
}