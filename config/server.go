package config

type Path struct{
	file string,
	response int
}

type Server struct{
	Port string,
	Paths []Path
}