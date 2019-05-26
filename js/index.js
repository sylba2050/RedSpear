window.onload = function() {
  var main = new Vue({
    el: "#main",
    data: {
      articles: [],
      article_detail_data: {},
      is_detail: false,
      html: ""
    },
    beforeCreate: function() {
      fetch("/articles", {
        method: "GET",
        mode: "cors"
      })
        .then(response => {
          return response.json();
        })
        .then(json => {
          this.articles = json;
        });
    },
    watch: {
      article_detail_data: function() {
        var request_json = {};
        request_json.md = this.article_detail_data.content;
        console.log(request_json);

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
            this.html = text;
          });
      }
    },
    methods: {
      detail: function(id) {
        this.is_detail = true;

        fetch("/article/" + id, {
          method: "GET",
          mode: "cors"
        })
          .then(response => {
            return response.json();
          })
          .then(json => {
            this.article_detail_data = json;
            console.log(this.article_detail_data);
          });
      }
    }
  });
};
