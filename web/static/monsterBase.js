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
            console.log("搜索的关键词是",key)
            if (key.length != 0){
                CleanSearchResult()
                websocket(wsurl,key,function(e){
                    //console.log(e.data))
                    /* 
                    (function(url,seq) {
				setTimeout( function timer() {
					console.log(seq,url)
					window.open("https://www.baidu.com/s?wd="+url,"_blank")
				}, seq*1000 );
			})(unescape(ttt[i]),i+1);	
                    */
                    r=JSON.parse(e.data)
                    for(let i=0;i<r.length;i++){
                        (function(line) {
                            setTimeout( function timer() {
                                addLine(line)
                            },i*300);
                        })(r[i]);
                    }
                    
                })
            }

        }
    }

})
function addLine (line){
    title=line.Title
    content=line.Content
    link=line.Link
    console.log(line)
    /*
    <div class="search-line">
        <div class="search-title">
            <a class="search-title-a" href="http://www.baidu.com/link?url=ZY97CuDE-k1y0bRV-B0J072daNadWQxcwHsRasoY_pysVgnZfourWTuCfZVFi2A8">中国网--网上中国</a>
        </div>

        <div class="search-content">
            <p>国家重点新闻网站,拥有十个语种独立新闻采编、报道和发布权;第一时间报道国家重大新闻事件;国情信息库服务全球读者了解中国;国务院新闻办公室发布会独家网络直播发布...</p>
        </div>
    </div>

    */
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
function CleanSearchResult(){
    var search_result= document.getElementById("search-result")
    while(search_result.hasChildNodes()) //当div下还存在子节点时 循环继续
    {
        search_result.removeChild(search_result.firstChild);
    }
}