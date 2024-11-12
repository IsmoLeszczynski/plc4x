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

// BACnetConstructedDataBACnetIPv6MulticastAddress is the corresponding interface of BACnetConstructedDataBACnetIPv6MulticastAddress
type BACnetConstructedDataBACnetIPv6MulticastAddress interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetIpv6MulticastAddress returns Ipv6MulticastAddress (property field)
	GetIpv6MulticastAddress() BACnetApplicationTagOctetString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagOctetString
	// IsBACnetConstructedDataBACnetIPv6MulticastAddress is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataBACnetIPv6MulticastAddress()
	// CreateBuilder creates a BACnetConstructedDataBACnetIPv6MulticastAddressBuilder
	CreateBACnetConstructedDataBACnetIPv6MulticastAddressBuilder() BACnetConstructedDataBACnetIPv6MulticastAddressBuilder
}

// _BACnetConstructedDataBACnetIPv6MulticastAddress is the data-structure of this message
type _BACnetConstructedDataBACnetIPv6MulticastAddress struct {
	BACnetConstructedDataContract
	Ipv6MulticastAddress BACnetApplicationTagOctetString
}

var _ BACnetConstructedDataBACnetIPv6MulticastAddress = (*_BACnetConstructedDataBACnetIPv6MulticastAddress)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBACnetIPv6MulticastAddress)(nil)

// NewBACnetConstructedDataBACnetIPv6MulticastAddress factory function for _BACnetConstructedDataBACnetIPv6MulticastAddress
func NewBACnetConstructedDataBACnetIPv6MulticastAddress(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, ipv6MulticastAddress BACnetApplicationTagOctetString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBACnetIPv6MulticastAddress {
	if ipv6MulticastAddress == nil {
		panic("ipv6MulticastAddress of type BACnetApplicationTagOctetString for BACnetConstructedDataBACnetIPv6MulticastAddress must not be nil")
	}
	_result := &_BACnetConstructedDataBACnetIPv6MulticastAddress{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Ipv6MulticastAddress:          ipv6MulticastAddress,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataBACnetIPv6MulticastAddressBuilder is a builder for BACnetConstructedDataBACnetIPv6MulticastAddress
type BACnetConstructedDataBACnetIPv6MulticastAddressBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(ipv6MulticastAddress BACnetApplicationTagOctetString) BACnetConstructedDataBACnetIPv6MulticastAddressBuilder
	// WithIpv6MulticastAddress adds Ipv6MulticastAddress (property field)
	WithIpv6MulticastAddress(BACnetApplicationTagOctetString) BACnetConstructedDataBACnetIPv6MulticastAddressBuilder
	// WithIpv6MulticastAddressBuilder adds Ipv6MulticastAddress (property field) which is build by the builder
	WithIpv6MulticastAddressBuilder(func(BACnetApplicationTagOctetStringBuilder) BACnetApplicationTagOctetStringBuilder) BACnetConstructedDataBACnetIPv6MulticastAddressBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataBACnetIPv6MulticastAddress or returns an error if something is wrong
	Build() (BACnetConstructedDataBACnetIPv6MulticastAddress, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataBACnetIPv6MulticastAddress
}

// NewBACnetConstructedDataBACnetIPv6MulticastAddressBuilder() creates a BACnetConstructedDataBACnetIPv6MulticastAddressBuilder
func NewBACnetConstructedDataBACnetIPv6MulticastAddressBuilder() BACnetConstructedDataBACnetIPv6MulticastAddressBuilder {
	return &_BACnetConstructedDataBACnetIPv6MulticastAddressBuilder{_BACnetConstructedDataBACnetIPv6MulticastAddress: new(_BACnetConstructedDataBACnetIPv6MulticastAddress)}
}

type _BACnetConstructedDataBACnetIPv6MulticastAddressBuilder struct {
	*_BACnetConstructedDataBACnetIPv6MulticastAddress

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataBACnetIPv6MulticastAddressBuilder) = (*_BACnetConstructedDataBACnetIPv6MulticastAddressBuilder)(nil)

func (b *_BACnetConstructedDataBACnetIPv6MulticastAddressBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataBACnetIPv6MulticastAddressBuilder) WithMandatoryFields(ipv6MulticastAddress BACnetApplicationTagOctetString) BACnetConstructedDataBACnetIPv6MulticastAddressBuilder {
	return b.WithIpv6MulticastAddress(ipv6MulticastAddress)
}

func (b *_BACnetConstructedDataBACnetIPv6MulticastAddressBuilder) WithIpv6MulticastAddress(ipv6MulticastAddress BACnetApplicationTagOctetString) BACnetConstructedDataBACnetIPv6MulticastAddressBuilder {
	b.Ipv6MulticastAddress = ipv6MulticastAddress
	return b
}

func (b *_BACnetConstructedDataBACnetIPv6MulticastAddressBuilder) WithIpv6MulticastAddressBuilder(builderSupplier func(BACnetApplicationTagOctetStringBuilder) BACnetApplicationTagOctetStringBuilder) BACnetConstructedDataBACnetIPv6MulticastAddressBuilder {
	builder := builderSupplier(b.Ipv6MulticastAddress.CreateBACnetApplicationTagOctetStringBuilder())
	var err error
	b.Ipv6MulticastAddress, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagOctetStringBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataBACnetIPv6MulticastAddressBuilder) Build() (BACnetConstructedDataBACnetIPv6MulticastAddress, error) {
	if b.Ipv6MulticastAddress == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'ipv6MulticastAddress' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataBACnetIPv6MulticastAddress.deepCopy(), nil
}

func (b *_BACnetConstructedDataBACnetIPv6MulticastAddressBuilder) MustBuild() BACnetConstructedDataBACnetIPv6MulticastAddress {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataBACnetIPv6MulticastAddressBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataBACnetIPv6MulticastAddressBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataBACnetIPv6MulticastAddressBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataBACnetIPv6MulticastAddressBuilder().(*_BACnetConstructedDataBACnetIPv6MulticastAddressBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataBACnetIPv6MulticastAddressBuilder creates a BACnetConstructedDataBACnetIPv6MulticastAddressBuilder
func (b *_BACnetConstructedDataBACnetIPv6MulticastAddress) CreateBACnetConstructedDataBACnetIPv6MulticastAddressBuilder() BACnetConstructedDataBACnetIPv6MulticastAddressBuilder {
	if b == nil {
		return NewBACnetConstructedDataBACnetIPv6MulticastAddressBuilder()
	}
	return &_BACnetConstructedDataBACnetIPv6MulticastAddressBuilder{_BACnetConstructedDataBACnetIPv6MulticastAddress: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BACNET_IPV6_MULTICAST_ADDRESS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) GetIpv6MulticastAddress() BACnetApplicationTagOctetString {
	return m.Ipv6MulticastAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) GetActualValue() BACnetApplicationTagOctetString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagOctetString(m.GetIpv6MulticastAddress())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBACnetIPv6MulticastAddress(structType any) BACnetConstructedDataBACnetIPv6MulticastAddress {
	if casted, ok := structType.(BACnetConstructedDataBACnetIPv6MulticastAddress); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBACnetIPv6MulticastAddress); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) GetTypeName() string {
	return "BACnetConstructedDataBACnetIPv6MulticastAddress"
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (ipv6MulticastAddress)
	lengthInBits += m.Ipv6MulticastAddress.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBACnetIPv6MulticastAddress BACnetConstructedDataBACnetIPv6MulticastAddress, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBACnetIPv6MulticastAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBACnetIPv6MulticastAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ipv6MulticastAddress, err := ReadSimpleField[BACnetApplicationTagOctetString](ctx, "ipv6MulticastAddress", ReadComplex[BACnetApplicationTagOctetString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagOctetString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ipv6MulticastAddress' field"))
	}
	m.Ipv6MulticastAddress = ipv6MulticastAddress

	actualValue, err := ReadVirtualField[BACnetApplicationTagOctetString](ctx, "actualValue", (*BACnetApplicationTagOctetString)(nil), ipv6MulticastAddress)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBACnetIPv6MulticastAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBACnetIPv6MulticastAddress")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBACnetIPv6MulticastAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBACnetIPv6MulticastAddress")
		}

		if err := WriteSimpleField[BACnetApplicationTagOctetString](ctx, "ipv6MulticastAddress", m.GetIpv6MulticastAddress(), WriteComplex[BACnetApplicationTagOctetString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'ipv6MulticastAddress' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBACnetIPv6MulticastAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBACnetIPv6MulticastAddress")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) IsBACnetConstructedDataBACnetIPv6MulticastAddress() {
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) deepCopy() *_BACnetConstructedDataBACnetIPv6MulticastAddress {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataBACnetIPv6MulticastAddressCopy := &_BACnetConstructedDataBACnetIPv6MulticastAddress{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagOctetString](m.Ipv6MulticastAddress),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataBACnetIPv6MulticastAddressCopy
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
