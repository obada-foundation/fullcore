import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "obadafoundation.fullcore.fullcore";
export interface Obit {
    did: string;
    serialNumberHash: string;
    manufacturer: string;
    partNumber: string;
    creator: string;
    owner: string;
    signature: string;
}
export declare const Obit: {
    encode(message: Obit, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Obit;
    fromJSON(object: any): Obit;
    toJSON(message: Obit): unknown;
    fromPartial(object: DeepPartial<Obit>): Obit;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
