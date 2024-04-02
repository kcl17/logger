package logger

import (
	"time"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var (
	log         *logrus.Logger
	AppLog      *logrus.Entry
	InitLog     *logrus.Entry
	ConfigLog   *logrus.Entry
)

const (
	// FieldRanAddr     string = "ran_addr"
	// FieldRanId       string = "ran_id"
	// FieldAmfUeNgapID string = "amf_ue_ngap_id"
	// FieldSupi        string = "supi"
	// FieldSuci        string = "suci"
)

func init() {
	log = logrus.New()
	log.SetReportCaller(false)

	log.Formatter = &formatter.Formatter{
		TimestampFormat: time.RFC3339,
		TrimMessages:    true,
		NoFieldsSpace:   true,
		HideKeys:        true,
		// FieldsOrder:     []string{"component", "category", FieldRanAddr, FieldRanId, FieldAmfUeNgapID, FieldSupi, FieldSuci},
	}

	AppLog = log.WithFields(logrus.Fields{"component": "HEXA_UPF", "category": "App"})
	InitLog = log.WithFields(logrus.Fields{"component": "HEXA_UPF", "category": "Init"})
	ConfigLog = log.WithFields(logrus.Fields{"component": "HEXA_UPF", "category": "CFG"})
}

func SetLogLevel(level logrus.Level) {
	log.SetLevel(level)
}

// func SetReportCaller(set bool) {
// 	log.SetReportCaller(set)
// }
