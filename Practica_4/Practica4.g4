grammar Practica4;
/*
 S ::= ( num , num ) L
 
 L ::= L ( num , num ) 
 | ( num , num ) 
 */

options{
    language = 'Go';

}

@parser::header {
    import ("Practica4/Node"

	"strings"
    "math"
    )
}

@parser::members {

    func mostrarResultado(out string){

        cadena := strings.Split(out, " ")


	    var arreglito []float64
	    for i := 0; i < len(cadena); i++ {

		flotante, _ := strconv.ParseFloat(cadena[i], 64)

		arreglito = append(arreglito, flotante)
	}
	    var res float64 = 0
	    for i := 0; i < len(arreglito)-1; i++ {
		    if i+2 >= (len(arreglito) - 1) {
			    break
		}

		op1 := math.Pow(arreglito[i+2]-arreglito[i], 2)
		op2 := math.Pow(arreglito[i+3]-arreglito[i+1], 2)
		res = res + math.Sqrt(op1+op2)

		i = i + 1
		if i >= (len(arreglito) - 1) {
			break
		}
	}

	    fmt.Println("Distancia: ", fmt.Sprintf("%.2f", res))
        fmt.Println("Sale el laboratorio jeje")

    }
    
}

//Rules
start
	returns[node.Nodo n]:
	un = lista EOF {
                        mostrarResultado($un.n.A + $un.n.B )
    };
lista
	returns[node.Nodo n]:
	na = lista PARA num = NUMBER COMA num2 = NUMBER PARC {
                                        x := $na.n.A + " " + $num.text + " "+$num2.text
                                        y := $na.n.B 
                                        $n = node.NewNodo(x, y)
    }
	| PARA num = NUMBER COMA num2 = NUMBER PARC {
                                    x := $num.text + " " + $num2.text
                                    y := "" 
                                    $n = node.NewNodo(x,y)
    };

//Tokens
PARA: '(';
PARC: ')';
COMA: ',';
NUMBER: ('-')? [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;
