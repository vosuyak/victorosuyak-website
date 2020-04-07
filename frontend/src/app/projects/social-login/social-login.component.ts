import { Component, OnInit } from '@angular/core';
import { AuthService } from "angularx-social-login";
import { GoogleLoginProvider } from "angularx-social-login";
import { IGmailUser } from './models/social-user';

@Component({
  selector: 'social-login',
  templateUrl: './social-login.component.html',
  styleUrls: ['./social-login.component.scss']
})
export class SocialLoginComponent implements OnInit {
  user: IGmailUser; //gmail user type
  loggedIn: boolean;
 
  constructor(private authService: AuthService) { }

  ngOnInit(): void {
    this.getSocialMediaUser();
  }

  getSocialMediaUser = ():void => {
    this.authService.authState.subscribe(
      (user) => {
        this.user = user;
        this.loggedIn = (user != null);
      },
      error =>{
        console.log('error: ', error);
      },
      () =>{

      }
    );
  }
  signInWithGoogle(): void {
    this.authService.signIn(GoogleLoginProvider.PROVIDER_ID);
  }
 
  signOut(): void {
    this.authService.signOut();
    this.loggedIn = false;
  }
}
