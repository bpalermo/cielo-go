package types

const (
	transactionalSandboxBaseUrl    = "https://apisandbox.cieloecommerce.cielo.com.br"
	transactionalProductionBaseUrl = "https://apisandbox.cieloecommerce.cielo.com.br"
	querySandboxBaseUrl            = "https://apiquerysandbox.cieloecommerce.cielo.com.br"
	queryProductionBaseUrl         = "https://apiquerysandbox.cieloecommerce.cielo.com.br"
)

type Environment struct {
	ApiUrl      string
	ApiQueryUrl string
}

var (
	// ProductionEnvironment sets the base URLs to production
	ProductionEnvironment = &Environment{
		ApiUrl:      transactionalProductionBaseUrl,
		ApiQueryUrl: queryProductionBaseUrl,
	}
	// SandboxEnvironment sets the base URLs to sandbox
	SandboxEnvironment = &Environment{
		ApiUrl:      transactionalSandboxBaseUrl,
		ApiQueryUrl: querySandboxBaseUrl,
	}
)
