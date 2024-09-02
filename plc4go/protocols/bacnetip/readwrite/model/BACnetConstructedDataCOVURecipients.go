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

// BACnetConstructedDataCOVURecipients is the corresponding interface of BACnetConstructedDataCOVURecipients
type BACnetConstructedDataCOVURecipients interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetCovuRecipients returns CovuRecipients (property field)
	GetCovuRecipients() []BACnetRecipient
}

// _BACnetConstructedDataCOVURecipients is the data-structure of this message
type _BACnetConstructedDataCOVURecipients struct {
	BACnetConstructedDataContract
	CovuRecipients []BACnetRecipient
}

var _ BACnetConstructedDataCOVURecipients = (*_BACnetConstructedDataCOVURecipients)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataCOVURecipients)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCOVURecipients) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCOVURecipients) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_COVU_RECIPIENTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCOVURecipients) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCOVURecipients) GetCovuRecipients() []BACnetRecipient {
	return m.CovuRecipients
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCOVURecipients factory function for _BACnetConstructedDataCOVURecipients
func NewBACnetConstructedDataCOVURecipients(covuRecipients []BACnetRecipient, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCOVURecipients {
	_result := &_BACnetConstructedDataCOVURecipients{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		CovuRecipients:                covuRecipients,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCOVURecipients(structType any) BACnetConstructedDataCOVURecipients {
	if casted, ok := structType.(BACnetConstructedDataCOVURecipients); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCOVURecipients); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCOVURecipients) GetTypeName() string {
	return "BACnetConstructedDataCOVURecipients"
}

func (m *_BACnetConstructedDataCOVURecipients) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.CovuRecipients) > 0 {
		for _, element := range m.CovuRecipients {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataCOVURecipients) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataCOVURecipients) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataCOVURecipients BACnetConstructedDataCOVURecipients, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCOVURecipients"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCOVURecipients")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	covuRecipients, err := ReadTerminatedArrayField[BACnetRecipient](ctx, "covuRecipients", ReadComplex[BACnetRecipient](BACnetRecipientParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'covuRecipients' field"))
	}
	m.CovuRecipients = covuRecipients

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCOVURecipients"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCOVURecipients")
	}

	return m, nil
}

func (m *_BACnetConstructedDataCOVURecipients) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCOVURecipients) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCOVURecipients"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCOVURecipients")
		}

		if err := WriteComplexTypeArrayField(ctx, "covuRecipients", m.GetCovuRecipients(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'covuRecipients' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCOVURecipients"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCOVURecipients")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCOVURecipients) IsBACnetConstructedDataCOVURecipients() {}

func (m *_BACnetConstructedDataCOVURecipients) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
