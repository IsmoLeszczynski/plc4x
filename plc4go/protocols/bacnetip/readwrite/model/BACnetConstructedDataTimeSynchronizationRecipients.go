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

// BACnetConstructedDataTimeSynchronizationRecipients is the corresponding interface of BACnetConstructedDataTimeSynchronizationRecipients
type BACnetConstructedDataTimeSynchronizationRecipients interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetTimeSynchronizationRecipients returns TimeSynchronizationRecipients (property field)
	GetTimeSynchronizationRecipients() []BACnetRecipient
}

// _BACnetConstructedDataTimeSynchronizationRecipients is the data-structure of this message
type _BACnetConstructedDataTimeSynchronizationRecipients struct {
	BACnetConstructedDataContract
	TimeSynchronizationRecipients []BACnetRecipient
}

var _ BACnetConstructedDataTimeSynchronizationRecipients = (*_BACnetConstructedDataTimeSynchronizationRecipients)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataTimeSynchronizationRecipients)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimeSynchronizationRecipients) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTimeSynchronizationRecipients) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TIME_SYNCHRONIZATION_RECIPIENTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimeSynchronizationRecipients) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimeSynchronizationRecipients) GetTimeSynchronizationRecipients() []BACnetRecipient {
	return m.TimeSynchronizationRecipients
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTimeSynchronizationRecipients factory function for _BACnetConstructedDataTimeSynchronizationRecipients
func NewBACnetConstructedDataTimeSynchronizationRecipients(timeSynchronizationRecipients []BACnetRecipient, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimeSynchronizationRecipients {
	_result := &_BACnetConstructedDataTimeSynchronizationRecipients{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		TimeSynchronizationRecipients: timeSynchronizationRecipients,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimeSynchronizationRecipients(structType any) BACnetConstructedDataTimeSynchronizationRecipients {
	if casted, ok := structType.(BACnetConstructedDataTimeSynchronizationRecipients); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimeSynchronizationRecipients); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimeSynchronizationRecipients) GetTypeName() string {
	return "BACnetConstructedDataTimeSynchronizationRecipients"
}

func (m *_BACnetConstructedDataTimeSynchronizationRecipients) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.TimeSynchronizationRecipients) > 0 {
		for _, element := range m.TimeSynchronizationRecipients {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataTimeSynchronizationRecipients) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataTimeSynchronizationRecipients) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataTimeSynchronizationRecipients BACnetConstructedDataTimeSynchronizationRecipients, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimeSynchronizationRecipients"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimeSynchronizationRecipients")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timeSynchronizationRecipients, err := ReadTerminatedArrayField[BACnetRecipient](ctx, "timeSynchronizationRecipients", ReadComplex[BACnetRecipient](BACnetRecipientParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeSynchronizationRecipients' field"))
	}
	m.TimeSynchronizationRecipients = timeSynchronizationRecipients

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimeSynchronizationRecipients"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimeSynchronizationRecipients")
	}

	return m, nil
}

func (m *_BACnetConstructedDataTimeSynchronizationRecipients) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTimeSynchronizationRecipients) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimeSynchronizationRecipients"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimeSynchronizationRecipients")
		}

		if err := WriteComplexTypeArrayField(ctx, "timeSynchronizationRecipients", m.GetTimeSynchronizationRecipients(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'timeSynchronizationRecipients' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimeSynchronizationRecipients"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimeSynchronizationRecipients")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimeSynchronizationRecipients) IsBACnetConstructedDataTimeSynchronizationRecipients() {
}

func (m *_BACnetConstructedDataTimeSynchronizationRecipients) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
