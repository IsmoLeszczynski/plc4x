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

// ModbusPDUReadExceptionStatusRequest is the corresponding interface of ModbusPDUReadExceptionStatusRequest
type ModbusPDUReadExceptionStatusRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
}

// ModbusPDUReadExceptionStatusRequestExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUReadExceptionStatusRequest.
// This is useful for switch cases.
type ModbusPDUReadExceptionStatusRequestExactly interface {
	ModbusPDUReadExceptionStatusRequest
	isModbusPDUReadExceptionStatusRequest() bool
}

// _ModbusPDUReadExceptionStatusRequest is the data-structure of this message
type _ModbusPDUReadExceptionStatusRequest struct {
	*_ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadExceptionStatusRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadExceptionStatusRequest) GetFunctionFlag() uint8 {
	return 0x07
}

func (m *_ModbusPDUReadExceptionStatusRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadExceptionStatusRequest) InitializeParent(parent ModbusPDU) {}

func (m *_ModbusPDUReadExceptionStatusRequest) GetParent() ModbusPDU {
	return m._ModbusPDU
}

// NewModbusPDUReadExceptionStatusRequest factory function for _ModbusPDUReadExceptionStatusRequest
func NewModbusPDUReadExceptionStatusRequest() *_ModbusPDUReadExceptionStatusRequest {
	_result := &_ModbusPDUReadExceptionStatusRequest{
		_ModbusPDU: NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReadExceptionStatusRequest(structType any) ModbusPDUReadExceptionStatusRequest {
	if casted, ok := structType.(ModbusPDUReadExceptionStatusRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadExceptionStatusRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadExceptionStatusRequest) GetTypeName() string {
	return "ModbusPDUReadExceptionStatusRequest"
}

func (m *_ModbusPDUReadExceptionStatusRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_ModbusPDUReadExceptionStatusRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUReadExceptionStatusRequestParse(ctx context.Context, theBytes []byte, response bool) (ModbusPDUReadExceptionStatusRequest, error) {
	return ModbusPDUReadExceptionStatusRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func ModbusPDUReadExceptionStatusRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (ModbusPDUReadExceptionStatusRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ModbusPDUReadExceptionStatusRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadExceptionStatusRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ModbusPDUReadExceptionStatusRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadExceptionStatusRequest")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUReadExceptionStatusRequest{
		_ModbusPDU: &_ModbusPDU{},
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUReadExceptionStatusRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadExceptionStatusRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadExceptionStatusRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadExceptionStatusRequest")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadExceptionStatusRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadExceptionStatusRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReadExceptionStatusRequest) isModbusPDUReadExceptionStatusRequest() bool {
	return true
}

func (m *_ModbusPDUReadExceptionStatusRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
