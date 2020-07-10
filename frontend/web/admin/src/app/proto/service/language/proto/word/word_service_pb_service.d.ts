// package: micro.service.language.word.v1
// file: service/language/proto/word/word_service.proto

import * as service_language_proto_word_word_service_pb from "../../../../service/language/proto/word/word_service_pb";
import {grpc} from "@improbable-eng/grpc-web";

type WordServiceExist = {
  readonly methodName: string;
  readonly service: typeof WordService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_language_proto_word_word_service_pb.ExistRequest;
  readonly responseType: typeof service_language_proto_word_word_service_pb.ExistResponse;
};

type WordServiceList = {
  readonly methodName: string;
  readonly service: typeof WordService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_language_proto_word_word_service_pb.ListRequest;
  readonly responseType: typeof service_language_proto_word_word_service_pb.ListResponse;
};

type WordServiceGet = {
  readonly methodName: string;
  readonly service: typeof WordService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_language_proto_word_word_service_pb.GetRequest;
  readonly responseType: typeof service_language_proto_word_word_service_pb.GetResponse;
};

type WordServiceCreate = {
  readonly methodName: string;
  readonly service: typeof WordService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_language_proto_word_word_service_pb.CreateRequest;
  readonly responseType: typeof service_language_proto_word_word_service_pb.CreateResponse;
};

type WordServiceUpdate = {
  readonly methodName: string;
  readonly service: typeof WordService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_language_proto_word_word_service_pb.UpdateRequest;
  readonly responseType: typeof service_language_proto_word_word_service_pb.UpdateResponse;
};

type WordServiceDelete = {
  readonly methodName: string;
  readonly service: typeof WordService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_language_proto_word_word_service_pb.DeleteRequest;
  readonly responseType: typeof service_language_proto_word_word_service_pb.DeleteResponse;
};

export class WordService {
  static readonly serviceName: string;
  static readonly Exist: WordServiceExist;
  static readonly List: WordServiceList;
  static readonly Get: WordServiceGet;
  static readonly Create: WordServiceCreate;
  static readonly Update: WordServiceUpdate;
  static readonly Delete: WordServiceDelete;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class WordServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  exist(
    requestMessage: service_language_proto_word_word_service_pb.ExistRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_language_proto_word_word_service_pb.ExistResponse|null) => void
  ): UnaryResponse;
  exist(
    requestMessage: service_language_proto_word_word_service_pb.ExistRequest,
    callback: (error: ServiceError|null, responseMessage: service_language_proto_word_word_service_pb.ExistResponse|null) => void
  ): UnaryResponse;
  list(
    requestMessage: service_language_proto_word_word_service_pb.ListRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_language_proto_word_word_service_pb.ListResponse|null) => void
  ): UnaryResponse;
  list(
    requestMessage: service_language_proto_word_word_service_pb.ListRequest,
    callback: (error: ServiceError|null, responseMessage: service_language_proto_word_word_service_pb.ListResponse|null) => void
  ): UnaryResponse;
  get(
    requestMessage: service_language_proto_word_word_service_pb.GetRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_language_proto_word_word_service_pb.GetResponse|null) => void
  ): UnaryResponse;
  get(
    requestMessage: service_language_proto_word_word_service_pb.GetRequest,
    callback: (error: ServiceError|null, responseMessage: service_language_proto_word_word_service_pb.GetResponse|null) => void
  ): UnaryResponse;
  create(
    requestMessage: service_language_proto_word_word_service_pb.CreateRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_language_proto_word_word_service_pb.CreateResponse|null) => void
  ): UnaryResponse;
  create(
    requestMessage: service_language_proto_word_word_service_pb.CreateRequest,
    callback: (error: ServiceError|null, responseMessage: service_language_proto_word_word_service_pb.CreateResponse|null) => void
  ): UnaryResponse;
  update(
    requestMessage: service_language_proto_word_word_service_pb.UpdateRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_language_proto_word_word_service_pb.UpdateResponse|null) => void
  ): UnaryResponse;
  update(
    requestMessage: service_language_proto_word_word_service_pb.UpdateRequest,
    callback: (error: ServiceError|null, responseMessage: service_language_proto_word_word_service_pb.UpdateResponse|null) => void
  ): UnaryResponse;
  delete(
    requestMessage: service_language_proto_word_word_service_pb.DeleteRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_language_proto_word_word_service_pb.DeleteResponse|null) => void
  ): UnaryResponse;
  delete(
    requestMessage: service_language_proto_word_word_service_pb.DeleteRequest,
    callback: (error: ServiceError|null, responseMessage: service_language_proto_word_word_service_pb.DeleteResponse|null) => void
  ): UnaryResponse;
}

