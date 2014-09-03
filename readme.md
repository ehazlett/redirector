# Redirector
Simple HTTP redirector in Go.

# Usage

`./redirector -url <url> -code <code>`

For example:

`./redirector -url http://foo.com -code 301`

There is also a Docker image for this also:

`docker run -P ehazlett/redirector -url <url> -code <code>`
