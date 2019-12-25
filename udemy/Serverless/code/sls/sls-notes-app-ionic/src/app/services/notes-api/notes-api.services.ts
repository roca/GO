import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs/Observable';

@Injectable()
export class NotesApiService {

    API_ROOT;
    STAGE;
    options;

    constructor(private httpClient: HttpClient) {
        this.API_ROOT = 'http://localhost:3000';
        this.STAGE = ''
        this.setOptions();
    }

    setOptions() {
        this.options = {
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                app_user_id: 'test_user',
                app_user_name: 'Test User'
            }
        };  
    }

    addNote(item) {
        let path = this.STAGE + '/note';
        let endpoint = this.API_ROOT + path;

        let itemData;
        itemData = {
            content: item.content,
            cat: item.cat
        };

        if (item.title != "") {
            itemData.title = item.title;
        }

        this.setOptions();
        return this.httpClient.post(endpoint, itemData, this.options);
    }

    updateNote(item) {
        let path = this.STAGE + '/note';
        let endpoint = this.API_ROOT + path;


        let itemData;
        itemData = {
            content: item.content,
            cat: item.cat,
            timestamp: parseInt(item.timestamp),
            note_id: item.note_id
        };

        if (item.title != "") {
            itemData.title = item.title;
        }

        this.setOptions();
        return this.httpClient.patch(endpoint, itemData, this.options);
    }

    deleteNote(timestamp): Observable<any> {
        let path = this.STAGE + '/note/t/' + timestamp;
        let endpoint = this.API_ROOT + path;

        this.setOptions();
        return this.httpClient.delete(endpoint, this.options);
    }

    getNotes(start?): Observable<any> {
        let path = this.STAGE + '/note?limit=24';
        let endpoint = this.API_ROOT + path;

        if (start > 0) {
            endpoint += '&start=' + start;
        }
        this.setOptions();
        return this.httpClient.get(endpoint, this.options);
    }

}