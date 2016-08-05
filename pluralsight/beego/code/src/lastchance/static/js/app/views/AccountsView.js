define(['dojo/_base/declare',
  'dijit/_WidgetBase',
  'dijit/_TemplatedMixin',
  'dijit/_WidgetsInTemplateMixin',
  'dojox/mobile/View',
  'dojox/mobile/ListItem',
  'dojo/text!./templates/AccountsView.html',
  'dojox/mobile/Heading',
  'dojox/mobile/Button',
  'dojox/mobile/ToolBarButton',
  'dojox/mobile/EdgeToEdgeList'
  ],
function(declare, _WidgetBase, _TemplatedMixin, _WidgetsInTemplateMixin,
  View, ListItem, template) {

    var view = declare([_WidgetBase, _TemplatedMixin, _WidgetsInTemplateMixin], {
      templateString: template,
      view: null,
      navtoTransfer: function() {
        this.view.performTransition('transfers', 1, 'slide');
      },

      postCreate: function() {
        accounts.forEach(function (account) {
          var item = new ListItem();
          item.containerNode.innerHTML =
          account.name + ' ... ' + account.number +
          '<span class="float-right">' +
          '$' + account.amount +
          '</span>';
          item.set('moveTo', 'accounts1');
          this.list.addChild(item);
        }.bind(this));
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
