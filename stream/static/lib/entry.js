var Entry = (function() {

  function Entry(data) {
    /** @type {String} */
    this.id = null;

    /** @type {Integer} */
    this.ts = null;

    /** @type {Feed} */
    this.parent = null;
  }

  // Management

  Entry.prototype.update = function() {
  }

  Entry.prototype.remove = function() {
  }

  Entry.prototype.hide = function() {
  }

  Entry.prototype.show = function() {
  }

  // Events callbacks

  Entry.prototype.onData = function(callback) {
  }

  // Handlers

  Entry.prototype.registerHandlers = function() {
    // bind to feed events
  }

  // Helpers

  Entry.prototype.getParent = function() {
  }

  Entry.prototype.getTimestamp = function() {
    return this.ts;
  }

  return Entry;

})();
