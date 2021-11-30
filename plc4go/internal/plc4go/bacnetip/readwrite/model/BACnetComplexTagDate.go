/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetComplexTagDate struct {
	YearMinus1900          int8
	Month                  int8
	DayOfMonth             int8
	DayOfWeek              int8
	Wildcard               int8
	YearIsWildcard         bool
	MonthIsWildcard        bool
	OddMonthWildcard       bool
	EvenMonthWildcard      bool
	DayOfMonthIsWildcard   bool
	LastDayOfMonthWildcard bool
	OddDayOfMonthWildcard  bool
	EvenDayOfMonthWildcard bool
	DayOfWeekIsWildcard    bool
	Parent                 *BACnetComplexTag
}

// The corresponding interface
type IBACnetComplexTagDate interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetComplexTagDate) DataType() BACnetDataType {
	return BACnetDataType_DATE
}

func (m *BACnetComplexTagDate) InitializeParent(parent *BACnetComplexTag, tagNumber uint8, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32, actualTagNumber uint8, actualLength uint32) {
	m.Parent.TagNumber = tagNumber
	m.Parent.TagClass = tagClass
	m.Parent.LengthValueType = lengthValueType
	m.Parent.ExtTagNumber = extTagNumber
	m.Parent.ExtLength = extLength
	m.Parent.ExtExtLength = extExtLength
	m.Parent.ExtExtExtLength = extExtExtLength
}

func NewBACnetComplexTagDate(yearMinus1900 int8, month int8, dayOfMonth int8, dayOfWeek int8, tagNumber uint8, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32) *BACnetComplexTag {
	child := &BACnetComplexTagDate{
		YearMinus1900: yearMinus1900,
		Month:         month,
		DayOfMonth:    dayOfMonth,
		DayOfWeek:     dayOfWeek,
		Parent:        NewBACnetComplexTag(tagNumber, tagClass, lengthValueType, extTagNumber, extLength, extExtLength, extExtExtLength),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetComplexTagDate(structType interface{}) *BACnetComplexTagDate {
	castFunc := func(typ interface{}) *BACnetComplexTagDate {
		if casted, ok := typ.(BACnetComplexTagDate); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetComplexTagDate); ok {
			return casted
		}
		if casted, ok := typ.(BACnetComplexTag); ok {
			return CastBACnetComplexTagDate(casted.Child)
		}
		if casted, ok := typ.(*BACnetComplexTag); ok {
			return CastBACnetComplexTagDate(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetComplexTagDate) GetTypeName() string {
	return "BACnetComplexTagDate"
}

func (m *BACnetComplexTagDate) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetComplexTagDate) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Simple field (yearMinus1900)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (month)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (dayOfMonth)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (dayOfWeek)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetComplexTagDate) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetComplexTagDateParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType) (*BACnetComplexTag, error) {
	if pullErr := readBuffer.PullContext("BACnetComplexTagDate"); pullErr != nil {
		return nil, pullErr
	}

	// Virtual field
	_wildcard := 0xFF
	wildcard := int8(_wildcard)

	// Simple Field (yearMinus1900)
	yearMinus1900, _yearMinus1900Err := readBuffer.ReadInt8("yearMinus1900", 8)
	if _yearMinus1900Err != nil {
		return nil, errors.Wrap(_yearMinus1900Err, "Error parsing 'yearMinus1900' field")
	}

	// Virtual field
	_yearIsWildcard := bool((yearMinus1900) == (wildcard))
	yearIsWildcard := bool(_yearIsWildcard)

	// Simple Field (month)
	month, _monthErr := readBuffer.ReadInt8("month", 8)
	if _monthErr != nil {
		return nil, errors.Wrap(_monthErr, "Error parsing 'month' field")
	}

	// Virtual field
	_monthIsWildcard := bool((month) == (wildcard))
	monthIsWildcard := bool(_monthIsWildcard)

	// Virtual field
	_oddMonthWildcard := bool((month) == (13))
	oddMonthWildcard := bool(_oddMonthWildcard)

	// Virtual field
	_evenMonthWildcard := bool((month) == (14))
	evenMonthWildcard := bool(_evenMonthWildcard)

	// Simple Field (dayOfMonth)
	dayOfMonth, _dayOfMonthErr := readBuffer.ReadInt8("dayOfMonth", 8)
	if _dayOfMonthErr != nil {
		return nil, errors.Wrap(_dayOfMonthErr, "Error parsing 'dayOfMonth' field")
	}

	// Virtual field
	_dayOfMonthIsWildcard := bool((dayOfMonth) == (wildcard))
	dayOfMonthIsWildcard := bool(_dayOfMonthIsWildcard)

	// Virtual field
	_lastDayOfMonthWildcard := bool((dayOfMonth) == (32))
	lastDayOfMonthWildcard := bool(_lastDayOfMonthWildcard)

	// Virtual field
	_oddDayOfMonthWildcard := bool((dayOfMonth) == (33))
	oddDayOfMonthWildcard := bool(_oddDayOfMonthWildcard)

	// Virtual field
	_evenDayOfMonthWildcard := bool((dayOfMonth) == (34))
	evenDayOfMonthWildcard := bool(_evenDayOfMonthWildcard)

	// Simple Field (dayOfWeek)
	dayOfWeek, _dayOfWeekErr := readBuffer.ReadInt8("dayOfWeek", 8)
	if _dayOfWeekErr != nil {
		return nil, errors.Wrap(_dayOfWeekErr, "Error parsing 'dayOfWeek' field")
	}

	// Virtual field
	_dayOfWeekIsWildcard := bool((dayOfWeek) == (wildcard))
	dayOfWeekIsWildcard := bool(_dayOfWeekIsWildcard)

	if closeErr := readBuffer.CloseContext("BACnetComplexTagDate"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetComplexTagDate{
		YearMinus1900:          yearMinus1900,
		Month:                  month,
		DayOfMonth:             dayOfMonth,
		DayOfWeek:              dayOfWeek,
		Wildcard:               wildcard,
		YearIsWildcard:         yearIsWildcard,
		MonthIsWildcard:        monthIsWildcard,
		OddMonthWildcard:       oddMonthWildcard,
		EvenMonthWildcard:      evenMonthWildcard,
		DayOfMonthIsWildcard:   dayOfMonthIsWildcard,
		LastDayOfMonthWildcard: lastDayOfMonthWildcard,
		OddDayOfMonthWildcard:  oddDayOfMonthWildcard,
		EvenDayOfMonthWildcard: evenDayOfMonthWildcard,
		DayOfWeekIsWildcard:    dayOfWeekIsWildcard,
		Parent:                 &BACnetComplexTag{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetComplexTagDate) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetComplexTagDate"); pushErr != nil {
			return pushErr
		}

		// Simple Field (yearMinus1900)
		yearMinus1900 := int8(m.YearMinus1900)
		_yearMinus1900Err := writeBuffer.WriteInt8("yearMinus1900", 8, (yearMinus1900))
		if _yearMinus1900Err != nil {
			return errors.Wrap(_yearMinus1900Err, "Error serializing 'yearMinus1900' field")
		}

		// Simple Field (month)
		month := int8(m.Month)
		_monthErr := writeBuffer.WriteInt8("month", 8, (month))
		if _monthErr != nil {
			return errors.Wrap(_monthErr, "Error serializing 'month' field")
		}

		// Simple Field (dayOfMonth)
		dayOfMonth := int8(m.DayOfMonth)
		_dayOfMonthErr := writeBuffer.WriteInt8("dayOfMonth", 8, (dayOfMonth))
		if _dayOfMonthErr != nil {
			return errors.Wrap(_dayOfMonthErr, "Error serializing 'dayOfMonth' field")
		}

		// Simple Field (dayOfWeek)
		dayOfWeek := int8(m.DayOfWeek)
		_dayOfWeekErr := writeBuffer.WriteInt8("dayOfWeek", 8, (dayOfWeek))
		if _dayOfWeekErr != nil {
			return errors.Wrap(_dayOfWeekErr, "Error serializing 'dayOfWeek' field")
		}

		if popErr := writeBuffer.PopContext("BACnetComplexTagDate"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetComplexTagDate) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
