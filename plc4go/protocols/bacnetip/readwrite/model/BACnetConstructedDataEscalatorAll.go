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

// BACnetConstructedDataEscalatorAll is the corresponding interface of BACnetConstructedDataEscalatorAll
type BACnetConstructedDataEscalatorAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
}

// _BACnetConstructedDataEscalatorAll is the data-structure of this message
type _BACnetConstructedDataEscalatorAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataEscalatorAll = (*_BACnetConstructedDataEscalatorAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataEscalatorAll)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEscalatorAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ESCALATOR
}

func (m *_BACnetConstructedDataEscalatorAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEscalatorAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// NewBACnetConstructedDataEscalatorAll factory function for _BACnetConstructedDataEscalatorAll
func NewBACnetConstructedDataEscalatorAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEscalatorAll {
	_result := &_BACnetConstructedDataEscalatorAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEscalatorAll(structType any) BACnetConstructedDataEscalatorAll {
	if casted, ok := structType.(BACnetConstructedDataEscalatorAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEscalatorAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEscalatorAll) GetTypeName() string {
	return "BACnetConstructedDataEscalatorAll"
}

func (m *_BACnetConstructedDataEscalatorAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataEscalatorAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataEscalatorAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataEscalatorAll BACnetConstructedDataEscalatorAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEscalatorAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEscalatorAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEscalatorAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEscalatorAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataEscalatorAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEscalatorAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEscalatorAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEscalatorAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEscalatorAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEscalatorAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEscalatorAll) IsBACnetConstructedDataEscalatorAll() {}

func (m *_BACnetConstructedDataEscalatorAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
