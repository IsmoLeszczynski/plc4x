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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BVLCForwardedNPDU is the corresponding interface of BVLCForwardedNPDU
type BVLCForwardedNPDU interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BVLC
	// GetIp returns Ip (property field)
	GetIp() []uint8
	// GetPort returns Port (property field)
	GetPort() uint16
	// GetNpdu returns Npdu (property field)
	GetNpdu() NPDU
	// IsBVLCForwardedNPDU is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBVLCForwardedNPDU()
	// CreateBuilder creates a BVLCForwardedNPDUBuilder
	CreateBVLCForwardedNPDUBuilder() BVLCForwardedNPDUBuilder
}

// _BVLCForwardedNPDU is the data-structure of this message
type _BVLCForwardedNPDU struct {
	BVLCContract
	Ip   []uint8
	Port uint16
	Npdu NPDU

	// Arguments.
	BvlcPayloadLength uint16
}

var _ BVLCForwardedNPDU = (*_BVLCForwardedNPDU)(nil)
var _ BVLCRequirements = (*_BVLCForwardedNPDU)(nil)

// NewBVLCForwardedNPDU factory function for _BVLCForwardedNPDU
func NewBVLCForwardedNPDU(ip []uint8, port uint16, npdu NPDU, bvlcPayloadLength uint16) *_BVLCForwardedNPDU {
	if npdu == nil {
		panic("npdu of type NPDU for BVLCForwardedNPDU must not be nil")
	}
	_result := &_BVLCForwardedNPDU{
		BVLCContract: NewBVLC(),
		Ip:           ip,
		Port:         port,
		Npdu:         npdu,
	}
	_result.BVLCContract.(*_BVLC)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BVLCForwardedNPDUBuilder is a builder for BVLCForwardedNPDU
type BVLCForwardedNPDUBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(ip []uint8, port uint16, npdu NPDU) BVLCForwardedNPDUBuilder
	// WithIp adds Ip (property field)
	WithIp(...uint8) BVLCForwardedNPDUBuilder
	// WithPort adds Port (property field)
	WithPort(uint16) BVLCForwardedNPDUBuilder
	// WithNpdu adds Npdu (property field)
	WithNpdu(NPDU) BVLCForwardedNPDUBuilder
	// WithNpduBuilder adds Npdu (property field) which is build by the builder
	WithNpduBuilder(func(NPDUBuilder) NPDUBuilder) BVLCForwardedNPDUBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BVLCBuilder
	// Build builds the BVLCForwardedNPDU or returns an error if something is wrong
	Build() (BVLCForwardedNPDU, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BVLCForwardedNPDU
}

// NewBVLCForwardedNPDUBuilder() creates a BVLCForwardedNPDUBuilder
func NewBVLCForwardedNPDUBuilder() BVLCForwardedNPDUBuilder {
	return &_BVLCForwardedNPDUBuilder{_BVLCForwardedNPDU: new(_BVLCForwardedNPDU)}
}

type _BVLCForwardedNPDUBuilder struct {
	*_BVLCForwardedNPDU

	parentBuilder *_BVLCBuilder

	err *utils.MultiError
}

var _ (BVLCForwardedNPDUBuilder) = (*_BVLCForwardedNPDUBuilder)(nil)

func (b *_BVLCForwardedNPDUBuilder) setParent(contract BVLCContract) {
	b.BVLCContract = contract
}

func (b *_BVLCForwardedNPDUBuilder) WithMandatoryFields(ip []uint8, port uint16, npdu NPDU) BVLCForwardedNPDUBuilder {
	return b.WithIp(ip...).WithPort(port).WithNpdu(npdu)
}

func (b *_BVLCForwardedNPDUBuilder) WithIp(ip ...uint8) BVLCForwardedNPDUBuilder {
	b.Ip = ip
	return b
}

func (b *_BVLCForwardedNPDUBuilder) WithPort(port uint16) BVLCForwardedNPDUBuilder {
	b.Port = port
	return b
}

func (b *_BVLCForwardedNPDUBuilder) WithNpdu(npdu NPDU) BVLCForwardedNPDUBuilder {
	b.Npdu = npdu
	return b
}

func (b *_BVLCForwardedNPDUBuilder) WithNpduBuilder(builderSupplier func(NPDUBuilder) NPDUBuilder) BVLCForwardedNPDUBuilder {
	builder := builderSupplier(b.Npdu.CreateNPDUBuilder())
	var err error
	b.Npdu, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NPDUBuilder failed"))
	}
	return b
}

func (b *_BVLCForwardedNPDUBuilder) Build() (BVLCForwardedNPDU, error) {
	if b.Npdu == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'npdu' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BVLCForwardedNPDU.deepCopy(), nil
}

func (b *_BVLCForwardedNPDUBuilder) MustBuild() BVLCForwardedNPDU {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BVLCForwardedNPDUBuilder) Done() BVLCBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBVLCBuilder().(*_BVLCBuilder)
	}
	return b.parentBuilder
}

func (b *_BVLCForwardedNPDUBuilder) buildForBVLC() (BVLC, error) {
	return b.Build()
}

func (b *_BVLCForwardedNPDUBuilder) DeepCopy() any {
	_copy := b.CreateBVLCForwardedNPDUBuilder().(*_BVLCForwardedNPDUBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBVLCForwardedNPDUBuilder creates a BVLCForwardedNPDUBuilder
func (b *_BVLCForwardedNPDU) CreateBVLCForwardedNPDUBuilder() BVLCForwardedNPDUBuilder {
	if b == nil {
		return NewBVLCForwardedNPDUBuilder()
	}
	return &_BVLCForwardedNPDUBuilder{_BVLCForwardedNPDU: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BVLCForwardedNPDU) GetBvlcFunction() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BVLCForwardedNPDU) GetParent() BVLCContract {
	return m.BVLCContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BVLCForwardedNPDU) GetIp() []uint8 {
	return m.Ip
}

func (m *_BVLCForwardedNPDU) GetPort() uint16 {
	return m.Port
}

func (m *_BVLCForwardedNPDU) GetNpdu() NPDU {
	return m.Npdu
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBVLCForwardedNPDU(structType any) BVLCForwardedNPDU {
	if casted, ok := structType.(BVLCForwardedNPDU); ok {
		return casted
	}
	if casted, ok := structType.(*BVLCForwardedNPDU); ok {
		return *casted
	}
	return nil
}

func (m *_BVLCForwardedNPDU) GetTypeName() string {
	return "BVLCForwardedNPDU"
}

func (m *_BVLCForwardedNPDU) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BVLCContract.(*_BVLC).getLengthInBits(ctx))

	// Array field
	if len(m.Ip) > 0 {
		lengthInBits += 8 * uint16(len(m.Ip))
	}

	// Simple field (port)
	lengthInBits += 16

	// Simple field (npdu)
	lengthInBits += m.Npdu.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BVLCForwardedNPDU) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BVLCForwardedNPDU) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BVLC, bvlcPayloadLength uint16) (__bVLCForwardedNPDU BVLCForwardedNPDU, err error) {
	m.BVLCContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLCForwardedNPDU"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCForwardedNPDU")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ip, err := ReadCountArrayField[uint8](ctx, "ip", ReadUnsignedByte(readBuffer, uint8(8)), uint64(int32(4)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ip' field"))
	}
	m.Ip = ip

	port, err := ReadSimpleField(ctx, "port", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'port' field"))
	}
	m.Port = port

	npdu, err := ReadSimpleField[NPDU](ctx, "npdu", ReadComplex[NPDU](NPDUParseWithBufferProducer((uint16)(uint16(bvlcPayloadLength)-uint16(uint16(6)))), readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'npdu' field"))
	}
	m.Npdu = npdu

	if closeErr := readBuffer.CloseContext("BVLCForwardedNPDU"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCForwardedNPDU")
	}

	return m, nil
}

func (m *_BVLCForwardedNPDU) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BVLCForwardedNPDU) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCForwardedNPDU"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BVLCForwardedNPDU")
		}

		if err := WriteSimpleTypeArrayField(ctx, "ip", m.GetIp(), WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'ip' field")
		}

		if err := WriteSimpleField[uint16](ctx, "port", m.GetPort(), WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'port' field")
		}

		if err := WriteSimpleField[NPDU](ctx, "npdu", m.GetNpdu(), WriteComplex[NPDU](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'npdu' field")
		}

		if popErr := writeBuffer.PopContext("BVLCForwardedNPDU"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BVLCForwardedNPDU")
		}
		return nil
	}
	return m.BVLCContract.(*_BVLC).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BVLCForwardedNPDU) GetBvlcPayloadLength() uint16 {
	return m.BvlcPayloadLength
}

//
////

func (m *_BVLCForwardedNPDU) IsBVLCForwardedNPDU() {}

func (m *_BVLCForwardedNPDU) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BVLCForwardedNPDU) deepCopy() *_BVLCForwardedNPDU {
	if m == nil {
		return nil
	}
	_BVLCForwardedNPDUCopy := &_BVLCForwardedNPDU{
		m.BVLCContract.(*_BVLC).deepCopy(),
		utils.DeepCopySlice[uint8, uint8](m.Ip),
		m.Port,
		utils.DeepCopy[NPDU](m.Npdu),
		m.BvlcPayloadLength,
	}
	m.BVLCContract.(*_BVLC)._SubType = m
	return _BVLCForwardedNPDUCopy
}

func (m *_BVLCForwardedNPDU) String() string {
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
