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

// SALDataTriggerControl is the corresponding interface of SALDataTriggerControl
type SALDataTriggerControl interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SALData
	// GetTriggerControlData returns TriggerControlData (property field)
	GetTriggerControlData() TriggerControlData
}

// SALDataTriggerControlExactly can be used when we want exactly this type and not a type which fulfills SALDataTriggerControl.
// This is useful for switch cases.
type SALDataTriggerControlExactly interface {
	SALDataTriggerControl
	isSALDataTriggerControl() bool
}

// _SALDataTriggerControl is the data-structure of this message
type _SALDataTriggerControl struct {
	SALDataContract
	TriggerControlData TriggerControlData
}

var _ SALDataTriggerControl = (*_SALDataTriggerControl)(nil)
var _ SALDataRequirements = (*_SALDataTriggerControl)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataTriggerControl) GetApplicationId() ApplicationId {
	return ApplicationId_TRIGGER_CONTROL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataTriggerControl) GetParent() SALDataContract {
	return m.SALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataTriggerControl) GetTriggerControlData() TriggerControlData {
	return m.TriggerControlData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALDataTriggerControl factory function for _SALDataTriggerControl
func NewSALDataTriggerControl(triggerControlData TriggerControlData, salData SALData) *_SALDataTriggerControl {
	_result := &_SALDataTriggerControl{
		SALDataContract:    NewSALData(salData),
		TriggerControlData: triggerControlData,
	}
	_result.SALDataContract.(*_SALData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataTriggerControl(structType any) SALDataTriggerControl {
	if casted, ok := structType.(SALDataTriggerControl); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataTriggerControl); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataTriggerControl) GetTypeName() string {
	return "SALDataTriggerControl"
}

func (m *_SALDataTriggerControl) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SALDataContract.(*_SALData).getLengthInBits(ctx))

	// Simple field (triggerControlData)
	lengthInBits += m.TriggerControlData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataTriggerControl) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SALDataTriggerControl) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SALData, applicationId ApplicationId) (__sALDataTriggerControl SALDataTriggerControl, err error) {
	m.SALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataTriggerControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataTriggerControl")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	triggerControlData, err := ReadSimpleField[TriggerControlData](ctx, "triggerControlData", ReadComplex[TriggerControlData](TriggerControlDataParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'triggerControlData' field"))
	}
	m.TriggerControlData = triggerControlData

	if closeErr := readBuffer.CloseContext("SALDataTriggerControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataTriggerControl")
	}

	return m, nil
}

func (m *_SALDataTriggerControl) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataTriggerControl) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataTriggerControl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataTriggerControl")
		}

		if err := WriteSimpleField[TriggerControlData](ctx, "triggerControlData", m.GetTriggerControlData(), WriteComplex[TriggerControlData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'triggerControlData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataTriggerControl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataTriggerControl")
		}
		return nil
	}
	return m.SALDataContract.(*_SALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataTriggerControl) isSALDataTriggerControl() bool {
	return true
}

func (m *_SALDataTriggerControl) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
