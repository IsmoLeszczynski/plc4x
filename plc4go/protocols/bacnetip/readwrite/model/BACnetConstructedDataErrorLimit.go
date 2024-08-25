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

// BACnetConstructedDataErrorLimit is the corresponding interface of BACnetConstructedDataErrorLimit
type BACnetConstructedDataErrorLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetErrorLimit returns ErrorLimit (property field)
	GetErrorLimit() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataErrorLimitExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataErrorLimit.
// This is useful for switch cases.
type BACnetConstructedDataErrorLimitExactly interface {
	BACnetConstructedDataErrorLimit
	isBACnetConstructedDataErrorLimit() bool
}

// _BACnetConstructedDataErrorLimit is the data-structure of this message
type _BACnetConstructedDataErrorLimit struct {
	*_BACnetConstructedData
	ErrorLimit BACnetApplicationTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataErrorLimit) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataErrorLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ERROR_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataErrorLimit) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataErrorLimit) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataErrorLimit) GetErrorLimit() BACnetApplicationTagReal {
	return m.ErrorLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataErrorLimit) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetErrorLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataErrorLimit factory function for _BACnetConstructedDataErrorLimit
func NewBACnetConstructedDataErrorLimit(errorLimit BACnetApplicationTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataErrorLimit {
	_result := &_BACnetConstructedDataErrorLimit{
		ErrorLimit:             errorLimit,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataErrorLimit(structType any) BACnetConstructedDataErrorLimit {
	if casted, ok := structType.(BACnetConstructedDataErrorLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataErrorLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataErrorLimit) GetTypeName() string {
	return "BACnetConstructedDataErrorLimit"
}

func (m *_BACnetConstructedDataErrorLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (errorLimit)
	lengthInBits += m.ErrorLimit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataErrorLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataErrorLimitParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataErrorLimit, error) {
	return BACnetConstructedDataErrorLimitParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataErrorLimitParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataErrorLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataErrorLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataErrorLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (errorLimit)
	if pullErr := readBuffer.PullContext("errorLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for errorLimit")
	}
	_errorLimit, _errorLimitErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _errorLimitErr != nil {
		return nil, errors.Wrap(_errorLimitErr, "Error parsing 'errorLimit' field of BACnetConstructedDataErrorLimit")
	}
	errorLimit := _errorLimit.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("errorLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for errorLimit")
	}

	// Virtual field
	_actualValue := errorLimit
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataErrorLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataErrorLimit")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataErrorLimit{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ErrorLimit: errorLimit,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataErrorLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataErrorLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataErrorLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataErrorLimit")
		}

		// Simple Field (errorLimit)
		if pushErr := writeBuffer.PushContext("errorLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for errorLimit")
		}
		_errorLimitErr := writeBuffer.WriteSerializable(ctx, m.GetErrorLimit())
		if popErr := writeBuffer.PopContext("errorLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for errorLimit")
		}
		if _errorLimitErr != nil {
			return errors.Wrap(_errorLimitErr, "Error serializing 'errorLimit' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataErrorLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataErrorLimit")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataErrorLimit) isBACnetConstructedDataErrorLimit() bool {
	return true
}

func (m *_BACnetConstructedDataErrorLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
