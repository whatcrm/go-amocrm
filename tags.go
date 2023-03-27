package go_amocrm

//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/gofiber/fiber/v2"
//	"log"
//	"whatcrm_notify/pkg/amocrm/models"
//)
//
//func (api *API) CreateTag(entity string, tag *[]models.Tag) (*models.TagResponse, error) {
//	api.log("EditTag request started...")
//
//	a, req := getAgent(fiber.MethodPost, api.AccessToken)
//	req.SetRequestURI(httpURL + api.Domain + "/" + noEntityURL + entity + "/tags")
//
//	api.log("marshalling and settings request body...")
//	req, err := marshal(tag, req)
//	if err != nil {
//		return nil, err
//	}
//
//	api.log("sending the request...")
//	body, errs := sendRequest(a)
//	if errs != nil {
//		return nil, fmt.Errorf("few errors: %v", errs)
//	}
//
//	api.log("unmarshalling the data...")
//	log.Println(string(body))
//	resp := models.TagResponse{}
//	if err = json.Unmarshal(body, &resp); err != nil {
//		return nil, err
//	}
//
//	api.log("returning the struct...")
//	log.Println(resp)
//
//	return &resp, nil
//}
