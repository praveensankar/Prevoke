package vc

import (
	"github.com/praveensankar/Revocation-Service/config"
	"go.uber.org/zap"
)

func TestVC(conf config.Config) {

	zap.S().Infoln("VC TEST")
	TestDiploma()

}
