function ca(info) {
    if (!info) {
        cookie.Del("token");
        location.href = "/";
        return
    }
    if (info.state) {
        window.myinfo = info;
        document.addEventListener("DOMContentLoaded", main.goon, false)
    } else {
        cookie.Del("token");
        location.href = "/"
    }
}
var main = {

    goon: function () {

        var url = "/api/200/hundred?cb=main.mystart&v=" + uuid(6, 10);
        jQuery.getScript(url, null)
    },
    mystart: function (obj) {
        if (obj.s == 2) return location.href = "6.html"
        window.my37 = new Object();
        for (var a in obj) {
            my37[a] = obj[a]
        }
        var ui = document.body;
        ui.info = document.querySelector("#info");
        ui.plat = document.querySelector("#main").querySelector(".platform");
        this.ui = ui;
        this.showUi()
    },

    showUi: function () {
        document.querySelector("#info .username").innerText = myinfo.u.Alias || myinfo.u.Mobile;
        document.querySelector("#info .wechat").innerText = myinfo.u.Wechat;

        if (!my37.r) {
            return this.noHaveRela()
        }
        if (!my37.r.CurrentMonad) {
            return this.noState()
        }
        if (!my37.m) {
            return this.noState()
        }
        if (!my37.r.Status) {
            return this.waitAccpcet()
        }
        if (my37.r.Status != 1 && my37.r.Status != 9) {
            this.showFreezeInfo(my37.r.Status, "income")
            this.Audit.UpdateTask()
            return
        }


        this.Audit.UpdateTask()
        this.MonadNormal()

    },
    MonadNormal: function () {
        this.defaultCountUp(my37.r.Income, my37.r.Spending, my37.r.Loss);
        this._showStartTag();
        if (my37.r.Status == 9) {
            document.querySelector(".recommand").style.display = 'none'
            document.querySelector(".one").style.display = 'none'
            document.querySelector('.retire').style.display = 'block'
            document.querySelector('#income').style.textDecoration = 'line-through'
            return;
        }
        if (my37.todos && my37.todos.length > 0) {
            var todocount = document.querySelector("#todo");
            todocount.innerText = my37.todos.length;
            var todo = document.querySelector(".todo")
            todo.className = "todo addRedpackage"
            todo.style.display = 'block'
        } else {
            var addRedpackage = document.querySelector(".addRedpackage")
            var todo = document.querySelector(".todo")
            if (addRedpackage) {
                addRedpackage.className = "todo"
                addRedpackage.style.display = 'none'
            }
            if (todo) {
                todo.style.display = 'none'
            }
            document.querySelector("#todo").innerText = 0
        }
    },
    showDFinfo: function () {
        var msg = my37;
        var sdfi = document.querySelector(".sdfi");
        var htm = [];
        htm.push("<h4>对方信息&nbsp;&nbsp;&nbsp;&nbsp;<a onclick='main.closeAddInfo()'>&nbsp;</a></h4>");
        htm.push('<p>昵称：<b onclick="copy()">' + msg.pi.alias + "</b></p>");
        htm.push('<p>手机：<b onclick="copy()">' + msg.pi.mob + "</b></p>");
        htm.push('<p>微信：<b onclick="copy()">' + msg.pi.wechat + "</b></p>");
        htm.push("<p>空闲时间：" + msg.pi.free + "</p>");
        htm.push("<h5>帮助别人快乐自己<br/>成就别人成就自己！</h5>");
        if (msg.pri) {
            htm.push("<hr>");
            htm.push("<p>提示：请务必在第一时间内完成此次任务！如果联系不上对方，请联系以下对方推荐人催促。</p>");
            htm.push("<h4>对方推荐人信息：</h4>");
            htm.push('<p>昵称：<b onclick="copy()">' + msg.pri.alias + "</b></p>");
            htm.push('<p>手机：<b onclick="copy()">' + msg.pri.mob + "</b></p>");
            htm.push('<p>微信：<b onclick="copy()">' + msg.pri.wechat + "</b></p>");
            htm.push("<p>空闲时间：" + msg.pri.free + "</p>")
        }
        sdfi.innerHTML = htm.join("");
        sdfi.style.display = "block";
        if (my37.isMain) {
            var sdfi = document.querySelector(".sdfi");
            sdfi.firstChild.querySelector("a").style.display = "none"
        }
    },
    closeAddInfo: function (e) {
        e = window.event || event;
        var el = e.target;
        el.parentElement.parentElement.style.display = "none";
        this.goon()
    },
    AddMonad: function () {
        var url = "/api/200/v1/task/new";
        ajax.GET(url, this._addMonadPost.bind(this))
    },
    _addMonadPost: function (msg) {
        if (msg.ok) {
            cookie.Set(my37.r.Mobile + "today", true, my37.interval);
            return alert("成功!")
        }
        this._copyTo37(msg);
        this.showDFinfo()
    },
    waitAccpcet: function () {
        document.querySelector(".process").style.display = "block";
        document.querySelector(".nohave").style.display = "none";
        document.querySelector("#info .recommand").style.display = "none";
        this.showDFinfo();
        var sdfi = document.querySelector(".sdfi");
        sdfi.firstChild.querySelector("a").style.display = "none"
    },
    noState: function () {
        this.ui.plat.querySelector(".nohave").style.display = "block";
        this.ui.plat.querySelector(".activeM").style.display = "none";
        this.ui.plat.querySelector("article").style.display = "none";
        this.ui.querySelector(".taskWarp").style.display = "none";
        document.querySelector("#info .recommand").style.display = "none";
        cookie.Del(my37.r.Mobile + "today")
    },
    noHaveRela: function () {
        this.toIndex()
    },
    reloadPage: function () {
        location.reload()
    },
    todo: function () {
        if (!my37.todos) {
            return
        }
        var warpAu = document.querySelector(".warpAu");
        var titAu = document.querySelector(".titAu");
        if (warpAu && titAu) {
            warpAu.style.display = "block";
            titAu.style.display = "block";
            return
        }
        var auid = [];
        for (var i = 0; i < my37.todos.length; i++) {
            if (my37.todos[i].Id) {
                auid.push(my37.todos[i].Id)
            }
        }
        var url = "/api/200/v1/todo/list?_=" + auid.join("|");
        ajax.GET(url, this._showtodo.bind(this))
    },
    _showtodo: function (msg) {
        var htm = [];
        var title = document.createElement("h2");
        title.innerHTML = '红包收到了吗？<a title="没收到！"  onclick="main.Audit.CloseAutidsList()">&nbsp;</a>';
        title.className = "tit titAu";
        var userinfo = document.querySelector("#info");
        userinfo.appendChild(title);
        for (var i = 0; i < my37.todos.length; i++) {
            if (my37.todos[i].Id && msg[my37.todos[i].Id]) {
                var className = i % 2 ? "odd" : "even";
                var refD = msg[my37.todos[i].Id].split("|");
                var mob = refD[0]
                    , wechat = refD[1]
                    , alias = refD[2]
                    , sum = refD[3]
                    , date = refD[4]
                    , cla = refD[5]
                    , isMa = refD[6];
                var datee = date.split("+")[0];
                alias = alias ? alias : "";
                datee = datee.replace(/[T]/img, " ").replace(/[Z]/img, "");
                htm.push('<li id="au_li_' + my37.todos[i].Id + '" integral=' + sum + ' class="' + className + '" >');
                htm.push('<div class="redPackget" >');
                htm.push('<h3 class="sum" >￥&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;' + sum + "</h3>");
                var tempStr = "对方：" + cla + "&nbsp;级，";
                htm.push('<p class="class"  >' + tempStr + "</b></p>");
                htm.push('<p class="alias"  >昵称：<b  onclick="copy()" >' + alias + "</b></p>");
                htm.push('<p class="wechat"  >微信：<b  onclick="copy()" >' + wechat + "</b></p>");
                htm.push('<p class="mob"  >手机：<b  onclick="copy()" >' + mob + "</b></p>");
                htm.push("<p>" + datee + "</p>");
                htm.push('<u   onclick="main.Audit.OK(' + my37.todos[i].Id + ')" >收到了</u>');
                htm.push('<u class="not"  onclick="main.Audit.Not(' + my37.todos[i].Id + ')" >没收到</u>');
                htm.push("</div>");
                htm.push("</li>")
            }
        }
        var ul = document.createElement("ul");
        ul.className = "warp warpAu";
        ul.innerHTML = htm.join("");
        var df = document.createDocumentFragment();
        var body = document.body;
        var h = body.getBoundingClientRect().bottom - title.getBoundingClientRect().bottom;
        df.appendChild(ul)
        body.appendChild(df)
        window.auditTodosCount = 0
    },
    _: function (dataArr, key, val) {
        for (var i = 0; i < dataArr.length; i++) {
            if (dataArr[i][key] == val) {
                return dataArr[i]
            }
        }
    },
    _copyTo37: function (msg) {
        for (var o in msg) {
            my37[o] = msg[o]
        }
    },
    toIndex: function () {
        location.href = "/"
    },
    ShowRecCode: function (event) {
        var url = "/api/200/v1/my/code";
        ajax.GET(url, this._showRecommandCode.bind(this))
    },
    _showRecommandCode: function (msg) {
        if (!msg) {
            return
        }
        var dialog = document.querySelector("#dialog");
        var shrea = dialog.querySelector(".shrea");
        var input = shrea.querySelector(".input");
        if (location.port) {
            input.innerText = location.protocol + "//" + location.hostname + ":" + location.port + "/?rh=" + msg
        } else {
            input.innerText = location.protocol + "//" + location.hostname + "/?rh=" + msg
        }
        dialog.style.display = "block";
        shrea.style.display = "block";
        shrea.style.width = document.body.clientWidth + "px"
    },
    shrea_hide: function () {
        var dialog = document.querySelector("#dialog");
        var shrea = dialog.querySelector(".shrea");
        dialog.style.display = "none";
        shrea.style.display = "none"
    },
    Task: {
        Ok: function (no) {
            return this.showTaskUi(1, 0, 6)
        },
        showTaskUi: function (tag, start, end) {
            var htm = [];
            if (!my37.tasks) {
                return
            }
            for (var m = 0; m < my37.tasks.length; m++) {
                var mt = my37.tasks[m];
                if (mt.Status != 2) {
                    continue
                }
                var mtc = mt.ProposerCount;
                if (mtc >= start && mtc <= end) {
                    var tmpHtm = '<p id="tmpid_' + mt.Id + '" onclick="main.Task.exec(' + mt.Id + ',true)" >&nbsp;</p>';
                    htm.push(tmpHtm)
                }
            }
            var tit = "<div class='areaTop' ><h4>未完成任务&nbsp;" + tag + "&nbsp;列表&nbsp;&nbsp;&nbsp;" + '<a onclick="main.closeTaskList()">&nbsp;</a></h4>';
            htm.push("</div>");
            var con = htm.join("");
            var needSubmitTaskslist = tit + con;
            var htm2 = [];
            for (var k = 0; k < my37.tasks.length; k++) {
                var mt = my37.tasks[k];
                if (mt.Status != 0) {
                    continue
                }
                var mtc = mt.ProposerCount;
                if (mtc >= start && mtc <= end) {
                    var tmpHtm = '<p  onclick="main.Task.exec(' + mt.Id + ',false)" >&nbsp;</p>';
                    htm2.push(tmpHtm)
                }
            }
            tit = "<div class='areaBottom' ><h4>待对方确认的任务&nbsp;" + tag + "&nbsp;列表</h4>";
            htm2.push("</div>");
            con = htm2.join("");
            var tasksunderway = tit + con;
            var sdfi = document.querySelector(".taskList");
            sdfi.className = "sdfi taskList";
            sdfi.innerHTML = needSubmitTaskslist + tasksunderway;
            sdfi.style.display = "block"
        },
        exec: function (mesId, statu) {
            this.statu = statu;
            this.auid = mesId;
            var url = "/api/200/v1/task/find/" + mesId;
            ajax.GET(url, this._findMes.bind(this))
        },
        _findMes: function (msg) {
            var pi = msg.pi.split("|");
            var pri = msg.pri.split("|");
            var sdfi = document.querySelector(".newRedpackage");
            var con = [];
            con.push('<div class="targetInfo" >');
            con.push("<h3>￥" + pi[3] + "</h3>");
            if (this.statu) {
                con.push("<a onclick='this.parentElement.parentElement.style.display=\"none\"'>&nbsp;</a>");
                con.push('<a class="commitTask" onclick=\'main.Todo.add(' + this.auid + ")'>&nbsp;</a>");
                con.push("<h4>对方信息</h4>")
            } else {
                con.push("<a onclick='this.parentElement.parentElement.style.display=\"none\"'>&nbsp;</a>");
                con.push("<h4>已发出，等待对方确认中：</h4>")
            }
            con.push('<p>昵称：<b onclick="copy()"  >' + pi[2] + "</b></p>");
            con.push('<p>手机：<b onclick="copy()" >' + pi[0] + "</b></p>");
            con.push('<p>微信：<b onclick="copy()" >' + pi[1] + "</b></p>");
            con.push("<p>空闲时间：" + pi[4] + "</p>");
            con.push("</div>");
            con.push('<div class="targetRefInfo">');
            con.push("<h5>帮助别人快乐自己<br/>成就别人成就自己</h5>");
            if (msg.pri) {
                con.push("<hr>");
                con.push("<p>提示：请务必在第一时间内完成此次任务！如果联系不上对方，请联系以下对方推荐人催促。</p>");
                con.push("<h4>对方推荐人信息：</h4>");
                con.push('<p>昵称：<b onclick="copy()">' + pri[2] + "</b></p>");
                con.push('<p>手机：<b onclick="copy()">' + pri[0] + "</b></p>");
                con.push('<p>微信：<b onclick="copy()">' + pri[1] + "</b></p>");
                con.push("<p>空闲时间：" + pri[3] + "</p>");
                con.push("</div>")
            }
            sdfi.innerHTML = con.join("");
            sdfi.style.display = "block"
        }
    },
    closeTaskList: function () {
        document.querySelector(".sdfi.taskList").style.display = "none";
        main.goon()
    },
    Todo: {
        add: function (id, cls) {
            this.tmpid = id;
            var areaBottom = document.querySelector(".areaBottom");
            var p = document.querySelector(".areaTop #tmpid_" + this.tmpid);
            var newP = p.cloneNode(true);
            p.style.display = "none";
            var tmpClick = newP.getAttribute("onclick").replace("true", "false");
            newP.setAttribute("onclick", tmpClick);
            areaBottom.appendChild(newP);
            var url = "/api/200/v1/task/submit/" + id;
            ajax.GET(url, this.addResult.bind(this));
            main.goon()
        },
        addResult: function (msg) {
            if (msg) {
                alert(msg);
                var sdfi = document.querySelector(".newRedpackage");
                sdfi.innerHTML = "";
                sdfi.style.display = "none"
            }
        }
    },
    Audit: {
        OK: function (id) {
            window.confirm('确定收到了！', "main.Audit._ok(" + id + ")", '_not()');
        },
        _not: function () {

        },
        _ok: function (id) {
            var url = "/api/200/v1/todo/submit/" + id;
            this.id = "au_li_" + id;
            ajax.GET(url, this._auditOk.bind(this));
        },
        _auditOk: function (msg) {
            var li = document.querySelector("#" + this.id);
            var integral = parseInt(li.getAttribute("integral"));
            window.auOkCount = window.auOkCount || 0;
            window.auOkCount = integral + window.auOkCount;
            li.parentElement.removeChild(li)
            var lis = document.querySelectorAll(".warpAu li");
            if (!lis.length) {
                main.Audit.CloseAutidsList()
            }
            if (msg.influence) {
                my37.tasks = my37.tasks || [];
                my37.tasks.push(msg.task);
                this.UpdateTask();
            }
            // 审核单子数量
            window.auditTodosCount++
        },
        Not: function (id) {
            var liId = "au_li_" + id;
            document.getElementById(liId).style.display = "none"
        },
        _notTodo: function (msg) {
            switch (msg) {
                case 1:
                    alert("无权限审核!");
                    break;
                default:
                    alert("ok");
                    break
            }
        },
        UpdateTask: function () {
            var tasks = my37.tasks;
            if (!tasks || !tasks.length) return;
            var taskCount = document.querySelector(".task .one u")
            taskCount.innerText = tasks.length;
            taskCount.style.display = 'inline-block'
            return;
            /*
            var taskUi = document.querySelector(".task");
            var oneN = 0
                , twoN = 0
                , threeN = 0;
            for (var i = 0; i < tasks.length; i++) {
                if (!tasks[i]) {
                    continue
                }
                var NO = parseInt(tasks[i].ProposerCount);
                switch (NO) {
                    case 0:
                        oneN += 1;
                        break;
                    case 1:
                        oneN += 1;
                        break;
                    case 2:
                        oneN += 1;
                        break;
                    case 3:
                        twoN += 1;
                        break;
                    case 4:
                        twoN += 1;
                        break;
                    case 5:
                        threeN += 1;
                        break;
                    case 6:
                        threeN += 1;
                        break
                }
            }
            var one = taskUi.querySelector(".one u");
            oneN ? (one.innerText = oneN) : one.style.dispaly = "none";
         */
        },
        CloseAutidsList: function () {

            var todo = document.querySelector("#todo");
            var count = parseInt(todo.innerText);
            count = count - window.auditTodosCount;
            if (count > 0) {
                todo.innerText = count;
            } else {
                todo.innerText = 0;
                var RedPackageList = document.querySelector(".addRedpackage");
                RedPackageList.className = 'todo'
                RedPackageList.style.display = 'none'
            }

            document.querySelector(".warpAu").style.display = "none";
            document.querySelector(".titAu").style.display = "none";
            window.auOkCount = window.auOkCount || 0;
            if (!window.auOkCount) return

            var tmpCount = window.auOkCount + my37.r.Income;
            main.CountUpdate("income", tmpCount);
            main.goon()
        }
    },
    defaultCountUp: function (i, s, l) {
        var options = {
            useEasing: true,
            useGrouping: true,
            separator: ",",
            decimal: ".",
            prefix: "",
            suffix: ".0"
        };

        window.income = new CountUp("income", 0, i, 0, 6, options);
        window.income.start()

        window.spending = new CountUp("spending", 0, s, 0, 2, options);
        window.spending.start();
        window.loss = new CountUp("loss", 0, l, 0, 1, options);
        window.loss.start()
    },
    showFreezeInfo: function (no, id) {
        var el = document.querySelector("#" + id);
        el.className = "freeze";
        if (no === 2) {
            var ref = my37.m.UnFreeze.split("T");
            var day = ref[0];
            var hourRef = ref[1].split(":");
            var hour = hourRef[0] + ":" + hourRef[1];
            el.innerHTML = "未按时出单已冻结，请在：<br>" + day + " " + hour + "之前出完&nbsp;" + my37.m.UnfreezePeriodCount + "&nbsp;单解冻；<br>逾期未解冻，该帐号将被注销！"
        }
        if (no === 4) {
            el.innerHTML = "任务数超限额被冻结，<br>请完成任务！否则将无限期冻结！"
        }
        document.querySelector(".recommand").style.display = 'none'
    },
    CountUpdate: function (el, count) {
        if (el.nodeName) {
            window[el.id].update(count);
            el.setAttribute("ref", count)
        } else {
            if (typeof (el) === "string") {
                window[el].update(count)
            }
            document.querySelector("#" + el).setAttribute("ref", count)
        }
    },
    _showStartTag: function () {
        var start = "";
        switch (my37.m.Class) {
            case 0:
                start = 0;
                break;
            case 1:
                start = 52;
                break;
            case 2:
                start = 105;
                break;
            case 3:
                start = 160;
                break;
            case 4:
                start = 210;
                break;
            case 5:
                start = 266;
                break;
            case 6:
                start = 318;
                break;
            default:
                start = 370;
                break
        }
        start = start.toFixed(0);
        document.querySelector("#main article h1").style.backgroundPosition = "50% -" + start + "px"
    },
    fullScreen: function () {
        launchFullscreen(document.documentElement)
    },
    normalScreen: function () {
        exitFullScreen()
    }
};
var my = {
    init: function () {
        window.zj = !window.zj;
        if (!window.zj) {
            surface.on("main");
            return
        }
        surface.on("my");
        var inputArr = document.querySelector("#my").querySelectorAll("input,textarea");
        for (var i = 0; i < inputArr.length; i++) {
            this.setDefault(inputArr[i], i)
        }
    },
    setDefault: function (el, index) {
        if (myinfo.u[el.name]) {
            if (el.name.toLowerCase() === "password") {
                el.value = ""
            } else {
                el.value = myinfo.u[el.name] || ""
            }
        } else {
            el.value = my37.r[el.name] || ""
        }
    },
    friendster: function (e) {
        e = window.event || event;
        var el = e.target || e.srcElement;
        var myF = document.querySelector("#myFriendster");
        var isShow = el.getAttribute("ok");
        if (isShow == 1) {
            el.setAttribute("ok", 0);
            myF.style.display = "none"
        } else {
            el.setAttribute("ok", 1);
            myF.style.display = "block"
        }
        var stat = myF.getAttribute("ref");
        if (stat) {
            return
        }
        var url = "/api/200/v1/my/friendster";
        ajax.GET(url, this.friendsterOk.bind(this))
    },
    friendsterOk: function (msg) {
        var first = msg.f ? msg.f.split("|") : "";
        var second = msg.s ? msg.s.split("|") : "";
        if (!first.length) {
            return alert("您的朋友圈0人，\n赶紧发您的推荐连接给别人吧！")
        }
        var liByData = function (arr, bb) {
            var tmplis = [];
            for (var i = 0; i < arr.length; i++) {
                if (arr[i]) {
                    var ref = arr[i].split("-");
                    tmplis.push("<div><b>" + ref[1] + "</b></div>");
                    if (!bb) {
                        continue
                    }
                    for (var k = 0; k < bb.length; k++) {
                        if (bb[k]) {
                            var refBb = bb[k].split("-");
                            if (refBb[0] === ref[0]) {
                                tmplis.push("<span><b>" + refBb[1] + "</b></span>")
                            }
                        }
                    }
                }
            }
            return tmplis.join("")
        }
            ;
        var myF = document.querySelector("#myFriendster");
        myF.setAttribute("ref", true);
        myF.innerHTML = liByData(first, second);
        var divs = myF.querySelectorAll("div,span");
        for (var i = 0; i < divs.length; i++) {
            divs[i].onclick = this.ShowUserInfoByMob.bind(this)
        }
    },
    ShowUserInfoByMob: function (event) {
        event = window.event || event;
        var el = event.target || event.srcElement;
        var lastC = el.lastChild;
        var myF = document.querySelector("#myFriendster");
        var pps = myF.querySelectorAll("p");
        for (var i = 0; i < pps.length; i++) {
            pps[i].style.display = "none"
        }
        if (lastC.nodeName.toLowerCase() === "p") {
            lastC.style.display = "block";
            return
        }
        var mob = el.innerText;
        var url = "/api/200/v1/my/relation/" + mob;
        ajax.GET(url, this._showUserInfo.bind(this, el))
    },
    _showUserInfo: function (el, msg) {
        var pp = document.createElement("p");
        pp.innerHTML = "手机：" + msg.Mobile + "<br>微信：" + msg.Wechat + "<br>昵称：" + msg.Alias;
        el.parentElement.appendChild(pp)
    },
    exit: function () {
        cookie.Del("token");
        cookie.Del(my37.r.Mobile + "today");
        var url = "/api/sso/exit";
        ajax.GET(url, this.exitCall)
    },
    exitCall: function (msg) {
        if (msg.toLowerCase() !== "ok") {
            return
        }
        location.href = "/"
    },
    edit: function () {
        var plat = document.querySelector("#my .platform");
        if (plat.style.display.toLowerCase() == "block") {
            plat.style.display = "none"
        } else {
            plat.style.display = "block"
        }
    },
    update: function () {
        var url = "/api/200/v1/my/edit";
        var data = DataCompp("edit");
        ajax.POST(url, data, this._updateSuccess)
    },
    _updateSuccess: function (msg) {
        if (("" + msg).toLowerCase() === "ok") {
            alert("成功!");
            document.querySelector("#edit .button.subBtn").style.display = "none"
        } else {
            alert("失败!")
        }
    }
};
