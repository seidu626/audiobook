// package: micro.service.language.language.v1
// file: service/language/proto/language/language_service.proto

var service_language_proto_language_language_service_pb = require("../../../../service/language/proto/language/language_service_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var LanguageService = (function () {
  function LanguageService() {}
  LanguageService.serviceName = "micro.service.language.language.v1.LanguageService";
  return LanguageService;
}());

LanguageService.Exist = {
  methodName: "Exist",
  service: LanguageService,
  requestStream: false,
  responseStream: false,
  requestType: service_language_proto_language_language_service_pb.ExistRequest,
  responseType: service_language_proto_language_language_service_pb.ExistResponse
};

LanguageService.List = {
  methodName: "List",
  service: LanguageService,
  requestStream: false,
  responseStream: false,
  requestType: service_language_proto_language_language_service_pb.ListRequest,
  responseType: service_language_proto_language_language_service_pb.ListResponse
};

LanguageService.Get = {
  methodName: "Get",
  service: LanguageService,
  requestStream: false,
  responseStream: false,
  requestType: service_language_proto_language_language_service_pb.GetRequest,
  responseType: service_language_proto_language_language_service_pb.GetResponse
};

LanguageService.Create = {
  methodName: "Create",
  service: LanguageService,
  requestStream: false,
  responseStream: false,
  requestType: service_language_proto_language_language_service_pb.CreateRequest,
  responseType: service_language_proto_language_language_service_pb.CreateResponse
};

LanguageService.Update = {
  methodName: "Update",
  service: LanguageService,
  requestStream: false,
  responseStream: false,
  requestType: service_language_proto_language_language_service_pb.UpdateRequest,
  responseType: service_language_proto_language_language_service_pb.UpdateResponse
};

LanguageService.Delete = {
  methodName: "Delete",
  service: LanguageService,
  requestStream: false,
  responseStream: false,
  requestType: service_language_proto_language_language_service_pb.DeleteRequest,
  responseType: service_language_proto_language_language_service_pb.DeleteResponse
};

exports.LanguageService = LanguageService;

function LanguageServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

LanguageServiceClient.prototype.exist = function exist(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(LanguageService.Exist, {
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

LanguageServiceClient.prototype.list = function list(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(LanguageService.List, {
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

LanguageServiceClient.prototype.get = function get(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(LanguageService.Get, {
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

LanguageServiceClient.prototype.create = function create(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(LanguageService.Create, {
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

LanguageServiceClient.prototype.update = function update(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(LanguageService.Update, {
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

LanguageServiceClient.prototype.delete = function pb_delete(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(LanguageService.Delete, {
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

exports.LanguageServiceClient = LanguageServiceClient;

