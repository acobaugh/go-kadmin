package kadmin

// #cgo LDFLAGS: -lkrb5 -lkadm5clnt
// #include <krb5.h>
// #include <kadm5/admin.h>
// #include <kadm5/kadm_err.h>
import "C"

import (
	"github.com/acobaugh/krb5-go"
)

// An Error is a krb5 library error. It may internally have an
// associated context.
type Error struct {
	context *krb5.Context
	code    int32
}

// ErrorCode returns the C error code for this library.
func (err *Error) ErrorCode() int32 {
	return err.code
}

// Error returns the KADM5 error string
func (err *Error) Error() string {
	/*var ctx C.krb5_context
	if err.context != nil {
		ctx = err.context.ctx
	}*/
	return ""
}

type Handle struct {
	context *krb5.Context
	handle  **C.void
}

// PrincipalEntry maps to kadm5_principal_ent_t
type PrincipalEntry struct {
	Principal        krb5.Principal
	ExpireTimeRaw    int32
	LastPwdChangeRaw int32
	PwExpirationRaw  int32
	MaxLifeRaw       int32
	ModName          krb5.Principal
	ModDate          int32
	Attributes       int32
	Kvno             uint
	Mkvno            uint
	Policy           string
	AuxAttributes    int32

	// version 2 fields
	MaxRenewableLifeRaw int32
	LastSuccessRaw      int32
	LastFailed          int32
	FailAuthCount       uint
}

func (p *PrincipalEntry) toC() C.kadm5_principal_ent_t {
	return nil
}

func principalEntryFromC(p *C.kadm5_principal_ent_t) *PrincipalEntry {
	return &PrincipalEntry{}
}

func (h *Handle) GetPrincipal(principal *krb5.Principal) (PrincipalEntry, error) {

	return PrincipalEntry{}, nil
}
