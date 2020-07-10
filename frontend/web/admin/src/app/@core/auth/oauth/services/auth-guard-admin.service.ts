import { Injectable } from '@angular/core';
import { CanActivate } from '@angular/router';
import { OAuthService } from 'angular-oauth2-oidc';

@Injectable()
export class AuthGuardAdmin implements CanActivate {
  constructor(private oauthService: OAuthService) { }

  canActivate() {
    if (this.oauthService.hasValidAccessToken()) {
      return true;
    }

    this.oauthService.initImplicitFlow();
  }
}
