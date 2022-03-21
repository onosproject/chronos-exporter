// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package rasa_action

type slots struct {
	SlotName string `json:"slot_name"`
}

type entities struct {
	Start      int    `json:"start"`
	End        int    `json:"end"`
	Value      string `json:"value"`
	Entity     string `json:"entity"`
	Confidence int    `json:"confidence"`
}

type intent struct {
	Confidence float64 `json:"confidence"`
	Name       string  `json:"name"`
}

type intent_ranking struct {
	Confidence float64 `json:"confidence"`
	Name       string  `json:"name"`
}
type latest_message struct {
	Entities []entities `json:"entities"`
	Intent intent `json:"intent"`
	IntentRanking []intent_ranking `json:"intent_ranking"`
	Text string `json:"text"`
}

type events struct {
	Event     string `json:"event"`
	Timestamp int    `json:"timestamp"`
}

type latest_action struct {
	ActionName string `json:"action_name"`
	ActionText string `json:"action_text"`
}

type active_loop struct {
	Name string `json:"name"`
}

type tracker struct {
	ConversationID string `json:"conversation_id"`
	Slots slots `json:"slots"`
	LatestMessage latest_message `json:"latest_message"`
	LatestEventTime float64 `json:"latest_event_time"`
	FollowupAction  string  `json:"followup_action"`
	Paused          bool    `json:"paused"`
	Events []events `json:"events"`
	LatestInputChannel string `json:"latest_input_channel"`
	LatestActionName   string `json:"latest_action_name"`
	LatestAction latest_action `json:"latest_action"`
	ActiveLoop active_loop `json:"active_loop"`
}

type config struct {
	StoreEntitiesAsSlots bool `json:"store_entities_as_slots"`
}

type property1 struct {
	UseEntities bool `json:"use_entities"`
}

type property2 struct {
	UseEntities bool `json:"use_entities"`
}

type intents struct {
	Property1 property1 `json:"property1"`
	Property2 property2 `json:"property2"`
}

type domain struct {
	Config config `json:"config"`
	Intents []intents `json:"intents"`
	Entities []string `json:"entities"`
	Slots    struct {
		Property1 struct {
			AutoFill     bool     `json:"auto_fill"`
			InitialValue string   `json:"initial_value"`
			Type         string   `json:"type"`
			Values       []string `json:"values"`
		} `json:"property1"`
		Property2 struct {
			AutoFill     bool     `json:"auto_fill"`
			InitialValue string   `json:"initial_value"`
			Type         string   `json:"type"`
			Values       []string `json:"values"`
		} `json:"property2"`
	} `json:"slots"`
	Responses struct {
		Property1 []struct {
			Text string `json:"text"`
		} `json:"property1"`
		Property2 []struct {
			Text string `json:"text"`
		} `json:"property2"`
	} `json:"responses"`
	Actions []string `json:"actions"`
}

type payload struct {
	NextAction string `json:"next_action"`
	SenderID   string `json:"sender_id"`
	Tracker tracker `json:"tracker"`
	Domain domain `json:"domain"`
}

type response struct {
	Events []events `json:"events"`
	Responses []struct {
		Text    string `json:"text"`
		Buttons []struct {
			Title   string `json:"title"`
			Payload string `json:"payload"`
		} `json:"buttons"`
		Elements []interface{} `json:"elements"`
		Custom   struct {
		} `json:"custom"`
		Image      string `json:"image"`
		Attachment string `json:"attachment"`
		Template   string `json:"template"`
		Property1  string `json:"property1"`
		Property2  string `json:"property2"`
		Response   string `json:"response"`
	} `json:"responses"`
}