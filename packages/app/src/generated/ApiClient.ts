/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { BaseHttpRequest } from './core/BaseHttpRequest';
import type { OpenAPIConfig } from './core/OpenAPI';
import { FetchHttpRequest } from './core/FetchHttpRequest';
import { DefaultService } from './services/DefaultService';
import { ServerService } from './services/ServerService';
import { SystemService } from './services/SystemService';
import { SystemsService } from './services/SystemsService';
import { TeamService } from './services/TeamService';
type HttpRequestConstructor = new (config: OpenAPIConfig) => BaseHttpRequest;
export class ApiClient {
    public readonly default: DefaultService;
    public readonly server: ServerService;
    public readonly system: SystemService;
    public readonly systems: SystemsService;
    public readonly team: TeamService;
    public readonly request: BaseHttpRequest;
    constructor(config?: Partial<OpenAPIConfig>, HttpRequest: HttpRequestConstructor = FetchHttpRequest) {
        this.request = new HttpRequest({
            BASE: config?.BASE ?? 'http://localhost:8080',
            VERSION: config?.VERSION ?? '1.0.0',
            WITH_CREDENTIALS: config?.WITH_CREDENTIALS ?? false,
            CREDENTIALS: config?.CREDENTIALS ?? 'include',
            TOKEN: config?.TOKEN,
            USERNAME: config?.USERNAME,
            PASSWORD: config?.PASSWORD,
            HEADERS: config?.HEADERS,
            ENCODE_PATH: config?.ENCODE_PATH,
        });
        this.default = new DefaultService(this.request);
        this.server = new ServerService(this.request);
        this.system = new SystemService(this.request);
        this.systems = new SystemsService(this.request);
        this.team = new TeamService(this.request);
    }
}

