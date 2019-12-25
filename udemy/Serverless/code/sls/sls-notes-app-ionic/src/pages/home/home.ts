import { Component } from '@angular/core';
import { NavController } from 'ionic-angular';
import { NotesPage } from '../notes/notes';

@Component({
  selector: 'page-home',
  templateUrl: 'home.html'
})
export class HomePage {

  constructor(public navCtrl: NavController) {

  }

  onGoToNotes() {
    // this.navCtrl.push(NotesPage);
    this.navCtrl.setRoot(NotesPage);
  }

}
