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

// PathSegment is the corresponding interface of PathSegment
type PathSegment interface {
	PathSegmentContract
	PathSegmentRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// PathSegmentContract provides a set of functions which can be overwritten by a sub struct
type PathSegmentContract interface {
	// IsPathSegment() is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsPathSegment()
}

// PathSegmentRequirements provides a set of functions which need to be implemented by a sub struct
type PathSegmentRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPathSegment returns PathSegment (discriminator field)
	GetPathSegment() uint8
}

// _PathSegment is the data-structure of this message
type _PathSegment struct {
	_SubType PathSegment
}

var _ PathSegmentContract = (*_PathSegment)(nil)

// NewPathSegment factory function for _PathSegment
func NewPathSegment() *_PathSegment {
	return &_PathSegment{}
}

// Deprecated: use the interface for direct cast
func CastPathSegment(structType any) PathSegment {
	if casted, ok := structType.(PathSegment); ok {
		return casted
	}
	if casted, ok := structType.(*PathSegment); ok {
		return *casted
	}
	return nil
}

func (m *_PathSegment) GetTypeName() string {
	return "PathSegment"
}

func (m *_PathSegment) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (pathSegment)
	lengthInBits += 3

	return lengthInBits
}

func (m *_PathSegment) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func PathSegmentParse[T PathSegment](ctx context.Context, theBytes []byte) (T, error) {
	return PathSegmentParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func PathSegmentParseWithBufferProducer[T PathSegment]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := PathSegmentParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func PathSegmentParseWithBuffer[T PathSegment](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_PathSegment{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_PathSegment) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__pathSegment PathSegment, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PathSegment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PathSegment")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	pathSegment, err := ReadDiscriminatorField[uint8](ctx, "pathSegment", ReadUnsignedByte(readBuffer, uint8(3)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pathSegment' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child PathSegment
	switch {
	case pathSegment == 0x00: // PortSegment
		if _child, err = (&_PortSegment{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type PortSegment for type-switch of PathSegment")
		}
	case pathSegment == 0x01: // LogicalSegment
		if _child, err = (&_LogicalSegment{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LogicalSegment for type-switch of PathSegment")
		}
	case pathSegment == 0x04: // DataSegment
		if _child, err = (&_DataSegment{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type DataSegment for type-switch of PathSegment")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [pathSegment=%v]", pathSegment)
	}

	if closeErr := readBuffer.CloseContext("PathSegment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PathSegment")
	}

	return _child, nil
}

func (pm *_PathSegment) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child PathSegment, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("PathSegment"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for PathSegment")
	}

	if err := WriteDiscriminatorField(ctx, "pathSegment", m.GetPathSegment(), WriteUnsignedByte(writeBuffer, 3)); err != nil {
		return errors.Wrap(err, "Error serializing 'pathSegment' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("PathSegment"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for PathSegment")
	}
	return nil
}

func (m *_PathSegment) IsPathSegment() {}
