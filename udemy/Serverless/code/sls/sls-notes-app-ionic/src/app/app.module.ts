import { BrowserModule } from '@angular/platform-browser';
import { ErrorHandler, NgModule } from '@angular/core';
import { IonicApp, IonicErrorHandler, IonicModule } from 'ionic-angular';
import { SplashScreen } from '@ionic-native/splash-screen';
import { StatusBar } from '@ionic-native/status-bar';

import { HttpClientModule } from '@angular/common/http';

import { MyApp } from './app.component';
import { HomePage } from '../pages/home/home';
import { NotesPage } from '../pages/notes/notes';
import { NotePage } from '../pages/note/note';
import { AutoSizeDirective } from './directives/auto-size/auto-size';
import { NotesApiService } from './services/notes-api/notes-api.services';
import { TitlePipe } from './pipes/extract-title.pipe';


@NgModule({
  declarations: [
    MyApp,
    HomePage,
    NotesPage,
    NotePage,
    AutoSizeDirective,
    TitlePipe
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    IonicModule.forRoot(MyApp)
  ],
  bootstrap: [IonicApp],
  entryComponents: [
    MyApp,
    HomePage,
    NotesPage,
    NotePage
  ],
  providers: [
    StatusBar,
    SplashScreen,
    {provide: ErrorHandler, useClass: IonicErrorHandler},
    NotesApiService
  ]
})
export class AppModule {}
