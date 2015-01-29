var SystemEvent = (function() {

  function SystemEvent(chid, event) {

    /** @type {String} */
    this.chid = chid;

    /** @type {Integer} */
    this.ts = event.Timestamp;

    /** @type {Integer} */
    this.actionGroup = null

    /** @type {Integer} */
    this.actionType = null

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

  SystemEvent.prototype.GetTimestamp = function() {
    return this.ts;
  }

  SystemEvent.prototype.PrintContent = function() {
    if (this.contentType == 'string') {
      return this.content
    }
    return JSON.stringify(this.Content)
  }

  return SystemEvent;

})();
