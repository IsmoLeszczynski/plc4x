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

// CALDataAcknowledge is the corresponding interface of CALDataAcknowledge
type CALDataAcknowledge interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CALData
	// GetParamNo returns ParamNo (property field)
	GetParamNo() Parameter
	// GetCode returns Code (property field)
	GetCode() uint8
}

// _CALDataAcknowledge is the data-structure of this message
type _CALDataAcknowledge struct {
	CALDataContract
	ParamNo Parameter
	Code    uint8
}

var _ CALDataAcknowledge = (*_CALDataAcknowledge)(nil)
var _ CALDataRequirements = (*_CALDataAcknowledge)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALDataAcknowledge) GetParent() CALDataContract {
	return m.CALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALDataAcknowledge) GetParamNo() Parameter {
	return m.ParamNo
}

func (m *_CALDataAcknowledge) GetCode() uint8 {
	return m.Code
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALDataAcknowledge factory function for _CALDataAcknowledge
func NewCALDataAcknowledge(paramNo Parameter, code uint8, commandTypeContainer CALCommandTypeContainer, additionalData CALData, requestContext RequestContext) *_CALDataAcknowledge {
	_result := &_CALDataAcknowledge{
		CALDataContract: NewCALData(commandTypeContainer, additionalData, requestContext),
		ParamNo:         paramNo,
		Code:            code,
	}
	_result.CALDataContract.(*_CALData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALDataAcknowledge(structType any) CALDataAcknowledge {
	if casted, ok := structType.(CALDataAcknowledge); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataAcknowledge); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataAcknowledge) GetTypeName() string {
	return "CALDataAcknowledge"
}

func (m *_CALDataAcknowledge) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CALDataContract.(*_CALData).getLengthInBits(ctx))

	// Simple field (paramNo)
	lengthInBits += 8

	// Simple field (code)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CALDataAcknowledge) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CALDataAcknowledge) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CALData, requestContext RequestContext) (__cALDataAcknowledge CALDataAcknowledge, err error) {
	m.CALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataAcknowledge"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataAcknowledge")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	paramNo, err := ReadEnumField[Parameter](ctx, "paramNo", "Parameter", ReadEnum(ParameterByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'paramNo' field"))
	}
	m.ParamNo = paramNo

	code, err := ReadSimpleField(ctx, "code", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'code' field"))
	}
	m.Code = code

	if closeErr := readBuffer.CloseContext("CALDataAcknowledge"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataAcknowledge")
	}

	return m, nil
}

func (m *_CALDataAcknowledge) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CALDataAcknowledge) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataAcknowledge"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataAcknowledge")
		}

		if err := WriteSimpleEnumField[Parameter](ctx, "paramNo", "Parameter", m.GetParamNo(), WriteEnum[Parameter, uint8](Parameter.GetValue, Parameter.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'paramNo' field")
		}

		if err := WriteSimpleField[uint8](ctx, "code", m.GetCode(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'code' field")
		}

		if popErr := writeBuffer.PopContext("CALDataAcknowledge"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataAcknowledge")
		}
		return nil
	}
	return m.CALDataContract.(*_CALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CALDataAcknowledge) IsCALDataAcknowledge() {}

func (m *_CALDataAcknowledge) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
