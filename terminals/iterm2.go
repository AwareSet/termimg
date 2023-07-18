package terminals

import (
	"github.com/srlehn/termimg/internal/environ"
	"github.com/srlehn/termimg/internal/propkeys"
	"github.com/srlehn/termimg/term"
)

////////////////////////////////////////////////////////////////////////////////
// iTerm2
////////////////////////////////////////////////////////////////////////////////

func init() {
	term.RegisterTermChecker(&termCheckerITerm2{term.NewTermCheckerCore(termNameITerm2)})
}

const termNameITerm2 = `iterm2`

var _ term.TermChecker = (*termCheckerITerm2)(nil)

type termCheckerITerm2 struct{ term.TermChecker }

func (t *termCheckerITerm2) CheckExclude(ci environ.Proprietor) (mightBe bool, p environ.Proprietor) {
	p = environ.NewProprietor()
	if t == nil || ci == nil {
		p.SetProperty(propkeys.CheckTermEnvExclPrefix+termNameITerm2, term.CheckTermFailed)
		return false, p
	}
	v, ok := ci.LookupEnv(`TERM_PROGRAM`)
	if ok && v == `iTerm.app` {
		p.SetProperty(propkeys.CheckTermEnvExclPrefix+termNameITerm2, term.CheckTermPassed)
		if ver, okV := ci.LookupEnv(`TERM_PROGRAM_VERSION`); okV && len(ver) > 0 {
			p.SetProperty(propkeys.ITerm2Version, ver)
			return true, p
		}
		return true, p
	}
	p.SetProperty(propkeys.CheckTermEnvExclPrefix+termNameITerm2, term.CheckTermFailed)
	return false, p
}

/*
https://github.com/kmgrant/macterm/issues/3#issuecomment-458387953
printf '\033[>q'
ESC P>|iTerm2 3.3.20200425-nightly ESC \
*/
