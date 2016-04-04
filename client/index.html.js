function cb(obj) {
  
      if(!obj) location.href = "/api/sso/?redirectUrl=" + location.href
    if (obj.state>1) {
        window.obj = obj
        window.user = obj.u
        var index = document.querySelector("#index");
        index.querySelector(".username").innerText = "" + user.Mobile
    } else {
        cookie.Del("token")
        var currentUrl = location.href
        location.href = "/api/sso/?redirectUrl=" + currentUrl
    }
    
    
}

function start() {
    var recomH = location.search.split("rh=")[1]
    var recomI = location.search.split("ri=")[1]
    var ref = location.search.split("token=")[1]
    
    if (recomI) {
        if (/^[\d|\w]{6,10}$/i.test(recomI)) {
            surface.on("in")
            document.querySelector('input[name="referrerId"]').value = recomI
            addListener(innreg,'click',inn.reg,bind(inn,37))
        }
        return
    }
    
     if (recomH) {
        if (/^[\d|\w]{6,10}$/i.test(recomH)) {
            surface.on("in")
            document.querySelector('input[name="referrerId"]').value = recomH
            addListener(innreg,'click',inn.reg.bind(inn,200))
        }
        return
    }
    
        var token, overdue;
        if (ref) {
            ref = ref.replace('&','')
            token = ref.split('overdue=')[0]
            overdue = ref.split('overdue=')[1]
        }
        if (!token) token = cookie.Get('token')
        if (token) {
            if (overdue) cookie.Set('token', token)
            surface.on("index")
            var userCb = document.createElement('script')
            userCb.src = '/api/sso/myinfo?cb=cb&token=' + token
            document.body.appendChild(userCb)
        } else {
            cookie.Del("token")
            var currentUrl = location.href
            location.href = "/api/sso/?redirectUrl=" + currentUrl
        }
 

}
window.onload = start

var inn = {
    join: function (paltform) {
      // if (!window.obj.r && !window.obj.s) return alert('非团队用户不能玩啥！')
   
      switch (paltform) {
          case 37:
              cookie.Set('paltform',37)
              location.href = '/37.html'
          break;
              case 24:
              cookie.Set('paltform',200)
              location.href = '/200.html'
              break;
       
      }
    }
    , reg: function (tag) {
        var url ;
        switch (tag) {
            case 200:
                url = '/api/200/signin'
                break;
            case 37:
                url = '/api/37/signin'
                break;
        }
      
        // 校验
        if (!this.validate("post")) return;
        var data = this._dataComp("post")
        ajax.POST(url, data, this._addSuccess)
    }
    , validate: function (id) {
        var inputs = document.querySelectorAll("#" + id + " input,#" + id + " textarea")
        var status = 0
        for (var n = 0; n < inputs.length; n++) {
            var input = inputs[n];
            var txt = input.value;
            input.style.border = '1px solid white'
            input.style.borderBottom = '1px solid #EFC4A6'
            switch (input.name.toLowerCase()) {
                case 'mobile':
                    if (/^[1|0][3|7|5|8|9]\d{8,9}$/i.test(txt)) {
                        status++
                    } else {
                        input.style.border = '1px solid red'
                        //   alert('手机号不正确！')
                    }
                    break;
                case 'wechat':
                    if (/^[a-zA-Z\d_]{6,}$/i.test(txt)) {
                        status++
                    } else {
                        input.style.border = '1px solid red'
                        //   alert('微信号格式错误，或长度小于6位！')
                    }
                    break;
                case 'password':
                    if (/^[a-zA-Z\d_]{6,}$/i.test(txt)) {
                        status++
                    } else {
                        input.style.border = '1px solid red'
                        //   alert('密码错误，不能用特殊字符并且不小于6位！')
                    }
                    break;
            }
        }
        if (status < 3) {
            return false
        }
        return true
    }
    , _addSuccess: function (msg) {
        if (('' + msg).toLowerCase() === 'ok') {
            alert('成功。')
            location.href = '/'
            return
        }

        if (msg === 4 || msg === 5 || msg === 6 || msg === 7) {
            alert('推荐链接不正确或已使用或过期,\n请获取新的推荐链接！')
            return
        }
        if (msg === 8 || msg === 9 || msg === 10) {
            alert('推荐人不存在或状态不正常')
            return
        }
        alert('系统失败，请重试！')


    }
    , _dataComp: function (id) {
        var inputs = document.querySelectorAll("#" + id + " input,#" + id + " textarea");
        var arg = []
        for (var n = 0; n < inputs.length; n++) {
            arg.push(inputs[n].name + '=' + inputs[n].value)
        }

        return arg.join("&")
    }
    , exit: function () {
        var url = '/api/sso/exit?' + (Math.random() * 10).toFixed(2)
        ajax.GET(url, this.exitCall)
    }
    , exitCall: function (msg) {
        cookie.Del('token')
        if (msg.toLowerCase() !== 'ok') return
        location.href = '/'
    }

}

var section = {
    sectionArr: {}
    , shift: function (showWho) {
        for (var se in this.sectionArr) {
            var section = document.querySelector("#" + this.sectionArr[se])
            section.style.display = "none"
        }
        if (!this.sectionArr[showWho]) return
        var section_show = document.querySelector("#" + showWho)
        section_show.style.display = "block"
    }


}
 