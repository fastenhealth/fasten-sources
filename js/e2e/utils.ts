import {LighthouseEndpointDefinition} from '@shared-library';

export const lighthouseEndpointLookup = {
    'sandbox': 'https://lighthouse.fastenhealth.com/sandbox/',
    'prod': 'https://lighthouse.fastenhealth.com/v1/'
}


export const connectAPIEndpoint = "https://api.connect.fastenhealth.com/v1/"

export async function getEndpointDefinition(endpoint_id: string): Promise<LighthouseEndpointDefinition> {

    let lighthouseEndpoint = lighthouseEndpointLookup['sandbox']

    return await fetch(  lighthouseEndpoint + 'connect/' + endpoint_id)
        .then(resp => resp.json())
        .then(response => {
            console.log(response)
            return response.data as LighthouseEndpointDefinition
        })
}