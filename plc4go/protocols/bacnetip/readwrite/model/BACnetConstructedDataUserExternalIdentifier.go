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

// BACnetConstructedDataUserExternalIdentifier is the corresponding interface of BACnetConstructedDataUserExternalIdentifier
type BACnetConstructedDataUserExternalIdentifier interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetUserExternalIdentifier returns UserExternalIdentifier (property field)
	GetUserExternalIdentifier() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
	// IsBACnetConstructedDataUserExternalIdentifier is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataUserExternalIdentifier()
	// CreateBuilder creates a BACnetConstructedDataUserExternalIdentifierBuilder
	CreateBACnetConstructedDataUserExternalIdentifierBuilder() BACnetConstructedDataUserExternalIdentifierBuilder
}

// _BACnetConstructedDataUserExternalIdentifier is the data-structure of this message
type _BACnetConstructedDataUserExternalIdentifier struct {
	BACnetConstructedDataContract
	UserExternalIdentifier BACnetApplicationTagCharacterString
}

var _ BACnetConstructedDataUserExternalIdentifier = (*_BACnetConstructedDataUserExternalIdentifier)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataUserExternalIdentifier)(nil)

// NewBACnetConstructedDataUserExternalIdentifier factory function for _BACnetConstructedDataUserExternalIdentifier
func NewBACnetConstructedDataUserExternalIdentifier(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, userExternalIdentifier BACnetApplicationTagCharacterString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataUserExternalIdentifier {
	if userExternalIdentifier == nil {
		panic("userExternalIdentifier of type BACnetApplicationTagCharacterString for BACnetConstructedDataUserExternalIdentifier must not be nil")
	}
	_result := &_BACnetConstructedDataUserExternalIdentifier{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		UserExternalIdentifier:        userExternalIdentifier,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataUserExternalIdentifierBuilder is a builder for BACnetConstructedDataUserExternalIdentifier
type BACnetConstructedDataUserExternalIdentifierBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(userExternalIdentifier BACnetApplicationTagCharacterString) BACnetConstructedDataUserExternalIdentifierBuilder
	// WithUserExternalIdentifier adds UserExternalIdentifier (property field)
	WithUserExternalIdentifier(BACnetApplicationTagCharacterString) BACnetConstructedDataUserExternalIdentifierBuilder
	// WithUserExternalIdentifierBuilder adds UserExternalIdentifier (property field) which is build by the builder
	WithUserExternalIdentifierBuilder(func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataUserExternalIdentifierBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataUserExternalIdentifier or returns an error if something is wrong
	Build() (BACnetConstructedDataUserExternalIdentifier, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataUserExternalIdentifier
}

// NewBACnetConstructedDataUserExternalIdentifierBuilder() creates a BACnetConstructedDataUserExternalIdentifierBuilder
func NewBACnetConstructedDataUserExternalIdentifierBuilder() BACnetConstructedDataUserExternalIdentifierBuilder {
	return &_BACnetConstructedDataUserExternalIdentifierBuilder{_BACnetConstructedDataUserExternalIdentifier: new(_BACnetConstructedDataUserExternalIdentifier)}
}

type _BACnetConstructedDataUserExternalIdentifierBuilder struct {
	*_BACnetConstructedDataUserExternalIdentifier

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataUserExternalIdentifierBuilder) = (*_BACnetConstructedDataUserExternalIdentifierBuilder)(nil)

func (b *_BACnetConstructedDataUserExternalIdentifierBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataUserExternalIdentifierBuilder) WithMandatoryFields(userExternalIdentifier BACnetApplicationTagCharacterString) BACnetConstructedDataUserExternalIdentifierBuilder {
	return b.WithUserExternalIdentifier(userExternalIdentifier)
}

func (b *_BACnetConstructedDataUserExternalIdentifierBuilder) WithUserExternalIdentifier(userExternalIdentifier BACnetApplicationTagCharacterString) BACnetConstructedDataUserExternalIdentifierBuilder {
	b.UserExternalIdentifier = userExternalIdentifier
	return b
}

func (b *_BACnetConstructedDataUserExternalIdentifierBuilder) WithUserExternalIdentifierBuilder(builderSupplier func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataUserExternalIdentifierBuilder {
	builder := builderSupplier(b.UserExternalIdentifier.CreateBACnetApplicationTagCharacterStringBuilder())
	var err error
	b.UserExternalIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagCharacterStringBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataUserExternalIdentifierBuilder) Build() (BACnetConstructedDataUserExternalIdentifier, error) {
	if b.UserExternalIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'userExternalIdentifier' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataUserExternalIdentifier.deepCopy(), nil
}

func (b *_BACnetConstructedDataUserExternalIdentifierBuilder) MustBuild() BACnetConstructedDataUserExternalIdentifier {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataUserExternalIdentifierBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataUserExternalIdentifierBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataUserExternalIdentifierBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataUserExternalIdentifierBuilder().(*_BACnetConstructedDataUserExternalIdentifierBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataUserExternalIdentifierBuilder creates a BACnetConstructedDataUserExternalIdentifierBuilder
func (b *_BACnetConstructedDataUserExternalIdentifier) CreateBACnetConstructedDataUserExternalIdentifierBuilder() BACnetConstructedDataUserExternalIdentifierBuilder {
	if b == nil {
		return NewBACnetConstructedDataUserExternalIdentifierBuilder()
	}
	return &_BACnetConstructedDataUserExternalIdentifierBuilder{_BACnetConstructedDataUserExternalIdentifier: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataUserExternalIdentifier) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataUserExternalIdentifier) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_USER_EXTERNAL_IDENTIFIER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataUserExternalIdentifier) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataUserExternalIdentifier) GetUserExternalIdentifier() BACnetApplicationTagCharacterString {
	return m.UserExternalIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataUserExternalIdentifier) GetActualValue() BACnetApplicationTagCharacterString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagCharacterString(m.GetUserExternalIdentifier())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataUserExternalIdentifier(structType any) BACnetConstructedDataUserExternalIdentifier {
	if casted, ok := structType.(BACnetConstructedDataUserExternalIdentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataUserExternalIdentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataUserExternalIdentifier) GetTypeName() string {
	return "BACnetConstructedDataUserExternalIdentifier"
}

func (m *_BACnetConstructedDataUserExternalIdentifier) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (userExternalIdentifier)
	lengthInBits += m.UserExternalIdentifier.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataUserExternalIdentifier) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataUserExternalIdentifier) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataUserExternalIdentifier BACnetConstructedDataUserExternalIdentifier, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataUserExternalIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataUserExternalIdentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	userExternalIdentifier, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "userExternalIdentifier", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'userExternalIdentifier' field"))
	}
	m.UserExternalIdentifier = userExternalIdentifier

	actualValue, err := ReadVirtualField[BACnetApplicationTagCharacterString](ctx, "actualValue", (*BACnetApplicationTagCharacterString)(nil), userExternalIdentifier)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataUserExternalIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataUserExternalIdentifier")
	}

	return m, nil
}

func (m *_BACnetConstructedDataUserExternalIdentifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataUserExternalIdentifier) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataUserExternalIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataUserExternalIdentifier")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "userExternalIdentifier", m.GetUserExternalIdentifier(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'userExternalIdentifier' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataUserExternalIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataUserExternalIdentifier")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataUserExternalIdentifier) IsBACnetConstructedDataUserExternalIdentifier() {
}

func (m *_BACnetConstructedDataUserExternalIdentifier) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataUserExternalIdentifier) deepCopy() *_BACnetConstructedDataUserExternalIdentifier {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataUserExternalIdentifierCopy := &_BACnetConstructedDataUserExternalIdentifier{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagCharacterString](m.UserExternalIdentifier),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataUserExternalIdentifierCopy
}

func (m *_BACnetConstructedDataUserExternalIdentifier) String() string {
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
