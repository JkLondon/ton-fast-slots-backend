package errorlist

import "errors"

var (
	ErrWrongAuthData            = errors.New("wrong auth data")
	ErrWrongAuthCredentials     = errors.New("wrong auth credentials")
	ErrUnknownSessionType       = errors.New("unknown session type")
	ErrInternalServerError      = errors.New("internal server error")
	ErrNotFoundAuthData         = errors.New("not found auth data")
	ErrWrongPasswordUpdateData  = errors.New("wrong password update data")
	ErrServiceInCart            = errors.New("service in cart")
	ErrNoUsernameInTG           = errors.New("no username in TG")
	ErrThisUsernameAlreadyTaken = errors.New("this username already taken")
	ErrNotFoundVerifyCode       = errors.New("not found verify code")
	ErrWrongCurrentPassword     = errors.New("wrong current password")
	ErrScheduleDateRangeBusy    = errors.New("schedule date range busy")
)
