import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs/Observable';

import { RequestSigner } from 'aws4';

import { AuthService } from '../auth/auth.service';

@Injectable()
export class NotesApiService {

    API_ROOT;
    STAGE;
    options;

    constructor(private httpClient: HttpClient,
        private authService: AuthService) {
        this.API_ROOT = 'https://notesapi.desertfoxdev.org';
        //this.API_ROOT = 'http://localhost:3000';
        this.STAGE = '/v1' // Put your API Stage path here
        this.setOptions();
    }

    setOptions(path = '/', method = '', body = '') {

        const host = new URL(this.API_ROOT);

        let args = {
            service: 'execute-api',
            region: 'us-east-1',
            hostname: host.hostname,
            path: path,
            method: method,
            body: body,
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json'
            }
        }; 

        if(method == 'GET') {
            delete args.body;
        }

        this.options = {};
        try {
            let savedCredsJson = this.authService.getCredentials();

            if(savedCredsJson) {
                let savedCreds = JSON.parse(savedCredsJson);
                let creds = {
                    accessKeyId: savedCreds.cognito_data.Credentials.AccessKeyId,
                    secretAccessKey: savedCreds.cognito_data.Credentials.SecretKey,
                    sessionToken: savedCreds.cognito_data.Credentials.SessionToken
                };
                
                let signer = new RequestSigner(args, creds);
                let signed = signer.sign();
                
                this.options.headers = signed.headers;
                delete this.options.headers.Host;
                this.options.headers.app_user_id = savedCreds.cognito_data.IdentityId;
                this.options.headers.app_user_name = savedCreds.user_name;
            }
        } catch (error) {
            // do nothing
        }        
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

        this.setOptions(path, 'POST', JSON.stringify(itemData));
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

        this.setOptions(path, 'PATCH', JSON.stringify(itemData));
        return this.httpClient.patch(endpoint, itemData, this.options);
    }

    deleteNote(timestamp): Observable<any> {
        let path = this.STAGE + '/note/t/' + timestamp;
        let endpoint = this.API_ROOT + path;

        this.setOptions(path, 'DELETE');
        return this.httpClient.delete(endpoint, this.options);
    }

    getNotes(start?): Observable<any> {
        let path = this.STAGE + '/note?limit=24';
        let endpoint = this.API_ROOT + path;

        if (start > 0) {
            endpoint += '&start=' + start;
        }
        this.setOptions(path, 'GET');
        return this.httpClient.get(endpoint, this.options);
    }

}