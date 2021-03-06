window.onload = function() {
  var app = new Vue({
    el: "#main",
    data: {
      title: "",
      markdown: "",
      html: "",
      hoge: false,
      markdownLastChangeCount: 0,
      markdown2htmlConvertFlg: false,
      count: 0
    },
    created: function() {
      setInterval(() => {
        this.markdownLastChangeCount++;
      }, 1);
      setInterval(() => {
        this.count++;
      }, 100);
    },
    watch: {
      markdown: function() {
        this.markdownLastChangeCount = 0;
        this.markdown2htmlConvertFlg = false;
      },
      count: function() {
        if (this.markdownLastChangeCount < 100) {
          return;
        }
        if (this.markdown2htmlConvertFlg) {
          return;
        }
        this.markdown2htmlConvertFlg = true;

        var request_json = {};
        request_json.md = this.markdown;

        fetch("/md", {
          method: "POST",
          mode: "cors",
          headers: {
            "Content-Type": "application/json"
          },
          body: JSON.stringify(request_json)
        })
          .then(response => {
            return response.text();
          })
          .then(text => {
            app.html = text;
          });
      }
    },
    methods: {
      postArticle: function() {
        var article = {};
        article.userid = "sylba2050";
        article.title = this.title;
        article.content = this.markdown;
        article.cp = 0;

        fetch("/article", {
          method: "POST",
          mode: "cors",
          headers: {
            "Content-Type": "application/json"
          },
          body: JSON.stringify(article)
        }).then(response => {
          if (response.ok) {
            location.href = "/";
          } else {
            console.log("failed");
          }
        });
      }
    }
  });
};
