var num = null;
var ws = null;
function websocket(url, data, callback) {
	function reconnect(url) {
		var timer = setTimeout(function() {
			if (ws.readyState == 2 || ws.readyState == 3) {
				console.log("readyState" + ws.readyState);
				createwebSocket();
			} else if (ws.readyState == 1) {
				clearTimeout(timer);
			}
		},
		5000);
	}

	createwebSocket();
	function createwebSocket() {
		ws = new WebSocket(url);
		runwebsocket();
	};

	function runwebsocket() {
		ws.onopen = function() {
			//console.log("发送数据");
			ws.send(data);
		};
		ws.onmessage = function(evt) {
			var result = {
				status: 0,
				msg: "ok",
				data: evt.data
			}
			callback(result);
		};
		ws.onclose = function(e) {
			//console.log("websocket连接关闭啦");
			var result = {
				status: 1,
				msg: "close",
				data: e.code
			}

			ws.close()
			//callback(result);
			//reconnect(url);
		};
		ws.οnerrοr = function(e) {
			//console.log("websocket发生错误");
			var result = {
				status: 1,
				msg: "error",
				data: e.code
			}
			callback(result);
		}
	}

};
function wsClose(){
	ws.onclose()
}
