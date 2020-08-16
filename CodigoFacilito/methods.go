package CodigoFacilito

func (self *Curso) getTitulo() string {
	return self.Titulo
}

func (self *Curso) ObtenerTitulo() string {
	return self.getTitulo()
}