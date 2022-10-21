<p align="center">
    <p align="center">
        <img src="https://www.guardrails.io/assets/images/logo-color.png" alt="GuardRails" title="GuardRails" height="100px" align="center"/>
    </p>
    <h1 align="center"><b>GuardRails CLI</b></h1>
</p>

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/guardrailsio/guardrails-cli)
![GitHub](https://img.shields.io/github/license/guardrailsio/guardrails-cli)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/guardrailsio/guardrails-cli?sort=semver)
[![GuardRails CLI CI](https://github.com/guardrailsio/guardrails-cli/actions/workflows/ci.yaml/badge.svg)](https://github.com/guardrailsio/guardrails-cli/actions/workflows/ci.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/guardrailsio/guardrails-cli)](https://goreportcard.com/report/github.com/guardrailsio/guardrails-cli) 
![Codecov](https://img.shields.io/codecov/c/github/guardrailsio/guardrails-cli?token=3c5e84bf-caa3-4a07-ace2-64f67b86a244)
[![GuardRails badge](https://api.guardrails.io/v2/badges/145474?token=5cdcd3c9f602bdf5dd4ec8a7d19e2e6599e571a73e8c9751f5b6d04deaf68aa6)](https://dashboard.guardrails.io/gh/guardrailsio/repos/145474)

GuardRails CLI provides tool to interact with [GuardRails](https://www.guardrails.io) via command line. See documentation for more details at https://docs.guardrails.io/docs
## Installation

### Installation scripts (Linux / OSX)

Just copy paste this command and you're good to go. We're assuming that you're using `bash` but you can change it accordingly based on the shell that you're using. You might be asked for a password for `sudo` in the process of installation.

```
curl -fsSL https://raw.githubusercontent.com/guardrailsio/guardrails-cli/master/etc/scripts/install.sh | bash
```

### Brew (Linux / OSX)

Alternatively, you can also install `guardrails` via brew:

```
brew tap guardrailsio/guardrails
brew install guardrails
```

### Windows

You need to have [scoop](https://scoop.sh) installed in order to install `guardrails`. The rest will be similar as using installation scripts for Linux / OSX. You just need to copy paste this command into your powershell:

```
iex ((new-object net.webclient).DownloadString('https://raw.githubusercontent.com/guardrailsio/guardrails-cli/master/etc/scripts/install.ps1'))
```

## Usage

GuardRails CLI have several commands:

- `scan` : scans a repository for vulnerabilities and output results
- `version` : shows build version
 
You can take a look at available commands and usage example via help menu: `guardrails --help` 

## License

GuardRails CLI is released under the Apache 2.0 license. See [LICENSE.txt](https://github.com/guardrailsio/guardrails-cli/blob/master/LICENSE.txt)

## Contribution