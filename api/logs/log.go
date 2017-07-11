package logs

import (
	"fmt"
	"net/url"

	"github.com/Sirupsen/logrus"
	"gitlab-odx.oracle.com/odx/functions/api/datastore/sql"
	"gitlab-odx.oracle.com/odx/functions/api/models"
)

func New(dbURL string) (models.FnLog, error) {
	u, err := url.Parse(dbURL)
	if err != nil {
		logrus.WithError(err).WithFields(logrus.Fields{"url": dbURL}).Fatal("bad DB URL")
	}
	logrus.WithFields(logrus.Fields{"db": u.Scheme}).Debug("creating log store")
	switch u.Scheme {
	case "sqlite3", "postgres", "mysql":
		return sql.New(u)
	default:
		return nil, fmt.Errorf("db type not supported %v", u.Scheme)
	}
}