import {JwksValidationHandler, ValidationParams} from 'angular-oauth2-oidc';

/**
 * Override the validateAtHash method as we don't request/use the access_token
 */
export class NoAccessTokenValidationHandler extends JwksValidationHandler {
    validateAtHash(validationParams: ValidationParams): Promise<boolean> {
        return Promise.resolve(true);
    }
}
