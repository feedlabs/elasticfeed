var StreamEvent = (function() {

  const ACTION_CHANNEL_JOIN = 0
  const ACTION_CHANNEL_LEAVE = 1
  const ACTION_CHANNEL_MESSAGE = 2

  function StreamEvent(event) {

    console.log(event)

    /** @type {String} */
    this.id = null;

    /** @type {Integer} */
    this.ts = event.Timestamp;

    /** @type {Integer} */
    this.actionGroup = null

    /** @type {Integer} */
    this.actionType = null

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
