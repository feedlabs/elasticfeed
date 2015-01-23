var StreamEvent = (function() {

  const JOIN = 0
  const LEAVE = 1
  const MESSAGE = 2

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
    } catch(e) {
      this.Content = event.Content
    }
  }

  StreamEvent.prototype.GetType = function() {
    switch (this.Type) {
      case JOIN:
        console.log(this.User + " joined the chat room");
        break;
      case LEAVE:
        console.log(this.User + " left the chat room");
        break;
      case MESSAGE:
        console.log(this.User + ", " + this.PrintContent());
        break;
    }
  }

  StreamEvent.prototype.GetTimestamp = function() {
    return this.ts;
  }

  StreamEvent.prototype.PrintContent = function() {
    return this.ContentType == 'json' ? JSON.stringify(this.Content) : this.Content
  }

  return StreamEvent;

})();
