/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Team } from '../models/Team';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class TeamService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * List all teams
     * @returns any Returns the created team
     * @throws ApiError
     */
    public listTeam(): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/teams',
        });
    }
    /**
     * Creates a new team
     * @param requestBody
     * @returns any Returns the created team
     * @throws ApiError
     */
    public createTeam(
        requestBody?: Team,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/teams',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
}
