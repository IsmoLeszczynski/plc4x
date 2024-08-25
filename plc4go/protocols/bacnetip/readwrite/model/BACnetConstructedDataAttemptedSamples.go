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

// BACnetConstructedDataAttemptedSamples is the corresponding interface of BACnetConstructedDataAttemptedSamples
type BACnetConstructedDataAttemptedSamples interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAttemptedSamples returns AttemptedSamples (property field)
	GetAttemptedSamples() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataAttemptedSamplesExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAttemptedSamples.
// This is useful for switch cases.
type BACnetConstructedDataAttemptedSamplesExactly interface {
	BACnetConstructedDataAttemptedSamples
	isBACnetConstructedDataAttemptedSamples() bool
}

// _BACnetConstructedDataAttemptedSamples is the data-structure of this message
type _BACnetConstructedDataAttemptedSamples struct {
	*_BACnetConstructedData
	AttemptedSamples BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAttemptedSamples) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAttemptedSamples) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ATTEMPTED_SAMPLES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAttemptedSamples) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAttemptedSamples) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAttemptedSamples) GetAttemptedSamples() BACnetApplicationTagUnsignedInteger {
	return m.AttemptedSamples
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAttemptedSamples) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetAttemptedSamples())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAttemptedSamples factory function for _BACnetConstructedDataAttemptedSamples
func NewBACnetConstructedDataAttemptedSamples(attemptedSamples BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAttemptedSamples {
	_result := &_BACnetConstructedDataAttemptedSamples{
		AttemptedSamples:       attemptedSamples,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAttemptedSamples(structType any) BACnetConstructedDataAttemptedSamples {
	if casted, ok := structType.(BACnetConstructedDataAttemptedSamples); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAttemptedSamples); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAttemptedSamples) GetTypeName() string {
	return "BACnetConstructedDataAttemptedSamples"
}

func (m *_BACnetConstructedDataAttemptedSamples) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (attemptedSamples)
	lengthInBits += m.AttemptedSamples.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAttemptedSamples) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataAttemptedSamplesParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAttemptedSamples, error) {
	return BACnetConstructedDataAttemptedSamplesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataAttemptedSamplesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAttemptedSamples, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAttemptedSamples"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAttemptedSamples")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (attemptedSamples)
	if pullErr := readBuffer.PullContext("attemptedSamples"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for attemptedSamples")
	}
	_attemptedSamples, _attemptedSamplesErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _attemptedSamplesErr != nil {
		return nil, errors.Wrap(_attemptedSamplesErr, "Error parsing 'attemptedSamples' field of BACnetConstructedDataAttemptedSamples")
	}
	attemptedSamples := _attemptedSamples.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("attemptedSamples"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for attemptedSamples")
	}

	// Virtual field
	_actualValue := attemptedSamples
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAttemptedSamples"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAttemptedSamples")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAttemptedSamples{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		AttemptedSamples: attemptedSamples,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAttemptedSamples) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAttemptedSamples) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAttemptedSamples"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAttemptedSamples")
		}

		// Simple Field (attemptedSamples)
		if pushErr := writeBuffer.PushContext("attemptedSamples"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for attemptedSamples")
		}
		_attemptedSamplesErr := writeBuffer.WriteSerializable(ctx, m.GetAttemptedSamples())
		if popErr := writeBuffer.PopContext("attemptedSamples"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for attemptedSamples")
		}
		if _attemptedSamplesErr != nil {
			return errors.Wrap(_attemptedSamplesErr, "Error serializing 'attemptedSamples' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAttemptedSamples"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAttemptedSamples")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAttemptedSamples) isBACnetConstructedDataAttemptedSamples() bool {
	return true
}

func (m *_BACnetConstructedDataAttemptedSamples) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
