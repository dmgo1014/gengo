// b only public and private packages. The latter it shouldn't.
package b

import (
	_ "github.com/dgodyna/gengo/examples/import-boss/tests/rules/c"
)
