elasticfeed
===========

what should be done
-------------------
![feed-overview - api 1](https://cloud.githubusercontent.com/assets/1843523/5789062/2c0d6ca2-9e5b-11e4-89fd-5d636ece1ed6.png)

#### Development
Create symbolic link from you directory to `GOPATH`
```
cd to-your-project
mkdir -p $GOPATH/src/github.com/feedlabs
ln -s $(pwd) $GOPATH/src/github.com/feedlabs/elasticfeed
```

#### Run
`go run api.go`

#### Dependencies
* `go get github.com/feedlabs/feedify` [repo](https://github.com/feedlabs/feedify)

#### Documentation

##### Install

```
$ npm install apidoc -g
```

##### Generate

```
$ apidoc -i public/ -o docs/api
```

##### View
```
$ open docs/api/index.html
```
