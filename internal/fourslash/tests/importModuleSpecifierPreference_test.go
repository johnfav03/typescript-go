package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/internal/core"
	"github.com/microsoft/typescript-go/internal/fourslash"
	. "github.com/microsoft/typescript-go/internal/fourslash/tests/util"
	"github.com/microsoft/typescript-go/internal/ls/lsutil"
	"github.com/microsoft/typescript-go/internal/modulespecifiers"
	"github.com/microsoft/typescript-go/internal/testutil"
)

func TestImportModuleSpecifierPreferenceShortest(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /project/src/utils/helper.ts
export const helperFunc = () => {};

// @Filename: /project/src/index.ts
helper/**/`

	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		UserPreferences: &lsutil.UserPreferences{
			IncludeCompletionsForModuleExports:    core.TSTrue,
			IncludeCompletionsForImportStatements: core.TSTrue,
			ImportModuleSpecifierPreference:       modulespecifiers.ImportModuleSpecifierPreferenceShortest,
		},
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{"helperFunc"},
		},
	})
	// Verify using baseline to capture the actual source
	f.BaselineAutoImportsCompletions(t, []string{""})
}

func TestImportModuleSpecifierPreferenceRelative(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /project/src/utils/helper.ts
export const helperFunc = () => {};

// @Filename: /project/src/index.ts
helper/**/`

	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		UserPreferences: &lsutil.UserPreferences{
			IncludeCompletionsForModuleExports:    core.TSTrue,
			IncludeCompletionsForImportStatements: core.TSTrue,
			ImportModuleSpecifierPreference:       modulespecifiers.ImportModuleSpecifierPreferenceRelative,
		},
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{"helperFunc"},
		},
	})
	// Verify using baseline to capture the actual source
	f.BaselineAutoImportsCompletions(t, []string{""})
}

func TestImportModuleSpecifierPreferenceNonRelative(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /project/tsconfig.json
{ "compilerOptions": { "baseUrl": "." } }

// @Filename: /project/src/utils/helper.ts
export const helperFunc = () => {};

// @Filename: /project/src/index.ts
helper/**/`

	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		UserPreferences: &lsutil.UserPreferences{
			IncludeCompletionsForModuleExports:    core.TSTrue,
			IncludeCompletionsForImportStatements: core.TSTrue,
			ImportModuleSpecifierPreference:       modulespecifiers.ImportModuleSpecifierPreferenceNonRelative,
		},
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{"helperFunc"},
		},
	})
	// Verify using baseline to capture the actual source
	f.BaselineAutoImportsCompletions(t, []string{""})
}

func TestImportModuleSpecifierPreferenceProjectRelative(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /project/tsconfig.json
{ "compilerOptions": { "baseUrl": ".", "paths": { "~/utils/*": ["./src/utils/*"] } } }

// @Filename: /project/src/utils/helper.ts
export const helperFunc = () => {};

// @Filename: /project/src/index.ts
helper/**/`

	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		UserPreferences: &lsutil.UserPreferences{
			IncludeCompletionsForModuleExports:    core.TSTrue,
			IncludeCompletionsForImportStatements: core.TSTrue,
			ImportModuleSpecifierPreference:       modulespecifiers.ImportModuleSpecifierPreferenceProjectRelative,
		},
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{"helperFunc"},
		},
	})
	// Verify using baseline to capture the actual source
	f.BaselineAutoImportsCompletions(t, []string{""})
}
