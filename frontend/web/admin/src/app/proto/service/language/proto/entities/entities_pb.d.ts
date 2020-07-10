// package: micro.service.language.entities.v1
// file: service/language/proto/entities/entities.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as google_protobuf_wrappers_pb from "google-protobuf/google/protobuf/wrappers_pb";
import * as third_party_proto_protoc_gen_gorm_options_gorm_pb from "../../../../backend/third_party/proto/protoc-gen-gorm/options/gorm_pb";
import * as third_party_proto_protoc_gen_gorm_types_types_pb from "../../../../backend/third_party/proto/protoc-gen-gorm/types/types_pb";
import * as third_party_proto_validate_validate_pb from "../../../../backend/third_party/proto/validate/validate_pb";

export class Exercise extends jspb.Message {
  hasId(): boolean;
  clearId(): void;
  getId(): third_party_proto_protoc_gen_gorm_types_types_pb.UUID | undefined;
  setId(value?: third_party_proto_protoc_gen_gorm_types_types_pb.UUID): void;

  hasCreatedAt(): boolean;
  clearCreatedAt(): void;
  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasUpdatedAt(): boolean;
  clearUpdatedAt(): void;
  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasDeletedAt(): boolean;
  clearDeletedAt(): void;
  getDeletedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDeletedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getContent(): string;
  setContent(value: string): void;

  getAudioSrc(): string;
  setAudioSrc(value: string): void;

  hasSkill(): boolean;
  clearSkill(): void;
  getSkill(): Skill | undefined;
  setSkill(value?: Skill): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Exercise.AsObject;
  static toObject(includeInstance: boolean, msg: Exercise): Exercise.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Exercise, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Exercise;
  static deserializeBinaryFromReader(message: Exercise, reader: jspb.BinaryReader): Exercise;
}

export namespace Exercise {
  export type AsObject = {
    id?: third_party_proto_protoc_gen_gorm_types_types_pb.UUID.AsObject,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    deletedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    content: string,
    audioSrc: string,
    skill?: Skill.AsObject,
  }
}

export class Skill extends jspb.Message {
  hasId(): boolean;
  clearId(): void;
  getId(): third_party_proto_protoc_gen_gorm_types_types_pb.UUID | undefined;
  setId(value?: third_party_proto_protoc_gen_gorm_types_types_pb.UUID): void;

  hasCreatedAt(): boolean;
  clearCreatedAt(): void;
  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasUpdatedAt(): boolean;
  clearUpdatedAt(): void;
  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasDeletedAt(): boolean;
  clearDeletedAt(): void;
  getDeletedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDeletedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasDependencies(): boolean;
  clearDependencies(): void;
  getDependencies(): google_protobuf_wrappers_pb.StringValue | undefined;
  setDependencies(value?: google_protobuf_wrappers_pb.StringValue): void;

  getDisabled(): boolean;
  setDisabled(value: boolean): void;

  getLessonNumber(): number;
  setLessonNumber(value: number): void;

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

  getTitle(): string;
  setTitle(value: string): void;

  getUrlTitle(): string;
  setUrlTitle(value: string): void;

  hasCategory(): boolean;
  clearCategory(): void;
  getCategory(): google_protobuf_wrappers_pb.StringValue | undefined;
  setCategory(value?: google_protobuf_wrappers_pb.StringValue): void;

  getIndex(): number;
  setIndex(value: number): void;

  clearWordsList(): void;
  getWordsList(): Array<Word>;
  setWordsList(value: Array<Word>): void;
  addWords(value?: Word, index?: number): Word;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Skill.AsObject;
  static toObject(includeInstance: boolean, msg: Skill): Skill.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Skill, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Skill;
  static deserializeBinaryFromReader(message: Skill, reader: jspb.BinaryReader): Skill;
}

export namespace Skill {
  export type AsObject = {
    id?: third_party_proto_protoc_gen_gorm_types_types_pb.UUID.AsObject,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    deletedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    dependencies?: google_protobuf_wrappers_pb.StringValue.AsObject,
    disabled: boolean,
    lessonNumber: number,
    description?: google_protobuf_wrappers_pb.StringValue.AsObject,
    locked: boolean,
    type?: google_protobuf_wrappers_pb.StringValue.AsObject,
    title: string,
    urlTitle: string,
    category?: google_protobuf_wrappers_pb.StringValue.AsObject,
    index: number,
    wordsList: Array<Word.AsObject>,
  }
}

export class Word extends jspb.Message {
  hasId(): boolean;
  clearId(): void;
  getId(): third_party_proto_protoc_gen_gorm_types_types_pb.UUID | undefined;
  setId(value?: third_party_proto_protoc_gen_gorm_types_types_pb.UUID): void;

  hasCreatedAt(): boolean;
  clearCreatedAt(): void;
  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasUpdatedAt(): boolean;
  clearUpdatedAt(): void;
  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasDeletedAt(): boolean;
  clearDeletedAt(): void;
  getDeletedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDeletedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getContent(): string;
  setContent(value: string): void;

  getAudioSrc(): string;
  setAudioSrc(value: string): void;

  hasSkill(): boolean;
  clearSkill(): void;
  getSkill(): Skill | undefined;
  setSkill(value?: Skill): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Word.AsObject;
  static toObject(includeInstance: boolean, msg: Word): Word.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Word, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Word;
  static deserializeBinaryFromReader(message: Word, reader: jspb.BinaryReader): Word;
}

export namespace Word {
  export type AsObject = {
    id?: third_party_proto_protoc_gen_gorm_types_types_pb.UUID.AsObject,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    deletedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    content: string,
    audioSrc: string,
    skill?: Skill.AsObject,
  }
}

export class Language extends jspb.Message {
  hasId(): boolean;
  clearId(): void;
  getId(): third_party_proto_protoc_gen_gorm_types_types_pb.UUID | undefined;
  setId(value?: third_party_proto_protoc_gen_gorm_types_types_pb.UUID): void;

  hasCreatedAt(): boolean;
  clearCreatedAt(): void;
  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasUpdatedAt(): boolean;
  clearUpdatedAt(): void;
  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasDeletedAt(): boolean;
  clearDeletedAt(): void;
  getDeletedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDeletedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getName(): string;
  setName(value: string): void;

  getAbbreviation(): string;
  setAbbreviation(value: string): void;

  getFlagSrc(): string;
  setFlagSrc(value: string): void;

  clearSkillsList(): void;
  getSkillsList(): Array<Skill>;
  setSkillsList(value: Array<Skill>): void;
  addSkills(value?: Skill, index?: number): Skill;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Language.AsObject;
  static toObject(includeInstance: boolean, msg: Language): Language.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Language, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Language;
  static deserializeBinaryFromReader(message: Language, reader: jspb.BinaryReader): Language;
}

export namespace Language {
  export type AsObject = {
    id?: third_party_proto_protoc_gen_gorm_types_types_pb.UUID.AsObject,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    deletedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    name: string,
    abbreviation: string,
    flagSrc: string,
    skillsList: Array<Skill.AsObject>,
  }
}

