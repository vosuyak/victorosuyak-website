import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ProjectsRoutingModule } from './projects-routing.module';
import { ProjectSnippetComponent } from './project-snippet/project-snippet.component';
import { SocialLoginComponent } from './social-login/social-login.component';
import { ProjectsComponent } from './projects.component';
import { TestComponent } from './test/test.component';
// login project
import { SocialLoginModule, AuthServiceConfig } from "angularx-social-login";
import { GoogleLoginProvider } from "angularx-social-login";


const clientId = '587421380986-qrun4in4pil2cieientqucocurju8cog.apps.googleusercontent.com';

let config = new AuthServiceConfig([
  {
    id: GoogleLoginProvider.PROVIDER_ID,
    provider: new GoogleLoginProvider(clientId)
  }
]);
export function provideConfig() {
  return config;
}
@NgModule({
  declarations: [ProjectsComponent, ProjectSnippetComponent, SocialLoginComponent, TestComponent],
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
