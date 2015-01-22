/*
 * Author: Feed Labs
 */

(function(window) {

  var Feed = {

    /** @type {String} */
    id: null,

    /** @type {String} */
    channelId: null,

    /** @type {Object} */
    entryList: {},

    /** @type {Array} */
    defaultEntryIds: [],

    /** @type {Object} */
    options: {
      feedId: '',
      outputContainerId: 'defaultContainerId',
      defaultElementLayout: '',
      defaultElementCount: 0
    },

    init: function(options, stylerFunction) {
      this.options = this._extend(this.options, options);
      this._stylerFunction = stylerFunction || this._stylerFunction;
      this.outputContainer = document.getElementById(this.options.outputContainerId);
      this._addDefaultEntries();

      var _this = this;
      setTimeout(function() {
        elasticfeed.load('http://www.feed.dev:10111/v1/feed/' + _this.options.feedId + '/entry', function(httpRequest) {
          _this._loadFirstEntries(JSON.parse(httpRequest.responseText));
        });
      }, 1500);
    },

    addEntry: function() {

    },

    findEntry: function(id) {

    },

    _stylerFunction: function(data) {
      return JSON.stringify(data.Data);
    }
  }

  var Entry = {

    /** @type {String} */
    id: null,

    init: function() {
    },

    update: function() {
    },

    remove: function() {
    }
  }

  var Channel = {

    /** @type {String} */
    id: null,

    init: function() {
    }
  }

  var MetricEvent = {

    /** @type {String} */
    id: null,

    init: function() {
    }
  }

  var StreamEvent = {

    /** @type {String} */
    id: null,

    init: function() {
    }
  }

  var elasticfeed = {

    /** @type {Object} */
    channelList: {},

    getFeed: function(id) {

    },

    getChannel: function(id) {

    },

    newFeed: function(options) {

    },

    newChannel: function(options) {

    },

    /**
     * @param {Object} a
     * @param {Object} b
     * @returns {Object}
     */
    _extend: function(a, b) {
      var c = {}, prop;
      for (prop in a) {
        if (a.hasOwnProperty(prop)) {
          c[prop] = a[prop];
        }
      }
      for (prop in b) {
        if (b.hasOwnProperty(prop)) {
          c[prop] = b[prop];
        }
      }
      return c;
    },

    _uniqueId: function() {
      return '_' + Math.random().toString(36).substr(2, 9);
    },


    load: function(url, callback) {
      var xhr;
      if (typeof XMLHttpRequest !== 'undefined') {
        xhr = new XMLHttpRequest();
      } else {
        var versions = ["MSXML2.XmlHttp.5.0", "MSXML2.XmlHttp.4.0", "MSXML2.XmlHttp.3.0", "MSXML2.XmlHttp.2.0", "Microsoft.XmlHttp"];
        for (var i = 0, len = versions.length; i < len; i++) {
          try {
            xhr = new ActiveXObject(versions[i]);
            break;
          }
          catch (e) {
          }
        }
      }
      xhr.onreadystatechange = ensureReadiness;
      function ensureReadiness() {
        if (xhr.readyState < 4) {
          return;
        }
        if (xhr.status !== 200) {
          return;
        }
        if (xhr.readyState === 4) {
          callback(xhr);
        }
      }

      xhr.open('GET', url, true);
      xhr.send('');
    }

  }; // elasticfeed

  // ===========================================================================

  if ("function" === typeof define) {
    define(function(require) {
      return elasticfeed;
    });
  } else {
    window.elasticfeed = elasticfeed;
  }
}(window));


// Helpers
// =======

Element.prototype.remove = function() {
  this.parentElement.removeChild(this);
};
NodeList.prototype.remove = HTMLCollection.prototype.remove = function() {
  for (var i = 0, len = this.length; i < len; i++) {
    if (this[i] && this[i].parentElement) {
      this[i].parentElement.removeChild(this[i]);
    }
  }
};
