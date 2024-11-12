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

// APDUAbort is the corresponding interface of APDUAbort
type APDUAbort interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	APDU
	// GetServer returns Server (property field)
	GetServer() bool
	// GetOriginalInvokeId returns OriginalInvokeId (property field)
	GetOriginalInvokeId() uint8
	// GetAbortReason returns AbortReason (property field)
	GetAbortReason() BACnetAbortReasonTagged
	// IsAPDUAbort is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAPDUAbort()
	// CreateBuilder creates a APDUAbortBuilder
	CreateAPDUAbortBuilder() APDUAbortBuilder
}

// _APDUAbort is the data-structure of this message
type _APDUAbort struct {
	APDUContract
	Server           bool
	OriginalInvokeId uint8
	AbortReason      BACnetAbortReasonTagged
	// Reserved Fields
	reservedField0 *uint8
}

var _ APDUAbort = (*_APDUAbort)(nil)
var _ APDURequirements = (*_APDUAbort)(nil)

// NewAPDUAbort factory function for _APDUAbort
func NewAPDUAbort(server bool, originalInvokeId uint8, abortReason BACnetAbortReasonTagged, apduLength uint16) *_APDUAbort {
	if abortReason == nil {
		panic("abortReason of type BACnetAbortReasonTagged for APDUAbort must not be nil")
	}
	_result := &_APDUAbort{
		APDUContract:     NewAPDU(apduLength),
		Server:           server,
		OriginalInvokeId: originalInvokeId,
		AbortReason:      abortReason,
	}
	_result.APDUContract.(*_APDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// APDUAbortBuilder is a builder for APDUAbort
type APDUAbortBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(server bool, originalInvokeId uint8, abortReason BACnetAbortReasonTagged) APDUAbortBuilder
	// WithServer adds Server (property field)
	WithServer(bool) APDUAbortBuilder
	// WithOriginalInvokeId adds OriginalInvokeId (property field)
	WithOriginalInvokeId(uint8) APDUAbortBuilder
	// WithAbortReason adds AbortReason (property field)
	WithAbortReason(BACnetAbortReasonTagged) APDUAbortBuilder
	// WithAbortReasonBuilder adds AbortReason (property field) which is build by the builder
	WithAbortReasonBuilder(func(BACnetAbortReasonTaggedBuilder) BACnetAbortReasonTaggedBuilder) APDUAbortBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() APDUBuilder
	// Build builds the APDUAbort or returns an error if something is wrong
	Build() (APDUAbort, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() APDUAbort
}

// NewAPDUAbortBuilder() creates a APDUAbortBuilder
func NewAPDUAbortBuilder() APDUAbortBuilder {
	return &_APDUAbortBuilder{_APDUAbort: new(_APDUAbort)}
}

type _APDUAbortBuilder struct {
	*_APDUAbort

	parentBuilder *_APDUBuilder

	err *utils.MultiError
}

var _ (APDUAbortBuilder) = (*_APDUAbortBuilder)(nil)

func (b *_APDUAbortBuilder) setParent(contract APDUContract) {
	b.APDUContract = contract
}

func (b *_APDUAbortBuilder) WithMandatoryFields(server bool, originalInvokeId uint8, abortReason BACnetAbortReasonTagged) APDUAbortBuilder {
	return b.WithServer(server).WithOriginalInvokeId(originalInvokeId).WithAbortReason(abortReason)
}

func (b *_APDUAbortBuilder) WithServer(server bool) APDUAbortBuilder {
	b.Server = server
	return b
}

func (b *_APDUAbortBuilder) WithOriginalInvokeId(originalInvokeId uint8) APDUAbortBuilder {
	b.OriginalInvokeId = originalInvokeId
	return b
}

func (b *_APDUAbortBuilder) WithAbortReason(abortReason BACnetAbortReasonTagged) APDUAbortBuilder {
	b.AbortReason = abortReason
	return b
}

func (b *_APDUAbortBuilder) WithAbortReasonBuilder(builderSupplier func(BACnetAbortReasonTaggedBuilder) BACnetAbortReasonTaggedBuilder) APDUAbortBuilder {
	builder := builderSupplier(b.AbortReason.CreateBACnetAbortReasonTaggedBuilder())
	var err error
	b.AbortReason, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetAbortReasonTaggedBuilder failed"))
	}
	return b
}

func (b *_APDUAbortBuilder) Build() (APDUAbort, error) {
	if b.AbortReason == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'abortReason' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._APDUAbort.deepCopy(), nil
}

func (b *_APDUAbortBuilder) MustBuild() APDUAbort {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_APDUAbortBuilder) Done() APDUBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewAPDUBuilder().(*_APDUBuilder)
	}
	return b.parentBuilder
}

func (b *_APDUAbortBuilder) buildForAPDU() (APDU, error) {
	return b.Build()
}

func (b *_APDUAbortBuilder) DeepCopy() any {
	_copy := b.CreateAPDUAbortBuilder().(*_APDUAbortBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAPDUAbortBuilder creates a APDUAbortBuilder
func (b *_APDUAbort) CreateAPDUAbortBuilder() APDUAbortBuilder {
	if b == nil {
		return NewAPDUAbortBuilder()
	}
	return &_APDUAbortBuilder{_APDUAbort: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_APDUAbort) GetApduType() ApduType {
	return ApduType_ABORT_PDU
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_APDUAbort) GetParent() APDUContract {
	return m.APDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_APDUAbort) GetServer() bool {
	return m.Server
}

func (m *_APDUAbort) GetOriginalInvokeId() uint8 {
	return m.OriginalInvokeId
}

func (m *_APDUAbort) GetAbortReason() BACnetAbortReasonTagged {
	return m.AbortReason
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAPDUAbort(structType any) APDUAbort {
	if casted, ok := structType.(APDUAbort); ok {
		return casted
	}
	if casted, ok := structType.(*APDUAbort); ok {
		return *casted
	}
	return nil
}

func (m *_APDUAbort) GetTypeName() string {
	return "APDUAbort"
}

func (m *_APDUAbort) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.APDUContract.(*_APDU).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 3

	// Simple field (server)
	lengthInBits += 1

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Simple field (abortReason)
	lengthInBits += m.AbortReason.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_APDUAbort) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_APDUAbort) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_APDU, apduLength uint16) (__aPDUAbort APDUAbort, err error) {
	m.APDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("APDUAbort"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for APDUAbort")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(3)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	server, err := ReadSimpleField(ctx, "server", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'server' field"))
	}
	m.Server = server

	originalInvokeId, err := ReadSimpleField(ctx, "originalInvokeId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'originalInvokeId' field"))
	}
	m.OriginalInvokeId = originalInvokeId

	abortReason, err := ReadSimpleField[BACnetAbortReasonTagged](ctx, "abortReason", ReadComplex[BACnetAbortReasonTagged](BACnetAbortReasonTaggedParseWithBufferProducer((uint32)(uint32(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'abortReason' field"))
	}
	m.AbortReason = abortReason

	if closeErr := readBuffer.CloseContext("APDUAbort"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for APDUAbort")
	}

	return m, nil
}

func (m *_APDUAbort) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_APDUAbort) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUAbort"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for APDUAbort")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 3)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "server", m.GetServer(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'server' field")
		}

		if err := WriteSimpleField[uint8](ctx, "originalInvokeId", m.GetOriginalInvokeId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'originalInvokeId' field")
		}

		if err := WriteSimpleField[BACnetAbortReasonTagged](ctx, "abortReason", m.GetAbortReason(), WriteComplex[BACnetAbortReasonTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'abortReason' field")
		}

		if popErr := writeBuffer.PopContext("APDUAbort"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for APDUAbort")
		}
		return nil
	}
	return m.APDUContract.(*_APDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_APDUAbort) IsAPDUAbort() {}

func (m *_APDUAbort) DeepCopy() any {
	return m.deepCopy()
}

func (m *_APDUAbort) deepCopy() *_APDUAbort {
	if m == nil {
		return nil
	}
	_APDUAbortCopy := &_APDUAbort{
		m.APDUContract.(*_APDU).deepCopy(),
		m.Server,
		m.OriginalInvokeId,
		utils.DeepCopy[BACnetAbortReasonTagged](m.AbortReason),
		m.reservedField0,
	}
	m.APDUContract.(*_APDU)._SubType = m
	return _APDUAbortCopy
}

func (m *_APDUAbort) String() string {
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
