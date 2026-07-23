# Command catalog

Generated from the govultr service interfaces by `go generate ./...`.
Do not edit by hand. Every operation is listed as
`vultr <service> <operation> <args>` followed by its govultr signature.

## Global flags

| flag | type | default | description |
| --- | --- | --- | --- |
| `--api-key` | string | — | Vultr API key (defaults to VULTR_API_KEY environment variable) |
| `-4`, `--ipv4` | bool | `false` | force IPv4 connections to the API (useful with IP-restricted API keys) |

## `vultr account`

Account API (2 operations)

- `vultr account get` — Get() -> *Account
- `vultr account get-bandwidth` — GetBandwidth() -> *AccountBandwidth

## `vultr application`

Application API (1 operations)

- `vultr application list` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — List(*ListOptions) -> []Application, *Meta

## `vultr backup`

Backup API (2 operations)

- `vultr backup get <string>` — Get(string) -> *Backup
- `vultr backup list` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — List(*ListOptions) -> []Backup, *Meta

## `vultr bare-metal-server`

BareMetalServer API (24 operations)

- `vultr bare-metal-server attach-vpc <string> <string>` — AttachVPC(string, string)
- `vultr bare-metal-server attach-vpc2 <string> <json:AttachVPC2Req>` `[--schema]` — AttachVPC2(string, *AttachVPC2Req)
- `vultr bare-metal-server create <json:BareMetalCreate>` `[--schema]` — Create(*BareMetalCreate) -> *BareMetalServer
- `vultr bare-metal-server delete <string>` — Delete(string)
- `vultr bare-metal-server detach-vpc <string> <string>` — DetachVPC(string, string)
- `vultr bare-metal-server detach-vpc2 <string> <string>` — DetachVPC2(string, string)
- `vultr bare-metal-server get <string>` — Get(string) -> *BareMetalServer
- `vultr bare-metal-server get-bandwidth <string>` — GetBandwidth(string) -> *Bandwidth
- `vultr bare-metal-server get-upgrades <string>` — GetUpgrades(string) -> *Upgrades
- `vultr bare-metal-server get-user-data <string>` — GetUserData(string) -> *UserData
- `vultr bare-metal-server get-vnc-url <string>` — GetVNCUrl(string) -> *VNCUrl
- `vultr bare-metal-server halt <string>` — Halt(string)
- `vultr bare-metal-server list` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — List(*ListOptions) -> []BareMetalServer, *Meta
- `vultr bare-metal-server list-ipv4s <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListIPv4s(string, *ListOptions) -> []IPv4, *Meta
- `vultr bare-metal-server list-ipv6s <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListIPv6s(string, *ListOptions) -> []IPv6, *Meta
- `vultr bare-metal-server list-vpc-info <string>` — ListVPCInfo(string) -> []VPCInfo
- `vultr bare-metal-server list-vpc2-info <string>` — ListVPC2Info(string) -> []VPC2Info
- `vultr bare-metal-server mass-halt <v1,v2,...>` — MassHalt([]string)
- `vultr bare-metal-server mass-reboot <v1,v2,...>` — MassReboot([]string)
- `vultr bare-metal-server mass-start <v1,v2,...>` — MassStart([]string)
- `vultr bare-metal-server reboot <string>` — Reboot(string)
- `vultr bare-metal-server reinstall <string>` — Reinstall(string) -> *BareMetalServer
- `vultr bare-metal-server start <string>` — Start(string)
- `vultr bare-metal-server update <string> <json:BareMetalUpdate>` `[--schema]` — Update(string, *BareMetalUpdate) -> *BareMetalServer

## `vultr billing`

Billing API (5 operations)

- `vultr billing get-invoice <string>` — GetInvoice(string) -> *Invoice
- `vultr billing list-history` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListHistory(*ListOptions) -> []History, *Meta
- `vultr billing list-invoice-items <int>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListInvoiceItems(int, *ListOptions) -> []InvoiceItem, *Meta
- `vultr billing list-invoices` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListInvoices(*ListOptions) -> []Invoice, *Meta
- `vultr billing list-pending-charges` `[--cursor --description --label --main-ip --per-page --region --tag]` — ListPendingCharges(*ListOptions) -> []InvoiceItem

## `vultr block-storage`

BlockStorage API (7 operations)

- `vultr block-storage attach <string> <json:BlockStorageAttach>` `[--schema]` — Attach(string, *BlockStorageAttach)
- `vultr block-storage create <json:BlockStorageCreate>` `[--schema]` — Create(*BlockStorageCreate) -> *BlockStorage
- `vultr block-storage delete <string>` — Delete(string)
- `vultr block-storage detach <string> <json:BlockStorageDetach>` `[--schema]` — Detach(string, *BlockStorageDetach)
- `vultr block-storage get <string>` — Get(string) -> *BlockStorage
- `vultr block-storage list` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — List(*ListOptions) -> []BlockStorage, *Meta
- `vultr block-storage update <string> <json:BlockStorageUpdate>` `[--schema]` — Update(string, *BlockStorageUpdate)

## `vultr cdn`

CDN API (15 operations)

- `vultr cdn create-pull-zone <json:CDNZoneReq>` `[--schema]` — CreatePullZone(*CDNZoneReq) -> *CDNZone
- `vultr cdn create-push-zone <json:CDNZoneReq>` `[--schema]` — CreatePushZone(*CDNZoneReq) -> *CDNZone
- `vultr cdn create-push-zone-file-endpoint <string> <json:CDNZoneEndpointReq>` `[--schema]` — CreatePushZoneFileEndpoint(string, *CDNZoneEndpointReq) -> *CDNZoneEndpoint
- `vultr cdn delete-pull-zone <string>` — DeletePullZone(string)
- `vultr cdn delete-push-zone <string>` — DeletePushZone(string)
- `vultr cdn delete-push-zone-file <string> <string>` — DeletePushZoneFile(string, string)
- `vultr cdn get-pull-zone <string>` — GetPullZone(string) -> *CDNZone
- `vultr cdn get-push-zone <string>` — GetPushZone(string) -> *CDNZone
- `vultr cdn get-push-zone-file <string> <string>` — GetPushZoneFile(string, string) -> *CDNZoneFile
- `vultr cdn list-pull-zones` — ListPullZones() -> []CDNZone, *Meta
- `vultr cdn list-push-zone-files <string>` — ListPushZoneFiles(string) -> *CDNZoneFileData
- `vultr cdn list-push-zones` — ListPushZones() -> []CDNZone, *Meta
- `vultr cdn purge-pull-zone <string>` — PurgePullZone(string)
- `vultr cdn update-pull-zone <string> <json:CDNZoneReq>` `[--schema]` — UpdatePullZone(string, *CDNZoneReq) -> *CDNZone
- `vultr cdn update-push-zone <string> <json:CDNZoneReq>` `[--schema]` — UpdatePushZone(string, *CDNZoneReq) -> *CDNZone

## `vultr container-registry`

ContainerRegistry API (30 operations)

- `vultr container-registry create <json:ContainerRegistryReq>` `[--schema]` — Create(*ContainerRegistryReq) -> *ContainerRegistry
- `vultr container-registry create-docker-credentials <string> <json:DockerCredentialsOpt>` `[--schema]` — CreateDockerCredentials(string, *DockerCredentialsOpt) -> *ContainerRegistryDockerCredentials
- `vultr container-registry create-replication <string> <string>` — CreateReplication(string, string) -> *ContainerRegistryReplication
- `vultr container-registry create-retention-rule <string> <json:ContainerRegistryRetentionRuleReq>` `[--schema]` — CreateRetentionRule(string, *ContainerRegistryRetentionRuleReq) -> *ContainerRegistryRetentionRule
- `vultr container-registry delete <string>` — Delete(string)
- `vultr container-registry delete-artifact <string> <string> <string>` — DeleteArtifact(string, string, string)
- `vultr container-registry delete-replication <string> <string>` — DeleteReplication(string, string)
- `vultr container-registry delete-repository <string> <string>` — DeleteRepository(string, string)
- `vultr container-registry delete-retention-rule <string> <int>` — DeleteRetentionRule(string, int)
- `vultr container-registry delete-robot <string> <string>` — DeleteRobot(string, string)
- `vultr container-registry execute-retention <string> <bool>` — ExecuteRetention(string, bool) -> *ContainerRegistryRetentionExecution
- `vultr container-registry get <string>` — Get(string) -> *ContainerRegistry
- `vultr container-registry get-artifact <string> <string> <string>` — GetArtifact(string, string, string) -> *ContainerRegistryArtifact
- `vultr container-registry get-replication <string> <string>` — GetReplication(string, string) -> *ContainerRegistryReplication
- `vultr container-registry get-repository <string> <string>` — GetRepository(string, string) -> *ContainerRegistryRepo
- `vultr container-registry get-retention-rule <string> <int>` — GetRetentionRule(string, int) -> *ContainerRegistryRetentionRule
- `vultr container-registry get-robot <string> <string>` — GetRobot(string, string) -> *ContainerRegistryRobot
- `vultr container-registry list` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — List(*ListOptions) -> []ContainerRegistry, *Meta
- `vultr container-registry list-artifacts <string> <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListArtifacts(string, string, *ListOptions) -> []ContainerRegistryArtifact, *Meta
- `vultr container-registry list-plans` — ListPlans() -> *ContainerRegistryPlans
- `vultr container-registry list-regions` — ListRegions() -> []ContainerRegistryRegion, *Meta
- `vultr container-registry list-replications <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListReplications(string, *ListOptions) -> []ContainerRegistryReplication, *Meta
- `vultr container-registry list-repositories <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListRepositories(string, *ListOptions) -> []ContainerRegistryRepo, *Meta
- `vultr container-registry list-retention-rules <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListRetentionRules(string, *ListOptions) -> []ContainerRegistryRetentionRule, *Meta
- `vultr container-registry list-robots <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListRobots(string, *ListOptions) -> []ContainerRegistryRobot, *Meta
- `vultr container-registry update <string> <json:ContainerRegistryUpdateReq>` `[--schema]` — Update(string, *ContainerRegistryUpdateReq) -> *ContainerRegistry
- `vultr container-registry update-repository <string> <string> <json:ContainerRegistryRepoUpdateReq>` `[--schema]` — UpdateRepository(string, string, *ContainerRegistryRepoUpdateReq) -> *ContainerRegistryRepo
- `vultr container-registry update-retention-rule <string> <int> <bool>` — UpdateRetentionRule(string, int, bool) -> *ContainerRegistryRetentionRule
- `vultr container-registry update-retention-schedule <string> <string>` — UpdateRetentionSchedule(string, string) -> *ContainerRegistryRetentionSchedule
- `vultr container-registry update-robot <string> <string> <json:ContainerRegistryRobotReq>` `[--schema]` — UpdateRobot(string, string, *ContainerRegistryRobotReq) -> *ContainerRegistryRobot

## `vultr database`

Database API (65 operations)

- `vultr database add-read-only-replica <string> <json:DatabaseAddReplicaReq>` `[--schema]` — AddReadOnlyReplica(string, *DatabaseAddReplicaReq) -> *Database
- `vultr database create <json:DatabaseCreateReq>` `[--schema]` — Create(*DatabaseCreateReq) -> *Database
- `vultr database create-connection-pool <string> <json:DatabaseConnectionPoolCreateReq>` `[--schema]` — CreateConnectionPool(string, *DatabaseConnectionPoolCreateReq) -> *DatabaseConnectionPool
- `vultr database create-connector <string> <json:DatabaseConnectorCreateReq>` `[--schema]` — CreateConnector(string, *DatabaseConnectorCreateReq) -> *DatabaseConnector
- `vultr database create-db <string> <json:DatabaseDBCreateReq>` `[--schema]` — CreateDB(string, *DatabaseDBCreateReq) -> *DatabaseDB
- `vultr database create-quota <string> <json:DatabaseQuotaCreateReq>` `[--schema]` — CreateQuota(string, *DatabaseQuotaCreateReq) -> *DatabaseQuota
- `vultr database create-topic <string> <json:DatabaseTopicCreateReq>` `[--schema]` — CreateTopic(string, *DatabaseTopicCreateReq) -> *DatabaseTopic
- `vultr database create-user <string> <json:DatabaseUserCreateReq>` `[--schema]` — CreateUser(string, *DatabaseUserCreateReq) -> *DatabaseUser
- `vultr database delete <string>` — Delete(string)
- `vultr database delete-connection-pool <string> <string>` — DeleteConnectionPool(string, string)
- `vultr database delete-connector <string> <string>` — DeleteConnector(string, string)
- `vultr database delete-db <string> <string>` — DeleteDB(string, string)
- `vultr database delete-quota <string> <string> <string>` — DeleteQuota(string, string, string)
- `vultr database delete-topic <string> <string>` — DeleteTopic(string, string)
- `vultr database delete-user <string> <string>` — DeleteUser(string, string)
- `vultr database detach-migration <string>` — DetachMigration(string)
- `vultr database fork <string> <json:DatabaseForkReq>` `[--schema]` — Fork(string, *DatabaseForkReq) -> *Database
- `vultr database get <string>` — Get(string) -> *Database
- `vultr database get-backup-information <string>` — GetBackupInformation(string) -> *DatabaseBackups
- `vultr database get-connection-pool <string> <string>` — GetConnectionPool(string, string) -> *DatabaseConnectionPool
- `vultr database get-connector <string> <string>` — GetConnector(string, string) -> *DatabaseConnector
- `vultr database get-connector-configuration-schema <string> <string>` — GetConnectorConfigurationSchema(string, string) -> []DatabaseConnectorConfigurationOption
- `vultr database get-connector-status <string> <string>` — GetConnectorStatus(string, string) -> *DatabaseConnectorStatus
- `vultr database get-db <string> <string>` — GetDB(string, string) -> *DatabaseDB
- `vultr database get-migration-status <string>` — GetMigrationStatus(string) -> *DatabaseMigration
- `vultr database get-quota <string> <string> <string>` — GetQuota(string, string, string) -> *DatabaseQuota
- `vultr database get-topic <string> <string>` — GetTopic(string, string) -> *DatabaseTopic
- `vultr database get-usage <string>` — GetUsage(string) -> *DatabaseUsage
- `vultr database get-user <string> <string>` — GetUser(string, string) -> *DatabaseUser
- `vultr database list <json:DBListOptions>` `[--schema]` — List(*DBListOptions) -> []Database, *Meta
- `vultr database list-advanced-options <string>` — ListAdvancedOptions(string) -> *DatabaseAdvancedOptions, []AvailableOption
- `vultr database list-available-connectors <string>` — ListAvailableConnectors(string) -> []DatabaseAvailableConnector
- `vultr database list-available-versions <string>` — ListAvailableVersions(string) -> []string
- `vultr database list-connection-pools <string>` — ListConnectionPools(string) -> *DatabaseConnections, []DatabaseConnectionPool, *Meta
- `vultr database list-connectors <string>` — ListConnectors(string) -> []DatabaseConnector, *Meta
- `vultr database list-d-bs <string>` — ListDBs(string) -> []DatabaseDB, *Meta
- `vultr database list-kafka-connect-advanced-options <string>` — ListKafkaConnectAdvancedOptions(string) -> *DatabaseKafkaConnectAdvancedOptions, []AvailableOption
- `vultr database list-kafka-rest-advanced-options <string>` — ListKafkaRESTAdvancedOptions(string) -> *DatabaseKafkaRESTAdvancedOptions, []AvailableOption
- `vultr database list-maintenance-updates <string>` — ListMaintenanceUpdates(string) -> []string
- `vultr database list-plans <json:DBPlanListOptions>` `[--schema]` — ListPlans(*DBPlanListOptions) -> []DatabasePlan, *Meta
- `vultr database list-quotas <string>` — ListQuotas(string) -> []DatabaseQuota, *Meta
- `vultr database list-schema-registry-advanced-options <string>` — ListSchemaRegistryAdvancedOptions(string) -> *DatabaseSchemaRegistryAdvancedOptions, []AvailableOption
- `vultr database list-service-alerts <string> <json:DatabaseListAlertsReq>` `[--schema]` — ListServiceAlerts(string, *DatabaseListAlertsReq) -> []DatabaseAlert
- `vultr database list-topics <string>` — ListTopics(string) -> []DatabaseTopic, *Meta
- `vultr database list-users <string>` — ListUsers(string) -> []DatabaseUser, *Meta
- `vultr database pause-connector <string> <string>` — PauseConnector(string, string)
- `vultr database promote-read-replica <string>` — PromoteReadReplica(string)
- `vultr database restart-connector <string> <string>` — RestartConnector(string, string)
- `vultr database restart-connector-task <string> <string> <int>` — RestartConnectorTask(string, string, int)
- `vultr database restore-from-backup <string> <json:DatabaseBackupRestoreReq>` `[--schema]` — RestoreFromBackup(string, *DatabaseBackupRestoreReq) -> *Database
- `vultr database resume-connector <string> <string>` — ResumeConnector(string, string)
- `vultr database start-maintenance <string>` — StartMaintenance(string) -> string
- `vultr database start-migration <string> <json:DatabaseMigrationStartReq>` `[--schema]` — StartMigration(string, *DatabaseMigrationStartReq) -> *DatabaseMigration
- `vultr database start-version-upgrade <string> <json:DatabaseVersionUpgradeReq>` `[--schema]` — StartVersionUpgrade(string, *DatabaseVersionUpgradeReq) -> string
- `vultr database update <string> <json:DatabaseUpdateReq>` `[--schema]` — Update(string, *DatabaseUpdateReq) -> *Database
- `vultr database update-advanced-options <string> <json:DatabaseAdvancedOptions>` `[--schema]` — UpdateAdvancedOptions(string, *DatabaseAdvancedOptions) -> *DatabaseAdvancedOptions, []AvailableOption
- `vultr database update-connection-pool <string> <string> <json:DatabaseConnectionPoolUpdateReq>` `[--schema]` — UpdateConnectionPool(string, string, *DatabaseConnectionPoolUpdateReq) -> *DatabaseConnectionPool
- `vultr database update-connector <string> <string> <json:DatabaseConnectorUpdateReq>` `[--schema]` — UpdateConnector(string, string, *DatabaseConnectorUpdateReq) -> *DatabaseConnector
- `vultr database update-kafka-connect-advanced-options <string> <json:DatabaseKafkaConnectAdvancedOptions>` `[--schema]` — UpdateKafkaConnectAdvancedOptions(string, *DatabaseKafkaConnectAdvancedOptions) -> *DatabaseKafkaConnectAdvancedOptions, []AvailableOption
- `vultr database update-kafka-rest-advanced-options <string> <json:DatabaseKafkaRESTAdvancedOptions>` `[--schema]` — UpdateKafkaRESTAdvancedOptions(string, *DatabaseKafkaRESTAdvancedOptions) -> *DatabaseKafkaRESTAdvancedOptions, []AvailableOption
- `vultr database update-quota <string> <string> <string> <json:DatabaseQuotaUpdateReq>` `[--schema]` — UpdateQuota(string, string, string, *DatabaseQuotaUpdateReq) -> *DatabaseQuota
- `vultr database update-schema-registry-advanced-options <string> <json:DatabaseSchemaRegistryAdvancedOptions>` `[--schema]` — UpdateSchemaRegistryAdvancedOptions(string, *DatabaseSchemaRegistryAdvancedOptions) -> *DatabaseSchemaRegistryAdvancedOptions, []AvailableOption
- `vultr database update-topic <string> <string> <json:DatabaseTopicUpdateReq>` `[--schema]` — UpdateTopic(string, string, *DatabaseTopicUpdateReq) -> *DatabaseTopic
- `vultr database update-user <string> <string> <json:DatabaseUserUpdateReq>` `[--schema]` — UpdateUser(string, string, *DatabaseUserUpdateReq) -> *DatabaseUser
- `vultr database update-user-acl <string> <string> <json:DatabaseUserACLReq>` `[--schema]` — UpdateUserACL(string, string, *DatabaseUserACLReq) -> *DatabaseUser

## `vultr domain`

Domain API (8 operations)

- `vultr domain create <json:DomainReq>` `[--schema]` — Create(*DomainReq) -> *Domain
- `vultr domain delete <string>` — Delete(string)
- `vultr domain get <string>` — Get(string) -> *Domain
- `vultr domain get-dns-sec <string>` — GetDNSSec(string) -> []string
- `vultr domain get-soa <string>` — GetSoa(string) -> *Soa
- `vultr domain list` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — List(*ListOptions) -> []Domain, *Meta
- `vultr domain update <string> <string>` — Update(string, string)
- `vultr domain update-soa <string> <json:Soa>` `[--schema]` — UpdateSoa(string, *Soa)

## `vultr domain-record`

DomainRecord API (5 operations)

- `vultr domain-record create <string> <json:DomainRecordCreateReq>` `[--schema]` — Create(string, *DomainRecordCreateReq) -> *DomainRecord
- `vultr domain-record delete <string> <string>` — Delete(string, string)
- `vultr domain-record get <string> <string>` — Get(string, string) -> *DomainRecord
- `vultr domain-record list <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — List(string, *ListOptions) -> []DomainRecord, *Meta
- `vultr domain-record update <string> <string> <json:DomainRecordUpdateReq>` `[--schema]` — Update(string, string, *DomainRecordUpdateReq)

## `vultr firewall-group`

FirewallGroup API (5 operations)

- `vultr firewall-group create <json:FirewallGroupReq>` `[--schema]` — Create(*FirewallGroupReq) -> *FirewallGroup
- `vultr firewall-group delete <string>` — Delete(string)
- `vultr firewall-group get <string>` — Get(string) -> *FirewallGroup
- `vultr firewall-group list` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — List(*ListOptions) -> []FirewallGroup, *Meta
- `vultr firewall-group update <string> <json:FirewallGroupReq>` `[--schema]` — Update(string, *FirewallGroupReq)

## `vultr firewall-rule`

FirewallRule API (4 operations)

- `vultr firewall-rule create <string> <json:FirewallRuleReq>` `[--schema]` — Create(string, *FirewallRuleReq) -> *FirewallRule
- `vultr firewall-rule delete <string> <int>` — Delete(string, int)
- `vultr firewall-rule get <string> <int>` — Get(string, int) -> *FirewallRule
- `vultr firewall-rule list <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — List(string, *ListOptions) -> []FirewallRule, *Meta

## `vultr inference`

Inference API (6 operations)

- `vultr inference create <json:InferenceCreateUpdateReq>` `[--schema]` — Create(*InferenceCreateUpdateReq) -> *Inference
- `vultr inference delete <string>` — Delete(string)
- `vultr inference get <string>` — Get(string) -> *Inference
- `vultr inference get-usage <string>` — GetUsage(string) -> *InferenceUsage
- `vultr inference list` — List() -> []Inference
- `vultr inference update <string> <json:InferenceCreateUpdateReq>` `[--schema]` — Update(string, *InferenceCreateUpdateReq) -> *Inference

## `vultr instance`

Instance API (37 operations)

- `vultr instance attach-iso <string> <string>` — AttachISO(string, string)
- `vultr instance attach-vpc <string> <string>` — AttachVPC(string, string)
- `vultr instance attach-vpc2 <string> <json:AttachVPC2Req>` `[--schema]` — AttachVPC2(string, *AttachVPC2Req)
- `vultr instance create <json:InstanceCreateReq>` `[--schema]` — Create(*InstanceCreateReq) -> *Instance
- `vultr instance create-ipv4 <string> [<bool|null>]` — CreateIPv4(string, *bool) -> *IPv4
- `vultr instance create-reverse-ipv4 <string> <json:ReverseIP>` `[--schema]` — CreateReverseIPv4(string, *ReverseIP)
- `vultr instance create-reverse-ipv6 <string> <json:ReverseIP>` `[--schema]` — CreateReverseIPv6(string, *ReverseIP)
- `vultr instance default-reverse-ipv4 <string> <string>` — DefaultReverseIPv4(string, string)
- `vultr instance delete <string>` — Delete(string)
- `vultr instance delete-ipv4 <string> <string>` — DeleteIPv4(string, string)
- `vultr instance delete-reverse-ipv6 <string> <string>` — DeleteReverseIPv6(string, string)
- `vultr instance detach-iso <string>` — DetachISO(string)
- `vultr instance detach-vpc <string> <string>` — DetachVPC(string, string)
- `vultr instance detach-vpc2 <string> <string>` — DetachVPC2(string, string)
- `vultr instance get <string>` — Get(string) -> *Instance
- `vultr instance get-backup-schedule <string>` — GetBackupSchedule(string) -> *BackupSchedule
- `vultr instance get-bandwidth <string>` — GetBandwidth(string) -> *Bandwidth
- `vultr instance get-neighbors <string>` — GetNeighbors(string) -> *Neighbors
- `vultr instance get-upgrades <string>` — GetUpgrades(string) -> *Upgrades
- `vultr instance get-user-data <string>` — GetUserData(string) -> *UserData
- `vultr instance halt <string>` — Halt(string)
- `vultr instance iso-status <string>` — ISOStatus(string) -> *Iso
- `vultr instance list` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — List(*ListOptions) -> []Instance, *Meta
- `vultr instance list-ipv4 <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListIPv4(string, *ListOptions) -> []IPv4, *Meta
- `vultr instance list-ipv6 <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListIPv6(string, *ListOptions) -> []IPv6, *Meta
- `vultr instance list-reverse-ipv6 <string>` — ListReverseIPv6(string) -> []ReverseIP
- `vultr instance list-vpc-info <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListVPCInfo(string, *ListOptions) -> []VPCInfo, *Meta
- `vultr instance list-vpc2-info <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListVPC2Info(string, *ListOptions) -> []VPC2Info, *Meta
- `vultr instance mass-halt <v1,v2,...>` — MassHalt([]string)
- `vultr instance mass-reboot <v1,v2,...>` — MassReboot([]string)
- `vultr instance mass-start <v1,v2,...>` — MassStart([]string)
- `vultr instance reboot <string>` — Reboot(string)
- `vultr instance reinstall <string> <json:ReinstallReq>` `[--schema]` — Reinstall(string, *ReinstallReq) -> *Instance
- `vultr instance restore <string> <json:RestoreReq>` `[--schema]` — Restore(string, *RestoreReq)
- `vultr instance set-backup-schedule <string> <json:BackupScheduleReq>` `[--schema]` — SetBackupSchedule(string, *BackupScheduleReq)
- `vultr instance start <string>` — Start(string)
- `vultr instance update <string> <json:InstanceUpdateReq>` `[--schema]` — Update(string, *InstanceUpdateReq) -> *Instance

## `vultr iso`

ISO API (5 operations)

- `vultr iso create <json:ISOReq>` `[--schema]` — Create(*ISOReq) -> *ISO
- `vultr iso delete <string>` — Delete(string)
- `vultr iso get <string>` — Get(string) -> *ISO
- `vultr iso list` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — List(*ListOptions) -> []ISO, *Meta
- `vultr iso list-public` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListPublic(*ListOptions) -> []PublicISO, *Meta

## `vultr kubernetes`

Kubernetes API (26 operations)

- `vultr kubernetes create-cluster <json:ClusterReq>` `[--schema]` — CreateCluster(*ClusterReq) -> *Cluster
- `vultr kubernetes create-node-pool <string> <json:NodePoolReq>` `[--schema]` — CreateNodePool(string, *NodePoolReq) -> *NodePool
- `vultr kubernetes create-node-pool-label <string> <string> <json:NodePoolLabelReq>` `[--schema]` — CreateNodePoolLabel(string, string, *NodePoolLabelReq) -> *NodePoolLabel
- `vultr kubernetes create-node-pool-taint <string> <string> <json:NodePoolTaintReq>` `[--schema]` — CreateNodePoolTaint(string, string, *NodePoolTaintReq) -> *NodePoolTaint
- `vultr kubernetes delete-cluster <string>` — DeleteCluster(string)
- `vultr kubernetes delete-cluster-with-resources <string>` — DeleteClusterWithResources(string)
- `vultr kubernetes delete-node-pool <string> <string>` — DeleteNodePool(string, string)
- `vultr kubernetes delete-node-pool-instance <string> <string> <string>` — DeleteNodePoolInstance(string, string, string)
- `vultr kubernetes delete-node-pool-label <string> <string> <string>` — DeleteNodePoolLabel(string, string, string)
- `vultr kubernetes delete-node-pool-taint <string> <string> <string>` — DeleteNodePoolTaint(string, string, string)
- `vultr kubernetes get-cluster <string>` — GetCluster(string) -> *Cluster
- `vultr kubernetes get-kube-config <string>` — GetKubeConfig(string) -> *KubeConfig
- `vultr kubernetes get-node-pool <string> <string>` — GetNodePool(string, string) -> *NodePool
- `vultr kubernetes get-node-pool-label <string> <string> <string>` — GetNodePoolLabel(string, string, string) -> *NodePoolLabel
- `vultr kubernetes get-node-pool-taint <string> <string> <string>` — GetNodePoolTaint(string, string, string) -> *NodePoolTaint
- `vultr kubernetes get-upgrades <string>` — GetUpgrades(string) -> []string
- `vultr kubernetes get-versions` — GetVersions() -> *Versions
- `vultr kubernetes list-clusters` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListClusters(*ListOptions) -> []Cluster, *Meta
- `vultr kubernetes list-node-pool-labels <string> <string>` — ListNodePoolLabels(string, string) -> []NodePoolLabel
- `vultr kubernetes list-node-pool-taints <string> <string>` — ListNodePoolTaints(string, string) -> []NodePoolTaint
- `vultr kubernetes list-node-pools <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListNodePools(string, *ListOptions) -> []NodePool, *Meta
- `vultr kubernetes list-worker-nodes <string> <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListWorkerNodes(string, string, *ListOptions) -> []Node, *Meta
- `vultr kubernetes recycle-node-pool-instance <string> <string> <string>` — RecycleNodePoolInstance(string, string, string)
- `vultr kubernetes update-cluster <string> <json:ClusterReqUpdate>` `[--schema]` — UpdateCluster(string, *ClusterReqUpdate)
- `vultr kubernetes update-node-pool <string> <string> <json:NodePoolReqUpdate>` `[--schema]` — UpdateNodePool(string, string, *NodePoolReqUpdate) -> *NodePool
- `vultr kubernetes upgrade <string> <json:ClusterUpgradeReq>` `[--schema]` — Upgrade(string, *ClusterUpgradeReq)

## `vultr load-balancer`

LoadBalancer API (15 operations)

- `vultr load-balancer create <json:LoadBalancerReq>` `[--schema]` — Create(*LoadBalancerReq) -> *LoadBalancer
- `vultr load-balancer create-firewall-rules <string> <json:[]LBFirewallRule>` `[--schema]` — CreateFirewallRules(string, []LBFirewallRule)
- `vultr load-balancer create-forwarding-rule <string> <json:ForwardingRule>` `[--schema]` — CreateForwardingRule(string, *ForwardingRule) -> *ForwardingRule
- `vultr load-balancer delete <string>` — Delete(string)
- `vultr load-balancer delete-auto-ssl <string>` — DeleteAutoSSL(string)
- `vultr load-balancer delete-firewall-rule <string> <string>` — DeleteFirewallRule(string, string)
- `vultr load-balancer delete-forwarding-rule <string> <string>` — DeleteForwardingRule(string, string)
- `vultr load-balancer delete-ssl <string>` — DeleteSSL(string)
- `vultr load-balancer get <string>` — Get(string) -> *LoadBalancer
- `vultr load-balancer get-firewall-rule <string> <string>` — GetFirewallRule(string, string) -> *LBFirewallRule
- `vultr load-balancer get-forwarding-rule <string> <string>` — GetForwardingRule(string, string) -> *ForwardingRule
- `vultr load-balancer list` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — List(*ListOptions) -> []LoadBalancer, *Meta
- `vultr load-balancer list-firewall-rules <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListFirewallRules(string, *ListOptions) -> []LBFirewallRule, *Meta
- `vultr load-balancer list-forwarding-rules <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListForwardingRules(string, *ListOptions) -> []ForwardingRule, *Meta
- `vultr load-balancer update <string> <json:LoadBalancerReq>` `[--schema]` — Update(string, *LoadBalancerReq)

## `vultr logs`

Logs API (1 operations)

- `vultr logs list <json:LogsOptions>` `[--schema]` — List(LogsOptions) -> []Log, *LogsMeta

## `vultr marketplace`

Marketplace API (1 operations)

- `vultr marketplace list-app-variables <string>` — ListAppVariables(string) -> []MarketplaceAppVariable

## `vultr object-storage`

ObjectStorage API (12 operations)

- `vultr object-storage create <json:ObjectStorageReq>` `[--schema]` — Create(*ObjectStorageReq) -> *ObjectStorage
- `vultr object-storage create-bucket <string> <json:ObjectStorageBucketReq>` `[--schema]` — CreateBucket(string, *ObjectStorageBucketReq)
- `vultr object-storage delete <string>` — Delete(string)
- `vultr object-storage delete-bucket <string> <string>` — DeleteBucket(string, string)
- `vultr object-storage get <string>` — Get(string) -> *ObjectStorage
- `vultr object-storage list` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — List(*ListOptions) -> []ObjectStorage, *Meta
- `vultr object-storage list-buckets <string>` — ListBuckets(string) -> []ObjectStorageBucket
- `vultr object-storage list-cluster` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListCluster(*ListOptions) -> []ObjectStorageCluster, *Meta
- `vultr object-storage list-cluster-tiers <int>` — ListClusterTiers(int) -> []ObjectStorageTier
- `vultr object-storage list-tiers` — ListTiers() -> []ObjectStorageTier
- `vultr object-storage regenerate-keys <string>` — RegenerateKeys(string) -> *S3Keys
- `vultr object-storage update <string> <json:ObjectStorageReq>` `[--schema]` — Update(string, *ObjectStorageReq)

## `vultr oidc`

OIDC API (12 operations)

- `vultr oidc authorize-oidc <json:OIDCAuthorizeParameters>` `[--schema]` — AuthorizeOIDC(*OIDCAuthorizeParameters)
- `vultr oidc create-oidc-issuer <json:OIDCIssuerReq>` `[--schema]` — CreateOIDCIssuer(*OIDCIssuerReq) -> *OIDCIssuer
- `vultr oidc create-oidc-provider <json:OIDCProviderReq>` `[--schema]` — CreateOIDCProvider(*OIDCProviderReq) -> *OIDCProvider
- `vultr oidc create-oidc-token <json:OIDCTokenReq>` `[--schema]` — CreateOIDCToken(*OIDCTokenReq) -> *OIDCToken
- `vultr oidc delete-oidc-issuer <string>` — DeleteOIDCIssuer(string)
- `vultr oidc delete-oidc-provider <string>` — DeleteOIDCProvider(string)
- `vultr oidc discovery-oidc <string>` — DiscoveryOIDC(string) -> *OIDCDocument
- `vultr oidc get-oidc-issuer <string>` — GetOIDCIssuer(string) -> *OIDCIssuer
- `vultr oidc get-oidc-provider <string>` — GetOIDCProvider(string) -> *OIDCProvider
- `vultr oidc get-oidc-token <string>` — GetOIDCToken(string) -> *OIDCToken
- `vultr oidc list-oidc-issuers` — ListOIDCIssuers() -> []OIDCIssuer
- `vultr oidc list-oidc-providers` — ListOIDCProviders() -> []OIDCProvider

## `vultr organization`

Organization API (68 operations)

- `vultr organization add-group-member <string> <json:OrganizationGroupMemberReq>` `[--schema]` — AddGroupMember(string, *OrganizationGroupMemberReq)
- `vultr organization attach-policy-group <string> <string>` — AttachPolicyGroup(string, string)
- `vultr organization attach-policy-user <string> <string>` — AttachPolicyUser(string, string)
- `vultr organization attach-role-group <string> <string>` — AttachRoleGroup(string, string) -> *OrganizationRoleGroupAssignment
- `vultr organization attach-role-policy <string> <string>` — AttachRolePolicy(string, string) -> *OrganizationRolePolicyAttachment
- `vultr organization attach-role-user <string> <string>` — AttachRoleUser(string, string) -> *OrganizationRoleUserAssignment
- `vultr organization create-group <json:OrganizationGroupReq>` `[--schema]` — CreateGroup(*OrganizationGroupReq) -> *OrganizationGroup
- `vultr organization create-invitation <json:OrganizationInvitationReq>` `[--schema]` — CreateInvitation(*OrganizationInvitationReq) -> *OrganizationInvitation
- `vultr organization create-organization <json:OrganizationReq>` `[--schema]` — CreateOrganization(*OrganizationReq) -> *Organization
- `vultr organization create-policy <json:OrganizationPolicyReq>` `[--schema]` — CreatePolicy(*OrganizationPolicyReq) -> *OrganizationPolicy
- `vultr organization create-role <json:OrganizationRoleReq>` `[--schema]` — CreateRole(*OrganizationRoleReq) -> *OrganizationRole
- `vultr organization create-role-session <json:OrganizationRoleSessionReq>` `[--schema]` — CreateRoleSession(*OrganizationRoleSessionReq) -> *OrganizationRoleSession
- `vultr organization create-role-trust <json:OrganizationRoleTrustCreateReq>` `[--schema]` — CreateRoleTrust(*OrganizationRoleTrustCreateReq) -> *OrganizationRoleTrust
- `vultr organization delete-group <string>` — DeleteGroup(string)
- `vultr organization delete-organization <string>` — DeleteOrganization(string)
- `vultr organization delete-policy <string>` — DeletePolicy(string)
- `vultr organization delete-role <string>` — DeleteRole(string)
- `vultr organization delete-role-trust <string>` — DeleteRoleTrust(string)
- `vultr organization delete-user <string> <string>` — DeleteUser(string, string)
- `vultr organization detach-policy-group <string> <string>` — DetachPolicyGroup(string, string)
- `vultr organization detach-policy-user <string> <string>` — DetachPolicyUser(string, string)
- `vultr organization detach-role-group <string> <string>` — DetachRoleGroup(string, string)
- `vultr organization detach-role-policy <string> <string>` — DetachRolePolicy(string, string)
- `vultr organization detach-role-user <string> <string>` — DetachRoleUser(string, string)
- `vultr organization get-group <string>` — GetGroup(string) -> *OrganizationGroup
- `vultr organization get-invitation <string>` — GetInvitation(string) -> *OrganizationInvitation
- `vultr organization get-organization <string>` — GetOrganization(string) -> *Organization
- `vultr organization get-policy <string>` — GetPolicy(string) -> *OrganizationPolicy
- `vultr organization get-role <string>` — GetRole(string) -> *OrganizationRole
- `vultr organization get-role-session <string>` — GetRoleSession(string) -> *OrganizationRoleSession
- `vultr organization get-role-trust <string>` — GetRoleTrust(string) -> *OrganizationRoleTrust
- `vultr organization list-current-user-groups` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListCurrentUserGroups(*ListOptions) -> []OrganizationGroup, *Meta
- `vultr organization list-current-user-policies` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListCurrentUserPolicies(*ListOptions) -> *OrganizationPoliciesForUser, *Meta
- `vultr organization list-current-user-roles` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListCurrentUserRoles(*ListOptions) -> *OrganizationRolesForUser, *Meta
- `vultr organization list-group-policies <string>` — ListGroupPolicies(string) -> *OrganizationGroupPolicies, *Meta
- `vultr organization list-group-roles <string>` — ListGroupRoles(string) -> *OrganizationGroupRoles, *Meta
- `vultr organization list-groups` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListGroups(*ListOptions) -> []OrganizationGroup, *Meta
- `vultr organization list-invitations` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListInvitations(*ListOptions) -> []OrganizationInvitation, *Meta
- `vultr organization list-organizations` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListOrganizations(*ListOptions) -> []Organization, *Meta
- `vultr organization list-policies` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListPolicies(*ListOptions) -> []OrganizationPolicy, *Meta
- `vultr organization list-policy-groups <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListPolicyGroups(string, *ListOptions) -> []OrganizationGroup, *Meta
- `vultr organization list-policy-users <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListPolicyUsers(string, *ListOptions) -> []OrganizationUser, *Meta
- `vultr organization list-role-groups <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListRoleGroups(string, *ListOptions) -> []OrganizationRoleGroupAssignment, *Meta
- `vultr organization list-role-policies <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListRolePolicies(string, *ListOptions) -> []OrganizationPolicy, *Meta
- `vultr organization list-role-sessions <string>` — ListRoleSessions(string) -> []OrganizationRoleSession, *Meta
- `vultr organization list-role-trusts` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListRoleTrusts(*ListOptions) -> []OrganizationRoleTrust, *Meta
- `vultr organization list-role-trusts-by-role <string>` — ListRoleTrustsByRole(string) -> []OrganizationRoleTrust, *Meta
- `vultr organization list-role-trusts-by-user <string>` — ListRoleTrustsByUser(string) -> []OrganizationRoleTrust, *Meta
- `vultr organization list-role-users <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListRoleUsers(string, *ListOptions) -> []OrganizationRoleUserAssignment, *Meta
- `vultr organization list-roles` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListRoles(*ListOptions) -> []OrganizationRole, *Meta
- `vultr organization list-suspended-users <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListSuspendedUsers(string, *ListOptions) -> []OrganizationUser, *Meta
- `vultr organization list-user-groups <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListUserGroups(string, *ListOptions) -> []OrganizationGroup, *Meta
- `vultr organization list-user-policies <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListUserPolicies(string, *ListOptions) -> *OrganizationPoliciesForUser, *Meta
- `vultr organization list-user-roles <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListUserRoles(string, *ListOptions) -> *OrganizationRolesForUser, *Meta
- `vultr organization remove-group-member <string> <string>` — RemoveGroupMember(string, string)
- `vultr organization resend-invitation <string>` — ResendInvitation(string) -> *OrganizationInvitation
- `vultr organization restore-organization <string>` — RestoreOrganization(string)
- `vultr organization restore-policy <string>` — RestorePolicy(string) -> *OrganizationPolicy
- `vultr organization restore-role <string>` — RestoreRole(string) -> *OrganizationRole
- `vultr organization restore-role-trust <string>` — RestoreRoleTrust(string) -> *OrganizationRoleTrust
- `vultr organization revoke-role-session <string>` — RevokeRoleSession(string)
- `vultr organization suspend-user <string> <string>` — SuspendUser(string, string)
- `vultr organization unsuspend-user <string> <string>` — UnsuspendUser(string, string)
- `vultr organization update-group <string> <json:OrganizationGroupReq>` `[--schema]` — UpdateGroup(string, *OrganizationGroupReq) -> *OrganizationGroup
- `vultr organization update-organization <string> <json:OrganizationReq>` `[--schema]` — UpdateOrganization(string, *OrganizationReq) -> *Organization
- `vultr organization update-policy <string> <json:OrganizationPolicyReq>` `[--schema]` — UpdatePolicy(string, *OrganizationPolicyReq) -> *OrganizationPolicy
- `vultr organization update-role <string> <json:OrganizationRoleReq>` `[--schema]` — UpdateRole(string, *OrganizationRoleReq) -> *OrganizationRole
- `vultr organization update-role-trust <string> <json:OrganizationRoleTrustUpdateReq>` `[--schema]` — UpdateRoleTrust(string, *OrganizationRoleTrustUpdateReq) -> *OrganizationRoleTrust

## `vultr os`

OS API (1 operations)

- `vultr os list` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — List(*ListOptions) -> []OS, *Meta

## `vultr plan`

Plan API (2 operations)

- `vultr plan list <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — List(string, *ListOptions) -> []Plan, *Meta
- `vultr plan list-bare-metal` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListBareMetal(*ListOptions) -> []BareMetalPlan, *Meta

## `vultr region`

Region API (2 operations)

- `vultr region availability <string> <string>` — Availability(string, string) -> *PlanAvailability
- `vultr region list` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — List(*ListOptions) -> []Region, *Meta

## `vultr reserved-ip`

ReservedIP API (8 operations)

- `vultr reserved-ip attach <string> <string>` — Attach(string, string)
- `vultr reserved-ip convert <json:ReservedIPConvertReq>` `[--schema]` — Convert(*ReservedIPConvertReq) -> *ReservedIP
- `vultr reserved-ip create <json:ReservedIPReq>` `[--schema]` — Create(*ReservedIPReq) -> *ReservedIP
- `vultr reserved-ip delete <string>` — Delete(string)
- `vultr reserved-ip detach <string>` — Detach(string)
- `vultr reserved-ip get <string>` — Get(string) -> *ReservedIP
- `vultr reserved-ip list` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — List(*ListOptions) -> []ReservedIP, *Meta
- `vultr reserved-ip update <string> <json:ReservedIPUpdateReq>` `[--schema]` — Update(string, *ReservedIPUpdateReq) -> *ReservedIP

## `vultr snapshot`

Snapshot API (5 operations)

- `vultr snapshot create <json:SnapshotReq>` `[--schema]` — Create(*SnapshotReq) -> *Snapshot
- `vultr snapshot create-from-url <json:SnapshotURLReq>` `[--schema]` — CreateFromURL(*SnapshotURLReq) -> *Snapshot
- `vultr snapshot delete <string>` — Delete(string)
- `vultr snapshot get <string>` — Get(string) -> *Snapshot
- `vultr snapshot list` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — List(*ListOptions) -> []Snapshot, *Meta

## `vultr ssh-key`

SSHKey API (5 operations)

- `vultr ssh-key create <json:SSHKeyReq>` `[--schema]` — Create(*SSHKeyReq) -> *SSHKey
- `vultr ssh-key delete <string>` — Delete(string)
- `vultr ssh-key get <string>` — Get(string) -> *SSHKey
- `vultr ssh-key list` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — List(*ListOptions) -> []SSHKey, *Meta
- `vultr ssh-key update <string> <json:SSHKeyReq>` `[--schema]` — Update(string, *SSHKeyReq)

## `vultr startup-script`

StartupScript API (5 operations)

- `vultr startup-script create <json:StartupScriptReq>` `[--schema]` — Create(*StartupScriptReq) -> *StartupScript
- `vultr startup-script delete <string>` — Delete(string)
- `vultr startup-script get <string>` — Get(string) -> *StartupScript
- `vultr startup-script list` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — List(*ListOptions) -> []StartupScript, *Meta
- `vultr startup-script update <string> <json:StartupScriptReq>` `[--schema]` — Update(string, *StartupScriptReq)

## `vultr sub-account`

SubAccount API (2 operations)

- `vultr sub-account create <json:SubAccountReq>` `[--schema]` — Create(*SubAccountReq) -> *SubAccount
- `vultr sub-account list` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — List(*ListOptions) -> []SubAccount, *Meta

## `vultr user`

User API (5 operations)

- `vultr user create <json:UserReq>` `[--schema]` — Create(*UserReq) -> *User
- `vultr user delete <string>` — Delete(string)
- `vultr user get <string>` — Get(string) -> *User
- `vultr user list` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — List(*ListOptions) -> []User, *Meta
- `vultr user update <string> <json:UserReq>` `[--schema]` — Update(string, *UserReq)

## `vultr virtual-file-system-storage`

VirtualFileSystemStorage API (9 operations)

- `vultr virtual-file-system-storage attach <string> <string>` — Attach(string, string) -> *VirtualFileSystemStorageAttachment
- `vultr virtual-file-system-storage attachment-get <string> <string>` — AttachmentGet(string, string) -> *VirtualFileSystemStorageAttachment
- `vultr virtual-file-system-storage attachment-list <string>` — AttachmentList(string) -> []VirtualFileSystemStorageAttachment
- `vultr virtual-file-system-storage create <json:VirtualFileSystemStorageReq>` `[--schema]` — Create(*VirtualFileSystemStorageReq) -> *VirtualFileSystemStorage
- `vultr virtual-file-system-storage delete <string>` — Delete(string)
- `vultr virtual-file-system-storage detach <string> <string>` — Detach(string, string)
- `vultr virtual-file-system-storage get <string>` — Get(string) -> *VirtualFileSystemStorage
- `vultr virtual-file-system-storage list` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — List(*ListOptions) -> []VirtualFileSystemStorage, *Meta
- `vultr virtual-file-system-storage update <string> <json:VirtualFileSystemStorageUpdateReq>` `[--schema]` — Update(string, *VirtualFileSystemStorageUpdateReq) -> *VirtualFileSystemStorage

## `vultr vpc`

VPC API (20 operations)

- `vultr vpc create <json:VPCReq>` `[--schema]` — Create(*VPCReq) -> *VPC
- `vultr vpc create-nat-gateway <string> <json:NATGatewayReq>` `[--schema]` — CreateNATGateway(string, *NATGatewayReq) -> *NATGateway
- `vultr vpc create-nat-gateway-firewall-rule <string> <string> <json:NATGatewayFirewallRuleCreateReq>` `[--schema]` — CreateNATGatewayFirewallRule(string, string, *NATGatewayFirewallRuleCreateReq) -> *NATGatewayFirewallRule
- `vultr vpc create-nat-gateway-port-forwarding-rule <string> <string> <json:NATGatewayPortForwardingRuleReq>` `[--schema]` — CreateNATGatewayPortForwardingRule(string, string, *NATGatewayPortForwardingRuleReq) -> *NATGatewayPortForwardingRule
- `vultr vpc delete <string>` — Delete(string)
- `vultr vpc delete-nat-gateway <string> <string>` — DeleteNATGateway(string, string)
- `vultr vpc delete-nat-gateway-firewall-rule <string> <string> <string>` — DeleteNATGatewayFirewallRule(string, string, string)
- `vultr vpc delete-nat-gateway-port-forwarding-rule <string> <string> <string>` — DeleteNATGatewayPortForwardingRule(string, string, string)
- `vultr vpc get <string>` — Get(string) -> *VPC
- `vultr vpc get-nat-gateway <string> <string>` — GetNATGateway(string, string) -> *NATGateway
- `vultr vpc get-nat-gateway-firewall-rule <string> <string> <string>` — GetNATGatewayFirewallRule(string, string, string) -> *NATGatewayFirewallRule
- `vultr vpc get-nat-gateway-port-forwarding-rule <string> <string> <string>` — GetNATGatewayPortForwardingRule(string, string, string) -> *NATGatewayPortForwardingRule
- `vultr vpc list` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — List(*ListOptions) -> []VPC, *Meta
- `vultr vpc list-nat-gateway-firewall-rules <string> <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListNATGatewayFirewallRules(string, string, *ListOptions) -> []NATGatewayFirewallRule, *Meta
- `vultr vpc list-nat-gateway-port-forwarding-rules <string> <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListNATGatewayPortForwardingRules(string, string, *ListOptions) -> []NATGatewayPortForwardingRule, *Meta
- `vultr vpc list-nat-gateways <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListNATGateways(string, *ListOptions) -> []NATGateway, *Meta
- `vultr vpc update <string> <string>` — Update(string, string)
- `vultr vpc update-nat-gateway <string> <string> <json:NATGatewayReq>` `[--schema]` — UpdateNATGateway(string, string, *NATGatewayReq) -> *NATGateway
- `vultr vpc update-nat-gateway-firewall-rule <string> <string> <string> <json:NATGatewayFirewallRuleUpdateReq>` `[--schema]` — UpdateNATGatewayFirewallRule(string, string, string, *NATGatewayFirewallRuleUpdateReq) -> *NATGatewayFirewallRule
- `vultr vpc update-nat-gateway-port-forwarding-rule <string> <string> <string> <json:NATGatewayPortForwardingRuleReq>` `[--schema]` — UpdateNATGatewayPortForwardingRule(string, string, string, *NATGatewayPortForwardingRuleReq) -> *NATGatewayPortForwardingRule

## `vultr vpc2`

VPC2 API (8 operations)

- `vultr vpc2 attach <string> <json:VPC2AttachDetachReq>` `[--schema]` — Attach(string, *VPC2AttachDetachReq)
- `vultr vpc2 create <json:VPC2Req>` `[--schema]` — Create(*VPC2Req) -> *VPC2
- `vultr vpc2 delete <string>` — Delete(string)
- `vultr vpc2 detach <string> <json:VPC2AttachDetachReq>` `[--schema]` — Detach(string, *VPC2AttachDetachReq)
- `vultr vpc2 get <string>` — Get(string) -> *VPC2
- `vultr vpc2 list` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — List(*ListOptions) -> []VPC2, *Meta
- `vultr vpc2 list-nodes <string>` `[--all --cursor --description --label --main-ip --per-page --region --tag]` — ListNodes(string, *ListOptions) -> []VPC2Node, *Meta
- `vultr vpc2 update <string> <string>` — Update(string, string)
