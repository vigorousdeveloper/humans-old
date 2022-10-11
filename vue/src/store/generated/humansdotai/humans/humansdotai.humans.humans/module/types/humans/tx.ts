/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "humansdotai.humans.humans";

export interface MsgRequestTransaction {
  creator: string;
  originChain: string;
  originAddress: string;
  targetChain: string;
  targetAddress: string;
  amount: string;
  fee: string;
}

export interface MsgRequestTransactionResponse {}

export interface MsgObservationVote {
  creator: string;
  txHash: string;
  chainName: string;
  from: string;
  to: string;
  amount: string;
}

export interface MsgObservationVoteResponse {}

export interface MsgUpdateBalance {
  creator: string;
  chainName: string;
  balance: string;
  decimal: string;
}

export interface MsgUpdateBalanceResponse {}

export interface MsgApproveTransaction {
  creator: string;
  txHash: string;
  success: string;
  signedKey: string;
}

export interface MsgApproveTransactionResponse {}

export interface MsgTransferPoolcoin {
  creator: string;
  addr: string;
  pool: string;
  amt: string;
}

export interface MsgTransferPoolcoinResponse {}

export interface MsgAddWhitelisted {
  creator: string;
  address: string;
}

export interface MsgAddWhitelistedResponse {}

export interface MsgSetAdmin {
  creator: string;
}

export interface MsgSetAdminResponse {}

const baseMsgRequestTransaction: object = {
  creator: "",
  originChain: "",
  originAddress: "",
  targetChain: "",
  targetAddress: "",
  amount: "",
  fee: "",
};

export const MsgRequestTransaction = {
  encode(
    message: MsgRequestTransaction,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.originChain !== "") {
      writer.uint32(18).string(message.originChain);
    }
    if (message.originAddress !== "") {
      writer.uint32(26).string(message.originAddress);
    }
    if (message.targetChain !== "") {
      writer.uint32(34).string(message.targetChain);
    }
    if (message.targetAddress !== "") {
      writer.uint32(42).string(message.targetAddress);
    }
    if (message.amount !== "") {
      writer.uint32(50).string(message.amount);
    }
    if (message.fee !== "") {
      writer.uint32(58).string(message.fee);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRequestTransaction {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRequestTransaction } as MsgRequestTransaction;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.originChain = reader.string();
          break;
        case 3:
          message.originAddress = reader.string();
          break;
        case 4:
          message.targetChain = reader.string();
          break;
        case 5:
          message.targetAddress = reader.string();
          break;
        case 6:
          message.amount = reader.string();
          break;
        case 7:
          message.fee = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRequestTransaction {
    const message = { ...baseMsgRequestTransaction } as MsgRequestTransaction;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.originChain !== undefined && object.originChain !== null) {
      message.originChain = String(object.originChain);
    } else {
      message.originChain = "";
    }
    if (object.originAddress !== undefined && object.originAddress !== null) {
      message.originAddress = String(object.originAddress);
    } else {
      message.originAddress = "";
    }
    if (object.targetChain !== undefined && object.targetChain !== null) {
      message.targetChain = String(object.targetChain);
    } else {
      message.targetChain = "";
    }
    if (object.targetAddress !== undefined && object.targetAddress !== null) {
      message.targetAddress = String(object.targetAddress);
    } else {
      message.targetAddress = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = String(object.fee);
    } else {
      message.fee = "";
    }
    return message;
  },

  toJSON(message: MsgRequestTransaction): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.originChain !== undefined &&
      (obj.originChain = message.originChain);
    message.originAddress !== undefined &&
      (obj.originAddress = message.originAddress);
    message.targetChain !== undefined &&
      (obj.targetChain = message.targetChain);
    message.targetAddress !== undefined &&
      (obj.targetAddress = message.targetAddress);
    message.amount !== undefined && (obj.amount = message.amount);
    message.fee !== undefined && (obj.fee = message.fee);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgRequestTransaction>
  ): MsgRequestTransaction {
    const message = { ...baseMsgRequestTransaction } as MsgRequestTransaction;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.originChain !== undefined && object.originChain !== null) {
      message.originChain = object.originChain;
    } else {
      message.originChain = "";
    }
    if (object.originAddress !== undefined && object.originAddress !== null) {
      message.originAddress = object.originAddress;
    } else {
      message.originAddress = "";
    }
    if (object.targetChain !== undefined && object.targetChain !== null) {
      message.targetChain = object.targetChain;
    } else {
      message.targetChain = "";
    }
    if (object.targetAddress !== undefined && object.targetAddress !== null) {
      message.targetAddress = object.targetAddress;
    } else {
      message.targetAddress = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = object.fee;
    } else {
      message.fee = "";
    }
    return message;
  },
};

const baseMsgRequestTransactionResponse: object = {};

export const MsgRequestTransactionResponse = {
  encode(
    _: MsgRequestTransactionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRequestTransactionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRequestTransactionResponse,
    } as MsgRequestTransactionResponse;
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

  fromJSON(_: any): MsgRequestTransactionResponse {
    const message = {
      ...baseMsgRequestTransactionResponse,
    } as MsgRequestTransactionResponse;
    return message;
  },

  toJSON(_: MsgRequestTransactionResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgRequestTransactionResponse>
  ): MsgRequestTransactionResponse {
    const message = {
      ...baseMsgRequestTransactionResponse,
    } as MsgRequestTransactionResponse;
    return message;
  },
};

const baseMsgObservationVote: object = {
  creator: "",
  txHash: "",
  chainName: "",
  from: "",
  to: "",
  amount: "",
};

export const MsgObservationVote = {
  encode(
    message: MsgObservationVote,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.txHash !== "") {
      writer.uint32(18).string(message.txHash);
    }
    if (message.chainName !== "") {
      writer.uint32(26).string(message.chainName);
    }
    if (message.from !== "") {
      writer.uint32(34).string(message.from);
    }
    if (message.to !== "") {
      writer.uint32(42).string(message.to);
    }
    if (message.amount !== "") {
      writer.uint32(50).string(message.amount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgObservationVote {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgObservationVote } as MsgObservationVote;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.txHash = reader.string();
          break;
        case 3:
          message.chainName = reader.string();
          break;
        case 4:
          message.from = reader.string();
          break;
        case 5:
          message.to = reader.string();
          break;
        case 6:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgObservationVote {
    const message = { ...baseMsgObservationVote } as MsgObservationVote;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.txHash !== undefined && object.txHash !== null) {
      message.txHash = String(object.txHash);
    } else {
      message.txHash = "";
    }
    if (object.chainName !== undefined && object.chainName !== null) {
      message.chainName = String(object.chainName);
    } else {
      message.chainName = "";
    }
    if (object.from !== undefined && object.from !== null) {
      message.from = String(object.from);
    } else {
      message.from = "";
    }
    if (object.to !== undefined && object.to !== null) {
      message.to = String(object.to);
    } else {
      message.to = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    return message;
  },

  toJSON(message: MsgObservationVote): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.txHash !== undefined && (obj.txHash = message.txHash);
    message.chainName !== undefined && (obj.chainName = message.chainName);
    message.from !== undefined && (obj.from = message.from);
    message.to !== undefined && (obj.to = message.to);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgObservationVote>): MsgObservationVote {
    const message = { ...baseMsgObservationVote } as MsgObservationVote;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.txHash !== undefined && object.txHash !== null) {
      message.txHash = object.txHash;
    } else {
      message.txHash = "";
    }
    if (object.chainName !== undefined && object.chainName !== null) {
      message.chainName = object.chainName;
    } else {
      message.chainName = "";
    }
    if (object.from !== undefined && object.from !== null) {
      message.from = object.from;
    } else {
      message.from = "";
    }
    if (object.to !== undefined && object.to !== null) {
      message.to = object.to;
    } else {
      message.to = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    return message;
  },
};

const baseMsgObservationVoteResponse: object = {};

export const MsgObservationVoteResponse = {
  encode(
    _: MsgObservationVoteResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgObservationVoteResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgObservationVoteResponse,
    } as MsgObservationVoteResponse;
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

  fromJSON(_: any): MsgObservationVoteResponse {
    const message = {
      ...baseMsgObservationVoteResponse,
    } as MsgObservationVoteResponse;
    return message;
  },

  toJSON(_: MsgObservationVoteResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgObservationVoteResponse>
  ): MsgObservationVoteResponse {
    const message = {
      ...baseMsgObservationVoteResponse,
    } as MsgObservationVoteResponse;
    return message;
  },
};

const baseMsgUpdateBalance: object = {
  creator: "",
  chainName: "",
  balance: "",
  decimal: "",
};

export const MsgUpdateBalance = {
  encode(message: MsgUpdateBalance, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.chainName !== "") {
      writer.uint32(18).string(message.chainName);
    }
    if (message.balance !== "") {
      writer.uint32(26).string(message.balance);
    }
    if (message.decimal !== "") {
      writer.uint32(34).string(message.decimal);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateBalance {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateBalance } as MsgUpdateBalance;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.chainName = reader.string();
          break;
        case 3:
          message.balance = reader.string();
          break;
        case 4:
          message.decimal = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateBalance {
    const message = { ...baseMsgUpdateBalance } as MsgUpdateBalance;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.chainName !== undefined && object.chainName !== null) {
      message.chainName = String(object.chainName);
    } else {
      message.chainName = "";
    }
    if (object.balance !== undefined && object.balance !== null) {
      message.balance = String(object.balance);
    } else {
      message.balance = "";
    }
    if (object.decimal !== undefined && object.decimal !== null) {
      message.decimal = String(object.decimal);
    } else {
      message.decimal = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateBalance): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.chainName !== undefined && (obj.chainName = message.chainName);
    message.balance !== undefined && (obj.balance = message.balance);
    message.decimal !== undefined && (obj.decimal = message.decimal);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateBalance>): MsgUpdateBalance {
    const message = { ...baseMsgUpdateBalance } as MsgUpdateBalance;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.chainName !== undefined && object.chainName !== null) {
      message.chainName = object.chainName;
    } else {
      message.chainName = "";
    }
    if (object.balance !== undefined && object.balance !== null) {
      message.balance = object.balance;
    } else {
      message.balance = "";
    }
    if (object.decimal !== undefined && object.decimal !== null) {
      message.decimal = object.decimal;
    } else {
      message.decimal = "";
    }
    return message;
  },
};

const baseMsgUpdateBalanceResponse: object = {};

export const MsgUpdateBalanceResponse = {
  encode(
    _: MsgUpdateBalanceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateBalanceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateBalanceResponse,
    } as MsgUpdateBalanceResponse;
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

  fromJSON(_: any): MsgUpdateBalanceResponse {
    const message = {
      ...baseMsgUpdateBalanceResponse,
    } as MsgUpdateBalanceResponse;
    return message;
  },

  toJSON(_: MsgUpdateBalanceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateBalanceResponse>
  ): MsgUpdateBalanceResponse {
    const message = {
      ...baseMsgUpdateBalanceResponse,
    } as MsgUpdateBalanceResponse;
    return message;
  },
};

const baseMsgApproveTransaction: object = {
  creator: "",
  txHash: "",
  success: "",
  signedKey: "",
};

export const MsgApproveTransaction = {
  encode(
    message: MsgApproveTransaction,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.txHash !== "") {
      writer.uint32(18).string(message.txHash);
    }
    if (message.success !== "") {
      writer.uint32(26).string(message.success);
    }
    if (message.signedKey !== "") {
      writer.uint32(34).string(message.signedKey);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgApproveTransaction {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgApproveTransaction } as MsgApproveTransaction;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.txHash = reader.string();
          break;
        case 3:
          message.success = reader.string();
          break;
        case 4:
          message.signedKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgApproveTransaction {
    const message = { ...baseMsgApproveTransaction } as MsgApproveTransaction;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.txHash !== undefined && object.txHash !== null) {
      message.txHash = String(object.txHash);
    } else {
      message.txHash = "";
    }
    if (object.success !== undefined && object.success !== null) {
      message.success = String(object.success);
    } else {
      message.success = "";
    }
    if (object.signedKey !== undefined && object.signedKey !== null) {
      message.signedKey = String(object.signedKey);
    } else {
      message.signedKey = "";
    }
    return message;
  },

  toJSON(message: MsgApproveTransaction): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.txHash !== undefined && (obj.txHash = message.txHash);
    message.success !== undefined && (obj.success = message.success);
    message.signedKey !== undefined && (obj.signedKey = message.signedKey);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgApproveTransaction>
  ): MsgApproveTransaction {
    const message = { ...baseMsgApproveTransaction } as MsgApproveTransaction;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.txHash !== undefined && object.txHash !== null) {
      message.txHash = object.txHash;
    } else {
      message.txHash = "";
    }
    if (object.success !== undefined && object.success !== null) {
      message.success = object.success;
    } else {
      message.success = "";
    }
    if (object.signedKey !== undefined && object.signedKey !== null) {
      message.signedKey = object.signedKey;
    } else {
      message.signedKey = "";
    }
    return message;
  },
};

const baseMsgApproveTransactionResponse: object = {};

export const MsgApproveTransactionResponse = {
  encode(
    _: MsgApproveTransactionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgApproveTransactionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgApproveTransactionResponse,
    } as MsgApproveTransactionResponse;
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

  fromJSON(_: any): MsgApproveTransactionResponse {
    const message = {
      ...baseMsgApproveTransactionResponse,
    } as MsgApproveTransactionResponse;
    return message;
  },

  toJSON(_: MsgApproveTransactionResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgApproveTransactionResponse>
  ): MsgApproveTransactionResponse {
    const message = {
      ...baseMsgApproveTransactionResponse,
    } as MsgApproveTransactionResponse;
    return message;
  },
};

const baseMsgTransferPoolcoin: object = {
  creator: "",
  addr: "",
  pool: "",
  amt: "",
};

export const MsgTransferPoolcoin = {
  encode(
    message: MsgTransferPoolcoin,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.addr !== "") {
      writer.uint32(18).string(message.addr);
    }
    if (message.pool !== "") {
      writer.uint32(26).string(message.pool);
    }
    if (message.amt !== "") {
      writer.uint32(34).string(message.amt);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgTransferPoolcoin {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgTransferPoolcoin } as MsgTransferPoolcoin;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.addr = reader.string();
          break;
        case 3:
          message.pool = reader.string();
          break;
        case 4:
          message.amt = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgTransferPoolcoin {
    const message = { ...baseMsgTransferPoolcoin } as MsgTransferPoolcoin;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.addr !== undefined && object.addr !== null) {
      message.addr = String(object.addr);
    } else {
      message.addr = "";
    }
    if (object.pool !== undefined && object.pool !== null) {
      message.pool = String(object.pool);
    } else {
      message.pool = "";
    }
    if (object.amt !== undefined && object.amt !== null) {
      message.amt = String(object.amt);
    } else {
      message.amt = "";
    }
    return message;
  },

  toJSON(message: MsgTransferPoolcoin): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.addr !== undefined && (obj.addr = message.addr);
    message.pool !== undefined && (obj.pool = message.pool);
    message.amt !== undefined && (obj.amt = message.amt);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgTransferPoolcoin>): MsgTransferPoolcoin {
    const message = { ...baseMsgTransferPoolcoin } as MsgTransferPoolcoin;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.addr !== undefined && object.addr !== null) {
      message.addr = object.addr;
    } else {
      message.addr = "";
    }
    if (object.pool !== undefined && object.pool !== null) {
      message.pool = object.pool;
    } else {
      message.pool = "";
    }
    if (object.amt !== undefined && object.amt !== null) {
      message.amt = object.amt;
    } else {
      message.amt = "";
    }
    return message;
  },
};

const baseMsgTransferPoolcoinResponse: object = {};

export const MsgTransferPoolcoinResponse = {
  encode(
    _: MsgTransferPoolcoinResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgTransferPoolcoinResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgTransferPoolcoinResponse,
    } as MsgTransferPoolcoinResponse;
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

  fromJSON(_: any): MsgTransferPoolcoinResponse {
    const message = {
      ...baseMsgTransferPoolcoinResponse,
    } as MsgTransferPoolcoinResponse;
    return message;
  },

  toJSON(_: MsgTransferPoolcoinResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgTransferPoolcoinResponse>
  ): MsgTransferPoolcoinResponse {
    const message = {
      ...baseMsgTransferPoolcoinResponse,
    } as MsgTransferPoolcoinResponse;
    return message;
  },
};

const baseMsgAddWhitelisted: object = { creator: "", address: "" };

export const MsgAddWhitelisted = {
  encode(message: MsgAddWhitelisted, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddWhitelisted {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddWhitelisted } as MsgAddWhitelisted;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddWhitelisted {
    const message = { ...baseMsgAddWhitelisted } as MsgAddWhitelisted;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: MsgAddWhitelisted): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddWhitelisted>): MsgAddWhitelisted {
    const message = { ...baseMsgAddWhitelisted } as MsgAddWhitelisted;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseMsgAddWhitelistedResponse: object = {};

export const MsgAddWhitelistedResponse = {
  encode(
    _: MsgAddWhitelistedResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgAddWhitelistedResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgAddWhitelistedResponse,
    } as MsgAddWhitelistedResponse;
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

  fromJSON(_: any): MsgAddWhitelistedResponse {
    const message = {
      ...baseMsgAddWhitelistedResponse,
    } as MsgAddWhitelistedResponse;
    return message;
  },

  toJSON(_: MsgAddWhitelistedResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgAddWhitelistedResponse>
  ): MsgAddWhitelistedResponse {
    const message = {
      ...baseMsgAddWhitelistedResponse,
    } as MsgAddWhitelistedResponse;
    return message;
  },
};

const baseMsgSetAdmin: object = { creator: "" };

export const MsgSetAdmin = {
  encode(message: MsgSetAdmin, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetAdmin {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSetAdmin } as MsgSetAdmin;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSetAdmin {
    const message = { ...baseMsgSetAdmin } as MsgSetAdmin;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: MsgSetAdmin): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSetAdmin>): MsgSetAdmin {
    const message = { ...baseMsgSetAdmin } as MsgSetAdmin;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseMsgSetAdminResponse: object = {};

export const MsgSetAdminResponse = {
  encode(_: MsgSetAdminResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetAdminResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSetAdminResponse } as MsgSetAdminResponse;
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

  fromJSON(_: any): MsgSetAdminResponse {
    const message = { ...baseMsgSetAdminResponse } as MsgSetAdminResponse;
    return message;
  },

  toJSON(_: MsgSetAdminResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgSetAdminResponse>): MsgSetAdminResponse {
    const message = { ...baseMsgSetAdminResponse } as MsgSetAdminResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  RequestTransaction(
    request: MsgRequestTransaction
  ): Promise<MsgRequestTransactionResponse>;
  ObservationVote(
    request: MsgObservationVote
  ): Promise<MsgObservationVoteResponse>;
  UpdateBalance(request: MsgUpdateBalance): Promise<MsgUpdateBalanceResponse>;
  ApproveTransaction(
    request: MsgApproveTransaction
  ): Promise<MsgApproveTransactionResponse>;
  TransferPoolcoin(
    request: MsgTransferPoolcoin
  ): Promise<MsgTransferPoolcoinResponse>;
  AddWhitelisted(
    request: MsgAddWhitelisted
  ): Promise<MsgAddWhitelistedResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SetAdmin(request: MsgSetAdmin): Promise<MsgSetAdminResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  RequestTransaction(
    request: MsgRequestTransaction
  ): Promise<MsgRequestTransactionResponse> {
    const data = MsgRequestTransaction.encode(request).finish();
    const promise = this.rpc.request(
      "humansdotai.humans.humans.Msg",
      "RequestTransaction",
      data
    );
    return promise.then((data) =>
      MsgRequestTransactionResponse.decode(new Reader(data))
    );
  }

  ObservationVote(
    request: MsgObservationVote
  ): Promise<MsgObservationVoteResponse> {
    const data = MsgObservationVote.encode(request).finish();
    const promise = this.rpc.request(
      "humansdotai.humans.humans.Msg",
      "ObservationVote",
      data
    );
    return promise.then((data) =>
      MsgObservationVoteResponse.decode(new Reader(data))
    );
  }

  UpdateBalance(request: MsgUpdateBalance): Promise<MsgUpdateBalanceResponse> {
    const data = MsgUpdateBalance.encode(request).finish();
    const promise = this.rpc.request(
      "humansdotai.humans.humans.Msg",
      "UpdateBalance",
      data
    );
    return promise.then((data) =>
      MsgUpdateBalanceResponse.decode(new Reader(data))
    );
  }

  ApproveTransaction(
    request: MsgApproveTransaction
  ): Promise<MsgApproveTransactionResponse> {
    const data = MsgApproveTransaction.encode(request).finish();
    const promise = this.rpc.request(
      "humansdotai.humans.humans.Msg",
      "ApproveTransaction",
      data
    );
    return promise.then((data) =>
      MsgApproveTransactionResponse.decode(new Reader(data))
    );
  }

  TransferPoolcoin(
    request: MsgTransferPoolcoin
  ): Promise<MsgTransferPoolcoinResponse> {
    const data = MsgTransferPoolcoin.encode(request).finish();
    const promise = this.rpc.request(
      "humansdotai.humans.humans.Msg",
      "TransferPoolcoin",
      data
    );
    return promise.then((data) =>
      MsgTransferPoolcoinResponse.decode(new Reader(data))
    );
  }

  AddWhitelisted(
    request: MsgAddWhitelisted
  ): Promise<MsgAddWhitelistedResponse> {
    const data = MsgAddWhitelisted.encode(request).finish();
    const promise = this.rpc.request(
      "humansdotai.humans.humans.Msg",
      "AddWhitelisted",
      data
    );
    return promise.then((data) =>
      MsgAddWhitelistedResponse.decode(new Reader(data))
    );
  }

  SetAdmin(request: MsgSetAdmin): Promise<MsgSetAdminResponse> {
    const data = MsgSetAdmin.encode(request).finish();
    const promise = this.rpc.request(
      "humansdotai.humans.humans.Msg",
      "SetAdmin",
      data
    );
    return promise.then((data) => MsgSetAdminResponse.decode(new Reader(data)));
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
