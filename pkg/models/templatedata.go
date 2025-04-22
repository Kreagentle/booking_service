package models

type TmpltData struct {
	MpString    map[string]string
	MpInt       map[string]int
	MpFloat     map[string]float64
	MpInterface map[string]interface{}
	CSRFToken   string
	Flash       string
	Warning     string
	Error       string
}
