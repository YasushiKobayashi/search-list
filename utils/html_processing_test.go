package utils

import (
	"testing"
)

func ReadTestFile() string {
	return `<a href="tel:090-1234-1234" >tel</a>
<a href="tel:09012341234" >tel</a>
<a href="tel:0120-222-222" >tel</a>
<a href="tel:0120222222" >tel</a>
<a href="tel:03-2222-2222" >tel</a>
<a href="tel:(03)2222-2222" >tel</a>
<a href="tel:0422-22-2222" >tel</a>
<a href="tel:(0422)22-2222" >tel</a>

<a href="mailto:ptpadan@gmail.com" >email</a>`
}

func Test_ExtractionTel(t *testing.T) {
	str := ReadTestFile()
	tels := ExtractionTel(str)
	if len(tels) != 8 {
		t.Fatalf("invalid get tels %#v", tels)
	}
}

func Test_ExtractionMobileTel(t *testing.T) {
	str := ReadTestFile()
	tels := ExtractionMobileTel(str)
	if len(tels) != 2 {
		t.Fatalf("invalid get tels %#v", tels)
	}
}

func Test_ExtractionFreeDial(t *testing.T) {
	str := ReadTestFile()
	tels := ExtractionFreeDial(str)
	if len(tels) != 2 {
		t.Fatalf("invalid get tels %#v", tels)
	}
}

func Test_ExtractionLandLine(t *testing.T) {
	str := ReadTestFile()
	tels := ExtractionLandLine(str)
	if len(tels) != 4 {
		t.Fatalf("invalid get tels %#v", tels)
	}
}

func Test_ExtractionEmail(t *testing.T) {
	str := ReadTestFile()
	tels := ExtractionEmail(str)
	if len(tels) != 1 {
		t.Fatalf("invalid get tels %#v", tels)
	}
}
