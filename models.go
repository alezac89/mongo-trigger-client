package mongoCloud

type Function struct {
	ID string `json:"_id"`
	Name string `json:"name"`
	Source string `json:"source"`
	Private bool `json:"private"`
	LastModified int64 `json:"last_modified"`
	ReadOnly bool `json:"read_only"`
}

type FunctionDTO struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Source string `json:"source"`
	Private bool `json:"private"`
	LastModified int64 `json:"last_modified"`
	ReadOnly bool `json:"read_only"`
}

type EventSubscriber struct {
	ID string `json:"_id"`
	Name string `json:"name"`
	Type string `json:"type"`
	FunctionID string `json:"function_id"`
	FunctionName string `json:"function_name"`
	Disabled bool `json:"disabled"`
	Config Config `json:"config"`
}

type EventSubscriberDTO struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	FunctionID string `json:"function_id"`
	FunctionName string `json:"function_name"`
	Disabled bool `json:"disabled"`
	Config Config `json:"config"`
}


type Config struct {
	Schedule string  `json:"schedule"`
	ScheduleType string `json:"schedule_type"`
	SkipCatchupEvents bool `json:"skip_catchup_events"`
}

//https://services.cloud.mongodb.com/api/admin/v3.0/groups/6499158adcf4dc6655293bbc/apps/65ebc1e4e0124310c57645ee/event_subscriptions

// {"name":"trigger0","type":"SCHEDULED","function_id":"65ef0d9f95f3de8dcfbcc6d3",
//"function_name":"Atlas_Triggers_trigger0_1710165406","disabled":false,
//"config":{"schedule":"*/5 * * * *","schedule_type":"BASIC","skip_catchup_events":false}}