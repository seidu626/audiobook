// package: micro.service.language.skill.v1
// file: service/language/proto/skill/skill_service.proto

var service_language_proto_skill_skill_service_pb = require("../../../../service/language/proto/skill/skill_service_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var SkillService = (function () {
  function SkillService() {}
  SkillService.serviceName = "micro.service.language.skill.v1.SkillService";
  return SkillService;
}());

SkillService.Exist = {
  methodName: "Exist",
  service: SkillService,
  requestStream: false,
  responseStream: false,
  requestType: service_language_proto_skill_skill_service_pb.ExistRequest,
  responseType: service_language_proto_skill_skill_service_pb.ExistResponse
};

SkillService.List = {
  methodName: "List",
  service: SkillService,
  requestStream: false,
  responseStream: false,
  requestType: service_language_proto_skill_skill_service_pb.ListRequest,
  responseType: service_language_proto_skill_skill_service_pb.ListResponse
};

SkillService.Get = {
  methodName: "Get",
  service: SkillService,
  requestStream: false,
  responseStream: false,
  requestType: service_language_proto_skill_skill_service_pb.GetRequest,
  responseType: service_language_proto_skill_skill_service_pb.GetResponse
};

SkillService.Create = {
  methodName: "Create",
  service: SkillService,
  requestStream: false,
  responseStream: false,
  requestType: service_language_proto_skill_skill_service_pb.CreateRequest,
  responseType: service_language_proto_skill_skill_service_pb.CreateResponse
};

SkillService.Update = {
  methodName: "Update",
  service: SkillService,
  requestStream: false,
  responseStream: false,
  requestType: service_language_proto_skill_skill_service_pb.UpdateRequest,
  responseType: service_language_proto_skill_skill_service_pb.UpdateResponse
};

SkillService.Delete = {
  methodName: "Delete",
  service: SkillService,
  requestStream: false,
  responseStream: false,
  requestType: service_language_proto_skill_skill_service_pb.DeleteRequest,
  responseType: service_language_proto_skill_skill_service_pb.DeleteResponse
};

exports.SkillService = SkillService;

function SkillServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

SkillServiceClient.prototype.exist = function exist(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SkillService.Exist, {
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

SkillServiceClient.prototype.list = function list(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SkillService.List, {
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

SkillServiceClient.prototype.get = function get(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SkillService.Get, {
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

SkillServiceClient.prototype.create = function create(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SkillService.Create, {
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

SkillServiceClient.prototype.update = function update(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SkillService.Update, {
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

SkillServiceClient.prototype.delete = function pb_delete(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SkillService.Delete, {
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

exports.SkillServiceClient = SkillServiceClient;

