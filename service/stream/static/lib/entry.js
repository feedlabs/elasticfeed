var Entry = (function() {

  const ACTION_GROUP_TYPE = 2

  const ACTION_ADD = 1
  const ACTION_DELETE = 2
  const ACTION_UPDATE = 3
  const ACTION_HIDE = 4
  const ACTION_SHOW = 5

  const PUBLISH_METRIC = 6

  /** @type {Entry} */
  var localCache = {}

  function Entry(data) {
    /** @type {String} */
    this.id = null;

    /** @type {String} */
    this.data = data;
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

  Entry.prototype.onBeforeUpdate = function(callback) {

  }

  Entry.prototype.onAfterUpdate = function(callback) {

  }

  Entry.prototype.onBeforeRemove = function(callback) {

  }

  Entry.prototype.onAfterRemove = function(callback) {

  }

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
