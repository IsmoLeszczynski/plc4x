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

// CEMI is the corresponding interface of CEMI
type CEMI interface {
	CEMIContract
	CEMIRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// CEMIContract provides a set of functions which can be overwritten by a sub struct
type CEMIContract interface {
	// GetSize() returns a parser argument
	GetSize() uint16
	// IsCEMI() is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCEMI()
}

// CEMIRequirements provides a set of functions which need to be implemented by a sub struct
type CEMIRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetMessageCode returns MessageCode (discriminator field)
	GetMessageCode() uint8
}

// _CEMI is the data-structure of this message
type _CEMI struct {
	_SubType CEMI

	// Arguments.
	Size uint16
}

var _ CEMIContract = (*_CEMI)(nil)

// NewCEMI factory function for _CEMI
func NewCEMI(size uint16) *_CEMI {
	return &_CEMI{Size: size}
}

// Deprecated: use the interface for direct cast
func CastCEMI(structType any) CEMI {
	if casted, ok := structType.(CEMI); ok {
		return casted
	}
	if casted, ok := structType.(*CEMI); ok {
		return *casted
	}
	return nil
}

func (m *_CEMI) GetTypeName() string {
	return "CEMI"
}

func (m *_CEMI) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (messageCode)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CEMI) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func CEMIParse[T CEMI](ctx context.Context, theBytes []byte, size uint16) (T, error) {
	return CEMIParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), size)
}

func CEMIParseWithBufferProducer[T CEMI](size uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := CEMIParseWithBuffer[T](ctx, readBuffer, size)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func CEMIParseWithBuffer[T CEMI](ctx context.Context, readBuffer utils.ReadBuffer, size uint16) (T, error) {
	v, err := (&_CEMI{Size: size}).parse(ctx, readBuffer, size)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_CEMI) parse(ctx context.Context, readBuffer utils.ReadBuffer, size uint16) (__cEMI CEMI, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CEMI"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CEMI")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	messageCode, err := ReadDiscriminatorField[uint8](ctx, "messageCode", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageCode' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child CEMI
	switch {
	case messageCode == 0x2B: // LBusmonInd
		if _child, err = (&_LBusmonInd{}).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LBusmonInd for type-switch of CEMI")
		}
	case messageCode == 0x11: // LDataReq
		if _child, err = (&_LDataReq{}).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LDataReq for type-switch of CEMI")
		}
	case messageCode == 0x29: // LDataInd
		if _child, err = (&_LDataInd{}).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LDataInd for type-switch of CEMI")
		}
	case messageCode == 0x2E: // LDataCon
		if _child, err = (&_LDataCon{}).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LDataCon for type-switch of CEMI")
		}
	case messageCode == 0x10: // LRawReq
		if _child, err = (&_LRawReq{}).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LRawReq for type-switch of CEMI")
		}
	case messageCode == 0x2D: // LRawInd
		if _child, err = (&_LRawInd{}).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LRawInd for type-switch of CEMI")
		}
	case messageCode == 0x2F: // LRawCon
		if _child, err = (&_LRawCon{}).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LRawCon for type-switch of CEMI")
		}
	case messageCode == 0x13: // LPollDataReq
		if _child, err = (&_LPollDataReq{}).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LPollDataReq for type-switch of CEMI")
		}
	case messageCode == 0x25: // LPollDataCon
		if _child, err = (&_LPollDataCon{}).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LPollDataCon for type-switch of CEMI")
		}
	case messageCode == 0x41: // TDataConnectedReq
		if _child, err = (&_TDataConnectedReq{}).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TDataConnectedReq for type-switch of CEMI")
		}
	case messageCode == 0x89: // TDataConnectedInd
		if _child, err = (&_TDataConnectedInd{}).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TDataConnectedInd for type-switch of CEMI")
		}
	case messageCode == 0x4A: // TDataIndividualReq
		if _child, err = (&_TDataIndividualReq{}).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TDataIndividualReq for type-switch of CEMI")
		}
	case messageCode == 0x94: // TDataIndividualInd
		if _child, err = (&_TDataIndividualInd{}).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TDataIndividualInd for type-switch of CEMI")
		}
	case messageCode == 0xFC: // MPropReadReq
		if _child, err = (&_MPropReadReq{}).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MPropReadReq for type-switch of CEMI")
		}
	case messageCode == 0xFB: // MPropReadCon
		if _child, err = (&_MPropReadCon{}).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MPropReadCon for type-switch of CEMI")
		}
	case messageCode == 0xF6: // MPropWriteReq
		if _child, err = (&_MPropWriteReq{}).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MPropWriteReq for type-switch of CEMI")
		}
	case messageCode == 0xF5: // MPropWriteCon
		if _child, err = (&_MPropWriteCon{}).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MPropWriteCon for type-switch of CEMI")
		}
	case messageCode == 0xF7: // MPropInfoInd
		if _child, err = (&_MPropInfoInd{}).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MPropInfoInd for type-switch of CEMI")
		}
	case messageCode == 0xF8: // MFuncPropCommandReq
		if _child, err = (&_MFuncPropCommandReq{}).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MFuncPropCommandReq for type-switch of CEMI")
		}
	case messageCode == 0xF9: // MFuncPropStateReadReq
		if _child, err = (&_MFuncPropStateReadReq{}).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MFuncPropStateReadReq for type-switch of CEMI")
		}
	case messageCode == 0xFA: // MFuncPropCon
		if _child, err = (&_MFuncPropCon{}).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MFuncPropCon for type-switch of CEMI")
		}
	case messageCode == 0xF1: // MResetReq
		if _child, err = (&_MResetReq{}).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MResetReq for type-switch of CEMI")
		}
	case messageCode == 0xF0: // MResetInd
		if _child, err = (&_MResetInd{}).parse(ctx, readBuffer, m, size); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MResetInd for type-switch of CEMI")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [messageCode=%v]", messageCode)
	}

	if closeErr := readBuffer.CloseContext("CEMI"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CEMI")
	}

	return _child, nil
}

func (pm *_CEMI) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CEMI, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CEMI"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CEMI")
	}

	if err := WriteDiscriminatorField(ctx, "messageCode", m.GetMessageCode(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'messageCode' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CEMI"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CEMI")
	}
	return nil
}

////
// Arguments Getter

func (m *_CEMI) GetSize() uint16 {
	return m.Size
}

//
////

func (m *_CEMI) IsCEMI() {}
