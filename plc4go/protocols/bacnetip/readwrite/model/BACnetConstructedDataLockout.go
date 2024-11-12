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

// BACnetConstructedDataLockout is the corresponding interface of BACnetConstructedDataLockout
type BACnetConstructedDataLockout interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLockout returns Lockout (property field)
	GetLockout() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataLockout is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLockout()
	// CreateBuilder creates a BACnetConstructedDataLockoutBuilder
	CreateBACnetConstructedDataLockoutBuilder() BACnetConstructedDataLockoutBuilder
}

// _BACnetConstructedDataLockout is the data-structure of this message
type _BACnetConstructedDataLockout struct {
	BACnetConstructedDataContract
	Lockout BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataLockout = (*_BACnetConstructedDataLockout)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLockout)(nil)

// NewBACnetConstructedDataLockout factory function for _BACnetConstructedDataLockout
func NewBACnetConstructedDataLockout(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, lockout BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLockout {
	if lockout == nil {
		panic("lockout of type BACnetApplicationTagBoolean for BACnetConstructedDataLockout must not be nil")
	}
	_result := &_BACnetConstructedDataLockout{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Lockout:                       lockout,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLockoutBuilder is a builder for BACnetConstructedDataLockout
type BACnetConstructedDataLockoutBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lockout BACnetApplicationTagBoolean) BACnetConstructedDataLockoutBuilder
	// WithLockout adds Lockout (property field)
	WithLockout(BACnetApplicationTagBoolean) BACnetConstructedDataLockoutBuilder
	// WithLockoutBuilder adds Lockout (property field) which is build by the builder
	WithLockoutBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataLockoutBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLockout or returns an error if something is wrong
	Build() (BACnetConstructedDataLockout, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLockout
}

// NewBACnetConstructedDataLockoutBuilder() creates a BACnetConstructedDataLockoutBuilder
func NewBACnetConstructedDataLockoutBuilder() BACnetConstructedDataLockoutBuilder {
	return &_BACnetConstructedDataLockoutBuilder{_BACnetConstructedDataLockout: new(_BACnetConstructedDataLockout)}
}

type _BACnetConstructedDataLockoutBuilder struct {
	*_BACnetConstructedDataLockout

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLockoutBuilder) = (*_BACnetConstructedDataLockoutBuilder)(nil)

func (b *_BACnetConstructedDataLockoutBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataLockoutBuilder) WithMandatoryFields(lockout BACnetApplicationTagBoolean) BACnetConstructedDataLockoutBuilder {
	return b.WithLockout(lockout)
}

func (b *_BACnetConstructedDataLockoutBuilder) WithLockout(lockout BACnetApplicationTagBoolean) BACnetConstructedDataLockoutBuilder {
	b.Lockout = lockout
	return b
}

func (b *_BACnetConstructedDataLockoutBuilder) WithLockoutBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataLockoutBuilder {
	builder := builderSupplier(b.Lockout.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	b.Lockout, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLockoutBuilder) Build() (BACnetConstructedDataLockout, error) {
	if b.Lockout == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lockout' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLockout.deepCopy(), nil
}

func (b *_BACnetConstructedDataLockoutBuilder) MustBuild() BACnetConstructedDataLockout {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLockoutBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLockoutBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLockoutBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLockoutBuilder().(*_BACnetConstructedDataLockoutBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLockoutBuilder creates a BACnetConstructedDataLockoutBuilder
func (b *_BACnetConstructedDataLockout) CreateBACnetConstructedDataLockoutBuilder() BACnetConstructedDataLockoutBuilder {
	if b == nil {
		return NewBACnetConstructedDataLockoutBuilder()
	}
	return &_BACnetConstructedDataLockoutBuilder{_BACnetConstructedDataLockout: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLockout) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLockout) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOCKOUT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLockout) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLockout) GetLockout() BACnetApplicationTagBoolean {
	return m.Lockout
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLockout) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetLockout())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLockout(structType any) BACnetConstructedDataLockout {
	if casted, ok := structType.(BACnetConstructedDataLockout); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLockout); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLockout) GetTypeName() string {
	return "BACnetConstructedDataLockout"
}

func (m *_BACnetConstructedDataLockout) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lockout)
	lengthInBits += m.Lockout.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLockout) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLockout) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLockout BACnetConstructedDataLockout, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLockout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLockout")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lockout, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "lockout", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lockout' field"))
	}
	m.Lockout = lockout

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), lockout)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLockout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLockout")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLockout) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLockout) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLockout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLockout")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "lockout", m.GetLockout(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lockout' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLockout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLockout")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLockout) IsBACnetConstructedDataLockout() {}

func (m *_BACnetConstructedDataLockout) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLockout) deepCopy() *_BACnetConstructedDataLockout {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLockoutCopy := &_BACnetConstructedDataLockout{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagBoolean](m.Lockout),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLockoutCopy
}

func (m *_BACnetConstructedDataLockout) String() string {
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
