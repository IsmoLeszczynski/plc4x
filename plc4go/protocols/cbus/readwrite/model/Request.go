/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// Request is the corresponding interface of Request
type Request interface {
	RequestContract
	RequestRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsRequest()
	// CreateBuilder creates a RequestBuilder
	CreateRequestBuilder() RequestBuilder
}

// RequestContract provides a set of functions which can be overwritten by a sub struct
type RequestContract interface {
	// GetPeekedByte returns PeekedByte (property field)
	GetPeekedByte() RequestType
	// GetStartingCR returns StartingCR (property field)
	GetStartingCR() *RequestType
	// GetResetMode returns ResetMode (property field)
	GetResetMode() *RequestType
	// GetSecondPeek returns SecondPeek (property field)
	GetSecondPeek() RequestType
	// GetTermination returns Termination (property field)
	GetTermination() RequestTermination
	// GetActualPeek returns ActualPeek (virtual field)
	GetActualPeek() RequestType
	// GetCBusOptions() returns a parser argument
	GetCBusOptions() CBusOptions
	// IsRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsRequest()
	// CreateBuilder creates a RequestBuilder
	CreateRequestBuilder() RequestBuilder
}

// RequestRequirements provides a set of functions which need to be implemented by a sub struct
type RequestRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetActualPeek returns ActualPeek (discriminator field)
	GetActualPeek() RequestType
}

// _Request is the data-structure of this message
type _Request struct {
	_SubType interface {
		RequestContract
		RequestRequirements
	}
	PeekedByte  RequestType
	StartingCR  *RequestType
	ResetMode   *RequestType
	SecondPeek  RequestType
	Termination RequestTermination

	// Arguments.
	CBusOptions CBusOptions
}

var _ RequestContract = (*_Request)(nil)

// NewRequest factory function for _Request
func NewRequest(peekedByte RequestType, startingCR *RequestType, resetMode *RequestType, secondPeek RequestType, termination RequestTermination, cBusOptions CBusOptions) *_Request {
	if termination == nil {
		panic("termination of type RequestTermination for Request must not be nil")
	}
	return &_Request{PeekedByte: peekedByte, StartingCR: startingCR, ResetMode: resetMode, SecondPeek: secondPeek, Termination: termination, CBusOptions: cBusOptions}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// RequestBuilder is a builder for Request
type RequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(peekedByte RequestType, secondPeek RequestType, termination RequestTermination) RequestBuilder
	// WithPeekedByte adds PeekedByte (property field)
	WithPeekedByte(RequestType) RequestBuilder
	// WithStartingCR adds StartingCR (property field)
	WithOptionalStartingCR(RequestType) RequestBuilder
	// WithResetMode adds ResetMode (property field)
	WithOptionalResetMode(RequestType) RequestBuilder
	// WithSecondPeek adds SecondPeek (property field)
	WithSecondPeek(RequestType) RequestBuilder
	// WithTermination adds Termination (property field)
	WithTermination(RequestTermination) RequestBuilder
	// WithTerminationBuilder adds Termination (property field) which is build by the builder
	WithTerminationBuilder(func(RequestTerminationBuilder) RequestTerminationBuilder) RequestBuilder
	// WithArgCBusOptions sets a parser argument
	WithArgCBusOptions(CBusOptions) RequestBuilder
	// AsRequestSmartConnectShortcut converts this build to a subType of Request. It is always possible to return to current builder using Done()
	AsRequestSmartConnectShortcut() RequestSmartConnectShortcutBuilder
	// AsRequestReset converts this build to a subType of Request. It is always possible to return to current builder using Done()
	AsRequestReset() RequestResetBuilder
	// AsRequestDirectCommandAccess converts this build to a subType of Request. It is always possible to return to current builder using Done()
	AsRequestDirectCommandAccess() RequestDirectCommandAccessBuilder
	// AsRequestCommand converts this build to a subType of Request. It is always possible to return to current builder using Done()
	AsRequestCommand() RequestCommandBuilder
	// AsRequestNull converts this build to a subType of Request. It is always possible to return to current builder using Done()
	AsRequestNull() RequestNullBuilder
	// AsRequestEmpty converts this build to a subType of Request. It is always possible to return to current builder using Done()
	AsRequestEmpty() RequestEmptyBuilder
	// AsRequestObsolete converts this build to a subType of Request. It is always possible to return to current builder using Done()
	AsRequestObsolete() RequestObsoleteBuilder
	// Build builds the Request or returns an error if something is wrong
	PartialBuild() (RequestContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() RequestContract
	// Build builds the Request or returns an error if something is wrong
	Build() (Request, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() Request
}

// NewRequestBuilder() creates a RequestBuilder
func NewRequestBuilder() RequestBuilder {
	return &_RequestBuilder{_Request: new(_Request)}
}

type _RequestChildBuilder interface {
	utils.Copyable
	setParent(RequestContract)
	buildForRequest() (Request, error)
}

type _RequestBuilder struct {
	*_Request

	childBuilder _RequestChildBuilder

	err *utils.MultiError
}

var _ (RequestBuilder) = (*_RequestBuilder)(nil)

func (b *_RequestBuilder) WithMandatoryFields(peekedByte RequestType, secondPeek RequestType, termination RequestTermination) RequestBuilder {
	return b.WithPeekedByte(peekedByte).WithSecondPeek(secondPeek).WithTermination(termination)
}

func (b *_RequestBuilder) WithPeekedByte(peekedByte RequestType) RequestBuilder {
	b.PeekedByte = peekedByte
	return b
}

func (b *_RequestBuilder) WithOptionalStartingCR(startingCR RequestType) RequestBuilder {
	b.StartingCR = &startingCR
	return b
}

func (b *_RequestBuilder) WithOptionalResetMode(resetMode RequestType) RequestBuilder {
	b.ResetMode = &resetMode
	return b
}

func (b *_RequestBuilder) WithSecondPeek(secondPeek RequestType) RequestBuilder {
	b.SecondPeek = secondPeek
	return b
}

func (b *_RequestBuilder) WithTermination(termination RequestTermination) RequestBuilder {
	b.Termination = termination
	return b
}

func (b *_RequestBuilder) WithTerminationBuilder(builderSupplier func(RequestTerminationBuilder) RequestTerminationBuilder) RequestBuilder {
	builder := builderSupplier(b.Termination.CreateRequestTerminationBuilder())
	var err error
	b.Termination, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "RequestTerminationBuilder failed"))
	}
	return b
}

func (b *_RequestBuilder) WithArgCBusOptions(cBusOptions CBusOptions) RequestBuilder {
	b.CBusOptions = cBusOptions
	return b
}

func (b *_RequestBuilder) PartialBuild() (RequestContract, error) {
	if b.Termination == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'termination' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._Request.deepCopy(), nil
}

func (b *_RequestBuilder) PartialMustBuild() RequestContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_RequestBuilder) AsRequestSmartConnectShortcut() RequestSmartConnectShortcutBuilder {
	if cb, ok := b.childBuilder.(RequestSmartConnectShortcutBuilder); ok {
		return cb
	}
	cb := NewRequestSmartConnectShortcutBuilder().(*_RequestSmartConnectShortcutBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_RequestBuilder) AsRequestReset() RequestResetBuilder {
	if cb, ok := b.childBuilder.(RequestResetBuilder); ok {
		return cb
	}
	cb := NewRequestResetBuilder().(*_RequestResetBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_RequestBuilder) AsRequestDirectCommandAccess() RequestDirectCommandAccessBuilder {
	if cb, ok := b.childBuilder.(RequestDirectCommandAccessBuilder); ok {
		return cb
	}
	cb := NewRequestDirectCommandAccessBuilder().(*_RequestDirectCommandAccessBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_RequestBuilder) AsRequestCommand() RequestCommandBuilder {
	if cb, ok := b.childBuilder.(RequestCommandBuilder); ok {
		return cb
	}
	cb := NewRequestCommandBuilder().(*_RequestCommandBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_RequestBuilder) AsRequestNull() RequestNullBuilder {
	if cb, ok := b.childBuilder.(RequestNullBuilder); ok {
		return cb
	}
	cb := NewRequestNullBuilder().(*_RequestNullBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_RequestBuilder) AsRequestEmpty() RequestEmptyBuilder {
	if cb, ok := b.childBuilder.(RequestEmptyBuilder); ok {
		return cb
	}
	cb := NewRequestEmptyBuilder().(*_RequestEmptyBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_RequestBuilder) AsRequestObsolete() RequestObsoleteBuilder {
	if cb, ok := b.childBuilder.(RequestObsoleteBuilder); ok {
		return cb
	}
	cb := NewRequestObsoleteBuilder().(*_RequestObsoleteBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_RequestBuilder) Build() (Request, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForRequest()
}

func (b *_RequestBuilder) MustBuild() Request {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_RequestBuilder) DeepCopy() any {
	_copy := b.CreateRequestBuilder().(*_RequestBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_RequestChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateRequestBuilder creates a RequestBuilder
func (b *_Request) CreateRequestBuilder() RequestBuilder {
	if b == nil {
		return NewRequestBuilder()
	}
	return &_RequestBuilder{_Request: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Request) GetPeekedByte() RequestType {
	return m.PeekedByte
}

func (m *_Request) GetStartingCR() *RequestType {
	return m.StartingCR
}

func (m *_Request) GetResetMode() *RequestType {
	return m.ResetMode
}

func (m *_Request) GetSecondPeek() RequestType {
	return m.SecondPeek
}

func (m *_Request) GetTermination() RequestTermination {
	return m.Termination
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_Request) GetActualPeek() RequestType {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	startingCR := m.GetStartingCR()
	_ = startingCR
	resetMode := m.GetResetMode()
	_ = resetMode
	return CastRequestType(CastRequestType(utils.InlineIf(bool((bool(bool((m.GetStartingCR()) == (nil))) && bool(bool((m.GetResetMode()) == (nil))))) || bool((bool(bool(bool((m.GetStartingCR()) == (nil))) && bool(bool((m.GetResetMode()) != (nil)))) && bool(bool((m.GetSecondPeek()) == (RequestType_EMPTY))))), func() any { return CastRequestType(m.GetPeekedByte()) }, func() any { return CastRequestType(m.GetSecondPeek()) })))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastRequest(structType any) Request {
	if casted, ok := structType.(Request); ok {
		return casted
	}
	if casted, ok := structType.(*Request); ok {
		return *casted
	}
	return nil
}

func (m *_Request) GetTypeName() string {
	return "Request"
}

func (m *_Request) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Optional Field (startingCR)
	if m.StartingCR != nil {
		lengthInBits += 8
	}

	// Optional Field (resetMode)
	if m.ResetMode != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// Simple field (termination)
	lengthInBits += m.Termination.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_Request) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_Request) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func RequestParse[T Request](ctx context.Context, theBytes []byte, cBusOptions CBusOptions) (T, error) {
	return RequestParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), cBusOptions)
}

func RequestParseWithBufferProducer[T Request](cBusOptions CBusOptions) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := RequestParseWithBuffer[T](ctx, readBuffer, cBusOptions)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func RequestParseWithBuffer[T Request](ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (T, error) {
	v, err := (&_Request{CBusOptions: cBusOptions}).parse(ctx, readBuffer, cBusOptions)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_Request) parse(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (__request Request, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Request"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Request")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	peekedByte, err := ReadPeekField[RequestType](ctx, "peekedByte", ReadEnum(RequestTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedByte' field"))
	}
	m.PeekedByte = peekedByte

	var startingCR *RequestType
	startingCR, err = ReadOptionalField[RequestType](ctx, "startingCR", ReadEnum(RequestTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))), bool((peekedByte) == (RequestType_EMPTY)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startingCR' field"))
	}
	m.StartingCR = startingCR

	var resetMode *RequestType
	resetMode, err = ReadOptionalField[RequestType](ctx, "resetMode", ReadEnum(RequestTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))), bool((peekedByte) == (RequestType_RESET)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'resetMode' field"))
	}
	m.ResetMode = resetMode

	secondPeek, err := ReadPeekField[RequestType](ctx, "secondPeek", ReadEnum(RequestTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'secondPeek' field"))
	}
	m.SecondPeek = secondPeek

	actualPeek, err := ReadVirtualField[RequestType](ctx, "actualPeek", (*RequestType)(nil), CastRequestType(utils.InlineIf(bool((bool(bool((startingCR) == (nil))) && bool(bool((resetMode) == (nil))))) || bool((bool(bool(bool((startingCR) == (nil))) && bool(bool((resetMode) != (nil)))) && bool(bool((secondPeek) == (RequestType_EMPTY))))), func() any { return CastRequestType(peekedByte) }, func() any { return CastRequestType(secondPeek) })))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualPeek' field"))
	}
	_ = actualPeek

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child Request
	switch {
	case actualPeek == RequestType_SMART_CONNECT_SHORTCUT: // RequestSmartConnectShortcut
		if _child, err = new(_RequestSmartConnectShortcut).parse(ctx, readBuffer, m, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type RequestSmartConnectShortcut for type-switch of Request")
		}
	case actualPeek == RequestType_RESET: // RequestReset
		if _child, err = new(_RequestReset).parse(ctx, readBuffer, m, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type RequestReset for type-switch of Request")
		}
	case actualPeek == RequestType_DIRECT_COMMAND: // RequestDirectCommandAccess
		if _child, err = new(_RequestDirectCommandAccess).parse(ctx, readBuffer, m, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type RequestDirectCommandAccess for type-switch of Request")
		}
	case actualPeek == RequestType_REQUEST_COMMAND: // RequestCommand
		if _child, err = new(_RequestCommand).parse(ctx, readBuffer, m, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type RequestCommand for type-switch of Request")
		}
	case actualPeek == RequestType_NULL: // RequestNull
		if _child, err = new(_RequestNull).parse(ctx, readBuffer, m, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type RequestNull for type-switch of Request")
		}
	case actualPeek == RequestType_EMPTY: // RequestEmpty
		if _child, err = new(_RequestEmpty).parse(ctx, readBuffer, m, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type RequestEmpty for type-switch of Request")
		}
	case 0 == 0: // RequestObsolete
		if _child, err = new(_RequestObsolete).parse(ctx, readBuffer, m, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type RequestObsolete for type-switch of Request")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [actualPeek=%v]", actualPeek)
	}

	termination, err := ReadSimpleField[RequestTermination](ctx, "termination", ReadComplex[RequestTermination](RequestTerminationParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'termination' field"))
	}
	m.Termination = termination

	if closeErr := readBuffer.CloseContext("Request"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Request")
	}

	return _child, nil
}

func (pm *_Request) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child Request, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("Request"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Request")
	}

	if err := WriteOptionalEnumField[RequestType](ctx, "startingCR", "RequestType", m.GetStartingCR(), WriteEnum[RequestType, uint8](RequestType.GetValue, RequestType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8)), bool((m.GetPeekedByte()) == (RequestType_EMPTY))); err != nil {
		return errors.Wrap(err, "Error serializing 'startingCR' field")
	}

	if err := WriteOptionalEnumField[RequestType](ctx, "resetMode", "RequestType", m.GetResetMode(), WriteEnum[RequestType, uint8](RequestType.GetValue, RequestType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8)), bool((m.GetPeekedByte()) == (RequestType_RESET))); err != nil {
		return errors.Wrap(err, "Error serializing 'resetMode' field")
	}
	// Virtual field
	actualPeek := m.GetActualPeek()
	_ = actualPeek
	if _actualPeekErr := writeBuffer.WriteVirtual(ctx, "actualPeek", m.GetActualPeek()); _actualPeekErr != nil {
		return errors.Wrap(_actualPeekErr, "Error serializing 'actualPeek' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if err := WriteSimpleField[RequestTermination](ctx, "termination", m.GetTermination(), WriteComplex[RequestTermination](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'termination' field")
	}

	if popErr := writeBuffer.PopContext("Request"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Request")
	}
	return nil
}

////
// Arguments Getter

func (m *_Request) GetCBusOptions() CBusOptions {
	return m.CBusOptions
}

//
////

func (m *_Request) IsRequest() {}

func (m *_Request) DeepCopy() any {
	return m.deepCopy()
}

func (m *_Request) deepCopy() *_Request {
	if m == nil {
		return nil
	}
	_RequestCopy := &_Request{
		nil, // will be set by child
		m.PeekedByte,
		utils.CopyPtr[RequestType](m.StartingCR),
		utils.CopyPtr[RequestType](m.ResetMode),
		m.SecondPeek,
		utils.DeepCopy[RequestTermination](m.Termination),
		m.CBusOptions,
	}
	return _RequestCopy
}
