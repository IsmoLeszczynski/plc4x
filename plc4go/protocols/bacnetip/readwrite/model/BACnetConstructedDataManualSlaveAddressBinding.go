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

// BACnetConstructedDataManualSlaveAddressBinding is the corresponding interface of BACnetConstructedDataManualSlaveAddressBinding
type BACnetConstructedDataManualSlaveAddressBinding interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetManualSlaveAddressBinding returns ManualSlaveAddressBinding (property field)
	GetManualSlaveAddressBinding() []BACnetAddressBinding
}

// BACnetConstructedDataManualSlaveAddressBindingExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataManualSlaveAddressBinding.
// This is useful for switch cases.
type BACnetConstructedDataManualSlaveAddressBindingExactly interface {
	BACnetConstructedDataManualSlaveAddressBinding
	isBACnetConstructedDataManualSlaveAddressBinding() bool
}

// _BACnetConstructedDataManualSlaveAddressBinding is the data-structure of this message
type _BACnetConstructedDataManualSlaveAddressBinding struct {
	BACnetConstructedDataContract
	ManualSlaveAddressBinding []BACnetAddressBinding
}

var _ BACnetConstructedDataManualSlaveAddressBinding = (*_BACnetConstructedDataManualSlaveAddressBinding)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataManualSlaveAddressBinding)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataManualSlaveAddressBinding) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MANUAL_SLAVE_ADDRESS_BINDING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataManualSlaveAddressBinding) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataManualSlaveAddressBinding) GetManualSlaveAddressBinding() []BACnetAddressBinding {
	return m.ManualSlaveAddressBinding
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataManualSlaveAddressBinding factory function for _BACnetConstructedDataManualSlaveAddressBinding
func NewBACnetConstructedDataManualSlaveAddressBinding(manualSlaveAddressBinding []BACnetAddressBinding, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataManualSlaveAddressBinding {
	_result := &_BACnetConstructedDataManualSlaveAddressBinding{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ManualSlaveAddressBinding:     manualSlaveAddressBinding,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataManualSlaveAddressBinding(structType any) BACnetConstructedDataManualSlaveAddressBinding {
	if casted, ok := structType.(BACnetConstructedDataManualSlaveAddressBinding); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataManualSlaveAddressBinding); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) GetTypeName() string {
	return "BACnetConstructedDataManualSlaveAddressBinding"
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.ManualSlaveAddressBinding) > 0 {
		for _, element := range m.ManualSlaveAddressBinding {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataManualSlaveAddressBinding BACnetConstructedDataManualSlaveAddressBinding, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataManualSlaveAddressBinding"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataManualSlaveAddressBinding")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	manualSlaveAddressBinding, err := ReadTerminatedArrayField[BACnetAddressBinding](ctx, "manualSlaveAddressBinding", ReadComplex[BACnetAddressBinding](BACnetAddressBindingParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'manualSlaveAddressBinding' field"))
	}
	m.ManualSlaveAddressBinding = manualSlaveAddressBinding

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataManualSlaveAddressBinding"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataManualSlaveAddressBinding")
	}

	return m, nil
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataManualSlaveAddressBinding"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataManualSlaveAddressBinding")
		}

		if err := WriteComplexTypeArrayField(ctx, "manualSlaveAddressBinding", m.GetManualSlaveAddressBinding(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'manualSlaveAddressBinding' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataManualSlaveAddressBinding"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataManualSlaveAddressBinding")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) isBACnetConstructedDataManualSlaveAddressBinding() bool {
	return true
}

func (m *_BACnetConstructedDataManualSlaveAddressBinding) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
