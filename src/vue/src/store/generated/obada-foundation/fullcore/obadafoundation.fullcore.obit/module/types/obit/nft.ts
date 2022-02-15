/* eslint-disable */
import { Any } from "../google/protobuf/any";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "obadafoundation.fullcore.obit";

export interface NFT {
  classId: string;
  id: string;
  uri: string;
  uriHash: string;
  data: Any | undefined;
}

export interface NFTData {
  ownerDid: string;
  obdDid: string;
  usn: string;
  rootHash: string;
}

const baseNFT: object = { classId: "", id: "", uri: "", uriHash: "" };

export const NFT = {
  encode(message: NFT, writer: Writer = Writer.create()): Writer {
    if (message.classId !== "") {
      writer.uint32(10).string(message.classId);
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
    }
    if (message.uri !== "") {
      writer.uint32(26).string(message.uri);
    }
    if (message.uriHash !== "") {
      writer.uint32(34).string(message.uriHash);
    }
    if (message.data !== undefined) {
      Any.encode(message.data, writer.uint32(82).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): NFT {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNFT } as NFT;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.classId = reader.string();
          break;
        case 2:
          message.id = reader.string();
          break;
        case 3:
          message.uri = reader.string();
          break;
        case 4:
          message.uriHash = reader.string();
          break;
        case 10:
          message.data = Any.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NFT {
    const message = { ...baseNFT } as NFT;
    if (object.classId !== undefined && object.classId !== null) {
      message.classId = String(object.classId);
    } else {
      message.classId = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.uri !== undefined && object.uri !== null) {
      message.uri = String(object.uri);
    } else {
      message.uri = "";
    }
    if (object.uriHash !== undefined && object.uriHash !== null) {
      message.uriHash = String(object.uriHash);
    } else {
      message.uriHash = "";
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = Any.fromJSON(object.data);
    } else {
      message.data = undefined;
    }
    return message;
  },

  toJSON(message: NFT): unknown {
    const obj: any = {};
    message.classId !== undefined && (obj.classId = message.classId);
    message.id !== undefined && (obj.id = message.id);
    message.uri !== undefined && (obj.uri = message.uri);
    message.uriHash !== undefined && (obj.uriHash = message.uriHash);
    message.data !== undefined &&
      (obj.data = message.data ? Any.toJSON(message.data) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<NFT>): NFT {
    const message = { ...baseNFT } as NFT;
    if (object.classId !== undefined && object.classId !== null) {
      message.classId = object.classId;
    } else {
      message.classId = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.uri !== undefined && object.uri !== null) {
      message.uri = object.uri;
    } else {
      message.uri = "";
    }
    if (object.uriHash !== undefined && object.uriHash !== null) {
      message.uriHash = object.uriHash;
    } else {
      message.uriHash = "";
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = Any.fromPartial(object.data);
    } else {
      message.data = undefined;
    }
    return message;
  },
};

const baseNFTData: object = { ownerDid: "", obdDid: "", usn: "", rootHash: "" };

export const NFTData = {
  encode(message: NFTData, writer: Writer = Writer.create()): Writer {
    if (message.ownerDid !== "") {
      writer.uint32(10).string(message.ownerDid);
    }
    if (message.obdDid !== "") {
      writer.uint32(18).string(message.obdDid);
    }
    if (message.usn !== "") {
      writer.uint32(26).string(message.usn);
    }
    if (message.rootHash !== "") {
      writer.uint32(34).string(message.rootHash);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): NFTData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNFTData } as NFTData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ownerDid = reader.string();
          break;
        case 2:
          message.obdDid = reader.string();
          break;
        case 3:
          message.usn = reader.string();
          break;
        case 4:
          message.rootHash = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NFTData {
    const message = { ...baseNFTData } as NFTData;
    if (object.ownerDid !== undefined && object.ownerDid !== null) {
      message.ownerDid = String(object.ownerDid);
    } else {
      message.ownerDid = "";
    }
    if (object.obdDid !== undefined && object.obdDid !== null) {
      message.obdDid = String(object.obdDid);
    } else {
      message.obdDid = "";
    }
    if (object.usn !== undefined && object.usn !== null) {
      message.usn = String(object.usn);
    } else {
      message.usn = "";
    }
    if (object.rootHash !== undefined && object.rootHash !== null) {
      message.rootHash = String(object.rootHash);
    } else {
      message.rootHash = "";
    }
    return message;
  },

  toJSON(message: NFTData): unknown {
    const obj: any = {};
    message.ownerDid !== undefined && (obj.ownerDid = message.ownerDid);
    message.obdDid !== undefined && (obj.obdDid = message.obdDid);
    message.usn !== undefined && (obj.usn = message.usn);
    message.rootHash !== undefined && (obj.rootHash = message.rootHash);
    return obj;
  },

  fromPartial(object: DeepPartial<NFTData>): NFTData {
    const message = { ...baseNFTData } as NFTData;
    if (object.ownerDid !== undefined && object.ownerDid !== null) {
      message.ownerDid = object.ownerDid;
    } else {
      message.ownerDid = "";
    }
    if (object.obdDid !== undefined && object.obdDid !== null) {
      message.obdDid = object.obdDid;
    } else {
      message.obdDid = "";
    }
    if (object.usn !== undefined && object.usn !== null) {
      message.usn = object.usn;
    } else {
      message.usn = "";
    }
    if (object.rootHash !== undefined && object.rootHash !== null) {
      message.rootHash = object.rootHash;
    } else {
      message.rootHash = "";
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
