API requests
------------


#### Examples
```
curl feed.dev:10111/v1/feed/
{}

curl -XPOST -H "content-type:application/json" -d '{"data":"new feed page"}' feed.dev:10111/v1/feed
{
  "id": "1409731977215584330"
}

curl -XPOST -H "content-type:application/json" -d '{"data":"new feed entry"}' feed.dev:10111/v1/feed/1409731977215584330/entry
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
