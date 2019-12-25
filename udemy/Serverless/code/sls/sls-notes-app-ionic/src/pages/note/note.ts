import { Component, ElementRef } from '@angular/core';
import { IonicPage, NavController, NavParams, ToastController, LoadingController } from 'ionic-angular';
import { FormBuilder, Validators } from '@angular/forms';
import { NotesApiService } from '../../app/services/notes-api/notes-api.services';

@IonicPage()
@Component({
	selector: 'page-note',
	templateUrl: 'note.html',
})
export class NotePage {
	note;
	notesCallback;
	disableSubmit;
	noteForm;
	defaultCat = 'general';
	loading;

	constructor(public navCtrl: NavController,
		public navParams: NavParams,
		public element: ElementRef,
		private formBuilder: FormBuilder,
		private notesApiService: NotesApiService,
		public toastCtrl: ToastController,
		public loadingCtrl: LoadingController) {
		this.noteForm = this.formBuilder.group({
			'title': [''],
			'content': ['', Validators.required],
			'cat': [this.defaultCat],
			'timestamp': [''],
			'note_id': ['']
		});
	}

	adjust(id): void {
		let textArea = this.element.nativeElement.getElementsByTagName('textarea')[id];
		if (textArea) {
			textArea.style.overflow = 'hidden';
			textArea.style.height = 'auto';
			textArea.style.height = textArea.scrollHeight + "px";
		}
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

	showLoadingControl() {
		//Loading control cannot be reused. Must be recreated each time.
		this.loading = this.loadingCtrl.create({
		  content: '<ion-spinner>Please wait...</ion-spinner>'
		});
		this.loading.present();
	  }

	ionViewWillEnter() {
		this.note = this.navParams.get('note');
		this.notesCallback = this.navParams.get('callback');
	}

	ionViewDidEnter() {
		this.adjust(0);
		this.adjust(1);
	}

	onSubmit() {
		this.disableSubmit = true;
		this.showLoadingControl();
		if (!this.noteForm.value.timestamp) {
			this.notesApiService.addNote(this.noteForm.value).subscribe(
				note => {
					this.notesCallback({ action: 'add', note: note }).then(() => {
						this.navCtrl.pop();
					}, ()=>{
						this.loading.dismiss();
					});
				},
				err => {
					if (err.error && err.error.message) {
						this.presentToast("An error occurred. " + err.error.message);
					} else {
						this.presentToast("An error occurred. " + err.message);
					}

					this.disableSubmit = false;
					this.loading.dismiss();
				},
				() => {
					this.loading.dismiss();
				}
			);
		} else {
			this.notesApiService.updateNote(this.noteForm.value).subscribe(
			    note => {
			        this.notesCallback({ action: 'update', note: note }).then(() => {
						this.navCtrl.pop();
					}, ()=>{
						this.loading.dismiss();
					});
			    },
			    err => {
			        if (err.error && err.error.message) {
						this.presentToast("An error occurred. " + err.error.message);
					} else {
						this.presentToast("An error occurred. " + err.message);
					}

					this.disableSubmit = false;
			        this.loading.dismiss();
			    },
			    () => {
			        this.loading.dismiss();
			    }
			);
		}
	}
}
