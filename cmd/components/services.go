package components

import (
	"github.com/rwcosta/folderz/pkg/service"
	"github.com/rwcosta/folderz/services/structure"
)

var Services = []service.Service {
	structure.FolderStructureService{},
}
