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

// MediaTransportControlDataRepeatOnOff is the corresponding interface of MediaTransportControlDataRepeatOnOff
type MediaTransportControlDataRepeatOnOff interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	MediaTransportControlData
	// GetRepeatType returns RepeatType (property field)
	GetRepeatType() byte
	// GetIsOff returns IsOff (virtual field)
	GetIsOff() bool
	// GetIsRepeatCurrent returns IsRepeatCurrent (virtual field)
	GetIsRepeatCurrent() bool
	// GetIsRepeatTracks returns IsRepeatTracks (virtual field)
	GetIsRepeatTracks() bool
	// IsMediaTransportControlDataRepeatOnOff is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMediaTransportControlDataRepeatOnOff()
	// CreateBuilder creates a MediaTransportControlDataRepeatOnOffBuilder
	CreateMediaTransportControlDataRepeatOnOffBuilder() MediaTransportControlDataRepeatOnOffBuilder
}

// _MediaTransportControlDataRepeatOnOff is the data-structure of this message
type _MediaTransportControlDataRepeatOnOff struct {
	MediaTransportControlDataContract
	RepeatType byte
}

var _ MediaTransportControlDataRepeatOnOff = (*_MediaTransportControlDataRepeatOnOff)(nil)
var _ MediaTransportControlDataRequirements = (*_MediaTransportControlDataRepeatOnOff)(nil)

// NewMediaTransportControlDataRepeatOnOff factory function for _MediaTransportControlDataRepeatOnOff
func NewMediaTransportControlDataRepeatOnOff(commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte, repeatType byte) *_MediaTransportControlDataRepeatOnOff {
	_result := &_MediaTransportControlDataRepeatOnOff{
		MediaTransportControlDataContract: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
		RepeatType:                        repeatType,
	}
	_result.MediaTransportControlDataContract.(*_MediaTransportControlData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MediaTransportControlDataRepeatOnOffBuilder is a builder for MediaTransportControlDataRepeatOnOff
type MediaTransportControlDataRepeatOnOffBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(repeatType byte) MediaTransportControlDataRepeatOnOffBuilder
	// WithRepeatType adds RepeatType (property field)
	WithRepeatType(byte) MediaTransportControlDataRepeatOnOffBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() MediaTransportControlDataBuilder
	// Build builds the MediaTransportControlDataRepeatOnOff or returns an error if something is wrong
	Build() (MediaTransportControlDataRepeatOnOff, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MediaTransportControlDataRepeatOnOff
}

// NewMediaTransportControlDataRepeatOnOffBuilder() creates a MediaTransportControlDataRepeatOnOffBuilder
func NewMediaTransportControlDataRepeatOnOffBuilder() MediaTransportControlDataRepeatOnOffBuilder {
	return &_MediaTransportControlDataRepeatOnOffBuilder{_MediaTransportControlDataRepeatOnOff: new(_MediaTransportControlDataRepeatOnOff)}
}

type _MediaTransportControlDataRepeatOnOffBuilder struct {
	*_MediaTransportControlDataRepeatOnOff

	parentBuilder *_MediaTransportControlDataBuilder

	err *utils.MultiError
}

var _ (MediaTransportControlDataRepeatOnOffBuilder) = (*_MediaTransportControlDataRepeatOnOffBuilder)(nil)

func (b *_MediaTransportControlDataRepeatOnOffBuilder) setParent(contract MediaTransportControlDataContract) {
	b.MediaTransportControlDataContract = contract
}

func (b *_MediaTransportControlDataRepeatOnOffBuilder) WithMandatoryFields(repeatType byte) MediaTransportControlDataRepeatOnOffBuilder {
	return b.WithRepeatType(repeatType)
}

func (b *_MediaTransportControlDataRepeatOnOffBuilder) WithRepeatType(repeatType byte) MediaTransportControlDataRepeatOnOffBuilder {
	b.RepeatType = repeatType
	return b
}

func (b *_MediaTransportControlDataRepeatOnOffBuilder) Build() (MediaTransportControlDataRepeatOnOff, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MediaTransportControlDataRepeatOnOff.deepCopy(), nil
}

func (b *_MediaTransportControlDataRepeatOnOffBuilder) MustBuild() MediaTransportControlDataRepeatOnOff {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_MediaTransportControlDataRepeatOnOffBuilder) Done() MediaTransportControlDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewMediaTransportControlDataBuilder().(*_MediaTransportControlDataBuilder)
	}
	return b.parentBuilder
}

func (b *_MediaTransportControlDataRepeatOnOffBuilder) buildForMediaTransportControlData() (MediaTransportControlData, error) {
	return b.Build()
}

func (b *_MediaTransportControlDataRepeatOnOffBuilder) DeepCopy() any {
	_copy := b.CreateMediaTransportControlDataRepeatOnOffBuilder().(*_MediaTransportControlDataRepeatOnOffBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMediaTransportControlDataRepeatOnOffBuilder creates a MediaTransportControlDataRepeatOnOffBuilder
func (b *_MediaTransportControlDataRepeatOnOff) CreateMediaTransportControlDataRepeatOnOffBuilder() MediaTransportControlDataRepeatOnOffBuilder {
	if b == nil {
		return NewMediaTransportControlDataRepeatOnOffBuilder()
	}
	return &_MediaTransportControlDataRepeatOnOffBuilder{_MediaTransportControlDataRepeatOnOff: b.deepCopy()}
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

func (m *_MediaTransportControlDataRepeatOnOff) GetParent() MediaTransportControlDataContract {
	return m.MediaTransportControlDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataRepeatOnOff) GetRepeatType() byte {
	return m.RepeatType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_MediaTransportControlDataRepeatOnOff) GetIsOff() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetRepeatType()) == (0x00)))
}

func (m *_MediaTransportControlDataRepeatOnOff) GetIsRepeatCurrent() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool(bool((m.GetRepeatType()) > (0x00))) && bool(bool((m.GetRepeatType()) <= (0xFE))))
}

func (m *_MediaTransportControlDataRepeatOnOff) GetIsRepeatTracks() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetRepeatType()) >= (0xFE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataRepeatOnOff(structType any) MediaTransportControlDataRepeatOnOff {
	if casted, ok := structType.(MediaTransportControlDataRepeatOnOff); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataRepeatOnOff); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataRepeatOnOff) GetTypeName() string {
	return "MediaTransportControlDataRepeatOnOff"
}

func (m *_MediaTransportControlDataRepeatOnOff) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MediaTransportControlDataContract.(*_MediaTransportControlData).getLengthInBits(ctx))

	// Simple field (repeatType)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_MediaTransportControlDataRepeatOnOff) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MediaTransportControlDataRepeatOnOff) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MediaTransportControlData) (__mediaTransportControlDataRepeatOnOff MediaTransportControlDataRepeatOnOff, err error) {
	m.MediaTransportControlDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlDataRepeatOnOff"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataRepeatOnOff")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	repeatType, err := ReadSimpleField(ctx, "repeatType", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'repeatType' field"))
	}
	m.RepeatType = repeatType

	isOff, err := ReadVirtualField[bool](ctx, "isOff", (*bool)(nil), bool((repeatType) == (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isOff' field"))
	}
	_ = isOff

	isRepeatCurrent, err := ReadVirtualField[bool](ctx, "isRepeatCurrent", (*bool)(nil), bool(bool((repeatType) > (0x00))) && bool(bool((repeatType) <= (0xFE))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isRepeatCurrent' field"))
	}
	_ = isRepeatCurrent

	isRepeatTracks, err := ReadVirtualField[bool](ctx, "isRepeatTracks", (*bool)(nil), bool((repeatType) >= (0xFE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isRepeatTracks' field"))
	}
	_ = isRepeatTracks

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataRepeatOnOff"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataRepeatOnOff")
	}

	return m, nil
}

func (m *_MediaTransportControlDataRepeatOnOff) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataRepeatOnOff) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataRepeatOnOff"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataRepeatOnOff")
		}

		if err := WriteSimpleField[byte](ctx, "repeatType", m.GetRepeatType(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'repeatType' field")
		}
		// Virtual field
		isOff := m.GetIsOff()
		_ = isOff
		if _isOffErr := writeBuffer.WriteVirtual(ctx, "isOff", m.GetIsOff()); _isOffErr != nil {
			return errors.Wrap(_isOffErr, "Error serializing 'isOff' field")
		}
		// Virtual field
		isRepeatCurrent := m.GetIsRepeatCurrent()
		_ = isRepeatCurrent
		if _isRepeatCurrentErr := writeBuffer.WriteVirtual(ctx, "isRepeatCurrent", m.GetIsRepeatCurrent()); _isRepeatCurrentErr != nil {
			return errors.Wrap(_isRepeatCurrentErr, "Error serializing 'isRepeatCurrent' field")
		}
		// Virtual field
		isRepeatTracks := m.GetIsRepeatTracks()
		_ = isRepeatTracks
		if _isRepeatTracksErr := writeBuffer.WriteVirtual(ctx, "isRepeatTracks", m.GetIsRepeatTracks()); _isRepeatTracksErr != nil {
			return errors.Wrap(_isRepeatTracksErr, "Error serializing 'isRepeatTracks' field")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataRepeatOnOff"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataRepeatOnOff")
		}
		return nil
	}
	return m.MediaTransportControlDataContract.(*_MediaTransportControlData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataRepeatOnOff) IsMediaTransportControlDataRepeatOnOff() {}

func (m *_MediaTransportControlDataRepeatOnOff) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MediaTransportControlDataRepeatOnOff) deepCopy() *_MediaTransportControlDataRepeatOnOff {
	if m == nil {
		return nil
	}
	_MediaTransportControlDataRepeatOnOffCopy := &_MediaTransportControlDataRepeatOnOff{
		m.MediaTransportControlDataContract.(*_MediaTransportControlData).deepCopy(),
		m.RepeatType,
	}
	m.MediaTransportControlDataContract.(*_MediaTransportControlData)._SubType = m
	return _MediaTransportControlDataRepeatOnOffCopy
}

func (m *_MediaTransportControlDataRepeatOnOff) String() string {
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
