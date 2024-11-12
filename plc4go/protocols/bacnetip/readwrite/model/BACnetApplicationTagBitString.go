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

// BACnetApplicationTagBitString is the corresponding interface of BACnetApplicationTagBitString
type BACnetApplicationTagBitString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetApplicationTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadBitString
	// IsBACnetApplicationTagBitString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetApplicationTagBitString()
	// CreateBuilder creates a BACnetApplicationTagBitStringBuilder
	CreateBACnetApplicationTagBitStringBuilder() BACnetApplicationTagBitStringBuilder
}

// _BACnetApplicationTagBitString is the data-structure of this message
type _BACnetApplicationTagBitString struct {
	BACnetApplicationTagContract
	Payload BACnetTagPayloadBitString
}

var _ BACnetApplicationTagBitString = (*_BACnetApplicationTagBitString)(nil)
var _ BACnetApplicationTagRequirements = (*_BACnetApplicationTagBitString)(nil)

// NewBACnetApplicationTagBitString factory function for _BACnetApplicationTagBitString
func NewBACnetApplicationTagBitString(header BACnetTagHeader, payload BACnetTagPayloadBitString) *_BACnetApplicationTagBitString {
	if payload == nil {
		panic("payload of type BACnetTagPayloadBitString for BACnetApplicationTagBitString must not be nil")
	}
	_result := &_BACnetApplicationTagBitString{
		BACnetApplicationTagContract: NewBACnetApplicationTag(header),
		Payload:                      payload,
	}
	_result.BACnetApplicationTagContract.(*_BACnetApplicationTag)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetApplicationTagBitStringBuilder is a builder for BACnetApplicationTagBitString
type BACnetApplicationTagBitStringBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(payload BACnetTagPayloadBitString) BACnetApplicationTagBitStringBuilder
	// WithPayload adds Payload (property field)
	WithPayload(BACnetTagPayloadBitString) BACnetApplicationTagBitStringBuilder
	// WithPayloadBuilder adds Payload (property field) which is build by the builder
	WithPayloadBuilder(func(BACnetTagPayloadBitStringBuilder) BACnetTagPayloadBitStringBuilder) BACnetApplicationTagBitStringBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetApplicationTagBuilder
	// Build builds the BACnetApplicationTagBitString or returns an error if something is wrong
	Build() (BACnetApplicationTagBitString, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetApplicationTagBitString
}

// NewBACnetApplicationTagBitStringBuilder() creates a BACnetApplicationTagBitStringBuilder
func NewBACnetApplicationTagBitStringBuilder() BACnetApplicationTagBitStringBuilder {
	return &_BACnetApplicationTagBitStringBuilder{_BACnetApplicationTagBitString: new(_BACnetApplicationTagBitString)}
}

type _BACnetApplicationTagBitStringBuilder struct {
	*_BACnetApplicationTagBitString

	parentBuilder *_BACnetApplicationTagBuilder

	err *utils.MultiError
}

var _ (BACnetApplicationTagBitStringBuilder) = (*_BACnetApplicationTagBitStringBuilder)(nil)

func (b *_BACnetApplicationTagBitStringBuilder) setParent(contract BACnetApplicationTagContract) {
	b.BACnetApplicationTagContract = contract
}

func (b *_BACnetApplicationTagBitStringBuilder) WithMandatoryFields(payload BACnetTagPayloadBitString) BACnetApplicationTagBitStringBuilder {
	return b.WithPayload(payload)
}

func (b *_BACnetApplicationTagBitStringBuilder) WithPayload(payload BACnetTagPayloadBitString) BACnetApplicationTagBitStringBuilder {
	b.Payload = payload
	return b
}

func (b *_BACnetApplicationTagBitStringBuilder) WithPayloadBuilder(builderSupplier func(BACnetTagPayloadBitStringBuilder) BACnetTagPayloadBitStringBuilder) BACnetApplicationTagBitStringBuilder {
	builder := builderSupplier(b.Payload.CreateBACnetTagPayloadBitStringBuilder())
	var err error
	b.Payload, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagPayloadBitStringBuilder failed"))
	}
	return b
}

func (b *_BACnetApplicationTagBitStringBuilder) Build() (BACnetApplicationTagBitString, error) {
	if b.Payload == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'payload' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetApplicationTagBitString.deepCopy(), nil
}

func (b *_BACnetApplicationTagBitStringBuilder) MustBuild() BACnetApplicationTagBitString {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetApplicationTagBitStringBuilder) Done() BACnetApplicationTagBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetApplicationTagBuilder().(*_BACnetApplicationTagBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetApplicationTagBitStringBuilder) buildForBACnetApplicationTag() (BACnetApplicationTag, error) {
	return b.Build()
}

func (b *_BACnetApplicationTagBitStringBuilder) DeepCopy() any {
	_copy := b.CreateBACnetApplicationTagBitStringBuilder().(*_BACnetApplicationTagBitStringBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetApplicationTagBitStringBuilder creates a BACnetApplicationTagBitStringBuilder
func (b *_BACnetApplicationTagBitString) CreateBACnetApplicationTagBitStringBuilder() BACnetApplicationTagBitStringBuilder {
	if b == nil {
		return NewBACnetApplicationTagBitStringBuilder()
	}
	return &_BACnetApplicationTagBitStringBuilder{_BACnetApplicationTagBitString: b.deepCopy()}
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

func (m *_BACnetApplicationTagBitString) GetParent() BACnetApplicationTagContract {
	return m.BACnetApplicationTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetApplicationTagBitString) GetPayload() BACnetTagPayloadBitString {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetApplicationTagBitString(structType any) BACnetApplicationTagBitString {
	if casted, ok := structType.(BACnetApplicationTagBitString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetApplicationTagBitString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetApplicationTagBitString) GetTypeName() string {
	return "BACnetApplicationTagBitString"
}

func (m *_BACnetApplicationTagBitString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetApplicationTagContract.(*_BACnetApplicationTag).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetApplicationTagBitString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetApplicationTagBitString) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetApplicationTag, header BACnetTagHeader) (__bACnetApplicationTagBitString BACnetApplicationTagBitString, err error) {
	m.BACnetApplicationTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetApplicationTagBitString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetApplicationTagBitString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[BACnetTagPayloadBitString](ctx, "payload", ReadComplex[BACnetTagPayloadBitString](BACnetTagPayloadBitStringParseWithBufferProducer((uint32)(header.GetActualLength())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagBitString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetApplicationTagBitString")
	}

	return m, nil
}

func (m *_BACnetApplicationTagBitString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetApplicationTagBitString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagBitString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetApplicationTagBitString")
		}

		if err := WriteSimpleField[BACnetTagPayloadBitString](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadBitString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagBitString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetApplicationTagBitString")
		}
		return nil
	}
	return m.BACnetApplicationTagContract.(*_BACnetApplicationTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetApplicationTagBitString) IsBACnetApplicationTagBitString() {}

func (m *_BACnetApplicationTagBitString) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetApplicationTagBitString) deepCopy() *_BACnetApplicationTagBitString {
	if m == nil {
		return nil
	}
	_BACnetApplicationTagBitStringCopy := &_BACnetApplicationTagBitString{
		m.BACnetApplicationTagContract.(*_BACnetApplicationTag).deepCopy(),
		utils.DeepCopy[BACnetTagPayloadBitString](m.Payload),
	}
	m.BACnetApplicationTagContract.(*_BACnetApplicationTag)._SubType = m
	return _BACnetApplicationTagBitStringCopy
}

func (m *_BACnetApplicationTagBitString) String() string {
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
