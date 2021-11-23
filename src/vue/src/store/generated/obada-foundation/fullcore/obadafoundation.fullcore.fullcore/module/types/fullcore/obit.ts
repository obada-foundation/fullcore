/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "obadafoundation.fullcore.fullcore";

export interface Obit {
  did: string;
  serialNumberHash: string;
  manufacturer: string;
  partNumber: string;
  creator: string;
  owner: string;
  signature: string;
}

const baseObit: object = {
  did: "",
  serialNumberHash: "",
  manufacturer: "",
  partNumber: "",
  creator: "",
  owner: "",
  signature: "",
};

export const Obit = {
  encode(message: Obit, writer: Writer = Writer.create()): Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    if (message.serialNumberHash !== "") {
      writer.uint32(18).string(message.serialNumberHash);
    }
    if (message.manufacturer !== "") {
      writer.uint32(26).string(message.manufacturer);
    }
    if (message.partNumber !== "") {
      writer.uint32(34).string(message.partNumber);
    }
    if (message.creator !== "") {
      writer.uint32(42).string(message.creator);
    }
    if (message.owner !== "") {
      writer.uint32(50).string(message.owner);
    }
    if (message.signature !== "") {
      writer.uint32(58).string(message.signature);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Obit {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseObit } as Obit;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        case 2:
          message.serialNumberHash = reader.string();
          break;
        case 3:
          message.manufacturer = reader.string();
          break;
        case 4:
          message.partNumber = reader.string();
          break;
        case 5:
          message.creator = reader.string();
          break;
        case 6:
          message.owner = reader.string();
          break;
        case 7:
          message.signature = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Obit {
    const message = { ...baseObit } as Obit;
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (
      object.serialNumberHash !== undefined &&
      object.serialNumberHash !== null
    ) {
      message.serialNumberHash = String(object.serialNumberHash);
    } else {
      message.serialNumberHash = "";
    }
    if (object.manufacturer !== undefined && object.manufacturer !== null) {
      message.manufacturer = String(object.manufacturer);
    } else {
      message.manufacturer = "";
    }
    if (object.partNumber !== undefined && object.partNumber !== null) {
      message.partNumber = String(object.partNumber);
    } else {
      message.partNumber = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = String(object.signature);
    } else {
      message.signature = "";
    }
    return message;
  },

  toJSON(message: Obit): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    message.serialNumberHash !== undefined &&
      (obj.serialNumberHash = message.serialNumberHash);
    message.manufacturer !== undefined &&
      (obj.manufacturer = message.manufacturer);
    message.partNumber !== undefined && (obj.partNumber = message.partNumber);
    message.creator !== undefined && (obj.creator = message.creator);
    message.owner !== undefined && (obj.owner = message.owner);
    message.signature !== undefined && (obj.signature = message.signature);
    return obj;
  },

  fromPartial(object: DeepPartial<Obit>): Obit {
    const message = { ...baseObit } as Obit;
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (
      object.serialNumberHash !== undefined &&
      object.serialNumberHash !== null
    ) {
      message.serialNumberHash = object.serialNumberHash;
    } else {
      message.serialNumberHash = "";
    }
    if (object.manufacturer !== undefined && object.manufacturer !== null) {
      message.manufacturer = object.manufacturer;
    } else {
      message.manufacturer = "";
    }
    if (object.partNumber !== undefined && object.partNumber !== null) {
      message.partNumber = object.partNumber;
    } else {
      message.partNumber = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = object.signature;
    } else {
      message.signature = "";
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
