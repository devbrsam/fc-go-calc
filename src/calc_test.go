package main
import "testing"

func TestCalc( t *testing.T ) {

	deveRetornar4 := sum(2,2)
	
	if deveRetornar4 != 4 {
		t.Errorf("sum(2,2) teste FALHOU: esperado %v, mas retornou %v", 4, deveRetornar4)
	} else {
		t.Logf("sum(2,2) teste OK")
	}	
}