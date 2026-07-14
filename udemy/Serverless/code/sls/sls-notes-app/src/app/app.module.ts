declare const gapi: any;
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { DatePipe } from '@angular/common';

import { InfiniteScrollModule } from 'ngx-infinite-scroll';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HomeComponent } from './components/home/home.component';
import { NavbarComponent } from './components/navbar/navbar.component';
import { NotesComponent } from './components/notes/notes.component';
import { NoteSnapshotComponent } from './components/notes/note-snapshot.component';
import { NoteComponent } from './components/notes/note.component';
import { SpinnerComponent } from './components/spinner/spinner-component';
import { NotesApiService } from './services/notes-api/notes-api.service';
import { NotesDataService } from './services/notes-data/notes-data.service';
import { LoginComponent } from './components/login/login.component';
import { AuthGuard } from './services/auth-guard/auth-guard.service';
import { AuthService } from './services/auth/auth.service';
import { TitlePipe } from './pipes/extract-title.pipe';

@NgModule({
    declarations: [
        AppComponent,
        HomeComponent,
        NavbarComponent,
        NotesComponent,
        NoteSnapshotComponent,
        NoteComponent,
        SpinnerComponent,
        TitlePipe,
        LoginComponent        
    ],
    imports: [
        BrowserModule,
        AppRoutingModule,
        FormsModule,
        ReactiveFormsModule,
        HttpClientModule,
        InfiniteScrollModule
    ],
    providers: [
        NotesApiService,
        NotesDataService,
        DatePipe,
        AuthGuard,
        AuthService
    ],
    bootstrap: [AppComponent]
})
export class AppModule {}