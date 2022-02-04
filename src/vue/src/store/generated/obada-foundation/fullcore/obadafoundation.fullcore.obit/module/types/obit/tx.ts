/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "obadafoundation.fullcore.obit";

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

export interface MsgUpdateTaResponse {}

export interface MsgDeleteTa {
  creator: string;
  id: number;
}

export interface MsgDeleteTaResponse {}

const baseMsgMintObit: object = {
  creator: "",
  serialNumberHash: "",
  manufacturer: "",
  partNumber: "",
  trustAnchor: "",
  ownerDid: "",
};

export const MsgMintObit = {
  encode(message: MsgMintObit, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
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
    if (message.trustAnchor !== "") {
      writer.uint32(42).string(message.trustAnchor);
    }
    if (message.ownerDid !== "") {
      writer.uint32(50).string(message.ownerDid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgMintObit {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgMintObit } as MsgMintObit;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
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
          message.trustAnchor = reader.string();
          break;
        case 6:
          message.ownerDid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgMintObit {
    const message = { ...baseMsgMintObit } as MsgMintObit;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
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
    if (object.trustAnchor !== undefined && object.trustAnchor !== null) {
      message.trustAnchor = String(object.trustAnchor);
    } else {
      message.trustAnchor = "";
    }
    if (object.ownerDid !== undefined && object.ownerDid !== null) {
      message.ownerDid = String(object.ownerDid);
    } else {
      message.ownerDid = "";
    }
    return message;
  },

  toJSON(message: MsgMintObit): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.serialNumberHash !== undefined &&
      (obj.serialNumberHash = message.serialNumberHash);
    message.manufacturer !== undefined &&
      (obj.manufacturer = message.manufacturer);
    message.partNumber !== undefined && (obj.partNumber = message.partNumber);
    message.trustAnchor !== undefined &&
      (obj.trustAnchor = message.trustAnchor);
    message.ownerDid !== undefined && (obj.ownerDid = message.ownerDid);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgMintObit>): MsgMintObit {
    const message = { ...baseMsgMintObit } as MsgMintObit;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
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
    if (object.trustAnchor !== undefined && object.trustAnchor !== null) {
      message.trustAnchor = object.trustAnchor;
    } else {
      message.trustAnchor = "";
    }
    if (object.ownerDid !== undefined && object.ownerDid !== null) {
      message.ownerDid = object.ownerDid;
    } else {
      message.ownerDid = "";
    }
    return message;
  },
};

const baseMsgMintObitResponse: object = { did: "" };

export const MsgMintObitResponse = {
  encode(
    message: MsgMintObitResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgMintObitResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgMintObitResponse } as MsgMintObitResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgMintObitResponse {
    const message = { ...baseMsgMintObitResponse } as MsgMintObitResponse;
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    return message;
  },

  toJSON(message: MsgMintObitResponse): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgMintObitResponse>): MsgMintObitResponse {
    const message = { ...baseMsgMintObitResponse } as MsgMintObitResponse;
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    return message;
  },
};

const baseMsgCreateTa: object = { creator: "", title: "", pubkey: "" };

export const MsgCreateTa = {
  encode(message: MsgCreateTa, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    if (message.pubkey !== "") {
      writer.uint32(26).string(message.pubkey);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateTa {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateTa } as MsgCreateTa;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.title = reader.string();
          break;
        case 3:
          message.pubkey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateTa {
    const message = { ...baseMsgCreateTa } as MsgCreateTa;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = String(object.title);
    } else {
      message.title = "";
    }
    if (object.pubkey !== undefined && object.pubkey !== null) {
      message.pubkey = String(object.pubkey);
    } else {
      message.pubkey = "";
    }
    return message;
  },

  toJSON(message: MsgCreateTa): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.title !== undefined && (obj.title = message.title);
    message.pubkey !== undefined && (obj.pubkey = message.pubkey);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateTa>): MsgCreateTa {
    const message = { ...baseMsgCreateTa } as MsgCreateTa;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = object.title;
    } else {
      message.title = "";
    }
    if (object.pubkey !== undefined && object.pubkey !== null) {
      message.pubkey = object.pubkey;
    } else {
      message.pubkey = "";
    }
    return message;
  },
};

const baseMsgCreateTaResponse: object = { id: 0 };

export const MsgCreateTaResponse = {
  encode(
    message: MsgCreateTaResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateTaResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateTaResponse } as MsgCreateTaResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateTaResponse {
    const message = { ...baseMsgCreateTaResponse } as MsgCreateTaResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateTaResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateTaResponse>): MsgCreateTaResponse {
    const message = { ...baseMsgCreateTaResponse } as MsgCreateTaResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgUpdateTa: object = { creator: "", id: 0, title: "", pubkey: "" };

export const MsgUpdateTa = {
  encode(message: MsgUpdateTa, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.title !== "") {
      writer.uint32(26).string(message.title);
    }
    if (message.pubkey !== "") {
      writer.uint32(34).string(message.pubkey);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateTa {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateTa } as MsgUpdateTa;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.title = reader.string();
          break;
        case 4:
          message.pubkey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateTa {
    const message = { ...baseMsgUpdateTa } as MsgUpdateTa;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = String(object.title);
    } else {
      message.title = "";
    }
    if (object.pubkey !== undefined && object.pubkey !== null) {
      message.pubkey = String(object.pubkey);
    } else {
      message.pubkey = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateTa): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.title !== undefined && (obj.title = message.title);
    message.pubkey !== undefined && (obj.pubkey = message.pubkey);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateTa>): MsgUpdateTa {
    const message = { ...baseMsgUpdateTa } as MsgUpdateTa;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = object.title;
    } else {
      message.title = "";
    }
    if (object.pubkey !== undefined && object.pubkey !== null) {
      message.pubkey = object.pubkey;
    } else {
      message.pubkey = "";
    }
    return message;
  },
};

const baseMsgUpdateTaResponse: object = {};

export const MsgUpdateTaResponse = {
  encode(_: MsgUpdateTaResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateTaResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateTaResponse } as MsgUpdateTaResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateTaResponse {
    const message = { ...baseMsgUpdateTaResponse } as MsgUpdateTaResponse;
    return message;
  },

  toJSON(_: MsgUpdateTaResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgUpdateTaResponse>): MsgUpdateTaResponse {
    const message = { ...baseMsgUpdateTaResponse } as MsgUpdateTaResponse;
    return message;
  },
};

const baseMsgDeleteTa: object = { creator: "", id: 0 };

export const MsgDeleteTa = {
  encode(message: MsgDeleteTa, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteTa {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteTa } as MsgDeleteTa;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteTa {
    const message = { ...baseMsgDeleteTa } as MsgDeleteTa;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgDeleteTa): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteTa>): MsgDeleteTa {
    const message = { ...baseMsgDeleteTa } as MsgDeleteTa;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgDeleteTaResponse: object = {};

export const MsgDeleteTaResponse = {
  encode(_: MsgDeleteTaResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteTaResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteTaResponse } as MsgDeleteTaResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteTaResponse {
    const message = { ...baseMsgDeleteTaResponse } as MsgDeleteTaResponse;
    return message;
  },

  toJSON(_: MsgDeleteTaResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgDeleteTaResponse>): MsgDeleteTaResponse {
    const message = { ...baseMsgDeleteTaResponse } as MsgDeleteTaResponse;
    return message;
  },
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

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  MintObit(request: MsgMintObit): Promise<MsgMintObitResponse> {
    const data = MsgMintObit.encode(request).finish();
    const promise = this.rpc.request(
      "obadafoundation.fullcore.obit.Msg",
      "MintObit",
      data
    );
    return promise.then((data) => MsgMintObitResponse.decode(new Reader(data)));
  }

  CreateTa(request: MsgCreateTa): Promise<MsgCreateTaResponse> {
    const data = MsgCreateTa.encode(request).finish();
    const promise = this.rpc.request(
      "obadafoundation.fullcore.obit.Msg",
      "CreateTa",
      data
    );
    return promise.then((data) => MsgCreateTaResponse.decode(new Reader(data)));
  }

  UpdateTa(request: MsgUpdateTa): Promise<MsgUpdateTaResponse> {
    const data = MsgUpdateTa.encode(request).finish();
    const promise = this.rpc.request(
      "obadafoundation.fullcore.obit.Msg",
      "UpdateTa",
      data
    );
    return promise.then((data) => MsgUpdateTaResponse.decode(new Reader(data)));
  }

  DeleteTa(request: MsgDeleteTa): Promise<MsgDeleteTaResponse> {
    const data = MsgDeleteTa.encode(request).finish();
    const promise = this.rpc.request(
      "obadafoundation.fullcore.obit.Msg",
      "DeleteTa",
      data
    );
    return promise.then((data) => MsgDeleteTaResponse.decode(new Reader(data)));
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
