// Code generated by Thrift Compiler (0.14.0). DO NOT EDIT.

package edgecontext

import(
	"bytes"
	"context"
	"fmt"
	"time"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

//A raw authentication token as returned by the authentication service.
//
type AuthenticationToken string

func AuthenticationTokenPtr(v AuthenticationToken) *AuthenticationToken { return &v }

//A two-character ISO 3166-1 country code
//
type CountryCode string

func CountryCodePtr(v CountryCode) *CountryCode { return &v }

// The components of the Reddit LoID cookie that we want to propagate between
// services.
// 
// This model is a component of the "Edge-Request" header.  You should not need to
// interact with this model directly, but rather through the EdgeRequestContext
// interface provided by baseplate.
// 
// 
// Attributes:
//  - ID: The ID of the LoID cookie.
// 
//  - CreatedMs: The time when the LoID cookie was created in epoch milliseconds.
// 
type Loid struct {
  ID string `thrift:"id,1" db:"id" json:"id"`
  CreatedMs int64 `thrift:"created_ms,2" db:"created_ms" json:"created_ms"`
}

func NewLoid() *Loid {
  return &Loid{}
}


func (p *Loid) GetID() string {
  return p.ID
}

func (p *Loid) GetCreatedMs() int64 {
  return p.CreatedMs
}
func (p *Loid) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Loid)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.ID = v
}
  return nil
}

func (p *Loid)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.CreatedMs = v
}
  return nil
}

func (p *Loid) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "Loid"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Loid) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "id", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:id: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.ID)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.id (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:id: ", p), err) }
  return err
}

func (p *Loid) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "created_ms", thrift.I64, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:created_ms: ", p), err) }
  if err := oprot.WriteI64(ctx, int64(p.CreatedMs)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.created_ms (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:created_ms: ", p), err) }
  return err
}

func (p *Loid) Equals(other *Loid) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.ID != other.ID { return false }
  if p.CreatedMs != other.CreatedMs { return false }
  return true
}

func (p *Loid) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Loid(%+v)", *p)
}

// The components of the Reddit Session tracker cookie that we want to
// propagate between services.
// 
// This model is a component of the "Edge-Request" header.  You should not need to
// interact with this model directly, but rather through the EdgeRequestContext
// interface provided by baseplate.
// 
// 
// Attributes:
//  - ID: The ID of the Session tracker cookie.
// 
type Session struct {
  ID string `thrift:"id,1" db:"id" json:"id"`
}

func NewSession() *Session {
  return &Session{}
}


func (p *Session) GetID() string {
  return p.ID
}
func (p *Session) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Session)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.ID = v
}
  return nil
}

func (p *Session) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "Session"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Session) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "id", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:id: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.ID)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.id (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:id: ", p), err) }
  return err
}

func (p *Session) Equals(other *Session) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.ID != other.ID { return false }
  return true
}

func (p *Session) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Session(%+v)", *p)
}

// The components of the device making a request to our services that we want to
// propogate between services.
// 
// This model is a component of the "Edge-Request" header.  You should not need to
// interact with this model directly, but rather through the EdgeRequestContext
// interface provided by baseplate.
// 
// 
// Attributes:
//  - ID: The ID of the device.
// 
type Device struct {
  ID string `thrift:"id,1" db:"id" json:"id"`
}

func NewDevice() *Device {
  return &Device{}
}


func (p *Device) GetID() string {
  return p.ID
}
func (p *Device) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Device)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.ID = v
}
  return nil
}

func (p *Device) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "Device"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Device) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "id", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:id: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.ID)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.id (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:id: ", p), err) }
  return err
}

func (p *Device) Equals(other *Device) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.ID != other.ID { return false }
  return true
}

func (p *Device) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Device(%+v)", *p)
}

// Metadata about the origin service for a request.
// 
// The "origin" service is the service responsible for handling the request from
// the client.
// 
// This model is a component of the "Edge-Request" header.  You should not need to
// interact with this model directly, but rather through the EdgeRequestContext
// interface provided by baseplate.
// 
// Attributes:
//  - Name: The name of the origin service.
// 
type OriginService struct {
  Name string `thrift:"name,1" db:"name" json:"name"`
}

func NewOriginService() *OriginService {
  return &OriginService{}
}


func (p *OriginService) GetName() string {
  return p.Name
}
func (p *OriginService) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *OriginService)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Name = v
}
  return nil
}

func (p *OriginService) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "OriginService"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *OriginService) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "name", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:name: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.Name)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.name (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:name: ", p), err) }
  return err
}

func (p *OriginService) Equals(other *OriginService) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.Name != other.Name { return false }
  return true
}

func (p *OriginService) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("OriginService(%+v)", *p)
}

// Geolocation data from a request to our services that we want to
// propagate between services.
// 
// This model is a component of the "Edge-Request" header.  You should not need to
// interact with this model directly, but rather through the EdgeRequestContext
// interface provided by baseplate.
// 
// 
// Attributes:
//  - CountryCode: The country code of the requesting client.
type Geolocation struct {
  CountryCode CountryCode `thrift:"country_code,1" db:"country_code" json:"country_code"`
}

func NewGeolocation() *Geolocation {
  return &Geolocation{}
}


func (p *Geolocation) GetCountryCode() CountryCode {
  return p.CountryCode
}
func (p *Geolocation) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Geolocation)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := CountryCode(v)
  p.CountryCode = temp
}
  return nil
}

func (p *Geolocation) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "Geolocation"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Geolocation) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "country_code", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:country_code: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.CountryCode)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.country_code (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:country_code: ", p), err) }
  return err
}

func (p *Geolocation) Equals(other *Geolocation) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.CountryCode != other.CountryCode { return false }
  return true
}

func (p *Geolocation) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Geolocation(%+v)", *p)
}

// Container model for the Edge-Request context header.
// 
// Baseplate will automatically parse this from the "Edge-Request" header and
// provides an interface that wraps this Thrift model.  You should not need to
// interact with this model directly, but rather through the EdgeRequestContext
// interface provided by baseplate.
// 
// 
// Attributes:
//  - Loid
//  - Session
//  - AuthenticationToken
//  - Device
//  - OriginService
//  - Geolocation
type Request struct {
  Loid *Loid `thrift:"loid,1" db:"loid" json:"loid"`
  Session *Session `thrift:"session,2" db:"session" json:"session"`
  AuthenticationToken AuthenticationToken `thrift:"authentication_token,3" db:"authentication_token" json:"authentication_token"`
  Device *Device `thrift:"device,4" db:"device" json:"device"`
  OriginService *OriginService `thrift:"origin_service,5" db:"origin_service" json:"origin_service"`
  Geolocation *Geolocation `thrift:"geolocation,6" db:"geolocation" json:"geolocation"`
}

func NewRequest() *Request {
  return &Request{}
}

var Request_Loid_DEFAULT *Loid
func (p *Request) GetLoid() *Loid {
  if !p.IsSetLoid() {
    return Request_Loid_DEFAULT
  }
return p.Loid
}
var Request_Session_DEFAULT *Session
func (p *Request) GetSession() *Session {
  if !p.IsSetSession() {
    return Request_Session_DEFAULT
  }
return p.Session
}

func (p *Request) GetAuthenticationToken() AuthenticationToken {
  return p.AuthenticationToken
}
var Request_Device_DEFAULT *Device
func (p *Request) GetDevice() *Device {
  if !p.IsSetDevice() {
    return Request_Device_DEFAULT
  }
return p.Device
}
var Request_OriginService_DEFAULT *OriginService
func (p *Request) GetOriginService() *OriginService {
  if !p.IsSetOriginService() {
    return Request_OriginService_DEFAULT
  }
return p.OriginService
}
var Request_Geolocation_DEFAULT *Geolocation
func (p *Request) GetGeolocation() *Geolocation {
  if !p.IsSetGeolocation() {
    return Request_Geolocation_DEFAULT
  }
return p.Geolocation
}
func (p *Request) IsSetLoid() bool {
  return p.Loid != nil
}

func (p *Request) IsSetSession() bool {
  return p.Session != nil
}

func (p *Request) IsSetDevice() bool {
  return p.Device != nil
}

func (p *Request) IsSetOriginService() bool {
  return p.OriginService != nil
}

func (p *Request) IsSetGeolocation() bool {
  return p.Geolocation != nil
}

func (p *Request) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField3(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 4:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField4(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 5:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField5(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 6:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField6(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Request)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  p.Loid = &Loid{}
  if err := p.Loid.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Loid), err)
  }
  return nil
}

func (p *Request)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  p.Session = &Session{}
  if err := p.Session.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Session), err)
  }
  return nil
}

func (p *Request)  ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  temp := AuthenticationToken(v)
  p.AuthenticationToken = temp
}
  return nil
}

func (p *Request)  ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
  p.Device = &Device{}
  if err := p.Device.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Device), err)
  }
  return nil
}

func (p *Request)  ReadField5(ctx context.Context, iprot thrift.TProtocol) error {
  p.OriginService = &OriginService{}
  if err := p.OriginService.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.OriginService), err)
  }
  return nil
}

func (p *Request)  ReadField6(ctx context.Context, iprot thrift.TProtocol) error {
  p.Geolocation = &Geolocation{}
  if err := p.Geolocation.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Geolocation), err)
  }
  return nil
}

func (p *Request) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "Request"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
    if err := p.writeField3(ctx, oprot); err != nil { return err }
    if err := p.writeField4(ctx, oprot); err != nil { return err }
    if err := p.writeField5(ctx, oprot); err != nil { return err }
    if err := p.writeField6(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Request) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "loid", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:loid: ", p), err) }
  if err := p.Loid.Write(ctx, oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Loid), err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:loid: ", p), err) }
  return err
}

func (p *Request) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "session", thrift.STRUCT, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:session: ", p), err) }
  if err := p.Session.Write(ctx, oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Session), err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:session: ", p), err) }
  return err
}

func (p *Request) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "authentication_token", thrift.STRING, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:authentication_token: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.AuthenticationToken)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.authentication_token (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:authentication_token: ", p), err) }
  return err
}

func (p *Request) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "device", thrift.STRUCT, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:device: ", p), err) }
  if err := p.Device.Write(ctx, oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Device), err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:device: ", p), err) }
  return err
}

func (p *Request) writeField5(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "origin_service", thrift.STRUCT, 5); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:origin_service: ", p), err) }
  if err := p.OriginService.Write(ctx, oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.OriginService), err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 5:origin_service: ", p), err) }
  return err
}

func (p *Request) writeField6(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "geolocation", thrift.STRUCT, 6); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:geolocation: ", p), err) }
  if err := p.Geolocation.Write(ctx, oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Geolocation), err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 6:geolocation: ", p), err) }
  return err
}

func (p *Request) Equals(other *Request) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if !p.Loid.Equals(other.Loid) { return false }
  if !p.Session.Equals(other.Session) { return false }
  if p.AuthenticationToken != other.AuthenticationToken { return false }
  if !p.Device.Equals(other.Device) { return false }
  if !p.OriginService.Equals(other.OriginService) { return false }
  if !p.Geolocation.Equals(other.Geolocation) { return false }
  return true
}

func (p *Request) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Request(%+v)", *p)
}

