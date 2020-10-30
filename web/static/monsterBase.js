// 如果是https,要换成wss://
const wsurl = "ws://116.62.78.44/ws"
var monster = new Vue({
    el:"#search",
    data:{
        name:           'Vue.js',

    },
    methods: {
        searchGo        :function() {
            // 别忘记清除空格
            key = document.search.key.value

            websocket(wsurl,key,function(e){
                console.log("听说这个样可以",e.data,"嗯哼")
            })
        }
    }

})