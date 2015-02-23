Feed API
--------

NEW ORG
-------
```
curl -u john:hello -XPOST -H "content-type:application/json" -d '{"data":"NEW ORG"}' localhost:10100/v1/org --digest
```

NEW APP
-------
```
curl -u john:hello -XPOST -H "content-type:application/json" -d '{"data":"NEW APP"}' localhost:10100/v1/application --digest
```

NEW FEED
-------
```
curl -u john:hello -XPOST -H "content-type:application/json" -d '{"data":"NEW FEED"}' localhost:10100/v1/application/XXX/feed --digest
```

NEW FEED
-------
```
curl -u john:hello -XPOST -H "content-type:application/json" -d '{"data":"NEW ENTRY"}' localhost:10100/v1/application/XXX/feed/XXX/entry --digest
```


POST
----
```
curl -XPOST -H "Content-type:application/json" -d '{"Data":"{\"channel\":\"iO5wshd5fFE5YXxJ/hfyKQ==:17\",\"event\":\"CM_Action_Abstract:SEND:31\",\"data\":{\"action\":{\"actor\":{\"_type\":33,\"_id\":{\"id\":\"1\"},\"id\":1,\"displayName\":\"user1\",\"visible\":true,\"_class\":\"Feed_Model_User\"},\"verb\":13,\"type\":31,\"_class\":\"Feed_Action_Feed\"},\"model\":{\"_type\":33,\"_id\":{\"id\":\"1\"},\"id\":1,\"displayName\":\"user1\",\"visible\":true,\"_class\":\"Feed_Model_User\"},\"data\":{\"action\":\"add\",\"clientData\":{\"type\":\"photo\",\"imageList\":[{\"url\":\"http://lorempixel.com/500/500/?4651277\",\"title\":\"Image\"},{\"url\":\"http://lorempixel.com/500/500/?8551474\",\"title\":\"Image\"}],\"likeCount\":0,\"id\":19}}}}"}' http://localhost:10100/v1/feed
```

Subscribe
---------
```
  message, _ = stream.NewStreamMessage()

	channels := []string{"socket-redis-down"}
	message.Subscribe(channels, func(timeout bool, message string, channel string) {
		if !timeout {
			fmt.Println("publish:", message, " channel:", channel)
		} else {
			fmt.Println("error: sub timedout")
		}
	})
```

Reload
------
```
curl -u john:hello 127.0.0.1:10100/v1/application/30/feed/89/reload --digest
{
  "result": "Feed reloaded",
  "status": "ok"
}
```

Empty
-----
```
curl -u john:hello 127.0.0.1:10100/v1/application/30/feed/86/empty --digest
{
  "result": "Feed empty done.",
  "status": "ok"
}
```
