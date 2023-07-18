package terminals

import (
	"github.com/srlehn/termimg/internal/environ"
	"github.com/srlehn/termimg/internal/propkeys"
	"github.com/srlehn/termimg/term"
)

////////////////////////////////////////////////////////////////////////////////
// MacTerm
////////////////////////////////////////////////////////////////////////////////

func init() {
	term.RegisterTermChecker(&termCheckerMacTerm{term.NewTermCheckerCore(termNameMacTerm)})
}

const termNameMacTerm = `macterm`

var _ term.TermChecker = (*termCheckerMacTerm)(nil)

type termCheckerMacTerm struct{ term.TermChecker }

func (t *termCheckerMacTerm) CheckExclude(ci environ.Proprietor) (mightBe bool, p environ.Proprietor) {
	p = environ.NewProprietor()
	if t == nil || ci == nil {
		p.SetProperty(propkeys.CheckTermEnvExclPrefix+termNameMacTerm, term.CheckTermFailed)
		return false, p
	}
	v, ok := ci.LookupEnv(`TERM_PROGRAM`)
	if ok && v == `MacTerm` {
		p.SetProperty(propkeys.CheckTermEnvExclPrefix+termNameMacTerm, term.CheckTermPassed)
		if ver, okV := ci.LookupEnv(`TERM_PROGRAM_VERSION`); okV && len(ver) > 0 {
			p.SetProperty(propkeys.MacTermBuildNr, ver) // YYYYMMDD
			return true, p
		}
		return true, p
	}
	p.SetProperty(propkeys.CheckTermEnvExclPrefix+termNameMacTerm, term.CheckTermFailed)
	return false, p
}
