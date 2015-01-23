var StreamEvent = (function() {

  const CHANNEL_JOIN = 0
  const CHANNEL_LEAVE = 1
  const CHANNEL_MESSAGE = 2

  const ACTION_FEED = 1
  const ACTION_ENTRY = 2

  const FEED_RELOAD = 1
  const FEED_RESET = 2
  const FEED_DATA_INIT = 3
  const FEED_DATA_MORE = 4
  const FEED_HIDE = 5
  const FEED_SHOW = 6

  const ENTRY_ADD = 1
  const ENTRY_DELETE = 2
  const ENTRY_UPDATE = 3
  const ENTRY_HIDE = 4
  const ENTRY_SHOW = 5

  function StreamEvent(event) {

    console.log(event)

    /** @type {String} */
    this.id = null;

    /** @type {Integer} */
    this.ts = event.Timestamp;

    /** @type {String} */
    this.User = event.User

    /** @type {String} */
    this.Type = event.Type

    /** @type {String} */
    this.ContentType = 'string'

    /** @type {String} */
    try {
      this.Content = JSON.parse(event.Content)
      this.ContentType = 'json'
    } catch (e) {
      this.Content = event.Content
    }
  }

  StreamEvent.prototype.GetType = function() {
    switch (this.Type) {
      case CHANNEL_JOIN:
        console.log(this.User + " joined the chat room");
        break;
      case CHANNEL_LEAVE:
        console.log(this.User + " left the chat room");
        break;
      case CHANNEL_MESSAGE:
        console.log(this.User + ", " + this.PrintContent());
        break;
    }
  }

  StreamEvent.prototype.GetTimestamp = function() {
    return this.ts;
  }

  StreamEvent.prototype.PrintContent = function() {

    if (this.ContentType == 'string') {
      return this.Content
    }

    return JSON.stringify(this.Content)

    switch (this.Content.ActionGroup) {
      case ACTION_FEED:
        switch (this.Content.Action) {
          case FEED_DATA_INIT:
            break;
        }
        break;
      case ACTION_ENTRY:
        switch (this.Content.Action) {
          case ENTRY_ADD:
            break;
        }
        break;
    }

  }

  return StreamEvent;

})();
