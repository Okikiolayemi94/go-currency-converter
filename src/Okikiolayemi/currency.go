package Okikiolayemi

type Naira float32

type Peso float32
type Real float32
type Cedi float32
type Euro float32
type Yen float32
type Quetzal float32
type Lempira float32
type Rupiah float32
type Rupee float32
type Krona float32

func (n Naira) PesoConverter() Peso {
	return Peso(n * 0.3424)
}
func (n Naira) RealConverter() Real {
	return Real(n * 0.01192)
}
func (n Naira) CediConverter() Cedi {
	return Cedi(n * 0.02391)
}
func (n Naira) EuroConverter() Euro {
	return Euro(n * 0.00233)
}
func (n Naira) YenConverter() Yen {
	return Yen(n * 0.3335)
}
func (n Naira) QuetzalConverter() Quetzal {
	return Quetzal(n * 0.0177)
}
func (n Naira) LempiraConverter() Lempira {
	return Lempira(n * 0.05571)
}
func (n Naira) RupiahConverter() Rupiah {
	return Rupiah(n * 35.129)
}
func (n Naira) RupeeConverter() Rupee {
	return Rupee(n * 0.18801)
}
func (n Naira) KronaConverter() Krona {
	return Krona(n * 0.3298)
}
