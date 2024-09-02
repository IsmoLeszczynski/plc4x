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

// BACnetConstructedDataSubscribedRecipients is the corresponding interface of BACnetConstructedDataSubscribedRecipients
type BACnetConstructedDataSubscribedRecipients interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetSubscribedRecipients returns SubscribedRecipients (property field)
	GetSubscribedRecipients() []BACnetEventNotificationSubscription
}

// _BACnetConstructedDataSubscribedRecipients is the data-structure of this message
type _BACnetConstructedDataSubscribedRecipients struct {
	BACnetConstructedDataContract
	SubscribedRecipients []BACnetEventNotificationSubscription
}

var _ BACnetConstructedDataSubscribedRecipients = (*_BACnetConstructedDataSubscribedRecipients)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataSubscribedRecipients)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSubscribedRecipients) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSubscribedRecipients) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SUBSCRIBED_RECIPIENTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSubscribedRecipients) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSubscribedRecipients) GetSubscribedRecipients() []BACnetEventNotificationSubscription {
	return m.SubscribedRecipients
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSubscribedRecipients factory function for _BACnetConstructedDataSubscribedRecipients
func NewBACnetConstructedDataSubscribedRecipients(subscribedRecipients []BACnetEventNotificationSubscription, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSubscribedRecipients {
	_result := &_BACnetConstructedDataSubscribedRecipients{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		SubscribedRecipients:          subscribedRecipients,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSubscribedRecipients(structType any) BACnetConstructedDataSubscribedRecipients {
	if casted, ok := structType.(BACnetConstructedDataSubscribedRecipients); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSubscribedRecipients); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSubscribedRecipients) GetTypeName() string {
	return "BACnetConstructedDataSubscribedRecipients"
}

func (m *_BACnetConstructedDataSubscribedRecipients) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.SubscribedRecipients) > 0 {
		for _, element := range m.SubscribedRecipients {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataSubscribedRecipients) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataSubscribedRecipients) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataSubscribedRecipients BACnetConstructedDataSubscribedRecipients, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSubscribedRecipients"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSubscribedRecipients")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	subscribedRecipients, err := ReadTerminatedArrayField[BACnetEventNotificationSubscription](ctx, "subscribedRecipients", ReadComplex[BACnetEventNotificationSubscription](BACnetEventNotificationSubscriptionParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscribedRecipients' field"))
	}
	m.SubscribedRecipients = subscribedRecipients

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSubscribedRecipients"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSubscribedRecipients")
	}

	return m, nil
}

func (m *_BACnetConstructedDataSubscribedRecipients) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataSubscribedRecipients) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSubscribedRecipients"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSubscribedRecipients")
		}

		if err := WriteComplexTypeArrayField(ctx, "subscribedRecipients", m.GetSubscribedRecipients(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'subscribedRecipients' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSubscribedRecipients"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSubscribedRecipients")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSubscribedRecipients) IsBACnetConstructedDataSubscribedRecipients() {}

func (m *_BACnetConstructedDataSubscribedRecipients) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
