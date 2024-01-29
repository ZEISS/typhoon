/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Version } from '../models/Version';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class ServerService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * Returns the current version of the API.
     * @returns Version
     * @throws ApiError
     */
    public version(): CancelablePromise<Version> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/version',
        });
    }
}
