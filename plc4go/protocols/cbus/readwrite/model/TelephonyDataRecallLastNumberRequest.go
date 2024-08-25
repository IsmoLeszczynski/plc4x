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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// TelephonyDataRecallLastNumberRequest is the corresponding interface of TelephonyDataRecallLastNumberRequest
type TelephonyDataRecallLastNumberRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	TelephonyData
	// GetRecallLastNumberType returns RecallLastNumberType (property field)
	GetRecallLastNumberType() byte
	// GetIsNumberOfLastOutgoingCall returns IsNumberOfLastOutgoingCall (virtual field)
	GetIsNumberOfLastOutgoingCall() bool
	// GetIsNumberOfLastIncomingCall returns IsNumberOfLastIncomingCall (virtual field)
	GetIsNumberOfLastIncomingCall() bool
}

// TelephonyDataRecallLastNumberRequestExactly can be used when we want exactly this type and not a type which fulfills TelephonyDataRecallLastNumberRequest.
// This is useful for switch cases.
type TelephonyDataRecallLastNumberRequestExactly interface {
	TelephonyDataRecallLastNumberRequest
	isTelephonyDataRecallLastNumberRequest() bool
}

// _TelephonyDataRecallLastNumberRequest is the data-structure of this message
type _TelephonyDataRecallLastNumberRequest struct {
	*_TelephonyData
	RecallLastNumberType byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TelephonyDataRecallLastNumberRequest) InitializeParent(parent TelephonyData, commandTypeContainer TelephonyCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_TelephonyDataRecallLastNumberRequest) GetParent() TelephonyData {
	return m._TelephonyData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TelephonyDataRecallLastNumberRequest) GetRecallLastNumberType() byte {
	return m.RecallLastNumberType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_TelephonyDataRecallLastNumberRequest) GetIsNumberOfLastOutgoingCall() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetRecallLastNumberType()) == (0x01)))
}

func (m *_TelephonyDataRecallLastNumberRequest) GetIsNumberOfLastIncomingCall() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetRecallLastNumberType()) == (0x02)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTelephonyDataRecallLastNumberRequest factory function for _TelephonyDataRecallLastNumberRequest
func NewTelephonyDataRecallLastNumberRequest(recallLastNumberType byte, commandTypeContainer TelephonyCommandTypeContainer, argument byte) *_TelephonyDataRecallLastNumberRequest {
	_result := &_TelephonyDataRecallLastNumberRequest{
		RecallLastNumberType: recallLastNumberType,
		_TelephonyData:       NewTelephonyData(commandTypeContainer, argument),
	}
	_result._TelephonyData._TelephonyDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTelephonyDataRecallLastNumberRequest(structType any) TelephonyDataRecallLastNumberRequest {
	if casted, ok := structType.(TelephonyDataRecallLastNumberRequest); ok {
		return casted
	}
	if casted, ok := structType.(*TelephonyDataRecallLastNumberRequest); ok {
		return *casted
	}
	return nil
}

func (m *_TelephonyDataRecallLastNumberRequest) GetTypeName() string {
	return "TelephonyDataRecallLastNumberRequest"
}

func (m *_TelephonyDataRecallLastNumberRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (recallLastNumberType)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_TelephonyDataRecallLastNumberRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TelephonyDataRecallLastNumberRequestParse(ctx context.Context, theBytes []byte) (TelephonyDataRecallLastNumberRequest, error) {
	return TelephonyDataRecallLastNumberRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TelephonyDataRecallLastNumberRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TelephonyDataRecallLastNumberRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("TelephonyDataRecallLastNumberRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TelephonyDataRecallLastNumberRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (recallLastNumberType)
	_recallLastNumberType, _recallLastNumberTypeErr := readBuffer.ReadByte("recallLastNumberType")
	if _recallLastNumberTypeErr != nil {
		return nil, errors.Wrap(_recallLastNumberTypeErr, "Error parsing 'recallLastNumberType' field of TelephonyDataRecallLastNumberRequest")
	}
	recallLastNumberType := _recallLastNumberType

	// Virtual field
	_isNumberOfLastOutgoingCall := bool((recallLastNumberType) == (0x01))
	isNumberOfLastOutgoingCall := bool(_isNumberOfLastOutgoingCall)
	_ = isNumberOfLastOutgoingCall

	// Virtual field
	_isNumberOfLastIncomingCall := bool((recallLastNumberType) == (0x02))
	isNumberOfLastIncomingCall := bool(_isNumberOfLastIncomingCall)
	_ = isNumberOfLastIncomingCall

	if closeErr := readBuffer.CloseContext("TelephonyDataRecallLastNumberRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TelephonyDataRecallLastNumberRequest")
	}

	// Create a partially initialized instance
	_child := &_TelephonyDataRecallLastNumberRequest{
		_TelephonyData:       &_TelephonyData{},
		RecallLastNumberType: recallLastNumberType,
	}
	_child._TelephonyData._TelephonyDataChildRequirements = _child
	return _child, nil
}

func (m *_TelephonyDataRecallLastNumberRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TelephonyDataRecallLastNumberRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TelephonyDataRecallLastNumberRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TelephonyDataRecallLastNumberRequest")
		}

		// Simple Field (recallLastNumberType)
		recallLastNumberType := byte(m.GetRecallLastNumberType())
		_recallLastNumberTypeErr := writeBuffer.WriteByte("recallLastNumberType", (recallLastNumberType))
		if _recallLastNumberTypeErr != nil {
			return errors.Wrap(_recallLastNumberTypeErr, "Error serializing 'recallLastNumberType' field")
		}
		// Virtual field
		isNumberOfLastOutgoingCall := m.GetIsNumberOfLastOutgoingCall()
		_ = isNumberOfLastOutgoingCall
		if _isNumberOfLastOutgoingCallErr := writeBuffer.WriteVirtual(ctx, "isNumberOfLastOutgoingCall", m.GetIsNumberOfLastOutgoingCall()); _isNumberOfLastOutgoingCallErr != nil {
			return errors.Wrap(_isNumberOfLastOutgoingCallErr, "Error serializing 'isNumberOfLastOutgoingCall' field")
		}
		// Virtual field
		isNumberOfLastIncomingCall := m.GetIsNumberOfLastIncomingCall()
		_ = isNumberOfLastIncomingCall
		if _isNumberOfLastIncomingCallErr := writeBuffer.WriteVirtual(ctx, "isNumberOfLastIncomingCall", m.GetIsNumberOfLastIncomingCall()); _isNumberOfLastIncomingCallErr != nil {
			return errors.Wrap(_isNumberOfLastIncomingCallErr, "Error serializing 'isNumberOfLastIncomingCall' field")
		}

		if popErr := writeBuffer.PopContext("TelephonyDataRecallLastNumberRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TelephonyDataRecallLastNumberRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TelephonyDataRecallLastNumberRequest) isTelephonyDataRecallLastNumberRequest() bool {
	return true
}

func (m *_TelephonyDataRecallLastNumberRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
