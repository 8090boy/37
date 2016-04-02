function ca(info) {
    if (!info) {
        cookie.Del("token")
        location.href = "/"
        return
    }
    if (info.state > 1) {
        window.my37 = info
        my.init()
    } else {
        cookie.Del("token")
        location.href = "/"
    }
}
 
var my = {
    GetInfo:function(){
        jQuery.get("/api/sso/myinfo?cb=ca&token="+cookie.Get('token'),null,null)
    },
    init: function () {
        this.user = window.my37.u
        surface.on('my')
        var inputArr = document.querySelector("#my").querySelectorAll("input,textarea")
        for (var i = 0; i < inputArr.length; i++) {
            this.setDefault(inputArr[i], i)
        }
    }
    , setDefault: function (el, index) {
        if (this.user[el.name]) {
            if (el.name.toLowerCase() === 'password') {
                el.value = ""
            } else {
                el.value = this.user[el.name] || ""
            }
        } else {
            el.value = window.my37.r[el.name] || ""
        }

    }
 

    , exit: function () {
        cookie.Del('token')
        cookie.Del(my37.r.Mobile + "today")
        var url = '/api/sso/exit';
        ajax.GET(url, this.exitCall);
    }
    , exitCall: function (msg) {
        if (msg.toLowerCase() !== 'ok') return
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
        var url = '/api/37/v1/my/edit'
        var data = DataCompp("edit")
        ajax.POST(url, data, this._updateSuccess)
    }
    , _updateSuccess: function (msg) {
        if (('' + msg).toLowerCase() === 'ok') {
            alert('成功!')
            document.querySelector("#edit .button.subBtn").style.display = 'none'

        } else {
            alert('失败!')
        }
    }


}


