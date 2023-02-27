// Code generated by "enumer -type ResourceType -addprefix azurerm_ -transform snake -linecomment"; DO NOT EDIT.

package azurerm

import (
	"fmt"
	"strings"
)

const _ResourceTypeName = "azurerm_resource_groupazurerm_availability_setazurerm_imageazurerm_managed_diskazurerm_virtual_machineazurerm_virtual_machine_data_disk_attachmentazurerm_virtual_machine_extensionazurerm_virtual_machine_scale_set_extensionazurerm_virtual_networkazurerm_linux_virtual_machineazurerm_linux_virtual_machine_scale_setazurerm_windows_virtual_machineazurerm_windows_virtual_machine_scale_setazurerm_subnetazurerm_network_interfaceazurerm_network_security_groupazurerm_application_gatewayazurerm_application_security_groupazurerm_network_ddos_protection_planazurerm_firewallazurerm_local_network_gatewayazurerm_nat_gatewayazurerm_network_profileazurerm_network_security_ruleazurerm_public_ipazurerm_public_ip_prefixazurerm_routeazurerm_route_tableazurerm_virtual_network_gatewayazurerm_virtual_network_gateway_connectionazurerm_virtual_network_peeringazurerm_web_application_firewall_policyazurerm_virtual_hubazurerm_virtual_hub_bgp_connectionazurerm_virtual_hub_connectionazurerm_virtual_hub_ipazurerm_virtual_hub_route_tableazurerm_virtual_hub_security_partner_providerazurerm_lbazurerm_lb_backend_address_poolazurerm_lb_ruleazurerm_lb_outbound_ruleazurerm_lb_nat_ruleazurerm_lb_nat_poolazurerm_lb_probeazurerm_virtual_desktop_host_poolazurerm_virtual_desktop_application_groupazurerm_logic_app_workflowazurerm_logic_app_trigger_customazurerm_logic_app_action_customazurerm_container_registryazurerm_container_registry_webhookazurerm_kubernetes_clusterazurerm_kubernetes_cluster_node_poolazurerm_storage_accountazurerm_storage_queueazurerm_storage_shareazurerm_storage_tableazurerm_storage_blobazurerm_mariadb_configurationazurerm_mariadb_databaseazurerm_mariadb_firewall_ruleazurerm_mariadb_serverazurerm_mariadb_virtual_network_ruleazurerm_mysql_configurationazurerm_mysql_databaseazurerm_mysql_firewall_ruleazurerm_mysql_serverazurerm_mysql_virtual_network_ruleazurerm_postgresql_configurationazurerm_postgresql_databaseazurerm_postgresql_firewall_ruleazurerm_postgresql_serverazurerm_postgresql_virtual_network_ruleazurerm_mssql_elasticpoolazurerm_mssql_databaseazurerm_mssql_firewall_ruleazurerm_mssql_serverazurerm_mssql_server_security_alert_policyazurerm_mssql_server_vulnerability_assessmentazurerm_mssql_virtual_machineazurerm_mssql_virtual_network_ruleazurerm_redis_cacheazurerm_redis_firewall_ruleazurerm_dns_zoneazurerm_dns_a_recordazurerm_dns_aaaa_recordazurerm_dns_caa_recordazurerm_dns_cname_recordazurerm_dns_mx_recordazurerm_dns_ns_recordazurerm_dns_ptr_recordazurerm_dns_srv_recordazurerm_dns_txt_recordazurerm_private_dns_zoneazurerm_private_dns_a_recordazurerm_private_dns_aaaa_recordazurerm_private_dns_cname_recordazurerm_private_dns_mx_recordazurerm_private_dns_ptr_recordazurerm_private_dns_srv_recordazurerm_private_dns_txt_recordazurerm_private_dns_zone_virtual_network_linkazurerm_policy_definitionazurerm_policy_remediationazurerm_policy_set_definitionazurerm_key_vaultazurerm_key_vault_access_policyazurerm_application_insightsazurerm_application_insights_api_keyazurerm_application_insights_analytics_itemazurerm_log_analytics_workspaceazurerm_log_analytics_linked_serviceazurerm_log_analytics_datasource_windows_performance_counterazurerm_log_analytics_datasource_windows_eventazurerm_monitor_action_groupazurerm_monitor_activity_log_alertazurerm_monitor_autoscale_settingazurerm_monitor_log_profileazurerm_monitor_metric_alertazurerm_windows_web_appazurerm_linux_web_appazurerm_linux_web_app_slotazurerm_windows_web_app_slotazurerm_web_app_active_slotazurerm_service_planazurerm_source_control_tokenazurerm_static_siteazurerm_static_site_custom_domainazurerm_web_app_hybrid_connectionazurerm_data_protection_backup_vault"

var _ResourceTypeIndex = [...]uint16{0, 22, 46, 59, 79, 102, 146, 179, 222, 245, 274, 313, 344, 385, 399, 424, 454, 481, 515, 551, 567, 596, 615, 638, 667, 684, 708, 721, 740, 771, 813, 844, 883, 902, 936, 966, 988, 1019, 1064, 1074, 1105, 1120, 1144, 1163, 1182, 1198, 1231, 1272, 1298, 1330, 1361, 1387, 1421, 1447, 1483, 1506, 1527, 1548, 1569, 1589, 1618, 1642, 1671, 1693, 1729, 1756, 1778, 1805, 1825, 1859, 1891, 1918, 1950, 1975, 2014, 2039, 2061, 2088, 2108, 2150, 2195, 2224, 2258, 2277, 2304, 2320, 2340, 2363, 2385, 2409, 2430, 2451, 2473, 2495, 2517, 2541, 2569, 2600, 2632, 2661, 2691, 2721, 2751, 2796, 2821, 2847, 2876, 2893, 2924, 2952, 2988, 3031, 3062, 3098, 3158, 3204, 3232, 3266, 3299, 3326, 3354, 3377, 3398, 3424, 3452, 3479, 3499, 3527, 3546, 3579, 3612, 3648}

const _ResourceTypeLowerName = "azurerm_resource_groupazurerm_availability_setazurerm_imageazurerm_managed_diskazurerm_virtual_machineazurerm_virtual_machine_data_disk_attachmentazurerm_virtual_machine_extensionazurerm_virtual_machine_scale_set_extensionazurerm_virtual_networkazurerm_linux_virtual_machineazurerm_linux_virtual_machine_scale_setazurerm_windows_virtual_machineazurerm_windows_virtual_machine_scale_setazurerm_subnetazurerm_network_interfaceazurerm_network_security_groupazurerm_application_gatewayazurerm_application_security_groupazurerm_network_ddos_protection_planazurerm_firewallazurerm_local_network_gatewayazurerm_nat_gatewayazurerm_network_profileazurerm_network_security_ruleazurerm_public_ipazurerm_public_ip_prefixazurerm_routeazurerm_route_tableazurerm_virtual_network_gatewayazurerm_virtual_network_gateway_connectionazurerm_virtual_network_peeringazurerm_web_application_firewall_policyazurerm_virtual_hubazurerm_virtual_hub_bgp_connectionazurerm_virtual_hub_connectionazurerm_virtual_hub_ipazurerm_virtual_hub_route_tableazurerm_virtual_hub_security_partner_providerazurerm_lbazurerm_lb_backend_address_poolazurerm_lb_ruleazurerm_lb_outbound_ruleazurerm_lb_nat_ruleazurerm_lb_nat_poolazurerm_lb_probeazurerm_virtual_desktop_host_poolazurerm_virtual_desktop_application_groupazurerm_logic_app_workflowazurerm_logic_app_trigger_customazurerm_logic_app_action_customazurerm_container_registryazurerm_container_registry_webhookazurerm_kubernetes_clusterazurerm_kubernetes_cluster_node_poolazurerm_storage_accountazurerm_storage_queueazurerm_storage_shareazurerm_storage_tableazurerm_storage_blobazurerm_mariadb_configurationazurerm_mariadb_databaseazurerm_mariadb_firewall_ruleazurerm_mariadb_serverazurerm_mariadb_virtual_network_ruleazurerm_mysql_configurationazurerm_mysql_databaseazurerm_mysql_firewall_ruleazurerm_mysql_serverazurerm_mysql_virtual_network_ruleazurerm_postgresql_configurationazurerm_postgresql_databaseazurerm_postgresql_firewall_ruleazurerm_postgresql_serverazurerm_postgresql_virtual_network_ruleazurerm_mssql_elasticpoolazurerm_mssql_databaseazurerm_mssql_firewall_ruleazurerm_mssql_serverazurerm_mssql_server_security_alert_policyazurerm_mssql_server_vulnerability_assessmentazurerm_mssql_virtual_machineazurerm_mssql_virtual_network_ruleazurerm_redis_cacheazurerm_redis_firewall_ruleazurerm_dns_zoneazurerm_dns_a_recordazurerm_dns_aaaa_recordazurerm_dns_caa_recordazurerm_dns_cname_recordazurerm_dns_mx_recordazurerm_dns_ns_recordazurerm_dns_ptr_recordazurerm_dns_srv_recordazurerm_dns_txt_recordazurerm_private_dns_zoneazurerm_private_dns_a_recordazurerm_private_dns_aaaa_recordazurerm_private_dns_cname_recordazurerm_private_dns_mx_recordazurerm_private_dns_ptr_recordazurerm_private_dns_srv_recordazurerm_private_dns_txt_recordazurerm_private_dns_zone_virtual_network_linkazurerm_policy_definitionazurerm_policy_remediationazurerm_policy_set_definitionazurerm_key_vaultazurerm_key_vault_access_policyazurerm_application_insightsazurerm_application_insights_api_keyazurerm_application_insights_analytics_itemazurerm_log_analytics_workspaceazurerm_log_analytics_linked_serviceazurerm_log_analytics_datasource_windows_performance_counterazurerm_log_analytics_datasource_windows_eventazurerm_monitor_action_groupazurerm_monitor_activity_log_alertazurerm_monitor_autoscale_settingazurerm_monitor_log_profileazurerm_monitor_metric_alertazurerm_windows_web_appazurerm_linux_web_appazurerm_linux_web_app_slotazurerm_windows_web_app_slotazurerm_web_app_active_slotazurerm_service_planazurerm_source_control_tokenazurerm_static_siteazurerm_static_site_custom_domainazurerm_web_app_hybrid_connectionazurerm_data_protection_backup_vault"

func (i ResourceType) String() string {
	if i < 0 || i >= ResourceType(len(_ResourceTypeIndex)-1) {
		return fmt.Sprintf("ResourceType(%d)", i)
	}
	return _ResourceTypeName[_ResourceTypeIndex[i]:_ResourceTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ResourceTypeNoOp() {
	var x [1]struct{}
	_ = x[ResourceGroup-(0)]
	_ = x[AvailabilitySet-(1)]
	_ = x[Image-(2)]
	_ = x[ManagedDisk-(3)]
	_ = x[VirtualMachine-(4)]
	_ = x[VirtualMachineDataDiskAttachment-(5)]
	_ = x[VirtualMachineExtension-(6)]
	_ = x[VirtualMachineScaleSetExtension-(7)]
	_ = x[VirtualNetwork-(8)]
	_ = x[LinuxVirtualMachine-(9)]
	_ = x[LinuxVirtualMachineScaleSet-(10)]
	_ = x[WindowsVirtualMachine-(11)]
	_ = x[WindowsVirtualMachineScaleSet-(12)]
	_ = x[Subnet-(13)]
	_ = x[NetworkInterface-(14)]
	_ = x[NetworkSecurityGroup-(15)]
	_ = x[ApplicationGateway-(16)]
	_ = x[ApplicationSecurityGroup-(17)]
	_ = x[NetworkDdosProtectionPlan-(18)]
	_ = x[Firewall-(19)]
	_ = x[LocalNetworkGateway-(20)]
	_ = x[NatGateway-(21)]
	_ = x[NetworkProfile-(22)]
	_ = x[NetworkSecurityRule-(23)]
	_ = x[PublicIP-(24)]
	_ = x[PublicIPPrefix-(25)]
	_ = x[Route-(26)]
	_ = x[RouteTable-(27)]
	_ = x[VirtualNetworkGateway-(28)]
	_ = x[VirtualNetworkGatewayConnection-(29)]
	_ = x[VirtualNetworkPeering-(30)]
	_ = x[WebApplicationFirewallPolicy-(31)]
	_ = x[VirtualHub-(32)]
	_ = x[VirtualHubBgpConnection-(33)]
	_ = x[VirtualHubConnection-(34)]
	_ = x[VirtualHubIP-(35)]
	_ = x[VirtualHubRouteTable-(36)]
	_ = x[VirtualHubSecurityPartnerProvider-(37)]
	_ = x[Lb-(38)]
	_ = x[LbBackendAddressPool-(39)]
	_ = x[LbRule-(40)]
	_ = x[LbOutboundRule-(41)]
	_ = x[LbNatRule-(42)]
	_ = x[LbNatPool-(43)]
	_ = x[LbProbe-(44)]
	_ = x[VirtualDesktopHostPool-(45)]
	_ = x[VirtualDesktopApplicationGroup-(46)]
	_ = x[LogicAppWorkflow-(47)]
	_ = x[LogicAppTriggerCustom-(48)]
	_ = x[LogicAppActionCustom-(49)]
	_ = x[ContainerRegistry-(50)]
	_ = x[ContainerRegistryWebhook-(51)]
	_ = x[KubernetesCluster-(52)]
	_ = x[KubernetesClusterNodePool-(53)]
	_ = x[StorageAccount-(54)]
	_ = x[StorageQueue-(55)]
	_ = x[StorageShare-(56)]
	_ = x[StorageTable-(57)]
	_ = x[StorageBlob-(58)]
	_ = x[MariadbConfiguration-(59)]
	_ = x[MariadbDatabase-(60)]
	_ = x[MariadbFirewallRule-(61)]
	_ = x[MariadbServer-(62)]
	_ = x[MariadbVirtualNetworkRule-(63)]
	_ = x[MysqlConfiguration-(64)]
	_ = x[MysqlDatabase-(65)]
	_ = x[MysqlFirewallRule-(66)]
	_ = x[MysqlServer-(67)]
	_ = x[MysqlVirtualNetworkRule-(68)]
	_ = x[PostgresqlConfiguration-(69)]
	_ = x[PostgresqlDatabase-(70)]
	_ = x[PostgresqlFirewallRule-(71)]
	_ = x[PostgresqlServer-(72)]
	_ = x[PostgresqlVirtualNetworkRule-(73)]
	_ = x[MssqlElasticpool-(74)]
	_ = x[MssqlDatabase-(75)]
	_ = x[MssqlFirewallRule-(76)]
	_ = x[MssqlServer-(77)]
	_ = x[MssqlServerSecurityAlertPolicy-(78)]
	_ = x[MssqlServerVulnerabilityAssessment-(79)]
	_ = x[MssqlVirtualMachine-(80)]
	_ = x[MssqlVirtualNetworkRule-(81)]
	_ = x[RedisCache-(82)]
	_ = x[RedisFirewallRule-(83)]
	_ = x[DNSZone-(84)]
	_ = x[DNSARecord-(85)]
	_ = x[DNSAaaaRecord-(86)]
	_ = x[DNSCaaRecord-(87)]
	_ = x[DNSCnameRecord-(88)]
	_ = x[DNSMxRecord-(89)]
	_ = x[DNSNsRecord-(90)]
	_ = x[DNSPtrRecord-(91)]
	_ = x[DNSSrvRecord-(92)]
	_ = x[DNSTxtRecord-(93)]
	_ = x[PrivateDNSZone-(94)]
	_ = x[PrivateDNSARecord-(95)]
	_ = x[PrivateDNSAaaaRecord-(96)]
	_ = x[PrivateDNSCnameRecord-(97)]
	_ = x[PrivateDNSMxRecord-(98)]
	_ = x[PrivateDNSPtrRecord-(99)]
	_ = x[PrivateDNSSrvRecord-(100)]
	_ = x[PrivateDNSTxtRecord-(101)]
	_ = x[PrivateDNSZoneVirtualNetworkLink-(102)]
	_ = x[PolicyDefinition-(103)]
	_ = x[PolicyRemediation-(104)]
	_ = x[PolicySetDefinition-(105)]
	_ = x[KeyVault-(106)]
	_ = x[KeyVaultAccessPolicy-(107)]
	_ = x[ApplicationInsights-(108)]
	_ = x[ApplicationInsightsAPIKey-(109)]
	_ = x[ApplicationInsightsAnalyticsItem-(110)]
	_ = x[LogAnalyticsWorkspace-(111)]
	_ = x[LogAnalyticsLinkedService-(112)]
	_ = x[LogAnalyticsDatasourceWindowsPerformanceCounter-(113)]
	_ = x[LogAnalyticsDatasourceWindowsEvent-(114)]
	_ = x[MonitorActionGroup-(115)]
	_ = x[MonitorActivityLogAlert-(116)]
	_ = x[MonitorAutoscaleSetting-(117)]
	_ = x[MonitorLogProfile-(118)]
	_ = x[MonitorMetricAlert-(119)]
	_ = x[WindowsWebApp-(120)]
	_ = x[LinuxWebApp-(121)]
	_ = x[LinuxWebAppSlot-(122)]
	_ = x[WindowsWebAppSlot-(123)]
	_ = x[WebAppActiveSlot-(124)]
	_ = x[ServicePlan-(125)]
	_ = x[SourceControlToken-(126)]
	_ = x[StaticSite-(127)]
	_ = x[StaticSiteCustomDomain-(128)]
	_ = x[WebAppHybridConnection-(129)]
	_ = x[DataProtectionBackupVault-(130)]
}

var _ResourceTypeValues = []ResourceType{ResourceGroup, AvailabilitySet, Image, ManagedDisk, VirtualMachine, VirtualMachineDataDiskAttachment, VirtualMachineExtension, VirtualMachineScaleSetExtension, VirtualNetwork, LinuxVirtualMachine, LinuxVirtualMachineScaleSet, WindowsVirtualMachine, WindowsVirtualMachineScaleSet, Subnet, NetworkInterface, NetworkSecurityGroup, ApplicationGateway, ApplicationSecurityGroup, NetworkDdosProtectionPlan, Firewall, LocalNetworkGateway, NatGateway, NetworkProfile, NetworkSecurityRule, PublicIP, PublicIPPrefix, Route, RouteTable, VirtualNetworkGateway, VirtualNetworkGatewayConnection, VirtualNetworkPeering, WebApplicationFirewallPolicy, VirtualHub, VirtualHubBgpConnection, VirtualHubConnection, VirtualHubIP, VirtualHubRouteTable, VirtualHubSecurityPartnerProvider, Lb, LbBackendAddressPool, LbRule, LbOutboundRule, LbNatRule, LbNatPool, LbProbe, VirtualDesktopHostPool, VirtualDesktopApplicationGroup, LogicAppWorkflow, LogicAppTriggerCustom, LogicAppActionCustom, ContainerRegistry, ContainerRegistryWebhook, KubernetesCluster, KubernetesClusterNodePool, StorageAccount, StorageQueue, StorageShare, StorageTable, StorageBlob, MariadbConfiguration, MariadbDatabase, MariadbFirewallRule, MariadbServer, MariadbVirtualNetworkRule, MysqlConfiguration, MysqlDatabase, MysqlFirewallRule, MysqlServer, MysqlVirtualNetworkRule, PostgresqlConfiguration, PostgresqlDatabase, PostgresqlFirewallRule, PostgresqlServer, PostgresqlVirtualNetworkRule, MssqlElasticpool, MssqlDatabase, MssqlFirewallRule, MssqlServer, MssqlServerSecurityAlertPolicy, MssqlServerVulnerabilityAssessment, MssqlVirtualMachine, MssqlVirtualNetworkRule, RedisCache, RedisFirewallRule, DNSZone, DNSARecord, DNSAaaaRecord, DNSCaaRecord, DNSCnameRecord, DNSMxRecord, DNSNsRecord, DNSPtrRecord, DNSSrvRecord, DNSTxtRecord, PrivateDNSZone, PrivateDNSARecord, PrivateDNSAaaaRecord, PrivateDNSCnameRecord, PrivateDNSMxRecord, PrivateDNSPtrRecord, PrivateDNSSrvRecord, PrivateDNSTxtRecord, PrivateDNSZoneVirtualNetworkLink, PolicyDefinition, PolicyRemediation, PolicySetDefinition, KeyVault, KeyVaultAccessPolicy, ApplicationInsights, ApplicationInsightsAPIKey, ApplicationInsightsAnalyticsItem, LogAnalyticsWorkspace, LogAnalyticsLinkedService, LogAnalyticsDatasourceWindowsPerformanceCounter, LogAnalyticsDatasourceWindowsEvent, MonitorActionGroup, MonitorActivityLogAlert, MonitorAutoscaleSetting, MonitorLogProfile, MonitorMetricAlert, WindowsWebApp, LinuxWebApp, LinuxWebAppSlot, WindowsWebAppSlot, WebAppActiveSlot, ServicePlan, SourceControlToken, StaticSite, StaticSiteCustomDomain, WebAppHybridConnection, DataProtectionBackupVault}

var _ResourceTypeNameToValueMap = map[string]ResourceType{
	_ResourceTypeName[0:22]:           ResourceGroup,
	_ResourceTypeLowerName[0:22]:      ResourceGroup,
	_ResourceTypeName[22:46]:          AvailabilitySet,
	_ResourceTypeLowerName[22:46]:     AvailabilitySet,
	_ResourceTypeName[46:59]:          Image,
	_ResourceTypeLowerName[46:59]:     Image,
	_ResourceTypeName[59:79]:          ManagedDisk,
	_ResourceTypeLowerName[59:79]:     ManagedDisk,
	_ResourceTypeName[79:102]:         VirtualMachine,
	_ResourceTypeLowerName[79:102]:    VirtualMachine,
	_ResourceTypeName[102:146]:        VirtualMachineDataDiskAttachment,
	_ResourceTypeLowerName[102:146]:   VirtualMachineDataDiskAttachment,
	_ResourceTypeName[146:179]:        VirtualMachineExtension,
	_ResourceTypeLowerName[146:179]:   VirtualMachineExtension,
	_ResourceTypeName[179:222]:        VirtualMachineScaleSetExtension,
	_ResourceTypeLowerName[179:222]:   VirtualMachineScaleSetExtension,
	_ResourceTypeName[222:245]:        VirtualNetwork,
	_ResourceTypeLowerName[222:245]:   VirtualNetwork,
	_ResourceTypeName[245:274]:        LinuxVirtualMachine,
	_ResourceTypeLowerName[245:274]:   LinuxVirtualMachine,
	_ResourceTypeName[274:313]:        LinuxVirtualMachineScaleSet,
	_ResourceTypeLowerName[274:313]:   LinuxVirtualMachineScaleSet,
	_ResourceTypeName[313:344]:        WindowsVirtualMachine,
	_ResourceTypeLowerName[313:344]:   WindowsVirtualMachine,
	_ResourceTypeName[344:385]:        WindowsVirtualMachineScaleSet,
	_ResourceTypeLowerName[344:385]:   WindowsVirtualMachineScaleSet,
	_ResourceTypeName[385:399]:        Subnet,
	_ResourceTypeLowerName[385:399]:   Subnet,
	_ResourceTypeName[399:424]:        NetworkInterface,
	_ResourceTypeLowerName[399:424]:   NetworkInterface,
	_ResourceTypeName[424:454]:        NetworkSecurityGroup,
	_ResourceTypeLowerName[424:454]:   NetworkSecurityGroup,
	_ResourceTypeName[454:481]:        ApplicationGateway,
	_ResourceTypeLowerName[454:481]:   ApplicationGateway,
	_ResourceTypeName[481:515]:        ApplicationSecurityGroup,
	_ResourceTypeLowerName[481:515]:   ApplicationSecurityGroup,
	_ResourceTypeName[515:551]:        NetworkDdosProtectionPlan,
	_ResourceTypeLowerName[515:551]:   NetworkDdosProtectionPlan,
	_ResourceTypeName[551:567]:        Firewall,
	_ResourceTypeLowerName[551:567]:   Firewall,
	_ResourceTypeName[567:596]:        LocalNetworkGateway,
	_ResourceTypeLowerName[567:596]:   LocalNetworkGateway,
	_ResourceTypeName[596:615]:        NatGateway,
	_ResourceTypeLowerName[596:615]:   NatGateway,
	_ResourceTypeName[615:638]:        NetworkProfile,
	_ResourceTypeLowerName[615:638]:   NetworkProfile,
	_ResourceTypeName[638:667]:        NetworkSecurityRule,
	_ResourceTypeLowerName[638:667]:   NetworkSecurityRule,
	_ResourceTypeName[667:684]:        PublicIP,
	_ResourceTypeLowerName[667:684]:   PublicIP,
	_ResourceTypeName[684:708]:        PublicIPPrefix,
	_ResourceTypeLowerName[684:708]:   PublicIPPrefix,
	_ResourceTypeName[708:721]:        Route,
	_ResourceTypeLowerName[708:721]:   Route,
	_ResourceTypeName[721:740]:        RouteTable,
	_ResourceTypeLowerName[721:740]:   RouteTable,
	_ResourceTypeName[740:771]:        VirtualNetworkGateway,
	_ResourceTypeLowerName[740:771]:   VirtualNetworkGateway,
	_ResourceTypeName[771:813]:        VirtualNetworkGatewayConnection,
	_ResourceTypeLowerName[771:813]:   VirtualNetworkGatewayConnection,
	_ResourceTypeName[813:844]:        VirtualNetworkPeering,
	_ResourceTypeLowerName[813:844]:   VirtualNetworkPeering,
	_ResourceTypeName[844:883]:        WebApplicationFirewallPolicy,
	_ResourceTypeLowerName[844:883]:   WebApplicationFirewallPolicy,
	_ResourceTypeName[883:902]:        VirtualHub,
	_ResourceTypeLowerName[883:902]:   VirtualHub,
	_ResourceTypeName[902:936]:        VirtualHubBgpConnection,
	_ResourceTypeLowerName[902:936]:   VirtualHubBgpConnection,
	_ResourceTypeName[936:966]:        VirtualHubConnection,
	_ResourceTypeLowerName[936:966]:   VirtualHubConnection,
	_ResourceTypeName[966:988]:        VirtualHubIP,
	_ResourceTypeLowerName[966:988]:   VirtualHubIP,
	_ResourceTypeName[988:1019]:       VirtualHubRouteTable,
	_ResourceTypeLowerName[988:1019]:  VirtualHubRouteTable,
	_ResourceTypeName[1019:1064]:      VirtualHubSecurityPartnerProvider,
	_ResourceTypeLowerName[1019:1064]: VirtualHubSecurityPartnerProvider,
	_ResourceTypeName[1064:1074]:      Lb,
	_ResourceTypeLowerName[1064:1074]: Lb,
	_ResourceTypeName[1074:1105]:      LbBackendAddressPool,
	_ResourceTypeLowerName[1074:1105]: LbBackendAddressPool,
	_ResourceTypeName[1105:1120]:      LbRule,
	_ResourceTypeLowerName[1105:1120]: LbRule,
	_ResourceTypeName[1120:1144]:      LbOutboundRule,
	_ResourceTypeLowerName[1120:1144]: LbOutboundRule,
	_ResourceTypeName[1144:1163]:      LbNatRule,
	_ResourceTypeLowerName[1144:1163]: LbNatRule,
	_ResourceTypeName[1163:1182]:      LbNatPool,
	_ResourceTypeLowerName[1163:1182]: LbNatPool,
	_ResourceTypeName[1182:1198]:      LbProbe,
	_ResourceTypeLowerName[1182:1198]: LbProbe,
	_ResourceTypeName[1198:1231]:      VirtualDesktopHostPool,
	_ResourceTypeLowerName[1198:1231]: VirtualDesktopHostPool,
	_ResourceTypeName[1231:1272]:      VirtualDesktopApplicationGroup,
	_ResourceTypeLowerName[1231:1272]: VirtualDesktopApplicationGroup,
	_ResourceTypeName[1272:1298]:      LogicAppWorkflow,
	_ResourceTypeLowerName[1272:1298]: LogicAppWorkflow,
	_ResourceTypeName[1298:1330]:      LogicAppTriggerCustom,
	_ResourceTypeLowerName[1298:1330]: LogicAppTriggerCustom,
	_ResourceTypeName[1330:1361]:      LogicAppActionCustom,
	_ResourceTypeLowerName[1330:1361]: LogicAppActionCustom,
	_ResourceTypeName[1361:1387]:      ContainerRegistry,
	_ResourceTypeLowerName[1361:1387]: ContainerRegistry,
	_ResourceTypeName[1387:1421]:      ContainerRegistryWebhook,
	_ResourceTypeLowerName[1387:1421]: ContainerRegistryWebhook,
	_ResourceTypeName[1421:1447]:      KubernetesCluster,
	_ResourceTypeLowerName[1421:1447]: KubernetesCluster,
	_ResourceTypeName[1447:1483]:      KubernetesClusterNodePool,
	_ResourceTypeLowerName[1447:1483]: KubernetesClusterNodePool,
	_ResourceTypeName[1483:1506]:      StorageAccount,
	_ResourceTypeLowerName[1483:1506]: StorageAccount,
	_ResourceTypeName[1506:1527]:      StorageQueue,
	_ResourceTypeLowerName[1506:1527]: StorageQueue,
	_ResourceTypeName[1527:1548]:      StorageShare,
	_ResourceTypeLowerName[1527:1548]: StorageShare,
	_ResourceTypeName[1548:1569]:      StorageTable,
	_ResourceTypeLowerName[1548:1569]: StorageTable,
	_ResourceTypeName[1569:1589]:      StorageBlob,
	_ResourceTypeLowerName[1569:1589]: StorageBlob,
	_ResourceTypeName[1589:1618]:      MariadbConfiguration,
	_ResourceTypeLowerName[1589:1618]: MariadbConfiguration,
	_ResourceTypeName[1618:1642]:      MariadbDatabase,
	_ResourceTypeLowerName[1618:1642]: MariadbDatabase,
	_ResourceTypeName[1642:1671]:      MariadbFirewallRule,
	_ResourceTypeLowerName[1642:1671]: MariadbFirewallRule,
	_ResourceTypeName[1671:1693]:      MariadbServer,
	_ResourceTypeLowerName[1671:1693]: MariadbServer,
	_ResourceTypeName[1693:1729]:      MariadbVirtualNetworkRule,
	_ResourceTypeLowerName[1693:1729]: MariadbVirtualNetworkRule,
	_ResourceTypeName[1729:1756]:      MysqlConfiguration,
	_ResourceTypeLowerName[1729:1756]: MysqlConfiguration,
	_ResourceTypeName[1756:1778]:      MysqlDatabase,
	_ResourceTypeLowerName[1756:1778]: MysqlDatabase,
	_ResourceTypeName[1778:1805]:      MysqlFirewallRule,
	_ResourceTypeLowerName[1778:1805]: MysqlFirewallRule,
	_ResourceTypeName[1805:1825]:      MysqlServer,
	_ResourceTypeLowerName[1805:1825]: MysqlServer,
	_ResourceTypeName[1825:1859]:      MysqlVirtualNetworkRule,
	_ResourceTypeLowerName[1825:1859]: MysqlVirtualNetworkRule,
	_ResourceTypeName[1859:1891]:      PostgresqlConfiguration,
	_ResourceTypeLowerName[1859:1891]: PostgresqlConfiguration,
	_ResourceTypeName[1891:1918]:      PostgresqlDatabase,
	_ResourceTypeLowerName[1891:1918]: PostgresqlDatabase,
	_ResourceTypeName[1918:1950]:      PostgresqlFirewallRule,
	_ResourceTypeLowerName[1918:1950]: PostgresqlFirewallRule,
	_ResourceTypeName[1950:1975]:      PostgresqlServer,
	_ResourceTypeLowerName[1950:1975]: PostgresqlServer,
	_ResourceTypeName[1975:2014]:      PostgresqlVirtualNetworkRule,
	_ResourceTypeLowerName[1975:2014]: PostgresqlVirtualNetworkRule,
	_ResourceTypeName[2014:2039]:      MssqlElasticpool,
	_ResourceTypeLowerName[2014:2039]: MssqlElasticpool,
	_ResourceTypeName[2039:2061]:      MssqlDatabase,
	_ResourceTypeLowerName[2039:2061]: MssqlDatabase,
	_ResourceTypeName[2061:2088]:      MssqlFirewallRule,
	_ResourceTypeLowerName[2061:2088]: MssqlFirewallRule,
	_ResourceTypeName[2088:2108]:      MssqlServer,
	_ResourceTypeLowerName[2088:2108]: MssqlServer,
	_ResourceTypeName[2108:2150]:      MssqlServerSecurityAlertPolicy,
	_ResourceTypeLowerName[2108:2150]: MssqlServerSecurityAlertPolicy,
	_ResourceTypeName[2150:2195]:      MssqlServerVulnerabilityAssessment,
	_ResourceTypeLowerName[2150:2195]: MssqlServerVulnerabilityAssessment,
	_ResourceTypeName[2195:2224]:      MssqlVirtualMachine,
	_ResourceTypeLowerName[2195:2224]: MssqlVirtualMachine,
	_ResourceTypeName[2224:2258]:      MssqlVirtualNetworkRule,
	_ResourceTypeLowerName[2224:2258]: MssqlVirtualNetworkRule,
	_ResourceTypeName[2258:2277]:      RedisCache,
	_ResourceTypeLowerName[2258:2277]: RedisCache,
	_ResourceTypeName[2277:2304]:      RedisFirewallRule,
	_ResourceTypeLowerName[2277:2304]: RedisFirewallRule,
	_ResourceTypeName[2304:2320]:      DNSZone,
	_ResourceTypeLowerName[2304:2320]: DNSZone,
	_ResourceTypeName[2320:2340]:      DNSARecord,
	_ResourceTypeLowerName[2320:2340]: DNSARecord,
	_ResourceTypeName[2340:2363]:      DNSAaaaRecord,
	_ResourceTypeLowerName[2340:2363]: DNSAaaaRecord,
	_ResourceTypeName[2363:2385]:      DNSCaaRecord,
	_ResourceTypeLowerName[2363:2385]: DNSCaaRecord,
	_ResourceTypeName[2385:2409]:      DNSCnameRecord,
	_ResourceTypeLowerName[2385:2409]: DNSCnameRecord,
	_ResourceTypeName[2409:2430]:      DNSMxRecord,
	_ResourceTypeLowerName[2409:2430]: DNSMxRecord,
	_ResourceTypeName[2430:2451]:      DNSNsRecord,
	_ResourceTypeLowerName[2430:2451]: DNSNsRecord,
	_ResourceTypeName[2451:2473]:      DNSPtrRecord,
	_ResourceTypeLowerName[2451:2473]: DNSPtrRecord,
	_ResourceTypeName[2473:2495]:      DNSSrvRecord,
	_ResourceTypeLowerName[2473:2495]: DNSSrvRecord,
	_ResourceTypeName[2495:2517]:      DNSTxtRecord,
	_ResourceTypeLowerName[2495:2517]: DNSTxtRecord,
	_ResourceTypeName[2517:2541]:      PrivateDNSZone,
	_ResourceTypeLowerName[2517:2541]: PrivateDNSZone,
	_ResourceTypeName[2541:2569]:      PrivateDNSARecord,
	_ResourceTypeLowerName[2541:2569]: PrivateDNSARecord,
	_ResourceTypeName[2569:2600]:      PrivateDNSAaaaRecord,
	_ResourceTypeLowerName[2569:2600]: PrivateDNSAaaaRecord,
	_ResourceTypeName[2600:2632]:      PrivateDNSCnameRecord,
	_ResourceTypeLowerName[2600:2632]: PrivateDNSCnameRecord,
	_ResourceTypeName[2632:2661]:      PrivateDNSMxRecord,
	_ResourceTypeLowerName[2632:2661]: PrivateDNSMxRecord,
	_ResourceTypeName[2661:2691]:      PrivateDNSPtrRecord,
	_ResourceTypeLowerName[2661:2691]: PrivateDNSPtrRecord,
	_ResourceTypeName[2691:2721]:      PrivateDNSSrvRecord,
	_ResourceTypeLowerName[2691:2721]: PrivateDNSSrvRecord,
	_ResourceTypeName[2721:2751]:      PrivateDNSTxtRecord,
	_ResourceTypeLowerName[2721:2751]: PrivateDNSTxtRecord,
	_ResourceTypeName[2751:2796]:      PrivateDNSZoneVirtualNetworkLink,
	_ResourceTypeLowerName[2751:2796]: PrivateDNSZoneVirtualNetworkLink,
	_ResourceTypeName[2796:2821]:      PolicyDefinition,
	_ResourceTypeLowerName[2796:2821]: PolicyDefinition,
	_ResourceTypeName[2821:2847]:      PolicyRemediation,
	_ResourceTypeLowerName[2821:2847]: PolicyRemediation,
	_ResourceTypeName[2847:2876]:      PolicySetDefinition,
	_ResourceTypeLowerName[2847:2876]: PolicySetDefinition,
	_ResourceTypeName[2876:2893]:      KeyVault,
	_ResourceTypeLowerName[2876:2893]: KeyVault,
	_ResourceTypeName[2893:2924]:      KeyVaultAccessPolicy,
	_ResourceTypeLowerName[2893:2924]: KeyVaultAccessPolicy,
	_ResourceTypeName[2924:2952]:      ApplicationInsights,
	_ResourceTypeLowerName[2924:2952]: ApplicationInsights,
	_ResourceTypeName[2952:2988]:      ApplicationInsightsAPIKey,
	_ResourceTypeLowerName[2952:2988]: ApplicationInsightsAPIKey,
	_ResourceTypeName[2988:3031]:      ApplicationInsightsAnalyticsItem,
	_ResourceTypeLowerName[2988:3031]: ApplicationInsightsAnalyticsItem,
	_ResourceTypeName[3031:3062]:      LogAnalyticsWorkspace,
	_ResourceTypeLowerName[3031:3062]: LogAnalyticsWorkspace,
	_ResourceTypeName[3062:3098]:      LogAnalyticsLinkedService,
	_ResourceTypeLowerName[3062:3098]: LogAnalyticsLinkedService,
	_ResourceTypeName[3098:3158]:      LogAnalyticsDatasourceWindowsPerformanceCounter,
	_ResourceTypeLowerName[3098:3158]: LogAnalyticsDatasourceWindowsPerformanceCounter,
	_ResourceTypeName[3158:3204]:      LogAnalyticsDatasourceWindowsEvent,
	_ResourceTypeLowerName[3158:3204]: LogAnalyticsDatasourceWindowsEvent,
	_ResourceTypeName[3204:3232]:      MonitorActionGroup,
	_ResourceTypeLowerName[3204:3232]: MonitorActionGroup,
	_ResourceTypeName[3232:3266]:      MonitorActivityLogAlert,
	_ResourceTypeLowerName[3232:3266]: MonitorActivityLogAlert,
	_ResourceTypeName[3266:3299]:      MonitorAutoscaleSetting,
	_ResourceTypeLowerName[3266:3299]: MonitorAutoscaleSetting,
	_ResourceTypeName[3299:3326]:      MonitorLogProfile,
	_ResourceTypeLowerName[3299:3326]: MonitorLogProfile,
	_ResourceTypeName[3326:3354]:      MonitorMetricAlert,
	_ResourceTypeLowerName[3326:3354]: MonitorMetricAlert,
	_ResourceTypeName[3354:3377]:      WindowsWebApp,
	_ResourceTypeLowerName[3354:3377]: WindowsWebApp,
	_ResourceTypeName[3377:3398]:      LinuxWebApp,
	_ResourceTypeLowerName[3377:3398]: LinuxWebApp,
	_ResourceTypeName[3398:3424]:      LinuxWebAppSlot,
	_ResourceTypeLowerName[3398:3424]: LinuxWebAppSlot,
	_ResourceTypeName[3424:3452]:      WindowsWebAppSlot,
	_ResourceTypeLowerName[3424:3452]: WindowsWebAppSlot,
	_ResourceTypeName[3452:3479]:      WebAppActiveSlot,
	_ResourceTypeLowerName[3452:3479]: WebAppActiveSlot,
	_ResourceTypeName[3479:3499]:      ServicePlan,
	_ResourceTypeLowerName[3479:3499]: ServicePlan,
	_ResourceTypeName[3499:3527]:      SourceControlToken,
	_ResourceTypeLowerName[3499:3527]: SourceControlToken,
	_ResourceTypeName[3527:3546]:      StaticSite,
	_ResourceTypeLowerName[3527:3546]: StaticSite,
	_ResourceTypeName[3546:3579]:      StaticSiteCustomDomain,
	_ResourceTypeLowerName[3546:3579]: StaticSiteCustomDomain,
	_ResourceTypeName[3579:3612]:      WebAppHybridConnection,
	_ResourceTypeLowerName[3579:3612]: WebAppHybridConnection,
	_ResourceTypeName[3612:3648]:      DataProtectionBackupVault,
	_ResourceTypeLowerName[3612:3648]: DataProtectionBackupVault,
}

var _ResourceTypeNames = []string{
	_ResourceTypeName[0:22],
	_ResourceTypeName[22:46],
	_ResourceTypeName[46:59],
	_ResourceTypeName[59:79],
	_ResourceTypeName[79:102],
	_ResourceTypeName[102:146],
	_ResourceTypeName[146:179],
	_ResourceTypeName[179:222],
	_ResourceTypeName[222:245],
	_ResourceTypeName[245:274],
	_ResourceTypeName[274:313],
	_ResourceTypeName[313:344],
	_ResourceTypeName[344:385],
	_ResourceTypeName[385:399],
	_ResourceTypeName[399:424],
	_ResourceTypeName[424:454],
	_ResourceTypeName[454:481],
	_ResourceTypeName[481:515],
	_ResourceTypeName[515:551],
	_ResourceTypeName[551:567],
	_ResourceTypeName[567:596],
	_ResourceTypeName[596:615],
	_ResourceTypeName[615:638],
	_ResourceTypeName[638:667],
	_ResourceTypeName[667:684],
	_ResourceTypeName[684:708],
	_ResourceTypeName[708:721],
	_ResourceTypeName[721:740],
	_ResourceTypeName[740:771],
	_ResourceTypeName[771:813],
	_ResourceTypeName[813:844],
	_ResourceTypeName[844:883],
	_ResourceTypeName[883:902],
	_ResourceTypeName[902:936],
	_ResourceTypeName[936:966],
	_ResourceTypeName[966:988],
	_ResourceTypeName[988:1019],
	_ResourceTypeName[1019:1064],
	_ResourceTypeName[1064:1074],
	_ResourceTypeName[1074:1105],
	_ResourceTypeName[1105:1120],
	_ResourceTypeName[1120:1144],
	_ResourceTypeName[1144:1163],
	_ResourceTypeName[1163:1182],
	_ResourceTypeName[1182:1198],
	_ResourceTypeName[1198:1231],
	_ResourceTypeName[1231:1272],
	_ResourceTypeName[1272:1298],
	_ResourceTypeName[1298:1330],
	_ResourceTypeName[1330:1361],
	_ResourceTypeName[1361:1387],
	_ResourceTypeName[1387:1421],
	_ResourceTypeName[1421:1447],
	_ResourceTypeName[1447:1483],
	_ResourceTypeName[1483:1506],
	_ResourceTypeName[1506:1527],
	_ResourceTypeName[1527:1548],
	_ResourceTypeName[1548:1569],
	_ResourceTypeName[1569:1589],
	_ResourceTypeName[1589:1618],
	_ResourceTypeName[1618:1642],
	_ResourceTypeName[1642:1671],
	_ResourceTypeName[1671:1693],
	_ResourceTypeName[1693:1729],
	_ResourceTypeName[1729:1756],
	_ResourceTypeName[1756:1778],
	_ResourceTypeName[1778:1805],
	_ResourceTypeName[1805:1825],
	_ResourceTypeName[1825:1859],
	_ResourceTypeName[1859:1891],
	_ResourceTypeName[1891:1918],
	_ResourceTypeName[1918:1950],
	_ResourceTypeName[1950:1975],
	_ResourceTypeName[1975:2014],
	_ResourceTypeName[2014:2039],
	_ResourceTypeName[2039:2061],
	_ResourceTypeName[2061:2088],
	_ResourceTypeName[2088:2108],
	_ResourceTypeName[2108:2150],
	_ResourceTypeName[2150:2195],
	_ResourceTypeName[2195:2224],
	_ResourceTypeName[2224:2258],
	_ResourceTypeName[2258:2277],
	_ResourceTypeName[2277:2304],
	_ResourceTypeName[2304:2320],
	_ResourceTypeName[2320:2340],
	_ResourceTypeName[2340:2363],
	_ResourceTypeName[2363:2385],
	_ResourceTypeName[2385:2409],
	_ResourceTypeName[2409:2430],
	_ResourceTypeName[2430:2451],
	_ResourceTypeName[2451:2473],
	_ResourceTypeName[2473:2495],
	_ResourceTypeName[2495:2517],
	_ResourceTypeName[2517:2541],
	_ResourceTypeName[2541:2569],
	_ResourceTypeName[2569:2600],
	_ResourceTypeName[2600:2632],
	_ResourceTypeName[2632:2661],
	_ResourceTypeName[2661:2691],
	_ResourceTypeName[2691:2721],
	_ResourceTypeName[2721:2751],
	_ResourceTypeName[2751:2796],
	_ResourceTypeName[2796:2821],
	_ResourceTypeName[2821:2847],
	_ResourceTypeName[2847:2876],
	_ResourceTypeName[2876:2893],
	_ResourceTypeName[2893:2924],
	_ResourceTypeName[2924:2952],
	_ResourceTypeName[2952:2988],
	_ResourceTypeName[2988:3031],
	_ResourceTypeName[3031:3062],
	_ResourceTypeName[3062:3098],
	_ResourceTypeName[3098:3158],
	_ResourceTypeName[3158:3204],
	_ResourceTypeName[3204:3232],
	_ResourceTypeName[3232:3266],
	_ResourceTypeName[3266:3299],
	_ResourceTypeName[3299:3326],
	_ResourceTypeName[3326:3354],
	_ResourceTypeName[3354:3377],
	_ResourceTypeName[3377:3398],
	_ResourceTypeName[3398:3424],
	_ResourceTypeName[3424:3452],
	_ResourceTypeName[3452:3479],
	_ResourceTypeName[3479:3499],
	_ResourceTypeName[3499:3527],
	_ResourceTypeName[3527:3546],
	_ResourceTypeName[3546:3579],
	_ResourceTypeName[3579:3612],
	_ResourceTypeName[3612:3648],
}

// ResourceTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ResourceTypeString(s string) (ResourceType, error) {
	if val, ok := _ResourceTypeNameToValueMap[s]; ok {
		return val, nil
	}
	s = strings.ToLower(s)
	if val, ok := _ResourceTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ResourceType values", s)
}

// ResourceTypeValues returns all values of the enum
func ResourceTypeValues() []ResourceType {
	return _ResourceTypeValues
}

// ResourceTypeStrings returns a slice of all String values of the enum
func ResourceTypeStrings() []string {
	strs := make([]string, len(_ResourceTypeNames))
	copy(strs, _ResourceTypeNames)
	return strs
}

// IsAResourceType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ResourceType) IsAResourceType() bool {
	for _, v := range _ResourceTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
