package controller

import (
	"fmt"
	"messenger/internal/utils"
	"net/http"

	"github.com/rs/zerolog/log"
)

func (hm *HandlersManager) RegistrationHandler(username, password string) {
	type Response struct {
		Token string `json:"token"`
	}

	var resp Response
	response, err := hm.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"username": username,
			"password": password,
		}).
		SetResult(&resp).
		Post(fmt.Sprintf("https://%s/register", hm.addr))

	if err != nil {
		log.Error().Err(err).Msg("cant request registration")
		return
	}

	if response.StatusCode() != http.StatusOK {
		log.Error().Msg(string(response.Body()))
		return
	}

	utils.WriteToken(resp.Token)
}

func (hm *HandlersManager) LoginHandler(username, password string) {
	type Response struct {
		Token string `json:"token"`
	}

	var resp Response
	response, err := hm.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"username": username,
			"password": password,
		}).
		SetResult(&resp).
		Post(fmt.Sprintf("https://%s/login", hm.addr))

	if err != nil {
		log.Error().Err(err).Msg("cant request login")
		return
	}

	if response.StatusCode() != http.StatusOK {
		log.Error().Msg(string(response.Body()))
		return
	}

	utils.WriteToken(resp.Token)
}

func (hm *HandlersManager) ValidateTokenHandler() bool {
	token := utils.GetToken()

	response, err := hm.client.R().
		SetHeader("Authorization", token).
		Get(fmt.Sprintf("https://%s/token", hm.addr))

	if err != nil {
		log.Error().Err(err).Msg("cant request token validation")
		return false
	}
	return response.StatusCode() == http.StatusOK
}
