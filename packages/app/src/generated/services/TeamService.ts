/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { PaginatedResult } from '../models/PaginatedResult';
import type { Team } from '../models/Team';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class TeamService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * List all teams
     * @param offset The number of items to skip before starting to collect the result set.
     * @param limit The numbers of items to return.
     * @returns any Successfull response
     * @throws ApiError
     */
    public listTeam(
        offset?: number,
        limit: number = 20,
    ): CancelablePromise<(PaginatedResult & {
        results?: Array<Team>;
    })> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/teams',
            query: {
                'offset': offset,
                'limit': limit,
            },
        });
    }
    /**
     * Creates a new team
     * @param requestBody
     * @returns Team Created
     * @throws ApiError
     */
    public createTeam(
        requestBody?: Team,
    ): CancelablePromise<Team> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/teams',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
}
