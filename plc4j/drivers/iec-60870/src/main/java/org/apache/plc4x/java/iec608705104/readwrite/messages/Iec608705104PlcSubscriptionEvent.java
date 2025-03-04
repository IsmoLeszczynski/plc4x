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

package org.apache.plc4x.java.iec608705104.readwrite.messages;

import org.apache.plc4x.java.api.model.PlcTag;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.spi.messages.DefaultPlcSubscriptionEvent;
import org.apache.plc4x.java.spi.messages.utils.PlcResponseItem;

import java.time.Instant;
import java.util.Map;

public class Iec608705104PlcSubscriptionEvent extends DefaultPlcSubscriptionEvent {

    private final Map<String, PlcTag> tags;
    public Iec608705104PlcSubscriptionEvent(Instant timestamp, Map<String, PlcTag> tags, Map<String, PlcResponseItem<PlcValue>> values) {
        super(timestamp, values);
        this.tags = tags;
    }

    @Override
    public PlcTag getTag(String name) {
        return tags.get(name);
    }

}
