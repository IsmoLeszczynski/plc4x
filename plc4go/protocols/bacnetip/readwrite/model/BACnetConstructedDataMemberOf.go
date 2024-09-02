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

// BACnetConstructedDataMemberOf is the corresponding interface of BACnetConstructedDataMemberOf
type BACnetConstructedDataMemberOf interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetZones returns Zones (property field)
	GetZones() []BACnetDeviceObjectReference
}

// _BACnetConstructedDataMemberOf is the data-structure of this message
type _BACnetConstructedDataMemberOf struct {
	BACnetConstructedDataContract
	Zones []BACnetDeviceObjectReference
}

var _ BACnetConstructedDataMemberOf = (*_BACnetConstructedDataMemberOf)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataMemberOf)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMemberOf) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMemberOf) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MEMBER_OF
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMemberOf) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMemberOf) GetZones() []BACnetDeviceObjectReference {
	return m.Zones
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMemberOf factory function for _BACnetConstructedDataMemberOf
func NewBACnetConstructedDataMemberOf(zones []BACnetDeviceObjectReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMemberOf {
	_result := &_BACnetConstructedDataMemberOf{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Zones:                         zones,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMemberOf(structType any) BACnetConstructedDataMemberOf {
	if casted, ok := structType.(BACnetConstructedDataMemberOf); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMemberOf); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMemberOf) GetTypeName() string {
	return "BACnetConstructedDataMemberOf"
}

func (m *_BACnetConstructedDataMemberOf) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.Zones) > 0 {
		for _, element := range m.Zones {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataMemberOf) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMemberOf) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMemberOf BACnetConstructedDataMemberOf, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMemberOf"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMemberOf")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zones, err := ReadTerminatedArrayField[BACnetDeviceObjectReference](ctx, "zones", ReadComplex[BACnetDeviceObjectReference](BACnetDeviceObjectReferenceParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zones' field"))
	}
	m.Zones = zones

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMemberOf"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMemberOf")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMemberOf) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMemberOf) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMemberOf"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMemberOf")
		}

		if err := WriteComplexTypeArrayField(ctx, "zones", m.GetZones(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'zones' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMemberOf"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMemberOf")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMemberOf) IsBACnetConstructedDataMemberOf() {}

func (m *_BACnetConstructedDataMemberOf) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
