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

// BACnetConstructedDataEscalatorFaultSignals is the corresponding interface of BACnetConstructedDataEscalatorFaultSignals
type BACnetConstructedDataEscalatorFaultSignals interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFaultSignals returns FaultSignals (property field)
	GetFaultSignals() []BACnetEscalatorFaultTagged
}

// BACnetConstructedDataEscalatorFaultSignalsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataEscalatorFaultSignals.
// This is useful for switch cases.
type BACnetConstructedDataEscalatorFaultSignalsExactly interface {
	BACnetConstructedDataEscalatorFaultSignals
	isBACnetConstructedDataEscalatorFaultSignals() bool
}

// _BACnetConstructedDataEscalatorFaultSignals is the data-structure of this message
type _BACnetConstructedDataEscalatorFaultSignals struct {
	BACnetConstructedDataContract
	FaultSignals []BACnetEscalatorFaultTagged
}

var _ BACnetConstructedDataEscalatorFaultSignals = (*_BACnetConstructedDataEscalatorFaultSignals)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataEscalatorFaultSignals)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEscalatorFaultSignals) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ESCALATOR
}

func (m *_BACnetConstructedDataEscalatorFaultSignals) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_SIGNALS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEscalatorFaultSignals) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEscalatorFaultSignals) GetFaultSignals() []BACnetEscalatorFaultTagged {
	return m.FaultSignals
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEscalatorFaultSignals factory function for _BACnetConstructedDataEscalatorFaultSignals
func NewBACnetConstructedDataEscalatorFaultSignals(faultSignals []BACnetEscalatorFaultTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEscalatorFaultSignals {
	_result := &_BACnetConstructedDataEscalatorFaultSignals{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FaultSignals:                  faultSignals,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEscalatorFaultSignals(structType any) BACnetConstructedDataEscalatorFaultSignals {
	if casted, ok := structType.(BACnetConstructedDataEscalatorFaultSignals); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEscalatorFaultSignals); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEscalatorFaultSignals) GetTypeName() string {
	return "BACnetConstructedDataEscalatorFaultSignals"
}

func (m *_BACnetConstructedDataEscalatorFaultSignals) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.FaultSignals) > 0 {
		for _, element := range m.FaultSignals {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataEscalatorFaultSignals) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataEscalatorFaultSignals) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataEscalatorFaultSignals BACnetConstructedDataEscalatorFaultSignals, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEscalatorFaultSignals"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEscalatorFaultSignals")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	faultSignals, err := ReadTerminatedArrayField[BACnetEscalatorFaultTagged](ctx, "faultSignals", ReadComplex[BACnetEscalatorFaultTagged](BACnetEscalatorFaultTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'faultSignals' field"))
	}
	m.FaultSignals = faultSignals

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEscalatorFaultSignals"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEscalatorFaultSignals")
	}

	return m, nil
}

func (m *_BACnetConstructedDataEscalatorFaultSignals) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEscalatorFaultSignals) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEscalatorFaultSignals"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEscalatorFaultSignals")
		}

		if err := WriteComplexTypeArrayField(ctx, "faultSignals", m.GetFaultSignals(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'faultSignals' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEscalatorFaultSignals"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEscalatorFaultSignals")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEscalatorFaultSignals) isBACnetConstructedDataEscalatorFaultSignals() bool {
	return true
}

func (m *_BACnetConstructedDataEscalatorFaultSignals) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
