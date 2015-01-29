var EntryEvent = (function() {

  function EntryEvent(event) {

    /** @type {String} */
    this.id = null;

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

  EntryEvent.prototype.GetTimestamp = function() {
    return this.ts;
  }

  EntryEvent.prototype.PrintContent = function() {
    if (this.contentType == 'string') {
      return this.content
    }
    return JSON.stringify(this.content)
  }

  return EntryEvent;

})();
