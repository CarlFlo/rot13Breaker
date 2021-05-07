package rot13Breaker

import (
	"testing"
)

func TestDecrypt(t *testing.T) {

	input := "QEB NRFZH YOLTK CLU GRJMP LSBO QEB IXWV ALD"
	expectedShifts := 3

	output, err := Decrypt(input)

	if err != nil {
		t.Fatalf("This should not have failed: %v", err)
	}

	if output[0].shift != expectedShifts {
		t.Fatalf("Expected first result with the result of %d shifts, got %d", expectedShifts, output[0].shift)
	}

}

func TestDecryptLong(t *testing.T) {

	input := "Zdlklu pz whya vm aol nlvnyhwopjhs hylh vm Mluuvzjhukph. Aol jspthal pz pu nlulyhs tpsk mvy paz uvyaolysf shapabkl kbl av zpnupmpjhua thypaptl pumsblujl. Pu zwpal vm aol opno shapabkl, Zdlklu vmalu ohz dhyt jvuapuluahs zbttlyz, ilpun svjhalk pu iladllu aol Uvyao Hashuapj, aol Ihsapj Zlh, huk chza Ybzzph. Aol nlulyhs jspthal huk lucpyvutlua chyf zpnupmpjhuasf myvt aol zvbao huk uvyao kbl av aol chza shapabkhs kpmmlylujl, huk tbjo vm Zdlklu ohz ylsphisf jvsk huk zuvdf dpualyz. Zvbaolyu Zdlklu pz wylkvtpuhuasf hnypjbsabyhs, dopsl aol uvyao pz olhcpsf mvylzalk huk pujsbklz h wvyapvu vm aol Zjhukpuhcphu Tvbuahpuz."
	expectedShifts := 19

	output, err := Decrypt(input)

	if err != nil {
		t.Fatalf("This should not have failed: %v", err)
	}

	if output[0].shift != expectedShifts {
		t.Fatalf("Expected first result with the result of %d shifts, got %d", expectedShifts, output[0].shift)
	}
}

func TestErrors(t *testing.T) {
	input := "short"

	_, err := Decrypt(input)

	if err == nil {
		t.Fatalf("The input is too short so this should have failed.")
	}

}
