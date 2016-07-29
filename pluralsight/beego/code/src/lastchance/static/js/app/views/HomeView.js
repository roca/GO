define(['dojo/_base/declare',
  'dijit/_WidgetBase',
  'dijit/_TemplatedMixin',
  'dijit/_WidgetsInTemplateMixin',
  'dojox/mobile/View',
  'dojo/text!./templates/HomeView.html',
  'dojox/mobile/Button'],
function(declare, _WidgetBase, _TemplatedMixin, _WidgetsInTemplateMixin,
  View, template) {

    var view = declare([_WidgetBase, _TemplatedMixin, _WidgetsInTemplateMixin], {
      loginUrl:  null,
      createUrl: null,
      templateString: template,
      onLogin: function() {
        window.location = this.loginUrl;
      }
    });

  return declare([View], {
    view: null,
    loginUrl:  null,
    createUrl: null,
    postCreate: function() {
      this.view = new view({
        loginUrl:  this.loginUrl,
        createUrl: this.createUrl,
      });
      this.containerNode.appendChild(this.view.domNode);
    }
  });
});
