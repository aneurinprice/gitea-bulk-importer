package internal

var Args struct {
	Input	string   `arg:"positional"`
	Type	string   `arg:"-t,--type" help:"User or Org to import" validate:"required,oneof=User user Org org"`
	LogLevel	string    `arg:"-l,--log-level" help:"Desired LogLevel" validate:"omitempty,oneof=Trace Debug Info Warning Error Fatal Panic"`
	DryRun	bool   `arg:"-d,--dry-run" help:"Do not import, just print what would be imported"`
	IncludeForks	bool `arg:"-f,--forks" help:"Include/Exclude forks in the import"`
	Rename	string	`arg:"--rename" help:"Rename User/Org in Gitea" validate:"omitempty,alphanum"`
	Regex	string	`arg:"-r,--regex" help:"Regex to filter repo names" validate:"omitempty"`
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
	Password: "################",
}

var GiteaLogin = giteaCredentials{
	Username:	"aneurinprice",
	Password:	"#############",
	GiteaUrl:	"https://gitea.services.nyeprice.local",
}