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
package org.apache.plc4x.java.opcua.readwrite;

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

public class MonitoredItemModifyRequest extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 757;
  }

  // Properties.
  protected final long monitoredItemId;
  protected final MonitoringParameters requestedParameters;

  public MonitoredItemModifyRequest(
      long monitoredItemId, MonitoringParameters requestedParameters) {
    super();
    this.monitoredItemId = monitoredItemId;
    this.requestedParameters = requestedParameters;
  }

  public long getMonitoredItemId() {
    return monitoredItemId;
  }

  public MonitoringParameters getRequestedParameters() {
    return requestedParameters;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("MonitoredItemModifyRequest");

    // Simple Field (monitoredItemId)
    writeSimpleField("monitoredItemId", monitoredItemId, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (requestedParameters)
    writeSimpleField("requestedParameters", requestedParameters, writeComplex(writeBuffer));

    writeBuffer.popContext("MonitoredItemModifyRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    MonitoredItemModifyRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (monitoredItemId)
    lengthInBits += 32;

    // Simple field (requestedParameters)
    lengthInBits += requestedParameters.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("MonitoredItemModifyRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long monitoredItemId = readSimpleField("monitoredItemId", readUnsignedLong(readBuffer, 32));

    MonitoringParameters requestedParameters =
        readSimpleField(
            "requestedParameters",
            readComplex(
                () ->
                    (MonitoringParameters)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (742)),
                readBuffer));

    readBuffer.closeContext("MonitoredItemModifyRequest");
    // Create the instance
    return new MonitoredItemModifyRequestBuilderImpl(monitoredItemId, requestedParameters);
  }

  public static class MonitoredItemModifyRequestBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final long monitoredItemId;
    private final MonitoringParameters requestedParameters;

    public MonitoredItemModifyRequestBuilderImpl(
        long monitoredItemId, MonitoringParameters requestedParameters) {
      this.monitoredItemId = monitoredItemId;
      this.requestedParameters = requestedParameters;
    }

    public MonitoredItemModifyRequest build() {
      MonitoredItemModifyRequest monitoredItemModifyRequest =
          new MonitoredItemModifyRequest(monitoredItemId, requestedParameters);
      return monitoredItemModifyRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof MonitoredItemModifyRequest)) {
      return false;
    }
    MonitoredItemModifyRequest that = (MonitoredItemModifyRequest) o;
    return (getMonitoredItemId() == that.getMonitoredItemId())
        && (getRequestedParameters() == that.getRequestedParameters())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getMonitoredItemId(), getRequestedParameters());
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
