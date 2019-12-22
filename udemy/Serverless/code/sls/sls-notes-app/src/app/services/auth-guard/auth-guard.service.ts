import { Injectable } from '@angular/core';
import { Router, CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';
import { AuthService } from '../auth/auth.service';

@Injectable()
export class AuthGuard {

    constructor(private authService: AuthService,
                private router: Router) { }

    async canActivate(route: ActivatedRouteSnapshot, 
                state: RouterStateSnapshot) {
        
        try {
            await this.authService.isLoggedIn();
            return true;
        } catch (err) {
            this.router.navigate(['login']);
            return false;
        }
    }
}