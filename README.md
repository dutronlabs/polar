# polar: API Layer for Container Vulnerability

## NOTE: still in active development

## Why?

First, I love coreos/clair - but I required something more for my CI builds. When developing a new app or workload in a container, having 
a CI build automatically check for vulnerabilities I may have introduced by adding in new linux packages, etc. was a massive improvement
over manually having to check with clair and analyze-image. This makes my life easier, and I hope it makes yours!

## License 

MIT License

## Install

### Quick
_TODO I'll get around to adding a docker compose file to this_
* docker run -p 6379:6379 redis:latest
* docker run -p 9001:9001 dutronlabs/polar:latest

### Source
* git clone github.com/dutronlabs/polar.git
* docker build . -t dutronlabs/polar
* docker run --name polar 