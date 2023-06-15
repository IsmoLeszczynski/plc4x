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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataUserType is the corresponding interface of BACnetConstructedDataUserType
type BACnetConstructedDataUserType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetUserType returns UserType (property field)
	GetUserType() BACnetAccessUserTypeTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAccessUserTypeTagged
}

// BACnetConstructedDataUserTypeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataUserType.
// This is useful for switch cases.
type BACnetConstructedDataUserTypeExactly interface {
	BACnetConstructedDataUserType
	isBACnetConstructedDataUserType() bool
}

// _BACnetConstructedDataUserType is the data-structure of this message
type _BACnetConstructedDataUserType struct {
	*_BACnetConstructedData
	UserType BACnetAccessUserTypeTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataUserType) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataUserType) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_USER_TYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataUserType) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataUserType) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataUserType) GetUserType() BACnetAccessUserTypeTagged {
	return m.UserType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataUserType) GetActualValue() BACnetAccessUserTypeTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetAccessUserTypeTagged(m.GetUserType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataUserType factory function for _BACnetConstructedDataUserType
func NewBACnetConstructedDataUserType(userType BACnetAccessUserTypeTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataUserType {
	_result := &_BACnetConstructedDataUserType{
		UserType:               userType,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataUserType(structType any) BACnetConstructedDataUserType {
	if casted, ok := structType.(BACnetConstructedDataUserType); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataUserType); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataUserType) GetTypeName() string {
	return "BACnetConstructedDataUserType"
}

func (m *_BACnetConstructedDataUserType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (userType)
	lengthInBits += m.UserType.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataUserType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataUserTypeParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataUserType, error) {
	return BACnetConstructedDataUserTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataUserTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataUserType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataUserType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataUserType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (userType)
	if pullErr := readBuffer.PullContext("userType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for userType")
	}
	_userType, _userTypeErr := BACnetAccessUserTypeTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _userTypeErr != nil {
		return nil, errors.Wrap(_userTypeErr, "Error parsing 'userType' field of BACnetConstructedDataUserType")
	}
	userType := _userType.(BACnetAccessUserTypeTagged)
	if closeErr := readBuffer.CloseContext("userType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for userType")
	}

	// Virtual field
	_actualValue := userType
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataUserType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataUserType")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataUserType{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		UserType: userType,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataUserType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataUserType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataUserType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataUserType")
		}

		// Simple Field (userType)
		if pushErr := writeBuffer.PushContext("userType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for userType")
		}
		_userTypeErr := writeBuffer.WriteSerializable(ctx, m.GetUserType())
		if popErr := writeBuffer.PopContext("userType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for userType")
		}
		if _userTypeErr != nil {
			return errors.Wrap(_userTypeErr, "Error serializing 'userType' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataUserType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataUserType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataUserType) isBACnetConstructedDataUserType() bool {
	return true
}

func (m *_BACnetConstructedDataUserType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
