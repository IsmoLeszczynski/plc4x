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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// ImageJPG is the corresponding interface of ImageJPG
type ImageJPG interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// _ImageJPG is the data-structure of this message
type _ImageJPG struct {
}

var _ ImageJPG = (*_ImageJPG)(nil)

// NewImageJPG factory function for _ImageJPG
func NewImageJPG() *_ImageJPG {
	return &_ImageJPG{}
}

// Deprecated: use the interface for direct cast
func CastImageJPG(structType any) ImageJPG {
	if casted, ok := structType.(ImageJPG); ok {
		return casted
	}
	if casted, ok := structType.(*ImageJPG); ok {
		return *casted
	}
	return nil
}

func (m *_ImageJPG) GetTypeName() string {
	return "ImageJPG"
}

func (m *_ImageJPG) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_ImageJPG) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ImageJPGParse(ctx context.Context, theBytes []byte) (ImageJPG, error) {
	return ImageJPGParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ImageJPGParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (ImageJPG, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ImageJPG, error) {
		return ImageJPGParseWithBuffer(ctx, readBuffer)
	}
}

func ImageJPGParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ImageJPG, error) {
	v, err := (&_ImageJPG{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_ImageJPG) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__imageJPG ImageJPG, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ImageJPG"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ImageJPG")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ImageJPG"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ImageJPG")
	}

	return m, nil
}

func (m *_ImageJPG) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ImageJPG) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ImageJPG"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ImageJPG")
	}

	if popErr := writeBuffer.PopContext("ImageJPG"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ImageJPG")
	}
	return nil
}

func (m *_ImageJPG) IsImageJPG() {}

func (m *_ImageJPG) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
