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

// BACnetConstructedDataLoadControlAll is the corresponding interface of BACnetConstructedDataLoadControlAll
type BACnetConstructedDataLoadControlAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataLoadControlAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLoadControlAll()
	// CreateBuilder creates a BACnetConstructedDataLoadControlAllBuilder
	CreateBACnetConstructedDataLoadControlAllBuilder() BACnetConstructedDataLoadControlAllBuilder
}

// _BACnetConstructedDataLoadControlAll is the data-structure of this message
type _BACnetConstructedDataLoadControlAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataLoadControlAll = (*_BACnetConstructedDataLoadControlAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLoadControlAll)(nil)

// NewBACnetConstructedDataLoadControlAll factory function for _BACnetConstructedDataLoadControlAll
func NewBACnetConstructedDataLoadControlAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLoadControlAll {
	_result := &_BACnetConstructedDataLoadControlAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLoadControlAllBuilder is a builder for BACnetConstructedDataLoadControlAll
type BACnetConstructedDataLoadControlAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataLoadControlAllBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLoadControlAll or returns an error if something is wrong
	Build() (BACnetConstructedDataLoadControlAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLoadControlAll
}

// NewBACnetConstructedDataLoadControlAllBuilder() creates a BACnetConstructedDataLoadControlAllBuilder
func NewBACnetConstructedDataLoadControlAllBuilder() BACnetConstructedDataLoadControlAllBuilder {
	return &_BACnetConstructedDataLoadControlAllBuilder{_BACnetConstructedDataLoadControlAll: new(_BACnetConstructedDataLoadControlAll)}
}

type _BACnetConstructedDataLoadControlAllBuilder struct {
	*_BACnetConstructedDataLoadControlAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLoadControlAllBuilder) = (*_BACnetConstructedDataLoadControlAllBuilder)(nil)

func (b *_BACnetConstructedDataLoadControlAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataLoadControlAllBuilder) WithMandatoryFields() BACnetConstructedDataLoadControlAllBuilder {
	return b
}

func (b *_BACnetConstructedDataLoadControlAllBuilder) Build() (BACnetConstructedDataLoadControlAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLoadControlAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataLoadControlAllBuilder) MustBuild() BACnetConstructedDataLoadControlAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLoadControlAllBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLoadControlAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLoadControlAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLoadControlAllBuilder().(*_BACnetConstructedDataLoadControlAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLoadControlAllBuilder creates a BACnetConstructedDataLoadControlAllBuilder
func (b *_BACnetConstructedDataLoadControlAll) CreateBACnetConstructedDataLoadControlAllBuilder() BACnetConstructedDataLoadControlAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataLoadControlAllBuilder()
	}
	return &_BACnetConstructedDataLoadControlAllBuilder{_BACnetConstructedDataLoadControlAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLoadControlAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LOAD_CONTROL
}

func (m *_BACnetConstructedDataLoadControlAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLoadControlAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLoadControlAll(structType any) BACnetConstructedDataLoadControlAll {
	if casted, ok := structType.(BACnetConstructedDataLoadControlAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLoadControlAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLoadControlAll) GetTypeName() string {
	return "BACnetConstructedDataLoadControlAll"
}

func (m *_BACnetConstructedDataLoadControlAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataLoadControlAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLoadControlAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLoadControlAll BACnetConstructedDataLoadControlAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLoadControlAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLoadControlAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLoadControlAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLoadControlAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLoadControlAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLoadControlAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLoadControlAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLoadControlAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLoadControlAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLoadControlAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLoadControlAll) IsBACnetConstructedDataLoadControlAll() {}

func (m *_BACnetConstructedDataLoadControlAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLoadControlAll) deepCopy() *_BACnetConstructedDataLoadControlAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLoadControlAllCopy := &_BACnetConstructedDataLoadControlAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLoadControlAllCopy
}

func (m *_BACnetConstructedDataLoadControlAll) String() string {
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
