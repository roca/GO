import { Component } from '@angular/core';
import { NavController, ToastController } from 'ionic-angular';

import { NotesPage } from '../notes/notes';

@Component({
  selector: 'page-home',
  templateUrl: 'home.html'
})
export class HomePage {
  res: any;
  err: any;
  

  constructor(public navCtrl: NavController,
              public toastCtrl: ToastController) {
  }

  presentToast(message) {
		let toast = this.toastCtrl.create({
			message: message,
			duration: 5000,
			position: 'top',
			cssClass: 'toast-danger'
		});
		toast.present();
	}

  goToNotes() {
    this.navCtrl.setRoot(NotesPage);
  }

  onLogin() {
    this.goToNotes();
  }

}
