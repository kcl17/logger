package logger

import (
	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	log     *logrus.Logger
	initlog *logrus.Entry
	applog  *logrus.Entry
	pfcplog *logrus.Entry
	xdplog  *logrus.Entry
	ebpflog *logrus.Entry
)

func init() {
	log = logrus.New()
	log.SetReportCaller(false)

	log.Formatter = &formatter.Formatter{
		TimestampFormat: time.RFC3339,
		TrimMessages:    true,
		NoFieldsSpace:   true,
		HideKeys:        true,
	}

	applog  = log.WithFields(logrus.Fields{"component": "HEXA-UPF", "category": "App"})
	initlog = log.WithFields(logrus.Fields{"component": "HEXA-UPF", "category": "Init"})
	pfcplog = log.WithFields(logrus.Fields{"component": "HEXA-UPF", "category": "PFCP"})
	xdplog  = log.WithFields(logrus.Fields{"component": "HEXA-UPF", "category": "XDP"})
	ebpflog = log.WithFields(logrus.Fields{"component": "HEXA-UPF", "category": "EBPF"})
}

func SetLogLevel(level logrus.Level) {
	log.SetLevel(level)
}
