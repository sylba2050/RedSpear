window.onload = function () {
  var app = new Vue({
    el: '#main',
    data: {
      markdown: "",
      html: "",
      hoge: false,
      markdownLastChangeCount: 0,
      markdown2htmlConvertFlg: false,
      count: 0,
    },
    created: function(){
        setInterval(() => { this.markdownLastChangeCount++ }, 1)
        setInterval(() => { this.count++ }, 100)
    },
    watch: {
        markdown: function(){
            this.markdownLastChangeCount = 0;
            this.markdown2htmlConvertFlg = false;
        },
        count: function(){
            if (this.markdownLastChangeCount < 100){
                return;
            }
            if (this.markdown2htmlConvertFlg){
                return;
            }
            this.markdown2htmlConvertFlg = true;

            var json = {};
            json.md = this.markdown;

            var request = new XMLHttpRequest();
            request.open("POST", "http://153.126.139.150:8080/api/md", true);
            request.setRequestHeader("Content-Type", "application/json");
            request.onreadystatechange = function() {
              if (this.readyState === XMLHttpRequest.DONE && this.status === 200) {
                app.html = request.responseText;
              }
            }
            request.send(JSON.stringify(json));
        },
    },
    methods: {
    }
  })
}
