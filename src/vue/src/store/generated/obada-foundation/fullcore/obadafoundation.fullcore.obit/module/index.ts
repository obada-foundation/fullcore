// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgSend } from "./types/obit/tx";
import { MsgDeleteTa } from "./types/obit/tx";
import { MsgCreateTa } from "./types/obit/tx";
import { MsgMintObit } from "./types/obit/tx";
import { MsgUpdateTa } from "./types/obit/tx";


const types = [
  ["/obadafoundation.fullcore.obit.MsgSend", MsgSend],
  ["/obadafoundation.fullcore.obit.MsgDeleteTa", MsgDeleteTa],
  ["/obadafoundation.fullcore.obit.MsgCreateTa", MsgCreateTa],
  ["/obadafoundation.fullcore.obit.MsgMintObit", MsgMintObit],
  ["/obadafoundation.fullcore.obit.MsgUpdateTa", MsgUpdateTa],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgSend: (data: MsgSend): EncodeObject => ({ typeUrl: "/obadafoundation.fullcore.obit.MsgSend", value: MsgSend.fromPartial( data ) }),
    msgDeleteTa: (data: MsgDeleteTa): EncodeObject => ({ typeUrl: "/obadafoundation.fullcore.obit.MsgDeleteTa", value: MsgDeleteTa.fromPartial( data ) }),
    msgCreateTa: (data: MsgCreateTa): EncodeObject => ({ typeUrl: "/obadafoundation.fullcore.obit.MsgCreateTa", value: MsgCreateTa.fromPartial( data ) }),
    msgMintObit: (data: MsgMintObit): EncodeObject => ({ typeUrl: "/obadafoundation.fullcore.obit.MsgMintObit", value: MsgMintObit.fromPartial( data ) }),
    msgUpdateTa: (data: MsgUpdateTa): EncodeObject => ({ typeUrl: "/obadafoundation.fullcore.obit.MsgUpdateTa", value: MsgUpdateTa.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
