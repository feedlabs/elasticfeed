/*
 * Author: Feed Labs
 */

(function(window) {
  var FeedPlugin = {

    /** @type {Object} */
    objectList: {},

    /** @type {Array} */
    defaultEntryIds: [],

    /** @type {HTMLElement|Null} */
    outputContainer: null,

    /** @type {Object} */
    options: {
      feedId: '',
      outputContainerId: 'defaultContainerId',
      defaultElementLayout: '',
      defaultElementCount: 0
    },

    /**
     * @param {Object|Null} options
     * @param {Function|Null} stylerFunction
     */
    init: function(options, stylerFunction) {
      this.options = this._extend(this.options, options);
      this._stylerFunction = stylerFunction || this._stylerFunction;
      this.outputContainer = document.getElementById(this.options.outputContainerId);
      this._addDefaultEntries();

      var _this = this;
      setTimeout(function() {
        FeedPlugin.load('http://www.feed.dev:10111/v1/feed/' + _this.options.feedId + '/entry', function(httpRequest) {
          _this._loadFirstEntries(JSON.parse(httpRequest.responseText));
        });
      }, 1500);
    },

    /**
     * @param {Array} firstEntries
     */
    _loadFirstEntries: function(firstEntries) {
      var _this = this;
      for (var key in firstEntries) {
        if (firstEntries.hasOwnProperty(key)) {
          var data = firstEntries[key];
          _this.add(data);
        }
      }
      this._removeDefaultEntries();
    },

    _addDefaultEntries: function() {
      for (var i = 1; i <= this.options.defaultElementCount; i++) {
        var objectId = this._uniqueId();
        var entry = document.createElement('div');
        entry.id = objectId;
        entry.innerHTML = this.options.defaultElementLayout;
        this.outputContainer.appendChild(entry);

        this.defaultEntryIds.push(objectId);
      }
    },

    _removeDefaultEntries: function() {
      this.defaultEntryIds.forEach(function(id) {
        var domObject = document.getElementById(id);
        domObject.remove();
      });
    },

    /**
     * @param {Object} data
     */
    processData: function(data) {
      switch (data.Action) {
        case 'add':
          this.add(data);
          break;
        case 'remove':
          this.remove(data);
          break;
        case 'update':
          this.update(data);
          break;
        default:
          console.log('Unknown action `' + data.Action + '`');
      }
    },

    /**
     * @param {Object} data
     */
    add: function(data) {
      var objectId = data.Id;
      this.objectList[objectId] = objectId;

      var domElement = document.createElement('div');
      domElement.id = objectId;
      domElement.innerHTML = this._stylerFunction(data.Data);

      this.outputContainer.insertBefore(domElement, this.outputContainer.firstChild);
    },

    /**
     * @param {Object} data
     */
    remove: function(data) {
      var domElement = document.getElementById(data.Id);
      domElement.remove();
      delete this.objectList[data.Id];
    },

    /**
     * @param {Object} data
     */
    update: function(data) {
      var domElement = document.getElementById(data.Id);
      domElement.innerHTML = this._stylerFunction(data.Data);
    },

    /**
     * @param {Object} data
     * @returns {String}
     */
    _stylerFunction: function(data) {
      return JSON.stringify(data.Data);
    },

    // Helpers
    // =======

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

  }; // FeedPlugin

  // ===========================================================================

  if ("function" === typeof define) {
    define(function(require) {
      return FeedPlugin;
    });
  } else {
    window.FeedPlugin = FeedPlugin;
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
