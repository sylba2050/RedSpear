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

            var request_json = {};
            request_json.md = this.markdown;

            fetch("/api/md", {
                method: 'POST',
                mode: 'cors',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(request_json)
            }).then((response) => {
                return response.text()
            }).then((text) => {
                console.log(text);
                app.html = text;
            });
        },
    },
    methods: {
        postArticle : function() {
            var article = {};
            article.userid = 'sylba2050';
            article.content = this.markdown;
            article.cp = 0;
            article.like = '';
            article.comment = '';

            fetch("/post", {
                method: 'POST',
                mode: 'cors',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(article)
            }).then(response => {
                if (response.ok) {
                    console.log("success");
                } else {
                    console.log("failed");
                }
            });
        }
    }
  })
}
