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

// MeteringDataDrinkingWaterConsumption is the corresponding interface of MeteringDataDrinkingWaterConsumption
type MeteringDataDrinkingWaterConsumption interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MeteringData
	// GetKL returns KL (property field)
	GetKL() uint32
}

// MeteringDataDrinkingWaterConsumptionExactly can be used when we want exactly this type and not a type which fulfills MeteringDataDrinkingWaterConsumption.
// This is useful for switch cases.
type MeteringDataDrinkingWaterConsumptionExactly interface {
	MeteringDataDrinkingWaterConsumption
	isMeteringDataDrinkingWaterConsumption() bool
}

// _MeteringDataDrinkingWaterConsumption is the data-structure of this message
type _MeteringDataDrinkingWaterConsumption struct {
	*_MeteringData
	KL uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MeteringDataDrinkingWaterConsumption) InitializeParent(parent MeteringData, commandTypeContainer MeteringCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_MeteringDataDrinkingWaterConsumption) GetParent() MeteringData {
	return m._MeteringData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MeteringDataDrinkingWaterConsumption) GetKL() uint32 {
	return m.KL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMeteringDataDrinkingWaterConsumption factory function for _MeteringDataDrinkingWaterConsumption
func NewMeteringDataDrinkingWaterConsumption(kL uint32, commandTypeContainer MeteringCommandTypeContainer, argument byte) *_MeteringDataDrinkingWaterConsumption {
	_result := &_MeteringDataDrinkingWaterConsumption{
		KL:            kL,
		_MeteringData: NewMeteringData(commandTypeContainer, argument),
	}
	_result._MeteringData._MeteringDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMeteringDataDrinkingWaterConsumption(structType any) MeteringDataDrinkingWaterConsumption {
	if casted, ok := structType.(MeteringDataDrinkingWaterConsumption); ok {
		return casted
	}
	if casted, ok := structType.(*MeteringDataDrinkingWaterConsumption); ok {
		return *casted
	}
	return nil
}

func (m *_MeteringDataDrinkingWaterConsumption) GetTypeName() string {
	return "MeteringDataDrinkingWaterConsumption"
}

func (m *_MeteringDataDrinkingWaterConsumption) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (kL)
	lengthInBits += 32

	return lengthInBits
}

func (m *_MeteringDataDrinkingWaterConsumption) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MeteringDataDrinkingWaterConsumptionParse(ctx context.Context, theBytes []byte) (MeteringDataDrinkingWaterConsumption, error) {
	return MeteringDataDrinkingWaterConsumptionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func MeteringDataDrinkingWaterConsumptionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (MeteringDataDrinkingWaterConsumption, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("MeteringDataDrinkingWaterConsumption"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MeteringDataDrinkingWaterConsumption")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (kL)
	_kL, _kLErr := readBuffer.ReadUint32("kL", 32)
	if _kLErr != nil {
		return nil, errors.Wrap(_kLErr, "Error parsing 'kL' field of MeteringDataDrinkingWaterConsumption")
	}
	kL := _kL

	if closeErr := readBuffer.CloseContext("MeteringDataDrinkingWaterConsumption"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MeteringDataDrinkingWaterConsumption")
	}

	// Create a partially initialized instance
	_child := &_MeteringDataDrinkingWaterConsumption{
		_MeteringData: &_MeteringData{},
		KL:            kL,
	}
	_child._MeteringData._MeteringDataChildRequirements = _child
	return _child, nil
}

func (m *_MeteringDataDrinkingWaterConsumption) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MeteringDataDrinkingWaterConsumption) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MeteringDataDrinkingWaterConsumption"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MeteringDataDrinkingWaterConsumption")
		}

		// Simple Field (kL)
		kL := uint32(m.GetKL())
		_kLErr := writeBuffer.WriteUint32("kL", 32, uint32((kL)))
		if _kLErr != nil {
			return errors.Wrap(_kLErr, "Error serializing 'kL' field")
		}

		if popErr := writeBuffer.PopContext("MeteringDataDrinkingWaterConsumption"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MeteringDataDrinkingWaterConsumption")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MeteringDataDrinkingWaterConsumption) isMeteringDataDrinkingWaterConsumption() bool {
	return true
}

func (m *_MeteringDataDrinkingWaterConsumption) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
