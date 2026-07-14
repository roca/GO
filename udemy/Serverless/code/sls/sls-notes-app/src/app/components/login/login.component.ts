import { Component, OnInit } from '@angular/core';
import { AuthService } from '../../services/auth/auth.service';

@Component({
    selector: 'login',
    templateUrl: 'login.component.html'
})

export class LoginComponent implements OnInit {

    constructor(private authService: AuthService) {
    }

    ngOnInit() { }

    onSignIn() {
        this.authService.login();
    }
}