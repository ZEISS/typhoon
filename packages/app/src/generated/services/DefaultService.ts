/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Team } from '../models/Team';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class DefaultService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * Gets a team by ID
     * @param teamId
     * @returns Team Successful
     * @throws ApiError
     */
    public getTeam(
        teamId: string,
    ): CancelablePromise<Team> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/team/{teamId}',
            path: {
                'teamId': teamId,
            },
        });
    }
}
