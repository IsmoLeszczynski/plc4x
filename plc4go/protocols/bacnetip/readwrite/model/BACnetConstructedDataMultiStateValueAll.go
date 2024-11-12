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

// BACnetConstructedDataMultiStateValueAll is the corresponding interface of BACnetConstructedDataMultiStateValueAll
type BACnetConstructedDataMultiStateValueAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataMultiStateValueAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataMultiStateValueAll()
	// CreateBuilder creates a BACnetConstructedDataMultiStateValueAllBuilder
	CreateBACnetConstructedDataMultiStateValueAllBuilder() BACnetConstructedDataMultiStateValueAllBuilder
}

// _BACnetConstructedDataMultiStateValueAll is the data-structure of this message
type _BACnetConstructedDataMultiStateValueAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataMultiStateValueAll = (*_BACnetConstructedDataMultiStateValueAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataMultiStateValueAll)(nil)

// NewBACnetConstructedDataMultiStateValueAll factory function for _BACnetConstructedDataMultiStateValueAll
func NewBACnetConstructedDataMultiStateValueAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMultiStateValueAll {
	_result := &_BACnetConstructedDataMultiStateValueAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataMultiStateValueAllBuilder is a builder for BACnetConstructedDataMultiStateValueAll
type BACnetConstructedDataMultiStateValueAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataMultiStateValueAllBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataMultiStateValueAll or returns an error if something is wrong
	Build() (BACnetConstructedDataMultiStateValueAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataMultiStateValueAll
}

// NewBACnetConstructedDataMultiStateValueAllBuilder() creates a BACnetConstructedDataMultiStateValueAllBuilder
func NewBACnetConstructedDataMultiStateValueAllBuilder() BACnetConstructedDataMultiStateValueAllBuilder {
	return &_BACnetConstructedDataMultiStateValueAllBuilder{_BACnetConstructedDataMultiStateValueAll: new(_BACnetConstructedDataMultiStateValueAll)}
}

type _BACnetConstructedDataMultiStateValueAllBuilder struct {
	*_BACnetConstructedDataMultiStateValueAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataMultiStateValueAllBuilder) = (*_BACnetConstructedDataMultiStateValueAllBuilder)(nil)

func (b *_BACnetConstructedDataMultiStateValueAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataMultiStateValueAllBuilder) WithMandatoryFields() BACnetConstructedDataMultiStateValueAllBuilder {
	return b
}

func (b *_BACnetConstructedDataMultiStateValueAllBuilder) Build() (BACnetConstructedDataMultiStateValueAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataMultiStateValueAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataMultiStateValueAllBuilder) MustBuild() BACnetConstructedDataMultiStateValueAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataMultiStateValueAllBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataMultiStateValueAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataMultiStateValueAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataMultiStateValueAllBuilder().(*_BACnetConstructedDataMultiStateValueAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataMultiStateValueAllBuilder creates a BACnetConstructedDataMultiStateValueAllBuilder
func (b *_BACnetConstructedDataMultiStateValueAll) CreateBACnetConstructedDataMultiStateValueAllBuilder() BACnetConstructedDataMultiStateValueAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataMultiStateValueAllBuilder()
	}
	return &_BACnetConstructedDataMultiStateValueAllBuilder{_BACnetConstructedDataMultiStateValueAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMultiStateValueAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_MULTI_STATE_VALUE
}

func (m *_BACnetConstructedDataMultiStateValueAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMultiStateValueAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMultiStateValueAll(structType any) BACnetConstructedDataMultiStateValueAll {
	if casted, ok := structType.(BACnetConstructedDataMultiStateValueAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMultiStateValueAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMultiStateValueAll) GetTypeName() string {
	return "BACnetConstructedDataMultiStateValueAll"
}

func (m *_BACnetConstructedDataMultiStateValueAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataMultiStateValueAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMultiStateValueAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMultiStateValueAll BACnetConstructedDataMultiStateValueAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMultiStateValueAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMultiStateValueAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMultiStateValueAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMultiStateValueAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMultiStateValueAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMultiStateValueAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMultiStateValueAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMultiStateValueAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMultiStateValueAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMultiStateValueAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMultiStateValueAll) IsBACnetConstructedDataMultiStateValueAll() {}

func (m *_BACnetConstructedDataMultiStateValueAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataMultiStateValueAll) deepCopy() *_BACnetConstructedDataMultiStateValueAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataMultiStateValueAllCopy := &_BACnetConstructedDataMultiStateValueAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataMultiStateValueAllCopy
}

func (m *_BACnetConstructedDataMultiStateValueAll) String() string {
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
