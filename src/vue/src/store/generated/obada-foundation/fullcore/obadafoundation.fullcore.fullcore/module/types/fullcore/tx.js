/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Obit } from "../fullcore/obit";
export const protobufPackage = "obadafoundation.fullcore.fullcore";
const baseMsgCreateObitResponse = {};
export const MsgCreateObitResponse = {
    encode(message, writer = Writer.create()) {
        if (message.Obit !== undefined) {
            Obit.encode(message.Obit, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateObitResponse };
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
        const message = { ...baseMsgCreateObitResponse };
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
        const message = { ...baseMsgCreateObitResponse };
        if (object.Obit !== undefined && object.Obit !== null) {
            message.Obit = Obit.fromPartial(object.Obit);
        }
        else {
            message.Obit = undefined;
        }
        return message;
    },
};
const baseMsgCreateObit = {
    serialNumberHash: "",
    manufacturer: "",
    partNumber: "",
    creator: "",
    owner: "",
    signature: "",
};
export const MsgCreateObit = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateObit };
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
    fromJSON(object) {
        const message = { ...baseMsgCreateObit };
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
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = String(object.owner);
        }
        else {
            message.owner = "";
        }
        if (object.signature !== undefined && object.signature !== null) {
            message.signature = String(object.signature);
        }
        else {
            message.signature = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
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
    fromPartial(object) {
        const message = { ...baseMsgCreateObit };
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
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = object.owner;
        }
        else {
            message.owner = "";
        }
        if (object.signature !== undefined && object.signature !== null) {
            message.signature = object.signature;
        }
        else {
            message.signature = "";
        }
        return message;
    },
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    CreateObit(request) {
        const data = MsgCreateObit.encode(request).finish();
        const promise = this.rpc.request("obadafoundation.fullcore.fullcore.Msg", "CreateObit", data);
        return promise.then((data) => MsgCreateObitResponse.decode(new Reader(data)));
    }
}
