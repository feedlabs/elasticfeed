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

  SystemEvent.prototype.GetTimestamp = function() {
    return this.ts;
  }

  SystemEvent.prototype.PrintContent = function() {
    if (this.ContentType == 'string') {
      return this.Content
    }
    return JSON.stringify(this.Content)
  }

  return SystemEvent;

})();
