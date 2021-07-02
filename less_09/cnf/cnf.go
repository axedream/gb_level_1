package cnf

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"gopkg.in/yaml.v2"
	"strings"
)

var configFileName *string

type Config struct {
	Port string
	Url string
	User string
}

func NewConfig() *Config  {
	return &Config{
	}
}

/**
	Проверяет дефолтные значения
 */
func testConfig(cn *Config)  {
	if cn.Port =="" {
		cn.Port = "8090"
	}

	if cn.Url == "" {
		cn.Url = "/"
	}

	if cn.User == "" {
		cn.User = "Guest"
	}
}


/**
	Загрузка конфигурации переменных из файла json
*/
func (t *Config) loadConfigFileYaml() {

	configData,err := ioutil.ReadFile(*configFileName)
	if err != nil {
		log.Fatalln("Ошибка чтения файла: ",err)
	}

	cn:=Config{}
	err = yaml.Unmarshal(configData, &cn)
	if err != nil {
		log.Fatalln("Ошибка парсинка yaml формата: ", err)
	}

	t.Port = cn.Port
	t.Url = cn.Url
	t.User = cn.User

	testConfig(t);

}


/**
	Загрузка конфигурации переменных из файла json
 */
func (t *Config) loadConfigFileJson() {

	configData,err := os.ReadFile(*configFileName)
	if err != nil {
		log.Fatalln("Ошибка чтения файла: ",err)
	}

	cn:=Config{}
	err = json.Unmarshal(configData, &cn)
	if err != nil {
		log.Fatalln("Ошибка парсинка json формата: ", err)
	}

	t.Port = cn.Port
	t.Url = cn.Url
	t.User = cn.User

	testConfig(t);
}

/**
	Загрузка конфигурации из переменных окружения
 */
func (t *Config) LoadConfigEnv() {

	port := os.Getenv("PORT")
	url := os.Getenv("NAME")
	user := os.Getenv("USER")

	t.Port = port
	t.Url = url
	t.User = user

	testConfig(t)
}


/**
Определяем тип файла и загрузка нужного формата
*/

func (t *Config) LoadConfiFile()  {

	configFileName = flag.String("config","","Config file")
	flag.Parse()

	ss := strings.Split(*configFileName,".")


	if len(ss) != 2 {
		log.Fatalln("Имя файла конфигурационного файла или его формат заданы не верно!")
	}


	switch ss[1] {
		case "yaml":
			t.loadConfigFileYaml()
		case "json":
			t.loadConfigFileJson()
		default:
			log.Fatalln("Имя файла конфигурационного файла или его формат заданы не верно!")
	}

}

