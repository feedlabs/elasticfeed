feedify
=======

what should be done
-------------------
![structure](https://cloud.githubusercontent.com/assets/1843523/2949716/c1d6ebd2-da11-11e3-9932-3138175c32e8.png)

#### Development
Create symbolic link from you directory to `GOPATH`
```
cd to-your-project
mkdir -p $GOPATH/src/github.com/feedlabs
ln -s $(pwd) $GOPATH/src/github.com/feedlabs/feedify
```

#### Run
`go run test/feedify.go`

#### Dependencies
`go get github.com/astaxie/beego`
