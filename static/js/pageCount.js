function showPages(name) { //初始化属性
    this.name = name;      //对象名称
    this.page = 1;         //当前页数
    this.pageCount = 1;    //总页数
    this.argName = 'page'; //参数名
    this.showTimes = 1;    //打印次数
}

showPages.prototype.getPage = function () { //丛url获得当前页数,如果变量重复只获取最后一个
    var args = location.search;
    var reg = new RegExp('[\?&]?' + this.argName + '=([^&]*)[&$]?', 'gi');
    var chk = args.match(reg);
    if (chk == null) {
        this.page = 1;
    } else {
        this.page = RegExp.$1;
    }
}
showPages.prototype.checkPages = function () { //进行当前页数和总页数的验证
    if (isNaN(parseInt(this.page))) this.page = 1;
    if (isNaN(parseInt(this.pageCount))) this.pageCount = 1;
    if (this.page < 1) this.page = 1;
    if (this.pageCount < 1) this.pageCount = 1;
    if (this.page > this.pageCount) this.page = this.pageCount;
    this.page = parseInt(this.page);
    this.pageCount = parseInt(this.pageCount);
}
showPages.prototype.createHtml = function (mode) { //生成html代码

    var strHtml = '', prevPage = this.page - 1, nextPage = this.page + 1;
    if (mode == '' || typeof (mode) == 'undefined') mode = 1;
    switch (mode) {
        case 1: //模式1 (3页缩略,首页,前页,后页,尾页)


            
            //strHtml += '<span class="count"> 共' + (this.page * this.pageCount) + '页 每页' + this.showPage + '条 页次: ' + this.page + '/' + this.pageCount + '</span>';
            strHtml += '<B style="color:#6a6a6a;">每页&nbsp;' + this.showPage + '&nbsp;条&nbsp;&nbsp;共&nbsp;' + (this.page * this.pageCount) + '&nbsp;页</B>&nbsp;&nbsp;';
            
            if (prevPage < 1) {
                strHtml += '<a title="上一页" style="" > <B>&lt;上一页</B> </a>';
            } else {
                strHtml += '<a title="上一页" style="" href="javascript:' + this.name + '.toPage(' + prevPage + ');"><B>&lt;上一页</B></a>';
            }
            if (this.page % 3 == 0) {
                var startPage = this.page - 2;
            } else {
                var startPage = this.page - this.page % 3 + 1;
            }
            if (startPage > 3) strHtml += '<a  href="javascript:' + this.name + '.toPage(' + (startPage - 1) + ');">...</a>';

            for (var i = startPage; i < startPage + 3; i++) {
                if (i > this.pageCount) break;
                if (i == this.page) {
                    strHtml += '<a title="页 ' + i + '" style="text-decoration: underline;color: red;font-weight: 600;" >' + i + '</a>';
                } else {
                    strHtml += '<a  title="页 ' + i + '"  href="javascript:' + this.name + '.toPage(' + i + ');">' + i + '</a>';
                }
            }


            var countNum = this.pageCount;
            if (this.pageCount >= startPage + 5) strHtml += '<a href="javascript:' + this.name + '.toPage(' + (startPage + 3) + ');">...</a>';
            if (nextPage > this.pageCount) {
                strHtml += '<a title="下一页"  style=""> <B>下一页&gt;</B> </a>&nbsp;&nbsp;&nbsp;&nbsp;<B style="color:6a6a6a;">页次:&nbsp; ' + this.page + '&nbsp;/&nbsp;' + this.pageCount + '</B>';
                //'<input id="PageGoTo" type="text" style="color:white; width:28px;" /><a href=""><b style="color:white;" >GO</b></a>';
            } else {
                strHtml += '<a title="下一页"  style="" href="javascript:' + this.name + '.toPage(' + nextPage + ');"><B>下一页&gt;</B></a>' +
                '&nbsp;&nbsp;<B style="color:#6a6a6a;">页次:&nbsp; ' + this.page + '&nbsp;/&nbsp;' + this.pageCount; //&nbsp;&nbsp;去</B><input id="PageGoTo" type="text" style="color:#6a6a6a;width:28px;" /><a href=""><b style="color:white;" >GO</b></a>';
            }

            break;
    }
    return strHtml;
}

showPages.prototype.createUrl = function (page) { //生成页面跳转url
    if (isNaN(parseInt(page))) page = 1;
    if (page < 1) page = 1;
    if (page > this.pageCount) page = this.pageCount;
    var url = location.protocol + '//' + location.host + location.pathname;

    var args = location.search;
    var reg = new RegExp('([\?&]?)' + this.argName + '=[^&]*[&$]?', 'gi');
    args = args.replace(reg, '$1');
    if (args == '' || args == null) {
        args += '?' + this.argName + '=' + page;
    } else if (args.substr(args.length - 1, 1) == '?' || args.substr(args.length - 1, 1) == '&') {
        args += this.argName + '=' + page;
    } else {
        args += '&' + this.argName + '=' + page;
    }

    return url + args + LocalPage;

}
showPages.prototype.toPage = function (page) { //页面跳转
    var turnTo = 1;
    if (typeof (page) == 'object') {
        turnTo = page.options[page.selectedIndex].value;
    } else {
        turnTo = page;
    }

    self.location.href = this.createUrl(turnTo);
}
showPages.prototype.printHtml = function (mode) { //显示html代码
    this.getPage();
    this.checkPages();
    this.showTimes += 1;
    document.write('<div id="pages_' + this.name + '_' + this.showTimes + '" class="page"></div>');
    document.getElementById('pages_' + this.name + '_' + this.showTimes).innerHTML = this.createHtml(mode);

}
showPages.prototype.formatInputPage = function (e) { //限定输入页数格式
    var ie = navigator.appName == "Microsoft Internet Explorer" ? true : false;
    if (!ie) var key = e.which;
    else var key = event.keyCode;
    if (key == 8 || key == 46 || (key >= 48 && key <= 57)) return true;
    return false;
}
/////////////////////////////////////////////////////////////////////////

var pg = new showPages('pg');
pg.pageCount = PageCount;  // 定义总页数(必要)
pg.showPage = showPage;  // 定义每页显示多少条
//pg.argName = 'p';  // 定义参数名(可选,默认为page)

document.write('  ');
pg.printHtml(1);