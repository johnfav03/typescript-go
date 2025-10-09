package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/internal/core"
	"github.com/microsoft/typescript-go/internal/fourslash"
	. "github.com/microsoft/typescript-go/internal/fourslash/tests/util"
	"github.com/microsoft/typescript-go/internal/ls"
	"github.com/microsoft/typescript-go/internal/testutil"
)

func TestAutoImportFileExcludePatterns1(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /project/a.ts
export const varA = 10;

// @Filename: /project/excluded/b.ts
export const varB = 20;

// @Filename: /project/c.ts
varA/**/
`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	
	// With exclusion pattern - only varA from ./a should be available as auto-import
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		UserPreferences: &ls.UserPreferences{
			IncludeCompletionsForModuleExports:    core.TSTrue,
			IncludeCompletionsForImportStatements: core.TSTrue,
			AutoImportFileExcludePatterns:         []string{"/project/excluded/**/*"},
		},
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{"varA"},
			Excludes: []string{"varB"},
		},
	})
}

func TestAutoImportFileExcludePatterns2(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /project/src/a.ts
export const varA = 10;

// @Filename: /project/tests/b.ts
export const varB = 20;

// @Filename: /project/node_modules/c/index.ts
export const varC = 30;

// @Filename: /project/src/main.ts
varA/**/
`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	
	// Exclude tests and node_modules directories
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		UserPreferences: &ls.UserPreferences{
			IncludeCompletionsForModuleExports:    core.TSTrue,
			IncludeCompletionsForImportStatements: core.TSTrue,
			AutoImportFileExcludePatterns:         []string{"/project/tests/**/*", "/project/node_modules/**/*"},
		},
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{"varA"},
			Excludes: []string{"varB", "varC"},
		},
	})
}
