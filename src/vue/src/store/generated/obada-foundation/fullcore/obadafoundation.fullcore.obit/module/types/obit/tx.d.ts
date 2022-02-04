import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "obadafoundation.fullcore.obit";
/** MsgMintObit handles obit information */
export interface MsgMintObit {
    creator: string;
    serialNumberHash: string;
    manufacturer: string;
    partNumber: string;
    trustAnchor: string;
    ownerDid: string;
}
/** MsgMintObitResponse success minting response */
export interface MsgMintObitResponse {
    did: string;
}
export interface MsgCreateTa {
    creator: string;
    title: string;
    pubkey: string;
}
export interface MsgCreateTaResponse {
    id: number;
}
export interface MsgUpdateTa {
    creator: string;
    id: number;
    title: string;
    pubkey: string;
}
export interface MsgUpdateTaResponse {
}
export interface MsgDeleteTa {
    creator: string;
    id: number;
}
export interface MsgDeleteTaResponse {
}
export declare const MsgMintObit: {
    encode(message: MsgMintObit, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgMintObit;
    fromJSON(object: any): MsgMintObit;
    toJSON(message: MsgMintObit): unknown;
    fromPartial(object: DeepPartial<MsgMintObit>): MsgMintObit;
};
export declare const MsgMintObitResponse: {
    encode(message: MsgMintObitResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgMintObitResponse;
    fromJSON(object: any): MsgMintObitResponse;
    toJSON(message: MsgMintObitResponse): unknown;
    fromPartial(object: DeepPartial<MsgMintObitResponse>): MsgMintObitResponse;
};
export declare const MsgCreateTa: {
    encode(message: MsgCreateTa, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateTa;
    fromJSON(object: any): MsgCreateTa;
    toJSON(message: MsgCreateTa): unknown;
    fromPartial(object: DeepPartial<MsgCreateTa>): MsgCreateTa;
};
export declare const MsgCreateTaResponse: {
    encode(message: MsgCreateTaResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateTaResponse;
    fromJSON(object: any): MsgCreateTaResponse;
    toJSON(message: MsgCreateTaResponse): unknown;
    fromPartial(object: DeepPartial<MsgCreateTaResponse>): MsgCreateTaResponse;
};
export declare const MsgUpdateTa: {
    encode(message: MsgUpdateTa, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateTa;
    fromJSON(object: any): MsgUpdateTa;
    toJSON(message: MsgUpdateTa): unknown;
    fromPartial(object: DeepPartial<MsgUpdateTa>): MsgUpdateTa;
};
export declare const MsgUpdateTaResponse: {
    encode(_: MsgUpdateTaResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateTaResponse;
    fromJSON(_: any): MsgUpdateTaResponse;
    toJSON(_: MsgUpdateTaResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateTaResponse>): MsgUpdateTaResponse;
};
export declare const MsgDeleteTa: {
    encode(message: MsgDeleteTa, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteTa;
    fromJSON(object: any): MsgDeleteTa;
    toJSON(message: MsgDeleteTa): unknown;
    fromPartial(object: DeepPartial<MsgDeleteTa>): MsgDeleteTa;
};
export declare const MsgDeleteTaResponse: {
    encode(_: MsgDeleteTaResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteTaResponse;
    fromJSON(_: any): MsgDeleteTaResponse;
    toJSON(_: MsgDeleteTaResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteTaResponse>): MsgDeleteTaResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** MintObit mint a new OBADA NFT (Obit) */
    MintObit(request: MsgMintObit): Promise<MsgMintObitResponse>;
    CreateTa(request: MsgCreateTa): Promise<MsgCreateTaResponse>;
    UpdateTa(request: MsgUpdateTa): Promise<MsgUpdateTaResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    DeleteTa(request: MsgDeleteTa): Promise<MsgDeleteTaResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    MintObit(request: MsgMintObit): Promise<MsgMintObitResponse>;
    CreateTa(request: MsgCreateTa): Promise<MsgCreateTaResponse>;
    UpdateTa(request: MsgUpdateTa): Promise<MsgUpdateTaResponse>;
    DeleteTa(request: MsgDeleteTa): Promise<MsgDeleteTaResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
