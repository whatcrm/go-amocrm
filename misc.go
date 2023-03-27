package go_amocrm

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"reflect"
	"regexp"
	"strings"
)

func (api *API) getAgent(method, baseURL string, params *Params) (*fiber.Agent, *fiber.Request) {
	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod(method)
	req.Header.SetContentType(fiber.MIMEApplicationJSON)
	req.Header.SetCanonical([]byte("Authorization"), []byte("Bearer "+api.AccessToken))
	req = isParams(req, api.Domain, baseURL, params)

	return a, req
}

func (api *API) makeRequest(options makeRequestOptions) (err error) {

	a, req := api.getAgent(options.Method, options.BaseURL, options.Params)

	if options.In != nil {
		api.log("marshaling the data...")
		req, err = marshal(options.In, req)
		if err != nil {
			return
		}
	}

	api.log("sending the data...")
	if err = a.Parse(); err != nil {
		log.Println(err)
		return
	}

	api.log("getting the answer...")
	_, body, errs := a.Bytes()
	if errs != nil {
		log.Println(errs)
		err = errs[0]
		return
	}

	if err = json.Unmarshal(body, options.Out); err != nil {
		log.Println(err)
		return
	}

	api.log("BODY: ", string(body))
	return
}

func marshal(data interface{}, req *fiber.Request) (*fiber.Request, error) {
	m, err := json.Marshal(&data)
	if err != nil {
		log.Println(err)
		return req, err
	}

	log.Println(data)
	req.SetBody(m)
	return req, nil
}

func isParams(req *fiber.Request, domain string, parameter string, params *Params) *fiber.Request {
	if params == nil {
		req.SetRequestURI(httpURL + domain + "/" + parameter)
		if parameter == "" {
			req.SetRequestURI(httpURL + domain)
		}
		log.Println(httpURL + domain + "/" + parameter)
		return req
	}

	var u *url.URL
	if parameter != "" {
		u, _ = url.Parse(httpURL + domain + "/" + parameter)
	} else {
		u, _ = url.Parse(httpURL + domain)
	}
	q := u.Query()

	s := reflect.TypeOf(params)
	v := reflect.ValueOf(params)
	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		value := v.Field(i).String()
		if value != "" {
			q.Set(strings.ToLower(field.Name), value)
		}
	}
	u.RawQuery = q.Encode()
	fmt.Println(u.String())
	req.SetRequestURI(u.String())
	return req
}

func (api *API) log(message ...string) {
	var res string
	for _, msg := range message {
		res += msg
	}

	if api.Debug == true {
		log.Println(res)
	}
}

func isRegex(text string) bool {
	re := regexp.MustCompile("^[a-zA-z]{1,}\\.(bitrix|amocrm)\\.(ru|com|kz|kg)")
	return re.MatchString(text)
}
