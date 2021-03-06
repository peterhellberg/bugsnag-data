# bugsnag-data

[![Build Status](https://travis-ci.org/peterhellberg/bugsnag-data.svg?branch=master)](https://travis-ci.org/peterhellberg/bugsnag-data)
[![Go Report Card](https://goreportcard.com/badge/github.com/peterhellberg/bugsnag-data)](https://goreportcard.com/report/github.com/peterhellberg/bugsnag-data)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/bugsnag-data#license-mit)

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

```sh
Usage of bugsnag-data:
  -delay duration
    	The delay between each request (default 1s)
  -key string
    	Data access API key (required)
  -max int
    	Max number of requests to make (0 means unlimited)
```

The first argument after the flags have been parsed is used as the path (of the first request to the Data access API).

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

`bugsnag-data` will follow next in the [Link header](https://tools.ietf.org/html/rfc5988)
until reaching a response with no Link header or having done the max number of requests
as specified in the max command line flag.

## License (MIT)

Copyright (c) 2017 [Peter Hellberg](https://c7.se/)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
