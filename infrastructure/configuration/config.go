package configuration

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

//[.'.]:> 📦 SISTEMA DE CONFIGURACIÓN
//[.'.]:> ========================
//[.'.]:> Este módulo maneja la configuración de nuestra aplicación siguiendo este flujo:
//[.'.]:>
//[.'.]:> 1️⃣ ARCHIVO CONFIG: Busca primero un archivo app.env en el directorio especificado
//[.'.]:>    Si lo encuentra, carga todas sus variables como configuración base
//[.'.]:>
//[.'.]:> 2️⃣ VARIABLES DE ENTORNO: Después de cargar el archivo (o si no existe):
//[.'.]:>    - Comprueba si existen variables de entorno con los mismos nombres
//[.'.]:>    - Las variables de entorno tienen PRIORIDAD y sobrescriben los valores del archivo
//[.'.]:>
//[.'.]:> 3️⃣ VALORES POR DEFECTO: Como último recurso, si algún valor sigue vacío,
//[.'.]:>    se aplican valores predeterminados para garantizar que la app pueda funcionar

// [.'.]:> 🧩 Estructura que contiene toda la configuración de la aplicación
// [.'.]:> Cada campo se mapea a una variable de entorno o valor en app.env del mismo nombre
type BoilerplateConfiguration struct {
	DatabaseDsn       string `mapstructure:"DATABASE_DSN"`        // Cadena de conexión a la base de datos
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"` // Dirección del servidor HTTP (ej: ":3000")
}

// [.'.]:> 🔄 Función principal que carga toda la configuración
// [.'.]:> Recibe la ruta donde buscar el archivo app.env y devuelve la configuración completa
// [.'.]:> Si hay algún error durante la carga, lo devuelve para que la aplicación pueda manejarlo
func LoadConfiguration(configFilesPath string) (config BoilerplateConfiguration, err error) {
	viper.AddConfigPath(configFilesPath)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.BindEnv("DATABASE_DSN")
	viper.BindEnv("HTTP_SERVER_ADDRESS")
	viper.AutomaticEnv()

	if err := loadConfigFile(); err != nil {
		return config, err
	}

	overrideWithEnvVars()

	err = viper.Unmarshal(&config)
	if err == nil {
		fmt.Println("['.']:> ✨ Configuración cargada correctamente ✨")
		fmt.Println("['.']:> ----------------------------------------")
		fmt.Printf("['.']:> 🌐 HTTP Server: %s\n", config.HTTPServerAddress)
		fmt.Println("['.']:> ----------------------------------------")
		applyFallbackEnvVars(&config)
	}

	return config, err
}

// [.'.]:> 📄 Intenta cargar el archivo de configuración app.env
// [.'.]:> Si el archivo no existe, lo maneja elegantemente y permite continuar
// [.'.]:> usando solo variables de entorno
func loadConfigFile() error {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
		fmt.Println("['.']:> 📋 Archivo de configuración no encontrado, usando solo variables de entorno")
	} else {
		fmt.Println("['.']:> 📋 Archivo de configuración cargado correctamente")
	}

	return nil
}

// [.'.]:> 🔀 Sobrescribe valores del archivo con variables de entorno
// [.'.]:> Esta es la clave para que las variables de entorno tengan prioridad
// [.'.]:> sobre el archivo de configuración
func overrideWithEnvVars() {
	dsnEnv := os.Getenv("DATABASE_DSN")
	if dsnEnv != "" && viper.GetString("DATABASE_DSN") == "" {
		viper.Set("DATABASE_DSN", dsnEnv)
	}

	httpEnv := os.Getenv("HTTP_SERVER_ADDRESS")
	if httpEnv != "" && viper.GetString("HTTP_SERVER_ADDRESS") == "" {
		viper.Set("HTTP_SERVER_ADDRESS", httpEnv)
	}
}

// [.'.]:> 🔒 Aplica valores de respaldo directamente desde variables de entorno
// [.'.]:> como última red de seguridad para los campos que aún estén vacíos
// [.'.]:> después de procesar el archivo y las variables a través de viper
func applyFallbackEnvVars(config *BoilerplateConfiguration) {
	if config.DatabaseDsn == "" {
		config.DatabaseDsn = os.Getenv("DATABASE_DSN")
	}
	if config.HTTPServerAddress == "" {
		config.HTTPServerAddress = os.Getenv("HTTP_SERVER_ADDRESS")
	}
}
