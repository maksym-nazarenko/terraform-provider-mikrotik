// This code was auto-generated.

package client

import (
	"github.com/go-routeros/routeros/v3"
)

// Constants for aaa.interim-update
const (
	InterfaceWiFiConfiguration_AaaInterimUpdate_Disabled = "disabled"
)

// Constants for aaa.mac-caching
const (
	InterfaceWiFiConfiguration_AaaMacCaching_Disabled = "disabled"
)

// Constants for chains
const (
	InterfaceWiFiConfiguration_Chains_0 = "0"
	InterfaceWiFiConfiguration_Chains_1 = "1"
	InterfaceWiFiConfiguration_Chains_2 = "2"
	InterfaceWiFiConfiguration_Chains_3 = "3"
	InterfaceWiFiConfiguration_Chains_4 = "4"
	InterfaceWiFiConfiguration_Chains_5 = "5"
	InterfaceWiFiConfiguration_Chains_6 = "6"
	InterfaceWiFiConfiguration_Chains_7 = "7"
)

// Constants for channel.band
const (
	InterfaceWiFiConfiguration_ChannelBand_2ghzAx = "2ghz-ax"
	InterfaceWiFiConfiguration_ChannelBand_2ghzBe = "2ghz-be"
	InterfaceWiFiConfiguration_ChannelBand_2ghzG  = "2ghz-g"
	InterfaceWiFiConfiguration_ChannelBand_2ghzN  = "2ghz-n"
	InterfaceWiFiConfiguration_ChannelBand_5ghzA  = "5ghz-a"
	InterfaceWiFiConfiguration_ChannelBand_5ghzAc = "5ghz-ac"
	InterfaceWiFiConfiguration_ChannelBand_5ghzAx = "5ghz-ax"
	InterfaceWiFiConfiguration_ChannelBand_5ghzBe = "5ghz-be"
	InterfaceWiFiConfiguration_ChannelBand_5ghzN  = "5ghz-n"
)

// Constants for channel.secondary-frequency
const (
	InterfaceWiFiConfiguration_ChannelSecondaryFrequency_Disabled = "disabled"
	InterfaceWiFiConfiguration_ChannelSecondaryFrequency_         = ","
)

// Constants for channel.skip-dfs-channels
const (
	InterfaceWiFiConfiguration_ChannelSkipDfsChannels_10minCac = "10min-cac"
	InterfaceWiFiConfiguration_ChannelSkipDfsChannels_All      = "all"
	InterfaceWiFiConfiguration_ChannelSkipDfsChannels_Disabled = "disabled"
)

// Constants for channel.width
const (
	InterfaceWiFiConfiguration_ChannelWidth_20408080mhz     = "20/40/80+80mhz"
	InterfaceWiFiConfiguration_ChannelWidth_204080160320mhz = "20/40/80/160/320mhz"
	InterfaceWiFiConfiguration_ChannelWidth_204080160mhz    = "20/40/80/160mhz"
	InterfaceWiFiConfiguration_ChannelWidth_204080mhz       = "20/40/80mhz"
	InterfaceWiFiConfiguration_ChannelWidth_2040mhz         = "20/40mhz"
	InterfaceWiFiConfiguration_ChannelWidth_2040mhzCe       = "20/40mhz-Ce"
	InterfaceWiFiConfiguration_ChannelWidth_2040mhzEC       = "20/40mhz-eC"
	InterfaceWiFiConfiguration_ChannelWidth_20mhz           = "20mhz"
)

// Constants for country
const (
	InterfaceWiFiConfiguration_Country_Afghanistan                           = "Afghanistan"
	InterfaceWiFiConfiguration_Country_AlandIslands                          = "Aland Islands"
	InterfaceWiFiConfiguration_Country_Albania                               = "Albania"
	InterfaceWiFiConfiguration_Country_Algeria                               = "Algeria"
	InterfaceWiFiConfiguration_Country_AmericanSamoa                         = "American Samoa"
	InterfaceWiFiConfiguration_Country_Andora                                = "Andora"
	InterfaceWiFiConfiguration_Country_Anguilla                              = "Anguilla"
	InterfaceWiFiConfiguration_Country_AntiguaAndBarbuda                     = "Antigua and Barbuda"
	InterfaceWiFiConfiguration_Country_Argentina                             = "Argentina"
	InterfaceWiFiConfiguration_Country_Armenia                               = "Armenia"
	InterfaceWiFiConfiguration_Country_Aruba                                 = "Aruba"
	InterfaceWiFiConfiguration_Country_Australia                             = "Australia"
	InterfaceWiFiConfiguration_Country_Austria                               = "Austria"
	InterfaceWiFiConfiguration_Country_Azerbaijan                            = "Azerbaijan"
	InterfaceWiFiConfiguration_Country_Bahamas                               = "Bahamas"
	InterfaceWiFiConfiguration_Country_Bahrain                               = "Bahrain"
	InterfaceWiFiConfiguration_Country_Bangladesh                            = "Bangladesh"
	InterfaceWiFiConfiguration_Country_Barbados                              = "Barbados"
	InterfaceWiFiConfiguration_Country_Belarus                               = "Belarus"
	InterfaceWiFiConfiguration_Country_Belgium                               = "Belgium"
	InterfaceWiFiConfiguration_Country_Belize                                = "Belize"
	InterfaceWiFiConfiguration_Country_Bermuda                               = "Bermuda"
	InterfaceWiFiConfiguration_Country_Bhutan                                = "Bhutan"
	InterfaceWiFiConfiguration_Country_Bolivia                               = "Bolivia"
	InterfaceWiFiConfiguration_Country_BosniaAndHerzegovina                  = "Bosnia and Herzegovina"
	InterfaceWiFiConfiguration_Country_Brazil                                = "Brazil"
	InterfaceWiFiConfiguration_Country_BritishVirginIslands                  = "British Virgin Islands"
	InterfaceWiFiConfiguration_Country_BruneiDarussalam                      = "Brunei Darussalam"
	InterfaceWiFiConfiguration_Country_Bulgaria                              = "Bulgaria"
	InterfaceWiFiConfiguration_Country_BurkinaFaso                           = "Burkina Faso"
	InterfaceWiFiConfiguration_Country_Cambodia                              = "Cambodia"
	InterfaceWiFiConfiguration_Country_Cameroon                              = "Cameroon"
	InterfaceWiFiConfiguration_Country_Canada                                = "Canada"
	InterfaceWiFiConfiguration_Country_CaymanIslands                         = "Cayman Islands"
	InterfaceWiFiConfiguration_Country_CentralAfricaRepublic                 = "Central Africa Republic"
	InterfaceWiFiConfiguration_Country_Chad                                  = "Chad"
	InterfaceWiFiConfiguration_Country_Chile                                 = "Chile"
	InterfaceWiFiConfiguration_Country_China                                 = "China"
	InterfaceWiFiConfiguration_Country_ChristmasIsland                       = "Christmas Island"
	InterfaceWiFiConfiguration_Country_Colombia                              = "Colombia"
	InterfaceWiFiConfiguration_Country_CookIslands                           = "Cook Islands"
	InterfaceWiFiConfiguration_Country_CostaRica                             = "Costa Rica"
	InterfaceWiFiConfiguration_Country_Croatia                               = "Croatia"
	InterfaceWiFiConfiguration_Country_Cyprus                                = "Cyprus"
	InterfaceWiFiConfiguration_Country_Czech                                 = "Czech"
	InterfaceWiFiConfiguration_Country_Denmark                               = "Denmark"
	InterfaceWiFiConfiguration_Country_Dominica                              = "Dominica"
	InterfaceWiFiConfiguration_Country_DominicanRepublic                     = "Dominican Republic"
	InterfaceWiFiConfiguration_Country_ETSI5557Outdoor                       = "ETSI 5.5-5.7 Outdoor"
	InterfaceWiFiConfiguration_Country_Ecuador                               = "Ecuador"
	InterfaceWiFiConfiguration_Country_Egypt                                 = "Egypt"
	InterfaceWiFiConfiguration_Country_ElSalvador                            = "El Salvador"
	InterfaceWiFiConfiguration_Country_Estonia                               = "Estonia"
	InterfaceWiFiConfiguration_Country_Ethiopia                              = "Ethiopia"
	InterfaceWiFiConfiguration_Country_FalklandIslands                       = "Falkland Islands"
	InterfaceWiFiConfiguration_Country_FaroeIslands                          = "Faroe Islands"
	InterfaceWiFiConfiguration_Country_Finland                               = "Finland"
	InterfaceWiFiConfiguration_Country_France                                = "France"
	InterfaceWiFiConfiguration_Country_FrenchGuiana                          = "French Guiana"
	InterfaceWiFiConfiguration_Country_FrenchPolynesia                       = "French Polynesia"
	InterfaceWiFiConfiguration_Country_FrenchSouthernTerritories             = "French Southern Territories"
	InterfaceWiFiConfiguration_Country_Georgia                               = "Georgia"
	InterfaceWiFiConfiguration_Country_Germany                               = "Germany"
	InterfaceWiFiConfiguration_Country_Ghana                                 = "Ghana"
	InterfaceWiFiConfiguration_Country_Gibraltar                             = "Gibraltar"
	InterfaceWiFiConfiguration_Country_Greece                                = "Greece"
	InterfaceWiFiConfiguration_Country_Greenland                             = "Greenland"
	InterfaceWiFiConfiguration_Country_Grenada                               = "Grenada"
	InterfaceWiFiConfiguration_Country_Guadeloupe                            = "Guadeloupe"
	InterfaceWiFiConfiguration_Country_Guam                                  = "Guam"
	InterfaceWiFiConfiguration_Country_Guatemala                             = "Guatemala"
	InterfaceWiFiConfiguration_Country_Guernsey                              = "Guernsey"
	InterfaceWiFiConfiguration_Country_Guyana                                = "Guyana"
	InterfaceWiFiConfiguration_Country_Haiti                                 = "Haiti"
	InterfaceWiFiConfiguration_Country_HeardIslandAndMcDonaldIslands         = "Heard Island and McDonald Islands"
	InterfaceWiFiConfiguration_Country_HolySee                               = "Holy See"
	InterfaceWiFiConfiguration_Country_Honduras                              = "Honduras"
	InterfaceWiFiConfiguration_Country_HongKong                              = "Hong Kong"
	InterfaceWiFiConfiguration_Country_Hungary                               = "Hungary"
	InterfaceWiFiConfiguration_Country_Iceland                               = "Iceland"
	InterfaceWiFiConfiguration_Country_India                                 = "India"
	InterfaceWiFiConfiguration_Country_Indonesia                             = "Indonesia"
	InterfaceWiFiConfiguration_Country_Iraq                                  = "Iraq"
	InterfaceWiFiConfiguration_Country_Ireland                               = "Ireland"
	InterfaceWiFiConfiguration_Country_IsleOfMan                             = "Isle of Man"
	InterfaceWiFiConfiguration_Country_Israel                                = "Israel"
	InterfaceWiFiConfiguration_Country_Italy                                 = "Italy"
	InterfaceWiFiConfiguration_Country_IvoryCoast                            = "Ivory Coast"
	InterfaceWiFiConfiguration_Country_Jamaica                               = "Jamaica"
	InterfaceWiFiConfiguration_Country_Japan                                 = "Japan"
	InterfaceWiFiConfiguration_Country_Jersey                                = "Jersey"
	InterfaceWiFiConfiguration_Country_Jordan                                = "Jordan"
	InterfaceWiFiConfiguration_Country_Kazakhstan                            = "Kazakhstan"
	InterfaceWiFiConfiguration_Country_Kenya                                 = "Kenya"
	InterfaceWiFiConfiguration_Country_Kuwait                                = "Kuwait"
	InterfaceWiFiConfiguration_Country_Latvia                                = "Latvia"
	InterfaceWiFiConfiguration_Country_Lebanon                               = "Lebanon"
	InterfaceWiFiConfiguration_Country_Lesotho                               = "Lesotho"
	InterfaceWiFiConfiguration_Country_Liechtenstein                         = "Liechtenstein"
	InterfaceWiFiConfiguration_Country_Lithuania                             = "Lithuania"
	InterfaceWiFiConfiguration_Country_Luxembourg                            = "Luxembourg"
	InterfaceWiFiConfiguration_Country_Macau                                 = "Macau"
	InterfaceWiFiConfiguration_Country_Malawi                                = "Malawi"
	InterfaceWiFiConfiguration_Country_Malaysia                              = "Malaysia"
	InterfaceWiFiConfiguration_Country_Maldives                              = "Maldives"
	InterfaceWiFiConfiguration_Country_Malta                                 = "Malta"
	InterfaceWiFiConfiguration_Country_MarshallIslands                       = "Marshall Islands"
	InterfaceWiFiConfiguration_Country_Martinique                            = "Martinique"
	InterfaceWiFiConfiguration_Country_Mauritania                            = "Mauritania"
	InterfaceWiFiConfiguration_Country_Mauritius                             = "Mauritius"
	InterfaceWiFiConfiguration_Country_Mayotte                               = "Mayotte"
	InterfaceWiFiConfiguration_Country_Mexico                                = "Mexico"
	InterfaceWiFiConfiguration_Country_Micronesia                            = "Micronesia"
	InterfaceWiFiConfiguration_Country_Moldova                               = "Moldova"
	InterfaceWiFiConfiguration_Country_Monaco                                = "Monaco"
	InterfaceWiFiConfiguration_Country_Mongolia                              = "Mongolia"
	InterfaceWiFiConfiguration_Country_Montenegro                            = "Montenegro"
	InterfaceWiFiConfiguration_Country_Montserrat                            = "Montserrat"
	InterfaceWiFiConfiguration_Country_Morocco                               = "Morocco"
	InterfaceWiFiConfiguration_Country_Myanmar                               = "Myanmar"
	InterfaceWiFiConfiguration_Country_Namibia                               = "Namibia"
	InterfaceWiFiConfiguration_Country_Nepal                                 = "Nepal"
	InterfaceWiFiConfiguration_Country_Netherlands                           = "Netherlands"
	InterfaceWiFiConfiguration_Country_NetherlandsAntilles                   = "Netherlands Antilles"
	InterfaceWiFiConfiguration_Country_NewCaledonia                          = "New Caledonia"
	InterfaceWiFiConfiguration_Country_NewZealand                            = "New Zealand"
	InterfaceWiFiConfiguration_Country_Nicaragua                             = "Nicaragua"
	InterfaceWiFiConfiguration_Country_Nigeria                               = "Nigeria"
	InterfaceWiFiConfiguration_Country_Niue                                  = "Niue"
	InterfaceWiFiConfiguration_Country_NorfolkIsland                         = "Norfolk Island"
	InterfaceWiFiConfiguration_Country_NorthMacedonia                        = "North Macedonia"
	InterfaceWiFiConfiguration_Country_NorthernMarianaIslands                = "Northern Mariana Islands"
	InterfaceWiFiConfiguration_Country_Norway                                = "Norway"
	InterfaceWiFiConfiguration_Country_Oman                                  = "Oman"
	InterfaceWiFiConfiguration_Country_Pakistan                              = "Pakistan"
	InterfaceWiFiConfiguration_Country_Palau                                 = "Palau"
	InterfaceWiFiConfiguration_Country_Panama                                = "Panama"
	InterfaceWiFiConfiguration_Country_PapuaNewGuinea                        = "Papua New Guinea"
	InterfaceWiFiConfiguration_Country_Paraguay                              = "Paraguay"
	InterfaceWiFiConfiguration_Country_Peru                                  = "Peru"
	InterfaceWiFiConfiguration_Country_Philippines                           = "Philippines"
	InterfaceWiFiConfiguration_Country_Poland                                = "Poland"
	InterfaceWiFiConfiguration_Country_Portugal                              = "Portugal"
	InterfaceWiFiConfiguration_Country_PuertoRico                            = "Puerto Rico"
	InterfaceWiFiConfiguration_Country_Qatar                                 = "Qatar"
	InterfaceWiFiConfiguration_Country_Reunion                               = "Reunion"
	InterfaceWiFiConfiguration_Country_Romania                               = "Romania"
	InterfaceWiFiConfiguration_Country_Russia                                = "Russia"
	InterfaceWiFiConfiguration_Country_Rwanda                                = "Rwanda"
	InterfaceWiFiConfiguration_Country_SaintBarthelemy                       = "Saint Barthelemy"
	InterfaceWiFiConfiguration_Country_SaintHelenaAscensionAndTristanDaCunha = "Saint Helena, Ascension and Tristan da Cunha"
	InterfaceWiFiConfiguration_Country_SaintKittsAndNevis                    = "Saint Kitts and Nevis"
	InterfaceWiFiConfiguration_Country_SaintLucia                            = "Saint Lucia"
	InterfaceWiFiConfiguration_Country_SaintMartin                           = "Saint Martin"
	InterfaceWiFiConfiguration_Country_SaintPierreAndMiquelon                = "Saint Pierre and Miquelon"
	InterfaceWiFiConfiguration_Country_SaintVincentAndTheGrenadines          = "Saint Vincent and the Grenadines"
	InterfaceWiFiConfiguration_Country_Samoa                                 = "Samoa"
	InterfaceWiFiConfiguration_Country_SanMarino                             = "San Marino"
	InterfaceWiFiConfiguration_Country_SaoTomeAndPrincipe                    = "Sao Tome and Principe"
	InterfaceWiFiConfiguration_Country_SaudiArabia                           = "Saudi Arabia"
	InterfaceWiFiConfiguration_Country_Senegal                               = "Senegal"
	InterfaceWiFiConfiguration_Country_Serbia                                = "Serbia"
	InterfaceWiFiConfiguration_Country_Singapore                             = "Singapore"
	InterfaceWiFiConfiguration_Country_SintMaarten                           = "Sint Maarten"
	InterfaceWiFiConfiguration_Country_Slovakia                              = "Slovakia"
	InterfaceWiFiConfiguration_Country_Slovenia                              = "Slovenia"
	InterfaceWiFiConfiguration_Country_SouthAfrica                           = "South Africa"
	InterfaceWiFiConfiguration_Country_SouthKorea                            = "South Korea"
	InterfaceWiFiConfiguration_Country_Spain                                 = "Spain"
	InterfaceWiFiConfiguration_Country_SriLanka                              = "Sri Lanka"
	InterfaceWiFiConfiguration_Country_Superchannel                          = "Superchannel"
	InterfaceWiFiConfiguration_Country_Suriname                              = "Suriname"
	InterfaceWiFiConfiguration_Country_SvalbardAndJanMayen                   = "Svalbard and Jan Mayen"
	InterfaceWiFiConfiguration_Country_Sweden                                = "Sweden"
	InterfaceWiFiConfiguration_Country_Switzerland                           = "Switzerland"
	InterfaceWiFiConfiguration_Country_Taiwan                                = "Taiwan"
	InterfaceWiFiConfiguration_Country_Tanzania                              = "Tanzania"
	InterfaceWiFiConfiguration_Country_Thailand                              = "Thailand"
	InterfaceWiFiConfiguration_Country_Togo                                  = "Togo"
	InterfaceWiFiConfiguration_Country_TrinidadAndTobago                     = "Trinidad and Tobago"
	InterfaceWiFiConfiguration_Country_Tunisia                               = "Tunisia"
	InterfaceWiFiConfiguration_Country_Turkey                                = "Turkey"
	InterfaceWiFiConfiguration_Country_TurksAndCaicos                        = "Turks and Caicos"
	InterfaceWiFiConfiguration_Country_UK58Fixed                             = "UK 5.8 fixed"
	InterfaceWiFiConfiguration_Country_Uganda                                = "Uganda"
	InterfaceWiFiConfiguration_Country_Ukraine                               = "Ukraine"
	InterfaceWiFiConfiguration_Country_UnitedArabEmirates                    = "United Arab Emirates"
	InterfaceWiFiConfiguration_Country_UnitedKingdom                         = "United Kingdom"
	InterfaceWiFiConfiguration_Country_UnitedStates                          = "United States"
	InterfaceWiFiConfiguration_Country_UnitedStatesMinorOutlyingIslands      = "United States Minor Outlying Islands"
	InterfaceWiFiConfiguration_Country_Uruguay                               = "Uruguay"
	InterfaceWiFiConfiguration_Country_Uzbekistan                            = "Uzbekistan"
	InterfaceWiFiConfiguration_Country_Vanuatu                               = "Vanuatu"
	InterfaceWiFiConfiguration_Country_Venezuela                             = "Venezuela"
	InterfaceWiFiConfiguration_Country_VietNam                               = "Viet Nam"
	InterfaceWiFiConfiguration_Country_VirginIslands                         = "Virgin Islands"
	InterfaceWiFiConfiguration_Country_WallisAndFutuna                       = "Wallis and Futuna"
	InterfaceWiFiConfiguration_Country_Yemen                                 = "Yemen"
	InterfaceWiFiConfiguration_Country_Zimbabwe                              = "Zimbabwe"
)

// Constants for datapath.bridge
const (
	InterfaceWiFiConfiguration_DatapathBridge_None = "none"
)

// Constants for datapath.bridge-horizon
const (
	InterfaceWiFiConfiguration_DatapathBridgeHorizon_None = "none"
)

// Constants for datapath.interface-list
const (
	InterfaceWiFiConfiguration_DatapathInterfaceList_All     = "all"
	InterfaceWiFiConfiguration_DatapathInterfaceList_Dynamic = "dynamic"
	InterfaceWiFiConfiguration_DatapathInterfaceList_None    = "none"
	InterfaceWiFiConfiguration_DatapathInterfaceList_Static  = "static"
)

// Constants for datapath.traffic-processing
const (
	InterfaceWiFiConfiguration_DatapathTrafficProcessing_OnCap     = "on-cap"
	InterfaceWiFiConfiguration_DatapathTrafficProcessing_OnCapsman = "on-capsman"
)

// Constants for datapath.vlan-id
const (
	InterfaceWiFiConfiguration_DatapathVlanId_None = "none"
)

// Constants for installation
const (
	InterfaceWiFiConfiguration_Installation_Indoor  = "indoor"
	InterfaceWiFiConfiguration_Installation_Outdoor = "outdoor"
)

// Constants for interworking.authentication-types
const (
	InterfaceWiFiConfiguration_InterworkingAuthenticationTypes_DnsRedirection     = "dns-redirection"
	InterfaceWiFiConfiguration_InterworkingAuthenticationTypes_HttpsRedirection   = "https-redirection"
	InterfaceWiFiConfiguration_InterworkingAuthenticationTypes_OnlineEnrollment   = "online-enrollment"
	InterfaceWiFiConfiguration_InterworkingAuthenticationTypes_TermsAndConditions = "terms-and-conditions"
)

// Constants for interworking.ipv4-availability
const (
	InterfaceWiFiConfiguration_InterworkingIpv4Availability_DoubleNated               = "double-nated"
	InterfaceWiFiConfiguration_InterworkingIpv4Availability_NotAvailable              = "not-available"
	InterfaceWiFiConfiguration_InterworkingIpv4Availability_PortRestricted            = "port-restricted"
	InterfaceWiFiConfiguration_InterworkingIpv4Availability_PortRestrictedDoubleNated = "port-restricted-double-nated"
	InterfaceWiFiConfiguration_InterworkingIpv4Availability_PortRestrictedSingleNated = "port-restricted-single-nated"
	InterfaceWiFiConfiguration_InterworkingIpv4Availability_Public                    = "public"
	InterfaceWiFiConfiguration_InterworkingIpv4Availability_SingleNated               = "single-nated"
	InterfaceWiFiConfiguration_InterworkingIpv4Availability_Unknown                   = "unknown"
)

// Constants for interworking.ipv6-availability
const (
	InterfaceWiFiConfiguration_InterworkingIpv6Availability_Available    = "available"
	InterfaceWiFiConfiguration_InterworkingIpv6Availability_NotAvailable = "not-available"
	InterfaceWiFiConfiguration_InterworkingIpv6Availability_Unknown      = "unknown"
)

// Constants for interworking.network-type
const (
	InterfaceWiFiConfiguration_InterworkingNetworkType_EmergencyOnly    = "emergency-only"
	InterfaceWiFiConfiguration_InterworkingNetworkType_PersonalDevice   = "personal-device"
	InterfaceWiFiConfiguration_InterworkingNetworkType_Private          = "private"
	InterfaceWiFiConfiguration_InterworkingNetworkType_PrivateWithGuest = "private-with-guest"
	InterfaceWiFiConfiguration_InterworkingNetworkType_PublicChargeable = "public-chargeable"
	InterfaceWiFiConfiguration_InterworkingNetworkType_PublicFree       = "public-free"
	InterfaceWiFiConfiguration_InterworkingNetworkType_Test             = "test"
	InterfaceWiFiConfiguration_InterworkingNetworkType_Wildcard         = "wildcard"
)

// Constants for interworking.venue
const (
	InterfaceWiFiConfiguration_InterworkingVenue_AssemblyAmphitheater           = "assembly-amphitheater"
	InterfaceWiFiConfiguration_InterworkingVenue_AssemblyAmusementPark          = "assembly-amusement-park"
	InterfaceWiFiConfiguration_InterworkingVenue_AssemblyArena                  = "assembly-arena"
	InterfaceWiFiConfiguration_InterworkingVenue_AssemblyBar                    = "assembly-bar"
	InterfaceWiFiConfiguration_InterworkingVenue_AssemblyCoffeeShop             = "assembly-coffee-shop"
	InterfaceWiFiConfiguration_InterworkingVenue_AssemblyConventionCenter       = "assembly-convention-center"
	InterfaceWiFiConfiguration_InterworkingVenue_AssemblyEmergencyCc            = "assembly-emergency-cc"
	InterfaceWiFiConfiguration_InterworkingVenue_AssemblyLibrary                = "assembly-library"
	InterfaceWiFiConfiguration_InterworkingVenue_AssemblyMuseum                 = "assembly-museum"
	InterfaceWiFiConfiguration_InterworkingVenue_AssemblyPlaceOfWorship         = "assembly-place-of-worship"
	InterfaceWiFiConfiguration_InterworkingVenue_AssemblyRestaurant             = "assembly-restaurant"
	InterfaceWiFiConfiguration_InterworkingVenue_AssemblyStadium                = "assembly-stadium"
	InterfaceWiFiConfiguration_InterworkingVenue_AssemblyTerminal               = "assembly-terminal"
	InterfaceWiFiConfiguration_InterworkingVenue_AssemblyTheater                = "assembly-theater"
	InterfaceWiFiConfiguration_InterworkingVenue_AssemblyUnspecified            = "assembly-unspecified"
	InterfaceWiFiConfiguration_InterworkingVenue_AssemblyZoo                    = "assembly-zoo"
	InterfaceWiFiConfiguration_InterworkingVenue_BusinessAttorneyOffice         = "business-attorney-office"
	InterfaceWiFiConfiguration_InterworkingVenue_BusinessBank                   = "business-bank"
	InterfaceWiFiConfiguration_InterworkingVenue_BusinessDoctor                 = "business-doctor"
	InterfaceWiFiConfiguration_InterworkingVenue_BusinessFireStation            = "business-fire-station"
	InterfaceWiFiConfiguration_InterworkingVenue_BusinessPoliceStation          = "business-police-station"
	InterfaceWiFiConfiguration_InterworkingVenue_BusinessPostOffice             = "business-post-office"
	InterfaceWiFiConfiguration_InterworkingVenue_BusinessProfessionalOffice     = "business-professional-office"
	InterfaceWiFiConfiguration_InterworkingVenue_BusinessRdFacility             = "business-rd-facility"
	InterfaceWiFiConfiguration_InterworkingVenue_BusinessUnspecified            = "business-unspecified"
	InterfaceWiFiConfiguration_InterworkingVenue_Disabled                       = "disabled"
	InterfaceWiFiConfiguration_InterworkingVenue_EducationalPrimarySchool       = "educational-primary-school"
	InterfaceWiFiConfiguration_InterworkingVenue_EducationalSecondarySchool     = "educational-secondary-school"
	InterfaceWiFiConfiguration_InterworkingVenue_EducationalUniversityOrCollege = "educational-university-or-college"
	InterfaceWiFiConfiguration_InterworkingVenue_EducationalUnspecified         = "educational-unspecified"
	InterfaceWiFiConfiguration_InterworkingVenue_IndustrialFactory              = "industrial-factory"
	InterfaceWiFiConfiguration_InterworkingVenue_IndustrialUnspecified          = "industrial-unspecified"
	InterfaceWiFiConfiguration_InterworkingVenue_InstitutionalGroupHome         = "institutional-group-home"
	InterfaceWiFiConfiguration_InterworkingVenue_InstitutionalHospital          = "institutional-hospital"
	InterfaceWiFiConfiguration_InterworkingVenue_InstitutionalLongTermCare      = "institutional-long-term-care"
	InterfaceWiFiConfiguration_InterworkingVenue_InstitutionalPrison            = "institutional-prison"
	InterfaceWiFiConfiguration_InterworkingVenue_InstitutionalRehabCenter       = "institutional-rehab-center"
	InterfaceWiFiConfiguration_InterworkingVenue_InstitutionalUnspecified       = "institutional-unspecified"
	InterfaceWiFiConfiguration_InterworkingVenue_MercantileAutomotiveService    = "mercantile-automotive-service"
	InterfaceWiFiConfiguration_InterworkingVenue_MercantileGasStation           = "mercantile-gas-station"
	InterfaceWiFiConfiguration_InterworkingVenue_MercantileGroceryMarket        = "mercantile-grocery-market"
	InterfaceWiFiConfiguration_InterworkingVenue_MercantileRetailStore          = "mercantile-retail-store"
	InterfaceWiFiConfiguration_InterworkingVenue_MercantileShoppingMall         = "mercantile-shopping-mall"
	InterfaceWiFiConfiguration_InterworkingVenue_MercantileUnspecified          = "mercantile-unspecified"
	InterfaceWiFiConfiguration_InterworkingVenue_OutdoorBusStop                 = "outdoor-bus-stop"
	InterfaceWiFiConfiguration_InterworkingVenue_OutdoorCityPark                = "outdoor-city-park"
	InterfaceWiFiConfiguration_InterworkingVenue_OutdoorKiosk                   = "outdoor-kiosk"
	InterfaceWiFiConfiguration_InterworkingVenue_OutdoorMuniMesh                = "outdoor-muni-mesh"
	InterfaceWiFiConfiguration_InterworkingVenue_OutdoorRestArea                = "outdoor-rest-area"
	InterfaceWiFiConfiguration_InterworkingVenue_OutdoorTrafficControl          = "outdoor-traffic-control"
	InterfaceWiFiConfiguration_InterworkingVenue_OutdoorUnspecified             = "outdoor-unspecified"
	InterfaceWiFiConfiguration_InterworkingVenue_ResidentialBoardingHouse       = "residential-boarding-house"
	InterfaceWiFiConfiguration_InterworkingVenue_ResidentialDormitory           = "residential-dormitory"
	InterfaceWiFiConfiguration_InterworkingVenue_ResidentialHotel               = "residential-hotel"
	InterfaceWiFiConfiguration_InterworkingVenue_ResidentialPrivate             = "residential-private"
	InterfaceWiFiConfiguration_InterworkingVenue_ResidentialUnspecified         = "residential-unspecified"
	InterfaceWiFiConfiguration_InterworkingVenue_StorageUnspecified             = "storage-unspecified"
	InterfaceWiFiConfiguration_InterworkingVenue_Unspecified                    = "unspecified"
	InterfaceWiFiConfiguration_InterworkingVenue_UtilityUnspecified             = "utility-unspecified"
	InterfaceWiFiConfiguration_InterworkingVenue_VehicularAirplane              = "vehicular-airplane"
	InterfaceWiFiConfiguration_InterworkingVenue_VehicularAutomobile            = "vehicular-automobile"
	InterfaceWiFiConfiguration_InterworkingVenue_VehicularBus                   = "vehicular-bus"
	InterfaceWiFiConfiguration_InterworkingVenue_VehicularFerry                 = "vehicular-ferry"
	InterfaceWiFiConfiguration_InterworkingVenue_VehicularMotorbike             = "vehicular-motorbike"
	InterfaceWiFiConfiguration_InterworkingVenue_VehicularShip                  = "vehicular-ship"
	InterfaceWiFiConfiguration_InterworkingVenue_VehicularTrain                 = "vehicular-train"
	InterfaceWiFiConfiguration_InterworkingVenue_VehicularUnspecified           = "vehicular-unspecified"
)

// Constants for interworking.wan-status
const (
	InterfaceWiFiConfiguration_InterworkingWanStatus_Down     = "down"
	InterfaceWiFiConfiguration_InterworkingWanStatus_Reserved = "reserved"
	InterfaceWiFiConfiguration_InterworkingWanStatus_Test     = "test"
	InterfaceWiFiConfiguration_InterworkingWanStatus_Up       = "up"
)

// Constants for manager
const (
	InterfaceWiFiConfiguration_Manager_Capsman        = "capsman"
	InterfaceWiFiConfiguration_Manager_CapsmanOrLocal = "capsman-or-local"
	InterfaceWiFiConfiguration_Manager_Local          = "local"
)

// Constants for mode
const (
	InterfaceWiFiConfiguration_Mode_Ap                  = "ap"
	InterfaceWiFiConfiguration_Mode_Station             = "station"
	InterfaceWiFiConfiguration_Mode_StationBridge       = "station-bridge"
	InterfaceWiFiConfiguration_Mode_StationPseudobridge = "station-pseudobridge"
)

// Constants for multicast-enhance
const (
	InterfaceWiFiConfiguration_MulticastEnhance_Disabled = "disabled"
	InterfaceWiFiConfiguration_MulticastEnhance_Enabled  = "enabled"
)

// Constants for qos-classifier
const (
	InterfaceWiFiConfiguration_QosClassifier_DscpHigh3Bits = "dscp-high-3-bits"
	InterfaceWiFiConfiguration_QosClassifier_Priority      = "priority"
)

// Constants for security.authentication-types
const (
	InterfaceWiFiConfiguration_SecurityAuthenticationTypes_WpaPsk      = "wpa-psk"
	InterfaceWiFiConfiguration_SecurityAuthenticationTypes_Wpa2Psk     = "wpa2-psk"
	InterfaceWiFiConfiguration_SecurityAuthenticationTypes_Wpa2PskSha2 = "wpa2-psk-sha2"
	InterfaceWiFiConfiguration_SecurityAuthenticationTypes_WpaEap      = "wpa-eap"
	InterfaceWiFiConfiguration_SecurityAuthenticationTypes_Wpa2Eap     = "wpa2-eap"
	InterfaceWiFiConfiguration_SecurityAuthenticationTypes_Wpa3Psk     = "wpa3-psk"
	InterfaceWiFiConfiguration_SecurityAuthenticationTypes_Owe         = "owe"
	InterfaceWiFiConfiguration_SecurityAuthenticationTypes_Wpa3Eap     = "wpa3-eap"
	InterfaceWiFiConfiguration_SecurityAuthenticationTypes_Wpa3Eap192  = "wpa3-eap-192"
)

// Constants for security.dh-groups
const (
	InterfaceWiFiConfiguration_SecurityDhGroups_19 = "19"
	InterfaceWiFiConfiguration_SecurityDhGroups_20 = "20"
	InterfaceWiFiConfiguration_SecurityDhGroups_21 = "21"
)

// Constants for security.eap-certificate-mode
const (
	InterfaceWiFiConfiguration_SecurityEapCertificateMode_DontVerifyCertificate    = "dont-verify-certificate"
	InterfaceWiFiConfiguration_SecurityEapCertificateMode_NoCertificates           = "no-certificates"
	InterfaceWiFiConfiguration_SecurityEapCertificateMode_VerifyCertificate        = "verify-certificate"
	InterfaceWiFiConfiguration_SecurityEapCertificateMode_VerifyCertificateWithCrl = "verify-certificate-with-crl"
)

// Constants for security.eap-methods
const (
	InterfaceWiFiConfiguration_SecurityEapMethods_Peap = "peap"
	InterfaceWiFiConfiguration_SecurityEapMethods_Tls  = "tls"
	InterfaceWiFiConfiguration_SecurityEapMethods_Ttls = "ttls"
)

// Constants for security.encryption
const (
	InterfaceWiFiConfiguration_SecurityEncryption_Tkip    = "tkip"
	InterfaceWiFiConfiguration_SecurityEncryption_Ccmp    = "ccmp"
	InterfaceWiFiConfiguration_SecurityEncryption_Gcmp    = "gcmp"
	InterfaceWiFiConfiguration_SecurityEncryption_Ccmp256 = "ccmp-256"
	InterfaceWiFiConfiguration_SecurityEncryption_Gcmp256 = "gcmp-256"
)

// Constants for security.group-encryption
const (
	InterfaceWiFiConfiguration_SecurityGroupEncryption_Ccmp    = "ccmp"
	InterfaceWiFiConfiguration_SecurityGroupEncryption_Ccmp256 = "ccmp-256"
	InterfaceWiFiConfiguration_SecurityGroupEncryption_Gcmp    = "gcmp"
	InterfaceWiFiConfiguration_SecurityGroupEncryption_Gcmp256 = "gcmp-256"
	InterfaceWiFiConfiguration_SecurityGroupEncryption_Tkip    = "tkip"
)

// Constants for security.management-encryption
const (
	InterfaceWiFiConfiguration_SecurityManagementEncryption_Cmac    = "cmac"
	InterfaceWiFiConfiguration_SecurityManagementEncryption_Cmac256 = "cmac-256"
	InterfaceWiFiConfiguration_SecurityManagementEncryption_Gmac    = "gmac"
	InterfaceWiFiConfiguration_SecurityManagementEncryption_Gmac256 = "gmac-256"
)

// Constants for security.management-protection
const (
	InterfaceWiFiConfiguration_SecurityManagementProtection_Allowed  = "allowed"
	InterfaceWiFiConfiguration_SecurityManagementProtection_Disabled = "disabled"
	InterfaceWiFiConfiguration_SecurityManagementProtection_Required = "required"
)

// Constants for security.owe-transition-interface
const (
	InterfaceWiFiConfiguration_SecurityOweTransitionInterface_Auto = "auto"
)

// Constants for security.sae-anti-clogging-threshold
const (
	InterfaceWiFiConfiguration_SecuritySaeAntiCloggingThreshold_Disabled = "disabled"
)

// Constants for security.sae-max-failure-rate
const (
	InterfaceWiFiConfiguration_SecuritySaeMaxFailureRate_Disabled = "disabled"
)

// Constants for security.sae-pwe
const (
	InterfaceWiFiConfiguration_SecuritySaePwe_Both              = "both"
	InterfaceWiFiConfiguration_SecuritySaePwe_HashToElement     = "hash-to-element"
	InterfaceWiFiConfiguration_SecuritySaePwe_HuntingAndPecking = "hunting-and-pecking"
)

// Constants for security.wps
const (
	InterfaceWiFiConfiguration_SecurityWps_Disable    = "disable"
	InterfaceWiFiConfiguration_SecurityWps_PushButton = "push-button"
)

// Constants for tx-chains
const (
	InterfaceWiFiConfiguration_TxChains_0 = "0"
	InterfaceWiFiConfiguration_TxChains_1 = "1"
	InterfaceWiFiConfiguration_TxChains_2 = "2"
	InterfaceWiFiConfiguration_TxChains_3 = "3"
	InterfaceWiFiConfiguration_TxChains_4 = "4"
	InterfaceWiFiConfiguration_TxChains_5 = "5"
	InterfaceWiFiConfiguration_TxChains_6 = "6"
	InterfaceWiFiConfiguration_TxChains_7 = "7"
)

// InterfaceWiFiConfiguration defines resource
type InterfaceWiFiConfiguration struct {
	Id                                 string `mikrotik:".id" codegen:"id"`
	Aaa                                string `mikrotik:"aaa" codegen:"aaa"`
	AaaCalledFormat                    string `mikrotik:"aaa.called-format" codegen:"aaa.called_format"`
	AaaCallingFormat                   string `mikrotik:"aaa.calling-format" codegen:"aaa.calling_format"`
	AaaInterimUpdate                   string `mikrotik:"aaa.interim-update" codegen:"aaa.interim_update"`
	AaaMacCaching                      string `mikrotik:"aaa.mac-caching" codegen:"aaa.mac_caching"`
	AaaNasIdentifier                   string `mikrotik:"aaa.nas-identifier" codegen:"aaa.nas_identifier"`
	AaaPasswordFormat                  string `mikrotik:"aaa.password-format" codegen:"aaa.password_format"`
	AaaUsernameFormat                  string `mikrotik:"aaa.username-format" codegen:"aaa.username_format"`
	AntennaGain                        string `mikrotik:"antenna-gain" codegen:"antenna_gain"`
	BeaconInterval                     string `mikrotik:"beacon-interval" codegen:"beacon_interval"`
	Chains                             string `mikrotik:"chains" codegen:"chains"`
	Channel                            string `mikrotik:"channel" codegen:"channel"`
	ChannelBand                        string `mikrotik:"channel.band" codegen:"channel.band"`
	ChannelDeprioritizeUnii34          bool   `mikrotik:"channel.deprioritize-unii-3-4" codegen:"channel.deprioritize_unii_3_4"`
	ChannelFrequency                   string `mikrotik:"channel.frequency" codegen:"channel.frequency"`
	ChannelReselectInterval            string `mikrotik:"channel.reselect-interval" codegen:"channel.reselect_interval"`
	ChannelReselectTime                string `mikrotik:"channel.reselect-time" codegen:"channel.reselect_time"`
	ChannelSecondaryFrequency          string `mikrotik:"channel.secondary-frequency" codegen:"channel.secondary_frequency"`
	ChannelSkipDfsChannels             string `mikrotik:"channel.skip-dfs-channels" codegen:"channel.skip_dfs_channels"`
	ChannelWidth                       string `mikrotik:"channel.width" codegen:"channel.width"`
	Comment                            string `mikrotik:"comment" codegen:"comment"`
	CopyFrom                           string `mikrotik:"copy-from" codegen:"copy_from"`
	Country                            string `mikrotik:"country" codegen:"country"`
	Datapath                           string `mikrotik:"datapath" codegen:"datapath"`
	DatapathBridge                     string `mikrotik:"datapath.bridge" codegen:"datapath.bridge"`
	DatapathBridgeCost                 string `mikrotik:"datapath.bridge-cost" codegen:"datapath.bridge_cost"`
	DatapathBridgeHorizon              string `mikrotik:"datapath.bridge-horizon" codegen:"datapath.bridge_horizon"`
	DatapathClientIsolation            bool   `mikrotik:"datapath.client-isolation" codegen:"datapath.client_isolation"`
	DatapathInterfaceList              string `mikrotik:"datapath.interface-list" codegen:"datapath.interface_list"`
	DatapathTrafficProcessing          string `mikrotik:"datapath.traffic-processing" codegen:"datapath.traffic_processing"`
	DatapathVlanId                     string `mikrotik:"datapath.vlan-id" codegen:"datapath.vlan_id"`
	Disabled                           bool   `mikrotik:"disabled" codegen:"disabled"`
	Distance                           string `mikrotik:"distance" codegen:"distance"`
	DtimPeriod                         string `mikrotik:"dtim-period" codegen:"dtim_period"`
	HideSsid                           bool   `mikrotik:"hide-ssid" codegen:"hide_ssid"`
	Installation                       string `mikrotik:"installation" codegen:"installation"`
	Interworking                       string `mikrotik:"interworking" codegen:"interworking"`
	Interworking3gppInfo               string `mikrotik:"interworking.3gpp-info" codegen:"interworking.3gpp_info"`
	InterworkingAuthenticationTypes    string `mikrotik:"interworking.authentication-types" codegen:"interworking.authentication_types"`
	InterworkingConnectionCapabilities string `mikrotik:"interworking.connection-capabilities" codegen:"interworking.connection_capabilities"`
	InterworkingDomainNames            string `mikrotik:"interworking.domain-names" codegen:"interworking.domain_names"`
	InterworkingEsr                    bool   `mikrotik:"interworking.esr" codegen:"interworking.esr"`
	InterworkingHessid                 string `mikrotik:"interworking.hessid" codegen:"interworking.hessid"`
	InterworkingHotspot20              bool   `mikrotik:"interworking.hotspot20" codegen:"interworking.hotspot20"`
	InterworkingHotspot20Dgaf          bool   `mikrotik:"interworking.hotspot20-dgaf" codegen:"interworking.hotspot20_dgaf"`
	InterworkingInternet               bool   `mikrotik:"interworking.internet" codegen:"interworking.internet"`
	InterworkingIpv4Availability       string `mikrotik:"interworking.ipv4-availability" codegen:"interworking.ipv4_availability"`
	InterworkingIpv6Availability       string `mikrotik:"interworking.ipv6-availability" codegen:"interworking.ipv6_availability"`
	InterworkingNetworkType            string `mikrotik:"interworking.network-type" codegen:"interworking.network_type"`
	InterworkingOperationalClasses     string `mikrotik:"interworking.operational-classes" codegen:"interworking.operational_classes"`
	InterworkingOperatorNames          string `mikrotik:"interworking.operator-names" codegen:"interworking.operator_names"`
	InterworkingRealms                 string `mikrotik:"interworking.realms" codegen:"interworking.realms"`
	InterworkingRoamingOis             string `mikrotik:"interworking.roaming-ois" codegen:"interworking.roaming_ois"`
	InterworkingUesa                   bool   `mikrotik:"interworking.uesa" codegen:"interworking.uesa"`
	InterworkingVenue                  string `mikrotik:"interworking.venue" codegen:"interworking.venue"`
	InterworkingVenueNames             string `mikrotik:"interworking.venue-names" codegen:"interworking.venue_names"`
	InterworkingWanAtCapacity          bool   `mikrotik:"interworking.wan-at-capacity" codegen:"interworking.wan_at_capacity"`
	InterworkingWanDownlink            string `mikrotik:"interworking.wan-downlink" codegen:"interworking.wan_downlink"`
	InterworkingWanDownlinkLoad        string `mikrotik:"interworking.wan-downlink-load" codegen:"interworking.wan_downlink_load"`
	InterworkingWanMeasurementDuration string `mikrotik:"interworking.wan-measurement-duration" codegen:"interworking.wan_measurement_duration"`
	InterworkingWanStatus              string `mikrotik:"interworking.wan-status" codegen:"interworking.wan_status"`
	InterworkingWanSymmetric           bool   `mikrotik:"interworking.wan-symmetric" codegen:"interworking.wan_symmetric"`
	InterworkingWanUplink              string `mikrotik:"interworking.wan-uplink" codegen:"interworking.wan_uplink"`
	InterworkingWanUplinkLoad          string `mikrotik:"interworking.wan-uplink-load" codegen:"interworking.wan_uplink_load"`
	Manager                            string `mikrotik:"manager" codegen:"manager"`
	MaxClients                         string `mikrotik:"max-clients" codegen:"max_clients"`
	Mode                               string `mikrotik:"mode" codegen:"mode"`
	MulticastEnhance                   string `mikrotik:"multicast-enhance" codegen:"multicast_enhance"`
	Name                               string `mikrotik:"name" codegen:"name"`
	QosClassifier                      string `mikrotik:"qos-classifier" codegen:"qos_classifier"`
	Security                           string `mikrotik:"security" codegen:"security"`
	SecurityAuthenticationTypes        string `mikrotik:"security.authentication-types" codegen:"security.authentication_types"`
	SecurityConnectGroup               string `mikrotik:"security.connect-group" codegen:"security.connect_group"`
	SecurityConnectPriority            string `mikrotik:"security.connect-priority" codegen:"security.connect_priority"`
	SecurityDhGroups                   string `mikrotik:"security.dh-groups" codegen:"security.dh_groups"`
	SecurityDisablePmkid               bool   `mikrotik:"security.disable-pmkid" codegen:"security.disable_pmkid"`
	SecurityEapAccounting              bool   `mikrotik:"security.eap-accounting" codegen:"security.eap_accounting"`
	SecurityEapAnonymousIdentity       string `mikrotik:"security.eap-anonymous-identity" codegen:"security.eap_anonymous_identity"`
	SecurityEapCertificateMode         string `mikrotik:"security.eap-certificate-mode" codegen:"security.eap_certificate_mode"`
	SecurityEapMethods                 string `mikrotik:"security.eap-methods" codegen:"security.eap_methods"`
	SecurityEapPassword                string `mikrotik:"security.eap-password" codegen:"security.eap_password"`
	SecurityEapTlsCertificate          string `mikrotik:"security.eap-tls-certificate" codegen:"security.eap_tls_certificate"`
	SecurityEapUsername                string `mikrotik:"security.eap-username" codegen:"security.eap_username"`
	SecurityEncryption                 string `mikrotik:"security.encryption" codegen:"security.encryption"`
	SecurityFt                         bool   `mikrotik:"security.ft" codegen:"security.ft"`
	SecurityFtMobilityDomain           string `mikrotik:"security.ft-mobility-domain" codegen:"security.ft_mobility_domain"`
	SecurityFtNasIdentifier            string `mikrotik:"security.ft-nas-identifier" codegen:"security.ft_nas_identifier"`
	SecurityFtOverDs                   bool   `mikrotik:"security.ft-over-ds" codegen:"security.ft_over_ds"`
	SecurityFtPreserveVlanid           bool   `mikrotik:"security.ft-preserve-vlanid" codegen:"security.ft_preserve_vlanid"`
	SecurityFtR0KeyLifetime            string `mikrotik:"security.ft-r0-key-lifetime" codegen:"security.ft_r0_key_lifetime"`
	SecurityFtReassociationDeadline    string `mikrotik:"security.ft-reassociation-deadline" codegen:"security.ft_reassociation_deadline"`
	SecurityGroupEncryption            string `mikrotik:"security.group-encryption" codegen:"security.group_encryption"`
	SecurityGroupKeyUpdate             string `mikrotik:"security.group-key-update" codegen:"security.group_key_update"`
	SecurityManagementEncryption       string `mikrotik:"security.management-encryption" codegen:"security.management_encryption"`
	SecurityManagementProtection       string `mikrotik:"security.management-protection" codegen:"security.management_protection"`
	SecurityMultiPassphraseGroup       string `mikrotik:"security.multi-passphrase-group" codegen:"security.multi_passphrase_group"`
	SecurityOweTransitionInterface     string `mikrotik:"security.owe-transition-interface" codegen:"security.owe_transition_interface"`
	SecurityPassphrase                 string `mikrotik:"security.passphrase" codegen:"security.passphrase"`
	SecuritySaeAntiCloggingThreshold   string `mikrotik:"security.sae-anti-clogging-threshold" codegen:"security.sae_anti_clogging_threshold"`
	SecuritySaeMaxFailureRate          string `mikrotik:"security.sae-max-failure-rate" codegen:"security.sae_max_failure_rate"`
	SecuritySaePwe                     string `mikrotik:"security.sae-pwe" codegen:"security.sae_pwe"`
	SecurityWps                        string `mikrotik:"security.wps" codegen:"security.wps"`
	Ssid                               string `mikrotik:"ssid" codegen:"ssid"`
	StationRoaming                     bool   `mikrotik:"station-roaming" codegen:"station_roaming"`
	Steering                           string `mikrotik:"steering" codegen:"steering"`
	Steering2gProbeDelay               bool   `mikrotik:"steering.2g-probe-delay" codegen:"steering.2g_probe_delay"`
	SteeringNeighborGroup              string `mikrotik:"steering.neighbor-group" codegen:"steering.neighbor_group"`
	SteeringRrm                        bool   `mikrotik:"steering.rrm" codegen:"steering.rrm"`
	SteeringWnm                        bool   `mikrotik:"steering.wnm" codegen:"steering.wnm"`
	TxChains                           string `mikrotik:"tx-chains" codegen:"tx_chains"`
	TxPower                            string `mikrotik:"tx-power" codegen:"tx_power"`

	About string `mikrotik:"about,readonly" codegen:"about,readonly"`
}

var _ Resource = (*InterfaceWiFiConfiguration)(nil)

func (b *InterfaceWiFiConfiguration) ActionToCommand(a Action) string {
	return map[Action]string{
		Add:    "/interface/wifi/configuration/add",
		Find:   "/interface/wifi/configuration/print",
		Update: "/interface/wifi/configuration/set",
		Delete: "/interface/wifi/configuration/remove",
	}[a]
}

func (b *InterfaceWiFiConfiguration) IDField() string {
	return ".id"
}

func (b *InterfaceWiFiConfiguration) ID() string {
	return b.Id
}

func (b *InterfaceWiFiConfiguration) SetID(id string) {
	b.Id = id
}

// Uncomment extra methods to satisfy more interfaces

// Adder
func (b *InterfaceWiFiConfiguration) AfterAddHook(r *routeros.Reply) {
	b.Id = r.Done.Map["ret"]
}

// Finder
// func (b *InterfaceWiFiConfiguration) FindField() string {
// 	return "name"
// }

// func (b *InterfaceWiFiConfiguration) FindFieldValue() string {
// 	return b.Name
// }

// Deleter
// func (b *InterfaceWiFiConfiguration) DeleteField() string {
// 	return "numbers"
// }

// func (b *InterfaceWiFiConfiguration) DeleteFieldValue() string {
// 	return b.Id
// }

// Typed wrappers
func (c Mikrotik) AddInterfaceWiFiConfiguration(r *InterfaceWiFiConfiguration) (*InterfaceWiFiConfiguration, error) {
	res, err := c.Add(r)
	if err != nil {
		return nil, err
	}

	return res.(*InterfaceWiFiConfiguration), nil
}

func (c Mikrotik) UpdateInterfaceWiFiConfiguration(r *InterfaceWiFiConfiguration) (*InterfaceWiFiConfiguration, error) {
	res, err := c.Update(r)
	if err != nil {
		return nil, err
	}

	return res.(*InterfaceWiFiConfiguration), nil
}

func (c Mikrotik) FindInterfaceWiFiConfiguration(id string) (*InterfaceWiFiConfiguration, error) {
	res, err := c.Find(&InterfaceWiFiConfiguration{Id: id})
	if err != nil {
		return nil, err
	}

	return res.(*InterfaceWiFiConfiguration), nil
}

func (c Mikrotik) ListInterfaceWiFiConfiguration() ([]InterfaceWiFiConfiguration, error) {
	res, err := c.List(&InterfaceWiFiConfiguration{})
	if err != nil {
		return nil, err
	}
	returnSlice := make([]InterfaceWiFiConfiguration, len(res))
	for i, v := range res {
		returnSlice[i] = *(v.(*InterfaceWiFiConfiguration))
	}

	return returnSlice, nil
}

func (c Mikrotik) DeleteInterfaceWiFiConfiguration(id string) error {
	return c.Delete(&InterfaceWiFiConfiguration{Id: id})
}
