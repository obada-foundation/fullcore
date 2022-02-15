import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../obit/params";
import { NFT } from "../obit/nft";
import { Ta } from "../obit/ta";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
export declare const protobufPackage = "obadafoundation.fullcore.obit";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
    /** params holds all the parameters of this module. */
    params: Params | undefined;
}
export interface QueryGetAllNftByOwnerRequest {
    owner: string;
}
export interface QueryGetAllNftByOwnerResponse {
    NFT: NFT[];
}
export interface QueryGetTaRequest {
    id: number;
}
export interface QueryGetTaResponse {
    Ta: Ta | undefined;
}
export interface QueryAllTaRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllTaResponse {
    Ta: Ta[];
    pagination: PageResponse | undefined;
}
export declare const QueryParamsRequest: {
    encode(_: QueryParamsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest;
    fromJSON(_: any): QueryParamsRequest;
    toJSON(_: QueryParamsRequest): unknown;
    fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest;
};
export declare const QueryParamsResponse: {
    encode(message: QueryParamsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse;
    fromJSON(object: any): QueryParamsResponse;
    toJSON(message: QueryParamsResponse): unknown;
    fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse;
};
export declare const QueryGetAllNftByOwnerRequest: {
    encode(message: QueryGetAllNftByOwnerRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetAllNftByOwnerRequest;
    fromJSON(object: any): QueryGetAllNftByOwnerRequest;
    toJSON(message: QueryGetAllNftByOwnerRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetAllNftByOwnerRequest>): QueryGetAllNftByOwnerRequest;
};
export declare const QueryGetAllNftByOwnerResponse: {
    encode(message: QueryGetAllNftByOwnerResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetAllNftByOwnerResponse;
    fromJSON(object: any): QueryGetAllNftByOwnerResponse;
    toJSON(message: QueryGetAllNftByOwnerResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetAllNftByOwnerResponse>): QueryGetAllNftByOwnerResponse;
};
export declare const QueryGetTaRequest: {
    encode(message: QueryGetTaRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetTaRequest;
    fromJSON(object: any): QueryGetTaRequest;
    toJSON(message: QueryGetTaRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetTaRequest>): QueryGetTaRequest;
};
export declare const QueryGetTaResponse: {
    encode(message: QueryGetTaResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetTaResponse;
    fromJSON(object: any): QueryGetTaResponse;
    toJSON(message: QueryGetTaResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetTaResponse>): QueryGetTaResponse;
};
export declare const QueryAllTaRequest: {
    encode(message: QueryAllTaRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllTaRequest;
    fromJSON(object: any): QueryAllTaRequest;
    toJSON(message: QueryAllTaRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllTaRequest>): QueryAllTaRequest;
};
export declare const QueryAllTaResponse: {
    encode(message: QueryAllTaResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllTaResponse;
    fromJSON(object: any): QueryAllTaResponse;
    toJSON(message: QueryAllTaResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllTaResponse>): QueryAllTaResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Parameters queries the parameters of the module. */
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    /** Queries a Ta by id. */
    Ta(request: QueryGetTaRequest): Promise<QueryGetTaResponse>;
    /** Queries a list of Ta items. */
    TaAll(request: QueryAllTaRequest): Promise<QueryAllTaResponse>;
    GetAllNftByOwner(request: QueryGetAllNftByOwnerRequest): Promise<QueryGetAllNftByOwnerResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    Ta(request: QueryGetTaRequest): Promise<QueryGetTaResponse>;
    TaAll(request: QueryAllTaRequest): Promise<QueryAllTaResponse>;
    GetAllNftByOwner(request: QueryGetAllNftByOwnerRequest): Promise<QueryGetAllNftByOwnerResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
