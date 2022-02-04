import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdateTa } from "./types/obit/tx";
import { MsgCreateTa } from "./types/obit/tx";
import { MsgMintObit } from "./types/obit/tx";
import { MsgDeleteTa } from "./types/obit/tx";
export declare const MissingWalletError: Error;
export declare const registry: Registry;
interface TxClientOptions {
    addr: string;
}
interface SignAndBroadcastOptions {
    fee: StdFee;
    memo?: string;
}
declare const txClient: (wallet: OfflineSigner, { addr: addr }?: TxClientOptions) => Promise<{
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }?: SignAndBroadcastOptions) => any;
    msgUpdateTa: (data: MsgUpdateTa) => EncodeObject;
    msgCreateTa: (data: MsgCreateTa) => EncodeObject;
    msgMintObit: (data: MsgMintObit) => EncodeObject;
    msgDeleteTa: (data: MsgDeleteTa) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
