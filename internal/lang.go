package internal

import (
	"net/http"
)

var CurrentLang string = "en"

// Data structure to store translations
type PageData struct {
	Title           string
	Description     string
	Placeholder     string
	ProcessButton   string
	CopyButton      string
    CopySuccessText string
	ResultHeading   string
	ToggleButton    string
	Instructions    string
    Footer          string
}

func GetTranslations(lang string) PageData {
	if lang == "fr" {
		return PageData{
			Title:         "NBSP Formateur",
			Description:   "Assurez-vous que votre texte suit les règles typographiques françaises.",
			Placeholder:   "Entrez votre texte ici...",
			ProcessButton: "Traiter le texte",
			CopyButton:    "Copier le texte",
            CopySuccessText: "✔ Copié",
			ResultHeading: "Texte traité :",
			ToggleButton:  "Switch to English",
			Instructions:  `
                <p><strong>Cliquez</strong> sur les espaces en surbrillance pour basculer entre l'espace normal et l'espace insécable (nbsp).</p>
                <p>Les espaces <strong class="highlight grey">Gris</strong> sont en dehors des règles typographiques françaises concernant les nbsp. Les <strong class="highlight green">Les espaces Verts</strong> sont concernés par ces règles et seront automatiquement convertis en nbsp.</p>
            `,
            Footer: "Créé pour un usage interne — <a href='github.com/craftidev/nbsp-highlight'>craftidev</a> © 2024",
		}
	}

	return PageData{
		Title:         "NBSP Formatter",
		Description:   "Ensure that your text follows the French typographic rules.",
		Placeholder:   "Enter your text here...",
		ProcessButton: "Process Text",
		CopyButton:    "Copy Text to Clipboard",
        CopySuccessText: "✔ Copied",
		ResultHeading: "Processed Text:",
		ToggleButton:  "Passer en Français",
		Instructions:  `
            <p><strong>Click</strong> on highlighted spaces to toggle between normal space and non-breaking space (nbsp).</p>
            <p><strong class="highlight grey">Grey</strong> spaces are outside the French typographic rules about nbsp. <strong class="highlight green">Green</strong> spaces are concerned by the rules and will be auto-converted to nbsp.</p>
        `,
        Footer: "Created for internal use — <a href='github.com/craftidev/nbsp-highlight'>craftidev</a> © 2024",
	}
}

func SwitchLanguageHandler(w http.ResponseWriter, r *http.Request) {
	lang := r.FormValue("lang")
	if lang == "fr" {
		CurrentLang = "fr"
	} else {
		CurrentLang = "en"
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
