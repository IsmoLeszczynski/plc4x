/**
 * Autogenerated by Thrift Compiler (0.12.0)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
using System;
using System.Collections;
using System.Collections.Generic;
using System.Text;
using System.IO;
using Thrift;
using Thrift.Collections;
using System.Runtime.Serialization;
using Thrift.Protocol;
using Thrift.Transport;

namespace org.apache.plc4x.interop
{

  #if !SILVERLIGHT
  [Serializable]
  #endif
  public partial class Response : TBase
  {
    private Dictionary<string, FieldResponse> _fields;

    public Dictionary<string, FieldResponse> Fields
    {
      get
      {
        return _fields;
      }
      set
      {
        __isset.fields = true;
        this._fields = value;
      }
    }


    public Isset __isset;
    #if !SILVERLIGHT
    [Serializable]
    #endif
    public struct Isset {
      public bool fields;
    }

    public Response() {
    }

    public void Read (TProtocol iprot)
    {
      iprot.IncrementRecursionDepth();
      try
      {
        TField field;
        iprot.ReadStructBegin();
        while (true)
        {
          field = iprot.ReadFieldBegin();
          if (field.Type == TType.Stop) { 
            break;
          }
          switch (field.ID)
          {
            case 1:
              if (field.Type == TType.Map) {
                {
                  Fields = new Dictionary<string, FieldResponse>();
                  TMap _map5 = iprot.ReadMapBegin();
                  for( int _i6 = 0; _i6 < _map5.Count; ++_i6)
                  {
                    string _key7;
                    FieldResponse _val8;
                    _key7 = iprot.ReadString();
                    _val8 = new FieldResponse();
                    _val8.Read(iprot);
                    Fields[_key7] = _val8;
                  }
                  iprot.ReadMapEnd();
                }
              } else { 
                TProtocolUtil.Skip(iprot, field.Type);
              }
              break;
            default: 
              TProtocolUtil.Skip(iprot, field.Type);
              break;
          }
          iprot.ReadFieldEnd();
        }
        iprot.ReadStructEnd();
      }
      finally
      {
        iprot.DecrementRecursionDepth();
      }
    }

    public void Write(TProtocol oprot) {
      oprot.IncrementRecursionDepth();
      try
      {
        TStruct struc = new TStruct("Response");
        oprot.WriteStructBegin(struc);
        TField field = new TField();
        if (Fields != null && __isset.fields) {
          field.Name = "fields";
          field.Type = TType.Map;
          field.ID = 1;
          oprot.WriteFieldBegin(field);
          {
            oprot.WriteMapBegin(new TMap(TType.String, TType.Struct, Fields.Count));
            foreach (string _iter9 in Fields.Keys)
            {
              oprot.WriteString(_iter9);
              Fields[_iter9].Write(oprot);
            }
            oprot.WriteMapEnd();
          }
          oprot.WriteFieldEnd();
        }
        oprot.WriteFieldStop();
        oprot.WriteStructEnd();
      }
      finally
      {
        oprot.DecrementRecursionDepth();
      }
    }

    public override string ToString() {
      StringBuilder __sb = new StringBuilder("Response(");
      bool __first = true;
      if (Fields != null && __isset.fields) {
        if(!__first) { __sb.Append(", "); }
        __first = false;
        __sb.Append("Fields: ");
        __sb.Append(Fields);
      }
      __sb.Append(")");
      return __sb.ToString();
    }

  }

}
