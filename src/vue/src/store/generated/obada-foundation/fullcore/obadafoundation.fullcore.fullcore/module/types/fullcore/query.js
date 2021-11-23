/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { PageRequest, PageResponse, } from "../cosmos/base/query/v1beta1/pagination";
import { Obit } from "../fullcore/obit";
export const protobufPackage = "obadafoundation.fullcore.fullcore";
const baseQueryAllObitRequest = {};
export const QueryAllObitRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllObitRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryAllObitRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageRequest.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllObitRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryAllObitResponse = {};
export const QueryAllObitResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.Obit) {
            Obit.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllObitResponse };
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
    fromJSON(object) {
        const message = { ...baseQueryAllObitResponse };
        message.Obit = [];
        if (object.Obit !== undefined && object.Obit !== null) {
            for (const e of object.Obit) {
                message.Obit.push(Obit.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.Obit) {
            obj.Obit = message.Obit.map((e) => (e ? Obit.toJSON(e) : undefined));
        }
        else {
            obj.Obit = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllObitResponse };
        message.Obit = [];
        if (object.Obit !== undefined && object.Obit !== null) {
            for (const e of object.Obit) {
                message.Obit.push(Obit.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryGetObitRequest = { did: "" };
export const QueryGetObitRequest = {
    encode(message, writer = Writer.create()) {
        if (message.did !== "") {
            writer.uint32(10).string(message.did);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetObitRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryGetObitRequest };
        if (object.did !== undefined && object.did !== null) {
            message.did = String(object.did);
        }
        else {
            message.did = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.did !== undefined && (obj.did = message.did);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetObitRequest };
        if (object.did !== undefined && object.did !== null) {
            message.did = object.did;
        }
        else {
            message.did = "";
        }
        return message;
    },
};
const baseQueryGetObitResponse = {};
export const QueryGetObitResponse = {
    encode(message, writer = Writer.create()) {
        if (message.Obit !== undefined) {
            Obit.encode(message.Obit, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetObitResponse };
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
    fromJSON(object) {
        const message = { ...baseQueryGetObitResponse };
        if (object.Obit !== undefined && object.Obit !== null) {
            message.Obit = Obit.fromJSON(object.Obit);
        }
        else {
            message.Obit = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.Obit !== undefined &&
            (obj.Obit = message.Obit ? Obit.toJSON(message.Obit) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetObitResponse };
        if (object.Obit !== undefined && object.Obit !== null) {
            message.Obit = Obit.fromPartial(object.Obit);
        }
        else {
            message.Obit = undefined;
        }
        return message;
    },
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    Obit(request) {
        const data = QueryGetObitRequest.encode(request).finish();
        const promise = this.rpc.request("obadafoundation.fullcore.fullcore.Query", "Obit", data);
        return promise.then((data) => QueryGetObitResponse.decode(new Reader(data)));
    }
    ObitAll(request) {
        const data = QueryAllObitRequest.encode(request).finish();
        const promise = this.rpc.request("obadafoundation.fullcore.fullcore.Query", "ObitAll", data);
        return promise.then((data) => QueryAllObitResponse.decode(new Reader(data)));
    }
}
