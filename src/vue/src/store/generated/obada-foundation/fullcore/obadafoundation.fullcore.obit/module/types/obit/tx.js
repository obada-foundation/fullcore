/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { NFTDocument } from "../obit/nft";
export const protobufPackage = "obadafoundation.fullcore.obit";
const baseMsgMintObit = {
    creator: "",
    serialNumberHash: "",
    manufacturer: "",
    partNumber: "",
    trustAnchorToken: "",
    uri: "",
    uriHash: "",
};
export const MsgMintObit = {
    encode(message, writer = Writer.create()) {
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
        if (message.trustAnchorToken !== "") {
            writer.uint32(42).string(message.trustAnchorToken);
        }
        if (message.uri !== "") {
            writer.uint32(50).string(message.uri);
        }
        if (message.uriHash !== "") {
            writer.uint32(58).string(message.uriHash);
        }
        for (const v of message.documents) {
            NFTDocument.encode(v, writer.uint32(82).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgMintObit };
        message.documents = [];
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
                    message.trustAnchorToken = reader.string();
                    break;
                case 6:
                    message.uri = reader.string();
                    break;
                case 7:
                    message.uriHash = reader.string();
                    break;
                case 10:
                    message.documents.push(NFTDocument.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgMintObit };
        message.documents = [];
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.serialNumberHash !== undefined &&
            object.serialNumberHash !== null) {
            message.serialNumberHash = String(object.serialNumberHash);
        }
        else {
            message.serialNumberHash = "";
        }
        if (object.manufacturer !== undefined && object.manufacturer !== null) {
            message.manufacturer = String(object.manufacturer);
        }
        else {
            message.manufacturer = "";
        }
        if (object.partNumber !== undefined && object.partNumber !== null) {
            message.partNumber = String(object.partNumber);
        }
        else {
            message.partNumber = "";
        }
        if (object.trustAnchorToken !== undefined &&
            object.trustAnchorToken !== null) {
            message.trustAnchorToken = String(object.trustAnchorToken);
        }
        else {
            message.trustAnchorToken = "";
        }
        if (object.uri !== undefined && object.uri !== null) {
            message.uri = String(object.uri);
        }
        else {
            message.uri = "";
        }
        if (object.uriHash !== undefined && object.uriHash !== null) {
            message.uriHash = String(object.uriHash);
        }
        else {
            message.uriHash = "";
        }
        if (object.documents !== undefined && object.documents !== null) {
            for (const e of object.documents) {
                message.documents.push(NFTDocument.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.serialNumberHash !== undefined &&
            (obj.serialNumberHash = message.serialNumberHash);
        message.manufacturer !== undefined &&
            (obj.manufacturer = message.manufacturer);
        message.partNumber !== undefined && (obj.partNumber = message.partNumber);
        message.trustAnchorToken !== undefined &&
            (obj.trustAnchorToken = message.trustAnchorToken);
        message.uri !== undefined && (obj.uri = message.uri);
        message.uriHash !== undefined && (obj.uriHash = message.uriHash);
        if (message.documents) {
            obj.documents = message.documents.map((e) => e ? NFTDocument.toJSON(e) : undefined);
        }
        else {
            obj.documents = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgMintObit };
        message.documents = [];
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.serialNumberHash !== undefined &&
            object.serialNumberHash !== null) {
            message.serialNumberHash = object.serialNumberHash;
        }
        else {
            message.serialNumberHash = "";
        }
        if (object.manufacturer !== undefined && object.manufacturer !== null) {
            message.manufacturer = object.manufacturer;
        }
        else {
            message.manufacturer = "";
        }
        if (object.partNumber !== undefined && object.partNumber !== null) {
            message.partNumber = object.partNumber;
        }
        else {
            message.partNumber = "";
        }
        if (object.trustAnchorToken !== undefined &&
            object.trustAnchorToken !== null) {
            message.trustAnchorToken = object.trustAnchorToken;
        }
        else {
            message.trustAnchorToken = "";
        }
        if (object.uri !== undefined && object.uri !== null) {
            message.uri = object.uri;
        }
        else {
            message.uri = "";
        }
        if (object.uriHash !== undefined && object.uriHash !== null) {
            message.uriHash = object.uriHash;
        }
        else {
            message.uriHash = "";
        }
        if (object.documents !== undefined && object.documents !== null) {
            for (const e of object.documents) {
                message.documents.push(NFTDocument.fromPartial(e));
            }
        }
        return message;
    },
};
const baseMsgMintObitResponse = { did: "" };
export const MsgMintObitResponse = {
    encode(message, writer = Writer.create()) {
        if (message.did !== "") {
            writer.uint32(10).string(message.did);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgMintObitResponse };
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
        const message = { ...baseMsgMintObitResponse };
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
        const message = { ...baseMsgMintObitResponse };
        if (object.did !== undefined && object.did !== null) {
            message.did = object.did;
        }
        else {
            message.did = "";
        }
        return message;
    },
};
const baseMsgCreateTa = { creator: "", title: "", pubkey: "" };
export const MsgCreateTa = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateTa };
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
    fromJSON(object) {
        const message = { ...baseMsgCreateTa };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.title !== undefined && object.title !== null) {
            message.title = String(object.title);
        }
        else {
            message.title = "";
        }
        if (object.pubkey !== undefined && object.pubkey !== null) {
            message.pubkey = String(object.pubkey);
        }
        else {
            message.pubkey = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.title !== undefined && (obj.title = message.title);
        message.pubkey !== undefined && (obj.pubkey = message.pubkey);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCreateTa };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.title !== undefined && object.title !== null) {
            message.title = object.title;
        }
        else {
            message.title = "";
        }
        if (object.pubkey !== undefined && object.pubkey !== null) {
            message.pubkey = object.pubkey;
        }
        else {
            message.pubkey = "";
        }
        return message;
    },
};
const baseMsgCreateTaResponse = { id: 0 };
export const MsgCreateTaResponse = {
    encode(message, writer = Writer.create()) {
        if (message.id !== 0) {
            writer.uint32(8).uint64(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateTaResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgCreateTaResponse };
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCreateTaResponse };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        return message;
    },
};
const baseMsgUpdateTa = { creator: "", id: 0, title: "", pubkey: "" };
export const MsgUpdateTa = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateTa };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.id = longToNumber(reader.uint64());
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
    fromJSON(object) {
        const message = { ...baseMsgUpdateTa };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        if (object.title !== undefined && object.title !== null) {
            message.title = String(object.title);
        }
        else {
            message.title = "";
        }
        if (object.pubkey !== undefined && object.pubkey !== null) {
            message.pubkey = String(object.pubkey);
        }
        else {
            message.pubkey = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.id !== undefined && (obj.id = message.id);
        message.title !== undefined && (obj.title = message.title);
        message.pubkey !== undefined && (obj.pubkey = message.pubkey);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgUpdateTa };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        if (object.title !== undefined && object.title !== null) {
            message.title = object.title;
        }
        else {
            message.title = "";
        }
        if (object.pubkey !== undefined && object.pubkey !== null) {
            message.pubkey = object.pubkey;
        }
        else {
            message.pubkey = "";
        }
        return message;
    },
};
const baseMsgUpdateTaResponse = {};
export const MsgUpdateTaResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateTaResponse };
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
    fromJSON(_) {
        const message = { ...baseMsgUpdateTaResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgUpdateTaResponse };
        return message;
    },
};
const baseMsgDeleteTa = { creator: "", id: 0 };
export const MsgDeleteTa = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.id !== 0) {
            writer.uint32(16).uint64(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgDeleteTa };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.id = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgDeleteTa };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.id !== undefined && (obj.id = message.id);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgDeleteTa };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        return message;
    },
};
const baseMsgDeleteTaResponse = {};
export const MsgDeleteTaResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgDeleteTaResponse };
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
    fromJSON(_) {
        const message = { ...baseMsgDeleteTaResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgDeleteTaResponse };
        return message;
    },
};
const baseMsgSend = { did: "", sender: "", receiver: "" };
export const MsgSend = {
    encode(message, writer = Writer.create()) {
        if (message.did !== "") {
            writer.uint32(10).string(message.did);
        }
        if (message.sender !== "") {
            writer.uint32(18).string(message.sender);
        }
        if (message.receiver !== "") {
            writer.uint32(26).string(message.receiver);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSend };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.did = reader.string();
                    break;
                case 2:
                    message.sender = reader.string();
                    break;
                case 3:
                    message.receiver = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgSend };
        if (object.did !== undefined && object.did !== null) {
            message.did = String(object.did);
        }
        else {
            message.did = "";
        }
        if (object.sender !== undefined && object.sender !== null) {
            message.sender = String(object.sender);
        }
        else {
            message.sender = "";
        }
        if (object.receiver !== undefined && object.receiver !== null) {
            message.receiver = String(object.receiver);
        }
        else {
            message.receiver = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.did !== undefined && (obj.did = message.did);
        message.sender !== undefined && (obj.sender = message.sender);
        message.receiver !== undefined && (obj.receiver = message.receiver);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgSend };
        if (object.did !== undefined && object.did !== null) {
            message.did = object.did;
        }
        else {
            message.did = "";
        }
        if (object.sender !== undefined && object.sender !== null) {
            message.sender = object.sender;
        }
        else {
            message.sender = "";
        }
        if (object.receiver !== undefined && object.receiver !== null) {
            message.receiver = object.receiver;
        }
        else {
            message.receiver = "";
        }
        return message;
    },
};
const baseMsgSendResponse = {};
export const MsgSendResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSendResponse };
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
    fromJSON(_) {
        const message = { ...baseMsgSendResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgSendResponse };
        return message;
    },
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    MintObit(request) {
        const data = MsgMintObit.encode(request).finish();
        const promise = this.rpc.request("obadafoundation.fullcore.obit.Msg", "MintObit", data);
        return promise.then((data) => MsgMintObitResponse.decode(new Reader(data)));
    }
    Send(request) {
        const data = MsgSend.encode(request).finish();
        const promise = this.rpc.request("obadafoundation.fullcore.obit.Msg", "Send", data);
        return promise.then((data) => MsgSendResponse.decode(new Reader(data)));
    }
    CreateTa(request) {
        const data = MsgCreateTa.encode(request).finish();
        const promise = this.rpc.request("obadafoundation.fullcore.obit.Msg", "CreateTa", data);
        return promise.then((data) => MsgCreateTaResponse.decode(new Reader(data)));
    }
    UpdateTa(request) {
        const data = MsgUpdateTa.encode(request).finish();
        const promise = this.rpc.request("obadafoundation.fullcore.obit.Msg", "UpdateTa", data);
        return promise.then((data) => MsgUpdateTaResponse.decode(new Reader(data)));
    }
    DeleteTa(request) {
        const data = MsgDeleteTa.encode(request).finish();
        const promise = this.rpc.request("obadafoundation.fullcore.obit.Msg", "DeleteTa", data);
        return promise.then((data) => MsgDeleteTaResponse.decode(new Reader(data)));
    }
}
var globalThis = (() => {
    if (typeof globalThis !== "undefined")
        return globalThis;
    if (typeof self !== "undefined")
        return self;
    if (typeof window !== "undefined")
        return window;
    if (typeof global !== "undefined")
        return global;
    throw "Unable to locate global object";
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
