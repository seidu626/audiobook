// package: micro.service.language.skill.v1
// file: service/language/proto/skill/skill_service.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_wrappers_pb from "google-protobuf/google/protobuf/wrappers_pb";
import * as service_language_proto_entities_entities_pb from "../../../../service/language/proto/entities/entities_pb";
import * as third_party_proto_validate_validate_pb from "../../../../third_party/proto/validate/validate_pb";

export class ExistRequest extends jspb.Message {
  hasId(): boolean;
  clearId(): void;
  getId(): google_protobuf_wrappers_pb.StringValue | undefined;
  setId(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasTitle(): boolean;
  clearTitle(): void;
  getTitle(): google_protobuf_wrappers_pb.StringValue | undefined;
  setTitle(value?: google_protobuf_wrappers_pb.StringValue): void;

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
    title?: google_protobuf_wrappers_pb.StringValue.AsObject,
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

  hasTitle(): boolean;
  clearTitle(): void;
  getTitle(): google_protobuf_wrappers_pb.StringValue | undefined;
  setTitle(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasUrlTitle(): boolean;
  clearUrlTitle(): void;
  getUrlTitle(): google_protobuf_wrappers_pb.StringValue | undefined;
  setUrlTitle(value?: google_protobuf_wrappers_pb.StringValue): void;

  getLessonNumber(): number;
  setLessonNumber(value: number): void;

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
    title?: google_protobuf_wrappers_pb.StringValue.AsObject,
    urlTitle?: google_protobuf_wrappers_pb.StringValue.AsObject,
    lessonNumber: number,
  }
}

export class ListResponse extends jspb.Message {
  clearResultsList(): void;
  getResultsList(): Array<service_language_proto_entities_entities_pb.Skill>;
  setResultsList(value: Array<service_language_proto_entities_entities_pb.Skill>): void;
  addResults(value?: service_language_proto_entities_entities_pb.Skill, index?: number): service_language_proto_entities_entities_pb.Skill;

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
    resultsList: Array<service_language_proto_entities_entities_pb.Skill.AsObject>,
    total: number,
  }
}

export class GetRequest extends jspb.Message {
  hasId(): boolean;
  clearId(): void;
  getId(): google_protobuf_wrappers_pb.StringValue | undefined;
  setId(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasTitle(): boolean;
  clearTitle(): void;
  getTitle(): google_protobuf_wrappers_pb.StringValue | undefined;
  setTitle(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasUrlTitle(): boolean;
  clearUrlTitle(): void;
  getUrlTitle(): google_protobuf_wrappers_pb.StringValue | undefined;
  setUrlTitle(value?: google_protobuf_wrappers_pb.StringValue): void;

  getLessonNumber(): number;
  setLessonNumber(value: number): void;

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
    title?: google_protobuf_wrappers_pb.StringValue.AsObject,
    urlTitle?: google_protobuf_wrappers_pb.StringValue.AsObject,
    lessonNumber: number,
  }
}

export class GetResponse extends jspb.Message {
  hasResult(): boolean;
  clearResult(): void;
  getResult(): service_language_proto_entities_entities_pb.Skill | undefined;
  setResult(value?: service_language_proto_entities_entities_pb.Skill): void;

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
    result?: service_language_proto_entities_entities_pb.Skill.AsObject,
  }
}

export class CreateRequest extends jspb.Message {
  hasId(): boolean;
  clearId(): void;
  getId(): google_protobuf_wrappers_pb.StringValue | undefined;
  setId(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasTitle(): boolean;
  clearTitle(): void;
  getTitle(): google_protobuf_wrappers_pb.StringValue | undefined;
  setTitle(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasUrlTitle(): boolean;
  clearUrlTitle(): void;
  getUrlTitle(): google_protobuf_wrappers_pb.StringValue | undefined;
  setUrlTitle(value?: google_protobuf_wrappers_pb.StringValue): void;

  getLessonNumber(): number;
  setLessonNumber(value: number): void;

  clearDependenciesList(): void;
  getDependenciesList(): Array<google_protobuf_wrappers_pb.StringValue>;
  setDependenciesList(value: Array<google_protobuf_wrappers_pb.StringValue>): void;
  addDependencies(value?: google_protobuf_wrappers_pb.StringValue, index?: number): google_protobuf_wrappers_pb.StringValue;

  getDisabled(): boolean;
  setDisabled(value: boolean): void;

  hasDescription(): boolean;
  clearDescription(): void;
  getDescription(): google_protobuf_wrappers_pb.StringValue | undefined;
  setDescription(value?: google_protobuf_wrappers_pb.StringValue): void;

  getLocked(): boolean;
  setLocked(value: boolean): void;

  hasType(): boolean;
  clearType(): void;
  getType(): google_protobuf_wrappers_pb.StringValue | undefined;
  setType(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasCategory(): boolean;
  clearCategory(): void;
  getCategory(): google_protobuf_wrappers_pb.StringValue | undefined;
  setCategory(value?: google_protobuf_wrappers_pb.StringValue): void;

  getIndex(): number;
  setIndex(value: number): void;

  hasLanguageId(): boolean;
  clearLanguageId(): void;
  getLanguageId(): google_protobuf_wrappers_pb.StringValue | undefined;
  setLanguageId(value?: google_protobuf_wrappers_pb.StringValue): void;

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
    id?: google_protobuf_wrappers_pb.StringValue.AsObject,
    title?: google_protobuf_wrappers_pb.StringValue.AsObject,
    urlTitle?: google_protobuf_wrappers_pb.StringValue.AsObject,
    lessonNumber: number,
    dependenciesList: Array<google_protobuf_wrappers_pb.StringValue.AsObject>,
    disabled: boolean,
    description?: google_protobuf_wrappers_pb.StringValue.AsObject,
    locked: boolean,
    type?: google_protobuf_wrappers_pb.StringValue.AsObject,
    category?: google_protobuf_wrappers_pb.StringValue.AsObject,
    index: number,
    languageId?: google_protobuf_wrappers_pb.StringValue.AsObject,
  }
}

export class CreateResponse extends jspb.Message {
  hasResult(): boolean;
  clearResult(): void;
  getResult(): service_language_proto_entities_entities_pb.Skill | undefined;
  setResult(value?: service_language_proto_entities_entities_pb.Skill): void;

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
    result?: service_language_proto_entities_entities_pb.Skill.AsObject,
  }
}

export class UpdateRequest extends jspb.Message {
  hasId(): boolean;
  clearId(): void;
  getId(): google_protobuf_wrappers_pb.StringValue | undefined;
  setId(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasTitle(): boolean;
  clearTitle(): void;
  getTitle(): google_protobuf_wrappers_pb.StringValue | undefined;
  setTitle(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasUrlTitle(): boolean;
  clearUrlTitle(): void;
  getUrlTitle(): google_protobuf_wrappers_pb.StringValue | undefined;
  setUrlTitle(value?: google_protobuf_wrappers_pb.StringValue): void;

  getLessonNumber(): number;
  setLessonNumber(value: number): void;

  clearDependenciesList(): void;
  getDependenciesList(): Array<google_protobuf_wrappers_pb.StringValue>;
  setDependenciesList(value: Array<google_protobuf_wrappers_pb.StringValue>): void;
  addDependencies(value?: google_protobuf_wrappers_pb.StringValue, index?: number): google_protobuf_wrappers_pb.StringValue;

  getDisabled(): boolean;
  setDisabled(value: boolean): void;

  hasDescription(): boolean;
  clearDescription(): void;
  getDescription(): google_protobuf_wrappers_pb.StringValue | undefined;
  setDescription(value?: google_protobuf_wrappers_pb.StringValue): void;

  getLocked(): boolean;
  setLocked(value: boolean): void;

  hasType(): boolean;
  clearType(): void;
  getType(): google_protobuf_wrappers_pb.StringValue | undefined;
  setType(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasCategory(): boolean;
  clearCategory(): void;
  getCategory(): google_protobuf_wrappers_pb.StringValue | undefined;
  setCategory(value?: google_protobuf_wrappers_pb.StringValue): void;

  getIndex(): number;
  setIndex(value: number): void;

  hasLanguageId(): boolean;
  clearLanguageId(): void;
  getLanguageId(): google_protobuf_wrappers_pb.StringValue | undefined;
  setLanguageId(value?: google_protobuf_wrappers_pb.StringValue): void;

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
    title?: google_protobuf_wrappers_pb.StringValue.AsObject,
    urlTitle?: google_protobuf_wrappers_pb.StringValue.AsObject,
    lessonNumber: number,
    dependenciesList: Array<google_protobuf_wrappers_pb.StringValue.AsObject>,
    disabled: boolean,
    description?: google_protobuf_wrappers_pb.StringValue.AsObject,
    locked: boolean,
    type?: google_protobuf_wrappers_pb.StringValue.AsObject,
    category?: google_protobuf_wrappers_pb.StringValue.AsObject,
    index: number,
    languageId?: google_protobuf_wrappers_pb.StringValue.AsObject,
  }
}

export class UpdateResponse extends jspb.Message {
  hasResult(): boolean;
  clearResult(): void;
  getResult(): service_language_proto_entities_entities_pb.Skill | undefined;
  setResult(value?: service_language_proto_entities_entities_pb.Skill): void;

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
    result?: service_language_proto_entities_entities_pb.Skill.AsObject,
  }
}

export class DeleteRequest extends jspb.Message {
  hasId(): boolean;
  clearId(): void;
  getId(): google_protobuf_wrappers_pb.StringValue | undefined;
  setId(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasTitle(): boolean;
  clearTitle(): void;
  getTitle(): google_protobuf_wrappers_pb.StringValue | undefined;
  setTitle(value?: google_protobuf_wrappers_pb.StringValue): void;

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
    title?: google_protobuf_wrappers_pb.StringValue.AsObject,
  }
}

export class DeleteResponse extends jspb.Message {
  hasResult(): boolean;
  clearResult(): void;
  getResult(): service_language_proto_entities_entities_pb.Skill | undefined;
  setResult(value?: service_language_proto_entities_entities_pb.Skill): void;

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
    result?: service_language_proto_entities_entities_pb.Skill.AsObject,
  }
}

