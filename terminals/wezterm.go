package terminals

import (
	"github.com/srlehn/termimg/internal/environ"
	"github.com/srlehn/termimg/internal/propkeys"
	"github.com/srlehn/termimg/term"
	"github.com/srlehn/termimg/wm"
)

////////////////////////////////////////////////////////////////////////////////
// WezTerm
////////////////////////////////////////////////////////////////////////////////

func init() {
	term.RegisterTermChecker(&termCheckerWezTerm{term.NewTermCheckerCore(termNameWezTerm)})
}

const termNameWezTerm = `wezterm`

var _ term.TermChecker = (*termCheckerWezTerm)(nil)

type termCheckerWezTerm struct{ term.TermChecker }

func (t *termCheckerWezTerm) CheckExclude(ci environ.Proprietor) (mightBe bool, p environ.Proprietor) {
	p = environ.NewProprietor()
	if t == nil || ci == nil {
		p.SetProperty(propkeys.CheckTermEnvExclPrefix+termNameWezTerm, term.CheckTermFailed)
		return false, p
	}

	var r bool
	pr := environ.NewProprietor()
	if v, ok := ci.LookupEnv(`TERM_PROGRAM`); ok && v == `WezTerm` {
		r = true
	}
	if v, err := ci.LookupEnv(`WEZTERM_EXECUTABLE`); err && len(v) > 0 {
		r = true
		pr.SetProperty(propkeys.WezTermExe, v)
	}
	if v, ok := ci.LookupEnv(`WEZTERM_UNIX_SOCKET`); ok && len(v) > 0 {
		r = true
		pr.SetProperty(propkeys.WezTermUnixSocket, v)
	}
	if v, ok := ci.LookupEnv(`WEZTERM_PANE`); ok && len(v) > 0 {
		r = true
		pr.SetProperty(propkeys.WezTermPane, v)
	}
	if v, ok := ci.LookupEnv(`WEZTERM_EXECUTABLE_DIR`); ok && len(v) > 0 {
		r = true
		pr.SetProperty(propkeys.WezTermExeDir, v)
	}
	if !r {
		p.SetProperty(propkeys.CheckTermEnvExclPrefix+termNameWezTerm, term.CheckTermFailed)
		return false, p
	}
	p.SetProperty(propkeys.CheckTermEnvExclPrefix+termNameWezTerm, term.CheckTermPassed)
	return true, p
}
func (t *termCheckerWezTerm) CheckIsWindow(w wm.Window) (is bool, p environ.Proprietor) {
	p = environ.NewProprietor()
	if t == nil || w == nil {
		p.SetProperty(propkeys.CheckTermWindowIsPrefix+termNameWezTerm, term.CheckTermFailed)
		return false, p
	}
	isWindow := w.WindowType() == `x11` &&
		w.WindowClass() == `org.wezfurlong.wezterm` &&
		w.WindowInstance() == `org.wezfurlong.wezterm`
	if isWindow {
		p.SetProperty(propkeys.CheckTermWindowIsPrefix+termNameWezTerm, term.CheckTermPassed)
	} else {
		p.SetProperty(propkeys.CheckTermWindowIsPrefix+termNameWezTerm, term.CheckTermFailed)
	}
	return isWindow, p
}
func (t *termCheckerWezTerm) Args(ci environ.Proprietor) []string { return []string{`--skip-config`} }
