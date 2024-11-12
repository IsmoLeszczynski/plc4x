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

// BACnetContextTagDate is the corresponding interface of BACnetContextTagDate
type BACnetContextTagDate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetContextTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadDate
	// IsBACnetContextTagDate is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetContextTagDate()
	// CreateBuilder creates a BACnetContextTagDateBuilder
	CreateBACnetContextTagDateBuilder() BACnetContextTagDateBuilder
}

// _BACnetContextTagDate is the data-structure of this message
type _BACnetContextTagDate struct {
	BACnetContextTagContract
	Payload BACnetTagPayloadDate
}

var _ BACnetContextTagDate = (*_BACnetContextTagDate)(nil)
var _ BACnetContextTagRequirements = (*_BACnetContextTagDate)(nil)

// NewBACnetContextTagDate factory function for _BACnetContextTagDate
func NewBACnetContextTagDate(header BACnetTagHeader, payload BACnetTagPayloadDate, tagNumberArgument uint8) *_BACnetContextTagDate {
	if payload == nil {
		panic("payload of type BACnetTagPayloadDate for BACnetContextTagDate must not be nil")
	}
	_result := &_BACnetContextTagDate{
		BACnetContextTagContract: NewBACnetContextTag(header, tagNumberArgument),
		Payload:                  payload,
	}
	_result.BACnetContextTagContract.(*_BACnetContextTag)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetContextTagDateBuilder is a builder for BACnetContextTagDate
type BACnetContextTagDateBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(payload BACnetTagPayloadDate) BACnetContextTagDateBuilder
	// WithPayload adds Payload (property field)
	WithPayload(BACnetTagPayloadDate) BACnetContextTagDateBuilder
	// WithPayloadBuilder adds Payload (property field) which is build by the builder
	WithPayloadBuilder(func(BACnetTagPayloadDateBuilder) BACnetTagPayloadDateBuilder) BACnetContextTagDateBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetContextTagBuilder
	// Build builds the BACnetContextTagDate or returns an error if something is wrong
	Build() (BACnetContextTagDate, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetContextTagDate
}

// NewBACnetContextTagDateBuilder() creates a BACnetContextTagDateBuilder
func NewBACnetContextTagDateBuilder() BACnetContextTagDateBuilder {
	return &_BACnetContextTagDateBuilder{_BACnetContextTagDate: new(_BACnetContextTagDate)}
}

type _BACnetContextTagDateBuilder struct {
	*_BACnetContextTagDate

	parentBuilder *_BACnetContextTagBuilder

	err *utils.MultiError
}

var _ (BACnetContextTagDateBuilder) = (*_BACnetContextTagDateBuilder)(nil)

func (b *_BACnetContextTagDateBuilder) setParent(contract BACnetContextTagContract) {
	b.BACnetContextTagContract = contract
}

func (b *_BACnetContextTagDateBuilder) WithMandatoryFields(payload BACnetTagPayloadDate) BACnetContextTagDateBuilder {
	return b.WithPayload(payload)
}

func (b *_BACnetContextTagDateBuilder) WithPayload(payload BACnetTagPayloadDate) BACnetContextTagDateBuilder {
	b.Payload = payload
	return b
}

func (b *_BACnetContextTagDateBuilder) WithPayloadBuilder(builderSupplier func(BACnetTagPayloadDateBuilder) BACnetTagPayloadDateBuilder) BACnetContextTagDateBuilder {
	builder := builderSupplier(b.Payload.CreateBACnetTagPayloadDateBuilder())
	var err error
	b.Payload, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagPayloadDateBuilder failed"))
	}
	return b
}

func (b *_BACnetContextTagDateBuilder) Build() (BACnetContextTagDate, error) {
	if b.Payload == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'payload' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetContextTagDate.deepCopy(), nil
}

func (b *_BACnetContextTagDateBuilder) MustBuild() BACnetContextTagDate {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetContextTagDateBuilder) Done() BACnetContextTagBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetContextTagBuilder().(*_BACnetContextTagBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetContextTagDateBuilder) buildForBACnetContextTag() (BACnetContextTag, error) {
	return b.Build()
}

func (b *_BACnetContextTagDateBuilder) DeepCopy() any {
	_copy := b.CreateBACnetContextTagDateBuilder().(*_BACnetContextTagDateBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetContextTagDateBuilder creates a BACnetContextTagDateBuilder
func (b *_BACnetContextTagDate) CreateBACnetContextTagDateBuilder() BACnetContextTagDateBuilder {
	if b == nil {
		return NewBACnetContextTagDateBuilder()
	}
	return &_BACnetContextTagDateBuilder{_BACnetContextTagDate: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagDate) GetDataType() BACnetDataType {
	return BACnetDataType_DATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagDate) GetParent() BACnetContextTagContract {
	return m.BACnetContextTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTagDate) GetPayload() BACnetTagPayloadDate {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetContextTagDate(structType any) BACnetContextTagDate {
	if casted, ok := structType.(BACnetContextTagDate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagDate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagDate) GetTypeName() string {
	return "BACnetContextTagDate"
}

func (m *_BACnetContextTagDate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetContextTagContract.(*_BACnetContextTag).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetContextTagDate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetContextTagDate) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetContextTag, tagNumberArgument uint8, dataType BACnetDataType) (__bACnetContextTagDate BACnetContextTagDate, err error) {
	m.BACnetContextTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagDate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[BACnetTagPayloadDate](ctx, "payload", ReadComplex[BACnetTagPayloadDate](BACnetTagPayloadDateParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	if closeErr := readBuffer.CloseContext("BACnetContextTagDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagDate")
	}

	return m, nil
}

func (m *_BACnetContextTagDate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetContextTagDate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagDate")
		}

		if err := WriteSimpleField[BACnetTagPayloadDate](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadDate](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagDate")
		}
		return nil
	}
	return m.BACnetContextTagContract.(*_BACnetContextTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetContextTagDate) IsBACnetContextTagDate() {}

func (m *_BACnetContextTagDate) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetContextTagDate) deepCopy() *_BACnetContextTagDate {
	if m == nil {
		return nil
	}
	_BACnetContextTagDateCopy := &_BACnetContextTagDate{
		m.BACnetContextTagContract.(*_BACnetContextTag).deepCopy(),
		utils.DeepCopy[BACnetTagPayloadDate](m.Payload),
	}
	m.BACnetContextTagContract.(*_BACnetContextTag)._SubType = m
	return _BACnetContextTagDateCopy
}

func (m *_BACnetContextTagDate) String() string {
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
