package generatoruuid

import (
	"github.com/google/uuid"

	"go-cli/pkg/logger"
)

func GenerateUUIDV4() {
	logger.InfoF("UUID is : %s", uuid.New().String())
}
