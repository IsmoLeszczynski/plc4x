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

// Constant values.
const CipUnconnectedRequest_ROUTE uint16 = 0x0001

// CipUnconnectedRequest is the corresponding interface of CipUnconnectedRequest
type CipUnconnectedRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
	// GetClassSegment returns ClassSegment (property field)
	GetClassSegment() PathSegment
	// GetInstanceSegment returns InstanceSegment (property field)
	GetInstanceSegment() PathSegment
	// GetUnconnectedService returns UnconnectedService (property field)
	GetUnconnectedService() CipService
	// GetBackPlane returns BackPlane (property field)
	GetBackPlane() int8
	// GetSlot returns Slot (property field)
	GetSlot() int8
}

// _CipUnconnectedRequest is the data-structure of this message
type _CipUnconnectedRequest struct {
	CipServiceContract
	ClassSegment       PathSegment
	InstanceSegment    PathSegment
	UnconnectedService CipService
	BackPlane          int8
	Slot               int8
	// Reserved Fields
	reservedField0 *uint16
}

var _ CipUnconnectedRequest = (*_CipUnconnectedRequest)(nil)
var _ CipServiceRequirements = (*_CipUnconnectedRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipUnconnectedRequest) GetService() uint8 {
	return 0x52
}

func (m *_CipUnconnectedRequest) GetResponse() bool {
	return bool(false)
}

func (m *_CipUnconnectedRequest) GetConnected() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipUnconnectedRequest) GetParent() CipServiceContract {
	return m.CipServiceContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipUnconnectedRequest) GetClassSegment() PathSegment {
	return m.ClassSegment
}

func (m *_CipUnconnectedRequest) GetInstanceSegment() PathSegment {
	return m.InstanceSegment
}

func (m *_CipUnconnectedRequest) GetUnconnectedService() CipService {
	return m.UnconnectedService
}

func (m *_CipUnconnectedRequest) GetBackPlane() int8 {
	return m.BackPlane
}

func (m *_CipUnconnectedRequest) GetSlot() int8 {
	return m.Slot
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_CipUnconnectedRequest) GetRoute() uint16 {
	return CipUnconnectedRequest_ROUTE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCipUnconnectedRequest factory function for _CipUnconnectedRequest
func NewCipUnconnectedRequest(classSegment PathSegment, instanceSegment PathSegment, unconnectedService CipService, backPlane int8, slot int8, serviceLen uint16) *_CipUnconnectedRequest {
	_result := &_CipUnconnectedRequest{
		CipServiceContract: NewCipService(serviceLen),
		ClassSegment:       classSegment,
		InstanceSegment:    instanceSegment,
		UnconnectedService: unconnectedService,
		BackPlane:          backPlane,
		Slot:               slot,
	}
	_result.CipServiceContract.(*_CipService)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCipUnconnectedRequest(structType any) CipUnconnectedRequest {
	if casted, ok := structType.(CipUnconnectedRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CipUnconnectedRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CipUnconnectedRequest) GetTypeName() string {
	return "CipUnconnectedRequest"
}

func (m *_CipUnconnectedRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CipServiceContract.(*_CipService).getLengthInBits(ctx))

	// Implicit Field (requestPathSize)
	lengthInBits += 8

	// Simple field (classSegment)
	lengthInBits += m.ClassSegment.GetLengthInBits(ctx)

	// Simple field (instanceSegment)
	lengthInBits += m.InstanceSegment.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 16

	// Implicit Field (messageSize)
	lengthInBits += 16

	// Simple field (unconnectedService)
	lengthInBits += m.UnconnectedService.GetLengthInBits(ctx)

	// Const Field (route)
	lengthInBits += 16

	// Simple field (backPlane)
	lengthInBits += 8

	// Simple field (slot)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CipUnconnectedRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CipUnconnectedRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CipService, connected bool, serviceLen uint16) (__cipUnconnectedRequest CipUnconnectedRequest, err error) {
	m.CipServiceContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipUnconnectedRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipUnconnectedRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestPathSize, err := ReadImplicitField[uint8](ctx, "requestPathSize", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestPathSize' field"))
	}
	_ = requestPathSize

	classSegment, err := ReadSimpleField[PathSegment](ctx, "classSegment", ReadComplex[PathSegment](PathSegmentParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'classSegment' field"))
	}
	m.ClassSegment = classSegment

	instanceSegment, err := ReadSimpleField[PathSegment](ctx, "instanceSegment", ReadComplex[PathSegment](PathSegmentParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'instanceSegment' field"))
	}
	m.InstanceSegment = instanceSegment

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedShort(readBuffer, uint8(16)), uint16(0x9D05))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	messageSize, err := ReadImplicitField[uint16](ctx, "messageSize", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageSize' field"))
	}
	_ = messageSize

	unconnectedService, err := ReadSimpleField[CipService](ctx, "unconnectedService", ReadComplex[CipService](CipServiceParseWithBufferProducer[CipService]((bool)(bool(false)), (uint16)(messageSize)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unconnectedService' field"))
	}
	m.UnconnectedService = unconnectedService

	route, err := ReadConstField[uint16](ctx, "route", ReadUnsignedShort(readBuffer, uint8(16)), CipUnconnectedRequest_ROUTE)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'route' field"))
	}
	_ = route

	backPlane, err := ReadSimpleField(ctx, "backPlane", ReadSignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'backPlane' field"))
	}
	m.BackPlane = backPlane

	slot, err := ReadSimpleField(ctx, "slot", ReadSignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'slot' field"))
	}
	m.Slot = slot

	if closeErr := readBuffer.CloseContext("CipUnconnectedRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipUnconnectedRequest")
	}

	return m, nil
}

func (m *_CipUnconnectedRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CipUnconnectedRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipUnconnectedRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipUnconnectedRequest")
		}
		requestPathSize := uint8(uint8((uint8(m.GetClassSegment().GetLengthInBytes(ctx)) + uint8(m.GetInstanceSegment().GetLengthInBytes(ctx)))) / uint8(uint8(2)))
		if err := WriteImplicitField(ctx, "requestPathSize", requestPathSize, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestPathSize' field")
		}

		if err := WriteSimpleField[PathSegment](ctx, "classSegment", m.GetClassSegment(), WriteComplex[PathSegment](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'classSegment' field")
		}

		if err := WriteSimpleField[PathSegment](ctx, "instanceSegment", m.GetInstanceSegment(), WriteComplex[PathSegment](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'instanceSegment' field")
		}

		if err := WriteReservedField[uint16](ctx, "reserved", uint16(0x9D05), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}
		messageSize := uint16(uint16(uint16(uint16(m.GetLengthInBytes(ctx)))-uint16(uint16(10))) - uint16(uint16(4)))
		if err := WriteImplicitField(ctx, "messageSize", messageSize, WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'messageSize' field")
		}

		if err := WriteSimpleField[CipService](ctx, "unconnectedService", m.GetUnconnectedService(), WriteComplex[CipService](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'unconnectedService' field")
		}

		if err := WriteConstField(ctx, "route", CipUnconnectedRequest_ROUTE, WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'route' field")
		}

		if err := WriteSimpleField[int8](ctx, "backPlane", m.GetBackPlane(), WriteSignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'backPlane' field")
		}

		if err := WriteSimpleField[int8](ctx, "slot", m.GetSlot(), WriteSignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'slot' field")
		}

		if popErr := writeBuffer.PopContext("CipUnconnectedRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipUnconnectedRequest")
		}
		return nil
	}
	return m.CipServiceContract.(*_CipService).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CipUnconnectedRequest) IsCipUnconnectedRequest() {}

func (m *_CipUnconnectedRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
