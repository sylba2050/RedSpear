window.onload = function() {
  var main = new Vue({
    el: "#main",
    data: {
      articles: [],
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
    },
    methods: {
      detail: function(id) {
        location.href = "/items/" + id;
      }
    }
  });
};
