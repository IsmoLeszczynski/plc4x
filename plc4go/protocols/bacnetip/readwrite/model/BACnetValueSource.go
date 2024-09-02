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
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetValueSource is the corresponding interface of BACnetValueSource
type BACnetValueSource interface {
	BACnetValueSourceContract
	BACnetValueSourceRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// BACnetValueSourceContract provides a set of functions which can be overwritten by a sub struct
type BACnetValueSourceContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// IsBACnetValueSource() is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetValueSource()
}

// BACnetValueSourceRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetValueSourceRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetValueSource is the data-structure of this message
type _BACnetValueSource struct {
	_SubType        BACnetValueSource
	PeekedTagHeader BACnetTagHeader
}

var _ BACnetValueSourceContract = (*_BACnetValueSource)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetValueSource) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetValueSource) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetValueSource factory function for _BACnetValueSource
func NewBACnetValueSource(peekedTagHeader BACnetTagHeader) *_BACnetValueSource {
	return &_BACnetValueSource{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetValueSource(structType any) BACnetValueSource {
	if casted, ok := structType.(BACnetValueSource); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetValueSource); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetValueSource) GetTypeName() string {
	return "BACnetValueSource"
}

func (m *_BACnetValueSource) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetValueSource) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetValueSourceParse[T BACnetValueSource](ctx context.Context, theBytes []byte) (T, error) {
	return BACnetValueSourceParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetValueSourceParseWithBufferProducer[T BACnetValueSource]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetValueSourceParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func BACnetValueSourceParseWithBuffer[T BACnetValueSource](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_BACnetValueSource{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_BACnetValueSource) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetValueSource BACnetValueSource, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetValueSource"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetValueSource")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetValueSource
	switch {
	case peekedTagNumber == uint8(0): // BACnetValueSourceNone
		if _child, err = (&_BACnetValueSourceNone{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetValueSourceNone for type-switch of BACnetValueSource")
		}
	case peekedTagNumber == uint8(1): // BACnetValueSourceObject
		if _child, err = (&_BACnetValueSourceObject{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetValueSourceObject for type-switch of BACnetValueSource")
		}
	case peekedTagNumber == uint8(2): // BACnetValueSourceAddress
		if _child, err = (&_BACnetValueSourceAddress{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetValueSourceAddress for type-switch of BACnetValueSource")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	if closeErr := readBuffer.CloseContext("BACnetValueSource"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetValueSource")
	}

	return _child, nil
}

func (pm *_BACnetValueSource) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetValueSource, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetValueSource"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetValueSource")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetValueSource"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetValueSource")
	}
	return nil
}

func (m *_BACnetValueSource) IsBACnetValueSource() {}
