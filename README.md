elasticfeed
===========
#### Overview
![feed-overview - api 4](https://cloud.githubusercontent.com/assets/1843523/6098783/26b3c180-afe9-11e4-9eb8-dc7908f28344.png)

#### Internal workflow
![feed-overview - mission](https://cloud.githubusercontent.com/assets/1843523/6098048/8b9a1fb6-afd2-11e4-8f80-fde4ad59b51d.png)

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
