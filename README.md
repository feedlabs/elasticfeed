api
===

what should be done
-------------------
![structure](https://cloud.githubusercontent.com/assets/1843523/4171533/32530bd4-353e-11e4-96a9-5be825db2923.png)

#### Development
Create symbolic link from you directory to `GOPATH`
```
cd to-your-project
mkdir -p $GOPATH/src/github.com/feedlabs
ln -s $(pwd) $GOPATH/src/github.com/feedlabs/api
```

#### Run
`go run api.go`

#### Dependencies
* `go get github.com/feedlabs/feedify`
