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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataDateValueAll is the corresponding interface of BACnetConstructedDataDateValueAll
type BACnetConstructedDataDateValueAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataDateValueAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataDateValueAll()
	// CreateBuilder creates a BACnetConstructedDataDateValueAllBuilder
	CreateBACnetConstructedDataDateValueAllBuilder() BACnetConstructedDataDateValueAllBuilder
}

// _BACnetConstructedDataDateValueAll is the data-structure of this message
type _BACnetConstructedDataDateValueAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataDateValueAll = (*_BACnetConstructedDataDateValueAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDateValueAll)(nil)

// NewBACnetConstructedDataDateValueAll factory function for _BACnetConstructedDataDateValueAll
func NewBACnetConstructedDataDateValueAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDateValueAll {
	_result := &_BACnetConstructedDataDateValueAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataDateValueAllBuilder is a builder for BACnetConstructedDataDateValueAll
type BACnetConstructedDataDateValueAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataDateValueAllBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataDateValueAll or returns an error if something is wrong
	Build() (BACnetConstructedDataDateValueAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataDateValueAll
}

// NewBACnetConstructedDataDateValueAllBuilder() creates a BACnetConstructedDataDateValueAllBuilder
func NewBACnetConstructedDataDateValueAllBuilder() BACnetConstructedDataDateValueAllBuilder {
	return &_BACnetConstructedDataDateValueAllBuilder{_BACnetConstructedDataDateValueAll: new(_BACnetConstructedDataDateValueAll)}
}

type _BACnetConstructedDataDateValueAllBuilder struct {
	*_BACnetConstructedDataDateValueAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataDateValueAllBuilder) = (*_BACnetConstructedDataDateValueAllBuilder)(nil)

func (b *_BACnetConstructedDataDateValueAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataDateValueAllBuilder) WithMandatoryFields() BACnetConstructedDataDateValueAllBuilder {
	return b
}

func (b *_BACnetConstructedDataDateValueAllBuilder) Build() (BACnetConstructedDataDateValueAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataDateValueAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataDateValueAllBuilder) MustBuild() BACnetConstructedDataDateValueAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataDateValueAllBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataDateValueAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataDateValueAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataDateValueAllBuilder().(*_BACnetConstructedDataDateValueAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataDateValueAllBuilder creates a BACnetConstructedDataDateValueAllBuilder
func (b *_BACnetConstructedDataDateValueAll) CreateBACnetConstructedDataDateValueAllBuilder() BACnetConstructedDataDateValueAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataDateValueAllBuilder()
	}
	return &_BACnetConstructedDataDateValueAllBuilder{_BACnetConstructedDataDateValueAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDateValueAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_DATE_VALUE
}

func (m *_BACnetConstructedDataDateValueAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDateValueAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDateValueAll(structType any) BACnetConstructedDataDateValueAll {
	if casted, ok := structType.(BACnetConstructedDataDateValueAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDateValueAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDateValueAll) GetTypeName() string {
	return "BACnetConstructedDataDateValueAll"
}

func (m *_BACnetConstructedDataDateValueAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataDateValueAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDateValueAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDateValueAll BACnetConstructedDataDateValueAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDateValueAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDateValueAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDateValueAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDateValueAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDateValueAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDateValueAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDateValueAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDateValueAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDateValueAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDateValueAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDateValueAll) IsBACnetConstructedDataDateValueAll() {}

func (m *_BACnetConstructedDataDateValueAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataDateValueAll) deepCopy() *_BACnetConstructedDataDateValueAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataDateValueAllCopy := &_BACnetConstructedDataDateValueAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataDateValueAllCopy
}

func (m *_BACnetConstructedDataDateValueAll) String() string {
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
