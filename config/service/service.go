package service

// // service :
// type service struct {
// 	Host string `yaml:"host"`
// 	Port string `yaml:"port"`
// }

type ServiceList struct {
	Articles struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"articles"`
}
