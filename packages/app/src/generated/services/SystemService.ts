/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Systems } from '../models/Systems';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class SystemService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * List all managed systems.
     * Returns a list of managed systems.
     * @returns Systems A paginated response of all systems
     * @throws ApiError
     */
    public listSystems(): CancelablePromise<Systems> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/systems',
        });
    }
}
