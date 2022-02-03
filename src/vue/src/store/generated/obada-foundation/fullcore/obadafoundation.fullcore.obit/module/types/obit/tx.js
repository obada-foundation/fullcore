/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
export const protobufPackage = "obadafoundation.fullcore.obit";
const baseMsgMintObit = { creator: "", did: "" };
export const MsgMintObit = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.did !== "") {
            writer.uint32(18).string(message.did);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgMintObit };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
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
        const message = { ...baseMsgMintObit };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
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
        message.creator !== undefined && (obj.creator = message.creator);
        message.did !== undefined && (obj.did = message.did);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgMintObit };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = object.did;
        }
        else {
            message.did = "";
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
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    MintObit(request) {
        const data = MsgMintObit.encode(request).finish();
        const promise = this.rpc.request("obadafoundation.fullcore.obit.Msg", "MintObit", data);
        return promise.then((data) => MsgMintObitResponse.decode(new Reader(data)));
    }
}
