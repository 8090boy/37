function ca(my) {
    if (my.state) {
        window.my37 = my
        main.goon.call(main,null)
    } else {
          cookie.Del("token")
        currentUrl = location.href
        location.href = "/api/sso/?redirectUrl=" + currentUrl
    }
}

var main = {
     goon :function(){
        var ref = Math.random() * 10000
        var url = '/api/'+ paltform +'/interaction?cb=main.init&v=' + ref.toFixed(0)
        jQuery.get(url,null)
    },
    init: function (obj) {
        
        for (var b in obj) {
            my37[b] = obj[b]
        }
        
        for (var a in my37) {
            this[a] = my37[a]
        }
 
        var ui = document.body
        ui.info = document.querySelector("#info")
        ui.plat = document.querySelector("#main").querySelector('.platform')
        this.ui = ui 
        this.showTodo()
        this.showUi()
        this.uTagStas()
    }
    , uTagStas: function () {
        var uus = document.querySelectorAll("u")
        for (var i = 0; i < uus.length; i++) {
            if (i == 1) continue
            var u = uus[i].innerHTML
            if (!u || u == 0) {
                uus[i].style.display = 'none'
            } else {
                uus[i].style.display = 'inline-block'
            }
        }
    }
    , showUi: function () {
        document.querySelector("#info .username").innerHTML = my37.u.Alias + "&nbsp;&nbsp;&nbsp;" + my37.u.Mobile
        document.querySelector("#info .wechat").innerText = my37.u.Wechat
        document.querySelector("#income").innerText = my37.income
    }
    ,showTodo:function(){
         var htm = [];
        var au,ii=0,str='';
        
        for (au in this.todos) {
            if(!this.todos[au]) continue
                var className = ii % 2 ? 'odd' : 'even'
                ii++
                var refD = this.todos[au].split('|')
                var mob=refD[0],wechat=refD[1],alias=refD[2],sum=refD[3],date=refD[4]                
                var datee = date.split("+")[0]
                date = datee.replace(/[T]/img, ' ').replace(/[Z]/img, '')
                //
          htm.push('<li id=\"li'+au+'\" >')
          htm.push('<input type="checkbox">')
          htm.push('<i>'+mob+'</i>')
          htm.push('<b>'+wechat+'</b>')
          htm.push('<u style=\"display: inline-block;\">'+alias+'</u>')
          htm.push('<p>'+date+'</p>')
          htm.push('<em>￥'+sum+'</em>')
          htm.push('<a onclick=\"main.Audit.Ok('+au+')\">确定</a>')
          htm.push('</li>')
        }
        
        var ul = document.querySelector('.todo')
        ul.innerHTML = htm.join("")

      
    }

    , MonadNormal: function () {
        this.defaultCountUp(this.r.Income, this.r.Spending, this.r.Loss)
        this._showStartTag()
        document.querySelector('#todo').innerText = this.a ? this.a.length : 0
        this.showCreateMonadInfo() // show monad
       
        this.Audit.UpdateTask() // show task
    }
//显示出单信息
    , showCreateMonadInfo: function () {
        var generate = document.querySelector('#generate')
        var uTip = generate.querySelector('u')
        var Tit = generate.querySelector('span')
        if (my37.pi) { //没有被审核的
            removeListener(generate, "click", this.AddMonad)
            addListener(generate, "click", this.showDFinfo.bind(this))
            Tit.innerText = '等待中'
            uTip.style.display = 'none'
        } else { // 可以出单
            removeListener(generate, "click", this.showDFinfo)
            addListener(generate, "click", this.AddMonad.bind(this))
            uTip.innerText = 1
            Tit.innerText = "出单"
        }

    }
    , showDFinfo: function () {

        var msg = my37
        var sdfi = document.querySelector('.sdfi')
        var htm = [];
        htm.push('<h4>对方信息&nbsp;&nbsp;&nbsp;&nbsp;<a onclick=\'main.closeAddInfo()\'>[关闭]</a></h4>')
        htm.push('<p>昵称：<b>'+msg.pi.alias+'</b></p>')
        htm.push('<p>手机：<b>'+msg.pi.mob+'</b></p>')
        htm.push('<p>微信：<b>'+msg.pi.wechat+'</b></p>')
        htm.push('<p>空闲时间：'+msg.pi.free+'</p>')
        //
        htm.push('<h5>帮助别人快乐自己<br/>成就别人成就自己！</h5>')
        if (msg.pri) {
            htm.push('<hr>')
            htm.push('<p>提示：请务必在24小时内完成此次任务！如果联系不上对方，请联系以下对方推荐人催促。</p>')
            htm.push('<h4>对方推荐人信息：</h4>')
            htm.push('<p>昵称：<b>'+msg.pri.alias+'</b></p>')
            htm.push('<p>手机：<b>'+msg.pri.mob+'</b></p>')
            htm.push('<p>微信：<b>'+msg.pri.wechat+'</b></p>')
            htm.push('<p>空闲时间：'+msg.pri.free+'</p>')
        }
        sdfi.innerHTML = htm.join("")
        sdfi.style.display = 'block'
        if (my37.isMain) {
            var sdfi = document.querySelector('.sdfi')
            sdfi.firstChild.querySelector('a').style.display = 'none'
        }
    }
    , closeAddInfo: function (e) {
        e = window.event || event
        var el = e.target
        el.parentElement.parentElement.style.display = "none"
        this.reloadPage()

    }
    , AddMonad: function () {
        var url = '/api/'+ paltform +'/v1/task/new'
        ajax.GET(url, this._addMonadPost.bind(this))
    }
    , _addMonadPost: function (msg) {
     
        
        this._copyTo37(msg)

        var generate = document.querySelector('#generate')
        removeListener(generate, "click", this.AddMonad)
        addListener(generate, "click", this.showDFinfo.bind(this))
        if (msg.pi) { //没有被审核的
            generate.querySelector('span').innerText = '等待中'
            generate.querySelector('u').style.display = 'none'

        }
        this.showDFinfo()
    }
//对方未确认主单
    , waitAccpcet: function () {
        document.querySelector('.process').style.display = 'block'
        document.querySelector('.nohave').style.display = 'none'
        this.showDFinfo()
        var sdfi = document.querySelector('.sdfi')
        sdfi.firstChild.querySelector('a').style.display = 'none'
    }
//没有主单
    , noState: function () {
        this.ui.plat.querySelector('.nohave').style.display = 'block'
        this.ui.plat.querySelector('.activeM').style.display = 'none'
        this.ui.plat.querySelector('article').style.display = 'none'
        this.ui.querySelector('.taskWarp').style.display = 'none'

    }
//非关系户
    , noHaveRela: function () {
        this.toIndex()
    }
    , reloadPage: function () {
        location.reload()
    }
//确认
    , todo: function () {
        if (!this.a) return

        var warpAu = document.querySelector('.warpAu')
        var titAu = document.querySelector('.titAu')
        if (warpAu && titAu) {
            warpAu.style.display = 'block'
            titAu.style.display = 'block'
            return
        }


        var sso = []
        var mon = []
        var rela = []
        for (var i = 0; i < this.a.length; i++) {
            if (this.a[i].ProposerSso)
                sso.push(this.a[i].ProposerSso)
            if (this.a[i].ProposerMonadId)
                mon.push(this.a[i].ProposerMonadId)
            if (this.a[i].ProposerRelationalId)
                rela.push(this.a[i].ProposerRelationalId)

        }

        var url = '/api/'+ paltform +'/v1/todo/list?sso=' + sso.join('|') + '&mon=' + mon.join('|') + '&rela=' + rela.join('|')
        ajax.GET(url, this._showtodo.bind(this))
    }
    , _showtodo:function(msg) {
        var htm = [];
        var title = document.createElement('h2')
        title.innerHTML = '待审核列表&nbsp;&nbsp;&nbsp;<a onclick="main.Audit.CloseAutidsList()">[关闭]</a>'
        title.className = 'tit titAu'
        var userinfo = document.querySelector('#info')
        userinfo.appendChild(title)
        var str = ''
        for (var i = 0; i < this.a.length; i++) {
            var clas = '20'
            var integral = 20
            switch (msg.m[i].Class + 1) {
                case 2:
                    clas = '20'
                    break;
                case 3:
                    clas = '20'
                    break;
                case 4:
                    clas = '50'
                    integral = 50
                    break;
                case 5:
                    clas = '50'
                    integral = 50
                    break;
                case 6:
                    clas = '100'
                    integral = 100
                    break;
                case 7:
                    clas = '100'
                    integral = 100
                    break;
            }
            var s = this._(msg.s, 'Id', this.a[i]['ProposerSso'])
            if (s) {
                var className = i % 2 ? 'odd' : 'even'
                var datee = this.a[i].Create.replace(/[T]/img, ' ').replace(/[Z]/img, '')
                str = '<li id="li_' + this.a[i].Id + '" integral=' + integral + ' class="' + className + '" ><i>' + s.Alias + '<br>' + s.Wechat + '<br>'
                + + s.Mobile + '</i><b>' + datee + '<br>' + clas + '</b><u onclick="main.Audit.Ok(' + this.a[i].Id + ')" >确认</u></li>';
            }
            htm.push(str)
        }
        var ul = document.createElement('ul')
        ul.className = 'warp warpAu'
        ul.innerHTML = htm.join("")
        var df = document.createDocumentFragment()
        //
        var body = document.body
        var h = body.getBoundingClientRect().bottom - title.getBoundingClientRect().bottom
        ul.style.height = h + 'px'
        df.appendChild(ul)
        body.appendChild(df)
    }
    , _:function(dataArr, key, val) {
        for (var i = 0; i < dataArr.length; i++) {
            if (dataArr[i][key] == val) {
                return dataArr[i]
            }
        }
    }
    , _copyTo37:function(msg) {
        for (var o in msg) {
            my37[o] = msg[o]
        }
    }

    , toIndex: function () {
        location.href = '/'
    }
    , ShowRecCode: function (event) {
        var url = '/api/'+ paltform +'/v1/my/code'
        ajax.GET(url, this._showRecommandCode.bind(this));
    }
    , _showRecommandCode: function (msg) {
        if (!msg) return
        var shrea = document.querySelector('#dialog .shrea')
        shrea.querySelector('input').value = location.origin + "/?r=" + msg
        shrea.style.display = 'block'
        shrea.style.width = document.body.clientWidth + 'px'
    }
    , shrea_hide: function () {
        var shrea = document.querySelector('#dialog .shrea')
        shrea.style.display = 'none'
    }

  
    , closeTaskList: function () {
        document.querySelector(".sdfi.taskList").style.display = 'none'
        main.reloadPage()

    }
    , Todo: {
        add: function (id, cls) {
            this.tmpid = id
            var areaBottom = document.querySelector('.areaBottom')
            var p = document.querySelector('.areaTop #tmpid_' + this.tmpid)
            var newP = p.cloneNode(true)
            p.style.display = 'none'
            var tmpClick = newP.getAttribute('onclick').replace('true', 'false')
            newP.setAttribute('onclick', tmpClick)
            areaBottom.appendChild(newP)
            var url = '/api/'+ paltform +'/v1/task/submit/' + id
            ajax.GET(url, this.addResult.bind(this));

        }
        , addResult: function (msg) {
            if (msg) {
                alert(msg)
                var sdfi = document.querySelectorAll('.sdfi')[1]
                sdfi.innerHTML = ''
                sdfi.style.display = 'none'
      
            }
        }
    }
    , Audit: {
        Ok: function (id) {
            
            var url = '/api/'+ paltform +'/v1/todo/submit/' + id
            this.id = "li" + id
             var ckb = document.querySelector("#"+this.id).firstChild
             if (ckb.checked){
            ajax.GET(url, this._auditOk.bind(this))
                 
             }
        }
        , _auditOk: function (msg) {
            document.querySelector("#"+this.id).style.display='none'
             alert(msg)
             location.reload()            
        }
        , UpdateTask:function() {
            var tasks = my37.tasks
            if (!tasks || !tasks.length) return
            var taskUi = document.querySelectorAll(".task")[1]
            var one = taskUi.querySelector(".one u")
            var two = taskUi.querySelector(".two u")
            var three = taskUi.querySelector(".three u")
            var oneN = 0, twoN = 0, threeN = 0;
            for (var i = 0; i < tasks.length; i++) {
                var NO = parseInt(tasks[i].MClass)
                switch (NO) {
                    case 1:
                        oneN += 1
                        break;
                    case 2:
                        oneN += 1
                        break;
                    case 3:
                        twoN += 1
                        break;
                    case 4:
                        twoN += 1
                        break;
                    case 5:
                        threeN += 1
                        break;
                    case 6:
                        threeN += 1
                        break;
                }
            }
            oneN ? (one.innerText = oneN) : one.style.dispaly = 'none'
            twoN ? (two.innerText = twoN) : two.style.dispaly = 'none'
            threeN ? (three.innerText = threeN) : three.style.dispaly = 'none'

        }
        , CloseAutidsList: function () {
            document.querySelector('.warpAu').style.display = 'none'
            document.querySelector('.titAu').style.display = 'none'
            window.auOkCount = window.auOkCount || 0
            if (!window.auOkCount) return
            var tmpCount = window.auOkCount + my37.r.Income
            main.CountUpdate('income', tmpCount)
            main.reloadPage()
        }


    }
     
//
//
//
//
//
    , defaultCountUp: function (i, s, l) {
        var options = {
            useEasing: true,
            useGrouping: true,
            separator: ',',
            decimal: '.',
            prefix: '',
            suffix: '.0'
        };

        window.income = new CountUp("income", 0, i, 0, 6, options);
        window.spending = new CountUp("spending", 0, s, 0, 2, options);
        window.loss = new CountUp("loss", 0, l, 0, 1, options);
        window.income.start()
        window.spending.start()
        window.loss.start()
    }
    , CountUpdate: function (el, count) {
        if (el.nodeName) {
            window[el.id].update(count)
            el.setAttribute('ref', count)
        } else {
            if (typeof (el) === "string") window[el].update(count)
            document.querySelector('#' + el).setAttribute('ref', count)
        }

    }
    , _showStartTag: function () {
        var start = '-' + ((my37.m.Class - 1) * 2.3).toFixed(2) + 'em'
        if (start)
            document.querySelector('#main article h1').style.backgroundPositionY = start

    }
}




var my = {
    init: function () {

        window.zj = !window.zj
        if (!window.zj) {
            surface.on('main')
            return
        }

        this.user = window.my37.u
        surface.on('my')
        var inputArr = document.querySelector("#my").querySelectorAll("input,textarea")

        for (var i = 0; i < inputArr.length; i++) {
            this.setDefault(inputArr[i], i)
        }
    }
    , friendster: function () {
        var url = '/api/'+ paltform +'/v1/my/friendster'
        ajax.GET(url, this.friendsterOk.bind(this))
    }
    , friendsterOk:function(msg) {
        var first = msg.f ? msg.f.split('|') : ''
        var second = msg.s ? msg.s.split('|') : ''
        if (!first.length) {
            return alert("目前没有推荐人员。")
        }
        var liByData = function (arr, bb) {
            var tmplis = []
            for (var i = 0; i < arr.length; i++) {
                if (arr[i]) {
                    var ref = arr[i].split('-')
                    tmplis.push('<div><b>'+ref[1]+'</b></div>')
                    if (!bb) continue
                    for (var k = 0; k < bb.length; k++) {
                        if (bb[k]) {
                            var refBb = bb[k].split('-')
                            if (refBb[0] === ref[0]) {
                                tmplis.push('<span><b>'+refBb[1]+'</b></span>')

                            }
                        }
                    }
                }
            }

            return tmplis.join('')
        }
        var myF = document.querySelector("#myFriendster");
        myF.innerHTML = liByData(first, second)
        var divs = myF.querySelectorAll("div,span")

        for (var i = 0; i < divs.length; i++) {
            divs[i].onclick = this.ShowUserInfoByMob.bind(this)
        }

    }
    , ShowUserInfoByMob: function (event) {
        event = window.event || event
        var el = event.target || event.srcElement
        var lastC = el.lastChild
        var myF = document.querySelector("#myFriendster");
        var pps = myF.querySelectorAll("p")
        for (var i = 0; i < pps.length; i++) {
            pps[i].style.display = 'none'
        }
        if (lastC.nodeName.toLowerCase() === 'p') {
            lastC.style.display = 'block'
            return
        }
        var mob = el.innerText
        var url = '/api/'+ paltform +'/v1/my/relation/' + mob
        ajax.GET(url, this._showUserInfo.bind(this, el))

    }
    , _showUserInfo: function (el, msg) {
        var pp = document.createElement("p")
        pp.innerHTML = msg.Mobile+","+msg.Wechat+","+msg.Alias;
        el.parentElement.appendChild(pp)
    }
    , setDefault: function (el, index) {
        if (this.user[el.name]) {
            el.value = this.user[el.name] || ""
        } else {
            if(!my37.r)    return
          
            el.value = window.my37.r[el.name] || ""

        }

    }
    , exit: function() {
        var url = '/api/sso/exit';
        ajax.GET(url, this.exitCall);
    }
    , exitCall: function(msg) {
        if (msg.toLowerCase() !== 'ok') return
        cookie.Del('token')
        location.href = '/'
    }

    , edit: function () {
        var plat = document.querySelector('#my .platform')
        if (plat.style.display.toLowerCase() == 'block') {
            plat.style.display = 'none'
        } else {
            plat.style.display = 'block'
        }
    }
    , update: function () {
        var url = '/api/'+ paltform +'/v1/my/edit'
        var data = DataComp("edit")
        ajax.POST(url, data, this._updateSuccess)
    }
    , _updateSuccess: function (msg) {
        if (('' + msg).toLowerCase() === 'ok') {
            alert('成功!')
        } else {
            alert('失败!')
        }
    }


}

function DataComp(id, type) {
    if (!type) {
        var inputs = document.querySelectorAll("#" + id + " input,#" + id + " textarea");
        var arg = [];
        for (var n = 0; n < inputs.length; n++) {
            arg.push(inputs[n].name.toLowerCase() + '=' + inputs[n].value)
        }
        return arg.join("&")
    }
    if (type === 'json') {

        var inputs = document.querySelectorAll("#" + id + " input,#" + id + " textarea");
        var arg = [];
        for (var n = 0; n < inputs.length; n++) {
            arg.push('\"' + inputs[n].name.toLowerCase() + '\":\"' + inputs[n].value + '\"')
        }
        return '{' + arg.join(",") + '}'
    }

}
