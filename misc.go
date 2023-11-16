package amocrm

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
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

	if baseURL == accountURL {
		req.SetRequestURI("https://www.amocrm.ru/oauth2/account/subdomain")
		return a, req
	} else {
		req = isParams(req, api.Domain, baseURL, params)
		return a, req
	}
}

func (api *API) makeRequest(options makeRequestOptions) (err error) {

	a, req := api.getAgent(options.Method, options.BaseURL, options.Params)
	api.log("req: %v", req)

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
	status, body, errs := a.Bytes()
	if errs != nil {
		log.Println(errs)
		err = errs[0]
		return
	}

	if body != nil && status != fiber.StatusCreated {
		api.log("BODY: ", string(body))
		if err = json.Unmarshal(body, options.Out); err != nil {
			log.Println(err)
			return fiber.NewError(400, string(body))
		}
	}

	if err = statusChecker(status); err != nil {
		return
	}
	api.log("statusChecker passed")
	return
}

func statusChecker(status int) error {
	switch status {
	case 400:
		return fiber.ErrBadRequest
	case 401:
		return fiber.ErrUnauthorized
	case 402:
		return fiber.ErrPaymentRequired
	case 403:
		return fiber.ErrForbidden
	case 404:
		return fiber.ErrNotFound
	case 201:
		return fiber.NewError(201, "Created")
	case 204:
		return fiber.NewError(204, "No content")
	case 200, 202, 302:
		return nil
	default:
		return fiber.NewError(status, "unknown status")
	}
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

func fixURI(domain, path string) (string, string) {
	//domain = " /https://checker.whatcrm.net//? "

	// fixing the domain name
	domain = strings.Trim(domain, " ")
	domain = strings.Trim(domain, "/")
	domain = strings.Trim(domain, ".")
	if strings.Contains(domain, "https://") || strings.Contains(domain, "http://") {
		uri, _ := url.Parse(domain)
		domain = uri.Hostname()
	}
	domain = strings.ReplaceAll(domain, "//", "/")

	//path = "//api//v4//?ths=ss"

	// fixing the parameter
	path = strings.Trim(path, " ")
	path = strings.Trim(path, "/")
	path = strings.Trim(path, ".")

	u, _ := url.Parse(path)
	if queryIndex := u.RawQuery; queryIndex != "" {
		path = path[:len(path)-len(u.RawQuery)-1]
	}

	if strings.Contains(path, "https://") || strings.Contains(path, "http://") {
		uri, _ := url.Parse(path)
		path = uri.Path
	}
	path = strings.ReplaceAll(path, "//", "/")

	return domain, path
}

func isParams(req *fiber.Request, domain string, parameter string, params *Params) *fiber.Request {
	// TODO fix domain and parameter

	domain, parameter = fixURI(domain, parameter)

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

	withSlice := []string{}
	if params.With.Contacts {
		withSlice = append(withSlice, "contacts")
	}
	if params.With.Leads {
		withSlice = append(withSlice, "leads")
	}
	if params.With.Companies {
		withSlice = append(withSlice, "companies")
	}
	if params.With.CatalogElements {
		withSlice = append(withSlice, "catalog_elements")
	}
	if len(withSlice) > 0 {
		q.Set("with", strings.Join(withSlice, ","))
	}

	if params.Filter != nil {
        for key, value := range params.Filter {
            q.Set("filter["+key+"]", value)
        }
    }

	s := reflect.TypeOf(*params)
	v := reflect.ValueOf(*params)
	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		value := v.Field(i).String()

		if field.Name == "ContactID" {
			field.Name = "contact_id"
		}
		if field.Name == "ChatID" {
			field.Name = "chat_id"
		}
		if value != "" && field.Type != reflect.TypeOf(With{}) {
			q.Set(strings.ToLower(field.Name), value)
		}
	}
	u.RawQuery = q.Encode()
	req.SetRequestURI(u.String())
	return req
}

func (api *API) log(message ...interface{}) {
	if api.Debug {
		log.Println(message...)
	}
}

func isRegex(text string) bool {
	re := regexp.MustCompile("^[a-zA-z0-9]{1,}\\.(bitrix|amocrm)\\.(ru|com|kz|kg)")
	return re.MatchString(text)
}
