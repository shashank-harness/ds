package dsapi

import (
	"bytes"
	"context"
	"database/sql"
	"ds/gen/ds"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"golang.org/x/exp/slog"
	"log"
	"net/url"
	"os/exec"
	"strings"
	"time"
)

// ds service example implementation.
// The example methods log the requests and return zero values.
type dssrvc struct {
	logger   *log.Logger
	ngClient *resty.Client
	cgClient *resty.Client
	db       *sqlx.DB
}

type AccountDTO struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

type User struct {
	MetaData struct {
	} `json:"metaData"`
	Resource struct {
		UUID             string `json:"uuid"`
		Name             string `json:"name"`
		Email            string `json:"email"`
		Token            string `json:"token"`
		DefaultAccountID string `json:"defaultAccountId"`
		Intent           any    `json:"intent"`
		Accounts         []struct {
			UUID              string `json:"uuid"`
			AccountName       string `json:"accountName"`
			CompanyName       string `json:"companyName"`
			DefaultExperience string `json:"defaultExperience"`
			CreatedFromNG     bool   `json:"createdFromNG"`
			NextGenEnabled    bool   `json:"nextGenEnabled"`
		} `json:"accounts"`
		Admin                          bool `json:"admin"`
		TwoFactorAuthenticationEnabled bool `json:"twoFactorAuthenticationEnabled"`
		EmailVerified                  bool `json:"emailVerified"`
		Locked                         bool `json:"locked"`
		Disabled                       bool `json:"disabled"`
		SignupAction                   any  `json:"signupAction"`
		Edition                        any  `json:"edition"`
		BillingFrequency               any  `json:"billingFrequency"`
		UtmInfo                        struct {
			UtmSource   any `json:"utmSource"`
			UtmContent  any `json:"utmContent"`
			UtmMedium   any `json:"utmMedium"`
			UtmTerm     any `json:"utmTerm"`
			UtmCampaign any `json:"utmCampaign"`
		} `json:"utmInfo"`
		ExternallyManaged bool  `json:"externallyManaged"`
		GivenName         any   `json:"givenName"`
		FamilyName        any   `json:"familyName"`
		ExternalID        any   `json:"externalId"`
		CreatedAt         int64 `json:"createdAt"`
		LastUpdatedAt     int64 `json:"lastUpdatedAt"`
		UserPreference    any   `json:"userPreference"`
	} `json:"resource"`
	ResponseMessages []any `json:"responseMessages"`
}

type Email struct {
	Resource string `json:"resource"`
}

// var authToken = "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdXRoVG9rZW4iOiI2NGE1NjM3NmU1MzAwZDQwZmU0YzcyMTIiLCJpc3MiOiJIYXJuZXNzIEluYyIsImV4cCI6MTY4ODY1NjIxNSwiZW52IjoiZ2F0ZXdheSIsImlhdCI6MTY4ODU2OTc1NX0.jYqKNLp0lhFaG2AVQHq6hviZ9bgsK8ZBdMYYyi1zVlE"
var authToken = "Bearer eyJhbGciOiJIUzI1NiJ9.eyJwcmluY2lwYWwiOiJTRVJWSUNFIiwiSXNzdWVyIjoiSXNzdWVyIiwiVXNlcm5hbWUiOiJzaGFzaGFuay5zaW5naEBoYXJuZXNzLmlvIiwiaXNzIjoiSGFybmVzcyBJbmMiLCJleHAiOjE2OTEyNTQxNjUsImlhdCI6IjE2ODg1Njk3NTUifQ.bMen9QcejpSGVqS93kDh7sQTtmAN05TlFHwdvtu93ss"

const clusterUrl = "http://localdev.harness.io/prod1"

const ngMgrUrl1 = "/user/currentUser?routingId=%s&accountIdentifier=%s"
const verifyToken = "http://localhost:7457/signup/ds/verify/%s"
const createAccountAndUser = "http://localhost:7457/signup/ds/account/%s"

// NewDs returns the ds service implementation.
func NewDs(logger *log.Logger, ngManagerUrl, cgManagerUrl string) ds.Service {
	// Check this eventually
	ngClient := resty.New()
	ngClient.SetBaseURL(ngManagerUrl)
	ngClient.SetHeader("Authorization", authToken)

	cgClient := resty.New()
	cgClient.SetBaseURL(cgManagerUrl)
	cgClient.SetHeader("Authorization", authToken)
	db := setUpDB()
	return &dssrvc{
		logger, ngClient, cgClient, db}
}

// List all stored bottles
func (s *dssrvc) List(ctx context.Context) (res ds.AccountMgmtCollection, err error) {
	s.logger.Print("ds.list")
	return
}

// Complete New Signup flow with token
func (s *dssrvc) Complete(ctx context.Context, p *ds.CompletePayload) (res *ds.UserResource, err error) {
	var userResource ds.UserResource
	res = &ds.UserResource{}

	//0Print the request from UI
	fmt.Printf("%s", *p.Token)

	//1- Verify Token from NG mgr
	resource := s.callGET(err, fmt.Sprintf(verifyToken, *p.Token))
	email := resource.Resource

	// 2- Create an Account object and Save it in the DB
	account, err := s.createAccount(email)
	fmt.Printf("account %s", account)

	// 3- Create Account and User in Cluster with same Account uuid
	var accountDTO = AccountDTO{account.UUID, account.Accountname}
	user := s.callPUT(err, fmt.Sprintf(createAccountAndUser, email), accountDTO)

	// 4- Prepare repsonse to be sent to caller - In this case, Gateway
	userResource = ds.UserResource{user.Resource.UUID, user.Resource.Email, user.Resource.Name, account.Clusterurl, account.Accountname}
	s.logger.Printf("%s", userResource)

	//defer s.db.Close()
	return &userResource, err
}

func (s *dssrvc) createAccount(email string) (*ds.AccountMgmt, error) {
	output, err := exec.Command("uuidgen").Output()
	s.logger.Printf("output: ", output)
	uuid := string(output)
	uuid = uuid[:22]
	var id int
	accountName := strings.Split(email, "@")[0]
	result := s.db.QueryRow("INSERT INTO ACCOUNT (UUID,CLUSTERURL,ACCOUNTNAME) VALUES ($1, $2, $3)", uuid, clusterUrl, accountName).Scan(&id)

	s.logger.Printf("result: ", result)
	s.logger.Printf("id: ", id)

	var account = ds.AccountMgmt{int(id), uuid, clusterUrl, accountName}
	return &account, err
}

func (s *dssrvc) callGET(err error, url string) *Email {
	var email = &Email{}
	resNg, err := s.ngClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", authToken).
		SetResult(&email).
		Get(url)

	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, resNg.Body(), "", "\t")
	log.Println("Response from NG:", string(prettyJSON.Bytes()))
	return email
}

func (s *dssrvc) callPUT(err error, url string, body AccountDTO) *User {
	var user = &User{}
	resNg, err := s.ngClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", authToken).
		SetBody(body).
		SetResult(&user).
		Put(url)

	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, resNg.Body(), "", "\t")
	log.Println("Response:", string(prettyJSON.Bytes()))
	return user
}

// Demo implements demo.
func (s *dssrvc) Demo(ctx context.Context, p *ds.DemoPayload) (res int, err error) {
	s.logger.Print("ds.demo")
	return
}

func setUpDB() *sqlx.DB {
	var db *sqlx.DB
	databaseUrl := "postgres://shashanksingh@localhost:5432/demo"
	//dbPass := "postgres"
	parsedURL, err := url.Parse(databaseUrl)
	if err != nil {
		panic(err)
	}
	queryParams := parsedURL.Query()
	queryParams.Add("sslmode", "disable")
	parsedURL.RawQuery = queryParams.Encode()
	connConfig, err := pgx.ParseConfig(parsedURL.String())
	if err != nil {
		panic(fmt.Errorf("error parsing db url '%s': %s", databaseUrl, err))
	}

	//connConfig.Password = dbPass
	connConfig.ConnectTimeout = time.Duration(5 * int(time.Second))
	connStr := stdlib.RegisterConnConfig(connConfig)

	sdb, err := sql.Open("pgx", connStr)

	db = sqlx.NewDb(sdb, "pgx")
	// TODO: make these into options
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	for i := 0; i < 30; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		slog.Warn("error checking database connection - retrying in 1s ...", "err", err)
		time.Sleep(time.Second)
	}
	return db
}
