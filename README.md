api
===

what should be done
-------------------
![structure](https://cloud.githubusercontent.com/assets/1843523/4171691/0072acc6-3540-11e4-92be-2f0e77233b12.png)

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
* `go get github.com/feedlabs/feedify` [repo](https://github.com/feedlabs/feedify)
