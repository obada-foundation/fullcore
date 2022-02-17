/* eslint-disable */
import { Any } from "../google/protobuf/any";
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "obadafoundation.fullcore.obit";
const baseNFT = { classId: "", id: "", uri: "", uriHash: "" };
export const NFT = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseNFT };
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
    fromJSON(object) {
        const message = { ...baseNFT };
        if (object.classId !== undefined && object.classId !== null) {
            message.classId = String(object.classId);
        }
        else {
            message.classId = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = "";
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
        if (object.data !== undefined && object.data !== null) {
            message.data = Any.fromJSON(object.data);
        }
        else {
            message.data = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.classId !== undefined && (obj.classId = message.classId);
        message.id !== undefined && (obj.id = message.id);
        message.uri !== undefined && (obj.uri = message.uri);
        message.uriHash !== undefined && (obj.uriHash = message.uriHash);
        message.data !== undefined &&
            (obj.data = message.data ? Any.toJSON(message.data) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseNFT };
        if (object.classId !== undefined && object.classId !== null) {
            message.classId = object.classId;
        }
        else {
            message.classId = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = "";
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
        if (object.data !== undefined && object.data !== null) {
            message.data = Any.fromPartial(object.data);
        }
        else {
            message.data = undefined;
        }
        return message;
    },
};
const baseNFTData = { trustAnchorToken: "", usn: "", checksum: "" };
export const NFTData = {
    encode(message, writer = Writer.create()) {
        if (message.trustAnchorToken !== "") {
            writer.uint32(10).string(message.trustAnchorToken);
        }
        if (message.usn !== "") {
            writer.uint32(18).string(message.usn);
        }
        if (message.checksum !== "") {
            writer.uint32(26).string(message.checksum);
        }
        for (const v of message.documents) {
            NFTDocument.encode(v, writer.uint32(34).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseNFTData };
        message.documents = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.trustAnchorToken = reader.string();
                    break;
                case 2:
                    message.usn = reader.string();
                    break;
                case 3:
                    message.checksum = reader.string();
                    break;
                case 4:
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
        const message = { ...baseNFTData };
        message.documents = [];
        if (object.trustAnchorToken !== undefined &&
            object.trustAnchorToken !== null) {
            message.trustAnchorToken = String(object.trustAnchorToken);
        }
        else {
            message.trustAnchorToken = "";
        }
        if (object.usn !== undefined && object.usn !== null) {
            message.usn = String(object.usn);
        }
        else {
            message.usn = "";
        }
        if (object.checksum !== undefined && object.checksum !== null) {
            message.checksum = String(object.checksum);
        }
        else {
            message.checksum = "";
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
        message.trustAnchorToken !== undefined &&
            (obj.trustAnchorToken = message.trustAnchorToken);
        message.usn !== undefined && (obj.usn = message.usn);
        message.checksum !== undefined && (obj.checksum = message.checksum);
        if (message.documents) {
            obj.documents = message.documents.map((e) => e ? NFTDocument.toJSON(e) : undefined);
        }
        else {
            obj.documents = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseNFTData };
        message.documents = [];
        if (object.trustAnchorToken !== undefined &&
            object.trustAnchorToken !== null) {
            message.trustAnchorToken = object.trustAnchorToken;
        }
        else {
            message.trustAnchorToken = "";
        }
        if (object.usn !== undefined && object.usn !== null) {
            message.usn = object.usn;
        }
        else {
            message.usn = "";
        }
        if (object.checksum !== undefined && object.checksum !== null) {
            message.checksum = object.checksum;
        }
        else {
            message.checksum = "";
        }
        if (object.documents !== undefined && object.documents !== null) {
            for (const e of object.documents) {
                message.documents.push(NFTDocument.fromPartial(e));
            }
        }
        return message;
    },
};
const baseNFTDocument = { name: "", uri: "", hash: "" };
export const NFTDocument = {
    encode(message, writer = Writer.create()) {
        if (message.name !== "") {
            writer.uint32(10).string(message.name);
        }
        if (message.uri !== "") {
            writer.uint32(18).string(message.uri);
        }
        if (message.hash !== "") {
            writer.uint32(26).string(message.hash);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseNFTDocument };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.name = reader.string();
                    break;
                case 2:
                    message.uri = reader.string();
                    break;
                case 3:
                    message.hash = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseNFTDocument };
        if (object.name !== undefined && object.name !== null) {
            message.name = String(object.name);
        }
        else {
            message.name = "";
        }
        if (object.uri !== undefined && object.uri !== null) {
            message.uri = String(object.uri);
        }
        else {
            message.uri = "";
        }
        if (object.hash !== undefined && object.hash !== null) {
            message.hash = String(object.hash);
        }
        else {
            message.hash = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.name !== undefined && (obj.name = message.name);
        message.uri !== undefined && (obj.uri = message.uri);
        message.hash !== undefined && (obj.hash = message.hash);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseNFTDocument };
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = "";
        }
        if (object.uri !== undefined && object.uri !== null) {
            message.uri = object.uri;
        }
        else {
            message.uri = "";
        }
        if (object.hash !== undefined && object.hash !== null) {
            message.hash = object.hash;
        }
        else {
            message.hash = "";
        }
        return message;
    },
};
