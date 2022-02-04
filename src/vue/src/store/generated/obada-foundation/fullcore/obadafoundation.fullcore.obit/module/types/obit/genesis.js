/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../obit/params";
import { Ta } from "../obit/ta";
export const protobufPackage = "obadafoundation.fullcore.obit";
const baseGenesisState = { taCount: 0 };
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        if (message.params !== undefined) {
            Params.encode(message.params, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.taList) {
            Ta.encode(v, writer.uint32(18).fork()).ldelim();
        }
        if (message.taCount !== 0) {
            writer.uint32(24).uint64(message.taCount);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.taList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.params = Params.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.taList.push(Ta.decode(reader, reader.uint32()));
                    break;
                case 3:
                    message.taCount = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.taList = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromJSON(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.taList !== undefined && object.taList !== null) {
            for (const e of object.taList) {
                message.taList.push(Ta.fromJSON(e));
            }
        }
        if (object.taCount !== undefined && object.taCount !== null) {
            message.taCount = Number(object.taCount);
        }
        else {
            message.taCount = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.params !== undefined &&
            (obj.params = message.params ? Params.toJSON(message.params) : undefined);
        if (message.taList) {
            obj.taList = message.taList.map((e) => (e ? Ta.toJSON(e) : undefined));
        }
        else {
            obj.taList = [];
        }
        message.taCount !== undefined && (obj.taCount = message.taCount);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.taList = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromPartial(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.taList !== undefined && object.taList !== null) {
            for (const e of object.taList) {
                message.taList.push(Ta.fromPartial(e));
            }
        }
        if (object.taCount !== undefined && object.taCount !== null) {
            message.taCount = object.taCount;
        }
        else {
            message.taCount = 0;
        }
        return message;
    },
};
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
