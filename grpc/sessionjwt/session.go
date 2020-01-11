package sessionjwt

import (
	"time"

	"emperror.dev/errors"
	"github.com/endiangroup/compandauth"
	"github.com/gbrlsnchs/jwt/v3"
	cmap "github.com/orcaman/concurrent-map"
)

var hs = jwt.NewHS256([]byte("secret"))

type UserSession struct {
	ID                string
	Role              string
	Secret            string
	MaxActiveSessions int64
	CAA               compandauth.CAA
}

type SessionStore struct {
	list cmap.ConcurrentMap
}

func NewSessionStore() *SessionStore {
	return &SessionStore{list: cmap.New()}
}
func (s *SessionStore) Insert(u *UserSession) {
	// u.Lock()
	s.list.Set(u.ID, u)
}
func (s *SessionStore) Query(id string) (*UserSession, bool) {
	v, ck := s.list.Get(id)
	if ck {
		return v.(*UserSession), ck
	}
	return nil, ck
}
func (s *SessionStore) Update(u *UserSession) error {
	cb := func(exists bool, valueInMap interface{}, newValue interface{}) interface{} {
		nv := newValue.(*UserSession)
		if !exists {
			return []*UserSession{nv}
		}
		res := valueInMap.(*UserSession)
		return res
	}
	s.list.Upsert(u.ID, u, cb)
	return nil
}

type JwtSessionCAA struct {
	jwt.Payload
	UserID string                 `json:"user_id"`
	Role   string                 `json:"role"`
	Secret string                 `json:"secret"`
	CAA    compandauth.SessionCAA `json:"caa"`
}

func setCounterCAA(i int64) *Counter {
	return NewCounter(i)
}

func NewSession(id, role, secret string, i int64) (*UserSession, error) {
	return &UserSession{
		ID:                id,
		Role:              role,
		Secret:            secret,
		MaxActiveSessions: i,
		CAA:               setCounterCAA(0),
	}, nil
}

func (u *UserSession) Lock() {
	u.CAA.Lock()
}

func (u *UserSession) LogoutAllSessions() {
	u.CAA.Revoke(u.MaxActiveSessions)
}

func (u *UserSession) Update() error {
	return nil
}

func (s *SessionStore) Login(id string) (*UserSession, []byte, error) {
	u, chk := s.Query(id)

	if chk {
		now := time.Now()
		j := JwtSessionCAA{
			Payload: jwt.Payload{
				Issuer:         "goUMS",
				Subject:        "goUMS",
				Audience:       jwt.Audience{"https://golang.org"},
				ExpirationTime: jwt.NumericDate(now.Add(24 * 30 * 12 * time.Hour)),
				NotBefore:      jwt.NumericDate(now.Add(30 * time.Minute)),
				IssuedAt:       jwt.NumericDate(now),
				JWTID:          "jwtID",
			},
			Role:   "access",
			Secret: "secret",
			UserID: u.ID,
		}
		j.CAA = u.CAA.Issue()
		// if err := s.Update(u); err != nil { // update user record with new issued CAA value
		// 	return nil, nil, err
		// }
		if !u.CAA.IsValid(j.CAA, u.MaxActiveSessions) {
			if u.CAA.IsLocked() {
				return nil, nil, errors.New("It appears your account has been locked")
			}
			return nil, nil, errors.New("Invalid session, please login again")
		}

		token, err := j.GenerateToken()
		return u, token, err
	}

	return nil, nil, errors.New("UserSession login failed")
}

func (s *SessionStore) Valid(j JwtSessionCAA) error {

	u, chk := s.Query(j.UserID)

	if chk {

		if !u.CAA.IsValid(j.CAA, u.MaxActiveSessions) {
			if u.CAA.IsLocked() {
				return errors.New("It appears your account has been locked")
			}
			return errors.New("Invalid session, please login again")
		}
	}
	return nil
}

func (j JwtSessionCAA) GenerateToken() ([]byte, error) {
	now := time.Now()
	j.Payload.Audience = []string{"https://golang.org"}
	j.Payload.ExpirationTime = jwt.NumericDate(now.Add(1 * time.Second))
	j.Payload.NotBefore = jwt.NumericDate(now.Add(1 * time.Second))
	j.Payload.IssuedAt = jwt.NumericDate(now)
	return jwt.Sign(j, hs)
}

func ParseVerifyToken(token []byte) (JwtSessionCAA, error) {
	var now = time.Now()
	var pl JwtSessionCAA

	// Validate claims "iat", "exp" and "aud".
	var iatValidator = jwt.IssuedAtValidator(now)
	var expValidator = jwt.ExpirationTimeValidator(now)
	var audValidator = jwt.AudienceValidator(jwt.Audience{"https://golang.org"})
	// Use jwt.ValidatePayload to build a jwt.VerifyOption.
	// Validators are run in the order informed.
	var validatePayload = jwt.ValidatePayload(&pl.Payload, iatValidator, expValidator, audValidator)

	_, err := jwt.Verify(token, hs, &pl, validatePayload)

	return pl, err
}
