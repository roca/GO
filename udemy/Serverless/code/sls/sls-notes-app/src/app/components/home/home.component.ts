import { Component, OnInit } from '@angular/core';
import { AuthService } from '../../services/auth/auth.service';
import { NotesDataService } from '../../services/notes-data/notes-data.service';

@Component({
    selector: 'app-home',
    templateUrl: 'home.component.html'
})
export class HomeComponent implements OnInit {

    showNoteModal = false;
    newNote;
    constructor(private notesDataService: NotesDataService,
        private authService: AuthService
    ) {  }

    ngOnInit() { 
        this.newNote = {};
        
    }

    onShowNoteModal($event) {
        this.showNoteModal = true;
    }

    onCloseNoteModal($event) {
        this.showNoteModal = false;
    }

    onSignOut() {
        this.authService.logout();
    }

}