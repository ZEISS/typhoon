/* tslint:disable */
/* eslint-disable */
/**
 * Typhoon API
 * # Introduction  The Typhoon API allows you to NATS.io 
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  System,
} from '../models/index';
import {
    SystemFromJSON,
    SystemToJSON,
} from '../models/index';

export interface ShowSystemRequest {
    systemId: string;
}

/**
 * 
 */
export class SystemsApi extends runtime.BaseAPI {

    /**
     * Returns a list of managed systems.
     * List all managed systems.
     */
    async listSystemsRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<System>>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/systems`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(SystemFromJSON));
    }

    /**
     * Returns a list of managed systems.
     * List all managed systems.
     */
    async listSystems(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<System>> {
        const response = await this.listSystemsRaw(initOverrides);
        return await response.value();
    }

    /**
     * Returns a list of managed systems.
     * List all managed systems.
     */
    async showSystemRaw(requestParameters: ShowSystemRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.systemId === null || requestParameters.systemId === undefined) {
            throw new runtime.RequiredError('systemId','Required parameter requestParameters.systemId was null or undefined when calling showSystem.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/systems/{systemId}`.replace(`{${"systemId"}}`, encodeURIComponent(String(requestParameters.systemId))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Returns a list of managed systems.
     * List all managed systems.
     */
    async showSystem(requestParameters: ShowSystemRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.showSystemRaw(requestParameters, initOverrides);
    }

}
