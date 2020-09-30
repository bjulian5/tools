package a

import (
	_ "github.com/bjulian5/tools/internal/lsp/circular/triple/b" //@diag("_ \"github.com/bjulian5/tools/internal/lsp/circular/triple/b\"", "compiler", "import cycle not allowed", "error")
)
