import { Component, OnInit } from '@angular/core';
import { IonicPage, NavController, NavParams } from 'ionic-angular';
import { NotePage } from '../note/note';
import { HomePage } from '../home/home';

@IonicPage()
@Component({
  selector: 'page-notes',
  templateUrl: 'notes.html',
})
export class NotesPage implements OnInit {
  userNotes;
  constructor(public navCtrl: NavController, 
              public navParams: NavParams) {
  }

  ionViewWillEnter() {
    // this.viewCtrl.showBackButton(false);
    this.userNotes = [
      {
        title: "Note 1",
        content: "This is the content of note 1",
        timestamp: 1538896786
      },
      {
        title: "Note 2",
        content: "This is the content of note 2",
        timestamp: 1538896775
      },
      {
        title: "Note 3",
        content: "This is the content of note 3",
        timestamp: 1538896600
      },
      {
        title: "Note 4",
        content: "This is the content of note 4",
        timestamp: 1538896517
      }
    ];
  }

  ngOnInit() {
    
  }

  ionViewDidLoad() {
    // console.log('ionViewDidLoad NotesPage');
  }

  onCreateNote () {
    this.openNote({});
  }

  openNote(note) {
    this.navCtrl.push(NotePage, {note: note});
  }

  deleteNote(note) {
    for(let i = 0; i < this.userNotes.length; i++) {
 
      if(this.userNotes[i] == note){
        this.userNotes.splice(i, 1);
      }
    }
  }

  onLogout() {
    this.navCtrl.setRoot(HomePage);
  }

}
