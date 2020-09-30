package b

import (
	_ "github.com/bjulian5/tools/internal/lsp/circular/double/one" //@diag("_ \"github.com/bjulian5/tools/internal/lsp/circular/double/one\"", "compiler", "import cycle not allowed", "error"),diag("\"github.com/bjulian5/tools/internal/lsp/circular/double/one\"", "compiler", "could not import github.com/bjulian5/tools/internal/lsp/circular/double/one (no package for import github.com/bjulian5/tools/internal/lsp/circular/double/one)", "error")
)
