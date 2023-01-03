package loggerfactory

import (
	"github.com/pkg/errors"
	"github.com/rahul-024/fund-transfer-poc/config"
	"github.com/rahul-024/fund-transfer-poc/loggerfactory/logrus"
)

// receiver for logrus factory
type LogrusFactory struct{}

// build logrus logger
func (mf *LogrusFactory) Build(lc *config.LogConfig) error {
	err := logrus.RegisterLog(*lc)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}
