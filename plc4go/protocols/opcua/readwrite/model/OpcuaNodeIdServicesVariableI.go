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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// OpcuaNodeIdServicesVariableI is an enum
type OpcuaNodeIdServicesVariableI int32

type IOpcuaNodeIdServicesVariableI interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableI_IOrderedObjectType_NumberInList                          OpcuaNodeIdServicesVariableI = 23517
	OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_AdminStatus                OpcuaNodeIdServicesVariableI = 24149
	OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_OperStatus                 OpcuaNodeIdServicesVariableI = 24150
	OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_PhysAddress                OpcuaNodeIdServicesVariableI = 24151
	OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed                      OpcuaNodeIdServicesVariableI = 24152
	OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed_Definition           OpcuaNodeIdServicesVariableI = 24153
	OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed_ValuePrecision       OpcuaNodeIdServicesVariableI = 24154
	OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed_InstrumentRange      OpcuaNodeIdServicesVariableI = 24155
	OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed_EURange              OpcuaNodeIdServicesVariableI = 24156
	OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed_EngineeringUnits     OpcuaNodeIdServicesVariableI = 24157
	OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed                          OpcuaNodeIdServicesVariableI = 24159
	OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed_Definition               OpcuaNodeIdServicesVariableI = 24160
	OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed_ValuePrecision           OpcuaNodeIdServicesVariableI = 24161
	OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed_InstrumentRange          OpcuaNodeIdServicesVariableI = 24162
	OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed_EURange                  OpcuaNodeIdServicesVariableI = 24163
	OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed_EngineeringUnits         OpcuaNodeIdServicesVariableI = 24164
	OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Duplex                         OpcuaNodeIdServicesVariableI = 24165
	OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_MaxFrameLength                 OpcuaNodeIdServicesVariableI = 24166
	OpcuaNodeIdServicesVariableI_IBaseEthernetCapabilitiesType_VlanTagCapable             OpcuaNodeIdServicesVariableI = 24168
	OpcuaNodeIdServicesVariableI_ISrClassType_Id                                          OpcuaNodeIdServicesVariableI = 24170
	OpcuaNodeIdServicesVariableI_ISrClassType_Priority                                    OpcuaNodeIdServicesVariableI = 24171
	OpcuaNodeIdServicesVariableI_ISrClassType_Vid                                         OpcuaNodeIdServicesVariableI = 24172
	OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStreamType_StreamId                          OpcuaNodeIdServicesVariableI = 24174
	OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStreamType_StreamName                        OpcuaNodeIdServicesVariableI = 24175
	OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStreamType_State                             OpcuaNodeIdServicesVariableI = 24176
	OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStreamType_AccumulatedLatency                OpcuaNodeIdServicesVariableI = 24177
	OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStreamType_SrClassId                         OpcuaNodeIdServicesVariableI = 24178
	OpcuaNodeIdServicesVariableI_IIeeeBaseTsnTrafficSpecificationType_MaxIntervalFrames   OpcuaNodeIdServicesVariableI = 24180
	OpcuaNodeIdServicesVariableI_IIeeeBaseTsnTrafficSpecificationType_MaxFrameSize        OpcuaNodeIdServicesVariableI = 24181
	OpcuaNodeIdServicesVariableI_IIeeeBaseTsnTrafficSpecificationType_Interval            OpcuaNodeIdServicesVariableI = 24182
	OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStatusStreamType_TalkerStatus                OpcuaNodeIdServicesVariableI = 24184
	OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStatusStreamType_ListenerStatus              OpcuaNodeIdServicesVariableI = 24185
	OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStatusStreamType_FailureCode                 OpcuaNodeIdServicesVariableI = 24186
	OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStatusStreamType_FailureSystemIdentifier     OpcuaNodeIdServicesVariableI = 24187
	OpcuaNodeIdServicesVariableI_IIeeeTsnInterfaceConfigurationType_MacAddress            OpcuaNodeIdServicesVariableI = 24189
	OpcuaNodeIdServicesVariableI_IIeeeTsnInterfaceConfigurationType_InterfaceName         OpcuaNodeIdServicesVariableI = 24190
	OpcuaNodeIdServicesVariableI_IIeeeTsnInterfaceConfigurationTalkerType_TimeAwareOffset OpcuaNodeIdServicesVariableI = 24194
	OpcuaNodeIdServicesVariableI_IIeeeTsnInterfaceConfigurationListenerType_ReceiveOffset OpcuaNodeIdServicesVariableI = 24198
	OpcuaNodeIdServicesVariableI_IIeeeTsnMacAddressType_DestinationAddress                OpcuaNodeIdServicesVariableI = 24200
	OpcuaNodeIdServicesVariableI_IIeeeTsnMacAddressType_SourceAddress                     OpcuaNodeIdServicesVariableI = 24201
	OpcuaNodeIdServicesVariableI_IIeeeTsnVlanTagType_VlanId                               OpcuaNodeIdServicesVariableI = 24203
	OpcuaNodeIdServicesVariableI_IIeeeTsnVlanTagType_PriorityCodePoint                    OpcuaNodeIdServicesVariableI = 24204
	OpcuaNodeIdServicesVariableI_IPriorityMappingEntryType_MappingUri                     OpcuaNodeIdServicesVariableI = 24206
	OpcuaNodeIdServicesVariableI_IPriorityMappingEntryType_PriorityLabel                  OpcuaNodeIdServicesVariableI = 24207
	OpcuaNodeIdServicesVariableI_IPriorityMappingEntryType_PriorityValue_PCP              OpcuaNodeIdServicesVariableI = 24208
	OpcuaNodeIdServicesVariableI_IPriorityMappingEntryType_PriorityValue_DSCP             OpcuaNodeIdServicesVariableI = 24209
	OpcuaNodeIdServicesVariableI_IIeeeAutoNegotiationStatusType_NegotiationStatus         OpcuaNodeIdServicesVariableI = 24234
	OpcuaNodeIdServicesVariableI_IVlanIdType_VlanId                                       OpcuaNodeIdServicesVariableI = 25219
)

var OpcuaNodeIdServicesVariableIValues []OpcuaNodeIdServicesVariableI

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableIValues = []OpcuaNodeIdServicesVariableI{
		OpcuaNodeIdServicesVariableI_IOrderedObjectType_NumberInList,
		OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_AdminStatus,
		OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_OperStatus,
		OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_PhysAddress,
		OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed,
		OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed_Definition,
		OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed_ValuePrecision,
		OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed_InstrumentRange,
		OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed_EURange,
		OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed_EngineeringUnits,
		OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed,
		OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed_Definition,
		OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed_ValuePrecision,
		OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed_InstrumentRange,
		OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed_EURange,
		OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed_EngineeringUnits,
		OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Duplex,
		OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_MaxFrameLength,
		OpcuaNodeIdServicesVariableI_IBaseEthernetCapabilitiesType_VlanTagCapable,
		OpcuaNodeIdServicesVariableI_ISrClassType_Id,
		OpcuaNodeIdServicesVariableI_ISrClassType_Priority,
		OpcuaNodeIdServicesVariableI_ISrClassType_Vid,
		OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStreamType_StreamId,
		OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStreamType_StreamName,
		OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStreamType_State,
		OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStreamType_AccumulatedLatency,
		OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStreamType_SrClassId,
		OpcuaNodeIdServicesVariableI_IIeeeBaseTsnTrafficSpecificationType_MaxIntervalFrames,
		OpcuaNodeIdServicesVariableI_IIeeeBaseTsnTrafficSpecificationType_MaxFrameSize,
		OpcuaNodeIdServicesVariableI_IIeeeBaseTsnTrafficSpecificationType_Interval,
		OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStatusStreamType_TalkerStatus,
		OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStatusStreamType_ListenerStatus,
		OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStatusStreamType_FailureCode,
		OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStatusStreamType_FailureSystemIdentifier,
		OpcuaNodeIdServicesVariableI_IIeeeTsnInterfaceConfigurationType_MacAddress,
		OpcuaNodeIdServicesVariableI_IIeeeTsnInterfaceConfigurationType_InterfaceName,
		OpcuaNodeIdServicesVariableI_IIeeeTsnInterfaceConfigurationTalkerType_TimeAwareOffset,
		OpcuaNodeIdServicesVariableI_IIeeeTsnInterfaceConfigurationListenerType_ReceiveOffset,
		OpcuaNodeIdServicesVariableI_IIeeeTsnMacAddressType_DestinationAddress,
		OpcuaNodeIdServicesVariableI_IIeeeTsnMacAddressType_SourceAddress,
		OpcuaNodeIdServicesVariableI_IIeeeTsnVlanTagType_VlanId,
		OpcuaNodeIdServicesVariableI_IIeeeTsnVlanTagType_PriorityCodePoint,
		OpcuaNodeIdServicesVariableI_IPriorityMappingEntryType_MappingUri,
		OpcuaNodeIdServicesVariableI_IPriorityMappingEntryType_PriorityLabel,
		OpcuaNodeIdServicesVariableI_IPriorityMappingEntryType_PriorityValue_PCP,
		OpcuaNodeIdServicesVariableI_IPriorityMappingEntryType_PriorityValue_DSCP,
		OpcuaNodeIdServicesVariableI_IIeeeAutoNegotiationStatusType_NegotiationStatus,
		OpcuaNodeIdServicesVariableI_IVlanIdType_VlanId,
	}
}

func OpcuaNodeIdServicesVariableIByValue(value int32) (enum OpcuaNodeIdServicesVariableI, ok bool) {
	switch value {
	case 23517:
		return OpcuaNodeIdServicesVariableI_IOrderedObjectType_NumberInList, true
	case 24149:
		return OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_AdminStatus, true
	case 24150:
		return OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_OperStatus, true
	case 24151:
		return OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_PhysAddress, true
	case 24152:
		return OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed, true
	case 24153:
		return OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed_Definition, true
	case 24154:
		return OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed_ValuePrecision, true
	case 24155:
		return OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed_InstrumentRange, true
	case 24156:
		return OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed_EURange, true
	case 24157:
		return OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed_EngineeringUnits, true
	case 24159:
		return OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed, true
	case 24160:
		return OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed_Definition, true
	case 24161:
		return OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed_ValuePrecision, true
	case 24162:
		return OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed_InstrumentRange, true
	case 24163:
		return OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed_EURange, true
	case 24164:
		return OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed_EngineeringUnits, true
	case 24165:
		return OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Duplex, true
	case 24166:
		return OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_MaxFrameLength, true
	case 24168:
		return OpcuaNodeIdServicesVariableI_IBaseEthernetCapabilitiesType_VlanTagCapable, true
	case 24170:
		return OpcuaNodeIdServicesVariableI_ISrClassType_Id, true
	case 24171:
		return OpcuaNodeIdServicesVariableI_ISrClassType_Priority, true
	case 24172:
		return OpcuaNodeIdServicesVariableI_ISrClassType_Vid, true
	case 24174:
		return OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStreamType_StreamId, true
	case 24175:
		return OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStreamType_StreamName, true
	case 24176:
		return OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStreamType_State, true
	case 24177:
		return OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStreamType_AccumulatedLatency, true
	case 24178:
		return OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStreamType_SrClassId, true
	case 24180:
		return OpcuaNodeIdServicesVariableI_IIeeeBaseTsnTrafficSpecificationType_MaxIntervalFrames, true
	case 24181:
		return OpcuaNodeIdServicesVariableI_IIeeeBaseTsnTrafficSpecificationType_MaxFrameSize, true
	case 24182:
		return OpcuaNodeIdServicesVariableI_IIeeeBaseTsnTrafficSpecificationType_Interval, true
	case 24184:
		return OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStatusStreamType_TalkerStatus, true
	case 24185:
		return OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStatusStreamType_ListenerStatus, true
	case 24186:
		return OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStatusStreamType_FailureCode, true
	case 24187:
		return OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStatusStreamType_FailureSystemIdentifier, true
	case 24189:
		return OpcuaNodeIdServicesVariableI_IIeeeTsnInterfaceConfigurationType_MacAddress, true
	case 24190:
		return OpcuaNodeIdServicesVariableI_IIeeeTsnInterfaceConfigurationType_InterfaceName, true
	case 24194:
		return OpcuaNodeIdServicesVariableI_IIeeeTsnInterfaceConfigurationTalkerType_TimeAwareOffset, true
	case 24198:
		return OpcuaNodeIdServicesVariableI_IIeeeTsnInterfaceConfigurationListenerType_ReceiveOffset, true
	case 24200:
		return OpcuaNodeIdServicesVariableI_IIeeeTsnMacAddressType_DestinationAddress, true
	case 24201:
		return OpcuaNodeIdServicesVariableI_IIeeeTsnMacAddressType_SourceAddress, true
	case 24203:
		return OpcuaNodeIdServicesVariableI_IIeeeTsnVlanTagType_VlanId, true
	case 24204:
		return OpcuaNodeIdServicesVariableI_IIeeeTsnVlanTagType_PriorityCodePoint, true
	case 24206:
		return OpcuaNodeIdServicesVariableI_IPriorityMappingEntryType_MappingUri, true
	case 24207:
		return OpcuaNodeIdServicesVariableI_IPriorityMappingEntryType_PriorityLabel, true
	case 24208:
		return OpcuaNodeIdServicesVariableI_IPriorityMappingEntryType_PriorityValue_PCP, true
	case 24209:
		return OpcuaNodeIdServicesVariableI_IPriorityMappingEntryType_PriorityValue_DSCP, true
	case 24234:
		return OpcuaNodeIdServicesVariableI_IIeeeAutoNegotiationStatusType_NegotiationStatus, true
	case 25219:
		return OpcuaNodeIdServicesVariableI_IVlanIdType_VlanId, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableIByName(value string) (enum OpcuaNodeIdServicesVariableI, ok bool) {
	switch value {
	case "IOrderedObjectType_NumberInList":
		return OpcuaNodeIdServicesVariableI_IOrderedObjectType_NumberInList, true
	case "IIetfBaseNetworkInterfaceType_AdminStatus":
		return OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_AdminStatus, true
	case "IIetfBaseNetworkInterfaceType_OperStatus":
		return OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_OperStatus, true
	case "IIetfBaseNetworkInterfaceType_PhysAddress":
		return OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_PhysAddress, true
	case "IIetfBaseNetworkInterfaceType_Speed":
		return OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed, true
	case "IIetfBaseNetworkInterfaceType_Speed_Definition":
		return OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed_Definition, true
	case "IIetfBaseNetworkInterfaceType_Speed_ValuePrecision":
		return OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed_ValuePrecision, true
	case "IIetfBaseNetworkInterfaceType_Speed_InstrumentRange":
		return OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed_InstrumentRange, true
	case "IIetfBaseNetworkInterfaceType_Speed_EURange":
		return OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed_EURange, true
	case "IIetfBaseNetworkInterfaceType_Speed_EngineeringUnits":
		return OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed_EngineeringUnits, true
	case "IIeeeBaseEthernetPortType_Speed":
		return OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed, true
	case "IIeeeBaseEthernetPortType_Speed_Definition":
		return OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed_Definition, true
	case "IIeeeBaseEthernetPortType_Speed_ValuePrecision":
		return OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed_ValuePrecision, true
	case "IIeeeBaseEthernetPortType_Speed_InstrumentRange":
		return OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed_InstrumentRange, true
	case "IIeeeBaseEthernetPortType_Speed_EURange":
		return OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed_EURange, true
	case "IIeeeBaseEthernetPortType_Speed_EngineeringUnits":
		return OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed_EngineeringUnits, true
	case "IIeeeBaseEthernetPortType_Duplex":
		return OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Duplex, true
	case "IIeeeBaseEthernetPortType_MaxFrameLength":
		return OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_MaxFrameLength, true
	case "IBaseEthernetCapabilitiesType_VlanTagCapable":
		return OpcuaNodeIdServicesVariableI_IBaseEthernetCapabilitiesType_VlanTagCapable, true
	case "ISrClassType_Id":
		return OpcuaNodeIdServicesVariableI_ISrClassType_Id, true
	case "ISrClassType_Priority":
		return OpcuaNodeIdServicesVariableI_ISrClassType_Priority, true
	case "ISrClassType_Vid":
		return OpcuaNodeIdServicesVariableI_ISrClassType_Vid, true
	case "IIeeeBaseTsnStreamType_StreamId":
		return OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStreamType_StreamId, true
	case "IIeeeBaseTsnStreamType_StreamName":
		return OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStreamType_StreamName, true
	case "IIeeeBaseTsnStreamType_State":
		return OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStreamType_State, true
	case "IIeeeBaseTsnStreamType_AccumulatedLatency":
		return OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStreamType_AccumulatedLatency, true
	case "IIeeeBaseTsnStreamType_SrClassId":
		return OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStreamType_SrClassId, true
	case "IIeeeBaseTsnTrafficSpecificationType_MaxIntervalFrames":
		return OpcuaNodeIdServicesVariableI_IIeeeBaseTsnTrafficSpecificationType_MaxIntervalFrames, true
	case "IIeeeBaseTsnTrafficSpecificationType_MaxFrameSize":
		return OpcuaNodeIdServicesVariableI_IIeeeBaseTsnTrafficSpecificationType_MaxFrameSize, true
	case "IIeeeBaseTsnTrafficSpecificationType_Interval":
		return OpcuaNodeIdServicesVariableI_IIeeeBaseTsnTrafficSpecificationType_Interval, true
	case "IIeeeBaseTsnStatusStreamType_TalkerStatus":
		return OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStatusStreamType_TalkerStatus, true
	case "IIeeeBaseTsnStatusStreamType_ListenerStatus":
		return OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStatusStreamType_ListenerStatus, true
	case "IIeeeBaseTsnStatusStreamType_FailureCode":
		return OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStatusStreamType_FailureCode, true
	case "IIeeeBaseTsnStatusStreamType_FailureSystemIdentifier":
		return OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStatusStreamType_FailureSystemIdentifier, true
	case "IIeeeTsnInterfaceConfigurationType_MacAddress":
		return OpcuaNodeIdServicesVariableI_IIeeeTsnInterfaceConfigurationType_MacAddress, true
	case "IIeeeTsnInterfaceConfigurationType_InterfaceName":
		return OpcuaNodeIdServicesVariableI_IIeeeTsnInterfaceConfigurationType_InterfaceName, true
	case "IIeeeTsnInterfaceConfigurationTalkerType_TimeAwareOffset":
		return OpcuaNodeIdServicesVariableI_IIeeeTsnInterfaceConfigurationTalkerType_TimeAwareOffset, true
	case "IIeeeTsnInterfaceConfigurationListenerType_ReceiveOffset":
		return OpcuaNodeIdServicesVariableI_IIeeeTsnInterfaceConfigurationListenerType_ReceiveOffset, true
	case "IIeeeTsnMacAddressType_DestinationAddress":
		return OpcuaNodeIdServicesVariableI_IIeeeTsnMacAddressType_DestinationAddress, true
	case "IIeeeTsnMacAddressType_SourceAddress":
		return OpcuaNodeIdServicesVariableI_IIeeeTsnMacAddressType_SourceAddress, true
	case "IIeeeTsnVlanTagType_VlanId":
		return OpcuaNodeIdServicesVariableI_IIeeeTsnVlanTagType_VlanId, true
	case "IIeeeTsnVlanTagType_PriorityCodePoint":
		return OpcuaNodeIdServicesVariableI_IIeeeTsnVlanTagType_PriorityCodePoint, true
	case "IPriorityMappingEntryType_MappingUri":
		return OpcuaNodeIdServicesVariableI_IPriorityMappingEntryType_MappingUri, true
	case "IPriorityMappingEntryType_PriorityLabel":
		return OpcuaNodeIdServicesVariableI_IPriorityMappingEntryType_PriorityLabel, true
	case "IPriorityMappingEntryType_PriorityValue_PCP":
		return OpcuaNodeIdServicesVariableI_IPriorityMappingEntryType_PriorityValue_PCP, true
	case "IPriorityMappingEntryType_PriorityValue_DSCP":
		return OpcuaNodeIdServicesVariableI_IPriorityMappingEntryType_PriorityValue_DSCP, true
	case "IIeeeAutoNegotiationStatusType_NegotiationStatus":
		return OpcuaNodeIdServicesVariableI_IIeeeAutoNegotiationStatusType_NegotiationStatus, true
	case "IVlanIdType_VlanId":
		return OpcuaNodeIdServicesVariableI_IVlanIdType_VlanId, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableIKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableIValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableI(structType any) OpcuaNodeIdServicesVariableI {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableI {
		if sOpcuaNodeIdServicesVariableI, ok := typ.(OpcuaNodeIdServicesVariableI); ok {
			return sOpcuaNodeIdServicesVariableI
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableI) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableI) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableIParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableI, error) {
	return OpcuaNodeIdServicesVariableIParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableIParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableI, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableI", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableI")
	}
	if enum, ok := OpcuaNodeIdServicesVariableIByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableI")
		return OpcuaNodeIdServicesVariableI(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableI) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableI) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableI", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableI) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableI_IOrderedObjectType_NumberInList:
		return "IOrderedObjectType_NumberInList"
	case OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_AdminStatus:
		return "IIetfBaseNetworkInterfaceType_AdminStatus"
	case OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_OperStatus:
		return "IIetfBaseNetworkInterfaceType_OperStatus"
	case OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_PhysAddress:
		return "IIetfBaseNetworkInterfaceType_PhysAddress"
	case OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed:
		return "IIetfBaseNetworkInterfaceType_Speed"
	case OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed_Definition:
		return "IIetfBaseNetworkInterfaceType_Speed_Definition"
	case OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed_ValuePrecision:
		return "IIetfBaseNetworkInterfaceType_Speed_ValuePrecision"
	case OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed_InstrumentRange:
		return "IIetfBaseNetworkInterfaceType_Speed_InstrumentRange"
	case OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed_EURange:
		return "IIetfBaseNetworkInterfaceType_Speed_EURange"
	case OpcuaNodeIdServicesVariableI_IIetfBaseNetworkInterfaceType_Speed_EngineeringUnits:
		return "IIetfBaseNetworkInterfaceType_Speed_EngineeringUnits"
	case OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed:
		return "IIeeeBaseEthernetPortType_Speed"
	case OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed_Definition:
		return "IIeeeBaseEthernetPortType_Speed_Definition"
	case OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed_ValuePrecision:
		return "IIeeeBaseEthernetPortType_Speed_ValuePrecision"
	case OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed_InstrumentRange:
		return "IIeeeBaseEthernetPortType_Speed_InstrumentRange"
	case OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed_EURange:
		return "IIeeeBaseEthernetPortType_Speed_EURange"
	case OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Speed_EngineeringUnits:
		return "IIeeeBaseEthernetPortType_Speed_EngineeringUnits"
	case OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_Duplex:
		return "IIeeeBaseEthernetPortType_Duplex"
	case OpcuaNodeIdServicesVariableI_IIeeeBaseEthernetPortType_MaxFrameLength:
		return "IIeeeBaseEthernetPortType_MaxFrameLength"
	case OpcuaNodeIdServicesVariableI_IBaseEthernetCapabilitiesType_VlanTagCapable:
		return "IBaseEthernetCapabilitiesType_VlanTagCapable"
	case OpcuaNodeIdServicesVariableI_ISrClassType_Id:
		return "ISrClassType_Id"
	case OpcuaNodeIdServicesVariableI_ISrClassType_Priority:
		return "ISrClassType_Priority"
	case OpcuaNodeIdServicesVariableI_ISrClassType_Vid:
		return "ISrClassType_Vid"
	case OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStreamType_StreamId:
		return "IIeeeBaseTsnStreamType_StreamId"
	case OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStreamType_StreamName:
		return "IIeeeBaseTsnStreamType_StreamName"
	case OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStreamType_State:
		return "IIeeeBaseTsnStreamType_State"
	case OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStreamType_AccumulatedLatency:
		return "IIeeeBaseTsnStreamType_AccumulatedLatency"
	case OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStreamType_SrClassId:
		return "IIeeeBaseTsnStreamType_SrClassId"
	case OpcuaNodeIdServicesVariableI_IIeeeBaseTsnTrafficSpecificationType_MaxIntervalFrames:
		return "IIeeeBaseTsnTrafficSpecificationType_MaxIntervalFrames"
	case OpcuaNodeIdServicesVariableI_IIeeeBaseTsnTrafficSpecificationType_MaxFrameSize:
		return "IIeeeBaseTsnTrafficSpecificationType_MaxFrameSize"
	case OpcuaNodeIdServicesVariableI_IIeeeBaseTsnTrafficSpecificationType_Interval:
		return "IIeeeBaseTsnTrafficSpecificationType_Interval"
	case OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStatusStreamType_TalkerStatus:
		return "IIeeeBaseTsnStatusStreamType_TalkerStatus"
	case OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStatusStreamType_ListenerStatus:
		return "IIeeeBaseTsnStatusStreamType_ListenerStatus"
	case OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStatusStreamType_FailureCode:
		return "IIeeeBaseTsnStatusStreamType_FailureCode"
	case OpcuaNodeIdServicesVariableI_IIeeeBaseTsnStatusStreamType_FailureSystemIdentifier:
		return "IIeeeBaseTsnStatusStreamType_FailureSystemIdentifier"
	case OpcuaNodeIdServicesVariableI_IIeeeTsnInterfaceConfigurationType_MacAddress:
		return "IIeeeTsnInterfaceConfigurationType_MacAddress"
	case OpcuaNodeIdServicesVariableI_IIeeeTsnInterfaceConfigurationType_InterfaceName:
		return "IIeeeTsnInterfaceConfigurationType_InterfaceName"
	case OpcuaNodeIdServicesVariableI_IIeeeTsnInterfaceConfigurationTalkerType_TimeAwareOffset:
		return "IIeeeTsnInterfaceConfigurationTalkerType_TimeAwareOffset"
	case OpcuaNodeIdServicesVariableI_IIeeeTsnInterfaceConfigurationListenerType_ReceiveOffset:
		return "IIeeeTsnInterfaceConfigurationListenerType_ReceiveOffset"
	case OpcuaNodeIdServicesVariableI_IIeeeTsnMacAddressType_DestinationAddress:
		return "IIeeeTsnMacAddressType_DestinationAddress"
	case OpcuaNodeIdServicesVariableI_IIeeeTsnMacAddressType_SourceAddress:
		return "IIeeeTsnMacAddressType_SourceAddress"
	case OpcuaNodeIdServicesVariableI_IIeeeTsnVlanTagType_VlanId:
		return "IIeeeTsnVlanTagType_VlanId"
	case OpcuaNodeIdServicesVariableI_IIeeeTsnVlanTagType_PriorityCodePoint:
		return "IIeeeTsnVlanTagType_PriorityCodePoint"
	case OpcuaNodeIdServicesVariableI_IPriorityMappingEntryType_MappingUri:
		return "IPriorityMappingEntryType_MappingUri"
	case OpcuaNodeIdServicesVariableI_IPriorityMappingEntryType_PriorityLabel:
		return "IPriorityMappingEntryType_PriorityLabel"
	case OpcuaNodeIdServicesVariableI_IPriorityMappingEntryType_PriorityValue_PCP:
		return "IPriorityMappingEntryType_PriorityValue_PCP"
	case OpcuaNodeIdServicesVariableI_IPriorityMappingEntryType_PriorityValue_DSCP:
		return "IPriorityMappingEntryType_PriorityValue_DSCP"
	case OpcuaNodeIdServicesVariableI_IIeeeAutoNegotiationStatusType_NegotiationStatus:
		return "IIeeeAutoNegotiationStatusType_NegotiationStatus"
	case OpcuaNodeIdServicesVariableI_IVlanIdType_VlanId:
		return "IVlanIdType_VlanId"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableI) String() string {
	return e.PLC4XEnumName()
}
