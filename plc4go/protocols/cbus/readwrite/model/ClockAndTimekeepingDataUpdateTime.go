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

// ClockAndTimekeepingDataUpdateTime is the corresponding interface of ClockAndTimekeepingDataUpdateTime
type ClockAndTimekeepingDataUpdateTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ClockAndTimekeepingData
	// GetHours returns Hours (property field)
	GetHours() uint8
	// GetMinute returns Minute (property field)
	GetMinute() uint8
	// GetSecond returns Second (property field)
	GetSecond() uint8
	// GetDaylightSaving returns DaylightSaving (property field)
	GetDaylightSaving() byte
	// GetIsNoDaylightSavings returns IsNoDaylightSavings (virtual field)
	GetIsNoDaylightSavings() bool
	// GetIsAdvancedBy1Hour returns IsAdvancedBy1Hour (virtual field)
	GetIsAdvancedBy1Hour() bool
	// GetIsReserved returns IsReserved (virtual field)
	GetIsReserved() bool
	// GetIsUnknown returns IsUnknown (virtual field)
	GetIsUnknown() bool
	// IsClockAndTimekeepingDataUpdateTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsClockAndTimekeepingDataUpdateTime()
	// CreateBuilder creates a ClockAndTimekeepingDataUpdateTimeBuilder
	CreateClockAndTimekeepingDataUpdateTimeBuilder() ClockAndTimekeepingDataUpdateTimeBuilder
}

// _ClockAndTimekeepingDataUpdateTime is the data-structure of this message
type _ClockAndTimekeepingDataUpdateTime struct {
	ClockAndTimekeepingDataContract
	Hours          uint8
	Minute         uint8
	Second         uint8
	DaylightSaving byte
}

var _ ClockAndTimekeepingDataUpdateTime = (*_ClockAndTimekeepingDataUpdateTime)(nil)
var _ ClockAndTimekeepingDataRequirements = (*_ClockAndTimekeepingDataUpdateTime)(nil)

// NewClockAndTimekeepingDataUpdateTime factory function for _ClockAndTimekeepingDataUpdateTime
func NewClockAndTimekeepingDataUpdateTime(commandTypeContainer ClockAndTimekeepingCommandTypeContainer, argument byte, hours uint8, minute uint8, second uint8, daylightSaving byte) *_ClockAndTimekeepingDataUpdateTime {
	_result := &_ClockAndTimekeepingDataUpdateTime{
		ClockAndTimekeepingDataContract: NewClockAndTimekeepingData(commandTypeContainer, argument),
		Hours:                           hours,
		Minute:                          minute,
		Second:                          second,
		DaylightSaving:                  daylightSaving,
	}
	_result.ClockAndTimekeepingDataContract.(*_ClockAndTimekeepingData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ClockAndTimekeepingDataUpdateTimeBuilder is a builder for ClockAndTimekeepingDataUpdateTime
type ClockAndTimekeepingDataUpdateTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(hours uint8, minute uint8, second uint8, daylightSaving byte) ClockAndTimekeepingDataUpdateTimeBuilder
	// WithHours adds Hours (property field)
	WithHours(uint8) ClockAndTimekeepingDataUpdateTimeBuilder
	// WithMinute adds Minute (property field)
	WithMinute(uint8) ClockAndTimekeepingDataUpdateTimeBuilder
	// WithSecond adds Second (property field)
	WithSecond(uint8) ClockAndTimekeepingDataUpdateTimeBuilder
	// WithDaylightSaving adds DaylightSaving (property field)
	WithDaylightSaving(byte) ClockAndTimekeepingDataUpdateTimeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ClockAndTimekeepingDataBuilder
	// Build builds the ClockAndTimekeepingDataUpdateTime or returns an error if something is wrong
	Build() (ClockAndTimekeepingDataUpdateTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ClockAndTimekeepingDataUpdateTime
}

// NewClockAndTimekeepingDataUpdateTimeBuilder() creates a ClockAndTimekeepingDataUpdateTimeBuilder
func NewClockAndTimekeepingDataUpdateTimeBuilder() ClockAndTimekeepingDataUpdateTimeBuilder {
	return &_ClockAndTimekeepingDataUpdateTimeBuilder{_ClockAndTimekeepingDataUpdateTime: new(_ClockAndTimekeepingDataUpdateTime)}
}

type _ClockAndTimekeepingDataUpdateTimeBuilder struct {
	*_ClockAndTimekeepingDataUpdateTime

	parentBuilder *_ClockAndTimekeepingDataBuilder

	err *utils.MultiError
}

var _ (ClockAndTimekeepingDataUpdateTimeBuilder) = (*_ClockAndTimekeepingDataUpdateTimeBuilder)(nil)

func (b *_ClockAndTimekeepingDataUpdateTimeBuilder) setParent(contract ClockAndTimekeepingDataContract) {
	b.ClockAndTimekeepingDataContract = contract
}

func (b *_ClockAndTimekeepingDataUpdateTimeBuilder) WithMandatoryFields(hours uint8, minute uint8, second uint8, daylightSaving byte) ClockAndTimekeepingDataUpdateTimeBuilder {
	return b.WithHours(hours).WithMinute(minute).WithSecond(second).WithDaylightSaving(daylightSaving)
}

func (b *_ClockAndTimekeepingDataUpdateTimeBuilder) WithHours(hours uint8) ClockAndTimekeepingDataUpdateTimeBuilder {
	b.Hours = hours
	return b
}

func (b *_ClockAndTimekeepingDataUpdateTimeBuilder) WithMinute(minute uint8) ClockAndTimekeepingDataUpdateTimeBuilder {
	b.Minute = minute
	return b
}

func (b *_ClockAndTimekeepingDataUpdateTimeBuilder) WithSecond(second uint8) ClockAndTimekeepingDataUpdateTimeBuilder {
	b.Second = second
	return b
}

func (b *_ClockAndTimekeepingDataUpdateTimeBuilder) WithDaylightSaving(daylightSaving byte) ClockAndTimekeepingDataUpdateTimeBuilder {
	b.DaylightSaving = daylightSaving
	return b
}

func (b *_ClockAndTimekeepingDataUpdateTimeBuilder) Build() (ClockAndTimekeepingDataUpdateTime, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ClockAndTimekeepingDataUpdateTime.deepCopy(), nil
}

func (b *_ClockAndTimekeepingDataUpdateTimeBuilder) MustBuild() ClockAndTimekeepingDataUpdateTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ClockAndTimekeepingDataUpdateTimeBuilder) Done() ClockAndTimekeepingDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewClockAndTimekeepingDataBuilder().(*_ClockAndTimekeepingDataBuilder)
	}
	return b.parentBuilder
}

func (b *_ClockAndTimekeepingDataUpdateTimeBuilder) buildForClockAndTimekeepingData() (ClockAndTimekeepingData, error) {
	return b.Build()
}

func (b *_ClockAndTimekeepingDataUpdateTimeBuilder) DeepCopy() any {
	_copy := b.CreateClockAndTimekeepingDataUpdateTimeBuilder().(*_ClockAndTimekeepingDataUpdateTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateClockAndTimekeepingDataUpdateTimeBuilder creates a ClockAndTimekeepingDataUpdateTimeBuilder
func (b *_ClockAndTimekeepingDataUpdateTime) CreateClockAndTimekeepingDataUpdateTimeBuilder() ClockAndTimekeepingDataUpdateTimeBuilder {
	if b == nil {
		return NewClockAndTimekeepingDataUpdateTimeBuilder()
	}
	return &_ClockAndTimekeepingDataUpdateTimeBuilder{_ClockAndTimekeepingDataUpdateTime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ClockAndTimekeepingDataUpdateTime) GetParent() ClockAndTimekeepingDataContract {
	return m.ClockAndTimekeepingDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ClockAndTimekeepingDataUpdateTime) GetHours() uint8 {
	return m.Hours
}

func (m *_ClockAndTimekeepingDataUpdateTime) GetMinute() uint8 {
	return m.Minute
}

func (m *_ClockAndTimekeepingDataUpdateTime) GetSecond() uint8 {
	return m.Second
}

func (m *_ClockAndTimekeepingDataUpdateTime) GetDaylightSaving() byte {
	return m.DaylightSaving
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_ClockAndTimekeepingDataUpdateTime) GetIsNoDaylightSavings() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetDaylightSaving()) == (0x00)))
}

func (m *_ClockAndTimekeepingDataUpdateTime) GetIsAdvancedBy1Hour() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetDaylightSaving()) == (0x01)))
}

func (m *_ClockAndTimekeepingDataUpdateTime) GetIsReserved() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool(bool((m.GetDaylightSaving()) > (0x01))) && bool(bool((m.GetDaylightSaving()) <= (0xFE))))
}

func (m *_ClockAndTimekeepingDataUpdateTime) GetIsUnknown() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetDaylightSaving()) > (0xFE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastClockAndTimekeepingDataUpdateTime(structType any) ClockAndTimekeepingDataUpdateTime {
	if casted, ok := structType.(ClockAndTimekeepingDataUpdateTime); ok {
		return casted
	}
	if casted, ok := structType.(*ClockAndTimekeepingDataUpdateTime); ok {
		return *casted
	}
	return nil
}

func (m *_ClockAndTimekeepingDataUpdateTime) GetTypeName() string {
	return "ClockAndTimekeepingDataUpdateTime"
}

func (m *_ClockAndTimekeepingDataUpdateTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ClockAndTimekeepingDataContract.(*_ClockAndTimekeepingData).getLengthInBits(ctx))

	// Simple field (hours)
	lengthInBits += 8

	// Simple field (minute)
	lengthInBits += 8

	// Simple field (second)
	lengthInBits += 8

	// Simple field (daylightSaving)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_ClockAndTimekeepingDataUpdateTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ClockAndTimekeepingDataUpdateTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ClockAndTimekeepingData) (__clockAndTimekeepingDataUpdateTime ClockAndTimekeepingDataUpdateTime, err error) {
	m.ClockAndTimekeepingDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ClockAndTimekeepingDataUpdateTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ClockAndTimekeepingDataUpdateTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	hours, err := ReadSimpleField(ctx, "hours", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hours' field"))
	}
	m.Hours = hours

	minute, err := ReadSimpleField(ctx, "minute", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minute' field"))
	}
	m.Minute = minute

	second, err := ReadSimpleField(ctx, "second", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'second' field"))
	}
	m.Second = second

	daylightSaving, err := ReadSimpleField(ctx, "daylightSaving", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'daylightSaving' field"))
	}
	m.DaylightSaving = daylightSaving

	isNoDaylightSavings, err := ReadVirtualField[bool](ctx, "isNoDaylightSavings", (*bool)(nil), bool((daylightSaving) == (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isNoDaylightSavings' field"))
	}
	_ = isNoDaylightSavings

	isAdvancedBy1Hour, err := ReadVirtualField[bool](ctx, "isAdvancedBy1Hour", (*bool)(nil), bool((daylightSaving) == (0x01)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isAdvancedBy1Hour' field"))
	}
	_ = isAdvancedBy1Hour

	isReserved, err := ReadVirtualField[bool](ctx, "isReserved", (*bool)(nil), bool(bool((daylightSaving) > (0x01))) && bool(bool((daylightSaving) <= (0xFE))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isReserved' field"))
	}
	_ = isReserved

	isUnknown, err := ReadVirtualField[bool](ctx, "isUnknown", (*bool)(nil), bool((daylightSaving) > (0xFE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isUnknown' field"))
	}
	_ = isUnknown

	if closeErr := readBuffer.CloseContext("ClockAndTimekeepingDataUpdateTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ClockAndTimekeepingDataUpdateTime")
	}

	return m, nil
}

func (m *_ClockAndTimekeepingDataUpdateTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ClockAndTimekeepingDataUpdateTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ClockAndTimekeepingDataUpdateTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ClockAndTimekeepingDataUpdateTime")
		}

		if err := WriteSimpleField[uint8](ctx, "hours", m.GetHours(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'hours' field")
		}

		if err := WriteSimpleField[uint8](ctx, "minute", m.GetMinute(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'minute' field")
		}

		if err := WriteSimpleField[uint8](ctx, "second", m.GetSecond(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'second' field")
		}

		if err := WriteSimpleField[byte](ctx, "daylightSaving", m.GetDaylightSaving(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'daylightSaving' field")
		}
		// Virtual field
		isNoDaylightSavings := m.GetIsNoDaylightSavings()
		_ = isNoDaylightSavings
		if _isNoDaylightSavingsErr := writeBuffer.WriteVirtual(ctx, "isNoDaylightSavings", m.GetIsNoDaylightSavings()); _isNoDaylightSavingsErr != nil {
			return errors.Wrap(_isNoDaylightSavingsErr, "Error serializing 'isNoDaylightSavings' field")
		}
		// Virtual field
		isAdvancedBy1Hour := m.GetIsAdvancedBy1Hour()
		_ = isAdvancedBy1Hour
		if _isAdvancedBy1HourErr := writeBuffer.WriteVirtual(ctx, "isAdvancedBy1Hour", m.GetIsAdvancedBy1Hour()); _isAdvancedBy1HourErr != nil {
			return errors.Wrap(_isAdvancedBy1HourErr, "Error serializing 'isAdvancedBy1Hour' field")
		}
		// Virtual field
		isReserved := m.GetIsReserved()
		_ = isReserved
		if _isReservedErr := writeBuffer.WriteVirtual(ctx, "isReserved", m.GetIsReserved()); _isReservedErr != nil {
			return errors.Wrap(_isReservedErr, "Error serializing 'isReserved' field")
		}
		// Virtual field
		isUnknown := m.GetIsUnknown()
		_ = isUnknown
		if _isUnknownErr := writeBuffer.WriteVirtual(ctx, "isUnknown", m.GetIsUnknown()); _isUnknownErr != nil {
			return errors.Wrap(_isUnknownErr, "Error serializing 'isUnknown' field")
		}

		if popErr := writeBuffer.PopContext("ClockAndTimekeepingDataUpdateTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ClockAndTimekeepingDataUpdateTime")
		}
		return nil
	}
	return m.ClockAndTimekeepingDataContract.(*_ClockAndTimekeepingData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ClockAndTimekeepingDataUpdateTime) IsClockAndTimekeepingDataUpdateTime() {}

func (m *_ClockAndTimekeepingDataUpdateTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ClockAndTimekeepingDataUpdateTime) deepCopy() *_ClockAndTimekeepingDataUpdateTime {
	if m == nil {
		return nil
	}
	_ClockAndTimekeepingDataUpdateTimeCopy := &_ClockAndTimekeepingDataUpdateTime{
		m.ClockAndTimekeepingDataContract.(*_ClockAndTimekeepingData).deepCopy(),
		m.Hours,
		m.Minute,
		m.Second,
		m.DaylightSaving,
	}
	m.ClockAndTimekeepingDataContract.(*_ClockAndTimekeepingData)._SubType = m
	return _ClockAndTimekeepingDataUpdateTimeCopy
}

func (m *_ClockAndTimekeepingDataUpdateTime) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
