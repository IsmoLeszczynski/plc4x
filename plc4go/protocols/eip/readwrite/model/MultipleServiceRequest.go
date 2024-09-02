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

// Constant values.
const MultipleServiceRequest_REQUESTPATHSIZE uint8 = 0x02
const MultipleServiceRequest_REQUESTPATH uint32 = 0x01240220

// MultipleServiceRequest is the corresponding interface of MultipleServiceRequest
type MultipleServiceRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
	// GetData returns Data (property field)
	GetData() Services
}

// MultipleServiceRequestExactly can be used when we want exactly this type and not a type which fulfills MultipleServiceRequest.
// This is useful for switch cases.
type MultipleServiceRequestExactly interface {
	MultipleServiceRequest
	isMultipleServiceRequest() bool
}

// _MultipleServiceRequest is the data-structure of this message
type _MultipleServiceRequest struct {
	CipServiceContract
	Data Services
}

var _ MultipleServiceRequest = (*_MultipleServiceRequest)(nil)
var _ CipServiceRequirements = (*_MultipleServiceRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MultipleServiceRequest) GetService() uint8 {
	return 0x0A
}

func (m *_MultipleServiceRequest) GetResponse() bool {
	return bool(false)
}

func (m *_MultipleServiceRequest) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MultipleServiceRequest) GetParent() CipServiceContract {
	return m.CipServiceContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MultipleServiceRequest) GetData() Services {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_MultipleServiceRequest) GetRequestPathSize() uint8 {
	return MultipleServiceRequest_REQUESTPATHSIZE
}

func (m *_MultipleServiceRequest) GetRequestPath() uint32 {
	return MultipleServiceRequest_REQUESTPATH
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMultipleServiceRequest factory function for _MultipleServiceRequest
func NewMultipleServiceRequest(data Services, serviceLen uint16) *_MultipleServiceRequest {
	_result := &_MultipleServiceRequest{
		CipServiceContract: NewCipService(serviceLen),
		Data:               data,
	}
	_result.CipServiceContract.(*_CipService)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMultipleServiceRequest(structType any) MultipleServiceRequest {
	if casted, ok := structType.(MultipleServiceRequest); ok {
		return casted
	}
	if casted, ok := structType.(*MultipleServiceRequest); ok {
		return *casted
	}
	return nil
}

func (m *_MultipleServiceRequest) GetTypeName() string {
	return "MultipleServiceRequest"
}

func (m *_MultipleServiceRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CipServiceContract.(*_CipService).getLengthInBits(ctx))

	// Const Field (requestPathSize)
	lengthInBits += 8

	// Const Field (requestPath)
	lengthInBits += 32

	// Simple field (data)
	lengthInBits += m.Data.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_MultipleServiceRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MultipleServiceRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CipService, connected bool, serviceLen uint16) (__multipleServiceRequest MultipleServiceRequest, err error) {
	m.CipServiceContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MultipleServiceRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MultipleServiceRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestPathSize, err := ReadConstField[uint8](ctx, "requestPathSize", ReadUnsignedByte(readBuffer, uint8(8)), MultipleServiceRequest_REQUESTPATHSIZE)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestPathSize' field"))
	}
	_ = requestPathSize

	requestPath, err := ReadConstField[uint32](ctx, "requestPath", ReadUnsignedInt(readBuffer, uint8(32)), MultipleServiceRequest_REQUESTPATH)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestPath' field"))
	}
	_ = requestPath

	data, err := ReadSimpleField[Services](ctx, "data", ReadComplex[Services](ServicesParseWithBufferProducer((uint16)(uint16(serviceLen)-uint16(uint16(6)))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("MultipleServiceRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MultipleServiceRequest")
	}

	return m, nil
}

func (m *_MultipleServiceRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MultipleServiceRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MultipleServiceRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MultipleServiceRequest")
		}

		if err := WriteConstField(ctx, "requestPathSize", MultipleServiceRequest_REQUESTPATHSIZE, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestPathSize' field")
		}

		if err := WriteConstField(ctx, "requestPath", MultipleServiceRequest_REQUESTPATH, WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestPath' field")
		}

		if err := WriteSimpleField[Services](ctx, "data", m.GetData(), WriteComplex[Services](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("MultipleServiceRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MultipleServiceRequest")
		}
		return nil
	}
	return m.CipServiceContract.(*_CipService).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MultipleServiceRequest) isMultipleServiceRequest() bool {
	return true
}

func (m *_MultipleServiceRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
