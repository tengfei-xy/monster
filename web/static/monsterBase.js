
var next = 0
//             "google"                  ,"baidu" 
const wsurl = ["wss://hk.monsters.vip/ws","wss://monsters.vip/ws"]

function searchGo(clear){
    
    
    if (clear) { next = 0}

    if (next == wsurl.length ){
        return
    }
    key = document.search.key.value.trim()
    
    if (surprise(key)) {return}
    console.log("key:",key," req:",wsurl[next])
    if (key.length != 0){
        if (clear){CleanSearchResult()} 
        
        websocket(wsurl[next],key,function(e){
            r=JSON.parse(e.data)
            for(let i=0;i<r.length;i++){
                (function(line) {
                    setTimeout( function timer() {
                        addLine(line)
                    },i*100);
                })(r[i]);
            }
            
        })
    }
    next++

}

function addLine (line){

    // 二次处理:解决对搜索结果中<>内容被解析成表单问题
    title=line.Title.replace(/\</gi,"&lt;")
    content=line.Content.replace(/\</gi,"&lt;")
    link=line.Link

    // 打印搜索结果
    //console.log(line)

    search_result = document.getElementById("search-result")
    
    cLine = document.createElement('div')
    cLine.className = "search-line"

    cTitle = document.createElement('div')
    cTitle.className = "search-title"

    cContent = document.createElement('div')
    cContent.className = "search-content"

    cTitle_a = document.createElement('a')
    cTitle_a.className = "search-title-a"
    cTitle_a.setAttribute('href',link)
    cTitle_a.setAttribute('target', '_blank');


    cTitle_a.innerHTML = title

    cContent_p = document.createElement('p')
    cContent_p.innerHTML = content

    cLine.appendChild(cTitle)
    cLine.appendChild(cContent)

    cTitle.appendChild(cTitle_a)

    cContent.appendChild(cContent_p)
    
    search_result.appendChild(cLine)




}
// 清除 已显示的搜索条目
function CleanSearchResult(){
    console.log("清除搜索条目")

    var search_result= document.getElementById("search-result")
    while(search_result.hasChildNodes()) //当div下还存在子节点时 循环继续
    {
        search_result.removeChild(search_result.firstChild);
    }
}
// 清除 已显示的搜索条目、已输入搜索关键词
function restore(){
    console.log("清除关键词")
    document.search.key.value=""
    CleanSearchResult()
}

// 滚动判断
window.onscroll = function()
{
    // 搜索框的宽度
    search_page_width = document.getElementById('search-page').clientWidth
    search_box = document.getElementById('search-box-id')
    search_result = document.getElementById('search-result')

    // 视角顶部
    let s = document.body.scrollTop

    // 总页面高度: document.documentElement.clientHeight
    let page_height = document.body.scrollHeight

    let exploer_height = document.body.clientHeight

    // 视角底部
    let bs = s + window.innerHeight

    //console.log("视角顶部:",s," 视角底部:",bs," 总页面高度:",page_height)

    // logo的偏移高度 + logo的高度 = 临界点
    let pd = document.getElementById('search-logo').offsetTop +  document.getElementById('search-logo').clientHeight

    if(s>pd) {
        search_box.style = "position:fixed;top:0px;"
        search_box.style.width=search_page_width
        search_result.style="margin-top:8.8rem;"
    }else {
        search_box.style = "position:relative;"
        search_result.style="margin-top:0rem;"
    }

    // 滑动到接近底部时(视角顶部+页面高度)
    // 开始请求下一个搜索引擎
    if ( bs > page_height - (page_height)/4){
        if (page_height > exploer_height+120){
            searchGo(false)
        }
    }
}

// 按键判断
document.onkeydown=function(e){
    if (e.code=="Enter"){
        searchGo(true)
        return
    }
}

function surprise(key){
    if (key == "朱宏宇" || key =="ZHY") {
        
        CleanSearchResult()
        let line={}

        line.Title = "魔镜魔镜，朱宏宇是不是有历史记载以来最可爱的人"
        line.Content = "魔镜说：“毋庸置疑”"
        line.Link="#"
        addLine(line)
        line.Title = "已为您搜索到这世界最最最最最优秀的人"
        line.Content = "——朱宏宇"
        line.Link="#"
        addLine(line)
        line.Title = "您是否在寻找全全全世界最cool的女孩  | Google搜索"
        line.Content = "全球最大的搜索引擎已匹配到您的搜索结构"
        line.Link="#"
        addLine(line)

        alert("TTF祝ZHY生日快乐！    -- 2020年11月20日")

        return true
    }else if (key=="monster" || key=="Monster"){
        alert("Monster搜索 — 谨此献给ZHY")
        alert("Monster搜索是一种基于百度和Google搜索引擎进行二次开发的，去除前者的广告、无用信息，提供纯粹的、链接式结果的，适配手机、平板、笔记本的搜索平台。")
        return true
        
    }else if (key=="TTF"){
        alert("你想知道的都在这里,不过github内网访问比较慢,https://github.com/tengfei-xy/monster")
        return true
        
    }

    return false
}