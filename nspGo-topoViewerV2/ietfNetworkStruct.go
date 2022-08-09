package nspgotopoviewernew

type IetfNetworkStruct struct {
	Response struct {
		Status    int `json:"status"`
		StartRow  int `json:"startRow"`
		EndRow    int `json:"endRow"`
		TotalRows int `json:"totalRows"`
		Data      struct {
			TeTopologiesAugment interface{} `json:"teTopologiesAugment"`
			Network             []struct {
				NetworkID         string `json:"networkId"`
				ServerProvided    bool   `json:"serverProvided"`
				TeTopologyAugment struct {
					Te struct {
						ProviderID struct {
							Uint32 int `json:"uint32"`
						} `json:"providerId"`
						ClientID     interface{} `json:"clientId"`
						TeTopologyID struct {
							String string `json:"string"`
						} `json:"teTopologyId"`
						Config interface{} `json:"config"`
						State  struct {
							TeTopologyConfig struct {
								Schedules             interface{} `json:"schedules"`
								Preference            int         `json:"preference"`
								OptimizationCriterion struct {
								} `json:"optimizationCriterion"`
							} `json:"teTopologyConfig"`
						} `json:"state"`
					} `json:"te"`
				} `json:"teTopologyAugment"`
				NetworkTypes struct {
					TeTopologyType struct {
						TeTopology struct {
						} `json:"teTopology"`
					} `json:"teTopologyType"`
				} `json:"networkTypes"`
				SupportingNetwork interface{} `json:"supportingNetwork"`
				Node              []struct {
					NodeID        string `json:"nodeId"`
					TeNodeAugment struct {
						Te struct {
							TeNodeID struct {
								DottedQuad struct {
									String string `json:"string"`
								} `json:"dottedQuad"`
							} `json:"teNodeId"`
							Config interface{} `json:"config"`
							State  struct {
								TeNodeConfig       interface{} `json:"teNodeConfig"`
								TeNodeStateDerived struct {
									InformationSource      interface{} `json:"informationSource"`
									InformationSourceState interface{} `json:"informationSourceState"`
									OperStatus             string      `json:"operStatus"`
									IsMultiAccessDr        interface{} `json:"isMultiAccessDr"`
									InformationSourceEntry interface{} `json:"informationSourceEntry"`
								} `json:"teNodeStateDerived"`
							} `json:"state"`
							TunnelTerminationPoint interface{} `json:"tunnelTerminationPoint"`
						} `json:"te"`
					} `json:"teNodeAugment"`
					SupportingNode   interface{} `json:"supportingNode"`
					TerminationPoint []struct {
						TpID                      string `json:"tpId"`
						TeTerminationPointAugment struct {
							Te struct {
								TeTpID struct {
									IPAddress struct {
										Ipv6Address interface{} `json:"ipv6Address"`
										Ipv4Address struct {
											String string `json:"string"`
										} `json:"ipv4Address"`
									} `json:"ipAddress"`
									Uint32 interface{} `json:"uint32"`
								} `json:"teTpId"`
								Config interface{} `json:"config"`
								State  struct {
									TeTerminationPointConfig struct {
										Schedules                        interface{} `json:"schedules"`
										InterfaceSwitchingCapabilityList struct {
											InterfaceSwitchingCapability struct {
												SwitchingPsc1 struct {
													TimeDivisionMultiplexCapable interface{} `json:"timeDivisionMultiplexCapable"`
													MaxLspBandwidth              interface{} `json:"maxLspBandwidth"`
													SwitchingCapability          struct {
													} `json:"switchingCapability"`
													Encoding interface{} `json:"encoding"`
												} `json:"SwitchingPsc1"`
											} `json:"interfaceSwitchingCapability"`
										} `json:"interfaceSwitchingCapabilityList"`
										InterLayerLockID interface{} `json:"interLayerLockId"`
									} `json:"teTerminationPointConfig"`
								} `json:"state"`
							} `json:"te"`
						} `json:"teTerminationPointAugment"`
						SupportingTerminationPoint interface{} `json:"supportingTerminationPoint"`
					} `json:"terminationPoint"`
				} `json:"node"`
				Link []struct {
					LinkID        string `json:"linkId"`
					TeLinkAugment struct {
						Te struct {
							Config struct {
								TeLinkConfig struct {
									TeLinkAttributes struct {
										Schedules                          interface{} `json:"schedules"`
										AccessType                         interface{} `json:"accessType"`
										ExternalDomain                     interface{} `json:"externalDomain"`
										IsAbstract                         interface{} `json:"isAbstract"`
										Name                               interface{} `json:"name"`
										Underlay                           interface{} `json:"underlay"`
										AdminStatus                        interface{} `json:"adminStatus"`
										PerformanceMetricThrottleContainer interface{} `json:"performanceMetricThrottleContainer"`
										TeLinkInfoAttributes               struct {
											InterfaceSwitchingCapabilityList interface{} `json:"interfaceSwitchingCapabilityList"`
											TeLinkConnectivityAttributes     struct {
												MaxLinkBandwidth     interface{} `json:"maxLinkBandwidth"`
												MaxResvLinkBandwidth interface{} `json:"maxResvLinkBandwidth"`
												UnreservedBandwidth  interface{} `json:"unreservedBandwidth"`
												TeDefaultMetric      interface{} `json:"teDefaultMetric"`
												PerformanceMetric    struct {
													TePerformanceMetric interface{} `json:"tePerformanceMetric"`
													Measurement         struct {
														PerformanceMetricAttributes struct {
															UnidirectionalDelay              int         `json:"unidirectionalDelay"`
															UnidirectionalMinDelay           interface{} `json:"unidirectionalMinDelay"`
															UnidirectionalMaxDelay           interface{} `json:"unidirectionalMaxDelay"`
															UnidirectionalDelayVariation     interface{} `json:"unidirectionalDelayVariation"`
															UnidirectionalPacketLoss         interface{} `json:"unidirectionalPacketLoss"`
															UnidirectionalResidualBandwidth  interface{} `json:"unidirectionalResidualBandwidth"`
															UnidirectionalAvailableBandwidth interface{} `json:"unidirectionalAvailableBandwidth"`
															UnidirectionalUtilizedBandwidth  interface{} `json:"unidirectionalUtilizedBandwidth"`
														} `json:"performanceMetricAttributes"`
													} `json:"measurement"`
													Normality interface{} `json:"normality"`
												} `json:"performanceMetric"`
												TeSrlgs struct {
													Value []interface{} `json:"value"`
												} `json:"teSrlgs"`
											} `json:"teLinkConnectivityAttributes"`
											LinkIndex           interface{} `json:"linkIndex"`
											AdministrativeGroup interface{} `json:"administrativeGroup"`
											LinkProtectionType  interface{} `json:"linkProtectionType"`
										} `json:"teLinkInfoAttributes"`
									} `json:"teLinkAttributes"`
									BundledLinks   interface{} `json:"bundledLinks"`
									ComponentLinks interface{} `json:"componentLinks"`
									TeLinkTemplate interface{} `json:"teLinkTemplate"`
									Template       interface{} `json:"template"`
								} `json:"teLinkConfig"`
							} `json:"config"`
							State struct {
								TeLinkConfig struct {
									TeLinkAttributes struct {
										Schedules                          interface{} `json:"schedules"`
										AccessType                         string      `json:"accessType"`
										ExternalDomain                     interface{} `json:"externalDomain"`
										IsAbstract                         bool        `json:"isAbstract"`
										Name                               string      `json:"name"`
										Underlay                           interface{} `json:"underlay"`
										AdminStatus                        string      `json:"adminStatus"`
										PerformanceMetricThrottleContainer interface{} `json:"performanceMetricThrottleContainer"`
										TeLinkInfoAttributes               struct {
											InterfaceSwitchingCapabilityList interface{} `json:"interfaceSwitchingCapabilityList"`
											TeLinkConnectivityAttributes     struct {
												MaxLinkBandwidth     interface{} `json:"maxLinkBandwidth"`
												MaxResvLinkBandwidth interface{} `json:"maxResvLinkBandwidth"`
												UnreservedBandwidth  struct {
													Num0 interface{} `json:"0"`
												} `json:"unreservedBandwidth"`
												TeDefaultMetric   int         `json:"teDefaultMetric"`
												PerformanceMetric interface{} `json:"performanceMetric"`
												TeSrlgs           struct {
													Value []interface{} `json:"value"`
												} `json:"teSrlgs"`
											} `json:"teLinkConnectivityAttributes"`
											LinkIndex           interface{} `json:"linkIndex"`
											AdministrativeGroup struct {
												AdminGroup struct {
													Binary []int `json:"binary"`
												} `json:"adminGroup"`
												ExtendedAdminGroup interface{} `json:"extendedAdminGroup"`
											} `json:"administrativeGroup"`
											LinkProtectionType interface{} `json:"linkProtectionType"`
										} `json:"teLinkInfoAttributes"`
									} `json:"teLinkAttributes"`
									BundledLinks   interface{} `json:"bundledLinks"`
									ComponentLinks interface{} `json:"componentLinks"`
									TeLinkTemplate interface{} `json:"teLinkTemplate"`
									Template       interface{} `json:"template"`
								} `json:"teLinkConfig"`
								TeLinkStateDerived struct {
									InformationSource      interface{} `json:"informationSource"`
									InformationSourceState interface{} `json:"informationSourceState"`
									OperStatus             string      `json:"operStatus"`
									IsTransitional         interface{} `json:"isTransitional"`
									InformationSourceEntry interface{} `json:"informationSourceEntry"`
									Recovery               interface{} `json:"recovery"`
									Underlay               interface{} `json:"underlay"`
								} `json:"teLinkStateDerived"`
							} `json:"state"`
						} `json:"te"`
					} `json:"teLinkAugment"`
					Source struct {
						SourceNode string `json:"sourceNode"`
						SourceTp   string `json:"sourceTp"`
					} `json:"source"`
					Destination struct {
						DestNode string `json:"destNode"`
						DestTp   string `json:"destTp"`
					} `json:"destination"`
					SupportingLink interface{} `json:"supportingLink"`
				} `json:"link"`
			} `json:"network"`
		} `json:"data"`
	} `json:"response"`
}
