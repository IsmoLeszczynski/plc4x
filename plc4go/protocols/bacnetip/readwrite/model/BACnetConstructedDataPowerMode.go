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

// BACnetConstructedDataPowerMode is the corresponding interface of BACnetConstructedDataPowerMode
type BACnetConstructedDataPowerMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetPowerMode returns PowerMode (property field)
	GetPowerMode() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataPowerMode is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataPowerMode()
	// CreateBuilder creates a BACnetConstructedDataPowerModeBuilder
	CreateBACnetConstructedDataPowerModeBuilder() BACnetConstructedDataPowerModeBuilder
}

// _BACnetConstructedDataPowerMode is the data-structure of this message
type _BACnetConstructedDataPowerMode struct {
	BACnetConstructedDataContract
	PowerMode BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataPowerMode = (*_BACnetConstructedDataPowerMode)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataPowerMode)(nil)

// NewBACnetConstructedDataPowerMode factory function for _BACnetConstructedDataPowerMode
func NewBACnetConstructedDataPowerMode(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, powerMode BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataPowerMode {
	if powerMode == nil {
		panic("powerMode of type BACnetApplicationTagBoolean for BACnetConstructedDataPowerMode must not be nil")
	}
	_result := &_BACnetConstructedDataPowerMode{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		PowerMode:                     powerMode,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataPowerModeBuilder is a builder for BACnetConstructedDataPowerMode
type BACnetConstructedDataPowerModeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(powerMode BACnetApplicationTagBoolean) BACnetConstructedDataPowerModeBuilder
	// WithPowerMode adds PowerMode (property field)
	WithPowerMode(BACnetApplicationTagBoolean) BACnetConstructedDataPowerModeBuilder
	// WithPowerModeBuilder adds PowerMode (property field) which is build by the builder
	WithPowerModeBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataPowerModeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataPowerMode or returns an error if something is wrong
	Build() (BACnetConstructedDataPowerMode, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataPowerMode
}

// NewBACnetConstructedDataPowerModeBuilder() creates a BACnetConstructedDataPowerModeBuilder
func NewBACnetConstructedDataPowerModeBuilder() BACnetConstructedDataPowerModeBuilder {
	return &_BACnetConstructedDataPowerModeBuilder{_BACnetConstructedDataPowerMode: new(_BACnetConstructedDataPowerMode)}
}

type _BACnetConstructedDataPowerModeBuilder struct {
	*_BACnetConstructedDataPowerMode

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataPowerModeBuilder) = (*_BACnetConstructedDataPowerModeBuilder)(nil)

func (b *_BACnetConstructedDataPowerModeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataPowerMode
}

func (b *_BACnetConstructedDataPowerModeBuilder) WithMandatoryFields(powerMode BACnetApplicationTagBoolean) BACnetConstructedDataPowerModeBuilder {
	return b.WithPowerMode(powerMode)
}

func (b *_BACnetConstructedDataPowerModeBuilder) WithPowerMode(powerMode BACnetApplicationTagBoolean) BACnetConstructedDataPowerModeBuilder {
	b.PowerMode = powerMode
	return b
}

func (b *_BACnetConstructedDataPowerModeBuilder) WithPowerModeBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataPowerModeBuilder {
	builder := builderSupplier(b.PowerMode.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	b.PowerMode, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataPowerModeBuilder) Build() (BACnetConstructedDataPowerMode, error) {
	if b.PowerMode == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'powerMode' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataPowerMode.deepCopy(), nil
}

func (b *_BACnetConstructedDataPowerModeBuilder) MustBuild() BACnetConstructedDataPowerMode {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataPowerModeBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataPowerModeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataPowerModeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataPowerModeBuilder().(*_BACnetConstructedDataPowerModeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataPowerModeBuilder creates a BACnetConstructedDataPowerModeBuilder
func (b *_BACnetConstructedDataPowerMode) CreateBACnetConstructedDataPowerModeBuilder() BACnetConstructedDataPowerModeBuilder {
	if b == nil {
		return NewBACnetConstructedDataPowerModeBuilder()
	}
	return &_BACnetConstructedDataPowerModeBuilder{_BACnetConstructedDataPowerMode: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPowerMode) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataPowerMode) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_POWER_MODE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPowerMode) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPowerMode) GetPowerMode() BACnetApplicationTagBoolean {
	return m.PowerMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPowerMode) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetPowerMode())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPowerMode(structType any) BACnetConstructedDataPowerMode {
	if casted, ok := structType.(BACnetConstructedDataPowerMode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPowerMode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPowerMode) GetTypeName() string {
	return "BACnetConstructedDataPowerMode"
}

func (m *_BACnetConstructedDataPowerMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (powerMode)
	lengthInBits += m.PowerMode.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataPowerMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataPowerMode) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataPowerMode BACnetConstructedDataPowerMode, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPowerMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPowerMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	powerMode, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "powerMode", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'powerMode' field"))
	}
	m.PowerMode = powerMode

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), powerMode)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPowerMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPowerMode")
	}

	return m, nil
}

func (m *_BACnetConstructedDataPowerMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataPowerMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPowerMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPowerMode")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "powerMode", m.GetPowerMode(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'powerMode' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPowerMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPowerMode")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataPowerMode) IsBACnetConstructedDataPowerMode() {}

func (m *_BACnetConstructedDataPowerMode) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataPowerMode) deepCopy() *_BACnetConstructedDataPowerMode {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataPowerModeCopy := &_BACnetConstructedDataPowerMode{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagBoolean](m.PowerMode),
	}
	_BACnetConstructedDataPowerModeCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataPowerModeCopy
}

func (m *_BACnetConstructedDataPowerMode) String() string {
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
