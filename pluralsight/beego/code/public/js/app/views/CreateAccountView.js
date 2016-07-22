define(['dojo/_base/declare',
  'dijit/_WidgetBase',
  'dijit/_TemplatedMixin',
  'dijit/_WidgetsInTemplateMixin',
  'dojox/mobile/View',
  'dojo/text!./templates/CreateAccountView.html',
  'dojox/mobile/Heading',
  'dojox/mobile/TextBox',
  'dojox/mobile/Button'],
function(declare, _WidgetBase, _TemplatedMixin, _WidgetsInTemplateMixin,
  View, template) {

    var view = declare([_WidgetBase, _TemplatedMixin, _WidgetsInTemplateMixin], {
      templateString: template
    });

  return declare([View], {
    view: null,
    postCreate: function() {
      this.view = new view();
      this.containerNode.appendChild(this.view.domNode);
    }
  });
});
