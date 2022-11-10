package cloud

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

func configureTokenFromFile(tokenPath string) (*AccessToken, error) {
	var content struct {
		User struct {
			AccessToken string `yaml:"access_token"`
		} `yaml:"user"`
	}

	data, err := os.ReadFile(tokenPath)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(data, &content); err != nil {
		return nil, errors.Wrap(err, "invalid token file")
	}

	if content.User.AccessToken == "" {
		return nil, ErrEmptyToken
	}

	return &AccessToken{
		Token: content.User.AccessToken,
	}, nil
}
