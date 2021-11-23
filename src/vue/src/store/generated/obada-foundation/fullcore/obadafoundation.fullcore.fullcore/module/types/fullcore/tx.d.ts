import { Reader, Writer } from "protobufjs/minimal";
import { Obit } from "../fullcore/obit";
export declare const protobufPackage = "obadafoundation.fullcore.fullcore";
export interface MsgCreateObitResponse {
    Obit: Obit | undefined;
}
export interface MsgCreateObit {
    serialNumberHash: string;
    manufacturer: string;
    partNumber: string;
    creator: string;
    owner: string;
    signature: string;
}
export declare const MsgCreateObitResponse: {
    encode(message: MsgCreateObitResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateObitResponse;
    fromJSON(object: any): MsgCreateObitResponse;
    toJSON(message: MsgCreateObitResponse): unknown;
    fromPartial(object: DeepPartial<MsgCreateObitResponse>): MsgCreateObitResponse;
};
export declare const MsgCreateObit: {
    encode(message: MsgCreateObit, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateObit;
    fromJSON(object: any): MsgCreateObit;
    toJSON(message: MsgCreateObit): unknown;
    fromPartial(object: DeepPartial<MsgCreateObit>): MsgCreateObit;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    CreateObit(request: MsgCreateObit): Promise<MsgCreateObitResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateObit(request: MsgCreateObit): Promise<MsgCreateObitResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
