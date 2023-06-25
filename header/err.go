package header

import "errors"

var (
	ErrInvalidPCSField         = errors.New("header: device other than DeviceLink should have PCS encoding of either PCSXYZ or PCSLAB")
	ErrInvalidProfileSignature = errors.New("header: proile file signature should be 'acsp'")
)
