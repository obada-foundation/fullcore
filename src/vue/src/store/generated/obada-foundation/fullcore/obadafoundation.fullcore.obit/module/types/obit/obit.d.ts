import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "obadafoundation.fullcore.obit";
/** Params defines the parameters for the module. */
export interface Obit {
    did: string;
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
