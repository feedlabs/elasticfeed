elasticfeed
===========
#### Overview
![elasticfeed-overview - server-engine 2](https://cloud.githubusercontent.com/assets/1843523/7103135/4fb3d170-e099-11e4-9c50-8fb7090331a5.png)

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
