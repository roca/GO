define(['dojo/_base/declare',
  'dijit/_WidgetBase',
  'dijit/_TemplatedMixin',
  'dijit/_WidgetsInTemplateMixin',
  'dojox/mobile/View',
  'dojo/request/xhr',
  'dojo/text!./templates/TransfersView.html',
  'dojox/mobile/Heading',
  'dojox/mobile/Button',
  'dojox/mobile/TextBox',
  'dojox/mobile/ComboBox',
  'dijit/form/DataList',
  'dijit/form/Form'],
function(declare, _WidgetBase, _TemplatedMixin, _WidgetsInTemplateMixin,
  View, request, template) {

    var view = declare([_WidgetBase, _TemplatedMixin, _WidgetsInTemplateMixin], {
      templateString: template,

      postCreate: function() {
        this.form.on("submit",function(e){
          var data = this.form.get("value");
          data.amount = parseFloat(data.amount);

          request.post('/api/transfer', {
            headers: {
              "Content-Type": "application/json"
            },
            data: JSON.stringify(data)
          });

          e.preventDefault();
        }.bind(this));
      }
    });

  return declare([View], {
    view: null,
    postCreate: function() {
      this.view = new view();
      this.containerNode.appendChild(this.view.domNode);
    }
  });
});
