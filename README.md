# esq

Esq is a command line utility to query elasticsearch, inspired by [elktail](https://github.com/knes1/elktail). Kibana is awesome as a search interface, but isn't that useful for scanning through a long stream of logs, and doesn't integrate with the myriad cli tools available. Esq is an opinionated way to query from the command line and pipe the output to other tools.

## Feature Requests

Please feel free to use the [Issue Tracker](https://github.com/astropuffin/esq/issues) if you have any feature ideas or requests (and, of course, to report bugs).

# Installation

#### Install Using Go

Esq is written in go, so you'll need to [set up your go environment](https://golang.org/doc/install) or use a [docker container](https://hub.docker.com/r/library/golang).

`go get github.com/dronedeploy/esq`

This will automatically download, compile and install the tool.
After that you should see `esq` in `$GOPATH/bin`.

#### Install Using Hombrew (OSX)

To install esq using homebrew run the following:

`brew update || brew update` # yes you have to update twice. This is a bug in brew itself.

`brew tap dronedeploy/tap`

`brew install esq`

*NOTE:* if you are getting an authentication error during `brew tap`, it may be because brew tries to use HTTPS instead of SSH to access github. To fix this, add the following two lines to your `~/.gitconfig` file:

```
[url "git@github.com:"]
    insteadOf = "https://github.com/"
```


#### Configuring

To configure your connection to elasticsearch, run the following:

`esq config --config ~/.esq.yml`

This will generate default config file in your home directory. Edit it, entering elasticsearch URL, username and password, and the fields you want to include in logs. 

Example config file contents:
```
url: http://elasticsearch.ddops.cool:9200
username: <your Kibana/elasticsearch username>
password: <your Kibana/elasticsearch password>
timestamp: '@timestamp'
index: kube-*
fields: '@timestamp,kubernetes.container,log'
```
