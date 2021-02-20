package resource

import "server/devops-tools/ldap/initconfig"

func ActionLdapRegister() {
	conn := initconfig.InitConfig()
	conn.Bind()
}
