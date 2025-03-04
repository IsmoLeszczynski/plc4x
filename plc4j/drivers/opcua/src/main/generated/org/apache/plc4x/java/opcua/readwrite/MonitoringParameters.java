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

public class MonitoringParameters extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 742;
  }

  // Properties.
  protected final long clientHandle;
  protected final double samplingInterval;
  protected final ExtensionObject filter;
  protected final long queueSize;
  protected final boolean discardOldest;

  public MonitoringParameters(
      long clientHandle,
      double samplingInterval,
      ExtensionObject filter,
      long queueSize,
      boolean discardOldest) {
    super();
    this.clientHandle = clientHandle;
    this.samplingInterval = samplingInterval;
    this.filter = filter;
    this.queueSize = queueSize;
    this.discardOldest = discardOldest;
  }

  public long getClientHandle() {
    return clientHandle;
  }

  public double getSamplingInterval() {
    return samplingInterval;
  }

  public ExtensionObject getFilter() {
    return filter;
  }

  public long getQueueSize() {
    return queueSize;
  }

  public boolean getDiscardOldest() {
    return discardOldest;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("MonitoringParameters");

    // Simple Field (clientHandle)
    writeSimpleField("clientHandle", clientHandle, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (samplingInterval)
    writeSimpleField("samplingInterval", samplingInterval, writeDouble(writeBuffer, 64));

    // Simple Field (filter)
    writeSimpleField("filter", filter, writeComplex(writeBuffer));

    // Simple Field (queueSize)
    writeSimpleField("queueSize", queueSize, writeUnsignedLong(writeBuffer, 32));

    // Reserved Field (reserved)
    writeReservedField("reserved", (byte) 0x00, writeUnsignedByte(writeBuffer, 7));

    // Simple Field (discardOldest)
    writeSimpleField("discardOldest", discardOldest, writeBoolean(writeBuffer));

    writeBuffer.popContext("MonitoringParameters");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    MonitoringParameters _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (clientHandle)
    lengthInBits += 32;

    // Simple field (samplingInterval)
    lengthInBits += 64;

    // Simple field (filter)
    lengthInBits += filter.getLengthInBits();

    // Simple field (queueSize)
    lengthInBits += 32;

    // Reserved Field (reserved)
    lengthInBits += 7;

    // Simple field (discardOldest)
    lengthInBits += 1;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("MonitoringParameters");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long clientHandle = readSimpleField("clientHandle", readUnsignedLong(readBuffer, 32));

    double samplingInterval = readSimpleField("samplingInterval", readDouble(readBuffer, 64));

    ExtensionObject filter =
        readSimpleField(
            "filter",
            readComplex(
                () -> ExtensionObject.staticParse(readBuffer, (boolean) (true)), readBuffer));

    long queueSize = readSimpleField("queueSize", readUnsignedLong(readBuffer, 32));

    Byte reservedField0 =
        readReservedField("reserved", readUnsignedByte(readBuffer, 7), (byte) 0x00);

    boolean discardOldest = readSimpleField("discardOldest", readBoolean(readBuffer));

    readBuffer.closeContext("MonitoringParameters");
    // Create the instance
    return new MonitoringParametersBuilderImpl(
        clientHandle, samplingInterval, filter, queueSize, discardOldest);
  }

  public static class MonitoringParametersBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final long clientHandle;
    private final double samplingInterval;
    private final ExtensionObject filter;
    private final long queueSize;
    private final boolean discardOldest;

    public MonitoringParametersBuilderImpl(
        long clientHandle,
        double samplingInterval,
        ExtensionObject filter,
        long queueSize,
        boolean discardOldest) {
      this.clientHandle = clientHandle;
      this.samplingInterval = samplingInterval;
      this.filter = filter;
      this.queueSize = queueSize;
      this.discardOldest = discardOldest;
    }

    public MonitoringParameters build() {
      MonitoringParameters monitoringParameters =
          new MonitoringParameters(
              clientHandle, samplingInterval, filter, queueSize, discardOldest);
      return monitoringParameters;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof MonitoringParameters)) {
      return false;
    }
    MonitoringParameters that = (MonitoringParameters) o;
    return (getClientHandle() == that.getClientHandle())
        && (getSamplingInterval() == that.getSamplingInterval())
        && (getFilter() == that.getFilter())
        && (getQueueSize() == that.getQueueSize())
        && (getDiscardOldest() == that.getDiscardOldest())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getClientHandle(),
        getSamplingInterval(),
        getFilter(),
        getQueueSize(),
        getDiscardOldest());
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
