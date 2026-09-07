package gen

import (
	"strconv"
	"strings"
)

// SVG renders one deterministic product image. The demo needs pictures and the
// demo has no network, so every image url in the catalogue points back at this
// gateway and this function answers it: a coloured card carrying the offer id,
// the image index and the archetype, with a shape that varies by hash.
//
// Same offer id, same bytes, forever.
func SVG(offerID int64, n int) string {
	h := hash64(offerID*1000003 + int64(n))
	arch := archetypes[int(uint64(offerID)%uint64(len(archetypes)))]

	hue := int(h % 360)
	hue2 := (hue + 30 + int((h>>9)%60)) % 360
	light := 46 + int((h>>17)%14)

	bg := hsl(hue, 62, 92)
	fg := hsl(hue, 70, 22)
	a := hsl(hue, 68, light)
	b := hsl(hue2, 64, light+12)

	var s strings.Builder
	s.Grow(1400)
	s.WriteString(`<svg xmlns="http://www.w3.org/2000/svg" width="600" height="600" viewBox="0 0 600 600" role="img">`)
	s.WriteString(`<title>` + esc(arch.NameEN) + ` ` + strconv.FormatInt(offerID, 10) + `</title>`)
	s.WriteString(`<defs><linearGradient id="g" x1="0" y1="0" x2="1" y2="1">`)
	s.WriteString(`<stop offset="0" stop-color="` + a + `"/><stop offset="1" stop-color="` + b + `"/>`)
	s.WriteString(`</linearGradient></defs>`)
	s.WriteString(`<rect width="600" height="600" fill="` + bg + `"/>`)

	// A shape family per image index, so an offer's images differ from each other.
	switch n % 4 {
	case 0:
		s.WriteString(`<rect x="110" y="110" width="380" height="380" rx="46" fill="url(#g)"/>`)
	case 1:
		s.WriteString(`<circle cx="300" cy="300" r="190" fill="url(#g)"/>`)
	case 2:
		s.WriteString(`<polygon points="300,90 510,300 300,510 90,300" fill="url(#g)"/>`)
	default:
		s.WriteString(`<rect x="120" y="160" width="360" height="280" rx="24" fill="url(#g)"/>`)
		s.WriteString(`<circle cx="300" cy="300" r="86" fill="` + bg + `" opacity="0.85"/>`)
	}

	// Deterministic speckle, so two offers with the same shape still differ.
	for i := 0; i < 5; i++ {
		hh := hash64(int64(h) + int64(i))
		cx := 80 + int(hh%440)
		cy := 80 + int((hh>>13)%440)
		r := 8 + int((hh>>27)%26)
		s.WriteString(`<circle cx="` + strconv.Itoa(cx) + `" cy="` + strconv.Itoa(cy) + `" r="` + strconv.Itoa(r) + `" fill="` + fg + `" opacity="0.10"/>`)
	}

	s.WriteString(`<text x="40" y="72" font-family="system-ui,-apple-system,Segoe UI,Helvetica,Arial,sans-serif" font-size="30" font-weight="600" fill="` + fg + `">` + esc(arch.NameEN) + `</text>`)
	s.WriteString(`<text x="40" y="548" font-family="ui-monospace,SFMono-Regular,Menlo,Consolas,monospace" font-size="30" fill="` + fg + `">` + strconv.FormatInt(offerID, 10) + `</text>`)
	s.WriteString(`<text x="520" y="548" text-anchor="end" font-family="ui-monospace,SFMono-Regular,Menlo,Consolas,monospace" font-size="30" fill="` + fg + `" opacity="0.7">#` + strconv.Itoa(n) + `</text>`)
	s.WriteString(`</svg>`)
	return s.String()
}

func hsl(h, s, l int) string {
	return "hsl(" + strconv.Itoa(h) + "," + strconv.Itoa(s) + "%," + strconv.Itoa(l) + "%)"
}

// esc escapes the handful of characters that matter inside SVG text.
func esc(s string) string {
	r := strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;", `"`, "&quot;")
	return r.Replace(s)
}
