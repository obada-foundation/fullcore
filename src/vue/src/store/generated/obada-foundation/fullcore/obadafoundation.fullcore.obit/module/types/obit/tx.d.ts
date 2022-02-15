import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "obadafoundation.fullcore.obit";
/** MsgMintObit handles obit information */
export interface MsgMintObit {
    creator: string;
    serialNumberHash: string;
    manufacturer: string;
    partNumber: string;
    obdDid: string;
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
/** MsgSend represents a message to send a nft from one account to another account. */
export interface MsgSend {
    /** id defines the unique identification of nft */
    did: string;
    /** sender is the address of the owner of nft */
    sender: string;
    /** receiver is the receiver address of nft */
    receiver: string;
}
/** MsgSendResponse defines the Msg/Send response type. */
export interface MsgSendResponse {
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
export declare const MsgSend: {
    encode(message: MsgSend, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSend;
    fromJSON(object: any): MsgSend;
    toJSON(message: MsgSend): unknown;
    fromPartial(object: DeepPartial<MsgSend>): MsgSend;
};
export declare const MsgSendResponse: {
    encode(_: MsgSendResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSendResponse;
    fromJSON(_: any): MsgSendResponse;
    toJSON(_: MsgSendResponse): unknown;
    fromPartial(_: DeepPartial<MsgSendResponse>): MsgSendResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** MintObit mint a new OBADA NFT (Obit) */
    MintObit(request: MsgMintObit): Promise<MsgMintObitResponse>;
    Send(request: MsgSend): Promise<MsgSendResponse>;
    CreateTa(request: MsgCreateTa): Promise<MsgCreateTaResponse>;
    UpdateTa(request: MsgUpdateTa): Promise<MsgUpdateTaResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpcxf */
    DeleteTa(request: MsgDeleteTa): Promise<MsgDeleteTaResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    MintObit(request: MsgMintObit): Promise<MsgMintObitResponse>;
    Send(request: MsgSend): Promise<MsgSendResponse>;
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
