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

// BACnetConstructedDataResolution is the corresponding interface of BACnetConstructedDataResolution
type BACnetConstructedDataResolution interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetResolution returns Resolution (property field)
	GetResolution() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
	// IsBACnetConstructedDataResolution is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataResolution()
	// CreateBuilder creates a BACnetConstructedDataResolutionBuilder
	CreateBACnetConstructedDataResolutionBuilder() BACnetConstructedDataResolutionBuilder
}

// _BACnetConstructedDataResolution is the data-structure of this message
type _BACnetConstructedDataResolution struct {
	BACnetConstructedDataContract
	Resolution BACnetApplicationTagReal
}

var _ BACnetConstructedDataResolution = (*_BACnetConstructedDataResolution)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataResolution)(nil)

// NewBACnetConstructedDataResolution factory function for _BACnetConstructedDataResolution
func NewBACnetConstructedDataResolution(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, resolution BACnetApplicationTagReal, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataResolution {
	if resolution == nil {
		panic("resolution of type BACnetApplicationTagReal for BACnetConstructedDataResolution must not be nil")
	}
	_result := &_BACnetConstructedDataResolution{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Resolution:                    resolution,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataResolutionBuilder is a builder for BACnetConstructedDataResolution
type BACnetConstructedDataResolutionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(resolution BACnetApplicationTagReal) BACnetConstructedDataResolutionBuilder
	// WithResolution adds Resolution (property field)
	WithResolution(BACnetApplicationTagReal) BACnetConstructedDataResolutionBuilder
	// WithResolutionBuilder adds Resolution (property field) which is build by the builder
	WithResolutionBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataResolutionBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataResolution or returns an error if something is wrong
	Build() (BACnetConstructedDataResolution, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataResolution
}

// NewBACnetConstructedDataResolutionBuilder() creates a BACnetConstructedDataResolutionBuilder
func NewBACnetConstructedDataResolutionBuilder() BACnetConstructedDataResolutionBuilder {
	return &_BACnetConstructedDataResolutionBuilder{_BACnetConstructedDataResolution: new(_BACnetConstructedDataResolution)}
}

type _BACnetConstructedDataResolutionBuilder struct {
	*_BACnetConstructedDataResolution

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataResolutionBuilder) = (*_BACnetConstructedDataResolutionBuilder)(nil)

func (b *_BACnetConstructedDataResolutionBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataResolutionBuilder) WithMandatoryFields(resolution BACnetApplicationTagReal) BACnetConstructedDataResolutionBuilder {
	return b.WithResolution(resolution)
}

func (b *_BACnetConstructedDataResolutionBuilder) WithResolution(resolution BACnetApplicationTagReal) BACnetConstructedDataResolutionBuilder {
	b.Resolution = resolution
	return b
}

func (b *_BACnetConstructedDataResolutionBuilder) WithResolutionBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataResolutionBuilder {
	builder := builderSupplier(b.Resolution.CreateBACnetApplicationTagRealBuilder())
	var err error
	b.Resolution, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataResolutionBuilder) Build() (BACnetConstructedDataResolution, error) {
	if b.Resolution == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'resolution' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataResolution.deepCopy(), nil
}

func (b *_BACnetConstructedDataResolutionBuilder) MustBuild() BACnetConstructedDataResolution {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataResolutionBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataResolutionBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataResolutionBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataResolutionBuilder().(*_BACnetConstructedDataResolutionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataResolutionBuilder creates a BACnetConstructedDataResolutionBuilder
func (b *_BACnetConstructedDataResolution) CreateBACnetConstructedDataResolutionBuilder() BACnetConstructedDataResolutionBuilder {
	if b == nil {
		return NewBACnetConstructedDataResolutionBuilder()
	}
	return &_BACnetConstructedDataResolutionBuilder{_BACnetConstructedDataResolution: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataResolution) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataResolution) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RESOLUTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataResolution) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataResolution) GetResolution() BACnetApplicationTagReal {
	return m.Resolution
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataResolution) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetResolution())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataResolution(structType any) BACnetConstructedDataResolution {
	if casted, ok := structType.(BACnetConstructedDataResolution); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataResolution); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataResolution) GetTypeName() string {
	return "BACnetConstructedDataResolution"
}

func (m *_BACnetConstructedDataResolution) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (resolution)
	lengthInBits += m.Resolution.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataResolution) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataResolution) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataResolution BACnetConstructedDataResolution, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataResolution"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataResolution")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	resolution, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "resolution", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'resolution' field"))
	}
	m.Resolution = resolution

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), resolution)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataResolution"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataResolution")
	}

	return m, nil
}

func (m *_BACnetConstructedDataResolution) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataResolution) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataResolution"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataResolution")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "resolution", m.GetResolution(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'resolution' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataResolution"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataResolution")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataResolution) IsBACnetConstructedDataResolution() {}

func (m *_BACnetConstructedDataResolution) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataResolution) deepCopy() *_BACnetConstructedDataResolution {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataResolutionCopy := &_BACnetConstructedDataResolution{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagReal](m.Resolution),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataResolutionCopy
}

func (m *_BACnetConstructedDataResolution) String() string {
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
