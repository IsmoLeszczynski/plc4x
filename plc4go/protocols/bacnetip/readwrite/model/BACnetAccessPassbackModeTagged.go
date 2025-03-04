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

// BACnetAccessPassbackModeTagged is the corresponding interface of BACnetAccessPassbackModeTagged
type BACnetAccessPassbackModeTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetAccessPassbackMode
	// IsBACnetAccessPassbackModeTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetAccessPassbackModeTagged()
	// CreateBuilder creates a BACnetAccessPassbackModeTaggedBuilder
	CreateBACnetAccessPassbackModeTaggedBuilder() BACnetAccessPassbackModeTaggedBuilder
}

// _BACnetAccessPassbackModeTagged is the data-structure of this message
type _BACnetAccessPassbackModeTagged struct {
	Header BACnetTagHeader
	Value  BACnetAccessPassbackMode

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetAccessPassbackModeTagged = (*_BACnetAccessPassbackModeTagged)(nil)

// NewBACnetAccessPassbackModeTagged factory function for _BACnetAccessPassbackModeTagged
func NewBACnetAccessPassbackModeTagged(header BACnetTagHeader, value BACnetAccessPassbackMode, tagNumber uint8, tagClass TagClass) *_BACnetAccessPassbackModeTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetAccessPassbackModeTagged must not be nil")
	}
	return &_BACnetAccessPassbackModeTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetAccessPassbackModeTaggedBuilder is a builder for BACnetAccessPassbackModeTagged
type BACnetAccessPassbackModeTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetAccessPassbackMode) BACnetAccessPassbackModeTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetAccessPassbackModeTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetAccessPassbackModeTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetAccessPassbackMode) BACnetAccessPassbackModeTaggedBuilder
	// WithArgTagNumber sets a parser argument
	WithArgTagNumber(uint8) BACnetAccessPassbackModeTaggedBuilder
	// WithArgTagClass sets a parser argument
	WithArgTagClass(TagClass) BACnetAccessPassbackModeTaggedBuilder
	// Build builds the BACnetAccessPassbackModeTagged or returns an error if something is wrong
	Build() (BACnetAccessPassbackModeTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetAccessPassbackModeTagged
}

// NewBACnetAccessPassbackModeTaggedBuilder() creates a BACnetAccessPassbackModeTaggedBuilder
func NewBACnetAccessPassbackModeTaggedBuilder() BACnetAccessPassbackModeTaggedBuilder {
	return &_BACnetAccessPassbackModeTaggedBuilder{_BACnetAccessPassbackModeTagged: new(_BACnetAccessPassbackModeTagged)}
}

type _BACnetAccessPassbackModeTaggedBuilder struct {
	*_BACnetAccessPassbackModeTagged

	err *utils.MultiError
}

var _ (BACnetAccessPassbackModeTaggedBuilder) = (*_BACnetAccessPassbackModeTaggedBuilder)(nil)

func (b *_BACnetAccessPassbackModeTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetAccessPassbackMode) BACnetAccessPassbackModeTaggedBuilder {
	return b.WithHeader(header).WithValue(value)
}

func (b *_BACnetAccessPassbackModeTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetAccessPassbackModeTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetAccessPassbackModeTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetAccessPassbackModeTaggedBuilder {
	builder := builderSupplier(b.Header.CreateBACnetTagHeaderBuilder())
	var err error
	b.Header, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagHeaderBuilder failed"))
	}
	return b
}

func (b *_BACnetAccessPassbackModeTaggedBuilder) WithValue(value BACnetAccessPassbackMode) BACnetAccessPassbackModeTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetAccessPassbackModeTaggedBuilder) WithArgTagNumber(tagNumber uint8) BACnetAccessPassbackModeTaggedBuilder {
	b.TagNumber = tagNumber
	return b
}
func (b *_BACnetAccessPassbackModeTaggedBuilder) WithArgTagClass(tagClass TagClass) BACnetAccessPassbackModeTaggedBuilder {
	b.TagClass = tagClass
	return b
}

func (b *_BACnetAccessPassbackModeTaggedBuilder) Build() (BACnetAccessPassbackModeTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetAccessPassbackModeTagged.deepCopy(), nil
}

func (b *_BACnetAccessPassbackModeTaggedBuilder) MustBuild() BACnetAccessPassbackModeTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetAccessPassbackModeTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetAccessPassbackModeTaggedBuilder().(*_BACnetAccessPassbackModeTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetAccessPassbackModeTaggedBuilder creates a BACnetAccessPassbackModeTaggedBuilder
func (b *_BACnetAccessPassbackModeTagged) CreateBACnetAccessPassbackModeTaggedBuilder() BACnetAccessPassbackModeTaggedBuilder {
	if b == nil {
		return NewBACnetAccessPassbackModeTaggedBuilder()
	}
	return &_BACnetAccessPassbackModeTaggedBuilder{_BACnetAccessPassbackModeTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAccessPassbackModeTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetAccessPassbackModeTagged) GetValue() BACnetAccessPassbackMode {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetAccessPassbackModeTagged(structType any) BACnetAccessPassbackModeTagged {
	if casted, ok := structType.(BACnetAccessPassbackModeTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAccessPassbackModeTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAccessPassbackModeTagged) GetTypeName() string {
	return "BACnetAccessPassbackModeTagged"
}

func (m *_BACnetAccessPassbackModeTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetAccessPassbackModeTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAccessPassbackModeTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetAccessPassbackModeTagged, error) {
	return BACnetAccessPassbackModeTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetAccessPassbackModeTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessPassbackModeTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessPassbackModeTagged, error) {
		return BACnetAccessPassbackModeTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetAccessPassbackModeTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetAccessPassbackModeTagged, error) {
	v, err := (&_BACnetAccessPassbackModeTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetAccessPassbackModeTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetAccessPassbackModeTagged BACnetAccessPassbackModeTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAccessPassbackModeTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAccessPassbackModeTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	value, err := ReadManualField[BACnetAccessPassbackMode](ctx, "value", readBuffer, EnsureType[BACnetAccessPassbackMode](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetAccessPassbackMode_PASSBACK_OFF)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetAccessPassbackModeTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAccessPassbackModeTagged")
	}

	return m, nil
}

func (m *_BACnetAccessPassbackModeTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAccessPassbackModeTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAccessPassbackModeTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAccessPassbackModeTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetAccessPassbackMode](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAccessPassbackModeTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAccessPassbackModeTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetAccessPassbackModeTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetAccessPassbackModeTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetAccessPassbackModeTagged) IsBACnetAccessPassbackModeTagged() {}

func (m *_BACnetAccessPassbackModeTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetAccessPassbackModeTagged) deepCopy() *_BACnetAccessPassbackModeTagged {
	if m == nil {
		return nil
	}
	_BACnetAccessPassbackModeTaggedCopy := &_BACnetAccessPassbackModeTagged{
		utils.DeepCopy[BACnetTagHeader](m.Header),
		m.Value,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetAccessPassbackModeTaggedCopy
}

func (m *_BACnetAccessPassbackModeTagged) String() string {
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
