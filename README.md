# Mail-GO
[![License](https://img.shields.io/github/license/Disconnect24/Mail-GO.svg?style=flat-square)](http://www.gnu.org/licenses/agpl-3.0)
![Production List](https://img.shields.io/discord/206934458954153984.svg?style=flat-square)
[![Go Report Card](https://goreportcard.com/badge/github.com/Disconnect24/Mail-GO?style=flat-square)](https://goreportcard.com/report/github.com/Disconnect24/Mail-GO)

This is an effort to rewrite Wii Mail legacy PHP scripts into golang.
Some reasons why:
- `apache2` has the fun tendency to go overboard on memory usage.
- `go` is fun.

### Read the wiki for setting up your own Mail-GO server!

# How to develop
The source is entirely here, with each individual cgi component in their own file.
A `Dockerfile` is available to create an image. You can use `docker-compose.yml` to develop on this specific component with its own mysql, or use *something that doesn't yet exist* to develop on DC24 as a whole.
You can use `docker-compose up` to start up both MySQL and Mail-GO.

# How can I use the patcher for my own usage?
You're welcome to `POST /patch` with a `nwc24msg.cfg` under form key `uploaded_config`.

# What should I do if I'm adding a new dependency?
We use Go's 1.11+ module feature. Make sure you have this enabled. For more information, see [the Go wiki](https://github.com/golang/go/wiki/Modules).