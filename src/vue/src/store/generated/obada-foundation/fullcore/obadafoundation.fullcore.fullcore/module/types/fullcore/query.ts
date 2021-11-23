/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { Obit } from "../fullcore/obit";

export const protobufPackage = "obadafoundation.fullcore.fullcore";

export interface QueryAllObitRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllObitResponse {
  Obit: Obit[];
  pagination: PageResponse | undefined;
}

export interface QueryGetObitRequest {
  did: string;
}

export interface QueryGetObitResponse {
  Obit: Obit | undefined;
}

const baseQueryAllObitRequest: object = {};

export const QueryAllObitRequest = {
  encode(
    message: QueryAllObitRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllObitRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllObitRequest } as QueryAllObitRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllObitRequest {
    const message = { ...baseQueryAllObitRequest } as QueryAllObitRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllObitRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllObitRequest>): QueryAllObitRequest {
    const message = { ...baseQueryAllObitRequest } as QueryAllObitRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllObitResponse: object = {};

export const QueryAllObitResponse = {
  encode(
    message: QueryAllObitResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.Obit) {
      Obit.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllObitResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllObitResponse } as QueryAllObitResponse;
    message.Obit = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Obit.push(Obit.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllObitResponse {
    const message = { ...baseQueryAllObitResponse } as QueryAllObitResponse;
    message.Obit = [];
    if (object.Obit !== undefined && object.Obit !== null) {
      for (const e of object.Obit) {
        message.Obit.push(Obit.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllObitResponse): unknown {
    const obj: any = {};
    if (message.Obit) {
      obj.Obit = message.Obit.map((e) => (e ? Obit.toJSON(e) : undefined));
    } else {
      obj.Obit = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllObitResponse>): QueryAllObitResponse {
    const message = { ...baseQueryAllObitResponse } as QueryAllObitResponse;
    message.Obit = [];
    if (object.Obit !== undefined && object.Obit !== null) {
      for (const e of object.Obit) {
        message.Obit.push(Obit.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetObitRequest: object = { did: "" };

export const QueryGetObitRequest = {
  encode(
    message: QueryGetObitRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetObitRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetObitRequest } as QueryGetObitRequest;
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

  fromJSON(object: any): QueryGetObitRequest {
    const message = { ...baseQueryGetObitRequest } as QueryGetObitRequest;
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    return message;
  },

  toJSON(message: QueryGetObitRequest): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetObitRequest>): QueryGetObitRequest {
    const message = { ...baseQueryGetObitRequest } as QueryGetObitRequest;
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    return message;
  },
};

const baseQueryGetObitResponse: object = {};

export const QueryGetObitResponse = {
  encode(
    message: QueryGetObitResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.Obit !== undefined) {
      Obit.encode(message.Obit, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetObitResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetObitResponse } as QueryGetObitResponse;
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

  fromJSON(object: any): QueryGetObitResponse {
    const message = { ...baseQueryGetObitResponse } as QueryGetObitResponse;
    if (object.Obit !== undefined && object.Obit !== null) {
      message.Obit = Obit.fromJSON(object.Obit);
    } else {
      message.Obit = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetObitResponse): unknown {
    const obj: any = {};
    message.Obit !== undefined &&
      (obj.Obit = message.Obit ? Obit.toJSON(message.Obit) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetObitResponse>): QueryGetObitResponse {
    const message = { ...baseQueryGetObitResponse } as QueryGetObitResponse;
    if (object.Obit !== undefined && object.Obit !== null) {
      message.Obit = Obit.fromPartial(object.Obit);
    } else {
      message.Obit = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  Obit(request: QueryGetObitRequest): Promise<QueryGetObitResponse>;
  ObitAll(request: QueryAllObitRequest): Promise<QueryAllObitResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Obit(request: QueryGetObitRequest): Promise<QueryGetObitResponse> {
    const data = QueryGetObitRequest.encode(request).finish();
    const promise = this.rpc.request(
      "obadafoundation.fullcore.fullcore.Query",
      "Obit",
      data
    );
    return promise.then((data) =>
      QueryGetObitResponse.decode(new Reader(data))
    );
  }

  ObitAll(request: QueryAllObitRequest): Promise<QueryAllObitResponse> {
    const data = QueryAllObitRequest.encode(request).finish();
    const promise = this.rpc.request(
      "obadafoundation.fullcore.fullcore.Query",
      "ObitAll",
      data
    );
    return promise.then((data) =>
      QueryAllObitResponse.decode(new Reader(data))
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
