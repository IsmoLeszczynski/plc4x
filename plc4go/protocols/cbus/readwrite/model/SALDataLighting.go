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

// SALDataLighting is the corresponding interface of SALDataLighting
type SALDataLighting interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SALData
	// GetLightingData returns LightingData (property field)
	GetLightingData() LightingData
}

// SALDataLightingExactly can be used when we want exactly this type and not a type which fulfills SALDataLighting.
// This is useful for switch cases.
type SALDataLightingExactly interface {
	SALDataLighting
	isSALDataLighting() bool
}

// _SALDataLighting is the data-structure of this message
type _SALDataLighting struct {
	*_SALData
	LightingData LightingData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataLighting) GetApplicationId() ApplicationId {
	return ApplicationId_LIGHTING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataLighting) InitializeParent(parent SALData, salData SALData) {
	m.SalData = salData
}

func (m *_SALDataLighting) GetParent() SALData {
	return m._SALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataLighting) GetLightingData() LightingData {
	return m.LightingData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALDataLighting factory function for _SALDataLighting
func NewSALDataLighting(lightingData LightingData, salData SALData) *_SALDataLighting {
	_result := &_SALDataLighting{
		LightingData: lightingData,
		_SALData:     NewSALData(salData),
	}
	_result._SALData._SALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataLighting(structType any) SALDataLighting {
	if casted, ok := structType.(SALDataLighting); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataLighting); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataLighting) GetTypeName() string {
	return "SALDataLighting"
}

func (m *_SALDataLighting) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (lightingData)
	lengthInBits += m.LightingData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataLighting) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SALDataLightingParse(ctx context.Context, theBytes []byte, applicationId ApplicationId) (SALDataLighting, error) {
	return SALDataLightingParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), applicationId)
}

func SALDataLightingParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, applicationId ApplicationId) (SALDataLighting, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("SALDataLighting"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataLighting")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lightingData)
	if pullErr := readBuffer.PullContext("lightingData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lightingData")
	}
	_lightingData, _lightingDataErr := LightingDataParseWithBuffer(ctx, readBuffer)
	if _lightingDataErr != nil {
		return nil, errors.Wrap(_lightingDataErr, "Error parsing 'lightingData' field of SALDataLighting")
	}
	lightingData := _lightingData.(LightingData)
	if closeErr := readBuffer.CloseContext("lightingData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lightingData")
	}

	if closeErr := readBuffer.CloseContext("SALDataLighting"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataLighting")
	}

	// Create a partially initialized instance
	_child := &_SALDataLighting{
		_SALData:     &_SALData{},
		LightingData: lightingData,
	}
	_child._SALData._SALDataChildRequirements = _child
	return _child, nil
}

func (m *_SALDataLighting) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataLighting) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataLighting"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataLighting")
		}

		// Simple Field (lightingData)
		if pushErr := writeBuffer.PushContext("lightingData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lightingData")
		}
		_lightingDataErr := writeBuffer.WriteSerializable(ctx, m.GetLightingData())
		if popErr := writeBuffer.PopContext("lightingData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lightingData")
		}
		if _lightingDataErr != nil {
			return errors.Wrap(_lightingDataErr, "Error serializing 'lightingData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataLighting"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataLighting")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataLighting) isSALDataLighting() bool {
	return true
}

func (m *_SALDataLighting) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
