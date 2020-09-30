package c

import "github.com/bjulian5/tools/internal/lsp/rename/b"

func _() {
	b.Hello() //@rename("Hello", "Goodbye")
}
