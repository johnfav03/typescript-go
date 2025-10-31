package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/internal/fourslash"
	. "github.com/microsoft/typescript-go/internal/fourslash/tests/util"
	"github.com/microsoft/typescript-go/internal/ls/lsutil"
	"github.com/microsoft/typescript-go/internal/testutil"
)

func TestQuotePreferenceSingle(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /project/helper.ts
export const helperFunc = () => {};

// @Filename: /project/index.ts
helperFunc/**/`

	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToMarker(t, "")
	f.VerifyApplyCodeActionFromCompletion(t, PtrTo(""), &fourslash.ApplyCodeActionFromCompletionOptions{
		Name:   "helperFunc",
		Source: "./helper",
		UserPreferences: &lsutil.UserPreferences{
			QuotePreference: lsutil.QuotePreferenceSingle,
		},
		NewFileContent: PtrTo(`import { helperFunc } from './helper';

helperFunc`),
	})
}

func TestQuotePreferenceDouble(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /project/helper.ts
export const helperFunc = () => {};

// @Filename: /project/index.ts
helperFunc/**/`

	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToMarker(t, "")
	f.VerifyApplyCodeActionFromCompletion(t, PtrTo(""), &fourslash.ApplyCodeActionFromCompletionOptions{
		Name:   "helperFunc",
		Source: "./helper",
		UserPreferences: &lsutil.UserPreferences{
			QuotePreference: lsutil.QuotePreferenceDouble,
		},
		NewFileContent: PtrTo(`import { helperFunc } from "./helper";

helperFunc`),
	})
}

func TestQuotePreferenceAuto(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /project/helper.ts
export const helperFunc = () => {};

// @Filename: /project/index.ts
import {} from 'existing';

helperFunc/**/`

	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToMarker(t, "")
	
	// Auto should detect single quotes from existing imports
	f.VerifyApplyCodeActionFromCompletion(t, PtrTo(""), &fourslash.ApplyCodeActionFromCompletionOptions{
		Name:   "helperFunc",
		Source: "./helper",
		UserPreferences: &lsutil.UserPreferences{
			QuotePreference: lsutil.QuotePreferenceAuto,
		},
		NewFileContent: PtrTo(`import {} from 'existing';
import { helperFunc } from './helper';

helperFunc`),
	})
}
