import { Injectable } from '@angular/core';
import { CanActivate, Router } from '@angular/router';
import { OAuthService } from 'angular-oauth2-oidc';

@Injectable()
export class AuthGuard implements CanActivate {
  constructor(private securityService: OAuthService) { }

  canActivate() {
    if (this.securityService.hasValidIdToken()) {
      return true;
    }
    this.securityService.initImplicitFlow();
    return false;
  }
}
