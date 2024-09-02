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

// SALDataEnableControl is the corresponding interface of SALDataEnableControl
type SALDataEnableControl interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SALData
	// GetEnableControlData returns EnableControlData (property field)
	GetEnableControlData() EnableControlData
}

// _SALDataEnableControl is the data-structure of this message
type _SALDataEnableControl struct {
	SALDataContract
	EnableControlData EnableControlData
}

var _ SALDataEnableControl = (*_SALDataEnableControl)(nil)
var _ SALDataRequirements = (*_SALDataEnableControl)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataEnableControl) GetApplicationId() ApplicationId {
	return ApplicationId_ENABLE_CONTROL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataEnableControl) GetParent() SALDataContract {
	return m.SALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataEnableControl) GetEnableControlData() EnableControlData {
	return m.EnableControlData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALDataEnableControl factory function for _SALDataEnableControl
func NewSALDataEnableControl(enableControlData EnableControlData, salData SALData) *_SALDataEnableControl {
	_result := &_SALDataEnableControl{
		SALDataContract:   NewSALData(salData),
		EnableControlData: enableControlData,
	}
	_result.SALDataContract.(*_SALData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataEnableControl(structType any) SALDataEnableControl {
	if casted, ok := structType.(SALDataEnableControl); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataEnableControl); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataEnableControl) GetTypeName() string {
	return "SALDataEnableControl"
}

func (m *_SALDataEnableControl) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SALDataContract.(*_SALData).getLengthInBits(ctx))

	// Simple field (enableControlData)
	lengthInBits += m.EnableControlData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataEnableControl) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SALDataEnableControl) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SALData, applicationId ApplicationId) (__sALDataEnableControl SALDataEnableControl, err error) {
	m.SALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataEnableControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataEnableControl")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	enableControlData, err := ReadSimpleField[EnableControlData](ctx, "enableControlData", ReadComplex[EnableControlData](EnableControlDataParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enableControlData' field"))
	}
	m.EnableControlData = enableControlData

	if closeErr := readBuffer.CloseContext("SALDataEnableControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataEnableControl")
	}

	return m, nil
}

func (m *_SALDataEnableControl) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataEnableControl) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataEnableControl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataEnableControl")
		}

		if err := WriteSimpleField[EnableControlData](ctx, "enableControlData", m.GetEnableControlData(), WriteComplex[EnableControlData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'enableControlData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataEnableControl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataEnableControl")
		}
		return nil
	}
	return m.SALDataContract.(*_SALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataEnableControl) IsSALDataEnableControl() {}

func (m *_SALDataEnableControl) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
