/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class SystemsService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * List all managed systems.
     * Returns a list of managed systems.
     * @param systemId The id of the system to retrieve
     * @returns any A paginated response of all systems
     * @throws ApiError
     */
    public showSystem(
        systemId: string,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/systems/{systemId}',
            path: {
                'systemId': systemId,
            },
        });
    }
}
