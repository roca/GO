import { Component, OnInit } from '@angular/core';
import { IonicPage, NavController, NavParams, LoadingController, ToastController } from 'ionic-angular';
import { NotePage } from '../note/note';
import { HomePage } from '../home/home';
import { NotesApiService } from '../../app/services/notes-api/notes-api.services';
import * as _ from "lodash";
import { AuthService } from '../../app/services/auth/auth.service';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/mergeMap';

@IonicPage()
@Component({
  selector: 'page-notes',
  templateUrl: 'notes.html',
})
export class NotesPage implements OnInit {
  userNotes;
  isLoading;
  startKey;
  loading;

  constructor(public navCtrl: NavController,
    public navParams: NavParams,
    public loadingCtrl: LoadingController,
    private notesApiService: NotesApiService,
    public toastCtrl: ToastController,
    private authService: AuthService) {
  }

  notePageCallback = (params) => {
    return new Promise((resolve, reject) => {
      if ("add" == params.action) {
        this.addNote(params.note);
      } else if ("update" == params.action) {
        this.updateNote(params.note);
      } else {
        reject();
      }

      resolve();
    });
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

  ngOnInit() {
    this.isLoading = true;
    this.refreshNotes();
  }

  showLoadingControl() {
    //Loading control cannot be reused. Must be recreated each time.
    this.loading = this.loadingCtrl.create({
      content: '<ion-spinner>Please wait...</ion-spinner>'
    });
    this.loading.present();
  }

  refreshNotes() {
    this.isLoading = true;
    this.showLoadingControl();
    this.userNotes = [];
    this.notesApiService.getNotes().subscribe(
      res => {
        if (_.has(res, 'LastEvaluatedKey')) {
          this.startKey = res.LastEvaluatedKey.timestamp;
        } else {
          this.startKey = 0;
        }

        if (_.has(res, 'notes')) {
          this.userNotes = _.union(this.userNotes, res.notes);
        }
      }, err => {
        if (err.error && err.error.message) {
          this.presentToast("An error occurred. " + err.error.message);
        } else {
          this.presentToast("An error occurred. " + err.message);
        }
        this.isLoading = false;
        this.loading.dismiss();
      }, () => {
        this.isLoading = false;
        this.loading.dismiss();
      }
    );
  }

  onScrollDown(infiniteScroll) {
    if (this.startKey == 0) {
      infiniteScroll.enable(false);
      return;
    }

    this.notesApiService.getNotes(this.startKey).subscribe(
      res => {
        if (this.startKey == 0) {
          infiniteScroll.complete();
          return;
        }

        if (_.has(res, 'LastEvaluatedKey')) {
          this.startKey = res.LastEvaluatedKey.timestamp;
        } else {
          this.startKey = 0;
        }

        if (_.has(res, 'notes')) {
          this.userNotes = _.union(this.userNotes, res.notes);
        }
      }, err => {
        if (err.error && err.error.message) {
          this.presentToast("An error occurred. " + err.error.message);
        } else {
          this.presentToast("An error occurred. " + err.message);
        }
        infiniteScroll.complete();
      }, () => {
        infiniteScroll.complete();
      }
    );
  }

  onCreateNote() {
    this.openNote({});
  }

  openNote(note) {
    this.navCtrl.push(NotePage, { note: note, callback: this.notePageCallback });
  }

  addNote(note) {
    this.userNotes.splice(0, 0, note);
  }

  updateNote(note) {
    let index = _.findIndex(this.userNotes, (x) => { return x.timestamp == note.timestamp; });
    this.userNotes.splice(index, 1, note);
  }

  deleteNote(note) {
    let index = _.findIndex(this.userNotes, (x) => { return x.timestamp == note.timestamp; });
    this.userNotes.splice(index, 1);

    this.notesApiService.deleteNote(note.timestamp).subscribe(
      res => {
        //do nothing
      }, err => {
        this.refreshNotes(); //refresh notes again
        if (err.error && err.error.message) {
          this.presentToast("An error occurred. " + err.error.message);
        } else {
          this.presentToast("An error occurred. " + err.message);
        }
      }, () => {
      }
    );
  }

  onLogout() {
    this.authService.logout().then(() => {
      this.userNotes = [];
      this.navCtrl.setRoot(HomePage);
    }).catch(() => {
      this.userNotes = [];
      this.navCtrl.setRoot(HomePage);
    });
  }

}
