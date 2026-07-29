// This code was generated. Review it carefully.

package mikrotik

import (
	"context"
	"github.com/ddelnano/terraform-provider-mikrotik/client"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	tftypes "github.com/hashicorp/terraform-plugin-framework/types"
)

type interfaceWiFiConfiguration struct {
	client *client.Mikrotik
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &interfaceWiFiConfiguration{}
	_ resource.ResourceWithConfigure   = &interfaceWiFiConfiguration{}
	_ resource.ResourceWithImportState = &interfaceWiFiConfiguration{}
)

// NewInterfaceWiFiConfigurationResource is a helper function to simplify the provider implementation.
func NewInterfaceWiFiConfigurationResource() resource.Resource {
	return &interfaceWiFiConfiguration{}
}

func (r *interfaceWiFiConfiguration) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Mikrotik)
}

// Metadata returns the resource type name.
func (r *interfaceWiFiConfiguration) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_wi_fi_configuration"
}

// Schema defines the schema for the resource.
func (s *interfaceWiFiConfiguration) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Creates a MikroTik InterfaceWiFiConfiguration.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: false,
				Optional: false,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "",
			},
			"aaa": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"aaa_called_format": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"aaa_calling_format": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"aaa_interim_update": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"aaa_mac_caching": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"aaa_nas_identifier": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"aaa_password_format": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"aaa_username_format": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"antenna_gain": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"beacon_interval": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"chains": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"channel": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"channel_band": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"channel_deprioritize_unii_3_4": schema.BoolAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"channel_frequency": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"channel_reselect_interval": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"channel_reselect_time": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"channel_secondary_frequency": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"channel_skip_dfs_channels": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"channel_width": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"comment": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"copy_from": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"country": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"datapath": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"datapath_bridge": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"datapath_bridge_cost": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"datapath_bridge_horizon": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"datapath_client_isolation": schema.BoolAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"datapath_interface_list": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"datapath_traffic_processing": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"datapath_vlan_id": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"distance": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"dtim_period": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"hide_ssid": schema.BoolAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"installation": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_3gpp_info": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_authentication_types": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_connection_capabilities": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_domain_names": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_esr": schema.BoolAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_hessid": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_hotspot20": schema.BoolAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_hotspot20_dgaf": schema.BoolAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_internet": schema.BoolAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_ipv4_availability": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_ipv6_availability": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_network_type": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_operational_classes": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_operator_names": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_realms": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_roaming_ois": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_uesa": schema.BoolAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_venue": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_venue_names": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_wan_at_capacity": schema.BoolAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_wan_downlink": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_wan_downlink_load": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_wan_measurement_duration": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_wan_status": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_wan_symmetric": schema.BoolAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_wan_uplink": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"interworking_wan_uplink_load": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"manager": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"max_clients": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"mode": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"multicast_enhance": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"name": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"qos_classifier": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_authentication_types": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_connect_group": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_connect_priority": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_dh_groups": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_disable_pmkid": schema.BoolAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_eap_accounting": schema.BoolAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_eap_anonymous_identity": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_eap_certificate_mode": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_eap_methods": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_eap_password": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_eap_tls_certificate": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_eap_username": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_encryption": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_ft": schema.BoolAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_ft_mobility_domain": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_ft_nas_identifier": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_ft_over_ds": schema.BoolAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_ft_preserve_vlanid": schema.BoolAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_ft_r0_key_lifetime": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_ft_reassociation_deadline": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_group_encryption": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_group_key_update": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_management_encryption": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_management_protection": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_multi_passphrase_group": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_owe_transition_interface": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_passphrase": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_sae_anti_clogging_threshold": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_sae_max_failure_rate": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_sae_pwe": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"security_wps": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"ssid": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"station_roaming": schema.BoolAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"steering": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"steering_2g_probe_delay": schema.BoolAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"steering_neighbor_group": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"steering_rrm": schema.BoolAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"steering_wnm": schema.BoolAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"tx_chains": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"tx_power": schema.StringAttribute{
				Required: false,
				Optional: true,
				Computed: false,

				Description: "",
			},
			"about": schema.StringAttribute{
				Required: false,
				Optional: false,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "",
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *interfaceWiFiConfiguration) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var terraformModel interfaceWiFiConfigurationModel
	var mikrotikModel client.InterfaceWiFiConfiguration
	GenericCreateResource(&terraformModel, &mikrotikModel, r.client, r)(ctx, req, resp)
}

// Read refreshes the Terraform state with the latest data.
func (r *interfaceWiFiConfiguration) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var terraformModel interfaceWiFiConfigurationModel
	var mikrotikModel client.InterfaceWiFiConfiguration
	GenericReadResource(&terraformModel, &mikrotikModel, r.client, r)(ctx, req, resp)
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *interfaceWiFiConfiguration) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var terraformModel interfaceWiFiConfigurationModel
	var mikrotikModel client.InterfaceWiFiConfiguration
	GenericUpdateResource(&terraformModel, &mikrotikModel, r.client, r)(ctx, req, resp)
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *interfaceWiFiConfiguration) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var terraformModel interfaceWiFiConfigurationModel
	var mikrotikModel client.InterfaceWiFiConfiguration
	GenericDeleteResource(&terraformModel, &mikrotikModel, r.client)(ctx, req, resp)
}

func (r *interfaceWiFiConfiguration) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

type interfaceWiFiConfigurationModel struct {
	Id                                 tftypes.String `tfsdk:"id"`
	Aaa                                tftypes.String `tfsdk:"aaa"`
	AaaCalledFormat                    tftypes.String `tfsdk:"aaa_called_format"`
	AaaCallingFormat                   tftypes.String `tfsdk:"aaa_calling_format"`
	AaaInterimUpdate                   tftypes.String `tfsdk:"aaa_interim_update"`
	AaaMacCaching                      tftypes.String `tfsdk:"aaa_mac_caching"`
	AaaNasIdentifier                   tftypes.String `tfsdk:"aaa_nas_identifier"`
	AaaPasswordFormat                  tftypes.String `tfsdk:"aaa_password_format"`
	AaaUsernameFormat                  tftypes.String `tfsdk:"aaa_username_format"`
	AntennaGain                        tftypes.String `tfsdk:"antenna_gain"`
	BeaconInterval                     tftypes.String `tfsdk:"beacon_interval"`
	Chains                             tftypes.String `tfsdk:"chains"`
	Channel                            tftypes.String `tfsdk:"channel"`
	ChannelBand                        tftypes.String `tfsdk:"channel_band"`
	ChannelDeprioritizeUnii34          tftypes.Bool   `tfsdk:"channel_deprioritize_unii_3_4"`
	ChannelFrequency                   tftypes.String `tfsdk:"channel_frequency"`
	ChannelReselectInterval            tftypes.String `tfsdk:"channel_reselect_interval"`
	ChannelReselectTime                tftypes.String `tfsdk:"channel_reselect_time"`
	ChannelSecondaryFrequency          tftypes.String `tfsdk:"channel_secondary_frequency"`
	ChannelSkipDfsChannels             tftypes.String `tfsdk:"channel_skip_dfs_channels"`
	ChannelWidth                       tftypes.String `tfsdk:"channel_width"`
	Comment                            tftypes.String `tfsdk:"comment"`
	CopyFrom                           tftypes.String `tfsdk:"copy_from"`
	Country                            tftypes.String `tfsdk:"country"`
	Datapath                           tftypes.String `tfsdk:"datapath"`
	DatapathBridge                     tftypes.String `tfsdk:"datapath_bridge"`
	DatapathBridgeCost                 tftypes.String `tfsdk:"datapath_bridge_cost"`
	DatapathBridgeHorizon              tftypes.String `tfsdk:"datapath_bridge_horizon"`
	DatapathClientIsolation            tftypes.Bool   `tfsdk:"datapath_client_isolation"`
	DatapathInterfaceList              tftypes.String `tfsdk:"datapath_interface_list"`
	DatapathTrafficProcessing          tftypes.String `tfsdk:"datapath_traffic_processing"`
	DatapathVlanId                     tftypes.String `tfsdk:"datapath_vlan_id"`
	Disabled                           tftypes.Bool   `tfsdk:"disabled"`
	Distance                           tftypes.String `tfsdk:"distance"`
	DtimPeriod                         tftypes.String `tfsdk:"dtim_period"`
	HideSsid                           tftypes.Bool   `tfsdk:"hide_ssid"`
	Installation                       tftypes.String `tfsdk:"installation"`
	Interworking                       tftypes.String `tfsdk:"interworking"`
	Interworking3gppInfo               tftypes.String `tfsdk:"interworking_3gpp_info"`
	InterworkingAuthenticationTypes    tftypes.String `tfsdk:"interworking_authentication_types"`
	InterworkingConnectionCapabilities tftypes.String `tfsdk:"interworking_connection_capabilities"`
	InterworkingDomainNames            tftypes.String `tfsdk:"interworking_domain_names"`
	InterworkingEsr                    tftypes.Bool   `tfsdk:"interworking_esr"`
	InterworkingHessid                 tftypes.String `tfsdk:"interworking_hessid"`
	InterworkingHotspot20              tftypes.Bool   `tfsdk:"interworking_hotspot20"`
	InterworkingHotspot20Dgaf          tftypes.Bool   `tfsdk:"interworking_hotspot20_dgaf"`
	InterworkingInternet               tftypes.Bool   `tfsdk:"interworking_internet"`
	InterworkingIpv4Availability       tftypes.String `tfsdk:"interworking_ipv4_availability"`
	InterworkingIpv6Availability       tftypes.String `tfsdk:"interworking_ipv6_availability"`
	InterworkingNetworkType            tftypes.String `tfsdk:"interworking_network_type"`
	InterworkingOperationalClasses     tftypes.String `tfsdk:"interworking_operational_classes"`
	InterworkingOperatorNames          tftypes.String `tfsdk:"interworking_operator_names"`
	InterworkingRealms                 tftypes.String `tfsdk:"interworking_realms"`
	InterworkingRoamingOis             tftypes.String `tfsdk:"interworking_roaming_ois"`
	InterworkingUesa                   tftypes.Bool   `tfsdk:"interworking_uesa"`
	InterworkingVenue                  tftypes.String `tfsdk:"interworking_venue"`
	InterworkingVenueNames             tftypes.String `tfsdk:"interworking_venue_names"`
	InterworkingWanAtCapacity          tftypes.Bool   `tfsdk:"interworking_wan_at_capacity"`
	InterworkingWanDownlink            tftypes.String `tfsdk:"interworking_wan_downlink"`
	InterworkingWanDownlinkLoad        tftypes.String `tfsdk:"interworking_wan_downlink_load"`
	InterworkingWanMeasurementDuration tftypes.String `tfsdk:"interworking_wan_measurement_duration"`
	InterworkingWanStatus              tftypes.String `tfsdk:"interworking_wan_status"`
	InterworkingWanSymmetric           tftypes.Bool   `tfsdk:"interworking_wan_symmetric"`
	InterworkingWanUplink              tftypes.String `tfsdk:"interworking_wan_uplink"`
	InterworkingWanUplinkLoad          tftypes.String `tfsdk:"interworking_wan_uplink_load"`
	Manager                            tftypes.String `tfsdk:"manager"`
	MaxClients                         tftypes.String `tfsdk:"max_clients"`
	Mode                               tftypes.String `tfsdk:"mode"`
	MulticastEnhance                   tftypes.String `tfsdk:"multicast_enhance"`
	Name                               tftypes.String `tfsdk:"name"`
	QosClassifier                      tftypes.String `tfsdk:"qos_classifier"`
	Security                           tftypes.String `tfsdk:"security"`
	SecurityAuthenticationTypes        tftypes.String `tfsdk:"security_authentication_types"`
	SecurityConnectGroup               tftypes.String `tfsdk:"security_connect_group"`
	SecurityConnectPriority            tftypes.String `tfsdk:"security_connect_priority"`
	SecurityDhGroups                   tftypes.String `tfsdk:"security_dh_groups"`
	SecurityDisablePmkid               tftypes.Bool   `tfsdk:"security_disable_pmkid"`
	SecurityEapAccounting              tftypes.Bool   `tfsdk:"security_eap_accounting"`
	SecurityEapAnonymousIdentity       tftypes.String `tfsdk:"security_eap_anonymous_identity"`
	SecurityEapCertificateMode         tftypes.String `tfsdk:"security_eap_certificate_mode"`
	SecurityEapMethods                 tftypes.String `tfsdk:"security_eap_methods"`
	SecurityEapPassword                tftypes.String `tfsdk:"security_eap_password"`
	SecurityEapTlsCertificate          tftypes.String `tfsdk:"security_eap_tls_certificate"`
	SecurityEapUsername                tftypes.String `tfsdk:"security_eap_username"`
	SecurityEncryption                 tftypes.String `tfsdk:"security_encryption"`
	SecurityFt                         tftypes.Bool   `tfsdk:"security_ft"`
	SecurityFtMobilityDomain           tftypes.String `tfsdk:"security_ft_mobility_domain"`
	SecurityFtNasIdentifier            tftypes.String `tfsdk:"security_ft_nas_identifier"`
	SecurityFtOverDs                   tftypes.Bool   `tfsdk:"security_ft_over_ds"`
	SecurityFtPreserveVlanid           tftypes.Bool   `tfsdk:"security_ft_preserve_vlanid"`
	SecurityFtR0KeyLifetime            tftypes.String `tfsdk:"security_ft_r0_key_lifetime"`
	SecurityFtReassociationDeadline    tftypes.String `tfsdk:"security_ft_reassociation_deadline"`
	SecurityGroupEncryption            tftypes.String `tfsdk:"security_group_encryption"`
	SecurityGroupKeyUpdate             tftypes.String `tfsdk:"security_group_key_update"`
	SecurityManagementEncryption       tftypes.String `tfsdk:"security_management_encryption"`
	SecurityManagementProtection       tftypes.String `tfsdk:"security_management_protection"`
	SecurityMultiPassphraseGroup       tftypes.String `tfsdk:"security_multi_passphrase_group"`
	SecurityOweTransitionInterface     tftypes.String `tfsdk:"security_owe_transition_interface"`
	SecurityPassphrase                 tftypes.String `tfsdk:"security_passphrase"`
	SecuritySaeAntiCloggingThreshold   tftypes.String `tfsdk:"security_sae_anti_clogging_threshold"`
	SecuritySaeMaxFailureRate          tftypes.String `tfsdk:"security_sae_max_failure_rate"`
	SecuritySaePwe                     tftypes.String `tfsdk:"security_sae_pwe"`
	SecurityWps                        tftypes.String `tfsdk:"security_wps"`
	Ssid                               tftypes.String `tfsdk:"ssid"`
	StationRoaming                     tftypes.Bool   `tfsdk:"station_roaming"`
	Steering                           tftypes.String `tfsdk:"steering"`
	Steering2gProbeDelay               tftypes.Bool   `tfsdk:"steering_2g_probe_delay"`
	SteeringNeighborGroup              tftypes.String `tfsdk:"steering_neighbor_group"`
	SteeringRrm                        tftypes.Bool   `tfsdk:"steering_rrm"`
	SteeringWnm                        tftypes.Bool   `tfsdk:"steering_wnm"`
	TxChains                           tftypes.String `tfsdk:"tx_chains"`
	TxPower                            tftypes.String `tfsdk:"tx_power"`
	About                              tftypes.String `tfsdk:"about"`
}
