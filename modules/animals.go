package modules

//config animals
type Config struct {
	Animals  []NameType `yaml:"animals"`
	Database Database   `yaml:"database"`
}

//config name and type
type NameType struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

// config Database
type Database struct {
	Host string `yaml:"host" validate:"required"`
	Port int    `yaml:"port" validate:"required"`
	Name string `yaml:"name" validate:"required"`
	User string `yaml:"user" validate:"required"`
	Pass string `yaml:"pass" validate:"required"`
	Char string `yaml:"char" validate:"required"`
}
