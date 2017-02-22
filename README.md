# bugsnag-data

Dump Bugsnag data using the Data access API

## API documentation

https://docs.bugsnag.com/api/data-access/

## Data access API key

The key can be found under the organization section of https://app.bugsnag.com/settings/

## Installation

```
go get -u github.com/peterhellberg/bugsnag-data
```

## Usage example

```bash
$ bugsnag-data -key [REDACTED] /accounts | jq .[0].account_creator

GET https://api.bugsnag.com/accounts
{
  "account_admin": true,
  "email": "[REDACTED]",
  "html_url": "https://app.bugsnag.com/settings/[REDACTED]/collaborators/[REDACTED]",
  "id": "[REDACTED]",
  "name": "[REDACTED]",
  "projects_url": "https://api.bugsnag.com/users/[REDACTED]/projects",
  "url": "https://api.bugsnag.com/users/[REDACTED]"
}
```
