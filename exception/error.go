package exception

import (
	"fmt"
	"net/http"
	"strings"
)

type ErrorCode string

var CustomErrors = map[ErrorCode]map[string]string{
	"auth-001": {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "invalid user id or secret"},
	"auth-002": {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "user does not exist"},
	"auth-003": {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "unexpected signing method"},
	"auth-004": {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "unable to parse token"},
	"auth-005": {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "unable to retrieve signed token"},
	"auth-006": {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "invalid token"},
	"auth-007": {"status": fmt.Sprintf("%d", http.StatusUnauthorized), "message": "unauthorized"},
	"auth-008": {"status": fmt.Sprintf("%d", http.StatusUnauthorized), "message": "refresh token count exceeded max limit"},
	"auth-009": {"status": fmt.Sprintf("%d", http.StatusUnauthorized), "message": "user not authorized to perform this action"},
	//"auth-010":  {"status": fmt.Sprintf("%d", http.StatusUnauthorized), "message": "payer agent not authorized to perform this action"},
	"query-001": {"status": fmt.Sprintf("%d", http.StatusInternalServerError), "message": "unable to create new record"},
	"query-002": {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "record already exists"},
	"query-003": {"status": fmt.Sprintf("%d", http.StatusNotFound), "message": "record does not exist"},
	"query-004": {"status": fmt.Sprintf("%d", http.StatusInternalServerError), "message": "unexpected query error"},
	"query-005": {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "user with email already exists"},
	//"query-006": {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "agent with fic already exists"},
	"mail-001": {"status": fmt.Sprintf("%d", http.StatusInternalServerError), "message": "could not send email"},
	"mail-002": {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "invalid user email"},
	"mail-003": {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "invalid college email"},
	"req-001":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "could not parse request query"},
	"req-002":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "improper input received in form fields"},
	"req-003":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "time not in valid format"},
	"req-004":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "agent not registered"},
	"req-005":  {"status": fmt.Sprintf("%d", http.StatusNotFound), "message": "route does not exist"},
	"req-006":  {"status": fmt.Sprintf("%d", http.StatusNotFound), "message": "agent with the id does not exist"},
	"req-007":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "request message id cannot be empty"},
	"req-008":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "request with the id does not exist"},
	"req-009":  {"status": fmt.Sprintf("%d", http.StatusUnauthorized), "message": "request with the id is not authorized for the payer agent"},
	"req-010":  {"status": fmt.Sprintf("%d", http.StatusNotFound), "message": "message with request id does not exist"},
	"req-011":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "r2px request with the id has already been rejected"},
	"req-012":  {"status": fmt.Sprintf("%d", http.StatusNotFound), "message": "this link has expired"},
	"req-013":  {"status": fmt.Sprintf("%d", http.StatusNotFound), "message": "client with client id does not exist"},
	"req-014":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "wrong client secret"},
	"req-015":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "request already declined"},
	"req-016":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "request already accepted"},
	"req-017":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "request already paid"},
	"req-018":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "payer already enrolled by agent"},
	"req-019":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "payer not enrolled in r2px"},
	"req-020":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "payee not present in r2px"},
	"req-021":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "country not present in r2px"},
	"req-022":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "request expired"},
	"req-023":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "webhook of the type already exists for agent"},
	"req-024":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "payee_request_status and payee_transaction_status webhooks are required"},
	"req-025":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "payer_request_arrived and payee_transaction_status webhooks are required"},
	"req-026":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "both payee and payer webhook types need are required"},
	"req-027":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "webhook to send request id not configured by payer agent"},
	"req-028":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "amount should be less than 10000 USD"},
	"req-029":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "amount should be greater than 10 USD"},
	"req-030":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "amount should be greater than 10 USD"},
	"req-031":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "amount should be a valid number"},
	"req-032":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "payee should be over 18 to send request"},
	"req-033":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "request with the end to end id already exists"},
	"req-034":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "request already accepted/rejected"},
	"sys-001":  {"status": fmt.Sprintf("%d", http.StatusInternalServerError), "message": "couldn't load env file"},
	"sys-002":  {"status": fmt.Sprintf("%d", http.StatusBadRequest), "message": "couldn't parse query"},
	"sys-003":  {"status": fmt.Sprintf("%d", http.StatusInternalServerError), "message": "couldn't parse JWT refresh time from env file"},
	"sys-004":  {"status": fmt.Sprintf("%d", http.StatusInternalServerError), "message": "couldn't parse JWT access time from env file"},
	"sys-005":  {"status": fmt.Sprintf("%d", http.StatusInternalServerError), "message": "couldn't parse JWT refresh count from env file"},
	"sys-006":  {"status": fmt.Sprintf("%d", http.StatusInternalServerError), "message": "error while encrypting client secret"},
	"sys-007":  {"status": fmt.Sprintf("%d", http.StatusInternalServerError), "message": "error while decrypting client secret"},
	"sys-008":  {"status": fmt.Sprintf("%d", http.StatusInternalServerError), "message": "error sending webhook"},
	"sys-009":  {"status": fmt.Sprintf("%d", http.StatusInternalServerError), "message": "email template id not configured in server"},
}

type UniqueConstraintError struct {
	Text string
}

func (u UniqueConstraintError) Error() string {
	return u.Text
}

type InvalidEmailError struct{}

func (InvalidEmailError) Error() string {
	return "invalid email"
}

type UnverifiedUser struct{}

func (UnverifiedUser) Error() string {
	return "user is not verified"
}

type InvalidPassword struct{}

func (InvalidPassword) Error() string {
	return "password os not valid"
}

type EmailAlreadyExists struct{}

func (EmailAlreadyExists) Error() string {
	return "email already exists"
}

func CheckUniqueConstraint(err error) bool {
	return strings.Contains(err.Error(), "SQLSTATE 23505")
}

func CheckForeignKeyConstraint(err error) bool {
	return strings.Contains(err.Error(), "SQLSTATE 23503")
}
