import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ProjectsRoutingModule } from './projects-routing.module';
import { ProjectSnippetComponent } from './project-snippet/project-snippet.component';
import { SocialLoginComponent } from './social-login/social-login.component';

// login project
import { SocialLoginModule, AuthServiceConfig } from "angularx-social-login";
import { GoogleLoginProvider, FacebookLoginProvider } from "angularx-social-login";
import { ProjectsComponent } from './projects.component';
let clientId = '587421380986-qrun4in4pil2cieientqucocurju8cog.apps.googleusercontent.com';
let clientSecret = 'RqmHn8jEUCeEPE97zR-b-gPM'; 

let config = new AuthServiceConfig([
  {
    id: GoogleLoginProvider.PROVIDER_ID,
    provider: new GoogleLoginProvider(clientId)
  },
  {
    id: FacebookLoginProvider.PROVIDER_ID,
    provider: new FacebookLoginProvider("Facebook-App-Id")
  }
]);
export function provideConfig() {
  return config;
}
@NgModule({
  declarations: [ProjectsComponent, ProjectSnippetComponent, SocialLoginComponent],
  imports: [
    CommonModule,
    SocialLoginModule,
    ProjectsRoutingModule
  ],
  providers: [
    {
      provide: AuthServiceConfig,
      useFactory: provideConfig
    }
  ],
})



export class ProjectsModule {
}
