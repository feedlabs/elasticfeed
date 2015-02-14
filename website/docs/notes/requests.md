API requests
------------

#### Auth
##### Superuser
`curl  -u x-super-user:supersecret 192.168.1.51:10100/v1/org --digest`
##### Admin
`curl  -u john:hello 192.168.1.51:10100/v1/application --digest`

#### User
`curl  -u john:hello -XPOST -H "content-type:application/json" -d '{"data":"NEW ADMIN", "Maintainer":true,"Username":"kris", "Whitelist":["127.0.0.1", "10.10.10.1"]}' 127.0.0.1:10100/v1/adming --digest`

#### Examples
```
curl feed.dev:10111/v1/feed/
{}

curl -XPOST -H "content-type:application/json" -d '{"id": <int>, "tags": [], "data":">new feed page"}' feed.dev:10111/v1/feed
{
  "id": "1409731977215584330"
}

curl -XPOST -H "content-type:application/json" -d '{"data":"new feed entry"}' feed.dev:10111/v1/feed/1409731977215584330/entry
{
  "id": "1409731990598877748"
}

curl -XPOST -H "content-type:application/json" -d '{"from-source-id": []<int>}' feed.dev:10111/v1/feed/1409731977215584330/entry
{
  "id": "1409731990598877748"
}

curl -XPUT -H "content-type:application/json" -d '{"data":"updated new feed page"}' feed.dev:10111/v1/feed/1409731977215584330
{
  "result": "update success",
  "status": "ok"
}

curl -XPUT -H "content-type:application/json" -d '{"data":"updated new feed entry"}' feed.dev:10111/v1/feed/1409731977215584330/entry/1409731990598877748
{
  "result": "update success",
  "status": "ok"
}

curl -XDELETE feed.dev:10111/v1/feed/1409731977215584330/entry/1409731990598877748
{
  "result": "delete success",
  "status": "ok"
}

curl -XDELETE feed.dev:10111/v1/feed/1409731977215584330
{
  "result": "delete success",
  "status": "ok"
}

curl feed.dev:10111/v1/feed/
{}
```
