// Autogenerated by Thrift Compiler (facebook)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
// @generated

package graph

import (
	"bytes"
	"sync"
	"fmt"
	thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift"
	nebula0 "github.com/vesoft-inc/nebula-go/nebula"

)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = sync.Mutex{}
var _ = bytes.Equal

var _ = nebula0.GoUnusedProtection__
type GraphService interface {
  // Parameters:
  //  - Username
  //  - Password
  Authenticate(username string, password string) (r *AuthResponse, err error)
  // Parameters:
  //  - SessionId
  Signout(sessionId int64) (err error)
  // Parameters:
  //  - SessionId
  //  - Stmt
  Execute(sessionId int64, stmt string) (r *ExecutionResponse, err error)
}

type GraphServiceClient struct {
  Transport thrift.Transport
  ProtocolFactory thrift.ProtocolFactory
  InputProtocol thrift.Protocol
  OutputProtocol thrift.Protocol
  SeqId int32
}

func (client *GraphServiceClient) Close() error {
  return client.Transport.Close()
}

func NewGraphServiceClientFactory(t thrift.Transport, f thrift.ProtocolFactory) *GraphServiceClient {
  return &GraphServiceClient{Transport: t,
    ProtocolFactory: f,
    InputProtocol: f.GetProtocol(t),
    OutputProtocol: f.GetProtocol(t),
    SeqId: 0,
  }
}

func NewGraphServiceClient(t thrift.Transport, iprot thrift.Protocol, oprot thrift.Protocol) *GraphServiceClient {
  return &GraphServiceClient{Transport: t,
    ProtocolFactory: nil,
    InputProtocol: iprot,
    OutputProtocol: oprot,
    SeqId: 0,
  }
}

// Parameters:
//  - Username
//  - Password
func (p *GraphServiceClient) Authenticate(username string, password string) (r *AuthResponse, err error) {
  if err = p.sendAuthenticate(username, password); err != nil { return }
  return p.recvAuthenticate()
}

func (p *GraphServiceClient) sendAuthenticate(username string, password string)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("authenticate", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := GraphServiceAuthenticateArgs{
  Username : username,
  Password : password,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *GraphServiceClient) recvAuthenticate() (value *AuthResponse, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "authenticate" {
    err = thrift.NewApplicationException(thrift.WRONG_METHOD_NAME, "authenticate failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewApplicationException(thrift.BAD_SEQUENCE_ID, "authenticate failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error5 := thrift.NewApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error6 error
    error6, err = error5.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error6
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "authenticate failed: invalid message type")
    return
  }
  result := GraphServiceAuthenticateResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}

// Parameters:
//  - SessionId
func (p *GraphServiceClient) Signout(sessionId int64) (err error) {
  if err = p.sendSignout(sessionId); err != nil { return }
  return
}

func (p *GraphServiceClient) sendSignout(sessionId int64)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("signout", thrift.ONEWAY, p.SeqId); err != nil {
      return
  }
  args := GraphServiceSignoutArgs{
  SessionId : sessionId,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}

// Parameters:
//  - SessionId
//  - Stmt
func (p *GraphServiceClient) Execute(sessionId int64, stmt string) (r *ExecutionResponse, err error) {
  if err = p.sendExecute(sessionId, stmt); err != nil { return }
  return p.recvExecute()
}

func (p *GraphServiceClient) sendExecute(sessionId int64, stmt string)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("execute", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := GraphServiceExecuteArgs{
  SessionId : sessionId,
  Stmt : stmt,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *GraphServiceClient) recvExecute() (value *ExecutionResponse, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "execute" {
    err = thrift.NewApplicationException(thrift.WRONG_METHOD_NAME, "execute failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewApplicationException(thrift.BAD_SEQUENCE_ID, "execute failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error7 := thrift.NewApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error8 error
    error8, err = error7.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error8
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "execute failed: invalid message type")
    return
  }
  result := GraphServiceExecuteResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}


type GraphServiceThreadsafeClient struct {
  Transport thrift.Transport
  ProtocolFactory thrift.ProtocolFactory
  InputProtocol thrift.Protocol
  OutputProtocol thrift.Protocol
  SeqId int32
  Mu sync.Mutex
}

func NewGraphServiceThreadsafeClientFactory(t thrift.Transport, f thrift.ProtocolFactory) *GraphServiceThreadsafeClient {
  return &GraphServiceThreadsafeClient{Transport: t,
    ProtocolFactory: f,
    InputProtocol: f.GetProtocol(t),
    OutputProtocol: f.GetProtocol(t),
    SeqId: 0,
  }
}

func NewGraphServiceThreadsafeClient(t thrift.Transport, iprot thrift.Protocol, oprot thrift.Protocol) *GraphServiceThreadsafeClient {
  return &GraphServiceThreadsafeClient{Transport: t,
    ProtocolFactory: nil,
    InputProtocol: iprot,
    OutputProtocol: oprot,
    SeqId: 0,
  }
}

func (p *GraphServiceThreadsafeClient) Threadsafe() {}

// Parameters:
//  - Username
//  - Password
func (p *GraphServiceThreadsafeClient) Authenticate(username string, password string) (r *AuthResponse, err error) {
  p.Mu.Lock()
  defer p.Mu.Unlock()
  if err = p.sendAuthenticate(username, password); err != nil { return }
  return p.recvAuthenticate()
}

func (p *GraphServiceThreadsafeClient) sendAuthenticate(username string, password string)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("authenticate", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := GraphServiceAuthenticateArgs{
  Username : username,
  Password : password,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *GraphServiceThreadsafeClient) recvAuthenticate() (value *AuthResponse, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "authenticate" {
    err = thrift.NewApplicationException(thrift.WRONG_METHOD_NAME, "authenticate failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewApplicationException(thrift.BAD_SEQUENCE_ID, "authenticate failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error9 := thrift.NewApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error10 error
    error10, err = error9.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error10
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "authenticate failed: invalid message type")
    return
  }
  result := GraphServiceAuthenticateResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}

// Parameters:
//  - SessionId
func (p *GraphServiceThreadsafeClient) Signout(sessionId int64) (err error) {
  p.Mu.Lock()
  defer p.Mu.Unlock()
  if err = p.sendSignout(sessionId); err != nil { return }
  return
}

func (p *GraphServiceThreadsafeClient) sendSignout(sessionId int64)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("signout", thrift.ONEWAY, p.SeqId); err != nil {
      return
  }
  args := GraphServiceSignoutArgs{
  SessionId : sessionId,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}

// Parameters:
//  - SessionId
//  - Stmt
func (p *GraphServiceThreadsafeClient) Execute(sessionId int64, stmt string) (r *ExecutionResponse, err error) {
  p.Mu.Lock()
  defer p.Mu.Unlock()
  if err = p.sendExecute(sessionId, stmt); err != nil { return }
  return p.recvExecute()
}

func (p *GraphServiceThreadsafeClient) sendExecute(sessionId int64, stmt string)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("execute", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := GraphServiceExecuteArgs{
  SessionId : sessionId,
  Stmt : stmt,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *GraphServiceThreadsafeClient) recvExecute() (value *ExecutionResponse, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "execute" {
    err = thrift.NewApplicationException(thrift.WRONG_METHOD_NAME, "execute failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewApplicationException(thrift.BAD_SEQUENCE_ID, "execute failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error11 := thrift.NewApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error12 error
    error12, err = error11.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error12
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "execute failed: invalid message type")
    return
  }
  result := GraphServiceExecuteResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}


type GraphServiceProcessor struct {
  processorMap map[string]thrift.ProcessorFunction
  handler GraphService
}

func (p *GraphServiceProcessor) AddToProcessorMap(key string, processor thrift.ProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *GraphServiceProcessor) GetProcessorFunction(key string) (processor thrift.ProcessorFunction, err error) {
  if processor, ok := p.processorMap[key]; ok {
    return processor, nil
  }
  return nil, nil // generic error message will be sent
}

func (p *GraphServiceProcessor) ProcessorMap() map[string]thrift.ProcessorFunction {
  return p.processorMap
}

func NewGraphServiceProcessor(handler GraphService) *GraphServiceProcessor {
  self13 := &GraphServiceProcessor{handler:handler, processorMap:make(map[string]thrift.ProcessorFunction)}
  self13.processorMap["authenticate"] = &graphServiceProcessorAuthenticate{handler:handler}
  self13.processorMap["signout"] = &graphServiceProcessorSignout{handler:handler}
  self13.processorMap["execute"] = &graphServiceProcessorExecute{handler:handler}
  return self13
}

type graphServiceProcessorAuthenticate struct {
  handler GraphService
}

func (p *graphServiceProcessorAuthenticate) Read(iprot thrift.Protocol) (thrift.Struct, thrift.Exception) {
  args := GraphServiceAuthenticateArgs{}
  if err := args.Read(iprot); err != nil {
    return nil, err
  }
  iprot.ReadMessageEnd()
  return &args, nil
}

func (p *graphServiceProcessorAuthenticate) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Protocol) (err thrift.Exception) {
  var err2 error
  messageType := thrift.REPLY
  switch result.(type) {
  case thrift.ApplicationException:
    messageType = thrift.EXCEPTION
  }
  if err2 = oprot.WriteMessageBegin("authenticate", messageType, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  return err
}

func (p *graphServiceProcessorAuthenticate) Run(argStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
  args := argStruct.(*GraphServiceAuthenticateArgs)
  var result GraphServiceAuthenticateResult
  if retval, err := p.handler.Authenticate(args.Username, args.Password); err != nil {
    switch err.(type) {
    default:
      x := thrift.NewApplicationException(thrift.INTERNAL_ERROR, "Internal error processing authenticate: " + err.Error())
      return x, x
    }
  } else {
    result.Success = retval
  }
  return &result, nil
}

type graphServiceProcessorSignout struct {
  handler GraphService
}

func (p *graphServiceProcessorSignout) Read(iprot thrift.Protocol) (thrift.Struct, thrift.Exception) {
  args := GraphServiceSignoutArgs{}
  if err := args.Read(iprot); err != nil {
    return nil, err
  }
  iprot.ReadMessageEnd()
  return &args, nil
}

func (p *graphServiceProcessorSignout) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Protocol) (err thrift.Exception) {
  var err2 error
  messageType := thrift.REPLY
  switch result.(type) {
  case thrift.ApplicationException:
    messageType = thrift.EXCEPTION
  }
  if err2 = oprot.WriteMessageBegin("signout", messageType, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  return err
}

func (p *graphServiceProcessorSignout) Run(argStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
  args := argStruct.(*GraphServiceSignoutArgs)
  if err := p.handler.Signout(args.SessionId); err != nil {
    switch err.(type) {
    default:
      x := thrift.NewApplicationException(thrift.INTERNAL_ERROR, "Internal error processing signout: " + err.Error())
      return x, x
    }
  }
  return nil, nil
}

type graphServiceProcessorExecute struct {
  handler GraphService
}

func (p *graphServiceProcessorExecute) Read(iprot thrift.Protocol) (thrift.Struct, thrift.Exception) {
  args := GraphServiceExecuteArgs{}
  if err := args.Read(iprot); err != nil {
    return nil, err
  }
  iprot.ReadMessageEnd()
  return &args, nil
}

func (p *graphServiceProcessorExecute) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Protocol) (err thrift.Exception) {
  var err2 error
  messageType := thrift.REPLY
  switch result.(type) {
  case thrift.ApplicationException:
    messageType = thrift.EXCEPTION
  }
  if err2 = oprot.WriteMessageBegin("execute", messageType, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  return err
}

func (p *graphServiceProcessorExecute) Run(argStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
  args := argStruct.(*GraphServiceExecuteArgs)
  var result GraphServiceExecuteResult
  if retval, err := p.handler.Execute(args.SessionId, args.Stmt); err != nil {
    switch err.(type) {
    default:
      x := thrift.NewApplicationException(thrift.INTERNAL_ERROR, "Internal error processing execute: " + err.Error())
      return x, x
    }
  } else {
    result.Success = retval
  }
  return &result, nil
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Username
//  - Password
type GraphServiceAuthenticateArgs struct {
  Username string `thrift:"username,1" db:"username" json:"username"`
  Password string `thrift:"password,2" db:"password" json:"password"`
}

func NewGraphServiceAuthenticateArgs() *GraphServiceAuthenticateArgs {
  return &GraphServiceAuthenticateArgs{}
}


func (p *GraphServiceAuthenticateArgs) GetUsername() string {
  return p.Username
}

func (p *GraphServiceAuthenticateArgs) GetPassword() string {
  return p.Password
}
func (p *GraphServiceAuthenticateArgs) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *GraphServiceAuthenticateArgs)  ReadField1(iprot thrift.Protocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Username = v
}
  return nil
}

func (p *GraphServiceAuthenticateArgs)  ReadField2(iprot thrift.Protocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Password = v
}
  return nil
}

func (p *GraphServiceAuthenticateArgs) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("authenticate_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := p.writeField2(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *GraphServiceAuthenticateArgs) writeField1(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("username", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:username: ", p), err) }
  if err := oprot.WriteString(string(p.Username)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.username (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:username: ", p), err) }
  return err
}

func (p *GraphServiceAuthenticateArgs) writeField2(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("password", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:password: ", p), err) }
  if err := oprot.WriteString(string(p.Password)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.password (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:password: ", p), err) }
  return err
}

func (p *GraphServiceAuthenticateArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("GraphServiceAuthenticateArgs(%+v)", *p)
}

// Attributes:
//  - Success
type GraphServiceAuthenticateResult struct {
  Success *AuthResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewGraphServiceAuthenticateResult() *GraphServiceAuthenticateResult {
  return &GraphServiceAuthenticateResult{}
}

var GraphServiceAuthenticateResult_Success_DEFAULT *AuthResponse
func (p *GraphServiceAuthenticateResult) GetSuccess() *AuthResponse {
  if !p.IsSetSuccess() {
    return GraphServiceAuthenticateResult_Success_DEFAULT
  }
return p.Success
}
func (p *GraphServiceAuthenticateResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *GraphServiceAuthenticateResult) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if err := p.ReadField0(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *GraphServiceAuthenticateResult)  ReadField0(iprot thrift.Protocol) error {
  p.Success = NewAuthResponse()
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *GraphServiceAuthenticateResult) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("authenticate_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField0(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *GraphServiceAuthenticateResult) writeField0(oprot thrift.Protocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *GraphServiceAuthenticateResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("GraphServiceAuthenticateResult(%+v)", *p)
}

// Attributes:
//  - SessionId
type GraphServiceSignoutArgs struct {
  SessionId int64 `thrift:"sessionId,1" db:"sessionId" json:"sessionId"`
}

func NewGraphServiceSignoutArgs() *GraphServiceSignoutArgs {
  return &GraphServiceSignoutArgs{}
}


func (p *GraphServiceSignoutArgs) GetSessionId() int64 {
  return p.SessionId
}
func (p *GraphServiceSignoutArgs) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *GraphServiceSignoutArgs)  ReadField1(iprot thrift.Protocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.SessionId = v
}
  return nil
}

func (p *GraphServiceSignoutArgs) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("signout_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *GraphServiceSignoutArgs) writeField1(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("sessionId", thrift.I64, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:sessionId: ", p), err) }
  if err := oprot.WriteI64(int64(p.SessionId)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.sessionId (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:sessionId: ", p), err) }
  return err
}

func (p *GraphServiceSignoutArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("GraphServiceSignoutArgs(%+v)", *p)
}

// Attributes:
//  - SessionId
//  - Stmt
type GraphServiceExecuteArgs struct {
  SessionId int64 `thrift:"sessionId,1" db:"sessionId" json:"sessionId"`
  Stmt string `thrift:"stmt,2" db:"stmt" json:"stmt"`
}

func NewGraphServiceExecuteArgs() *GraphServiceExecuteArgs {
  return &GraphServiceExecuteArgs{}
}


func (p *GraphServiceExecuteArgs) GetSessionId() int64 {
  return p.SessionId
}

func (p *GraphServiceExecuteArgs) GetStmt() string {
  return p.Stmt
}
func (p *GraphServiceExecuteArgs) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *GraphServiceExecuteArgs)  ReadField1(iprot thrift.Protocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.SessionId = v
}
  return nil
}

func (p *GraphServiceExecuteArgs)  ReadField2(iprot thrift.Protocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Stmt = v
}
  return nil
}

func (p *GraphServiceExecuteArgs) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("execute_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := p.writeField2(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *GraphServiceExecuteArgs) writeField1(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("sessionId", thrift.I64, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:sessionId: ", p), err) }
  if err := oprot.WriteI64(int64(p.SessionId)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.sessionId (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:sessionId: ", p), err) }
  return err
}

func (p *GraphServiceExecuteArgs) writeField2(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("stmt", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:stmt: ", p), err) }
  if err := oprot.WriteString(string(p.Stmt)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.stmt (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:stmt: ", p), err) }
  return err
}

func (p *GraphServiceExecuteArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("GraphServiceExecuteArgs(%+v)", *p)
}

// Attributes:
//  - Success
type GraphServiceExecuteResult struct {
  Success *ExecutionResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewGraphServiceExecuteResult() *GraphServiceExecuteResult {
  return &GraphServiceExecuteResult{}
}

var GraphServiceExecuteResult_Success_DEFAULT *ExecutionResponse
func (p *GraphServiceExecuteResult) GetSuccess() *ExecutionResponse {
  if !p.IsSetSuccess() {
    return GraphServiceExecuteResult_Success_DEFAULT
  }
return p.Success
}
func (p *GraphServiceExecuteResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *GraphServiceExecuteResult) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if err := p.ReadField0(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *GraphServiceExecuteResult)  ReadField0(iprot thrift.Protocol) error {
  p.Success = NewExecutionResponse()
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *GraphServiceExecuteResult) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("execute_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField0(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *GraphServiceExecuteResult) writeField0(oprot thrift.Protocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *GraphServiceExecuteResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("GraphServiceExecuteResult(%+v)", *p)
}


