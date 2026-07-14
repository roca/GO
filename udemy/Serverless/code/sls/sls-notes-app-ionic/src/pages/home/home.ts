import { Component } from '@angular/core';
import { NavController, ToastController } from 'ionic-angular';

import { NotesPage } from '../notes/notes';
import { AuthService } from '../../app/services/auth/auth.service';

@Component({
  selector: 'page-home',
  templateUrl: 'home.html'
})
export class HomePage {
  res: any;
  err: any;
  

  constructor(public navCtrl: NavController,
              private authService: AuthService,
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
    this.authService.login().then(
      (res)=> {
        this.goToNotes();
      },
      (err) => {
        this.presentToast("We encounted an error while logging you in. Please try again.");
      }
    );
  }

}
