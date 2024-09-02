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

// BACnetConstructedDataIPv6Address is the corresponding interface of BACnetConstructedDataIPv6Address
type BACnetConstructedDataIPv6Address interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetIpv6Address returns Ipv6Address (property field)
	GetIpv6Address() BACnetApplicationTagOctetString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagOctetString
}

// BACnetConstructedDataIPv6AddressExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataIPv6Address.
// This is useful for switch cases.
type BACnetConstructedDataIPv6AddressExactly interface {
	BACnetConstructedDataIPv6Address
	isBACnetConstructedDataIPv6Address() bool
}

// _BACnetConstructedDataIPv6Address is the data-structure of this message
type _BACnetConstructedDataIPv6Address struct {
	BACnetConstructedDataContract
	Ipv6Address BACnetApplicationTagOctetString
}

var _ BACnetConstructedDataIPv6Address = (*_BACnetConstructedDataIPv6Address)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIPv6Address)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPv6Address) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPv6Address) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IPV6_ADDRESS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPv6Address) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6Address) GetIpv6Address() BACnetApplicationTagOctetString {
	return m.Ipv6Address
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6Address) GetActualValue() BACnetApplicationTagOctetString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagOctetString(m.GetIpv6Address())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIPv6Address factory function for _BACnetConstructedDataIPv6Address
func NewBACnetConstructedDataIPv6Address(ipv6Address BACnetApplicationTagOctetString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPv6Address {
	_result := &_BACnetConstructedDataIPv6Address{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Ipv6Address:                   ipv6Address,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPv6Address(structType any) BACnetConstructedDataIPv6Address {
	if casted, ok := structType.(BACnetConstructedDataIPv6Address); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPv6Address); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPv6Address) GetTypeName() string {
	return "BACnetConstructedDataIPv6Address"
}

func (m *_BACnetConstructedDataIPv6Address) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (ipv6Address)
	lengthInBits += m.Ipv6Address.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIPv6Address) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIPv6Address) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIPv6Address BACnetConstructedDataIPv6Address, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPv6Address"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPv6Address")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ipv6Address, err := ReadSimpleField[BACnetApplicationTagOctetString](ctx, "ipv6Address", ReadComplex[BACnetApplicationTagOctetString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagOctetString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ipv6Address' field"))
	}
	m.Ipv6Address = ipv6Address

	actualValue, err := ReadVirtualField[BACnetApplicationTagOctetString](ctx, "actualValue", (*BACnetApplicationTagOctetString)(nil), ipv6Address)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPv6Address"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPv6Address")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIPv6Address) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIPv6Address) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPv6Address"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPv6Address")
		}

		if err := WriteSimpleField[BACnetApplicationTagOctetString](ctx, "ipv6Address", m.GetIpv6Address(), WriteComplex[BACnetApplicationTagOctetString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'ipv6Address' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPv6Address"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPv6Address")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPv6Address) isBACnetConstructedDataIPv6Address() bool {
	return true
}

func (m *_BACnetConstructedDataIPv6Address) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
