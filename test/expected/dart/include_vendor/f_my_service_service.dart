// Autogenerated by Frugal Compiler (2.3.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING



import 'dart:async';

import 'dart:typed_data' show Uint8List;
import 'package:logging/logging.dart' as logging;
import 'package:thrift/thrift.dart' as thrift;
import 'package:frugal/frugal.dart' as frugal;

import 'package:some_vendored_place/vendor_namespace.dart' as t_vendor_namespace;
import 'package:excepts/excepts.dart' as t_excepts;
import 'package:include_vendor/include_vendor.dart' as t_include_vendor;


abstract class FMyService {

  Future<t_vendor_namespace.Item> getItem(frugal.FContext ctx);
}

class FMyServiceClient implements FMyService {
  static final logging.Logger _log = new logging.Logger('MyService');
  Map<String, frugal.FMethod> _methods;

  FMyServiceClient(frugal.FServiceProvider provider, [List<frugal.Middleware> middleware]) {
    _transport = provider.transport;
    _protocolFactory = provider.protocolFactory;
    var combined = middleware ?? [];
    combined.addAll(provider.middleware);
    this._methods = {};
    this._methods['getItem'] = new frugal.FMethod(this._getItem, 'MyService', 'getItem', combined);
  }

  frugal.FTransport _transport;
  frugal.FProtocolFactory _protocolFactory;

  Future<t_vendor_namespace.Item> getItem(frugal.FContext ctx) {
    return this._methods['getItem']([ctx]) as Future<t_vendor_namespace.Item>;
  }

  Future<t_vendor_namespace.Item> _getItem(frugal.FContext ctx) async {
    var memoryBuffer = new frugal.TMemoryOutputBuffer(_transport.requestSizeLimit);
    var oprot = _protocolFactory.getProtocol(memoryBuffer);
    oprot.writeRequestHeader(ctx);
    oprot.writeMessageBegin(new thrift.TMessage("getItem", thrift.TMessageType.CALL, 0));
    getItem_args args = new getItem_args();
    args.write(oprot);
    oprot.writeMessageEnd();
    var response = await _transport.request(ctx, memoryBuffer.writeBytes);

    var iprot = _protocolFactory.getProtocol(response);
    iprot.readResponseHeader(ctx);
    thrift.TMessage msg = iprot.readMessageBegin();
    if (msg.type == thrift.TMessageType.EXCEPTION) {
      thrift.TApplicationError error = thrift.TApplicationError.read(iprot);
      iprot.readMessageEnd();
      if (error.type == frugal.FrugalTTransportErrorType.REQUEST_TOO_LARGE) {
        throw new thrift.TTransportError(frugal.FrugalTTransportErrorType.RESPONSE_TOO_LARGE, error.message);
      }
      throw error;
    }

    getItem_result result = new getItem_result();
    result.read(iprot);
    iprot.readMessageEnd();
    if (result.isSetSuccess()) {
      return result.success;
    }

    if (result.d != null) {
      throw result.d;
    }
    throw new thrift.TApplicationError(
      frugal.FrugalTApplicationErrorType.MISSING_RESULT, "getItem failed: unknown result"
    );
  }
}

class getItem_args implements thrift.TBase {
  static final thrift.TStruct _STRUCT_DESC = new thrift.TStruct("getItem_args");



  getItem_args() {
  }

  getFieldValue(int fieldID) {
    switch (fieldID) {
      default:
        throw new ArgumentError("Field $fieldID doesn't exist!");
    }
  }

  setFieldValue(int fieldID, Object value) {
    switch(fieldID) {
      default:
        throw new ArgumentError("Field $fieldID doesn't exist!");
    }
  }

  // Returns true if the field corresponding to fieldID is set (has been assigned a value) and false otherwise
  bool isSet(int fieldID) {
    switch(fieldID) {
      default:
        throw new ArgumentError("Field $fieldID doesn't exist!");
    }
  }

  read(thrift.TProtocol iprot) {
    thrift.TField field;
    iprot.readStructBegin();
    while(true) {
      field = iprot.readFieldBegin();
      if(field.type == thrift.TType.STOP) {
        break;
      }
      switch(field.id) {
        default:
          thrift.TProtocolUtil.skip(iprot, field.type);
          break;
      }
      iprot.readFieldEnd();
    }
    iprot.readStructEnd();

    // check for required fields of primitive type, which can't be checked in the validate method
    validate();
  }

  write(thrift.TProtocol oprot) {
    validate();

    oprot.writeStructBegin(_STRUCT_DESC);
    oprot.writeFieldStop();
    oprot.writeStructEnd();
  }

  String toString() {
    StringBuffer ret = new StringBuffer("getItem_args(");

    ret.write(")");

    return ret.toString();
  }

  validate() {
    // check for required fields
    // check that fields of type enum have valid values
  }
}
class getItem_result implements thrift.TBase {
  static final thrift.TStruct _STRUCT_DESC = new thrift.TStruct("getItem_result");
  static final thrift.TField _SUCCESS_FIELD_DESC = new thrift.TField("success", thrift.TType.STRUCT, 0);
  static final thrift.TField _D_FIELD_DESC = new thrift.TField("d", thrift.TType.STRUCT, 1);

  t_vendor_namespace.Item _success;
  static const int SUCCESS = 0;
  t_excepts.InvalidData _d;
  static const int D = 1;


  getItem_result() {
  }

  t_vendor_namespace.Item get success => this._success;

  set success(t_vendor_namespace.Item success) {
    this._success = success;
  }

  bool isSetSuccess() => this.success != null;

  unsetSuccess() {
    this.success = null;
  }

  t_excepts.InvalidData get d => this._d;

  set d(t_excepts.InvalidData d) {
    this._d = d;
  }

  bool isSetD() => this.d != null;

  unsetD() {
    this.d = null;
  }

  getFieldValue(int fieldID) {
    switch (fieldID) {
      case SUCCESS:
        return this.success;
      case D:
        return this.d;
      default:
        throw new ArgumentError("Field $fieldID doesn't exist!");
    }
  }

  setFieldValue(int fieldID, Object value) {
    switch(fieldID) {
      case SUCCESS:
        if(value == null) {
          unsetSuccess();
        } else {
          this.success = value as t_vendor_namespace.Item;
        }
        break;

      case D:
        if(value == null) {
          unsetD();
        } else {
          this.d = value as t_excepts.InvalidData;
        }
        break;

      default:
        throw new ArgumentError("Field $fieldID doesn't exist!");
    }
  }

  // Returns true if the field corresponding to fieldID is set (has been assigned a value) and false otherwise
  bool isSet(int fieldID) {
    switch(fieldID) {
      case SUCCESS:
        return isSetSuccess();
      case D:
        return isSetD();
      default:
        throw new ArgumentError("Field $fieldID doesn't exist!");
    }
  }

  read(thrift.TProtocol iprot) {
    thrift.TField field;
    iprot.readStructBegin();
    while(true) {
      field = iprot.readFieldBegin();
      if(field.type == thrift.TType.STOP) {
        break;
      }
      switch(field.id) {
        case SUCCESS:
          if(field.type == thrift.TType.STRUCT) {
            success = new t_vendor_namespace.Item();
            success.read(iprot);
          } else {
            thrift.TProtocolUtil.skip(iprot, field.type);
          }
          break;
        case D:
          if(field.type == thrift.TType.STRUCT) {
            d = new t_excepts.InvalidData();
            d.read(iprot);
          } else {
            thrift.TProtocolUtil.skip(iprot, field.type);
          }
          break;
        default:
          thrift.TProtocolUtil.skip(iprot, field.type);
          break;
      }
      iprot.readFieldEnd();
    }
    iprot.readStructEnd();

    // check for required fields of primitive type, which can't be checked in the validate method
    validate();
  }

  write(thrift.TProtocol oprot) {
    validate();

    oprot.writeStructBegin(_STRUCT_DESC);
    if(isSetSuccess() && this.success != null) {
      oprot.writeFieldBegin(_SUCCESS_FIELD_DESC);
      success.write(oprot);
      oprot.writeFieldEnd();
    }
    if(isSetD() && this.d != null) {
      oprot.writeFieldBegin(_D_FIELD_DESC);
      d.write(oprot);
      oprot.writeFieldEnd();
    }
    oprot.writeFieldStop();
    oprot.writeStructEnd();
  }

  String toString() {
    StringBuffer ret = new StringBuffer("getItem_result(");

    if(isSetSuccess()) {
      ret.write("success:");
      if(this.success == null) {
        ret.write("null");
      } else {
        ret.write(this.success);
      }
    }

    if(isSetD()) {
      ret.write(", ");
      ret.write("d:");
      if(this.d == null) {
        ret.write("null");
      } else {
        ret.write(this.d);
      }
    }

    ret.write(")");

    return ret.toString();
  }

  validate() {
    // check for required fields
    // check that fields of type enum have valid values
  }
}
