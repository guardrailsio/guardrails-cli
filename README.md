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

The GuardRails CLI allows you to interact with [GuardRails](https://www.guardrails.io) via command line.
## Table of Contents

- [Pre-Requisites](#pre-requisites)
- [Installation](#installation)
- [Usage](#usage)
- [Documentation](#documentation)
- [License](#license)

## Pre-Requisites

To use the GuardRails CLI, you need to have an active account and a CLI token.

More information on how to get started can be found [here](https://www.guardrails.io/docs/en/getting-started).

Your GuardRails account CLI token can be found in the account's setting page on the dashboard.

## Installation

### Installation scripts (Linux / OSX)

Just paste this command and you're good to go. We're assuming that you're using `bash` but you can change it accordingly based on the shell that you're using. You might be asked for a password for `sudo` in the process of installation.

```
curl -fsSL https://raw.githubusercontent.com/guardrailsio/guardrails-cli/master/etc/scripts/install.sh | bash
```

### Brew (Linux / OSX)

Alternatively, you can also install `guardrails` via `brew`:

```
brew tap guardrailsio/guardrails
brew install guardrails
```

### Windows

You need to have [scoop](https://scoop.sh) installed in order to install `guardrails`. The rest will be similar to the installation scripts for Linux / OSX. You just need to paste this command into your powershell:

```
iex ((new-object net.webclient).DownloadString('https://raw.githubusercontent.com/guardrailsio/guardrails-cli/master/etc/scripts/install.ps1'))
```

## Usage

Here are the main GuardRails CLI commands:

- `scan` : scans a repository for vulnerabilities and output results
- `version` : displays build version
 
For more information on all the options and arguments available please check the help menu with: `guardrails --help` 

### How to read the results

The CLI will output the total number of vulnerabilities detected, if any.

Vulnerabilities detected are grouped by category, i.e. Hard-Coded Secrets.

For each item within a category, there will be a severity index (see table below), the type of vulnerability and a link to a fixing advice in our documentation; and finally the file path and line number.

Example: `(M) Hard-coded Secret - awesome-product/config.js:2`

Here we're looking at a vulnerability of type `Hard-coded secret` with a `Medium` severity in the file `awesome-product/config.js` at line `2`.

For Vulnerable Libraries specifically, the type of vulnerability will be replaced by the dependency name and version.

Example: `(C) pkg:gem/mypackage@2.5.2 - awesome-product/Gemfile.lock:14`

Here we're looking at the vulnerable `mypackage` dependancy in version `2.5.2` with a `Critical` severity declared in the file `awesome-product/Gemfile.lock` at line `14`.

#### Severity index table

| Index | Severity      |
|-------|---------------|
| (I)   | Informational |
| (L)   | Low           |
| (M)   | Medium        |
| (H)   | High          |
| (C)   | Critical      |

## Documentation

https://www.guardrails.io/docs/en/cli/introduction

## License

GuardRails CLI is released under the Apache 2.0 license. See [LICENSE.txt](https://github.com/guardrailsio/guardrails-cli/blob/master/LICENSE.txt)