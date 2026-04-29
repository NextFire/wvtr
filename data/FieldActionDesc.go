package data

import (
	"fmt"
)

func (fad *FieldActionDesc) String() string {
	res := ""
	from := fad.FromH
	target := fad.TargetH
	status := fad.TargetStatus

	critTxt := ""
	if status&Crit == Crit {
		critTxt = "(crit)"
	}

	targetPVChange := fad.TargetPVChange
	if status&TookDamage == TookDamage {
		res += fmt.Sprintf("%s has inflicted %f dmg%s to %s.", from.Name, targetPVChange, critTxt, target.Name)
	}
	if status&Died == Died {
		res += fmt.Sprintf("%s died.", target.Name)
	}
	if status&Dodged == Dodged {
		res += fmt.Sprintf("%s dodged.", target.Name)
	}
	if status&Blocked == Blocked {
		res += fmt.Sprintf("%s blocked.", target.Name)
	}
	return res
}
