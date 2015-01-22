var Entry = (function() {

  function Entry(data) {
    /** @type {String} */
    this.id = null;

    /** @type {Integer} */
    this.ts = null;

    /** @type {Feed} */
    this.parent = null;
  }

  Entry.prototype.update = function() {
  }

  Entry.prototype.remove = function() {
  }

  Entry.prototype.getTimestamp = function() {
    return this.ts;
  }

  return Entry;

})();
