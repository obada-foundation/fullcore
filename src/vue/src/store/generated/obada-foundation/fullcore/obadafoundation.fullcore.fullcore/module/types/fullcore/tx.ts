/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Obit } from "../fullcore/obit";

export const protobufPackage = "obadafoundation.fullcore.fullcore";

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

const baseMsgCreateObitResponse: object = {};

export const MsgCreateObitResponse = {
  encode(
    message: MsgCreateObitResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.Obit !== undefined) {
      Obit.encode(message.Obit, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateObitResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateObitResponse } as MsgCreateObitResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Obit = Obit.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateObitResponse {
    const message = { ...baseMsgCreateObitResponse } as MsgCreateObitResponse;
    if (object.Obit !== undefined && object.Obit !== null) {
      message.Obit = Obit.fromJSON(object.Obit);
    } else {
      message.Obit = undefined;
    }
    return message;
  },

  toJSON(message: MsgCreateObitResponse): unknown {
    const obj: any = {};
    message.Obit !== undefined &&
      (obj.Obit = message.Obit ? Obit.toJSON(message.Obit) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateObitResponse>
  ): MsgCreateObitResponse {
    const message = { ...baseMsgCreateObitResponse } as MsgCreateObitResponse;
    if (object.Obit !== undefined && object.Obit !== null) {
      message.Obit = Obit.fromPartial(object.Obit);
    } else {
      message.Obit = undefined;
    }
    return message;
  },
};

const baseMsgCreateObit: object = {
  serialNumberHash: "",
  manufacturer: "",
  partNumber: "",
  creator: "",
  owner: "",
  signature: "",
};

export const MsgCreateObit = {
  encode(message: MsgCreateObit, writer: Writer = Writer.create()): Writer {
    if (message.serialNumberHash !== "") {
      writer.uint32(10).string(message.serialNumberHash);
    }
    if (message.manufacturer !== "") {
      writer.uint32(18).string(message.manufacturer);
    }
    if (message.partNumber !== "") {
      writer.uint32(26).string(message.partNumber);
    }
    if (message.creator !== "") {
      writer.uint32(34).string(message.creator);
    }
    if (message.owner !== "") {
      writer.uint32(42).string(message.owner);
    }
    if (message.signature !== "") {
      writer.uint32(50).string(message.signature);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateObit {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateObit } as MsgCreateObit;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.serialNumberHash = reader.string();
          break;
        case 2:
          message.manufacturer = reader.string();
          break;
        case 3:
          message.partNumber = reader.string();
          break;
        case 4:
          message.creator = reader.string();
          break;
        case 5:
          message.owner = reader.string();
          break;
        case 6:
          message.signature = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateObit {
    const message = { ...baseMsgCreateObit } as MsgCreateObit;
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

  toJSON(message: MsgCreateObit): unknown {
    const obj: any = {};
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

  fromPartial(object: DeepPartial<MsgCreateObit>): MsgCreateObit {
    const message = { ...baseMsgCreateObit } as MsgCreateObit;
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

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateObit(request: MsgCreateObit): Promise<MsgCreateObitResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateObit(request: MsgCreateObit): Promise<MsgCreateObitResponse> {
    const data = MsgCreateObit.encode(request).finish();
    const promise = this.rpc.request(
      "obadafoundation.fullcore.fullcore.Msg",
      "CreateObit",
      data
    );
    return promise.then((data) =>
      MsgCreateObitResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
