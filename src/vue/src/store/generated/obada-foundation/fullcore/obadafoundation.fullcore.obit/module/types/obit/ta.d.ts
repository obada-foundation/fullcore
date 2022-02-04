import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "obadafoundation.fullcore.obit";
export interface Ta {
    id: number;
    title: string;
    pubkey: string;
    creator: string;
}
export declare const Ta: {
    encode(message: Ta, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Ta;
    fromJSON(object: any): Ta;
    toJSON(message: Ta): unknown;
    fromPartial(object: DeepPartial<Ta>): Ta;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
