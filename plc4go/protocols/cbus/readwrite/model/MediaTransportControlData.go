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

// MediaTransportControlData is the corresponding interface of MediaTransportControlData
type MediaTransportControlData interface {
	MediaTransportControlDataContract
	MediaTransportControlDataRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// MediaTransportControlDataContract provides a set of functions which can be overwritten by a sub struct
type MediaTransportControlDataContract interface {
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() MediaTransportControlCommandTypeContainer
	// GetMediaLinkGroup returns MediaLinkGroup (property field)
	GetMediaLinkGroup() byte
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() MediaTransportControlCommandType
	// IsMediaTransportControlData() is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMediaTransportControlData()
}

// MediaTransportControlDataRequirements provides a set of functions which need to be implemented by a sub struct
type MediaTransportControlDataRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetCommandType returns CommandType (discriminator field)
	GetCommandType() MediaTransportControlCommandType
}

// _MediaTransportControlData is the data-structure of this message
type _MediaTransportControlData struct {
	_SubType             MediaTransportControlData
	CommandTypeContainer MediaTransportControlCommandTypeContainer
	MediaLinkGroup       byte
}

var _ MediaTransportControlDataContract = (*_MediaTransportControlData)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlData) GetCommandTypeContainer() MediaTransportControlCommandTypeContainer {
	return m.CommandTypeContainer
}

func (m *_MediaTransportControlData) GetMediaLinkGroup() byte {
	return m.MediaLinkGroup
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_MediaTransportControlData) GetCommandType() MediaTransportControlCommandType {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return CastMediaTransportControlCommandType(m.GetCommandTypeContainer().CommandType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMediaTransportControlData factory function for _MediaTransportControlData
func NewMediaTransportControlData(commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) *_MediaTransportControlData {
	return &_MediaTransportControlData{CommandTypeContainer: commandTypeContainer, MediaLinkGroup: mediaLinkGroup}
}

// Deprecated: use the interface for direct cast
func CastMediaTransportControlData(structType any) MediaTransportControlData {
	if casted, ok := structType.(MediaTransportControlData); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlData); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlData) GetTypeName() string {
	return "MediaTransportControlData"
}

func (m *_MediaTransportControlData) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (mediaLinkGroup)
	lengthInBits += 8

	return lengthInBits
}

func (m *_MediaTransportControlData) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func MediaTransportControlDataParse[T MediaTransportControlData](ctx context.Context, theBytes []byte) (T, error) {
	return MediaTransportControlDataParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func MediaTransportControlDataParseWithBufferProducer[T MediaTransportControlData]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := MediaTransportControlDataParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func MediaTransportControlDataParseWithBuffer[T MediaTransportControlData](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_MediaTransportControlData{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_MediaTransportControlData) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__mediaTransportControlData MediaTransportControlData, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(KnowsMediaTransportControlCommandTypeContainer(ctx, readBuffer)) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "no command type could be found"})
	}

	commandTypeContainer, err := ReadEnumField[MediaTransportControlCommandTypeContainer](ctx, "commandTypeContainer", "MediaTransportControlCommandTypeContainer", ReadEnum(MediaTransportControlCommandTypeContainerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandTypeContainer' field"))
	}
	m.CommandTypeContainer = commandTypeContainer

	commandType, err := ReadVirtualField[MediaTransportControlCommandType](ctx, "commandType", (*MediaTransportControlCommandType)(nil), commandTypeContainer.CommandType())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandType' field"))
	}
	_ = commandType

	mediaLinkGroup, err := ReadSimpleField(ctx, "mediaLinkGroup", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'mediaLinkGroup' field"))
	}
	m.MediaLinkGroup = mediaLinkGroup

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child MediaTransportControlData
	switch {
	case commandType == MediaTransportControlCommandType_STOP: // MediaTransportControlDataStop
		if _child, err = (&_MediaTransportControlDataStop{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataStop for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_PLAY: // MediaTransportControlDataPlay
		if _child, err = (&_MediaTransportControlDataPlay{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataPlay for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_PAUSE_RESUME: // MediaTransportControlDataPauseResume
		if _child, err = (&_MediaTransportControlDataPauseResume{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataPauseResume for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_SELECT_CATEGORY: // MediaTransportControlDataSetCategory
		if _child, err = (&_MediaTransportControlDataSetCategory{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataSetCategory for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_SELECT_SELECTION: // MediaTransportControlDataSetSelection
		if _child, err = (&_MediaTransportControlDataSetSelection{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataSetSelection for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_SELECT_TRACK: // MediaTransportControlDataSetTrack
		if _child, err = (&_MediaTransportControlDataSetTrack{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataSetTrack for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_SHUFFLE_ON_OFF: // MediaTransportControlDataShuffleOnOff
		if _child, err = (&_MediaTransportControlDataShuffleOnOff{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataShuffleOnOff for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_REPEAT_ON_OFF: // MediaTransportControlDataRepeatOnOff
		if _child, err = (&_MediaTransportControlDataRepeatOnOff{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataRepeatOnOff for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_NEXT_PREVIOUS_CATEGORY: // MediaTransportControlDataNextPreviousCategory
		if _child, err = (&_MediaTransportControlDataNextPreviousCategory{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataNextPreviousCategory for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_NEXT_PREVIOUS_SELECTION: // MediaTransportControlDataNextPreviousSelection
		if _child, err = (&_MediaTransportControlDataNextPreviousSelection{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataNextPreviousSelection for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_NEXT_PREVIOUS_TRACK: // MediaTransportControlDataNextPreviousTrack
		if _child, err = (&_MediaTransportControlDataNextPreviousTrack{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataNextPreviousTrack for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_FAST_FORWARD: // MediaTransportControlDataFastForward
		if _child, err = (&_MediaTransportControlDataFastForward{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataFastForward for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_REWIND: // MediaTransportControlDataRewind
		if _child, err = (&_MediaTransportControlDataRewind{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataRewind for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_SOURCE_POWER_CONTROL: // MediaTransportControlDataSourcePowerControl
		if _child, err = (&_MediaTransportControlDataSourcePowerControl{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataSourcePowerControl for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_TOTAL_TRACKS: // MediaTransportControlDataTotalTracks
		if _child, err = (&_MediaTransportControlDataTotalTracks{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataTotalTracks for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_STATUS_REQUEST: // MediaTransportControlDataStatusRequest
		if _child, err = (&_MediaTransportControlDataStatusRequest{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataStatusRequest for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_ENUMERATE_CATEGORIES_SELECTIONS_TRACKS: // MediaTransportControlDataEnumerateCategoriesSelectionTracks
		if _child, err = (&_MediaTransportControlDataEnumerateCategoriesSelectionTracks{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataEnumerateCategoriesSelectionTracks for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_ENUMERATION_SIZE: // MediaTransportControlDataEnumerationsSize
		if _child, err = (&_MediaTransportControlDataEnumerationsSize{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataEnumerationsSize for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_TRACK_NAME: // MediaTransportControlDataTrackName
		if _child, err = (&_MediaTransportControlDataTrackName{}).parse(ctx, readBuffer, m, commandTypeContainer); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataTrackName for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_SELECTION_NAME: // MediaTransportControlDataSelectionName
		if _child, err = (&_MediaTransportControlDataSelectionName{}).parse(ctx, readBuffer, m, commandTypeContainer); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataSelectionName for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_CATEGORY_NAME: // MediaTransportControlDataCategoryName
		if _child, err = (&_MediaTransportControlDataCategoryName{}).parse(ctx, readBuffer, m, commandTypeContainer); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataCategoryName for type-switch of MediaTransportControlData")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [commandType=%v]", commandType)
	}

	if closeErr := readBuffer.CloseContext("MediaTransportControlData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlData")
	}

	return _child, nil
}

func (pm *_MediaTransportControlData) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child MediaTransportControlData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("MediaTransportControlData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for MediaTransportControlData")
	}

	if err := WriteSimpleEnumField[MediaTransportControlCommandTypeContainer](ctx, "commandTypeContainer", "MediaTransportControlCommandTypeContainer", m.GetCommandTypeContainer(), WriteEnum[MediaTransportControlCommandTypeContainer, uint8](MediaTransportControlCommandTypeContainer.GetValue, MediaTransportControlCommandTypeContainer.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'commandTypeContainer' field")
	}
	// Virtual field
	commandType := m.GetCommandType()
	_ = commandType
	if _commandTypeErr := writeBuffer.WriteVirtual(ctx, "commandType", m.GetCommandType()); _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}

	if err := WriteSimpleField[byte](ctx, "mediaLinkGroup", m.GetMediaLinkGroup(), WriteByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'mediaLinkGroup' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("MediaTransportControlData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for MediaTransportControlData")
	}
	return nil
}

func (m *_MediaTransportControlData) IsMediaTransportControlData() {}
