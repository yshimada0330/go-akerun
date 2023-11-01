package akerun

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

type Akerun struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	ImageURL            string `json:"image_url"`
	OpenDoorAlert       bool   `json:"open_door_alert"`
	OpenDoorAlertSecond int    `json:"open_door_alert_second"`
	PushButton          bool   `json:"push_button"`
	NormalSoundVolume   int    `json:"normal_sound_volume"`
	AlertSoundVolume    int    `json:"alert_sound_volume"`
	BatteryPercentage   int    `json:"battery_percentage"`
	Autolock            bool   `json:"autolock"`
	AutolockOffSchedule struct {
		StartTime  string `json:"start_time"`
		EndTime    string `json:"end_time"`
		DaysOfWeek []int  `json:"days_of_week"`
	} `json:"autolock_off_schedule"`
	AkerunRemote struct {
		ID string `json:"id"`
	} `json:"akerun_remote"`
	NFCReaderInside struct {
		ID                string `json:"id"`
		BatteryPercentage int    `json:"battery_percentage"`
	} `json:"nfc_reader_inside"`
	NFCReaderOutside struct {
		ID                string `json:"id"`
		BatteryPercentage int    `json:"battery_percentage"`
	} `json:"nfc_reader_outside"`
	DoorSensor struct {
		ID                string `json:"id"`
		BatteryPercentage int    `json:"battery_percentage"`
	} `json:"door_sensor"`
}

type AkerunList struct {
	Akeruns []Akerun `json:"akeruns"`
}

type AkerunListParameter struct {
	Limit     uint32   `url:"limit,omitempty"`
	AkerunIds []string `url:"akerun_ids[],omitempty"`
	IdAfter   string   `url:"id_after,omitempty"`
	IdBefore  string   `url:"id_before,omitempty"`
}

func (c *Client) GetAkeruns(ctx context.Context, oauth2Token *oauth2.Token, organizationId string, params AkerunListParameter) (*AkerunList, error) {
	var result AkerunList
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("organizations/%s/akeruns", organizationId)
	err = c.callVersion(ctx, url, http.MethodGet, oauth2Token, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
