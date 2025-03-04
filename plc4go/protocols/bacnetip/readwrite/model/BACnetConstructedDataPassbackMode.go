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

// BACnetConstructedDataPassbackMode is the corresponding interface of BACnetConstructedDataPassbackMode
type BACnetConstructedDataPassbackMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetPassbackMode returns PassbackMode (property field)
	GetPassbackMode() BACnetAccessPassbackModeTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAccessPassbackModeTagged
	// IsBACnetConstructedDataPassbackMode is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataPassbackMode()
	// CreateBuilder creates a BACnetConstructedDataPassbackModeBuilder
	CreateBACnetConstructedDataPassbackModeBuilder() BACnetConstructedDataPassbackModeBuilder
}

// _BACnetConstructedDataPassbackMode is the data-structure of this message
type _BACnetConstructedDataPassbackMode struct {
	BACnetConstructedDataContract
	PassbackMode BACnetAccessPassbackModeTagged
}

var _ BACnetConstructedDataPassbackMode = (*_BACnetConstructedDataPassbackMode)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataPassbackMode)(nil)

// NewBACnetConstructedDataPassbackMode factory function for _BACnetConstructedDataPassbackMode
func NewBACnetConstructedDataPassbackMode(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, passbackMode BACnetAccessPassbackModeTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataPassbackMode {
	if passbackMode == nil {
		panic("passbackMode of type BACnetAccessPassbackModeTagged for BACnetConstructedDataPassbackMode must not be nil")
	}
	_result := &_BACnetConstructedDataPassbackMode{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		PassbackMode:                  passbackMode,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataPassbackModeBuilder is a builder for BACnetConstructedDataPassbackMode
type BACnetConstructedDataPassbackModeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(passbackMode BACnetAccessPassbackModeTagged) BACnetConstructedDataPassbackModeBuilder
	// WithPassbackMode adds PassbackMode (property field)
	WithPassbackMode(BACnetAccessPassbackModeTagged) BACnetConstructedDataPassbackModeBuilder
	// WithPassbackModeBuilder adds PassbackMode (property field) which is build by the builder
	WithPassbackModeBuilder(func(BACnetAccessPassbackModeTaggedBuilder) BACnetAccessPassbackModeTaggedBuilder) BACnetConstructedDataPassbackModeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataPassbackMode or returns an error if something is wrong
	Build() (BACnetConstructedDataPassbackMode, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataPassbackMode
}

// NewBACnetConstructedDataPassbackModeBuilder() creates a BACnetConstructedDataPassbackModeBuilder
func NewBACnetConstructedDataPassbackModeBuilder() BACnetConstructedDataPassbackModeBuilder {
	return &_BACnetConstructedDataPassbackModeBuilder{_BACnetConstructedDataPassbackMode: new(_BACnetConstructedDataPassbackMode)}
}

type _BACnetConstructedDataPassbackModeBuilder struct {
	*_BACnetConstructedDataPassbackMode

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataPassbackModeBuilder) = (*_BACnetConstructedDataPassbackModeBuilder)(nil)

func (b *_BACnetConstructedDataPassbackModeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataPassbackMode
}

func (b *_BACnetConstructedDataPassbackModeBuilder) WithMandatoryFields(passbackMode BACnetAccessPassbackModeTagged) BACnetConstructedDataPassbackModeBuilder {
	return b.WithPassbackMode(passbackMode)
}

func (b *_BACnetConstructedDataPassbackModeBuilder) WithPassbackMode(passbackMode BACnetAccessPassbackModeTagged) BACnetConstructedDataPassbackModeBuilder {
	b.PassbackMode = passbackMode
	return b
}

func (b *_BACnetConstructedDataPassbackModeBuilder) WithPassbackModeBuilder(builderSupplier func(BACnetAccessPassbackModeTaggedBuilder) BACnetAccessPassbackModeTaggedBuilder) BACnetConstructedDataPassbackModeBuilder {
	builder := builderSupplier(b.PassbackMode.CreateBACnetAccessPassbackModeTaggedBuilder())
	var err error
	b.PassbackMode, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetAccessPassbackModeTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataPassbackModeBuilder) Build() (BACnetConstructedDataPassbackMode, error) {
	if b.PassbackMode == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'passbackMode' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataPassbackMode.deepCopy(), nil
}

func (b *_BACnetConstructedDataPassbackModeBuilder) MustBuild() BACnetConstructedDataPassbackMode {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataPassbackModeBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataPassbackModeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataPassbackModeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataPassbackModeBuilder().(*_BACnetConstructedDataPassbackModeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataPassbackModeBuilder creates a BACnetConstructedDataPassbackModeBuilder
func (b *_BACnetConstructedDataPassbackMode) CreateBACnetConstructedDataPassbackModeBuilder() BACnetConstructedDataPassbackModeBuilder {
	if b == nil {
		return NewBACnetConstructedDataPassbackModeBuilder()
	}
	return &_BACnetConstructedDataPassbackModeBuilder{_BACnetConstructedDataPassbackMode: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPassbackMode) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataPassbackMode) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PASSBACK_MODE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPassbackMode) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPassbackMode) GetPassbackMode() BACnetAccessPassbackModeTagged {
	return m.PassbackMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPassbackMode) GetActualValue() BACnetAccessPassbackModeTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetAccessPassbackModeTagged(m.GetPassbackMode())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPassbackMode(structType any) BACnetConstructedDataPassbackMode {
	if casted, ok := structType.(BACnetConstructedDataPassbackMode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPassbackMode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPassbackMode) GetTypeName() string {
	return "BACnetConstructedDataPassbackMode"
}

func (m *_BACnetConstructedDataPassbackMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (passbackMode)
	lengthInBits += m.PassbackMode.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataPassbackMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataPassbackMode) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataPassbackMode BACnetConstructedDataPassbackMode, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPassbackMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPassbackMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	passbackMode, err := ReadSimpleField[BACnetAccessPassbackModeTagged](ctx, "passbackMode", ReadComplex[BACnetAccessPassbackModeTagged](BACnetAccessPassbackModeTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'passbackMode' field"))
	}
	m.PassbackMode = passbackMode

	actualValue, err := ReadVirtualField[BACnetAccessPassbackModeTagged](ctx, "actualValue", (*BACnetAccessPassbackModeTagged)(nil), passbackMode)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPassbackMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPassbackMode")
	}

	return m, nil
}

func (m *_BACnetConstructedDataPassbackMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataPassbackMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPassbackMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPassbackMode")
		}

		if err := WriteSimpleField[BACnetAccessPassbackModeTagged](ctx, "passbackMode", m.GetPassbackMode(), WriteComplex[BACnetAccessPassbackModeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'passbackMode' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPassbackMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPassbackMode")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataPassbackMode) IsBACnetConstructedDataPassbackMode() {}

func (m *_BACnetConstructedDataPassbackMode) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataPassbackMode) deepCopy() *_BACnetConstructedDataPassbackMode {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataPassbackModeCopy := &_BACnetConstructedDataPassbackMode{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetAccessPassbackModeTagged](m.PassbackMode),
	}
	_BACnetConstructedDataPassbackModeCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataPassbackModeCopy
}

func (m *_BACnetConstructedDataPassbackMode) String() string {
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
