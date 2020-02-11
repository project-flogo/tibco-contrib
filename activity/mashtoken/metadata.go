package mashtoken

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct{}

type Input struct {
	Username  string `md:"username"`
	Password  string `md:"password"`
	Scope     string `md:"scope"`
	BasicAuth string `md:"basicauth"`
}

type Output struct {
	AccessToken  string `md:"accesstoken"`
	TokenType    string `md:"tokentype"`
	ExpiresIn    string `md:"expiresin"`
	RefreshToken string `md:"refreshtoken"`
	Scope        string `md:"scope"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"username":  i.Username,
		"password":  i.Password,
		"scope":     i.Scope,
		"basicauth": i.BasicAuth,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {
	var err error

	i.Username, err = coerce.ToString(values["username"])
	if err != nil {
		return err
	}

	i.Password, err = coerce.ToString(values["password"])
	if err != nil {
		return err
	}

	i.Scope, err = coerce.ToString(values["scope"])
	if err != nil {
		return err
	}

	i.BasicAuth, err = coerce.ToString(values["basicauth"])
	if err != nil {
		return err
	}

	return err
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"accesstoken":  o.AccessToken,
		"tokentype":    o.TokenType,
		"expiresin":    o.ExpiresIn,
		"refreshtoken": o.RefreshToken,
		"scope":        o.Scope,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {
	var err error

	o.AccessToken, err = coerce.ToString(values["accesstoken"])
	if err != nil {
		return err
	}

	o.TokenType, err = coerce.ToString(values["tokentype"])
	if err != nil {
		return err
	}

	o.ExpiresIn, err = coerce.ToString(values["expiresin"])
	if err != nil {
		return err
	}

	o.RefreshToken, err = coerce.ToString(values["refreshtoken"])
	if err != nil {
		return err
	}

	o.Scope, err = coerce.ToString(values["scope"])
	if err != nil {
		return err
	}

	return err
}
