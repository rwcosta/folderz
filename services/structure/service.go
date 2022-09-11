package structure

type FolderStructureService struct{}

func (f FolderStructureService) Execute() error {
    return mountFolderStructure()
}
