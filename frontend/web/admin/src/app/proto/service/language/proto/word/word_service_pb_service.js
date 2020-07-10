// package: micro.service.language.word.v1
// file: service/language/proto/word/word_service.proto

var service_language_proto_word_word_service_pb = require("../../../../service/language/proto/word/word_service_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var WordService = (function () {
  function WordService() {}
  WordService.serviceName = "micro.service.language.word.v1.WordService";
  return WordService;
}());

WordService.Exist = {
  methodName: "Exist",
  service: WordService,
  requestStream: false,
  responseStream: false,
  requestType: service_language_proto_word_word_service_pb.ExistRequest,
  responseType: service_language_proto_word_word_service_pb.ExistResponse
};

WordService.List = {
  methodName: "List",
  service: WordService,
  requestStream: false,
  responseStream: false,
  requestType: service_language_proto_word_word_service_pb.ListRequest,
  responseType: service_language_proto_word_word_service_pb.ListResponse
};

WordService.Get = {
  methodName: "Get",
  service: WordService,
  requestStream: false,
  responseStream: false,
  requestType: service_language_proto_word_word_service_pb.GetRequest,
  responseType: service_language_proto_word_word_service_pb.GetResponse
};

WordService.Create = {
  methodName: "Create",
  service: WordService,
  requestStream: false,
  responseStream: false,
  requestType: service_language_proto_word_word_service_pb.CreateRequest,
  responseType: service_language_proto_word_word_service_pb.CreateResponse
};

WordService.Update = {
  methodName: "Update",
  service: WordService,
  requestStream: false,
  responseStream: false,
  requestType: service_language_proto_word_word_service_pb.UpdateRequest,
  responseType: service_language_proto_word_word_service_pb.UpdateResponse
};

WordService.Delete = {
  methodName: "Delete",
  service: WordService,
  requestStream: false,
  responseStream: false,
  requestType: service_language_proto_word_word_service_pb.DeleteRequest,
  responseType: service_language_proto_word_word_service_pb.DeleteResponse
};

exports.WordService = WordService;

function WordServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

WordServiceClient.prototype.exist = function exist(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(WordService.Exist, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

WordServiceClient.prototype.list = function list(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(WordService.List, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

WordServiceClient.prototype.get = function get(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(WordService.Get, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

WordServiceClient.prototype.create = function create(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(WordService.Create, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

WordServiceClient.prototype.update = function update(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(WordService.Update, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

WordServiceClient.prototype.delete = function pb_delete(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(WordService.Delete, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.WordServiceClient = WordServiceClient;

