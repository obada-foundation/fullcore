import { Reader, Writer } from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { Obit } from "../fullcore/obit";
export declare const protobufPackage = "obadafoundation.fullcore.fullcore";
export interface QueryAllObitRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllObitResponse {
    Obit: Obit[];
    pagination: PageResponse | undefined;
}
export interface QueryGetObitRequest {
    did: string;
}
export interface QueryGetObitResponse {
    Obit: Obit | undefined;
}
export declare const QueryAllObitRequest: {
    encode(message: QueryAllObitRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllObitRequest;
    fromJSON(object: any): QueryAllObitRequest;
    toJSON(message: QueryAllObitRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllObitRequest>): QueryAllObitRequest;
};
export declare const QueryAllObitResponse: {
    encode(message: QueryAllObitResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllObitResponse;
    fromJSON(object: any): QueryAllObitResponse;
    toJSON(message: QueryAllObitResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllObitResponse>): QueryAllObitResponse;
};
export declare const QueryGetObitRequest: {
    encode(message: QueryGetObitRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetObitRequest;
    fromJSON(object: any): QueryGetObitRequest;
    toJSON(message: QueryGetObitRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetObitRequest>): QueryGetObitRequest;
};
export declare const QueryGetObitResponse: {
    encode(message: QueryGetObitResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetObitResponse;
    fromJSON(object: any): QueryGetObitResponse;
    toJSON(message: QueryGetObitResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetObitResponse>): QueryGetObitResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    Obit(request: QueryGetObitRequest): Promise<QueryGetObitResponse>;
    ObitAll(request: QueryAllObitRequest): Promise<QueryAllObitResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Obit(request: QueryGetObitRequest): Promise<QueryGetObitResponse>;
    ObitAll(request: QueryAllObitRequest): Promise<QueryAllObitResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
