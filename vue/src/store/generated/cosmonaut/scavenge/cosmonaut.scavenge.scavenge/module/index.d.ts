import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCommitSolution } from "./types/scavenge/tx";
import { MsgSubmitScavenge } from "./types/scavenge/tx";
import { MsgRevealSolution } from "./types/scavenge/tx";
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
    msgCommitSolution: (data: MsgCommitSolution) => EncodeObject;
    msgSubmitScavenge: (data: MsgSubmitScavenge) => EncodeObject;
    msgRevealSolution: (data: MsgRevealSolution) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };