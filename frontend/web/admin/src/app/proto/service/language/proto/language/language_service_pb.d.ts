// package: micro.service.language.language.v1
// file: service/language/proto/language/language_service.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_wrappers_pb from "google-protobuf/google/protobuf/wrappers_pb";
import * as service_language_proto_entities_entities_pb from "../../../../backend/service/language/proto/entities/entities_pb";
import * as third_party_proto_validate_validate_pb from "../../../../backend/third_party/proto/validate/validate_pb";

export class ExistRequest extends jspb.Message {
  hasId(): boolean;
  clearId(): void;
  getId(): google_protobuf_wrappers_pb.StringValue | undefined;
  setId(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasName(): boolean;
  clearName(): void;
  getName(): google_protobuf_wrappers_pb.StringValue | undefined;
  setName(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasAbbreviation(): boolean;
  clearAbbreviation(): void;
  getAbbreviation(): google_protobuf_wrappers_pb.StringValue | undefined;
  setAbbreviation(value?: google_protobuf_wrappers_pb.StringValue): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExistRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ExistRequest): ExistRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ExistRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExistRequest;
  static deserializeBinaryFromReader(message: ExistRequest, reader: jspb.BinaryReader): ExistRequest;
}

export namespace ExistRequest {
  export type AsObject = {
    id?: google_protobuf_wrappers_pb.StringValue.AsObject,
    name?: google_protobuf_wrappers_pb.StringValue.AsObject,
    abbreviation?: google_protobuf_wrappers_pb.StringValue.AsObject,
  }
}

export class ExistResponse extends jspb.Message {
  getResult(): boolean;
  setResult(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExistResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ExistResponse): ExistResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ExistResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExistResponse;
  static deserializeBinaryFromReader(message: ExistResponse, reader: jspb.BinaryReader): ExistResponse;
}

export namespace ExistResponse {
  export type AsObject = {
    result: boolean,
  }
}

export class ListRequest extends jspb.Message {
  hasLimit(): boolean;
  clearLimit(): void;
  getLimit(): google_protobuf_wrappers_pb.UInt32Value | undefined;
  setLimit(value?: google_protobuf_wrappers_pb.UInt32Value): void;

  hasPage(): boolean;
  clearPage(): void;
  getPage(): google_protobuf_wrappers_pb.UInt32Value | undefined;
  setPage(value?: google_protobuf_wrappers_pb.UInt32Value): void;

  hasSort(): boolean;
  clearSort(): void;
  getSort(): google_protobuf_wrappers_pb.StringValue | undefined;
  setSort(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasName(): boolean;
  clearName(): void;
  getName(): google_protobuf_wrappers_pb.StringValue | undefined;
  setName(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasAbbreviation(): boolean;
  clearAbbreviation(): void;
  getAbbreviation(): google_protobuf_wrappers_pb.StringValue | undefined;
  setAbbreviation(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasFlagSrc(): boolean;
  clearFlagSrc(): void;
  getFlagSrc(): google_protobuf_wrappers_pb.StringValue | undefined;
  setFlagSrc(value?: google_protobuf_wrappers_pb.StringValue): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListRequest): ListRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRequest;
  static deserializeBinaryFromReader(message: ListRequest, reader: jspb.BinaryReader): ListRequest;
}

export namespace ListRequest {
  export type AsObject = {
    limit?: google_protobuf_wrappers_pb.UInt32Value.AsObject,
    page?: google_protobuf_wrappers_pb.UInt32Value.AsObject,
    sort?: google_protobuf_wrappers_pb.StringValue.AsObject,
    name?: google_protobuf_wrappers_pb.StringValue.AsObject,
    abbreviation?: google_protobuf_wrappers_pb.StringValue.AsObject,
    flagSrc?: google_protobuf_wrappers_pb.StringValue.AsObject,
  }
}

export class ListResponse extends jspb.Message {
  clearResultsList(): void;
  getResultsList(): Array<service_language_proto_entities_entities_pb.Language>;
  setResultsList(value: Array<service_language_proto_entities_entities_pb.Language>): void;
  addResults(value?: service_language_proto_entities_entities_pb.Language, index?: number): service_language_proto_entities_entities_pb.Language;

  getTotal(): number;
  setTotal(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListResponse): ListResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListResponse;
  static deserializeBinaryFromReader(message: ListResponse, reader: jspb.BinaryReader): ListResponse;
}

export namespace ListResponse {
  export type AsObject = {
    resultsList: Array<service_language_proto_entities_entities_pb.Language.AsObject>,
    total: number,
  }
}

export class GetRequest extends jspb.Message {
  hasId(): boolean;
  clearId(): void;
  getId(): google_protobuf_wrappers_pb.StringValue | undefined;
  setId(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasName(): boolean;
  clearName(): void;
  getName(): google_protobuf_wrappers_pb.StringValue | undefined;
  setName(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasAbbreviation(): boolean;
  clearAbbreviation(): void;
  getAbbreviation(): google_protobuf_wrappers_pb.StringValue | undefined;
  setAbbreviation(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasFlagSrc(): boolean;
  clearFlagSrc(): void;
  getFlagSrc(): google_protobuf_wrappers_pb.StringValue | undefined;
  setFlagSrc(value?: google_protobuf_wrappers_pb.StringValue): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetRequest): GetRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetRequest;
  static deserializeBinaryFromReader(message: GetRequest, reader: jspb.BinaryReader): GetRequest;
}

export namespace GetRequest {
  export type AsObject = {
    id?: google_protobuf_wrappers_pb.StringValue.AsObject,
    name?: google_protobuf_wrappers_pb.StringValue.AsObject,
    abbreviation?: google_protobuf_wrappers_pb.StringValue.AsObject,
    flagSrc?: google_protobuf_wrappers_pb.StringValue.AsObject,
  }
}

export class GetResponse extends jspb.Message {
  hasResult(): boolean;
  clearResult(): void;
  getResult(): service_language_proto_entities_entities_pb.Language | undefined;
  setResult(value?: service_language_proto_entities_entities_pb.Language): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetResponse): GetResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetResponse;
  static deserializeBinaryFromReader(message: GetResponse, reader: jspb.BinaryReader): GetResponse;
}

export namespace GetResponse {
  export type AsObject = {
    result?: service_language_proto_entities_entities_pb.Language.AsObject,
  }
}

export class CreateRequest extends jspb.Message {
  hasName(): boolean;
  clearName(): void;
  getName(): google_protobuf_wrappers_pb.StringValue | undefined;
  setName(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasAbbreviation(): boolean;
  clearAbbreviation(): void;
  getAbbreviation(): google_protobuf_wrappers_pb.StringValue | undefined;
  setAbbreviation(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasFlagSrc(): boolean;
  clearFlagSrc(): void;
  getFlagSrc(): google_protobuf_wrappers_pb.StringValue | undefined;
  setFlagSrc(value?: google_protobuf_wrappers_pb.StringValue): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateRequest): CreateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateRequest;
  static deserializeBinaryFromReader(message: CreateRequest, reader: jspb.BinaryReader): CreateRequest;
}

export namespace CreateRequest {
  export type AsObject = {
    name?: google_protobuf_wrappers_pb.StringValue.AsObject,
    abbreviation?: google_protobuf_wrappers_pb.StringValue.AsObject,
    flagSrc?: google_protobuf_wrappers_pb.StringValue.AsObject,
  }
}

export class CreateResponse extends jspb.Message {
  hasResult(): boolean;
  clearResult(): void;
  getResult(): service_language_proto_entities_entities_pb.Language | undefined;
  setResult(value?: service_language_proto_entities_entities_pb.Language): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateResponse): CreateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateResponse;
  static deserializeBinaryFromReader(message: CreateResponse, reader: jspb.BinaryReader): CreateResponse;
}

export namespace CreateResponse {
  export type AsObject = {
    result?: service_language_proto_entities_entities_pb.Language.AsObject,
  }
}

export class UpdateRequest extends jspb.Message {
  hasId(): boolean;
  clearId(): void;
  getId(): google_protobuf_wrappers_pb.StringValue | undefined;
  setId(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasName(): boolean;
  clearName(): void;
  getName(): google_protobuf_wrappers_pb.StringValue | undefined;
  setName(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasAbbreviation(): boolean;
  clearAbbreviation(): void;
  getAbbreviation(): google_protobuf_wrappers_pb.StringValue | undefined;
  setAbbreviation(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasFlagSrc(): boolean;
  clearFlagSrc(): void;
  getFlagSrc(): google_protobuf_wrappers_pb.StringValue | undefined;
  setFlagSrc(value?: google_protobuf_wrappers_pb.StringValue): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateRequest): UpdateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateRequest;
  static deserializeBinaryFromReader(message: UpdateRequest, reader: jspb.BinaryReader): UpdateRequest;
}

export namespace UpdateRequest {
  export type AsObject = {
    id?: google_protobuf_wrappers_pb.StringValue.AsObject,
    name?: google_protobuf_wrappers_pb.StringValue.AsObject,
    abbreviation?: google_protobuf_wrappers_pb.StringValue.AsObject,
    flagSrc?: google_protobuf_wrappers_pb.StringValue.AsObject,
  }
}

export class UpdateResponse extends jspb.Message {
  hasResult(): boolean;
  clearResult(): void;
  getResult(): service_language_proto_entities_entities_pb.Language | undefined;
  setResult(value?: service_language_proto_entities_entities_pb.Language): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateResponse): UpdateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateResponse;
  static deserializeBinaryFromReader(message: UpdateResponse, reader: jspb.BinaryReader): UpdateResponse;
}

export namespace UpdateResponse {
  export type AsObject = {
    result?: service_language_proto_entities_entities_pb.Language.AsObject,
  }
}

export class DeleteRequest extends jspb.Message {
  hasId(): boolean;
  clearId(): void;
  getId(): google_protobuf_wrappers_pb.StringValue | undefined;
  setId(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasName(): boolean;
  clearName(): void;
  getName(): google_protobuf_wrappers_pb.StringValue | undefined;
  setName(value?: google_protobuf_wrappers_pb.StringValue): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteRequest): DeleteRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteRequest;
  static deserializeBinaryFromReader(message: DeleteRequest, reader: jspb.BinaryReader): DeleteRequest;
}

export namespace DeleteRequest {
  export type AsObject = {
    id?: google_protobuf_wrappers_pb.StringValue.AsObject,
    name?: google_protobuf_wrappers_pb.StringValue.AsObject,
  }
}

export class DeleteResponse extends jspb.Message {
  hasResult(): boolean;
  clearResult(): void;
  getResult(): service_language_proto_entities_entities_pb.Language | undefined;
  setResult(value?: service_language_proto_entities_entities_pb.Language): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteResponse): DeleteResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteResponse;
  static deserializeBinaryFromReader(message: DeleteResponse, reader: jspb.BinaryReader): DeleteResponse;
}

export namespace DeleteResponse {
  export type AsObject = {
    result?: service_language_proto_entities_entities_pb.Language.AsObject,
  }
}

