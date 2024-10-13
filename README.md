# gitea-bulk-exporter
This should probably be `gitea-bulk-importer` but I have already created the repo now.

# Description
A tool to configure gitea mirroring of a github user/org into gitea. It will read the remote user/org and create a corresponding org in gitea. The tool will then (as configured) import repositories from the github user/org into gitea as repository mirrors.

## Installation
To install this project, follow these steps:

1. Clone the repository:
    ```
    git clone https://github.com/aneurinprice/gitea-bulk-exporter.git
    ```
2. Navigate to the project directory:
    ```
    cd gitea-bulk-exporter
    ```
3. Install the package (optional):
    ```
    go install cmd/gitea-bulk-importer/gitea-bulk-exporter.go 
    ```

## Usage

To use this project, follow these steps:

1. Run the Binary:
    ```
    gitea-bulk-exporter --help
or
    go run cmd/gitea-bulk-importer/gitea-bulk-exporter.go --help
    ```

## Runtime Options

```
Usage: gitea-bulk-exporter [--type TYPE] [--log-level LOG-LEVEL] [--dry-run] [--forks] [--rename RENAME] [--regex REGEX] [INPUT]

Positional arguments:
  INPUT

Options:
  --type TYPE, -t TYPE   User or Org to import
  --log-level LOG-LEVEL, -l LOG-LEVEL
                         Desired LogLevel
  --dry-run, -d          Do not import, just print what would be imported
  --forks, -f            Include/Exclude forks in the import
  --rename RENAME, -r RENAME
                         Rename User/Org in Gitea
  --regex REGEX, -s REGEX
                         Regex to filter repo names
  --help, -h             display this help and exit
```

## Credentials
Set the following environmental variables:

```
GITHUB_USERNAME="YourUsername"
GITHUB_PASSWORD="YourGithubPassword"
GITEA_USERNAME="YourUsername"
GITEA_PASSWORD="YourGiteaPassword"
GITEA_URL="https://your.gitea.host"
```

## License
This project is licensed under the MIT License.