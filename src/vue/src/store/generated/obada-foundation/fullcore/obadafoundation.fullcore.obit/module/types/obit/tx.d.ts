import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "obadafoundation.fullcore.obit";
export interface MsgMintObit {
    creator: string;
    did: string;
}
export interface MsgMintObitResponse {
    did: string;
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
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    MintObit(request: MsgMintObit): Promise<MsgMintObitResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    MintObit(request: MsgMintObit): Promise<MsgMintObitResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
