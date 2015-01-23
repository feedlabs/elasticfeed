var StreamEvent = (function() {

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

  StreamEvent.prototype.GetTimestamp = function() {
    return this.ts;
  }

  StreamEvent.prototype.PrintContent = function() {
    if (this.ContentType == 'string') {
      return this.Content
    }
    return JSON.stringify(this.Content)
  }

  return StreamEvent;

})();
