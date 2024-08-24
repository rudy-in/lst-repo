# lst-repo

## Description

**`lst-repo`** is a command-line utility written in Go that allows users to fetch and list GitHub repositories for a specified GitHub user. It uses the GitHub API to retrieve and display repository information in a user-friendly format.

## Features

- Fetches public repositories from GitHub.
- Utilizes GitHub's API with a personal access token for authentication.
- Simple command-line interface for efficient usage.

## Requirements

- [Go](https://golang.org/doc/install) 1.18 or higher.
- A GitHub personal access token with appropriate permissions.

# Setup

## 1. Clone the Repository

Clone this repository to your local machine:

```bash
git clone https://github.com/yourusername/lst-repo.git
cd lst-repo
```

# 2. Setup ENV variable 

This tool is available for both **Unix** and **DOS** environments

## Unix
to run this tool on unix

```bash
export GITHUB_TOKEN="your github token"
chmod 775 ./run
chmod +x ./run
./run
```


# Command Prompt and Powershell (For DOS)

## CMD PROMPT
```bash
set GITHUB_TOKEN=your_generated_token_here
cd src
go run *.go
```

## POWERSHELL
```bash
$env:GITHUB_TOKEN="your_generated_token_here"
cd src
go run *.go
```


### [IMPORTANT] Change the value of the following variable before running 

**In main.go line 21 `username := "your_username"` change `your_username` to your own github username example `username := "rudy-in"`**

### Contact

If any problem occurs please contact the developer at atsharma623@gmail.com

### Acknowledgments

- Special thanks to the GitHub community for their robust API documentation and support.

- Thanks to the Go community for their excellent development tools and libraries.


**Happy hacking!**