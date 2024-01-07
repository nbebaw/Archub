package lib

const (
	ColorRed  = "\033[0;31m"
	ColorNone = "\033[0m"
)

var (
	Help            bool
	ShowVersion     bool
	CleanUp         bool
	CleanUpAlias    bool
	ListPkg         bool
	ListPkgAlias    bool
	CheckUpdatePkg  bool
	UpdateAllPkg    bool
	UpdatePkg       bool
	UpdatePkgAlias  bool
	DeletePkg       bool
	DeletePkgAlias  bool
	SearchPkg       bool
	SearchPkgAlias  bool
	InstallPkg      bool
	InstallPkgAlias bool
)
