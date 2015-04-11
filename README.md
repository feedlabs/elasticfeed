elasticfeed
===========
#### Overview
![elasticfeed-overview - server-engine 1](https://cloud.githubusercontent.com/assets/1843523/7102998/033b0ace-e095-11e4-90d8-16d46a694792.png)

#### Internal workflow
![elasticfeed-overview - mission](https://cloud.githubusercontent.com/assets/1843523/7103001/212978e0-e095-11e4-8b23-091adefe3cb7.png)

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
