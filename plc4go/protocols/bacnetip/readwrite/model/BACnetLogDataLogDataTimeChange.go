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

// BACnetLogDataLogDataTimeChange is the corresponding interface of BACnetLogDataLogDataTimeChange
type BACnetLogDataLogDataTimeChange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetLogData
	// GetTimeChange returns TimeChange (property field)
	GetTimeChange() BACnetContextTagReal
	// IsBACnetLogDataLogDataTimeChange is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLogDataLogDataTimeChange()
	// CreateBuilder creates a BACnetLogDataLogDataTimeChangeBuilder
	CreateBACnetLogDataLogDataTimeChangeBuilder() BACnetLogDataLogDataTimeChangeBuilder
}

// _BACnetLogDataLogDataTimeChange is the data-structure of this message
type _BACnetLogDataLogDataTimeChange struct {
	BACnetLogDataContract
	TimeChange BACnetContextTagReal
}

var _ BACnetLogDataLogDataTimeChange = (*_BACnetLogDataLogDataTimeChange)(nil)
var _ BACnetLogDataRequirements = (*_BACnetLogDataLogDataTimeChange)(nil)

// NewBACnetLogDataLogDataTimeChange factory function for _BACnetLogDataLogDataTimeChange
func NewBACnetLogDataLogDataTimeChange(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, timeChange BACnetContextTagReal, tagNumber uint8) *_BACnetLogDataLogDataTimeChange {
	if timeChange == nil {
		panic("timeChange of type BACnetContextTagReal for BACnetLogDataLogDataTimeChange must not be nil")
	}
	_result := &_BACnetLogDataLogDataTimeChange{
		BACnetLogDataContract: NewBACnetLogData(openingTag, peekedTagHeader, closingTag, tagNumber),
		TimeChange:            timeChange,
	}
	_result.BACnetLogDataContract.(*_BACnetLogData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetLogDataLogDataTimeChangeBuilder is a builder for BACnetLogDataLogDataTimeChange
type BACnetLogDataLogDataTimeChangeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(timeChange BACnetContextTagReal) BACnetLogDataLogDataTimeChangeBuilder
	// WithTimeChange adds TimeChange (property field)
	WithTimeChange(BACnetContextTagReal) BACnetLogDataLogDataTimeChangeBuilder
	// WithTimeChangeBuilder adds TimeChange (property field) which is build by the builder
	WithTimeChangeBuilder(func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetLogDataLogDataTimeChangeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetLogDataBuilder
	// Build builds the BACnetLogDataLogDataTimeChange or returns an error if something is wrong
	Build() (BACnetLogDataLogDataTimeChange, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetLogDataLogDataTimeChange
}

// NewBACnetLogDataLogDataTimeChangeBuilder() creates a BACnetLogDataLogDataTimeChangeBuilder
func NewBACnetLogDataLogDataTimeChangeBuilder() BACnetLogDataLogDataTimeChangeBuilder {
	return &_BACnetLogDataLogDataTimeChangeBuilder{_BACnetLogDataLogDataTimeChange: new(_BACnetLogDataLogDataTimeChange)}
}

type _BACnetLogDataLogDataTimeChangeBuilder struct {
	*_BACnetLogDataLogDataTimeChange

	parentBuilder *_BACnetLogDataBuilder

	err *utils.MultiError
}

var _ (BACnetLogDataLogDataTimeChangeBuilder) = (*_BACnetLogDataLogDataTimeChangeBuilder)(nil)

func (b *_BACnetLogDataLogDataTimeChangeBuilder) setParent(contract BACnetLogDataContract) {
	b.BACnetLogDataContract = contract
}

func (b *_BACnetLogDataLogDataTimeChangeBuilder) WithMandatoryFields(timeChange BACnetContextTagReal) BACnetLogDataLogDataTimeChangeBuilder {
	return b.WithTimeChange(timeChange)
}

func (b *_BACnetLogDataLogDataTimeChangeBuilder) WithTimeChange(timeChange BACnetContextTagReal) BACnetLogDataLogDataTimeChangeBuilder {
	b.TimeChange = timeChange
	return b
}

func (b *_BACnetLogDataLogDataTimeChangeBuilder) WithTimeChangeBuilder(builderSupplier func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetLogDataLogDataTimeChangeBuilder {
	builder := builderSupplier(b.TimeChange.CreateBACnetContextTagRealBuilder())
	var err error
	b.TimeChange, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetLogDataLogDataTimeChangeBuilder) Build() (BACnetLogDataLogDataTimeChange, error) {
	if b.TimeChange == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeChange' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetLogDataLogDataTimeChange.deepCopy(), nil
}

func (b *_BACnetLogDataLogDataTimeChangeBuilder) MustBuild() BACnetLogDataLogDataTimeChange {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetLogDataLogDataTimeChangeBuilder) Done() BACnetLogDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetLogDataBuilder().(*_BACnetLogDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetLogDataLogDataTimeChangeBuilder) buildForBACnetLogData() (BACnetLogData, error) {
	return b.Build()
}

func (b *_BACnetLogDataLogDataTimeChangeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetLogDataLogDataTimeChangeBuilder().(*_BACnetLogDataLogDataTimeChangeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetLogDataLogDataTimeChangeBuilder creates a BACnetLogDataLogDataTimeChangeBuilder
func (b *_BACnetLogDataLogDataTimeChange) CreateBACnetLogDataLogDataTimeChangeBuilder() BACnetLogDataLogDataTimeChangeBuilder {
	if b == nil {
		return NewBACnetLogDataLogDataTimeChangeBuilder()
	}
	return &_BACnetLogDataLogDataTimeChangeBuilder{_BACnetLogDataLogDataTimeChange: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogDataLogDataTimeChange) GetParent() BACnetLogDataContract {
	return m.BACnetLogDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogDataLogDataTimeChange) GetTimeChange() BACnetContextTagReal {
	return m.TimeChange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLogDataLogDataTimeChange(structType any) BACnetLogDataLogDataTimeChange {
	if casted, ok := structType.(BACnetLogDataLogDataTimeChange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogDataLogDataTimeChange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogDataLogDataTimeChange) GetTypeName() string {
	return "BACnetLogDataLogDataTimeChange"
}

func (m *_BACnetLogDataLogDataTimeChange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetLogDataContract.(*_BACnetLogData).getLengthInBits(ctx))

	// Simple field (timeChange)
	lengthInBits += m.TimeChange.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogDataLogDataTimeChange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetLogDataLogDataTimeChange) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetLogData, tagNumber uint8) (__bACnetLogDataLogDataTimeChange BACnetLogDataLogDataTimeChange, err error) {
	m.BACnetLogDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogDataTimeChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogDataTimeChange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timeChange, err := ReadSimpleField[BACnetContextTagReal](ctx, "timeChange", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeChange' field"))
	}
	m.TimeChange = timeChange

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogDataTimeChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogDataTimeChange")
	}

	return m, nil
}

func (m *_BACnetLogDataLogDataTimeChange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogDataLogDataTimeChange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogDataTimeChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogDataTimeChange")
		}

		if err := WriteSimpleField[BACnetContextTagReal](ctx, "timeChange", m.GetTimeChange(), WriteComplex[BACnetContextTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeChange' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogDataTimeChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogDataTimeChange")
		}
		return nil
	}
	return m.BACnetLogDataContract.(*_BACnetLogData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogDataLogDataTimeChange) IsBACnetLogDataLogDataTimeChange() {}

func (m *_BACnetLogDataLogDataTimeChange) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLogDataLogDataTimeChange) deepCopy() *_BACnetLogDataLogDataTimeChange {
	if m == nil {
		return nil
	}
	_BACnetLogDataLogDataTimeChangeCopy := &_BACnetLogDataLogDataTimeChange{
		m.BACnetLogDataContract.(*_BACnetLogData).deepCopy(),
		utils.DeepCopy[BACnetContextTagReal](m.TimeChange),
	}
	m.BACnetLogDataContract.(*_BACnetLogData)._SubType = m
	return _BACnetLogDataLogDataTimeChangeCopy
}

func (m *_BACnetLogDataLogDataTimeChange) String() string {
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
