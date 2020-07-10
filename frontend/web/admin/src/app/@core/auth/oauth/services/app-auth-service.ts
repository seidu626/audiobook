import { Injector, Injectable } from '@angular/core';
import { OAuthService, JwksValidationHandler } from 'angular-oauth2-oidc';
import { filter } from 'rxjs/operators';
import { SessionStorage } from '../../../reactive-storage';
import { authConfig } from '../idsvr-authconfig';
import { Constants } from '../../../constants';


@Injectable({
  providedIn: 'root',
})
// https://medium.com/@philip_lysenko/initial-auth-with-angular-oauth2-a8740efe9264
export class AppAuthService {

  constructor(
    private oauthService: OAuthService,
    private sessionStorage: SessionStorage,
    private injector: Injector,
  ) {}

  async initAuth(): Promise<any> {
    return new Promise((resolveFn, rejectFn) => {
      // setup oauthService
      this.oauthService.configure(authConfig);
      this.oauthService.setupAutomaticSilentRefresh();
      this.oauthService.setStorage(sessionStorage);
      this.oauthService.tokenValidationHandler = new JwksValidationHandler();

      // subscribe to token events
      this.oauthService.events
      .pipe(filter(e => e.type === 'token_received'))
      .subscribe(_ => {
        this.oauthService.loadUserProfile().then(user => {
          if (user) {
            this.sessionStorage.setItem(Constants.USER_INFO_KEY, user);
          }
        });
       // this.handleNewToken();
      });

      this.oauthService.events
      .pipe(filter(e => e.type === 'session_terminated'))
      .subscribe(() => {
        // tslint:disable-next-line:no-console
        console.debug('Your session has been terminated!');
        this.sessionStorage.removeItem(Constants.USER_INFO_KEY);
      });


      // continue initializing app (provoking a token_received event) or redirect to login-page
      this.oauthService.loadDiscoveryDocumentAndTryLogin({
        disableOAuth2StateCheck: true
      }).then(isLoggedIn => {
        if (isLoggedIn) {
          resolveFn();
          // if you don't use clearHashAfterLogin from angular-oauth2-oidc you can remove the #hash or route to some other URL manually:
          // const lazyPath = this.injector.get(LAZY_PATH) as string;
          // this.injector.get(Router).navigateByUrl(lazyPath + '/a') // remove login hash fragments from url
          //   .then(() => resolveFn()); // callback only once login state is resolved
        } else {
          this.oauthService.initImplicitFlow();
          // rejectFn(); //TODO: his throws an error: APP_INITIALIZER: rejected promise throws "Cannot read property 'ngOriginalError' of undefined"
        }
      });

      // this.oauthService.setupAutomaticSilentRefresh();
    });
  }
}
