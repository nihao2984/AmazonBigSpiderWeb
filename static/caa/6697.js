
(function() {

function getTime() {
    return (new Date).getTime();
}
function getCl() {
    var cl = "undefined" != typeof screen.pixelDepth ? screen.pixelDepth : screen.colorDepth;
    return "undefined" != typeof cl ? cl : '0';
}
function getDs() {
    return screen.width + "x" + screen.height;
}
function getLn() {
    var ln = "undefined" != typeof navigator.language ? navigator.language : navigator.browserLanguage;
    return "undefined" != typeof ln ? ln.toLowerCase() : 'unknown';
}
function getCkAfterLt() {
    return getCookie('qs_lvt_' + qsSi.toString()) == lt.toString ? 0 : 1
}
function getJa() {
    return navigator.javaEnabled() ? 1 : 0;
}
function getFl() {
    var flashVer = '0.0';
    if (window.ActiveXObject) {
        var swf = new ActiveXObject('ShockwaveFlash.ShockwaveFlash');
        if (swf) {
            var flashVerArr = swf.GetVariable("$"+"version")
                .split(' ')[1].replace(/,/g, '.').split('.');
            flashVer = flashVerArr[0] + '.' + flashVerArr[1]
        }
    } else {
        if (navigator.plugins && navigator.plugins.length > 0) {
            var swf = navigator.plugins['Shockwave Flash'];
            if (swf) {
                var arr = swf.description.split(' ');
                for (var i = 0, len = arr.length; i < len; i++) {
                    var ver = Number(arr[i]);

                    if (!isNaN(ver)) {
                        flashVer = arr[i];
                        break;
                    }
                }
            }
        }
    }
    return flashVer;
}

function getRandom(len) {
    return Math.random().toString().substr(2, len);
}
function isSameDomain(refer) {
    var domainArr = document.domain.split('.');     
    var ls = ['com', 'gov', 'net', 'org', 'ac', 'edu'];
    if (refer != null && refer != '') {
        if (refer.indexOf('//') > -1) {
            refer = refer.split("//")[1];    
        }
        var referArr = refer.split('/')[0].split(':')[0].split('.');
        var domLen = domainArr.length;
        var referLen = referArr.length;
        var ckLen = Math.min(domLen, referLen, 2);
        for (i in ls) {
            if (referArr[referLen - 2] == ls[i]) {
                ckLen = Math.min(domLen, referLen, 3); 
                break;
            }
        }
        for (var ind = 1; ind <= ckLen; ind++) {
            if (domainArr[domLen - ind] != referArr[referLen - ind]) {
                return false;
            }
        }
        return true;
    } else {
        return true;
    }
}
function getCookie(key) {
    var strCookie = document.cookie;
	if(strCookie != null)
	{
		var arrCookie = strCookie.split(";");
      for ( var i =0; i < arrCookie.length; i ++) {
			var arr = arrCookie[i].replace(/^\s*((?:[\S\s]*\S)?)\s*$/, '$1' ).split("=");
			if (arr[0]== key) {
				return arr[1];
			}
		}
	}
    return '';
}
function setCookie(key, value, expiresDays, path) {
    var strCookie = key + "=" + decodeURI(value);
    if (expiresDays > 0) {
        var date = new Date();
        date.setTime(date.getTime() + expiresDays * 3600 * 1000 * 24);
        strCookie = strCookie + ";expires=" + date.toUTCString();
    }
    strCookie = strCookie + ";path=" + path;
    document.cookie = strCookie;
}

function getLt() {
    var key = 'qs_lvt_' + qsSi.toString();
    var timestamp = parseInt(+new Date / 1000).toString();
    var ltlist = getCookie(key).split(',');
    var ltLast = parseInt(ltlist[ltlist.length - 1]); 

    if (!isNaN(ltLast)) { 
        if (Math.abs(ltLast-qsPageOpenTime) > 8 * 60 * 60 || !isSameDomain(top.document.referrer)) {
            if (ltlist.push(timestamp) > 5) {
                ltlist.shift()
            }
            setCookie(key, ltlist.toString(), 365, '/');
            isSameVisit = false;
            return ltlist;
        } else {
            isSameVisit = true;
            return ltlist;
        }
    } else {
        setCookie(key, timestamp, 365, '/');
        isSameVisit = false;
        return [timestamp];
    }
}

function setE360() {
    if (!isSameVisit) {
        var e_uid = "undefined" == typeof _e360_uid ? '' : _e360_uid;
        var e_cid = "undefined" == typeof _e360_campaignid ? '' : _e360_campaignid;
        var e_gid = "undefined" == typeof _e360_groupid ? '' : _e360_groupid;
        var e_yid = "undefined" == typeof _e360_creativeid ? '' : _e360_creativeid;
        var e_kid = "undefined" == typeof _e360_keywordid ? '' : _e360_keywordid;
        var e_com = "undefined" == typeof _e360_commerce ? '' : _e360_commerce;
        var e360 = '';
        if (e_uid != '') {
            e360 = e_uid + ',' + e_cid + ',' + e_gid + ',' + e_yid + ',' + e_kid;
        }
        if ( e_com != '') {
            if (e360 != '') {
               e360 = e360 + ',' + e_com;  
            }  
            else {
              e360 = e_com;  
            }
        }
        if (window.localStorage) {
            localStorage.setItem("s_e360", e360);
        } else {
            setCookie("s_e360", e360, 1, '/');
        }
    }
}

function getE360() {
    var e360 = window.localStorage ? localStorage.getItem("s_e360") : getCookie('s_e360');
    if (e360 == null) {
        e360 = '';
    }
	var arr = new Array();
    var arrE360 = e360.split(',');
    if ( arrE360.length == 1 && arrE360[0] != '') {
        arr['e_com'] = arrE360[0];  
    }
    else if (arrE360.length == 5 && arrE360[0] != '') {
        arr['e_uid'] = arrE360[0];
        arr['e_cid'] = arrE360[1];
        arr['e_gid'] = arrE360[2];
        arr['e_yid'] = arrE360[3];
        arr['e_kid'] = arrE360[4];
    }
    else if (arrE360.length == 6 && arrE360[0] != '') {
        arr['e_uid'] = arrE360[0];
        arr['e_cid'] = arrE360[1];
        arr['e_gid'] = arrE360[2];
        arr['e_yid'] = arrE360[3];
        arr['e_kid'] = arrE360[4];
        arr['e_com'] = arrE360[5];  
    }
    return arr;
}

function getBrowser() {
    var arr = new Array();
    arr["cl"] = getCl();
    arr["ds"] = getDs();
    arr["ln"] = getLn();
    arr["ck"] = getCkAfterLt();
    arr["ja"] = getJa();
    arr["fl"] = getFl();
    return arr;
}

function getPageStats() {
    var arr = getBrowser();
    arr["si"] = qsSi;
    arr["v"] = qsV;

    var refer = top.document.referrer;
    arr["su"] = refer;
    arr["lt"] = lt[lt.length - 1];
    if (lt.length > 1) {
        arr["lt2"] = lt[lt.length - 2];
    }
    return arr;
}
function execE360js(e360_u, isSameVisit) {
    if (e360_u != '' && !isSameVisit) {
        var element = document.createElement('script');
        element.type = 'text/javascript';
        element.async = true;
        element.src = ('https:' == document.location.protocol ? 'https://' : 'http://') + 'stat.tf.360.cn/search/c.js?u=' + e360_u;
        var node = document.getElementsByTagName('script')[0];
        node.parentNode.insertBefore(element, node);
		attachListener(element,'load',saveLoad);
        attachListener(element,'error',saveLoad);
		attachListener(element,"readystatechange",(function(){if(element.readyState=="complete" || element.readyState=="loaded"){saveLoad();}})); 
    }else{
        saveLoad();
    }
}
function combineArray(ls1, ls2) {
    for (key in ls2) {
        if (typeof ls1[key] == 'undefined')
            ls1[key] = ls2[key];
        else
            ls1[key + '2'] = ls2[key];
    }
    return ls1;
}



var saveLoadCount=0;
var win = window,
    doc = document,
    protocol = document.location.protocol;
    baseUrl = 'http://s.union.360.cn/s.gif';

var qsLoad = 0;
var qsClick = 2;
var qsUnload = 3;
var qsDocument = 87;

var qsSi = '6697';
var qsV = '1.0.2';
var qsPageClk = false;
var e360_u = '0';
var isSameVisit = false;   
var qsPageOpenTime = parseInt(getTime()/1000);
var lt = getLt();      

var qsPageStats = getPageStats();
execE360js(e360_u, isSameVisit);
function attachListener(dom, type, handler) {
    if (dom.addEventListener) {
        dom.addEventListener(type, handler, false);
    } else if (dom.attachEvent) {
        dom.attachEvent('on' + type, handler);
    } else {
        dom["on" + type] = handler;
    }
}

function getEventCoord(event) {
    if (event.pageX == null && event.clientX != null) {
        var target = event.target || event.srcElement,
            eventDoc = target.ownerDocument || doc,
            doc = eventDoc.documentElement,
            body = eventDoc.body;
    }
    var coord = new Array();
    coord['x'] = 'null';
    coord['y'] = 'null';

    if (event.pageX != null) {
        coord['x'] = event.pageX;
    }

    if (event.pageY != null) {
        coord['y'] = event.pageY;
    }

    return coord;
}

function findAnchor(event) {
    var i, maxdepth = 3,
        target = event.target || event.srcElement;

    var arr = new Array();
    arr['t'] = target.nodeName;

    for (i = 0; i < maxdepth && target && target.nodeName != 'A'; target = target.parentNode) {
    }
    arr['u'] = target && target.nodeName == 'A' ? target : '';
    return arr;
}

function request(url) {
    var key = 'qstats_' + getTime(),
        img = new Image;

    win[key] = img;
    img.onload = img.onerror = function () {
        win[key] = null;
        try {
            delete win[key];
        } catch (e) {
        }
    };
    img.src = url;
}

function makeRequest(ext) {
    ext["rnd"] = getRandom(10);
    var pos = 0;
    var str = "";
    for (var i in ext) {
        if (pos++) {
            str += "&";
        }
        str += i + "=" + encodeURIComponent(ext[i]);
    }
    request(baseUrl + '?' + str);
}

function saveLoad(event) {
	if(saveLoadCount==0){ 
		var arr = qsPageStats;
		setE360();
		arr = combineArray(arr,getE360());
		arr["et"] = qsLoad;
    arr['flt'] = qsPageOpenTime;
		makeRequest(arr);
	}
	saveLoadCount=saveLoadCount+1;
}

function saveUnload(event) {
    var ep = parseInt(getTime()/1000) - qsPageOpenTime;
    var arr = qsPageStats;
    arr["et"] = qsUnload;
    arr["ep"] = ep;
    arr['flt'] = qsPageOpenTime;
    makeRequest(arr);
}
function formatClkEp(arr1, arr2) {
    var str = '';
    var pos = 0;
    for (var k in arr1) {
        if (pos++) {
            str += ",";
        }
        str += k + ":" + encodeURIComponent(arr1[k]);
    }

    for (var k in arr2) {
        if (pos++) {
            str += ",";
        }
        str += k + ":" + encodeURIComponent(arr2[k]);
    }

    return str;
}
function saveClick(event) {
    if (!qsPageClk) {
        return;
    }

    var arr = findAnchor(event);
    if (arr['u'] == ''|| arr['u'].indexOf('javascript')>-1) {
        return;
    }

    var coord = getEventCoord(event);
    var ep = formatClkEp(arr, coord);

    var arr = qsPageStats;
    arr["et"] = qsClick;
    arr["ep"] = ep;
    arr['flt'] = qsPageOpenTime;
    makeRequest(arr);
}

attachListener(win, 'load', saveLoad);
attachListener(win, 'beforeunload', saveUnload);//chrome,ie
attachListener(doc, 'click', saveClick);

})();