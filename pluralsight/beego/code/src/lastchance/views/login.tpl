{{define "content"}}
  <div id="login" data-dojo-type="app/views/LoginView"></div>
  <script>
    require([
        "dojox/mobile/parser",
        "dojox/mobile/compat",
        "dojo/domReady!",
        'app/views/LoginView'
    ], function (parser) {
        // now parse the page for widgets
        parser.parse();
    });
</script>
{{end}}
{{template "_layout.tpl"}}
