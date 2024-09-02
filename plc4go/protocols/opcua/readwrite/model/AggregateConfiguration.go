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

// AggregateConfiguration is the corresponding interface of AggregateConfiguration
type AggregateConfiguration interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetTreatUncertainAsBad returns TreatUncertainAsBad (property field)
	GetTreatUncertainAsBad() bool
	// GetUseServerCapabilitiesDefaults returns UseServerCapabilitiesDefaults (property field)
	GetUseServerCapabilitiesDefaults() bool
	// GetPercentDataBad returns PercentDataBad (property field)
	GetPercentDataBad() uint8
	// GetPercentDataGood returns PercentDataGood (property field)
	GetPercentDataGood() uint8
	// GetUseSlopedExtrapolation returns UseSlopedExtrapolation (property field)
	GetUseSlopedExtrapolation() bool
}

// AggregateConfigurationExactly can be used when we want exactly this type and not a type which fulfills AggregateConfiguration.
// This is useful for switch cases.
type AggregateConfigurationExactly interface {
	AggregateConfiguration
	isAggregateConfiguration() bool
}

// _AggregateConfiguration is the data-structure of this message
type _AggregateConfiguration struct {
	ExtensionObjectDefinitionContract
	TreatUncertainAsBad           bool
	UseServerCapabilitiesDefaults bool
	PercentDataBad                uint8
	PercentDataGood               uint8
	UseSlopedExtrapolation        bool
	// Reserved Fields
	reservedField0 *uint8
	reservedField1 *uint8
}

var _ AggregateConfiguration = (*_AggregateConfiguration)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_AggregateConfiguration)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AggregateConfiguration) GetIdentifier() string {
	return "950"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AggregateConfiguration) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AggregateConfiguration) GetTreatUncertainAsBad() bool {
	return m.TreatUncertainAsBad
}

func (m *_AggregateConfiguration) GetUseServerCapabilitiesDefaults() bool {
	return m.UseServerCapabilitiesDefaults
}

func (m *_AggregateConfiguration) GetPercentDataBad() uint8 {
	return m.PercentDataBad
}

func (m *_AggregateConfiguration) GetPercentDataGood() uint8 {
	return m.PercentDataGood
}

func (m *_AggregateConfiguration) GetUseSlopedExtrapolation() bool {
	return m.UseSlopedExtrapolation
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAggregateConfiguration factory function for _AggregateConfiguration
func NewAggregateConfiguration(treatUncertainAsBad bool, useServerCapabilitiesDefaults bool, percentDataBad uint8, percentDataGood uint8, useSlopedExtrapolation bool) *_AggregateConfiguration {
	_result := &_AggregateConfiguration{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		TreatUncertainAsBad:               treatUncertainAsBad,
		UseServerCapabilitiesDefaults:     useServerCapabilitiesDefaults,
		PercentDataBad:                    percentDataBad,
		PercentDataGood:                   percentDataGood,
		UseSlopedExtrapolation:            useSlopedExtrapolation,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAggregateConfiguration(structType any) AggregateConfiguration {
	if casted, ok := structType.(AggregateConfiguration); ok {
		return casted
	}
	if casted, ok := structType.(*AggregateConfiguration); ok {
		return *casted
	}
	return nil
}

func (m *_AggregateConfiguration) GetTypeName() string {
	return "AggregateConfiguration"
}

func (m *_AggregateConfiguration) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 6

	// Simple field (treatUncertainAsBad)
	lengthInBits += 1

	// Simple field (useServerCapabilitiesDefaults)
	lengthInBits += 1

	// Simple field (percentDataBad)
	lengthInBits += 8

	// Simple field (percentDataGood)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (useSlopedExtrapolation)
	lengthInBits += 1

	return lengthInBits
}

func (m *_AggregateConfiguration) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AggregateConfiguration) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__aggregateConfiguration AggregateConfiguration, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AggregateConfiguration"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AggregateConfiguration")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(6)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	treatUncertainAsBad, err := ReadSimpleField(ctx, "treatUncertainAsBad", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'treatUncertainAsBad' field"))
	}
	m.TreatUncertainAsBad = treatUncertainAsBad

	useServerCapabilitiesDefaults, err := ReadSimpleField(ctx, "useServerCapabilitiesDefaults", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'useServerCapabilitiesDefaults' field"))
	}
	m.UseServerCapabilitiesDefaults = useServerCapabilitiesDefaults

	percentDataBad, err := ReadSimpleField(ctx, "percentDataBad", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'percentDataBad' field"))
	}
	m.PercentDataBad = percentDataBad

	percentDataGood, err := ReadSimpleField(ctx, "percentDataGood", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'percentDataGood' field"))
	}
	m.PercentDataGood = percentDataGood

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField1 = reservedField1

	useSlopedExtrapolation, err := ReadSimpleField(ctx, "useSlopedExtrapolation", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'useSlopedExtrapolation' field"))
	}
	m.UseSlopedExtrapolation = useSlopedExtrapolation

	if closeErr := readBuffer.CloseContext("AggregateConfiguration"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AggregateConfiguration")
	}

	return m, nil
}

func (m *_AggregateConfiguration) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AggregateConfiguration) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AggregateConfiguration"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AggregateConfiguration")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 6)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "treatUncertainAsBad", m.GetTreatUncertainAsBad(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'treatUncertainAsBad' field")
		}

		if err := WriteSimpleField[bool](ctx, "useServerCapabilitiesDefaults", m.GetUseServerCapabilitiesDefaults(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'useServerCapabilitiesDefaults' field")
		}

		if err := WriteSimpleField[uint8](ctx, "percentDataBad", m.GetPercentDataBad(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'percentDataBad' field")
		}

		if err := WriteSimpleField[uint8](ctx, "percentDataGood", m.GetPercentDataGood(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'percentDataGood' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 2")
		}

		if err := WriteSimpleField[bool](ctx, "useSlopedExtrapolation", m.GetUseSlopedExtrapolation(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'useSlopedExtrapolation' field")
		}

		if popErr := writeBuffer.PopContext("AggregateConfiguration"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AggregateConfiguration")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AggregateConfiguration) isAggregateConfiguration() bool {
	return true
}

func (m *_AggregateConfiguration) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
