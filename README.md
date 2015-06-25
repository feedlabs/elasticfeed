![logo - new page 6](https://cloud.githubusercontent.com/assets/1843523/7351311/b4db8698-ed04-11e4-8a01-a5b72ba55163.png)
===========
#### Overview
![elasticfeed-overview - server-engine 3](https://cloud.githubusercontent.com/assets/1843523/8316372/d38e7464-19f6-11e5-8170-a1111ce0f31c.png)

#### Internal workflow
![elasticfeed-overview - mission](https://cloud.githubusercontent.com/assets/1843523/7103001/212978e0-e095-11e4-8b23-091adefe3cb7.png)

#### Development environment
Create symbolic link from you directory to `GOPATH`
```
cd to-your-project
mkdir -p $GOPATH/src/github.com/feedlabs
ln -s $(pwd) $GOPATH/src/github.com/feedlabs/elasticfeed
```

#### Developing Elasticfeed
If you wish to work on Elasticfeed itself or any of its built-in providers,
you'll first need [Go](http://www.golang.org) installed (version 1.2+ is
_required_). Make sure Go is properly installed, including setting up
a [GOPATH](http://golang.org/doc/code.html#GOPATH).

Next, install the following software packages, which are needed for some dependencies:

- [Bazaar](http://bazaar.canonical.com/en/)
- [Git](http://git-scm.com/)
- [Mercurial](http://mercurial.selenic.com/)

Then, install [Gox](https://github.com/mitchellh/gox), which is used
as a compilation tool on top of Go:

    $ go get -u github.com/mitchellh/gox

Next, clone this repository into `$GOPATH/src/github.com/elasticfeed/elasticfeed`.
Install the necessary dependencies by running `make updatedeps` and then just
type `make`. This will compile some more dependencies and then run the tests. If
this exits with exit status 0, then everything is working!

    $ make updatedeps
    ...
    $ make
    ...

To compile a development version of Elasticfeed and the built-in plugins,
run `make dev`. This will put Elasticfeed binaries in the `bin` folder:

    $ make dev
    ...
    $ bin/elasticfeed
    ...


If you're developing a specific package, you can run tests for just that
package by specifying the `TEST` variable. For example below, only
`elasticfeed` package tests will be run.

    $ make test TEST=./elasticfeed
    ...


#### Run
Environment variables for plugins
```
export ELASTICFEED_PLUGIN_MIN_PORT=40000
export ELASTICFEED_PLUGIN_MAX_PORT=41000
```
Compilation and running
```
go run api.go
```

#### Dependencies
* `go get github.com/feedlabs/feedify` [repo](https://github.com/feedlabs/feedify)

#### Documentation

##### Install

```
$ npm install apidoc -g
```

##### Generate

```
$ apidoc -i service/db -o docs/api
```

##### View
```
$ open docs/api/index.html
```
