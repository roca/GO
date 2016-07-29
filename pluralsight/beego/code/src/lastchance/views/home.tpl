{{define "content"}}
  <div id="home" data-dojo-type="app/views/HomeView"
      data-dojo-props="
          loginUrl:  '{{urlfor "AccountController.Login"}}',
          createUrl: '{{urlfor "AccountController.Create"}}'
      "
  ></div>
  <script>
    require([
        "dojox/mobile/parser",
        "dojox/mobile/compat",
        "dojo/domReady!",
        'app/views/HomeView'
    ], function (parser) {
        // now parse the page for widgets
        parser.parse();
    });
  </script>
{{end}}
{{template "_layout.tpl"}}
