import { Any } from "../google/protobuf/any";
import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "obadafoundation.fullcore.obit";
export interface NFT {
    classId: string;
    id: string;
    uri: string;
    uriHash: string;
    data: Any | undefined;
}
export interface NFTData {
    trustAnchorToken: string;
    usn: string;
    checksum: string;
    documents: NFTDocument[];
}
export interface NFTDocument {
    name: string;
    uri: string;
    hash: string;
}
export declare const NFT: {
    encode(message: NFT, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): NFT;
    fromJSON(object: any): NFT;
    toJSON(message: NFT): unknown;
    fromPartial(object: DeepPartial<NFT>): NFT;
};
export declare const NFTData: {
    encode(message: NFTData, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): NFTData;
    fromJSON(object: any): NFTData;
    toJSON(message: NFTData): unknown;
    fromPartial(object: DeepPartial<NFTData>): NFTData;
};
export declare const NFTDocument: {
    encode(message: NFTDocument, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): NFTDocument;
    fromJSON(object: any): NFTDocument;
    toJSON(message: NFTDocument): unknown;
    fromPartial(object: DeepPartial<NFTDocument>): NFTDocument;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
