package rlib

// AppendComment adds the supplied string to a.Comment, separating it from
// any existing comment already in the field.
func (a *Assessment) AppendComment(s string) {
	if len(a.Comment) > 0 {
		a.Comment += " | "
	}
	a.Comment += s
}
