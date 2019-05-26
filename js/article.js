window.onload = function() {
  var main = new Vue({
    el: "#main",
    data: {
      id: "",
      article_detail_data: {},
      html: ""
    },
    beforeCreate: function() {
      var s = location.href.split("/")
      this.id = s[s.length - 1];

      fetch("/article/" + this.id, {
        method: "GET",
        mode: "cors"
      })
        .then(response => {
          return response.json();
        })
        .then(json => {
          this.article_detail_data = json;
        });
    },
    watch: {
      article_detail_data: function() {
        var request_json = {};
        request_json.md = this.article_detail_data.content;

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
    }
  });
};
