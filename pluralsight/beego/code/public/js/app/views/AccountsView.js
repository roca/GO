define(['dojo/_base/declare',
  'dijit/_WidgetBase',
  'dijit/_TemplatedMixin',
  'dijit/_WidgetsInTemplateMixin',
  'dojox/mobile/View',
  'dojo/text!./templates/AccountsView.html',
  'dojox/mobile/Heading',
  'dojox/mobile/Button',
  'dojox/mobile/ToolBarButton',
  'dojox/mobile/EdgeToEdgeList',
  'dojox/mobile/ListItem'],
function(declare, _WidgetBase, _TemplatedMixin, _WidgetsInTemplateMixin,
  View, template) {

    var view = declare([_WidgetBase, _TemplatedMixin, _WidgetsInTemplateMixin], {
      templateString: template,
      view: null,
      navtoTransfer: function() {
        this.view.performTransition('transfers', 1, 'slide');
      }
    });

  return declare([View], {
    view: null,
    postCreate: function() {
      this.view = new view({view: this});
      this.containerNode.appendChild(this.view.domNode);
    }
  });
});
