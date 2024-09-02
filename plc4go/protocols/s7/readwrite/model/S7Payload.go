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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// S7Payload is the corresponding interface of S7Payload
type S7Payload interface {
	S7PayloadContract
	S7PayloadRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// S7PayloadContract provides a set of functions which can be overwritten by a sub struct
type S7PayloadContract interface {
	// GetParameter() returns a parser argument
	GetParameter() S7Parameter
	// IsS7Payload() is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7Payload()
}

// S7PayloadRequirements provides a set of functions which need to be implemented by a sub struct
type S7PayloadRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetMessageType returns MessageType (discriminator field)
	GetMessageType() uint8
	// GetParameterParameterType returns ParameterParameterType (discriminator field)
	GetParameterParameterType() uint8
}

// _S7Payload is the data-structure of this message
type _S7Payload struct {
	_SubType S7Payload

	// Arguments.
	Parameter S7Parameter
}

var _ S7PayloadContract = (*_S7Payload)(nil)

// NewS7Payload factory function for _S7Payload
func NewS7Payload(parameter S7Parameter) *_S7Payload {
	return &_S7Payload{Parameter: parameter}
}

// Deprecated: use the interface for direct cast
func CastS7Payload(structType any) S7Payload {
	if casted, ok := structType.(S7Payload); ok {
		return casted
	}
	if casted, ok := structType.(*S7Payload); ok {
		return *casted
	}
	return nil
}

func (m *_S7Payload) GetTypeName() string {
	return "S7Payload"
}

func (m *_S7Payload) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_S7Payload) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func S7PayloadParse[T S7Payload](ctx context.Context, theBytes []byte, messageType uint8, parameter S7Parameter) (T, error) {
	return S7PayloadParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), messageType, parameter)
}

func S7PayloadParseWithBufferProducer[T S7Payload](messageType uint8, parameter S7Parameter) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := S7PayloadParseWithBuffer[T](ctx, readBuffer, messageType, parameter)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func S7PayloadParseWithBuffer[T S7Payload](ctx context.Context, readBuffer utils.ReadBuffer, messageType uint8, parameter S7Parameter) (T, error) {
	v, err := (&_S7Payload{Parameter: parameter}).parse(ctx, readBuffer, messageType, parameter)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_S7Payload) parse(ctx context.Context, readBuffer utils.ReadBuffer, messageType uint8, parameter S7Parameter) (__s7Payload S7Payload, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7Payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7Payload")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child S7Payload
	switch {
	case CastS7Parameter(parameter).GetParameterType() == 0x04 && messageType == 0x03: // S7PayloadReadVarResponse
		if _child, err = (&_S7PayloadReadVarResponse{}).parse(ctx, readBuffer, m, messageType, parameter); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type S7PayloadReadVarResponse for type-switch of S7Payload")
		}
	case CastS7Parameter(parameter).GetParameterType() == 0x05 && messageType == 0x01: // S7PayloadWriteVarRequest
		if _child, err = (&_S7PayloadWriteVarRequest{}).parse(ctx, readBuffer, m, messageType, parameter); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type S7PayloadWriteVarRequest for type-switch of S7Payload")
		}
	case CastS7Parameter(parameter).GetParameterType() == 0x05 && messageType == 0x03: // S7PayloadWriteVarResponse
		if _child, err = (&_S7PayloadWriteVarResponse{}).parse(ctx, readBuffer, m, messageType, parameter); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type S7PayloadWriteVarResponse for type-switch of S7Payload")
		}
	case CastS7Parameter(parameter).GetParameterType() == 0x00 && messageType == 0x07: // S7PayloadUserData
		if _child, err = (&_S7PayloadUserData{}).parse(ctx, readBuffer, m, messageType, parameter); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type S7PayloadUserData for type-switch of S7Payload")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [parameterparameterType=%v, messageType=%v]", CastS7Parameter(parameter).GetParameterType(), messageType)
	}

	if closeErr := readBuffer.CloseContext("S7Payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7Payload")
	}

	return _child, nil
}

func (pm *_S7Payload) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child S7Payload, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("S7Payload"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for S7Payload")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("S7Payload"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for S7Payload")
	}
	return nil
}

////
// Arguments Getter

func (m *_S7Payload) GetParameter() S7Parameter {
	return m.Parameter
}

//
////

func (m *_S7Payload) IsS7Payload() {}
