package mongoCloud

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateTriggers - Create new trigger with function
func (c *Client) CreateTriggers(groupId string, appId string, f FunctionDTO, es EventSubscriberDTO) (*EventSubscriberDTO, error) {


	incomingfunction := Function{
		ID: f.ID,
		Name: f.Name,
		Source: f.Source,
		Private: f.Private,
		LastModified: f.LastModified,
		ReadOnly: f.ReadOnly,
	}

	rb, err := json.Marshal(incomingfunction)
	if err != nil {
		return nil, err
	}

	funReq, funErr := http.NewRequest("POST", fmt.Sprintf("%s/api/admin/v3.0/groups/%s/apps/%s/functions", c.HostURL, groupId, appId), strings.NewReader(string(rb)))
	if funErr != nil {
		return nil, funErr
	}

	funBody, funBodyErr := c.doRequest(funReq)
	if funBodyErr != nil {
		return nil, funBodyErr
	}

	function := Function{}
	funBodyErr = json.Unmarshal(funBody, &function)
	if funBodyErr != nil {
		return nil, funBodyErr
	}

	functionDTO := FunctionDTO{ID: function.ID, 
		Name: function.Name, 
		Source: function.Source, 
		Private: function.Private, LastModified: function.LastModified, ReadOnly: function.ReadOnly}


	incomingEventSubscriber := EventSubscriber{
		Name: es.Name,
		Type: es.Type,
		FunctionID: functionDTO.ID,
		FunctionName: functionDTO.Name,
		Disabled: es.Disabled,
		Config: Config{
			Schedule: es.Config.Schedule,
			ScheduleType: es.Config.ScheduleType,
			SkipCatchupEvents: es.Config.SkipCatchupEvents,
		},
	}

	esrb, err := json.Marshal(incomingEventSubscriber)
	if err != nil {
		return nil, err
	}

	esReq, esErr := http.NewRequest("POST", fmt.Sprintf("%s/api/admin/v3.0/groups/%s/apps/%s/event_subscriptions", c.HostURL, groupId, appId), strings.NewReader(string(esrb)))


	if esErr != nil {
		return nil, esErr
	}

	esBody, esBodyErr := c.doRequest(esReq)
	if esBodyErr != nil {
		return nil, esBodyErr
	}

	eventSubscriber := EventSubscriber{}
	esBodyErr = json.Unmarshal(esBody, &eventSubscriber)
	if esBodyErr != nil {
		return nil, esBodyErr
	}

	eventSubscriberDTO := EventSubscriberDTO{
		ID: eventSubscriber.ID,
		Name: eventSubscriber.Name,
		Type: eventSubscriber.Type,
		FunctionID: eventSubscriber.FunctionID,
		FunctionName: eventSubscriber.FunctionName,
		Disabled: eventSubscriber.Disabled,
		Config: Config{
			Schedule: eventSubscriber.Config.Schedule,
			ScheduleType: eventSubscriber.Config.ScheduleType,
			SkipCatchupEvents: eventSubscriber.Config.SkipCatchupEvents,
		},
	}

	return &eventSubscriberDTO, nil
}