package glob

import (
	"github.com/argoproj/argo-cd/v2/common"
	"github.com/gobwas/glob"
	log "github.com/sirupsen/logrus"
)

func Match(pattern, text string, separators ...rune) bool {
	compiledGlob, err := glob.Compile(pattern, separators...)
	if err != nil {
		log.WithFields(log.Fields{
			common.SecurityField:    common.SecurityLow,
			common.SecurityCWEField: 775,
		}).Warnf("failed to compile pattern %s due to error %v", pattern, err)
		return false
	}
	return compiledGlob.Match(text)
}
