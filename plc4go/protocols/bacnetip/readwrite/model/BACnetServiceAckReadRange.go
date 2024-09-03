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

// BACnetServiceAckReadRange is the corresponding interface of BACnetServiceAckReadRange
type BACnetServiceAckReadRange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetServiceAck
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetPropertyArrayIndex returns PropertyArrayIndex (property field)
	GetPropertyArrayIndex() BACnetContextTagUnsignedInteger
	// GetResultFlags returns ResultFlags (property field)
	GetResultFlags() BACnetResultFlagsTagged
	// GetItemCount returns ItemCount (property field)
	GetItemCount() BACnetContextTagUnsignedInteger
	// GetItemData returns ItemData (property field)
	GetItemData() BACnetConstructedData
	// GetFirstSequenceNumber returns FirstSequenceNumber (property field)
	GetFirstSequenceNumber() BACnetContextTagUnsignedInteger
	// IsBACnetServiceAckReadRange is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetServiceAckReadRange()
}

// _BACnetServiceAckReadRange is the data-structure of this message
type _BACnetServiceAckReadRange struct {
	BACnetServiceAckContract
	ObjectIdentifier    BACnetContextTagObjectIdentifier
	PropertyIdentifier  BACnetPropertyIdentifierTagged
	PropertyArrayIndex  BACnetContextTagUnsignedInteger
	ResultFlags         BACnetResultFlagsTagged
	ItemCount           BACnetContextTagUnsignedInteger
	ItemData            BACnetConstructedData
	FirstSequenceNumber BACnetContextTagUnsignedInteger
}

var _ BACnetServiceAckReadRange = (*_BACnetServiceAckReadRange)(nil)
var _ BACnetServiceAckRequirements = (*_BACnetServiceAckReadRange)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckReadRange) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_READ_RANGE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckReadRange) GetParent() BACnetServiceAckContract {
	return m.BACnetServiceAckContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckReadRange) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetServiceAckReadRange) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_BACnetServiceAckReadRange) GetPropertyArrayIndex() BACnetContextTagUnsignedInteger {
	return m.PropertyArrayIndex
}

func (m *_BACnetServiceAckReadRange) GetResultFlags() BACnetResultFlagsTagged {
	return m.ResultFlags
}

func (m *_BACnetServiceAckReadRange) GetItemCount() BACnetContextTagUnsignedInteger {
	return m.ItemCount
}

func (m *_BACnetServiceAckReadRange) GetItemData() BACnetConstructedData {
	return m.ItemData
}

func (m *_BACnetServiceAckReadRange) GetFirstSequenceNumber() BACnetContextTagUnsignedInteger {
	return m.FirstSequenceNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckReadRange factory function for _BACnetServiceAckReadRange
func NewBACnetServiceAckReadRange(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged, propertyArrayIndex BACnetContextTagUnsignedInteger, resultFlags BACnetResultFlagsTagged, itemCount BACnetContextTagUnsignedInteger, itemData BACnetConstructedData, firstSequenceNumber BACnetContextTagUnsignedInteger, serviceAckLength uint32) *_BACnetServiceAckReadRange {
	if objectIdentifier == nil {
		panic("objectIdentifier of type BACnetContextTagObjectIdentifier for BACnetServiceAckReadRange must not be nil")
	}
	if propertyIdentifier == nil {
		panic("propertyIdentifier of type BACnetPropertyIdentifierTagged for BACnetServiceAckReadRange must not be nil")
	}
	if resultFlags == nil {
		panic("resultFlags of type BACnetResultFlagsTagged for BACnetServiceAckReadRange must not be nil")
	}
	if itemCount == nil {
		panic("itemCount of type BACnetContextTagUnsignedInteger for BACnetServiceAckReadRange must not be nil")
	}
	_result := &_BACnetServiceAckReadRange{
		BACnetServiceAckContract: NewBACnetServiceAck(serviceAckLength),
		ObjectIdentifier:         objectIdentifier,
		PropertyIdentifier:       propertyIdentifier,
		PropertyArrayIndex:       propertyArrayIndex,
		ResultFlags:              resultFlags,
		ItemCount:                itemCount,
		ItemData:                 itemData,
		FirstSequenceNumber:      firstSequenceNumber,
	}
	_result.BACnetServiceAckContract.(*_BACnetServiceAck)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckReadRange(structType any) BACnetServiceAckReadRange {
	if casted, ok := structType.(BACnetServiceAckReadRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckReadRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckReadRange) GetTypeName() string {
	return "BACnetServiceAckReadRange"
}

func (m *_BACnetServiceAckReadRange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetServiceAckContract.(*_BACnetServiceAck).getLengthInBits(ctx))

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits(ctx)

	// Optional Field (propertyArrayIndex)
	if m.PropertyArrayIndex != nil {
		lengthInBits += m.PropertyArrayIndex.GetLengthInBits(ctx)
	}

	// Simple field (resultFlags)
	lengthInBits += m.ResultFlags.GetLengthInBits(ctx)

	// Simple field (itemCount)
	lengthInBits += m.ItemCount.GetLengthInBits(ctx)

	// Optional Field (itemData)
	if m.ItemData != nil {
		lengthInBits += m.ItemData.GetLengthInBits(ctx)
	}

	// Optional Field (firstSequenceNumber)
	if m.FirstSequenceNumber != nil {
		lengthInBits += m.FirstSequenceNumber.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetServiceAckReadRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetServiceAckReadRange) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetServiceAck, serviceAckLength uint32) (__bACnetServiceAckReadRange BACnetServiceAckReadRange, err error) {
	m.BACnetServiceAckContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckReadRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckReadRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}
	m.ObjectIdentifier = objectIdentifier

	propertyIdentifier, err := ReadSimpleField[BACnetPropertyIdentifierTagged](ctx, "propertyIdentifier", ReadComplex[BACnetPropertyIdentifierTagged](BACnetPropertyIdentifierTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyIdentifier' field"))
	}
	m.PropertyIdentifier = propertyIdentifier

	var propertyArrayIndex BACnetContextTagUnsignedInteger
	_propertyArrayIndex, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "propertyArrayIndex", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyArrayIndex' field"))
	}
	if _propertyArrayIndex != nil {
		propertyArrayIndex = *_propertyArrayIndex
		m.PropertyArrayIndex = propertyArrayIndex
	}

	resultFlags, err := ReadSimpleField[BACnetResultFlagsTagged](ctx, "resultFlags", ReadComplex[BACnetResultFlagsTagged](BACnetResultFlagsTaggedParseWithBufferProducer((uint8)(uint8(3)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'resultFlags' field"))
	}
	m.ResultFlags = resultFlags

	itemCount, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "itemCount", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(4)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemCount' field"))
	}
	m.ItemCount = itemCount

	var itemData BACnetConstructedData
	_itemData, err := ReadOptionalField[BACnetConstructedData](ctx, "itemData", ReadComplex[BACnetConstructedData](BACnetConstructedDataParseWithBufferProducer[BACnetConstructedData]((uint8)(uint8(5)), (BACnetObjectType)(objectIdentifier.GetObjectType()), (BACnetPropertyIdentifier)(propertyIdentifier.GetValue()), (BACnetTagPayloadUnsignedInteger)((CastBACnetTagPayloadUnsignedInteger(utils.InlineIf(bool((propertyArrayIndex) != (nil)), func() any { return CastBACnetTagPayloadUnsignedInteger((propertyArrayIndex).GetPayload()) }, func() any { return CastBACnetTagPayloadUnsignedInteger(nil) }))))), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemData' field"))
	}
	if _itemData != nil {
		itemData = *_itemData
		m.ItemData = itemData
	}

	var firstSequenceNumber BACnetContextTagUnsignedInteger
	_firstSequenceNumber, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "firstSequenceNumber", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(6)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'firstSequenceNumber' field"))
	}
	if _firstSequenceNumber != nil {
		firstSequenceNumber = *_firstSequenceNumber
		m.FirstSequenceNumber = firstSequenceNumber
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckReadRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckReadRange")
	}

	return m, nil
}

func (m *_BACnetServiceAckReadRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckReadRange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckReadRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckReadRange")
		}

		if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", m.GetObjectIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectIdentifier' field")
		}

		if err := WriteSimpleField[BACnetPropertyIdentifierTagged](ctx, "propertyIdentifier", m.GetPropertyIdentifier(), WriteComplex[BACnetPropertyIdentifierTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'propertyIdentifier' field")
		}

		if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "propertyArrayIndex", GetRef(m.GetPropertyArrayIndex()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'propertyArrayIndex' field")
		}

		if err := WriteSimpleField[BACnetResultFlagsTagged](ctx, "resultFlags", m.GetResultFlags(), WriteComplex[BACnetResultFlagsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'resultFlags' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "itemCount", m.GetItemCount(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemCount' field")
		}

		if err := WriteOptionalField[BACnetConstructedData](ctx, "itemData", GetRef(m.GetItemData()), WriteComplex[BACnetConstructedData](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'itemData' field")
		}

		if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "firstSequenceNumber", GetRef(m.GetFirstSequenceNumber()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'firstSequenceNumber' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckReadRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckReadRange")
		}
		return nil
	}
	return m.BACnetServiceAckContract.(*_BACnetServiceAck).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetServiceAckReadRange) IsBACnetServiceAckReadRange() {}

func (m *_BACnetServiceAckReadRange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
