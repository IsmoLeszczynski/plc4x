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

// BACnetConstructedDataOctetstringValueAll is the corresponding interface of BACnetConstructedDataOctetstringValueAll
type BACnetConstructedDataOctetstringValueAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataOctetstringValueAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataOctetstringValueAll()
	// CreateBuilder creates a BACnetConstructedDataOctetstringValueAllBuilder
	CreateBACnetConstructedDataOctetstringValueAllBuilder() BACnetConstructedDataOctetstringValueAllBuilder
}

// _BACnetConstructedDataOctetstringValueAll is the data-structure of this message
type _BACnetConstructedDataOctetstringValueAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataOctetstringValueAll = (*_BACnetConstructedDataOctetstringValueAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataOctetstringValueAll)(nil)

// NewBACnetConstructedDataOctetstringValueAll factory function for _BACnetConstructedDataOctetstringValueAll
func NewBACnetConstructedDataOctetstringValueAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataOctetstringValueAll {
	_result := &_BACnetConstructedDataOctetstringValueAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataOctetstringValueAllBuilder is a builder for BACnetConstructedDataOctetstringValueAll
type BACnetConstructedDataOctetstringValueAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataOctetstringValueAllBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataOctetstringValueAll or returns an error if something is wrong
	Build() (BACnetConstructedDataOctetstringValueAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataOctetstringValueAll
}

// NewBACnetConstructedDataOctetstringValueAllBuilder() creates a BACnetConstructedDataOctetstringValueAllBuilder
func NewBACnetConstructedDataOctetstringValueAllBuilder() BACnetConstructedDataOctetstringValueAllBuilder {
	return &_BACnetConstructedDataOctetstringValueAllBuilder{_BACnetConstructedDataOctetstringValueAll: new(_BACnetConstructedDataOctetstringValueAll)}
}

type _BACnetConstructedDataOctetstringValueAllBuilder struct {
	*_BACnetConstructedDataOctetstringValueAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataOctetstringValueAllBuilder) = (*_BACnetConstructedDataOctetstringValueAllBuilder)(nil)

func (b *_BACnetConstructedDataOctetstringValueAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataOctetstringValueAllBuilder) WithMandatoryFields() BACnetConstructedDataOctetstringValueAllBuilder {
	return b
}

func (b *_BACnetConstructedDataOctetstringValueAllBuilder) Build() (BACnetConstructedDataOctetstringValueAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataOctetstringValueAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataOctetstringValueAllBuilder) MustBuild() BACnetConstructedDataOctetstringValueAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataOctetstringValueAllBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataOctetstringValueAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataOctetstringValueAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataOctetstringValueAllBuilder().(*_BACnetConstructedDataOctetstringValueAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataOctetstringValueAllBuilder creates a BACnetConstructedDataOctetstringValueAllBuilder
func (b *_BACnetConstructedDataOctetstringValueAll) CreateBACnetConstructedDataOctetstringValueAllBuilder() BACnetConstructedDataOctetstringValueAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataOctetstringValueAllBuilder()
	}
	return &_BACnetConstructedDataOctetstringValueAllBuilder{_BACnetConstructedDataOctetstringValueAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataOctetstringValueAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_OCTETSTRING_VALUE
}

func (m *_BACnetConstructedDataOctetstringValueAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataOctetstringValueAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataOctetstringValueAll(structType any) BACnetConstructedDataOctetstringValueAll {
	if casted, ok := structType.(BACnetConstructedDataOctetstringValueAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOctetstringValueAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataOctetstringValueAll) GetTypeName() string {
	return "BACnetConstructedDataOctetstringValueAll"
}

func (m *_BACnetConstructedDataOctetstringValueAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataOctetstringValueAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataOctetstringValueAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataOctetstringValueAll BACnetConstructedDataOctetstringValueAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOctetstringValueAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataOctetstringValueAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOctetstringValueAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataOctetstringValueAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataOctetstringValueAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataOctetstringValueAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOctetstringValueAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataOctetstringValueAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOctetstringValueAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataOctetstringValueAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataOctetstringValueAll) IsBACnetConstructedDataOctetstringValueAll() {}

func (m *_BACnetConstructedDataOctetstringValueAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataOctetstringValueAll) deepCopy() *_BACnetConstructedDataOctetstringValueAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataOctetstringValueAllCopy := &_BACnetConstructedDataOctetstringValueAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataOctetstringValueAllCopy
}

func (m *_BACnetConstructedDataOctetstringValueAll) String() string {
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
