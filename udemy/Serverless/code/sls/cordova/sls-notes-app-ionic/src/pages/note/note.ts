import { Component } from '@angular/core';
import { IonicPage, NavController, NavParams } from 'ionic-angular';

@IonicPage()
@Component({
  selector: 'page-note',
  templateUrl: 'note.html',
})
export class NotePage {
  note;

  constructor(public navCtrl: NavController, public navParams: NavParams) {
  }

  ionViewWillEnter() {
    // console.log(this.navParams.get('note'));
    this.note = this.navParams.get('note');
  }

  ionViewDidLoad() {
    // console.log('ionViewDidLoad NotePage');
  }

  onSaveNote() {
    this.navCtrl.pop();
  }

}
