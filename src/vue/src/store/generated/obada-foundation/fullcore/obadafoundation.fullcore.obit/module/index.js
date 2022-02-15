// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgSend } from "./types/obit/tx";
import { MsgUpdateTa } from "./types/obit/tx";
import { MsgCreateTa } from "./types/obit/tx";
import { MsgMintObit } from "./types/obit/tx";
import { MsgDeleteTa } from "./types/obit/tx";
const types = [
    ["/obadafoundation.fullcore.obit.MsgSend", MsgSend],
    ["/obadafoundation.fullcore.obit.MsgUpdateTa", MsgUpdateTa],
    ["/obadafoundation.fullcore.obit.MsgCreateTa", MsgCreateTa],
    ["/obadafoundation.fullcore.obit.MsgMintObit", MsgMintObit],
    ["/obadafoundation.fullcore.obit.MsgDeleteTa", MsgDeleteTa],
];
export const MissingWalletError = new Error("wallet is required");
export const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    let client;
    if (addr) {
        client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    }
    else {
        client = await SigningStargateClient.offline(wallet, { registry });
    }
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgSend: (data) => ({ typeUrl: "/obadafoundation.fullcore.obit.MsgSend", value: MsgSend.fromPartial(data) }),
        msgUpdateTa: (data) => ({ typeUrl: "/obadafoundation.fullcore.obit.MsgUpdateTa", value: MsgUpdateTa.fromPartial(data) }),
        msgCreateTa: (data) => ({ typeUrl: "/obadafoundation.fullcore.obit.MsgCreateTa", value: MsgCreateTa.fromPartial(data) }),
        msgMintObit: (data) => ({ typeUrl: "/obadafoundation.fullcore.obit.MsgMintObit", value: MsgMintObit.fromPartial(data) }),
        msgDeleteTa: (data) => ({ typeUrl: "/obadafoundation.fullcore.obit.MsgDeleteTa", value: MsgDeleteTa.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
