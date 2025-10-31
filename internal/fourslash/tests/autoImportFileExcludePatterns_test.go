package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/internal/core"
	"github.com/microsoft/typescript-go/internal/fourslash"
	. "github.com/microsoft/typescript-go/internal/fourslash/tests/util"
	"github.com/microsoft/typescript-go/internal/ls/lsutil"
	"github.com/microsoft/typescript-go/internal/testutil"
)

func TestAutoImportFileExcludePatternsBasic(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @module: commonjs
// @Filename: /project/node_modules/excluded-package/index.d.ts
export declare class ExcludedClass {}

// @Filename: /project/src/helper.ts
export const helperFunc = () => {};

// @Filename: /project/index.ts
Excluded/**/`

	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	
	// Without exclusion pattern, ExcludedClass should be available
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		UserPreferences: &lsutil.UserPreferences{
			IncludeCompletionsForModuleExports:    core.TSTrue,
			IncludeCompletionsForImportStatements: core.TSTrue,
		},
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{"ExcludedClass"},
		},
	})
	
	// With exclusion pattern, ExcludedClass should not be available
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		UserPreferences: &lsutil.UserPreferences{
			IncludeCompletionsForModuleExports:    core.TSTrue,
			IncludeCompletionsForImportStatements: core.TSTrue,
			AutoImportFileExcludePatterns:         []string{"/**/node_modules/excluded-package"},
		},
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Excludes: []string{"ExcludedClass"},
		},
	})
}

func TestAutoImportFileExcludePatternsMultiple(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @module: commonjs
// @Filename: /project/node_modules/package-a/index.d.ts
export declare class ClassA {}

// @Filename: /project/node_modules/package-b/index.d.ts
export declare class ClassB {}

// @Filename: /project/src/helper.ts
export const helperFunc = () => {};

// @Filename: /project/index.ts
Class/**/`

	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	
	// Exclude multiple packages
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		UserPreferences: &lsutil.UserPreferences{
			IncludeCompletionsForModuleExports:    core.TSTrue,
			IncludeCompletionsForImportStatements: core.TSTrue,
			AutoImportFileExcludePatterns:         []string{"/**/node_modules/package-a", "/**/node_modules/package-b"},
		},
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Excludes: []string{"ClassA", "ClassB"},
		},
	})
}
