package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/internal/fourslash"
	"github.com/microsoft/typescript-go/internal/ls/lsutil"
	"github.com/microsoft/typescript-go/internal/testutil"
)

func TestAllowRenameOfImportPathEnabled(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /project/old-module.ts
export const value = 42;

// @Filename: /project/index.ts
import { value } from "./old-/**/module";`

	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToMarker(t, "")
	
	// When enabled, renaming import path should work
	f.VerifyRenameSucceeded(t, &lsutil.UserPreferences{
		AllowRenameOfImportPath: true,
	})
}

func TestAllowRenameOfImportPathDisabled(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /project/old-module.ts
export const value = 42;

// @Filename: /project/index.ts
import { value } from "./old-/**/module";`

	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToMarker(t, "")
	
	// When disabled, renaming import path should not work
	f.VerifyRenameFailed(t, &lsutil.UserPreferences{
		AllowRenameOfImportPath: false,
	})
}
