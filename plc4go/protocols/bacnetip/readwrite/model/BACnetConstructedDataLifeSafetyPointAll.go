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

// BACnetConstructedDataLifeSafetyPointAll is the corresponding interface of BACnetConstructedDataLifeSafetyPointAll
type BACnetConstructedDataLifeSafetyPointAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
}

// BACnetConstructedDataLifeSafetyPointAllExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLifeSafetyPointAll.
// This is useful for switch cases.
type BACnetConstructedDataLifeSafetyPointAllExactly interface {
	BACnetConstructedDataLifeSafetyPointAll
	isBACnetConstructedDataLifeSafetyPointAll() bool
}

// _BACnetConstructedDataLifeSafetyPointAll is the data-structure of this message
type _BACnetConstructedDataLifeSafetyPointAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataLifeSafetyPointAll = (*_BACnetConstructedDataLifeSafetyPointAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLifeSafetyPointAll)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLifeSafetyPointAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LIFE_SAFETY_POINT
}

func (m *_BACnetConstructedDataLifeSafetyPointAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLifeSafetyPointAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// NewBACnetConstructedDataLifeSafetyPointAll factory function for _BACnetConstructedDataLifeSafetyPointAll
func NewBACnetConstructedDataLifeSafetyPointAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLifeSafetyPointAll {
	_result := &_BACnetConstructedDataLifeSafetyPointAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLifeSafetyPointAll(structType any) BACnetConstructedDataLifeSafetyPointAll {
	if casted, ok := structType.(BACnetConstructedDataLifeSafetyPointAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLifeSafetyPointAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLifeSafetyPointAll) GetTypeName() string {
	return "BACnetConstructedDataLifeSafetyPointAll"
}

func (m *_BACnetConstructedDataLifeSafetyPointAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataLifeSafetyPointAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLifeSafetyPointAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLifeSafetyPointAll BACnetConstructedDataLifeSafetyPointAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLifeSafetyPointAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLifeSafetyPointAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLifeSafetyPointAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLifeSafetyPointAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLifeSafetyPointAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLifeSafetyPointAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLifeSafetyPointAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLifeSafetyPointAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLifeSafetyPointAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLifeSafetyPointAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLifeSafetyPointAll) isBACnetConstructedDataLifeSafetyPointAll() bool {
	return true
}

func (m *_BACnetConstructedDataLifeSafetyPointAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
