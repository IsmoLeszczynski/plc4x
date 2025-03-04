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

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// DateAndTime is the corresponding interface of DateAndTime
type DateAndTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetYear returns Year (property field)
	GetYear() uint8
	// GetMonth returns Month (property field)
	GetMonth() uint8
	// GetDay returns Day (property field)
	GetDay() uint8
	// GetHour returns Hour (property field)
	GetHour() uint8
	// GetMinutes returns Minutes (property field)
	GetMinutes() uint8
	// GetSeconds returns Seconds (property field)
	GetSeconds() uint8
	// GetMsec returns Msec (property field)
	GetMsec() uint16
	// GetDow returns Dow (property field)
	GetDow() uint8
	// IsDateAndTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDateAndTime()
	// CreateBuilder creates a DateAndTimeBuilder
	CreateDateAndTimeBuilder() DateAndTimeBuilder
}

// _DateAndTime is the data-structure of this message
type _DateAndTime struct {
	Year    uint8
	Month   uint8
	Day     uint8
	Hour    uint8
	Minutes uint8
	Seconds uint8
	Msec    uint16
	Dow     uint8
}

var _ DateAndTime = (*_DateAndTime)(nil)

// NewDateAndTime factory function for _DateAndTime
func NewDateAndTime(year uint8, month uint8, day uint8, hour uint8, minutes uint8, seconds uint8, msec uint16, dow uint8) *_DateAndTime {
	return &_DateAndTime{Year: year, Month: month, Day: day, Hour: hour, Minutes: minutes, Seconds: seconds, Msec: msec, Dow: dow}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DateAndTimeBuilder is a builder for DateAndTime
type DateAndTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(year uint8, month uint8, day uint8, hour uint8, minutes uint8, seconds uint8, msec uint16, dow uint8) DateAndTimeBuilder
	// WithYear adds Year (property field)
	WithYear(uint8) DateAndTimeBuilder
	// WithMonth adds Month (property field)
	WithMonth(uint8) DateAndTimeBuilder
	// WithDay adds Day (property field)
	WithDay(uint8) DateAndTimeBuilder
	// WithHour adds Hour (property field)
	WithHour(uint8) DateAndTimeBuilder
	// WithMinutes adds Minutes (property field)
	WithMinutes(uint8) DateAndTimeBuilder
	// WithSeconds adds Seconds (property field)
	WithSeconds(uint8) DateAndTimeBuilder
	// WithMsec adds Msec (property field)
	WithMsec(uint16) DateAndTimeBuilder
	// WithDow adds Dow (property field)
	WithDow(uint8) DateAndTimeBuilder
	// Build builds the DateAndTime or returns an error if something is wrong
	Build() (DateAndTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DateAndTime
}

// NewDateAndTimeBuilder() creates a DateAndTimeBuilder
func NewDateAndTimeBuilder() DateAndTimeBuilder {
	return &_DateAndTimeBuilder{_DateAndTime: new(_DateAndTime)}
}

type _DateAndTimeBuilder struct {
	*_DateAndTime

	err *utils.MultiError
}

var _ (DateAndTimeBuilder) = (*_DateAndTimeBuilder)(nil)

func (b *_DateAndTimeBuilder) WithMandatoryFields(year uint8, month uint8, day uint8, hour uint8, minutes uint8, seconds uint8, msec uint16, dow uint8) DateAndTimeBuilder {
	return b.WithYear(year).WithMonth(month).WithDay(day).WithHour(hour).WithMinutes(minutes).WithSeconds(seconds).WithMsec(msec).WithDow(dow)
}

func (b *_DateAndTimeBuilder) WithYear(year uint8) DateAndTimeBuilder {
	b.Year = year
	return b
}

func (b *_DateAndTimeBuilder) WithMonth(month uint8) DateAndTimeBuilder {
	b.Month = month
	return b
}

func (b *_DateAndTimeBuilder) WithDay(day uint8) DateAndTimeBuilder {
	b.Day = day
	return b
}

func (b *_DateAndTimeBuilder) WithHour(hour uint8) DateAndTimeBuilder {
	b.Hour = hour
	return b
}

func (b *_DateAndTimeBuilder) WithMinutes(minutes uint8) DateAndTimeBuilder {
	b.Minutes = minutes
	return b
}

func (b *_DateAndTimeBuilder) WithSeconds(seconds uint8) DateAndTimeBuilder {
	b.Seconds = seconds
	return b
}

func (b *_DateAndTimeBuilder) WithMsec(msec uint16) DateAndTimeBuilder {
	b.Msec = msec
	return b
}

func (b *_DateAndTimeBuilder) WithDow(dow uint8) DateAndTimeBuilder {
	b.Dow = dow
	return b
}

func (b *_DateAndTimeBuilder) Build() (DateAndTime, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DateAndTime.deepCopy(), nil
}

func (b *_DateAndTimeBuilder) MustBuild() DateAndTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_DateAndTimeBuilder) DeepCopy() any {
	_copy := b.CreateDateAndTimeBuilder().(*_DateAndTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDateAndTimeBuilder creates a DateAndTimeBuilder
func (b *_DateAndTime) CreateDateAndTimeBuilder() DateAndTimeBuilder {
	if b == nil {
		return NewDateAndTimeBuilder()
	}
	return &_DateAndTimeBuilder{_DateAndTime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DateAndTime) GetYear() uint8 {
	return m.Year
}

func (m *_DateAndTime) GetMonth() uint8 {
	return m.Month
}

func (m *_DateAndTime) GetDay() uint8 {
	return m.Day
}

func (m *_DateAndTime) GetHour() uint8 {
	return m.Hour
}

func (m *_DateAndTime) GetMinutes() uint8 {
	return m.Minutes
}

func (m *_DateAndTime) GetSeconds() uint8 {
	return m.Seconds
}

func (m *_DateAndTime) GetMsec() uint16 {
	return m.Msec
}

func (m *_DateAndTime) GetDow() uint8 {
	return m.Dow
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDateAndTime(structType any) DateAndTime {
	if casted, ok := structType.(DateAndTime); ok {
		return casted
	}
	if casted, ok := structType.(*DateAndTime); ok {
		return *casted
	}
	return nil
}

func (m *_DateAndTime) GetTypeName() string {
	return "DateAndTime"
}

func (m *_DateAndTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (year)
	lengthInBits += 8

	// Simple field (month)
	lengthInBits += 8

	// Simple field (day)
	lengthInBits += 8

	// Simple field (hour)
	lengthInBits += 8

	// Simple field (minutes)
	lengthInBits += 8

	// Simple field (seconds)
	lengthInBits += 8

	// Simple field (msec)
	lengthInBits += 12

	// Simple field (dow)
	lengthInBits += 4

	return lengthInBits
}

func (m *_DateAndTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DateAndTimeParse(ctx context.Context, theBytes []byte) (DateAndTime, error) {
	return DateAndTimeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DateAndTimeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (DateAndTime, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (DateAndTime, error) {
		return DateAndTimeParseWithBuffer(ctx, readBuffer)
	}
}

func DateAndTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DateAndTime, error) {
	v, err := (&_DateAndTime{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_DateAndTime) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__dateAndTime DateAndTime, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DateAndTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DateAndTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	year, err := ReadSimpleField(ctx, "year", ReadUnsignedByte(readBuffer, uint8(8)), codegen.WithEncoding("BCD"))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'year' field"))
	}
	m.Year = year

	month, err := ReadSimpleField(ctx, "month", ReadUnsignedByte(readBuffer, uint8(8)), codegen.WithEncoding("BCD"))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'month' field"))
	}
	m.Month = month

	day, err := ReadSimpleField(ctx, "day", ReadUnsignedByte(readBuffer, uint8(8)), codegen.WithEncoding("BCD"))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'day' field"))
	}
	m.Day = day

	hour, err := ReadSimpleField(ctx, "hour", ReadUnsignedByte(readBuffer, uint8(8)), codegen.WithEncoding("BCD"))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hour' field"))
	}
	m.Hour = hour

	minutes, err := ReadSimpleField(ctx, "minutes", ReadUnsignedByte(readBuffer, uint8(8)), codegen.WithEncoding("BCD"))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minutes' field"))
	}
	m.Minutes = minutes

	seconds, err := ReadSimpleField(ctx, "seconds", ReadUnsignedByte(readBuffer, uint8(8)), codegen.WithEncoding("BCD"))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'seconds' field"))
	}
	m.Seconds = seconds

	msec, err := ReadSimpleField(ctx, "msec", ReadUnsignedShort(readBuffer, uint8(12)), codegen.WithEncoding("BCD"))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'msec' field"))
	}
	m.Msec = msec

	dow, err := ReadSimpleField(ctx, "dow", ReadUnsignedByte(readBuffer, uint8(4)), codegen.WithEncoding("BCD"))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dow' field"))
	}
	m.Dow = dow

	if closeErr := readBuffer.CloseContext("DateAndTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DateAndTime")
	}

	return m, nil
}

func (m *_DateAndTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DateAndTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("DateAndTime"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DateAndTime")
	}

	if err := WriteSimpleField[uint8](ctx, "year", m.GetYear(), WriteUnsignedByte(writeBuffer, 8), codegen.WithEncoding("BCD")); err != nil {
		return errors.Wrap(err, "Error serializing 'year' field")
	}

	if err := WriteSimpleField[uint8](ctx, "month", m.GetMonth(), WriteUnsignedByte(writeBuffer, 8), codegen.WithEncoding("BCD")); err != nil {
		return errors.Wrap(err, "Error serializing 'month' field")
	}

	if err := WriteSimpleField[uint8](ctx, "day", m.GetDay(), WriteUnsignedByte(writeBuffer, 8), codegen.WithEncoding("BCD")); err != nil {
		return errors.Wrap(err, "Error serializing 'day' field")
	}

	if err := WriteSimpleField[uint8](ctx, "hour", m.GetHour(), WriteUnsignedByte(writeBuffer, 8), codegen.WithEncoding("BCD")); err != nil {
		return errors.Wrap(err, "Error serializing 'hour' field")
	}

	if err := WriteSimpleField[uint8](ctx, "minutes", m.GetMinutes(), WriteUnsignedByte(writeBuffer, 8), codegen.WithEncoding("BCD")); err != nil {
		return errors.Wrap(err, "Error serializing 'minutes' field")
	}

	if err := WriteSimpleField[uint8](ctx, "seconds", m.GetSeconds(), WriteUnsignedByte(writeBuffer, 8), codegen.WithEncoding("BCD")); err != nil {
		return errors.Wrap(err, "Error serializing 'seconds' field")
	}

	if err := WriteSimpleField[uint16](ctx, "msec", m.GetMsec(), WriteUnsignedShort(writeBuffer, 12), codegen.WithEncoding("BCD")); err != nil {
		return errors.Wrap(err, "Error serializing 'msec' field")
	}

	if err := WriteSimpleField[uint8](ctx, "dow", m.GetDow(), WriteUnsignedByte(writeBuffer, 4), codegen.WithEncoding("BCD")); err != nil {
		return errors.Wrap(err, "Error serializing 'dow' field")
	}

	if popErr := writeBuffer.PopContext("DateAndTime"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DateAndTime")
	}
	return nil
}

func (m *_DateAndTime) IsDateAndTime() {}

func (m *_DateAndTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DateAndTime) deepCopy() *_DateAndTime {
	if m == nil {
		return nil
	}
	_DateAndTimeCopy := &_DateAndTime{
		m.Year,
		m.Month,
		m.Day,
		m.Hour,
		m.Minutes,
		m.Seconds,
		m.Msec,
		m.Dow,
	}
	return _DateAndTimeCopy
}

func (m *_DateAndTime) String() string {
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
