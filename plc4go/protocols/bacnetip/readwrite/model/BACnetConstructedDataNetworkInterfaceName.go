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

// BACnetConstructedDataNetworkInterfaceName is the corresponding interface of BACnetConstructedDataNetworkInterfaceName
type BACnetConstructedDataNetworkInterfaceName interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNetworkInterfaceName returns NetworkInterfaceName (property field)
	GetNetworkInterfaceName() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
	// IsBACnetConstructedDataNetworkInterfaceName is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataNetworkInterfaceName()
	// CreateBuilder creates a BACnetConstructedDataNetworkInterfaceNameBuilder
	CreateBACnetConstructedDataNetworkInterfaceNameBuilder() BACnetConstructedDataNetworkInterfaceNameBuilder
}

// _BACnetConstructedDataNetworkInterfaceName is the data-structure of this message
type _BACnetConstructedDataNetworkInterfaceName struct {
	BACnetConstructedDataContract
	NetworkInterfaceName BACnetApplicationTagCharacterString
}

var _ BACnetConstructedDataNetworkInterfaceName = (*_BACnetConstructedDataNetworkInterfaceName)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataNetworkInterfaceName)(nil)

// NewBACnetConstructedDataNetworkInterfaceName factory function for _BACnetConstructedDataNetworkInterfaceName
func NewBACnetConstructedDataNetworkInterfaceName(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, networkInterfaceName BACnetApplicationTagCharacterString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataNetworkInterfaceName {
	if networkInterfaceName == nil {
		panic("networkInterfaceName of type BACnetApplicationTagCharacterString for BACnetConstructedDataNetworkInterfaceName must not be nil")
	}
	_result := &_BACnetConstructedDataNetworkInterfaceName{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NetworkInterfaceName:          networkInterfaceName,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataNetworkInterfaceNameBuilder is a builder for BACnetConstructedDataNetworkInterfaceName
type BACnetConstructedDataNetworkInterfaceNameBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(networkInterfaceName BACnetApplicationTagCharacterString) BACnetConstructedDataNetworkInterfaceNameBuilder
	// WithNetworkInterfaceName adds NetworkInterfaceName (property field)
	WithNetworkInterfaceName(BACnetApplicationTagCharacterString) BACnetConstructedDataNetworkInterfaceNameBuilder
	// WithNetworkInterfaceNameBuilder adds NetworkInterfaceName (property field) which is build by the builder
	WithNetworkInterfaceNameBuilder(func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataNetworkInterfaceNameBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataNetworkInterfaceName or returns an error if something is wrong
	Build() (BACnetConstructedDataNetworkInterfaceName, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataNetworkInterfaceName
}

// NewBACnetConstructedDataNetworkInterfaceNameBuilder() creates a BACnetConstructedDataNetworkInterfaceNameBuilder
func NewBACnetConstructedDataNetworkInterfaceNameBuilder() BACnetConstructedDataNetworkInterfaceNameBuilder {
	return &_BACnetConstructedDataNetworkInterfaceNameBuilder{_BACnetConstructedDataNetworkInterfaceName: new(_BACnetConstructedDataNetworkInterfaceName)}
}

type _BACnetConstructedDataNetworkInterfaceNameBuilder struct {
	*_BACnetConstructedDataNetworkInterfaceName

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataNetworkInterfaceNameBuilder) = (*_BACnetConstructedDataNetworkInterfaceNameBuilder)(nil)

func (b *_BACnetConstructedDataNetworkInterfaceNameBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataNetworkInterfaceNameBuilder) WithMandatoryFields(networkInterfaceName BACnetApplicationTagCharacterString) BACnetConstructedDataNetworkInterfaceNameBuilder {
	return b.WithNetworkInterfaceName(networkInterfaceName)
}

func (b *_BACnetConstructedDataNetworkInterfaceNameBuilder) WithNetworkInterfaceName(networkInterfaceName BACnetApplicationTagCharacterString) BACnetConstructedDataNetworkInterfaceNameBuilder {
	b.NetworkInterfaceName = networkInterfaceName
	return b
}

func (b *_BACnetConstructedDataNetworkInterfaceNameBuilder) WithNetworkInterfaceNameBuilder(builderSupplier func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataNetworkInterfaceNameBuilder {
	builder := builderSupplier(b.NetworkInterfaceName.CreateBACnetApplicationTagCharacterStringBuilder())
	var err error
	b.NetworkInterfaceName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagCharacterStringBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataNetworkInterfaceNameBuilder) Build() (BACnetConstructedDataNetworkInterfaceName, error) {
	if b.NetworkInterfaceName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'networkInterfaceName' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataNetworkInterfaceName.deepCopy(), nil
}

func (b *_BACnetConstructedDataNetworkInterfaceNameBuilder) MustBuild() BACnetConstructedDataNetworkInterfaceName {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataNetworkInterfaceNameBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataNetworkInterfaceNameBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataNetworkInterfaceNameBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataNetworkInterfaceNameBuilder().(*_BACnetConstructedDataNetworkInterfaceNameBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataNetworkInterfaceNameBuilder creates a BACnetConstructedDataNetworkInterfaceNameBuilder
func (b *_BACnetConstructedDataNetworkInterfaceName) CreateBACnetConstructedDataNetworkInterfaceNameBuilder() BACnetConstructedDataNetworkInterfaceNameBuilder {
	if b == nil {
		return NewBACnetConstructedDataNetworkInterfaceNameBuilder()
	}
	return &_BACnetConstructedDataNetworkInterfaceNameBuilder{_BACnetConstructedDataNetworkInterfaceName: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataNetworkInterfaceName) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataNetworkInterfaceName) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_NETWORK_INTERFACE_NAME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataNetworkInterfaceName) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataNetworkInterfaceName) GetNetworkInterfaceName() BACnetApplicationTagCharacterString {
	return m.NetworkInterfaceName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataNetworkInterfaceName) GetActualValue() BACnetApplicationTagCharacterString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagCharacterString(m.GetNetworkInterfaceName())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataNetworkInterfaceName(structType any) BACnetConstructedDataNetworkInterfaceName {
	if casted, ok := structType.(BACnetConstructedDataNetworkInterfaceName); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNetworkInterfaceName); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataNetworkInterfaceName) GetTypeName() string {
	return "BACnetConstructedDataNetworkInterfaceName"
}

func (m *_BACnetConstructedDataNetworkInterfaceName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (networkInterfaceName)
	lengthInBits += m.NetworkInterfaceName.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataNetworkInterfaceName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataNetworkInterfaceName) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataNetworkInterfaceName BACnetConstructedDataNetworkInterfaceName, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNetworkInterfaceName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNetworkInterfaceName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	networkInterfaceName, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "networkInterfaceName", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkInterfaceName' field"))
	}
	m.NetworkInterfaceName = networkInterfaceName

	actualValue, err := ReadVirtualField[BACnetApplicationTagCharacterString](ctx, "actualValue", (*BACnetApplicationTagCharacterString)(nil), networkInterfaceName)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNetworkInterfaceName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNetworkInterfaceName")
	}

	return m, nil
}

func (m *_BACnetConstructedDataNetworkInterfaceName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataNetworkInterfaceName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNetworkInterfaceName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNetworkInterfaceName")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "networkInterfaceName", m.GetNetworkInterfaceName(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'networkInterfaceName' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNetworkInterfaceName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNetworkInterfaceName")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataNetworkInterfaceName) IsBACnetConstructedDataNetworkInterfaceName() {}

func (m *_BACnetConstructedDataNetworkInterfaceName) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataNetworkInterfaceName) deepCopy() *_BACnetConstructedDataNetworkInterfaceName {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataNetworkInterfaceNameCopy := &_BACnetConstructedDataNetworkInterfaceName{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagCharacterString](m.NetworkInterfaceName),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataNetworkInterfaceNameCopy
}

func (m *_BACnetConstructedDataNetworkInterfaceName) String() string {
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
