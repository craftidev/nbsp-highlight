package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
    "html"
)

func serveFrontend(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func handleText(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	text := html.EscapeString(r.FormValue("inputText"))
	highlightedText := highlightText(text)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprintf(w, "%s", highlightedText)
}

func toggleSpaceHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	currentClass := r.FormValue("class")

	var newSpace string
	if text == " " {
		newSpace = fmt.Sprintf(
            `<span
                hx-post="/toggle"
                hx-target="this"
                hx-swap="outerHTML"
                hx-vals='{"text":"&amp;nbsp;", "class":"%s"}'
                class="highlight %s"
            >&amp;nbsp;</span>`, currentClass, currentClass)
	} else {
		newSpace = fmt.Sprintf(
            `<span
                hx-post="/toggle"
                hx-target="this"
                hx-swap="outerHTML"
                hx-vals='{"text":" ", "class":"%s"}'
                class="highlight %s"
            > </span>`, currentClass, currentClass)
	}

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprintf(w, "%s", newSpace)
}


func highlightText(text string) string {
    nbspRegex := regexp.MustCompile(`&nbsp;`)
    ampNbspRegex := regexp.MustCompile(`&amp;nbsp;`)
    greenBefore := regexp.MustCompile(`[«\d]`)
    greenAfter := regexp.MustCompile(`[»—–!?%°:\d]`)
    var highlightedText strings.Builder

    runes := []rune(text)
    lastPos := 0
    for i := 0; i < len(runes); i++ {
        r := runes[i]

        isNbsp := false
        isAmpNbsp := false
        if i + 5 < len(runes) {
            isNbsp = nbspRegex.MatchString(string(runes[i:i + 6]))
        }
        if i + 9 < len(runes) {
            isAmpNbsp = ampNbspRegex.MatchString(string(runes[i:i + 10]))
        }
        if r == ' ' || r == '\u00A0' || isNbsp || isAmpNbsp {
            highlightedText.WriteString(string(runes[lastPos:i]))
            if isNbsp {
                i += 5
            }
            if isAmpNbsp {
                i += 9
            }

            var before, after rune
            if i > 0  {
                before = runes[i - 1]
            }
            if i + 1 < len(runes) {
                after = runes[i + 1]
            }

            if greenBefore.MatchString(string(before)) || greenAfter.MatchString(string(after)) {
                highlightedText.WriteString(
                    `<span
                        hx-post="/toggle"
                        hx-target="this"
                        hx-swap="outerHTML"
                        hx-vals='{"text":"&amp;nbsp;", "class":"highlight green"}'
                        class="highlight green"
                    >&amp;nbsp;</span>`)
            } else {
                if isNbsp {
                    highlightedText.WriteString(
                        `<span
                            hx-post="/toggle"
                            hx-target="this"
                            hx-swap="outerHTML"
                            hx-vals='{"text":"&amp;nbsp;", "class":"highlight grey"}'
                            class="highlight grey"
                        >&amp;nbsp;</span>`)
                } else {
                    highlightedText.WriteString(
                        `<span
                            hx-post="/toggle"
                            hx-target="this"
                            hx-swap="outerHTML"
                            hx-vals='{"text":" ", "class":"highlight grey"}'
                            class="highlight grey"
                        > </span>`)
                }
            }
            lastPos = i + 1
        }
    }
    highlightedText.WriteString(string(runes[lastPos:]))
    return highlightedText.String()
}

func main() {
	http.HandleFunc("/", serveFrontend)
	http.HandleFunc("/process", handleText)
	http.HandleFunc("/toggle", toggleSpaceHandler)

    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server started at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
