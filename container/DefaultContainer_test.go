package container

import (
	"github.com/thingfu/hub/api"
	"github.com/thingfu/hub/data/source"
	"testing"
)

func TestInitializeContainer(t *testing.T) {
	cfg := new(api.Configuration)
	cfg.Db = "localhost:27017"
	cfg.MaxProcs = 1
	cfg.NodeID = "LocalDev"
	mailConfig := new(api.MailConfig)
	mailConfig.Host = "smtp.gmail.com"
	mailConfig.User = "joe"
	mailConfig.Pass = "secret"
	mailConfig.Port = 579
	cfg.Mail = *mailConfig

	container, env := Initialize("./home", *cfg)
	if container == nil {
		t.Fail()
	}

	if env == nil {
		t.Fail()
	}

	if container.DataSource() == nil {
		t.Fail()
	}

	if container.Env() == nil {
		t.Fail()
	}

	if container.Factory() == nil {
		t.Fail()
	}

	if container.ProtocolHandlers() == nil {
		t.Fail()
	}

	if container.RulesService() == nil {
		t.Fail()
	}

	if container.ScheduleService() == nil {
		t.Fail()
	}

	if container.ThingManager() == nil {
		t.Fail()
	}

	// ---------------------------------------------
}

func TestWiringContainer(t *testing.T) {
	container := new(DefaultContainer)

	container.Register(new(source.MongoDataSource), "api.DataSource")
	_, ok := container.DataSource().(api.DataSource)
	if !ok {
		t.Fatal("Failed registering Data Source Type")
	}

	/*
		case t == "api.Environment":
			c.environment = svc.(api.Environment)

		case t == "api.RulesService":
			c.rulesService = svc.(api.RulesService)

		case t == "api.ThingManager":
			c.thingManager = svc.(api.ThingManager)

		case t == "api.ScheduleService":
			c.scheduleService = svc.(api.ScheduleService)

		case t == "api.DataSource":
			c.dataSource = svc.(api.DataSource)

		case t == "api.Factory":
			c.factory = svc.(api.Factory)

		case t == "api.ProtocolHandler":
			name := svc.(api.ProtocolHandler).GetName()
			c.protocolHandlers[name] = svc.(api.ProtocolHandler)
	*/

	/*
		rulesService := c.RulesService()
			factory := c.Factory()
			thingManager := c.ThingManager()
			env := c.Env()
			dataSource := c.DataSource()
			scheduleServices := c.ScheduleService()

			// Wire Up Services
			// Rules Service
			rulesService.SetThingManager(thingManager)
			rulesService.SetFactory(factory)

			// Factory

			// ThingManager
			thingManager.SetRulesService(rulesService)
			thingManager.SetFactory(factory)
			thingManager.SetDataSource(dataSource)

			// DataSource
			dataSource.SetEnvironment(env)

			// ScheduleService
			scheduleServices.SetRulesService(rulesService)
			scheduleServices.SetThingManager(thingManager)

			// Protocol Handlers
			c.protocolHandlers = make(map[string]api.ProtocolHandler)

			services := make([]api.ContainerAware, 6)
			services[0] = rulesService
			services[1] = factory
			services[2] = thingManager
			services[3] = env
			services[4] = dataSource
			services[5] = scheduleServices

			// Inject Container into all ContainerAware Services
			for _, service := range services {
				service.SetContainer(c)
			}

			// Validate Wiring
			for _, service := range services {
				service.ValidateWiring()
			}
	*/
}

/*
{
    "NodeID": "LocalDev",
    "MaxProcs": 1,
    "DB": "localhost:27017",
    "Protocols": {
        "sim": {
            "Enabled": true,
            "Conn": "http"
        },
        "RF433": {
            "Enabled": true,
            "Conn": "Serial",
            "Port": "/dev/tty.usbmodem1411",
            "Baud": 9600
        }
    },
    "Mail": {
        "User": "username@host",
        "Pass": "passwd",
        "Host": "smtp_host",
        "Port": 587
    }
}
*/
