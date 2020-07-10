import { AuthConfig } from 'angular-oauth2-oidc';
import { environment } from '../../../../environments/environment';

export const authConfig: AuthConfig = {
  // oidc: true,

  requestAccessToken: true,

  postLogoutRedirectUri: `${window.location.origin}`,

  clearHashAfterLogin: true,

  strictDiscoveryDocumentValidation: false,

  // Url of the Identity Provider
  issuer: environment.identityEndpoint,

  requireHttps: true,

  loginUrl: `${environment.identityEndpoint}/connect/authorize`,

  responseType: 'id_token token',

  userinfoEndpoint: `${environment.identityEndpoint}/connect/userinfo`,

  tokenEndpoint: `${environment.identityEndpoint}/connect/token`,

  // dummyClientSecret: '5074F040C4F4B79B45010B33EE7597CD15C8CEADB68793323F3CD79E52C4A7E2',

  // URL of the SPA to redirect the user to after login
  redirectUri: `${window.location.origin}`,

  // URL of the SPA to redirect the user after silent refresh
  silentRefreshRedirectUri: `${window.location.origin}/silent-refresh.html`,

  silentRefreshShowIFrame: false,

  // The SPA's id. The SPA is registerd with this id at the auth-server
  clientId: environment.clientId,

  // set the scope for the permissions the client should request
  // The first three are defined by OIDC. The 4th is a usecase-specific one
  scope: 'openid profile email skoruba_identity_admin_api account_token employee_category',

  skipIssuerCheck: true,

  // silentRefreshShowIFrame: true,

  showDebugInformation: true
};
