var Event = (function() {

  function Event(event) {

    /** @type {String} */
    this.id = event.Id || null;

    /** @type {Integer} */
    this.ts = event.Timestamp;

    /** @type {Integer} */
    this.actionGroup = null

    /** @type {Integer} */
    this.actionType = null

    /** @type {String} */
    this.user = event.User

    /** @type {String} */
    this.type = event.Type

    /** @type {String} */
    this.contentType = 'string'

    /** @type {String} */
    try {
      this.content = JSON.parse(event.Content)
      this.contentType = 'json'
    } catch (e) {
      this.content = event.Content
    }
  }

  Event.prototype.GetTimestamp = function() {
    return this.ts;
  }

  Event.prototype.PrintContent = function() {
    if (this.contentType == 'string') {
      return this.content
    }
    return JSON.stringify(this.content)
  }

  return Event;

})();
