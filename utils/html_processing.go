package utils

import "regexp"

func ExtractionTel(html string) []string {
	mobile := ExtractionMobileTel(html)
	freeDial := ExtractionFreeDial(html)
	landLine := ExtractionLandLine(html)

	appended := append(mobile, freeDial...)
	appended = append(appended, landLine...)
	return UniqStringArray(appended)
}

func ExtractionMobileTel(html string) []string {
	r := regexp.MustCompile(`0[5789]0(-|)\d{4}(-|)\d{4}`)
	return r.FindAllString(html, -1)
}

func ExtractionFreeDial(html string) []string {
	r := regexp.MustCompile(`0120(-|)\d{3}(-|)\d{3}`)
	return r.FindAllString(html, -1)
}

func ExtractionLandLine(html string) []string {
	// 03-2222-2222
	// (03)2222-2222
	r1 := regexp.MustCompile(`(\(|)0\d(\)|-)\d{4}(-|)\d{4}`)
	s1 := r1.FindAllString(html, -1)

	// 0422-22-2222
	// (0422)22-2222
	r2 := regexp.MustCompile(`(\(|)0\d{3}(\)|-)\d{2}(-|)\d{4}`)
	s2 := r2.FindAllString(html, -1)
	return append(s1, s2...)
}

func ExtractionEmail(html string) []string {
	r := regexp.MustCompile(`[a-zA-Z0-9.!#$%&'*+\/=?^_{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*`)
	s := r.FindAllString(html, -1)
	return UniqStringArray(s)
}
