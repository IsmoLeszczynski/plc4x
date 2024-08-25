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

// BACnetConstructedDataLastCredentialAddedTime is the corresponding interface of BACnetConstructedDataLastCredentialAddedTime
type BACnetConstructedDataLastCredentialAddedTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLastCredentialAddedTime returns LastCredentialAddedTime (property field)
	GetLastCredentialAddedTime() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
}

// BACnetConstructedDataLastCredentialAddedTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLastCredentialAddedTime.
// This is useful for switch cases.
type BACnetConstructedDataLastCredentialAddedTimeExactly interface {
	BACnetConstructedDataLastCredentialAddedTime
	isBACnetConstructedDataLastCredentialAddedTime() bool
}

// _BACnetConstructedDataLastCredentialAddedTime is the data-structure of this message
type _BACnetConstructedDataLastCredentialAddedTime struct {
	*_BACnetConstructedData
	LastCredentialAddedTime BACnetDateTime
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLastCredentialAddedTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLastCredentialAddedTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_CREDENTIAL_ADDED_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLastCredentialAddedTime) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLastCredentialAddedTime) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLastCredentialAddedTime) GetLastCredentialAddedTime() BACnetDateTime {
	return m.LastCredentialAddedTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLastCredentialAddedTime) GetActualValue() BACnetDateTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDateTime(m.GetLastCredentialAddedTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLastCredentialAddedTime factory function for _BACnetConstructedDataLastCredentialAddedTime
func NewBACnetConstructedDataLastCredentialAddedTime(lastCredentialAddedTime BACnetDateTime, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLastCredentialAddedTime {
	_result := &_BACnetConstructedDataLastCredentialAddedTime{
		LastCredentialAddedTime: lastCredentialAddedTime,
		_BACnetConstructedData:  NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLastCredentialAddedTime(structType any) BACnetConstructedDataLastCredentialAddedTime {
	if casted, ok := structType.(BACnetConstructedDataLastCredentialAddedTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastCredentialAddedTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLastCredentialAddedTime) GetTypeName() string {
	return "BACnetConstructedDataLastCredentialAddedTime"
}

func (m *_BACnetConstructedDataLastCredentialAddedTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (lastCredentialAddedTime)
	lengthInBits += m.LastCredentialAddedTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLastCredentialAddedTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataLastCredentialAddedTimeParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLastCredentialAddedTime, error) {
	return BACnetConstructedDataLastCredentialAddedTimeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataLastCredentialAddedTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLastCredentialAddedTime, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastCredentialAddedTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastCredentialAddedTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lastCredentialAddedTime)
	if pullErr := readBuffer.PullContext("lastCredentialAddedTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lastCredentialAddedTime")
	}
	_lastCredentialAddedTime, _lastCredentialAddedTimeErr := BACnetDateTimeParseWithBuffer(ctx, readBuffer)
	if _lastCredentialAddedTimeErr != nil {
		return nil, errors.Wrap(_lastCredentialAddedTimeErr, "Error parsing 'lastCredentialAddedTime' field of BACnetConstructedDataLastCredentialAddedTime")
	}
	lastCredentialAddedTime := _lastCredentialAddedTime.(BACnetDateTime)
	if closeErr := readBuffer.CloseContext("lastCredentialAddedTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lastCredentialAddedTime")
	}

	// Virtual field
	_actualValue := lastCredentialAddedTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastCredentialAddedTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastCredentialAddedTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLastCredentialAddedTime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		LastCredentialAddedTime: lastCredentialAddedTime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLastCredentialAddedTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLastCredentialAddedTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastCredentialAddedTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastCredentialAddedTime")
		}

		// Simple Field (lastCredentialAddedTime)
		if pushErr := writeBuffer.PushContext("lastCredentialAddedTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lastCredentialAddedTime")
		}
		_lastCredentialAddedTimeErr := writeBuffer.WriteSerializable(ctx, m.GetLastCredentialAddedTime())
		if popErr := writeBuffer.PopContext("lastCredentialAddedTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lastCredentialAddedTime")
		}
		if _lastCredentialAddedTimeErr != nil {
			return errors.Wrap(_lastCredentialAddedTimeErr, "Error serializing 'lastCredentialAddedTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastCredentialAddedTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastCredentialAddedTime")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLastCredentialAddedTime) isBACnetConstructedDataLastCredentialAddedTime() bool {
	return true
}

func (m *_BACnetConstructedDataLastCredentialAddedTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
