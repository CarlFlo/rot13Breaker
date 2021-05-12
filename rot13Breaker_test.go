package rot13Breaker

import (
	"testing"
)

func TestDecrypt(t *testing.T) {

	input := "Znoy oy g zkyz ul znk iujk"
	expectedShifts := 20

	output := Decrypt(input)

	if output[0].Shift != expectedShifts {
		t.Fatalf("Expected first result to show %d shifts, got %d", expectedShifts, output[0].Shift)
	}

}

func TestDecryptLong(t *testing.T) {

	input := "Lpxwxg bl itkm hy max zxhzktiabvte tkxt hy Yxgghlvtgwbt. Max vebftmx bl bg zxgxkte fbew yhk bml ghkmaxker etmbmnwx wnx mh lbzgbybvtgm ftkbmbfx bgyenxgvx. Bg libmx hy max abza etmbmnwx, Lpxwxg hymxg atl ptkf vhgmbgxgmte lnffxkl, uxbgz ehvtmxw bg uxmpxxg max Ghkma Tmetgmbv, max Utembv Lxt, tgw otlm Knllbt. Max zxgxkte vebftmx tgw xgobkhgfxgm otkr lbzgbybvtgmer ykhf max lhnma tgw ghkma wnx mh max otlm etmbmnwte wbyyxkxgvx, tgw fnva hy Lpxwxg atl kxebtuer vhew tgw lghpr pbgmxkl. Lhnmaxkg Lpxwxg bl ikxwhfbgtgmer tzkbvnemnkte, pabex max ghkma bl axtober yhkxlmxw tgw bgvenwxl t ihkmbhg hy max Lvtgwbgtobtg Fhngmtbgl."
	expectedShifts := 7

	output := Decrypt(input)

	if output[0].Shift != expectedShifts {
		t.Fatalf("Expected first result to show %d shifts, got %d", expectedShifts, output[0].Shift)
	}
}
