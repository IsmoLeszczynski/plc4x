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

// AdsDiscoveryBlock is the corresponding interface of AdsDiscoveryBlock
type AdsDiscoveryBlock interface {
	AdsDiscoveryBlockContract
	AdsDiscoveryBlockRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsAdsDiscoveryBlock is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsDiscoveryBlock()
	// CreateBuilder creates a AdsDiscoveryBlockBuilder
	CreateAdsDiscoveryBlockBuilder() AdsDiscoveryBlockBuilder
}

// AdsDiscoveryBlockContract provides a set of functions which can be overwritten by a sub struct
type AdsDiscoveryBlockContract interface {
	// IsAdsDiscoveryBlock is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsDiscoveryBlock()
	// CreateBuilder creates a AdsDiscoveryBlockBuilder
	CreateAdsDiscoveryBlockBuilder() AdsDiscoveryBlockBuilder
}

// AdsDiscoveryBlockRequirements provides a set of functions which need to be implemented by a sub struct
type AdsDiscoveryBlockRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetBlockType returns BlockType (discriminator field)
	GetBlockType() AdsDiscoveryBlockType
}

// _AdsDiscoveryBlock is the data-structure of this message
type _AdsDiscoveryBlock struct {
	_SubType interface {
		AdsDiscoveryBlockContract
		AdsDiscoveryBlockRequirements
	}
}

var _ AdsDiscoveryBlockContract = (*_AdsDiscoveryBlock)(nil)

// NewAdsDiscoveryBlock factory function for _AdsDiscoveryBlock
func NewAdsDiscoveryBlock() *_AdsDiscoveryBlock {
	return &_AdsDiscoveryBlock{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AdsDiscoveryBlockBuilder is a builder for AdsDiscoveryBlock
type AdsDiscoveryBlockBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() AdsDiscoveryBlockBuilder
	// AsAdsDiscoveryBlockStatus converts this build to a subType of AdsDiscoveryBlock. It is always possible to return to current builder using Done()
	AsAdsDiscoveryBlockStatus() AdsDiscoveryBlockStatusBuilder
	// AsAdsDiscoveryBlockPassword converts this build to a subType of AdsDiscoveryBlock. It is always possible to return to current builder using Done()
	AsAdsDiscoveryBlockPassword() AdsDiscoveryBlockPasswordBuilder
	// AsAdsDiscoveryBlockVersion converts this build to a subType of AdsDiscoveryBlock. It is always possible to return to current builder using Done()
	AsAdsDiscoveryBlockVersion() AdsDiscoveryBlockVersionBuilder
	// AsAdsDiscoveryBlockOsData converts this build to a subType of AdsDiscoveryBlock. It is always possible to return to current builder using Done()
	AsAdsDiscoveryBlockOsData() AdsDiscoveryBlockOsDataBuilder
	// AsAdsDiscoveryBlockHostName converts this build to a subType of AdsDiscoveryBlock. It is always possible to return to current builder using Done()
	AsAdsDiscoveryBlockHostName() AdsDiscoveryBlockHostNameBuilder
	// AsAdsDiscoveryBlockAmsNetId converts this build to a subType of AdsDiscoveryBlock. It is always possible to return to current builder using Done()
	AsAdsDiscoveryBlockAmsNetId() AdsDiscoveryBlockAmsNetIdBuilder
	// AsAdsDiscoveryBlockRouteName converts this build to a subType of AdsDiscoveryBlock. It is always possible to return to current builder using Done()
	AsAdsDiscoveryBlockRouteName() AdsDiscoveryBlockRouteNameBuilder
	// AsAdsDiscoveryBlockUserName converts this build to a subType of AdsDiscoveryBlock. It is always possible to return to current builder using Done()
	AsAdsDiscoveryBlockUserName() AdsDiscoveryBlockUserNameBuilder
	// AsAdsDiscoveryBlockFingerprint converts this build to a subType of AdsDiscoveryBlock. It is always possible to return to current builder using Done()
	AsAdsDiscoveryBlockFingerprint() AdsDiscoveryBlockFingerprintBuilder
	// Build builds the AdsDiscoveryBlock or returns an error if something is wrong
	PartialBuild() (AdsDiscoveryBlockContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() AdsDiscoveryBlockContract
	// Build builds the AdsDiscoveryBlock or returns an error if something is wrong
	Build() (AdsDiscoveryBlock, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AdsDiscoveryBlock
}

// NewAdsDiscoveryBlockBuilder() creates a AdsDiscoveryBlockBuilder
func NewAdsDiscoveryBlockBuilder() AdsDiscoveryBlockBuilder {
	return &_AdsDiscoveryBlockBuilder{_AdsDiscoveryBlock: new(_AdsDiscoveryBlock)}
}

type _AdsDiscoveryBlockChildBuilder interface {
	utils.Copyable
	setParent(AdsDiscoveryBlockContract)
	buildForAdsDiscoveryBlock() (AdsDiscoveryBlock, error)
}

type _AdsDiscoveryBlockBuilder struct {
	*_AdsDiscoveryBlock

	childBuilder _AdsDiscoveryBlockChildBuilder

	err *utils.MultiError
}

var _ (AdsDiscoveryBlockBuilder) = (*_AdsDiscoveryBlockBuilder)(nil)

func (b *_AdsDiscoveryBlockBuilder) WithMandatoryFields() AdsDiscoveryBlockBuilder {
	return b
}

func (b *_AdsDiscoveryBlockBuilder) PartialBuild() (AdsDiscoveryBlockContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AdsDiscoveryBlock.deepCopy(), nil
}

func (b *_AdsDiscoveryBlockBuilder) PartialMustBuild() AdsDiscoveryBlockContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AdsDiscoveryBlockBuilder) AsAdsDiscoveryBlockStatus() AdsDiscoveryBlockStatusBuilder {
	if cb, ok := b.childBuilder.(AdsDiscoveryBlockStatusBuilder); ok {
		return cb
	}
	cb := NewAdsDiscoveryBlockStatusBuilder().(*_AdsDiscoveryBlockStatusBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_AdsDiscoveryBlockBuilder) AsAdsDiscoveryBlockPassword() AdsDiscoveryBlockPasswordBuilder {
	if cb, ok := b.childBuilder.(AdsDiscoveryBlockPasswordBuilder); ok {
		return cb
	}
	cb := NewAdsDiscoveryBlockPasswordBuilder().(*_AdsDiscoveryBlockPasswordBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_AdsDiscoveryBlockBuilder) AsAdsDiscoveryBlockVersion() AdsDiscoveryBlockVersionBuilder {
	if cb, ok := b.childBuilder.(AdsDiscoveryBlockVersionBuilder); ok {
		return cb
	}
	cb := NewAdsDiscoveryBlockVersionBuilder().(*_AdsDiscoveryBlockVersionBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_AdsDiscoveryBlockBuilder) AsAdsDiscoveryBlockOsData() AdsDiscoveryBlockOsDataBuilder {
	if cb, ok := b.childBuilder.(AdsDiscoveryBlockOsDataBuilder); ok {
		return cb
	}
	cb := NewAdsDiscoveryBlockOsDataBuilder().(*_AdsDiscoveryBlockOsDataBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_AdsDiscoveryBlockBuilder) AsAdsDiscoveryBlockHostName() AdsDiscoveryBlockHostNameBuilder {
	if cb, ok := b.childBuilder.(AdsDiscoveryBlockHostNameBuilder); ok {
		return cb
	}
	cb := NewAdsDiscoveryBlockHostNameBuilder().(*_AdsDiscoveryBlockHostNameBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_AdsDiscoveryBlockBuilder) AsAdsDiscoveryBlockAmsNetId() AdsDiscoveryBlockAmsNetIdBuilder {
	if cb, ok := b.childBuilder.(AdsDiscoveryBlockAmsNetIdBuilder); ok {
		return cb
	}
	cb := NewAdsDiscoveryBlockAmsNetIdBuilder().(*_AdsDiscoveryBlockAmsNetIdBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_AdsDiscoveryBlockBuilder) AsAdsDiscoveryBlockRouteName() AdsDiscoveryBlockRouteNameBuilder {
	if cb, ok := b.childBuilder.(AdsDiscoveryBlockRouteNameBuilder); ok {
		return cb
	}
	cb := NewAdsDiscoveryBlockRouteNameBuilder().(*_AdsDiscoveryBlockRouteNameBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_AdsDiscoveryBlockBuilder) AsAdsDiscoveryBlockUserName() AdsDiscoveryBlockUserNameBuilder {
	if cb, ok := b.childBuilder.(AdsDiscoveryBlockUserNameBuilder); ok {
		return cb
	}
	cb := NewAdsDiscoveryBlockUserNameBuilder().(*_AdsDiscoveryBlockUserNameBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_AdsDiscoveryBlockBuilder) AsAdsDiscoveryBlockFingerprint() AdsDiscoveryBlockFingerprintBuilder {
	if cb, ok := b.childBuilder.(AdsDiscoveryBlockFingerprintBuilder); ok {
		return cb
	}
	cb := NewAdsDiscoveryBlockFingerprintBuilder().(*_AdsDiscoveryBlockFingerprintBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_AdsDiscoveryBlockBuilder) Build() (AdsDiscoveryBlock, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForAdsDiscoveryBlock()
}

func (b *_AdsDiscoveryBlockBuilder) MustBuild() AdsDiscoveryBlock {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AdsDiscoveryBlockBuilder) DeepCopy() any {
	_copy := b.CreateAdsDiscoveryBlockBuilder().(*_AdsDiscoveryBlockBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_AdsDiscoveryBlockChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAdsDiscoveryBlockBuilder creates a AdsDiscoveryBlockBuilder
func (b *_AdsDiscoveryBlock) CreateAdsDiscoveryBlockBuilder() AdsDiscoveryBlockBuilder {
	if b == nil {
		return NewAdsDiscoveryBlockBuilder()
	}
	return &_AdsDiscoveryBlockBuilder{_AdsDiscoveryBlock: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAdsDiscoveryBlock(structType any) AdsDiscoveryBlock {
	if casted, ok := structType.(AdsDiscoveryBlock); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDiscoveryBlock); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDiscoveryBlock) GetTypeName() string {
	return "AdsDiscoveryBlock"
}

func (m *_AdsDiscoveryBlock) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (blockType)
	lengthInBits += 16

	return lengthInBits
}

func (m *_AdsDiscoveryBlock) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_AdsDiscoveryBlock) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func AdsDiscoveryBlockParse[T AdsDiscoveryBlock](ctx context.Context, theBytes []byte) (T, error) {
	return AdsDiscoveryBlockParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func AdsDiscoveryBlockParseWithBufferProducer[T AdsDiscoveryBlock]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := AdsDiscoveryBlockParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func AdsDiscoveryBlockParseWithBuffer[T AdsDiscoveryBlock](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_AdsDiscoveryBlock{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_AdsDiscoveryBlock) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__adsDiscoveryBlock AdsDiscoveryBlock, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDiscoveryBlock"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDiscoveryBlock")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	blockType, err := ReadDiscriminatorEnumField[AdsDiscoveryBlockType](ctx, "blockType", "AdsDiscoveryBlockType", ReadEnum(AdsDiscoveryBlockTypeByValue, ReadUnsignedShort(readBuffer, uint8(16))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'blockType' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child AdsDiscoveryBlock
	switch {
	case blockType == AdsDiscoveryBlockType_STATUS: // AdsDiscoveryBlockStatus
		if _child, err = new(_AdsDiscoveryBlockStatus).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsDiscoveryBlockStatus for type-switch of AdsDiscoveryBlock")
		}
	case blockType == AdsDiscoveryBlockType_PASSWORD: // AdsDiscoveryBlockPassword
		if _child, err = new(_AdsDiscoveryBlockPassword).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsDiscoveryBlockPassword for type-switch of AdsDiscoveryBlock")
		}
	case blockType == AdsDiscoveryBlockType_VERSION: // AdsDiscoveryBlockVersion
		if _child, err = new(_AdsDiscoveryBlockVersion).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsDiscoveryBlockVersion for type-switch of AdsDiscoveryBlock")
		}
	case blockType == AdsDiscoveryBlockType_OS_DATA: // AdsDiscoveryBlockOsData
		if _child, err = new(_AdsDiscoveryBlockOsData).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsDiscoveryBlockOsData for type-switch of AdsDiscoveryBlock")
		}
	case blockType == AdsDiscoveryBlockType_HOST_NAME: // AdsDiscoveryBlockHostName
		if _child, err = new(_AdsDiscoveryBlockHostName).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsDiscoveryBlockHostName for type-switch of AdsDiscoveryBlock")
		}
	case blockType == AdsDiscoveryBlockType_AMS_NET_ID: // AdsDiscoveryBlockAmsNetId
		if _child, err = new(_AdsDiscoveryBlockAmsNetId).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsDiscoveryBlockAmsNetId for type-switch of AdsDiscoveryBlock")
		}
	case blockType == AdsDiscoveryBlockType_ROUTE_NAME: // AdsDiscoveryBlockRouteName
		if _child, err = new(_AdsDiscoveryBlockRouteName).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsDiscoveryBlockRouteName for type-switch of AdsDiscoveryBlock")
		}
	case blockType == AdsDiscoveryBlockType_USER_NAME: // AdsDiscoveryBlockUserName
		if _child, err = new(_AdsDiscoveryBlockUserName).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsDiscoveryBlockUserName for type-switch of AdsDiscoveryBlock")
		}
	case blockType == AdsDiscoveryBlockType_FINGERPRINT: // AdsDiscoveryBlockFingerprint
		if _child, err = new(_AdsDiscoveryBlockFingerprint).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsDiscoveryBlockFingerprint for type-switch of AdsDiscoveryBlock")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [blockType=%v]", blockType)
	}

	if closeErr := readBuffer.CloseContext("AdsDiscoveryBlock"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDiscoveryBlock")
	}

	return _child, nil
}

func (pm *_AdsDiscoveryBlock) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child AdsDiscoveryBlock, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AdsDiscoveryBlock"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AdsDiscoveryBlock")
	}

	if err := WriteDiscriminatorEnumField(ctx, "blockType", "AdsDiscoveryBlockType", m.GetBlockType(), WriteEnum[AdsDiscoveryBlockType, uint16](AdsDiscoveryBlockType.GetValue, AdsDiscoveryBlockType.PLC4XEnumName, WriteUnsignedShort(writeBuffer, 16))); err != nil {
		return errors.Wrap(err, "Error serializing 'blockType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("AdsDiscoveryBlock"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AdsDiscoveryBlock")
	}
	return nil
}

func (m *_AdsDiscoveryBlock) IsAdsDiscoveryBlock() {}

func (m *_AdsDiscoveryBlock) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsDiscoveryBlock) deepCopy() *_AdsDiscoveryBlock {
	if m == nil {
		return nil
	}
	_AdsDiscoveryBlockCopy := &_AdsDiscoveryBlock{
		nil, // will be set by child
	}
	return _AdsDiscoveryBlockCopy
}
