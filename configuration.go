package configuration

import (
	"os"
	"strconv"

	"github.com/prometheus/common/log"
)

// Configuration is a structure of environment settings for services
// and topics channel name.
type Configuration struct {
	APIVersion,
	ServiceName string

	Production struct {
		InitialTopic      string
		SprootTopic       string
		LogunaTopic       string
		LogFilePath       string
		HecatoncheirTopic string

		Broker struct {
			Host string
			Port int
		}

		HTTPServer struct {
			Host                 string
			Port                 int
			StaticFilesDirectory string
		}

		SocketServer struct {
			Host string
			Port int
		}

		Database struct {
			Host string
			Port int
		}
	}

	Development struct {
		InitialTopic,
		SprootTopic,
		LogunaTopic,
		LogFilePath,
		HecatoncheirTopic string

		Broker struct {
			Host string
			Port int
		}

		HTTPServer struct {
			Host string
			Port int
		}

		SocketServer struct {
			Host string
			Port int
		}

		Database struct {
			Host string
			Port int
		}
	}
}

// New a constructor for *Configuration
func New() *Configuration {
	configuration := Configuration{}

	apiVersion := os.Getenv("API-Version")
	if apiVersion == "" {
		configuration.APIVersion = "1.0.0"
	} else {
		configuration.APIVersion = apiVersion
	}

	serviceName := os.Getenv("Service-Name")
	if serviceName == "" {
		configuration.ServiceName = ""
	} else {
		configuration.ServiceName = serviceName
	}

	// Services production
	productionInitialTopic := os.Getenv("Production-Initial-Topic")
	if productionInitialTopic == "" {
		configuration.Production.InitialTopic = "Initial"
	} else {
		configuration.Production.InitialTopic = productionInitialTopic
	}

	productionSprootTopic := os.Getenv("Production-Sproot-Topic")
	if productionSprootTopic == "" {
		configuration.Production.SprootTopic = "Sproot"
	} else {
		configuration.Production.SprootTopic = productionSprootTopic
	}

	productionLogunaTopic := os.Getenv("Production-Loguna-Topic")
	if productionLogunaTopic == "" {
		configuration.Production.LogunaTopic = "Loguna"
	} else {
		configuration.Production.LogunaTopic = productionLogunaTopic
	}

	productionHecatoncheirTopic := os.Getenv("Production-Hecatoncheir-Topic")
	if productionHecatoncheirTopic == "" {
		configuration.Production.HecatoncheirTopic = "Hecatoncheir"
	} else {
		configuration.Production.HecatoncheirTopic = productionHecatoncheirTopic
	}

	// Services Development
	developmentInitialTopic := os.Getenv("Development-Initial-Topic")
	if developmentInitialTopic == "" {
		configuration.Development.InitialTopic = "DevInitial"
	} else {
		configuration.Development.InitialTopic = developmentInitialTopic
	}

	developmentSprootTopic := os.Getenv("Development-Sproot-Topic")
	if developmentSprootTopic == "" {
		configuration.Development.SprootTopic = "DevSproot"
	} else {
		configuration.Development.SprootTopic = developmentSprootTopic
	}

	developmentLogunaTopic := os.Getenv("Development-Loguna-Topic")
	if developmentLogunaTopic == "" {
		configuration.Development.LogunaTopic = "DevLoguna"
	} else {
		configuration.Development.LogunaTopic = developmentLogunaTopic
	}

	developmentHecatoncheirTopic := os.Getenv("Development-Hecatoncheir-Topic")
	if developmentHecatoncheirTopic == "" {
		configuration.Development.HecatoncheirTopic = "DevHecatoncheir"
	} else {
		configuration.Development.HecatoncheirTopic = developmentHecatoncheirTopic
	}

	// Broker Production
	productionBrokerHostFromEnvironment := os.Getenv("Production-Broker-Host")
	if productionBrokerHostFromEnvironment == "" {
		configuration.Production.Broker.Host = "192.168.99.100"
	} else {
		configuration.Production.Broker.Host = productionBrokerHostFromEnvironment
	}

	productionBrokerPortFromEnvironment := os.Getenv("Production-Broker-Port")
	if productionBrokerPortFromEnvironment == "" {
		configuration.Production.Broker.Port = 4222
	} else {
		port, err := strconv.Atoi(productionBrokerPortFromEnvironment)
		if err != nil {
			log.Fatal(err)
		}

		configuration.Production.Broker.Port = port
	}

	// HTTPServer Production
	productionHTTPServerHostFromEnvironment := os.Getenv("Production-HTTPServer-Host")
	if productionHTTPServerHostFromEnvironment == "" {
		configuration.Production.HTTPServer.Host = "localhost"
	} else {
		configuration.Production.HTTPServer.Host = productionHTTPServerHostFromEnvironment
	}

	productionHTTPServerPortFromEnvironment := os.Getenv("Production-HTTPServer-Port")
	if productionHTTPServerPortFromEnvironment == "" {
		configuration.Production.HTTPServer.Port = 80
	} else {
		port, err := strconv.Atoi(productionHTTPServerPortFromEnvironment)
		if err != nil {
			log.Fatal(err)
		}

		configuration.Production.HTTPServer.Port = port
	}

	productionHTTPServerStaticFilesDirectoryFromEnvironment := os.Getenv("Production-HTTPServer-StaticFilesDirectory")
	if productionHTTPServerStaticFilesDirectoryFromEnvironment == "" {
		configuration.Production.HTTPServer.StaticFilesDirectory = "build/web"
	} else {
		configuration.Production.HTTPServer.StaticFilesDirectory = productionHTTPServerStaticFilesDirectoryFromEnvironment
	}

	// SocketServer Production
	productionSocketServerHostFromEnvironment := os.Getenv("Production-SocketServer-Host")
	if productionSocketServerHostFromEnvironment == "" {
		configuration.Production.SocketServer.Host = "localhost"
	} else {
		configuration.Production.SocketServer.Host = productionSocketServerHostFromEnvironment
	}

	productionSocketServerPortFromEnvironment := os.Getenv("Production-SocketServer-Port")
	if productionSocketServerPortFromEnvironment == "" {
		configuration.Production.SocketServer.Port = 81
	} else {
		port, err := strconv.Atoi(productionSocketServerPortFromEnvironment)
		if err != nil {
			log.Fatal(err)
		}

		configuration.Production.SocketServer.Port = port
	}

	developmentBrokerHostFromEnvironment := os.Getenv("Development-Broker-Host")
	if developmentBrokerHostFromEnvironment == "" {
		configuration.Development.Broker.Host = "192.168.99.100"
	} else {
		configuration.Development.Broker.Host = developmentBrokerHostFromEnvironment
	}

	developmentBrokerPortFromEnvironment := os.Getenv("Development-Broker-Port")
	if developmentBrokerPortFromEnvironment == "" {
		configuration.Development.Broker.Port = 4222
	} else {
		port, err := strconv.Atoi(developmentBrokerPortFromEnvironment)
		if err != nil {
			log.Fatal(err)
		}

		configuration.Development.Broker.Port = port
	}

	// HTTPServer Development
	developmentHTTPHostFromEnvironment := os.Getenv("Development-HTTPServer-Port")
	if developmentHTTPHostFromEnvironment == "" {
		configuration.Development.HTTPServer.Host = "localhost"
	} else {
		configuration.Development.HTTPServer.Host = developmentHTTPHostFromEnvironment
	}

	developmentHTTPServerPortFromEnvironment := os.Getenv("Development-HTTPServer-Port")
	if developmentHTTPServerPortFromEnvironment == "" {
		configuration.Development.HTTPServer.Port = 8000
	} else {
		port, err := strconv.Atoi(developmentBrokerPortFromEnvironment)
		if err != nil {
			log.Fatal(err)
		}

		configuration.Development.HTTPServer.Port = port
	}

	// SocketServer Development
	developmentSocketHostFromEnvironment := os.Getenv("Development-SocketServer-Port")
	if developmentSocketHostFromEnvironment == "" {
		configuration.Development.SocketServer.Host = "localhost"
	} else {
		configuration.Development.SocketServer.Host = developmentSocketHostFromEnvironment
	}

	developmentSocketServerPortFromEnvironment := os.Getenv("Development-SocketServer-Port")
	if developmentSocketServerPortFromEnvironment == "" {
		configuration.Development.SocketServer.Port = 8001
	} else {
		port, err := strconv.Atoi(developmentSocketServerPortFromEnvironment)
		if err != nil {
			log.Fatal(err)
		}

		configuration.Development.SocketServer.Port = port
	}

	// Database Production
	productionDatabaseHostFromEnvironment := os.Getenv("Production-Database-Host")
	if productionDatabaseHostFromEnvironment == "" {
		configuration.Production.Database.Host = "192.168.99.100"
	} else {
		configuration.Production.Database.Host = productionDatabaseHostFromEnvironment
	}

	productionDatabasePortFromEnvironment := os.Getenv("Production-Database-Port")
	if productionDatabasePortFromEnvironment == "" {
		configuration.Production.Database.Port = 9080
	} else {
		port, err := strconv.Atoi(productionDatabasePortFromEnvironment)
		if err != nil {
			log.Fatal(err)
		}

		configuration.Production.Database.Port = port
	}

	// Database Development
	developmentDatabaseHostFromEnvironment := os.Getenv("Development-Database-Host")
	if developmentDatabaseHostFromEnvironment == "" {
		configuration.Development.Database.Host = "192.168.99.100"
	} else {
		configuration.Development.Database.Host = developmentDatabaseHostFromEnvironment
	}

	developmentDatabasePortFromEnvironment := os.Getenv("Development-Database-Port")
	if developmentDatabasePortFromEnvironment == "" {
		configuration.Development.Database.Port = 9080
	} else {
		port, err := strconv.Atoi(developmentDatabasePortFromEnvironment)
		if err != nil {
			log.Fatal(err)
		}

		configuration.Development.Database.Port = port
	}

	// LogFilePath
	productionLogFilePath := os.Getenv("Production-Log-File-Path")
	if productionLogFilePath == "" {
		configuration.Production.LogFilePath = "log"
	} else {
		configuration.Production.LogFilePath = productionLogFilePath
	}

	developmentLogFilePath := os.Getenv("Development-Log-File-Path")
	if developmentLogFilePath == "" {
		configuration.Development.LogFilePath = "dev_log"
	} else {
		configuration.Development.LogFilePath = developmentLogFilePath
	}

	return &configuration
}
