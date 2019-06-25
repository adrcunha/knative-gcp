/*
Copyright 2019 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
)

// Topic implements pubsub.Client.Topic
func (c *PubSubClient) Topic(id string) Topic {
	return &PubSubTopic{topic: c.client.Topic(id)}
}

// PubSubTopic wrapps pubsub.Topic
type PubSubTopic struct {
	topic *pubsub.Topic
}

// Exists implements pubsub.Topic.Exists
func (t *PubSubTopic) Exists(ctx context.Context) (bool, error) {
	return t.topic.Exists(ctx)
}
