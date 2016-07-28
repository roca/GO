{{define "content"}}
<div id="accounts" data-dojo-type="app/views/AccountsView"></div>
<div id="accounts1" data-dojo-type="app/views/AccountsView"></div>
<div id="accounts2" data-dojo-type="app/views/AccountsView"></div>
<div id="transfers" data-dojo-type="app/views/TransfersView"></div>
<script>
  require([
      "dojox/mobile/parser",
      "dojox/mobile/compat",
      "dojo/domReady!",
      'app/views/AccountsView',
      'app/views/TransfersView'
  ], function (parser) {
      parser.parse();
  });
</script>
{{end}}
{{template "_layout.tpl"}}
