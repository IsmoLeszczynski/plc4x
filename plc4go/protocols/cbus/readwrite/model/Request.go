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
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// Request is the corresponding interface of Request
type Request interface {
	RequestContract
	RequestRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
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
	// IsRequest() is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsRequest()
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
	_SubType    Request
	PeekedByte  RequestType
	StartingCR  *RequestType
	ResetMode   *RequestType
	SecondPeek  RequestType
	Termination RequestTermination

	// Arguments.
	CBusOptions CBusOptions
}

var _ RequestContract = (*_Request)(nil)

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

// NewRequest factory function for _Request
func NewRequest(peekedByte RequestType, startingCR *RequestType, resetMode *RequestType, secondPeek RequestType, termination RequestTermination, cBusOptions CBusOptions) *_Request {
	return &_Request{PeekedByte: peekedByte, StartingCR: startingCR, ResetMode: resetMode, SecondPeek: secondPeek, Termination: termination, CBusOptions: cBusOptions}
}

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
		return v, err
	}
}

func RequestParseWithBuffer[T Request](ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (T, error) {
	v, err := (&_Request{CBusOptions: cBusOptions}).parse(ctx, readBuffer, cBusOptions)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
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
		if _child, err = (&_RequestSmartConnectShortcut{}).parse(ctx, readBuffer, m, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type RequestSmartConnectShortcut for type-switch of Request")
		}
	case actualPeek == RequestType_RESET: // RequestReset
		if _child, err = (&_RequestReset{}).parse(ctx, readBuffer, m, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type RequestReset for type-switch of Request")
		}
	case actualPeek == RequestType_DIRECT_COMMAND: // RequestDirectCommandAccess
		if _child, err = (&_RequestDirectCommandAccess{}).parse(ctx, readBuffer, m, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type RequestDirectCommandAccess for type-switch of Request")
		}
	case actualPeek == RequestType_REQUEST_COMMAND: // RequestCommand
		if _child, err = (&_RequestCommand{}).parse(ctx, readBuffer, m, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type RequestCommand for type-switch of Request")
		}
	case actualPeek == RequestType_NULL: // RequestNull
		if _child, err = (&_RequestNull{}).parse(ctx, readBuffer, m, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type RequestNull for type-switch of Request")
		}
	case actualPeek == RequestType_EMPTY: // RequestEmpty
		if _child, err = (&_RequestEmpty{}).parse(ctx, readBuffer, m, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type RequestEmpty for type-switch of Request")
		}
	case 0 == 0: // RequestObsolete
		if _child, err = (&_RequestObsolete{}).parse(ctx, readBuffer, m, cBusOptions); err != nil {
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
