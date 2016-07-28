{{define "content"}}
  <div id="createAccount" data-dojo-type="app/views/CreateAccountView"></div>
  <script>
    require([
        "dojox/mobile/parser",
        "dojox/mobile/compat",
        "dojo/domReady!",
        'app/views/CreateAccountView'
    ], function (parser) {
        // now parse the page for widgets
        parser.parse();
    });
</script>
{{end}}
{{template "_layout.tpl"}}
