import * as Oauth from '@panva/oauth4webapi';
import {uuidV4} from '../utils/uuid';
import {SourceState} from '../models/source-state';
import {LighthouseEndpointDefinition} from '../models/lighthouse';

export async function generateSourceAuthorizeUrl(lighthouseSource: LighthouseEndpointDefinition, reconnectSourceId?: string): Promise<{url: URL, sourceState: SourceState}> {
    const state = uuidV4()
    let sourceStateInfo = new SourceState()
    sourceStateInfo.state = state
    sourceStateInfo.endpoint_id = lighthouseSource.id
    sourceStateInfo.portal_id = lighthouseSource.portal_id
    sourceStateInfo.brand_id = lighthouseSource.brand_id
    if(reconnectSourceId){
        //if the source already exists, and we want to re-connect it (because of an expiration), we need to pass the existing source id
        sourceStateInfo.reconnect_source_id = reconnectSourceId
    }

    // generate the authorization url
    const authorizationUrl = new URL(lighthouseSource.authorization_endpoint);
    authorizationUrl.searchParams.set('redirect_uri', lighthouseSource.redirect_uri);
    authorizationUrl.searchParams.set('response_type', lighthouseSource.response_types_supported[0]);
    authorizationUrl.searchParams.set('response_mode', lighthouseSource.response_modes_supported[0]);
    authorizationUrl.searchParams.set('state', state);
    authorizationUrl.searchParams.set('client_id', lighthouseSource.client_id);
    if(lighthouseSource.scopes_supported && lighthouseSource.scopes_supported.length){
        authorizationUrl.searchParams.set('scope', lighthouseSource.scopes_supported.join(' '));
    } else {
        authorizationUrl.searchParams.set('scope', '');
    }
    if (lighthouseSource.aud) {
        authorizationUrl.searchParams.set('aud', lighthouseSource.aud);
    }

    //this is for providers that support CORS and PKCE (public client auth)
    if(!lighthouseSource.confidential || (lighthouseSource.code_challenge_methods_supported || []).length > 0){
        // https://github.com/panva/oauth4webapi/blob/8eba19eac408bdec5c1fe8abac2710c50bfadcc3/examples/public.ts
        const codeVerifier = Oauth.generateRandomCodeVerifier();
        const codeChallenge = await Oauth.calculatePKCECodeChallenge(codeVerifier);
        const codeChallengeMethod = lighthouseSource.code_challenge_methods_supported?.[0] || 'S256'

        sourceStateInfo.code_verifier = codeVerifier
        sourceStateInfo.code_challenge = codeChallenge
        sourceStateInfo.code_challenge_method = codeChallengeMethod

        authorizationUrl.searchParams.set('code_challenge', codeChallenge);
        authorizationUrl.searchParams.set('code_challenge_method', codeChallengeMethod);
    }

    return {url: authorizationUrl, sourceState: sourceStateInfo}
}
