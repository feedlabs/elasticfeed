var Entry = (function() {

  const GROUP_TYPE = 2

  const ADD = 1
  const DELETE = 2
  const UPDATE = 3
  const HIDE = 4
  const SHOW = 5

  /** @type {Entry} */
  var localCache = {}

  function Entry(feed, id, data, styler) {
    /** @type {String} */
    this.id = id;

    /** @type {String} */
    this.data = data;

    /** @type {Object} */
    this._feed = feed;

    /** @type {Function} */
    this._styler = styler;

    this.bindFeedMessages();
  }

  // UI

  Entry.prototype.render = function() {
    document.getElementById(this.id).innerHTML = this._styler.call(this, this.data);
  }

  // Events

  Entry.prototype.on = function(type, callback) {
    switch (name) {
      case 'add':
        type = ADD
        break;
      case 'delete':
        type = DELETE
        break;
      case 'update':
        type = UPDATE
        break;
      case 'hide':
        type = HIDE
        break;
      case 'show':
        type = SHOW
        break;
      default:
        return false;
        break;
    }
    if (this._handlers[type] == undefined) {
      this._handlers[type] = []
    }
    this._handlers[type].push(callback);
    return true;
  }

  Entry.prototype.onData = function(entryEvent) {
    switch (entryEvent.type) {
      case ADD:
        this.onAdd(entryEvent.ts, entryEvent.content)
        break;
      case DELETE:
        this.onDelete(entryEvent.ts)
        break;
      case UPDATE:
        this.onUpdate(entryEvent.ts, entryEvent.content)
        break;
      case HIDE:
        this.onHide(entryEvent.ts)
        break;
      case SHOW:
        this.onShow(entryEvent.ts)
        break;
    }
  }

  // Management

  Entry.prototype.update = function() {
  }

  Entry.prototype.delete = function() {
  }

  Entry.prototype.hide = function() {
  }

  Entry.prototype.show = function() {
  }

  // Events callbacks

  Entry.prototype.onAdd = function(timestamp, data) {
    for (var i in this._handlers[ADD]) {
      this._handlers[ADD][i].call(this, timestamp, data);
    }
  }

  Entry.prototype.onUpdate = function(timestamp, data) {
    for (var i in this._handlers[UPDATE]) {
      this._handlers[UPDATE][i].call(this, timestamp, data);
    }
  }

  Entry.prototype.onDelete = function(timestamp) {
    for (var i in this._handlers[DELETE]) {
      this._handlers[DELETE][i].call(this, timestamp);
    }
  }

  Entry.prototype.onDelete = function(timestamp) {
    for (var i in this._handlers[HIDE]) {
      this._handlers[HIDE][i].call(this, timestamp);
    }
  }

  Entry.prototype.onDelete = function(timestamp) {
    for (var i in this._handlers[SHOW]) {
      this._handlers[SHOW][i].call(this, timestamp);
    }
  }

  // Handlers

  Entry.prototype.bindFeedMessages = function() {
    var self = this;
    this._feed.on('entry', function(ts, entryEvent) {
      if (entryEvent.id == self.id || entryEvent.id == '*') {
        self.onData(entryEvent);
      }
    });
  }

  // Helpers

  Entry.prototype.getTimestamp = function() {
    return this.ts;
  }

  return Entry;

})();
