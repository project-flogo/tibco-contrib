package mashtoken

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/project-flogo/core/activity"
)

func init() {
	_ = activity.Register(&Activity{}, New)
}

const (
	grantType      = "password"
	url            = "https://api.mashery.com/v3/token"
	ovToken        = "accesstoken"
	ovTokenType    = "tokentype"
	ovExpires      = "expiresin"
	ovRefreshToken = "refreshtoken"
	ovScope        = "scope"
)

type Activity struct {
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

func New(ctx activity.InitContext) (activity.Activity, error) {
	return &Activity{}, nil
}

func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {
	in := &Input{}
	err = ctx.GetInputObject(in)
	if err != nil {
		return false, err
	}

	// Encode BasicAuth
	encodedAuth := b64.StdEncoding.EncodeToString([]byte(in.BasicAuth))

	// Get the token from TIBCO Cloud Mashery
	payload := strings.NewReader(fmt.Sprintf("grant_type=%s&username=%s&password=%s&scope=%s", grantType, in.Username, in.Password, in.Scope))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return false, err
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("authorization", fmt.Sprintf("Basic %s", encodedAuth))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}

	// Set the output value in the context
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return false, err
	}
	ctx.SetOutput(ovExpires, data["expires_in"])
	ctx.SetOutput(ovRefreshToken, data["refresh_token"])
	ctx.SetOutput(ovScope, data["scope"])
	ctx.SetOutput(ovToken, data["access_token"])
	ctx.SetOutput(ovTokenType, data["token_type"])

	return true, nil
}
