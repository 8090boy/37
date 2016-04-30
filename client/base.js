var Sys = (function(ua) {
    var s = {};
    s.IE = ua.match(/msie/ig) ? true : false;
    s.Firefox = ua.match(/Firefox/ig) ? true : false;
    s.Chrome = ua.match(/Chrome/ig) ? true : false;
    s.Safari = ua.match(/Safari/ig) ? true : false;
    return s
}
)(navigator.userAgent.toLowerCase());
function launchFullscreen(element) {
    if (element.requestFullscreen) {
        element.requestFullscreen()
    } else {
        if (element.mozRequestFullScreen) {
            element.mozRequestFullScreen()
        } else {
            if (element.webkitRequestFullscreen) {
                element.webkitRequestFullscreen()
            } else {
                if (element.msRequestFullscreen) {
                    element.msRequestFullscreen()
                }
            }
        }
    }
}
function exitFullScreen() {
    if (document.exitFullscreen) {
        document.exitFullscreen()
    } else {
        if (document.mozCancelFullScreen) {
            document.mozCancelFullScreen()
        } else {
            if (document.webkitCancelFullScreen) {
                document.webkitCancelFullScreen()
            } else {
                if (document.msExitFullscreen) {
                    document.msExitFullscreen()
                }
            }
        }
    }
}
var $ = function(id) {
    return "string" == typeof id ? document.getElementById(id) : id
}
;
function Each(list, fun) {
    for (var i = 0, len = list.length; i < len; i++) {
        fun(list[i], i)
    }
}
var Css = function(e, o) {
    if (typeof o == "string") {
        e.style.cssText = o;
        return
    }
    for (var i in o) {
        e.style[i] = o[i]
    }
}
;
var Attr = function(e, o) {
    for (var i in o) {
        e.setAttribute(i, o[i])
    }
}
;
var $$ = function(p, e) {
    return p.getElementsByTagName(e)
}
;
var Extend = function(destination, source) {
    for (var property in source) {
        destination[property] = source[property]
    }
}
;
var Bind = function(object, fun) {
    var args = Array.prototype.slice.call(arguments).slice(2);
    return function() {
        return fun.apply(object, args)
    }
}
;
var BindAsEventListener = function(object, fun, args) {
    var args = Array.prototype.slice.call(arguments).slice(2);
    return function(event) {
        return fun.apply(object, [event || window.event].concat(args))
    }
}
;
var CurrentStyle = function(element) {
    return element.currentStyle || document.defaultView.getComputedStyle(element, null )
}
;
var Getpos = function(o) {
    var x = 0
      , y = 0;
    do {
        x += o.offsetLeft;
        y += o.offsetTop
    } while ((o = o.offsetParent));return {
        "x": x,
        "y": y
    }
}
;
function addListener(element, e, fn) {
    element.addEventListener ? element.addEventListener(e, fn, false) : element.attachEvent("on" + e, fn)
}
function removeListener(element, e, fn) {
    element.removeEventListener ? element.removeEventListener(e, fn, false) : element.detachEvent("on" + e, fn)
}
var Class = function(properties) {
    var _class = function() {
        return (arguments[0] !== null  && this.initialize && typeof (this.initialize) == "function") ? this.initialize.apply(this, arguments) : this
    }
    ;
    _class.prototype = properties;
    return _class
}
;
function loadjs(url) {
    var scriptTag = document.getElementById("loadjsid");
    if (!scriptTag) {
        scriptTag = document.createElement("script")
    }
    scriptTag.id = "loadjsid";
    scriptTag.src = url;
    document.head.appendChild(scriptTag)
}
function uuid(len, radix) {
    var chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz".split("");
    var uuid = [], i;
    radix = radix || chars.length;
    if (len) {
        for (i = 0; i < len; i++) {
            uuid[i] = chars[0 | Math.random() * radix]
        }
    } else {
        var r;
        uuid[8] = uuid[13] = uuid[18] = uuid[23] = "-";
        uuid[14] = "4";
        for (i = 0; i < 36; i++) {
            if (!uuid[i]) {
                r = 0 | Math.random() * 16;
                uuid[i] = chars[(i == 19) ? (r & 3) | 8 : r]
            }
        }
    }
    return uuid.join("")
}
function copy(e) {
    e = window.event || event;
    var el = e.srcElement || e.target;
    el.setAttribute("contenteditable", true);
    var s = document.getSelection();
    var r;
    if (s.type.toLowerCase() === "none") {
        r = document.createRange()
    } else {
        r = s.getRangeAt(0)
    }
    r.selectNodeContents(el);
    s.removeAllRanges();
    s.addRange(r);
    document.execCommand("selectAll", false, "");
    document.execCommand("copy", false, "");
    if (Sys.Chrome) {
        log.copyok(true, 3000, s.toString())
    }
}
var log = {
    copyok: function(status, time, txt) {
        var dialog = document.querySelector("#dialog");
        var mask = document.querySelector(".mask");
        var copytip = document.getElementById("copytip");
        if (/\?rh=|ri=/img.test(txt)) {
            txt = txt + "<br><a href='" + txt + '\' target="_self" >帮注册</a>'
        }
        copytip.innerHTML = "已复制：<h3>" + txt + "</h3>可直接粘贴！";
        if (status) {
            var t = setTimeout(function() {
                clearTimeout(t);
                log.copyok(false)
            }
            , time);
            dialog.style.display = "block";
            mask.style.display = "block";
            copytip.style.display = "block"
        } else {
            copytip.style.display = "none";
            mask.style.display = "none";
            dialog.style.display = "none"
        }
    }
};
var ajax = {
    GET: function(url, fn, o) {
        this._getXMLHttpRequest();
        var xht = this.xmlHttp;
        xht.onreadystatechange = this._forwordFunInIndex.bind(xht, fn);
        xht.open("GET", url, true);
        xht.withCredentials = true;
        xht.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
        xht.send(null )
    },
    POST: function(url, data, fn, o) {
        this._getXMLHttpRequest();
        var xht = this.xmlHttp;
        xht.onreadystatechange = this._forwordFunInIndex.bind(xht, fn);
        xht.open("POST", url, true);
        xht.withCredentials = true;
        xht.setRequestHeader("Content-Type", o ? o.type : "application/x-www-form-urlencoded");
        xht.send(data)
    },
    _forwordFunInIndex: function(fn) {
        if (this.readyState === 4 && this.status === 200) {
            var msg;
            try {
                msg = eval(this.responseText)
            } catch (e) {
                msg = JSON.parse(this.responseText)
            }
            if (!msg) {
                return
            }
            fn(msg)
        }
    },
    _getXMLHttpRequest: function() {
        if (this.xmlHttp) {
            this.xmlHttp = null 
        }
        if (window.XMLHttpRequest) {
            this.xmlHttp = new XMLHttpRequest();
            if (this.xmlHttp.overrideMimeType) {
                this.xmlHttp.overrideMimeType("text/xml")
            }
        } else {
            if (window.ActiveXObject) {
                try {
                    this.xmlHttp = new ActiveXObject("Msxml2.XMLHTTP")
                } catch (e) {
                    try {
                        this.xmlHttp = new ActiveXObject("Microsoft.XMLHTTP")
                    } catch (e) {}
                }
            }
        }
        return this.xmlHttp
    }
};
function DataCompp(id, type) {
    if (!type) {
        var inputs = document.querySelectorAll("#" + id + " input,#" + id + " textarea");
        var arg = [];
        for (var n = 0; n < inputs.length; n++) {
            arg.push(inputs[n].name.toLowerCase() + "=" + inputs[n].value)
        }
        return arg.join("&")
    }
    if (type === "json") {
        var inputs = document.querySelectorAll("#" + id + " input,#" + id + " textarea");
        var arg = [];
        for (var n = 0; n < inputs.length; n++) {
            arg.push('"' + inputs[n].name.toLowerCase() + '":"' + inputs[n].value + '"')
        }
        return "{" + arg.join(",") + "}"
    }
}
var cookie = {
    Get: function(name) {
        var arr, reg = new RegExp("(^| )" + name + "=([^;]*)(;|$)");
        if (arr = document.cookie.match(reg)) {
            return unescape(arr[2])
        } else {
            return null 
        }
    },
    Del: function(name) {
        var exp = new Date();
        exp.setTime(exp.getTime() - 1);
        var cval = this.Get(name);
        if (cval != null ) {
            document.cookie = name + "=" + cval + ";expires=" + exp.toGMTString()
        }
    },
    Set: function(name, value, time) {
        time = time || "h12";
        var strsec = this._getsec(time);
        var exp = new Date();
        exp.setTime(exp.getTime() + strsec * 1);
        document.cookie = name + "=" + escape(value) + ";expires=" + exp.toGMTString()
    },
    _getsec: function(str) {
        var str1 = str.substring(1, str.length) * 1;
        var str2 = str.substring(0, 1);
        if (str2 == "s") {
            return str1 * 1000
        } else {
            if (str2 == "h") {
                return str1 * 60 * 60 * 1000
            } else {
                if (str2 == "d") {
                    return str1 * 24 * 60 * 60 * 1000
                }
            }
        }
    }
};
var surface = {
    sectionArr: {
        my: "my",
        main: "main",
        index: "index",
        regpage: "regpage"
    },
    on: function(showWho) {
        for (var se in this.sectionArr) {
            var section = document.querySelector("#" + this.sectionArr[se]);
            if (section) {
                section.style.display = "none"
            }
        }
        if (!this.sectionArr[showWho]) {
            return
        }
        var section_show = document.querySelector("#" + showWho);
        section_show.style.display = "block"
    }
};
function StrToDate(str) {
    var tempStrs = str.split(" ");
    var dateStrs = tempStrs[0].split("-");
    var year = parseInt(dateStrs[0], 10);
    var month = parseInt(dateStrs[1], 10) - 1;
    var day = parseInt(dateStrs[2], 10);
    var timeStrs = tempStrs[1].split("-");
    var hour = parseInt(timeStrs[0], 10);
    var minute = parseInt(timeStrs[1], 10) - 1;
    var second = parseInt(timeStrs[2], 10);
    var date = new Date(year,month,day,hour,minute,second);
    return date
}
function count_down(obj) {
    var time_distance = obj.time_end - new Date();
    var int_day, int_hour, int_minute, int_second;
    if (time_distance > 0) {
        int_day = Math.floor(time_distance / 86400000);
        time_distance -= int_day * 86400000;
        int_hour = Math.floor(time_distance / 3600000);
        time_distance -= int_hour * 3600000;
        int_minute = Math.floor(time_distance / 60000);
        time_distance -= int_minute * 60000;
        int_second = Math.floor(time_distance / 1000);
        if (int_hour < 10) {
            int_hour = "0" + int_hour
        }
        if (int_minute < 10) {
            int_minute = "0" + int_minute
        }
        if (int_second < 10) {
            int_second = "0" + int_second
        }
        obj.time_day.innerHTML = int_day;
        obj.time_hour.innerHTML = int_hour;
        obj.time_minute.innerHTML = int_minute;
        obj.time_second.innerHTML = int_second;
        setTimeout(count_down.bind(this, obj), 1000)
    } else {
        obj.abcd.style.display = "none"
    }
}
function _WXShare(img, width, height, title, desc, url, appid) {
    img = img || "http://a.zhixun.in/plug/img/ico-share.png";
    width = width || 100;
    height = height || 100;
    title = title || document.title;
    desc = desc || document.title;
    url = url || document.location.href;
    appid = appid || "";
    function _ShareFriend() {
        WeixinJSBridge.invoke("sendAppMessage", {
            "appid": appid,
            "img_url": img,
            "img_width": width,
            "img_height": height,
            "link": url,
            "desc": desc,
            "title": title
        }, function(res) {
            _report("send_msg", res.err_msg)
        }
        )
    }
    function _ShareTL() {
        WeixinJSBridge.invoke("shareTimeline", {
            "img_url": img,
            "img_width": width,
            "img_height": height,
            "link": url,
            "desc": desc,
            "title": title
        }, function(res) {
            _report("timeline", res.err_msg)
        }
        )
    }
    function _ShareWB() {
        WeixinJSBridge.invoke("shareWeibo", {
            "content": desc,
            "url": url,
        }, function(res) {
            _report("weibo", res.err_msg)
        }
        )
    }
    document.addEventListener("WeixinJSBridgeReady", function onBridgeReady() {
        WeixinJSBridge.on("menu:share:appmessage", function(argv) {
            _ShareFriend()
        }
        );
        WeixinJSBridge.on("menu:share:timeline", function(argv) {
            _ShareTL()
        }
        );
        WeixinJSBridge.on("menu:share:weibo", function(argv) {
            _ShareWB()
        }
        )
    }
    , false)
}
;