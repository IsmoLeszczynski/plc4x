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

// SALDataAirConditioning is the corresponding interface of SALDataAirConditioning
type SALDataAirConditioning interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SALData
	// GetAirConditioningData returns AirConditioningData (property field)
	GetAirConditioningData() AirConditioningData
}

// SALDataAirConditioningExactly can be used when we want exactly this type and not a type which fulfills SALDataAirConditioning.
// This is useful for switch cases.
type SALDataAirConditioningExactly interface {
	SALDataAirConditioning
	isSALDataAirConditioning() bool
}

// _SALDataAirConditioning is the data-structure of this message
type _SALDataAirConditioning struct {
	SALDataContract
	AirConditioningData AirConditioningData
}

var _ SALDataAirConditioning = (*_SALDataAirConditioning)(nil)
var _ SALDataRequirements = (*_SALDataAirConditioning)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataAirConditioning) GetApplicationId() ApplicationId {
	return ApplicationId_AIR_CONDITIONING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataAirConditioning) GetParent() SALDataContract {
	return m.SALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataAirConditioning) GetAirConditioningData() AirConditioningData {
	return m.AirConditioningData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALDataAirConditioning factory function for _SALDataAirConditioning
func NewSALDataAirConditioning(airConditioningData AirConditioningData, salData SALData) *_SALDataAirConditioning {
	_result := &_SALDataAirConditioning{
		SALDataContract:     NewSALData(salData),
		AirConditioningData: airConditioningData,
	}
	_result.SALDataContract.(*_SALData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataAirConditioning(structType any) SALDataAirConditioning {
	if casted, ok := structType.(SALDataAirConditioning); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataAirConditioning); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataAirConditioning) GetTypeName() string {
	return "SALDataAirConditioning"
}

func (m *_SALDataAirConditioning) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SALDataContract.(*_SALData).getLengthInBits(ctx))

	// Simple field (airConditioningData)
	lengthInBits += m.AirConditioningData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataAirConditioning) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SALDataAirConditioning) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SALData, applicationId ApplicationId) (__sALDataAirConditioning SALDataAirConditioning, err error) {
	m.SALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataAirConditioning"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataAirConditioning")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	airConditioningData, err := ReadSimpleField[AirConditioningData](ctx, "airConditioningData", ReadComplex[AirConditioningData](AirConditioningDataParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'airConditioningData' field"))
	}
	m.AirConditioningData = airConditioningData

	if closeErr := readBuffer.CloseContext("SALDataAirConditioning"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataAirConditioning")
	}

	return m, nil
}

func (m *_SALDataAirConditioning) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataAirConditioning) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataAirConditioning"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataAirConditioning")
		}

		if err := WriteSimpleField[AirConditioningData](ctx, "airConditioningData", m.GetAirConditioningData(), WriteComplex[AirConditioningData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'airConditioningData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataAirConditioning"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataAirConditioning")
		}
		return nil
	}
	return m.SALDataContract.(*_SALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataAirConditioning) isSALDataAirConditioning() bool {
	return true
}

func (m *_SALDataAirConditioning) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
