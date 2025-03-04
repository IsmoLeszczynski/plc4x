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
package org.apache.plc4x.java.s7discovery.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public abstract class PnDcp_Block implements Message {

  // Abstract accessors for discriminator values.
  public abstract PnDcp_BlockOptions getOption();

  public abstract Short getSuboption();

  public PnDcp_Block() {
    super();
  }

  protected abstract void serializePnDcp_BlockChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("PnDcp_Block");

    // Discriminator Field (option) (Used as input to a switch field)
    writeDiscriminatorEnumField(
        "option",
        "PnDcp_BlockOptions",
        getOption(),
        writeEnum(
            PnDcp_BlockOptions::getValue,
            PnDcp_BlockOptions::name,
            writeUnsignedShort(writeBuffer, 8)),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Discriminator Field (suboption) (Used as input to a switch field)
    writeDiscriminatorField(
        "suboption",
        getSuboption(),
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Implicit Field (blockLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int blockLength = (int) ((getLengthInBytes()) - (4));
    writeImplicitField(
        "blockLength",
        blockLength,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Switch field (Serialize the sub-type)
    serializePnDcp_BlockChild(writeBuffer);

    writeBuffer.popContext("PnDcp_Block");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    PnDcp_Block _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Discriminator Field (option)
    lengthInBits += 8;

    // Discriminator Field (suboption)
    lengthInBits += 8;

    // Implicit Field (blockLength)
    lengthInBits += 16;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static PnDcp_Block staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("PnDcp_Block");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    PnDcp_BlockOptions option =
        readDiscriminatorEnumField(
            "option",
            "PnDcp_BlockOptions",
            readEnum(PnDcp_BlockOptions::enumForValue, readUnsignedShort(readBuffer, 8)),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short suboption =
        readDiscriminatorField(
            "suboption",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int blockLength =
        readImplicitField(
            "blockLength",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    PnDcp_BlockBuilder builder = null;
    if (EvaluationHelper.equals(option, PnDcp_BlockOptions.IP_OPTION)
        && EvaluationHelper.equals(suboption, (short) 1)) {
      builder = PnDcp_Block_IpMacAddress.staticParsePnDcp_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.IP_OPTION)
        && EvaluationHelper.equals(suboption, (short) 2)) {
      builder = PnDcp_Block_IpParameter.staticParsePnDcp_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.IP_OPTION)
        && EvaluationHelper.equals(suboption, (short) 3)) {
      builder = PnDcp_Block_FullIpSuite.staticParsePnDcp_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.DEVICE_PROPERTIES_OPTION)
        && EvaluationHelper.equals(suboption, (short) 1)) {
      builder =
          PnDcp_Block_DevicePropertiesDeviceVendor.staticParsePnDcp_BlockBuilder(
              readBuffer, blockLength);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.DEVICE_PROPERTIES_OPTION)
        && EvaluationHelper.equals(suboption, (short) 2)) {
      builder =
          PnDcp_Block_DevicePropertiesNameOfStation.staticParsePnDcp_BlockBuilder(
              readBuffer, blockLength);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.DEVICE_PROPERTIES_OPTION)
        && EvaluationHelper.equals(suboption, (short) 3)) {
      builder = PnDcp_Block_DevicePropertiesDeviceId.staticParsePnDcp_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.DEVICE_PROPERTIES_OPTION)
        && EvaluationHelper.equals(suboption, (short) 4)) {
      builder = PnDcp_Block_DevicePropertiesDeviceRole.staticParsePnDcp_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.DEVICE_PROPERTIES_OPTION)
        && EvaluationHelper.equals(suboption, (short) 5)) {
      builder =
          PnDcp_Block_DevicePropertiesDeviceOptions.staticParsePnDcp_BlockBuilder(
              readBuffer, blockLength);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.DEVICE_PROPERTIES_OPTION)
        && EvaluationHelper.equals(suboption, (short) 6)) {
      builder =
          PnDcp_Block_DevicePropertiesAliasName.staticParsePnDcp_BlockBuilder(
              readBuffer, blockLength);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.DEVICE_PROPERTIES_OPTION)
        && EvaluationHelper.equals(suboption, (short) 7)) {
      builder =
          PnDcp_Block_DevicePropertiesDeviceInstance.staticParsePnDcp_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.DEVICE_PROPERTIES_OPTION)
        && EvaluationHelper.equals(suboption, (short) 8)) {
      builder = PnDcp_Block_DevicePropertiesOemDeviceId.staticParsePnDcp_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.DEVICE_PROPERTIES_OPTION)
        && EvaluationHelper.equals(suboption, (short) 9)) {
      builder =
          PnDcp_Block_DevicePropertiesStandardGateway.staticParsePnDcp_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.DCP_OPTION)
        && EvaluationHelper.equals(suboption, (short) 12)) {
      builder = PnDcp_Block_DhcpOptionHostName.staticParsePnDcp_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.DCP_OPTION)
        && EvaluationHelper.equals(suboption, (short) 43)) {
      builder =
          PnDcp_Block_DhcpOptionVendorSpecificInformation.staticParsePnDcp_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.DCP_OPTION)
        && EvaluationHelper.equals(suboption, (short) 54)) {
      builder = PnDcp_Block_DhcpOptionServerIdentifier.staticParsePnDcp_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.DCP_OPTION)
        && EvaluationHelper.equals(suboption, (short) 55)) {
      builder =
          PnDcp_Block_DhcpOptionParameterRequestList.staticParsePnDcp_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.DCP_OPTION)
        && EvaluationHelper.equals(suboption, (short) 60)) {
      builder = PnDcp_Block_DhcpOptionClassIdentifier.staticParsePnDcp_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.DCP_OPTION)
        && EvaluationHelper.equals(suboption, (short) 61)) {
      builder =
          PnDcp_Block_DhcpOptionDhcpClientIdentifier.staticParsePnDcp_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.DCP_OPTION)
        && EvaluationHelper.equals(suboption, (short) 81)) {
      builder =
          PnDcp_Block_DhcpOptionFullyQualifiedDomainName.staticParsePnDcp_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.DCP_OPTION)
        && EvaluationHelper.equals(suboption, (short) 97)) {
      builder = PnDcp_Block_DhcpOptionUuidBasedClient.staticParsePnDcp_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.CONTROL_OPTION)
        && EvaluationHelper.equals(suboption, (short) 1)) {
      builder = PnDcp_Block_ControlOptionStart.staticParsePnDcp_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.CONTROL_OPTION)
        && EvaluationHelper.equals(suboption, (short) 2)) {
      builder = PnDcp_Block_ControlOptionStop.staticParsePnDcp_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.CONTROL_OPTION)
        && EvaluationHelper.equals(suboption, (short) 3)) {
      builder = PnDcp_Block_ControlOptionSignal.staticParsePnDcp_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.CONTROL_OPTION)
        && EvaluationHelper.equals(suboption, (short) 4)) {
      builder = PnDcp_Block_ControlOptionResponse.staticParsePnDcp_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.CONTROL_OPTION)
        && EvaluationHelper.equals(suboption, (short) 5)) {
      builder = PnDcp_Block_ControlOptionFactoryReset.staticParsePnDcp_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.CONTROL_OPTION)
        && EvaluationHelper.equals(suboption, (short) 6)) {
      builder = PnDcp_Block_ControlOptionResetToFactory.staticParsePnDcp_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.DEVICE_INITIATIVE_OPTION)
        && EvaluationHelper.equals(suboption, (short) 1)) {
      builder = PnDcp_Block_DeviceInitiativeOption.staticParsePnDcp_BlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(option, PnDcp_BlockOptions.ALL_SELECTOR_OPTION)
        && EvaluationHelper.equals(suboption, (short) 0xFF)) {
      builder = PnDcp_Block_ALLSelector.staticParsePnDcp_BlockBuilder(readBuffer);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "option="
              + option
              + " "
              + "suboption="
              + suboption
              + "]");
    }

    readBuffer.closeContext("PnDcp_Block");
    // Create the instance
    PnDcp_Block _pnDcp_Block = builder.build();
    return _pnDcp_Block;
  }

  public interface PnDcp_BlockBuilder {
    PnDcp_Block build();
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PnDcp_Block)) {
      return false;
    }
    PnDcp_Block that = (PnDcp_Block) o;
    return true;
  }

  @Override
  public int hashCode() {
    return Objects.hash();
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
