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

public class NodeAttributes extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 351;
  }

  // Properties.
  protected final long specifiedAttributes;
  protected final LocalizedText displayName;
  protected final LocalizedText description;
  protected final long writeMask;
  protected final long userWriteMask;

  public NodeAttributes(
      long specifiedAttributes,
      LocalizedText displayName,
      LocalizedText description,
      long writeMask,
      long userWriteMask) {
    super();
    this.specifiedAttributes = specifiedAttributes;
    this.displayName = displayName;
    this.description = description;
    this.writeMask = writeMask;
    this.userWriteMask = userWriteMask;
  }

  public long getSpecifiedAttributes() {
    return specifiedAttributes;
  }

  public LocalizedText getDisplayName() {
    return displayName;
  }

  public LocalizedText getDescription() {
    return description;
  }

  public long getWriteMask() {
    return writeMask;
  }

  public long getUserWriteMask() {
    return userWriteMask;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("NodeAttributes");

    // Simple Field (specifiedAttributes)
    writeSimpleField(
        "specifiedAttributes", specifiedAttributes, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (displayName)
    writeSimpleField("displayName", displayName, writeComplex(writeBuffer));

    // Simple Field (description)
    writeSimpleField("description", description, writeComplex(writeBuffer));

    // Simple Field (writeMask)
    writeSimpleField("writeMask", writeMask, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (userWriteMask)
    writeSimpleField("userWriteMask", userWriteMask, writeUnsignedLong(writeBuffer, 32));

    writeBuffer.popContext("NodeAttributes");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    NodeAttributes _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (specifiedAttributes)
    lengthInBits += 32;

    // Simple field (displayName)
    lengthInBits += displayName.getLengthInBits();

    // Simple field (description)
    lengthInBits += description.getLengthInBits();

    // Simple field (writeMask)
    lengthInBits += 32;

    // Simple field (userWriteMask)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("NodeAttributes");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long specifiedAttributes =
        readSimpleField("specifiedAttributes", readUnsignedLong(readBuffer, 32));

    LocalizedText displayName =
        readSimpleField(
            "displayName", readComplex(() -> LocalizedText.staticParse(readBuffer), readBuffer));

    LocalizedText description =
        readSimpleField(
            "description", readComplex(() -> LocalizedText.staticParse(readBuffer), readBuffer));

    long writeMask = readSimpleField("writeMask", readUnsignedLong(readBuffer, 32));

    long userWriteMask = readSimpleField("userWriteMask", readUnsignedLong(readBuffer, 32));

    readBuffer.closeContext("NodeAttributes");
    // Create the instance
    return new NodeAttributesBuilderImpl(
        specifiedAttributes, displayName, description, writeMask, userWriteMask);
  }

  public static class NodeAttributesBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final long specifiedAttributes;
    private final LocalizedText displayName;
    private final LocalizedText description;
    private final long writeMask;
    private final long userWriteMask;

    public NodeAttributesBuilderImpl(
        long specifiedAttributes,
        LocalizedText displayName,
        LocalizedText description,
        long writeMask,
        long userWriteMask) {
      this.specifiedAttributes = specifiedAttributes;
      this.displayName = displayName;
      this.description = description;
      this.writeMask = writeMask;
      this.userWriteMask = userWriteMask;
    }

    public NodeAttributes build() {
      NodeAttributes nodeAttributes =
          new NodeAttributes(
              specifiedAttributes, displayName, description, writeMask, userWriteMask);
      return nodeAttributes;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof NodeAttributes)) {
      return false;
    }
    NodeAttributes that = (NodeAttributes) o;
    return (getSpecifiedAttributes() == that.getSpecifiedAttributes())
        && (getDisplayName() == that.getDisplayName())
        && (getDescription() == that.getDescription())
        && (getWriteMask() == that.getWriteMask())
        && (getUserWriteMask() == that.getUserWriteMask())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getSpecifiedAttributes(),
        getDisplayName(),
        getDescription(),
        getWriteMask(),
        getUserWriteMask());
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
